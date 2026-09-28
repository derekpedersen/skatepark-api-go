package skatepark_api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestAdmin_NewAdminServer_Validation(t *testing.T) {
	_, err := NewAdminServer(AdminConfig{})
	assert.Error(t, err)

	_, err = NewAdminServer(AdminConfig{
		GitHubClientID:     "id",
		GitHubClientSecret: "secret",
		GitHubRedirectURL:  "http://localhost/admin/callback",
		SessionSecret:      "short",
		AllowedGitHubUser:  "derekpedersen",
	})
	assert.Error(t, err)
}

func TestAdmin_UnauthenticatedMeReturns401(t *testing.T) {
	admin := newTestAdminServer(t)

	router := mux.NewRouter()
	AddAdminRoutes(router, admin)

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	var payload map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &payload)
	assert.NoError(t, err)
	assert.Equal(t, false, payload["authenticated"])
}

func TestAdmin_LoginRedirectsToGitHub(t *testing.T) {
	admin := newTestAdminServer(t)

	router := mux.NewRouter()
	AddAdminRoutes(router, admin)

	req := httptest.NewRequest(http.MethodGet, "/admin/login", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusFound, rr.Code)
	assert.Contains(t, rr.Header().Get("Location"), "github.com/login/oauth/authorize")
	assert.NotEmpty(t, rr.Result().Cookies())
}

func TestAdmin_CallbackRejectsInvalidState(t *testing.T) {
	admin := newTestAdminServer(t)

	router := mux.NewRouter()
	AddAdminRoutes(router, admin)

	req := httptest.NewRequest(http.MethodGet, "/admin/callback?state=bad&code=abc", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestAdmin_CallbackRejectsNonAllowedUser(t *testing.T) {
	admin := newTestAdminServer(t)
	admin.exchangeToken = func(_ context.Context, _ string) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "token"}, nil
	}
	admin.fetchGitHub = func(_ context.Context, _ *oauth2.Token) (string, error) {
		return "someone-else", nil
	}

	router := mux.NewRouter()
	AddAdminRoutes(router, admin)

	state, loginCookies := performLogin(t, router, "")

	req := httptest.NewRequest(http.MethodGet, "/admin/callback?state="+url.QueryEscape(state)+"&code=abc", nil)
	for _, c := range loginCookies {
		req.AddCookie(c)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestAdmin_SuccessfulLoginAndAdminAccess(t *testing.T) {
	admin := newTestAdminServer(t)
	admin.exchangeToken = func(_ context.Context, _ string) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "token"}, nil
	}
	admin.fetchGitHub = func(_ context.Context, _ *oauth2.Token) (string, error) {
		return "derekpedersen", nil
	}

	router := mux.NewRouter()
	AddAdminRoutes(router, admin)

	state, loginCookies := performLogin(t, router, "")

	callbackReq := httptest.NewRequest(http.MethodGet, "/admin/callback?state="+url.QueryEscape(state)+"&code=abc", nil)
	for _, c := range loginCookies {
		callbackReq.AddCookie(c)
	}
	callbackRR := httptest.NewRecorder()
	router.ServeHTTP(callbackRR, callbackReq)

	assert.Equal(t, http.StatusFound, callbackRR.Code)
	assert.Equal(t, "/admin/me", callbackRR.Header().Get("Location"))

	adminReq := httptest.NewRequest(http.MethodGet, "/admin/me", nil)
	for _, c := range callbackRR.Result().Cookies() {
		adminReq.AddCookie(c)
	}
	adminRR := httptest.NewRecorder()
	router.ServeHTTP(adminRR, adminReq)

	assert.Equal(t, http.StatusOK, adminRR.Code)
	var payload map[string]interface{}
	err := json.Unmarshal(adminRR.Body.Bytes(), &payload)
	assert.NoError(t, err)
	assert.Equal(t, true, payload["authenticated"])
	assert.Equal(t, "derekpedersen", payload["login"])
}

func TestAdmin_LoginAndLogoutUseRedirectURI(t *testing.T) {
	admin := newTestAdminServer(t)
	admin.allowedRedirectURI = "http://localhost:4200/admin"
	admin.postLogoutRedirect = "http://localhost:4200/login"
	admin.exchangeToken = func(_ context.Context, _ string) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "token"}, nil
	}
	admin.fetchGitHub = func(_ context.Context, _ *oauth2.Token) (string, error) {
		return "derekpedersen", nil
	}

	router := mux.NewRouter()
	AddAdminRoutes(router, admin)

	state, loginCookies := performLogin(t, router, "http://localhost:4200/admin")

	callbackReq := httptest.NewRequest(http.MethodGet, "/admin/callback?state="+url.QueryEscape(state)+"&code=abc", nil)
	for _, c := range loginCookies {
		callbackReq.AddCookie(c)
	}
	callbackRR := httptest.NewRecorder()
	router.ServeHTTP(callbackRR, callbackReq)

	assert.Equal(t, http.StatusFound, callbackRR.Code)
	assert.Equal(t, "http://localhost:4200/admin", callbackRR.Header().Get("Location"))

	logoutReq := httptest.NewRequest(http.MethodGet, "/admin/logout?redirect_uri=http://localhost:4200/login", nil)
	for _, c := range callbackRR.Result().Cookies() {
		logoutReq.AddCookie(c)
	}
	logoutRR := httptest.NewRecorder()
	router.ServeHTTP(logoutRR, logoutReq)

	assert.Equal(t, http.StatusFound, logoutRR.Code)
	assert.Equal(t, admin.postLogoutRedirect, logoutRR.Header().Get("Location"))
}

func newTestAdminServer(t *testing.T) *AdminServer {
	t.Helper()

	admin, err := NewAdminServer(AdminConfig{
		GitHubClientID:     "client-id",
		GitHubClientSecret: "client-secret",
		GitHubRedirectURL:  "http://localhost:8080/admin/callback",
		SessionSecret:      strings.Repeat("a", 32),
		AllowedGitHubUser:  "derekpedersen",
		CookieSecure:       false,
	})
	if err != nil {
		t.Fatalf("expected test admin server: %v", err)
	}

	return admin
}

func performLogin(t *testing.T, router *mux.Router, redirectURI string) (string, []*http.Cookie) {
	t.Helper()

	path := "/admin/login"
	if redirectURI != "" {
		path += "?redirect_uri=" + url.QueryEscape(redirectURI)
	}

	req := httptest.NewRequest(http.MethodGet, path, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("expected login redirect, got %d", rr.Code)
	}

	location := rr.Header().Get("Location")
	parsed, err := url.Parse(location)
	if err != nil {
		t.Fatalf("failed to parse login redirect location: %v", err)
	}

	state := parsed.Query().Get("state")
	if state == "" {
		t.Fatalf("state was not set in oauth redirect")
	}

	return state, rr.Result().Cookies()
}