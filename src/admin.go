package skatepark_api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	defaultAllowedGitHubUser = "derekpedersen"
	defaultAdminSessionName  = "skatepark-admin"
	oauthStateSessionKey     = "oauth_state"
	oauthRedirectSessionKey  = "oauth_redirect"
	authenticatedSessionKey  = "authenticated"
	githubLoginSessionKey    = "github_login"
)

type tokenExchanger func(context.Context, string) (*oauth2.Token, error)
type githubUserFetcher func(context.Context, *oauth2.Token) (string, error)

// AdminConfig controls /admin authentication behavior.
type AdminConfig struct {
	GitHubClientID     string
	GitHubClientSecret string
	GitHubRedirectURL  string
	SessionSecret      string
	AllowedGitHubUser  string
	PostLoginRedirect  string
	PostLogoutRedirect string
	AllowedRedirectURI string
	CookieSecure       bool
}

// Enabled reports if all required OAuth/session settings are present.
func (c AdminConfig) Enabled() bool {
	return c.GitHubClientID != "" &&
		c.GitHubClientSecret != "" &&
		c.GitHubRedirectURL != "" &&
		c.SessionSecret != ""
}

// NewAdminConfigFromEnv loads admin configuration from environment variables.
func NewAdminConfigFromEnv() AdminConfig {
	allowedUser := strings.TrimSpace(os.Getenv("ADMIN_ALLOWED_GITHUB_USER"))
	if allowedUser == "" {
		allowedUser = defaultAllowedGitHubUser
	}

	return AdminConfig{
		GitHubClientID:     strings.TrimSpace(os.Getenv("GITHUB_OAUTH_CLIENT_ID")),
		GitHubClientSecret: strings.TrimSpace(os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")),
		GitHubRedirectURL:  strings.TrimSpace(os.Getenv("GITHUB_OAUTH_REDIRECT_URL")),
		SessionSecret:      os.Getenv("ADMIN_SESSION_SECRET"),
		AllowedGitHubUser:  allowedUser,
		PostLoginRedirect:  strings.TrimSpace(os.Getenv("ADMIN_POST_LOGIN_REDIRECT")),
		PostLogoutRedirect: strings.TrimSpace(os.Getenv("ADMIN_POST_LOGOUT_REDIRECT")),
		AllowedRedirectURI: strings.TrimSpace(os.Getenv("ADMIN_ALLOWED_REDIRECT_URI")),
		CookieSecure:       parseEnvBool(os.Getenv("ADMIN_COOKIE_SECURE")),
	}
}

// AdminServer handles /admin OAuth flow and protected portal pages.
type AdminServer struct {
	oauthConfig *oauth2.Config
	store       *sessions.CookieStore

	allowedGitHubUser string
	postLoginRedirect string
	postLogoutRedirect string
	allowedRedirectURI string
	sessionName       string

	exchangeToken tokenExchanger
	fetchGitHub   githubUserFetcher
}

// NewAdminServer creates a new admin auth server using GitHub OAuth.
func NewAdminServer(cfg AdminConfig) (*AdminServer, error) {
	if !cfg.Enabled() {
		return nil, errors.New("admin config is incomplete")
	}
	if len(cfg.SessionSecret) < 32 {
		return nil, errors.New("ADMIN_SESSION_SECRET must be at least 32 characters")
	}

	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 8,
	}

	oauthConfig := &oauth2.Config{
		ClientID:     cfg.GitHubClientID,
		ClientSecret: cfg.GitHubClientSecret,
		RedirectURL:  cfg.GitHubRedirectURL,
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
	}

	admin := &AdminServer{
		oauthConfig:       oauthConfig,
		store:             store,
		allowedGitHubUser: cfg.AllowedGitHubUser,
		postLoginRedirect: cfg.PostLoginRedirect,
		postLogoutRedirect: cfg.PostLogoutRedirect,
		allowedRedirectURI: cfg.AllowedRedirectURI,
		sessionName:       defaultAdminSessionName,
	}
	if admin.postLoginRedirect == "" {
		admin.postLoginRedirect = "/admin/me"
	}
	if admin.postLogoutRedirect == "" {
		admin.postLogoutRedirect = "/admin/me"
	}

	admin.exchangeToken = func(ctx context.Context, code string) (*oauth2.Token, error) {
		return admin.oauthConfig.Exchange(ctx, code)
	}
	admin.fetchGitHub = admin.fetchGitHubLogin

	return admin, nil
}

