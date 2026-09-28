# SKATEPARK API GO

This is my skateparks api written in golang. It powers [https://www.celebrityskateboards.com](https://www.celebrityskateboards.com).

## REPOSITORY

These are the skateparks that I've been to in my life. 

### How to Add

1. Add entry to appropriate state folder
2. Run locally (with load imgur albums enabled)
3. Commit imgur imaage update to json files
4. How part of the repository

## ADMIN PORTAL (IN PROGRESS)

The service now supports an internal `/admin` route that can be protected by GitHub OAuth.

Set these environment variables to enable admin auth:

- `GITHUB_OAUTH_CLIENT_ID`
- `GITHUB_OAUTH_CLIENT_SECRET`
- `GITHUB_OAUTH_REDIRECT_URL` (example: `https://your-domain.com/admin/callback`)
- `ADMIN_SESSION_SECRET` (minimum 32 chars)
- `ADMIN_ALLOWED_GITHUB_USER` (defaults to `derekpedersen`)
- `ADMIN_COOKIE_SECURE` (`true` in production HTTPS)
- `ADMIN_POST_LOGIN_REDIRECT` (default `/admin/me`)
- `ADMIN_POST_LOGOUT_REDIRECT` (default `/admin/me`)
- `ADMIN_ALLOWED_REDIRECT_URI` (optional exact external SPA URL allowed via `redirect_uri` query parameter)

If these values are not set, `/admin` auth routes stay disabled and the API continues to run normally.

## BACKUP LEGACY DB JSON

Before migrating off filesystem JSON, archive `.db` into `.old-db`:

```bash
make backup-db
```

This creates a timestamped backup zip file like `.old-db/skatepark-db-YYYYMMDD-HHMMSS.zip`.