// AddAdminRoutes mounts /admin handlers if an admin server is configured.
func AddAdminRoutes(router *mux.Router, admin *AdminServer) {
	if admin == nil {
		return
	}

	adminRouter := router.StrictSlash(true).PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("", admin.meHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/", admin.meHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/me", admin.meHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/login", admin.loginHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/callback", admin.callbackHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/logout", admin.logoutHandler).Methods(http.MethodGet, http.MethodPost)
}

func (a *AdminServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, a.sessionName)
	if err != nil {
		http.Error(w, "failed to initialize session", http.StatusInternalServerError)
		return
	}

	state, err := generateState()
	if err != nil {
		http.Error(w, "failed to initialize oauth state", http.StatusInternalServerError)
		return
	}

	session.Values[oauthStateSessionKey] = state
	session.Values[oauthRedirectSessionKey] = a.redirectTarget(r.URL.Query().Get("redirect_uri"), a.postLoginRedirect)
	if err := session.Save(r, w); err != nil {
		http.Error(w, "failed to store oauth state", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, a.oauthConfig.AuthCodeURL(state), http.StatusFound)
}

func (a *AdminServer) callbackHandler(w http.ResponseWriter, r *http.Request) {
	state := strings.TrimSpace(r.URL.Query().Get("state"))
	code := strings.TrimSpace(r.URL.Query().Get("code"))
	if state == "" || code == "" {
		http.Error(w, "missing oauth state or code", http.StatusBadRequest)
		return
	}

	session, err := a.store.Get(r, a.sessionName)
	if err != nil {
		http.Error(w, "failed to load session", http.StatusInternalServerError)
		return
	}

	savedState, _ := session.Values[oauthStateSessionKey].(string)
	if savedState == "" || savedState != state {
		http.Error(w, "invalid oauth state", http.StatusUnauthorized)
		return
	}

	token, err := a.exchangeToken(r.Context(), code)
	if err != nil {
		http.Error(w, "failed to exchange oauth code", http.StatusUnauthorized)
		return
	}

	login, err := a.fetchGitHub(r.Context(), token)
	if err != nil {
		http.Error(w, "failed to fetch github profile", http.StatusUnauthorized)
		return
	}

	if !strings.EqualFold(login, a.allowedGitHubUser) {
		a.clearSession(session)
		_ = session.Save(r, w)
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	session.Values[authenticatedSessionKey] = true
	session.Values[githubLoginSessionKey] = login
	redirectTarget, _ := session.Values[oauthRedirectSessionKey].(string)
	delete(session.Values, oauthStateSessionKey)
	delete(session.Values, oauthRedirectSessionKey)

	if err := session.Save(r, w); err != nil {
		http.Error(w, "failed to save session", http.StatusInternalServerError)
		return
	}
	if redirectTarget == "" {
		redirectTarget = a.postLoginRedirect
	}

	http.Redirect(w, r, redirectTarget, http.StatusFound)
}

func (a *AdminServer) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, a.sessionName)
	if err != nil {
		http.Error(w, "failed to load session", http.StatusInternalServerError)
		return
	}

	a.clearSession(session)
	if err := session.Save(r, w); err != nil {
		http.Error(w, "failed to clear session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, a.redirectTarget(r.URL.Query().Get("redirect_uri"), a.postLogoutRedirect), http.StatusFound)
}

func (a *AdminServer) requireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := a.store.Get(r, a.sessionName)
		if err != nil {
			http.Error(w, "failed to load session", http.StatusInternalServerError)
			return
		}

		authenticated, _ := session.Values[authenticatedSessionKey].(bool)
		login, _ := session.Values[githubLoginSessionKey].(string)
		if !authenticated || !strings.EqualFold(login, a.allowedGitHubUser) {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (a *AdminServer) meHandler(w http.ResponseWriter, r *http.Request) {
	login, ok, err := a.currentUser(r)
	if err != nil {
		http.Error(w, "failed to load session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"authenticated": false,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"authenticated": true,
		"login":         login,
	})
}

func (a *AdminServer) fetchGitHubLogin(ctx context.Context, token *oauth2.Token) (string, error) {
	client := a.oauthConfig.Client(ctx, token)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return "", fmt.Errorf("github user request failed: status=%d body=%s", resp.StatusCode, string(body))
	}

	var profile struct {
		Login string `json:"login"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return "", err
	}
	if strings.TrimSpace(profile.Login) == "" {
		return "", errors.New("github profile missing login")
	}

	return profile.Login, nil
}

func (a *AdminServer) clearSession(session *sessions.Session) {
	delete(session.Values, oauthStateSessionKey)
	delete(session.Values, oauthRedirectSessionKey)
	delete(session.Values, authenticatedSessionKey)
	delete(session.Values, githubLoginSessionKey)
	session.Options.MaxAge = -1
}

func (a *AdminServer) currentUser(r *http.Request) (string, bool, error) {
	session, err := a.store.Get(r, a.sessionName)
	if err != nil {
		return "", false, err
	}

	authenticated, _ := session.Values[authenticatedSessionKey].(bool)
	login, _ := session.Values[githubLoginSessionKey].(string)
	if !authenticated || !strings.EqualFold(login, a.allowedGitHubUser) {
		return "", false, nil
	}

	return login, true, nil
}

func (a *AdminServer) redirectTarget(candidate string, fallback string) string {
	target := strings.TrimSpace(candidate)
	if target == "" {
		return fallback
	}

	if strings.HasPrefix(target, "/") {
		return target
	}

	if a.allowedRedirectURI != "" && strings.EqualFold(target, a.allowedRedirectURI) {
		return target
	}

	return fallback
}

func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func parseEnvBool(value string) bool {
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return parsed
}