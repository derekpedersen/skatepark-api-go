# AGENTS

This file is the contributor and AI-agent contract for this repository.

## Purpose

Skatepark API Go is a Go HTTP API service. Primary application code is in `src/`, with the executable entrypoint in `cmd/main.go`.

## Repository Layout

- `cmd/main.go`: application bootstrap and startup wiring.
- `src/`: HTTP handlers, router, domain logic, data access, and tests.
- `Makefile`: canonical local and CI workflow commands.
- `Jenkinsfile`: CI pipeline stages and order.

## Canonical Workflows

Use Make targets whenever possible instead of ad-hoc command variants.

- Install and sync dependencies: `make dependencies`
- Run tests with coverage artifacts: `make test`
- Build binary to `.bin/skatepark-api-go`: `make build`
- Run the built binary: `make run`
- Build image: `make docker`
- Publish image: `make publish`
- Deploy with Helm: `make deploy`

CI currently runs `make dependencies`, `make test`, and `make build`.

## Change Rules

- Keep diffs small and scoped to the request.
- Do not refactor unrelated files while addressing a task.
- Preserve public behavior unless the task explicitly requests a change.
- Do not modify deploy or infrastructure behavior unless requested.
- Prefer existing patterns in `src/*.go` and corresponding tests.

## Testing Expectations

- Add or update tests when behavior changes.
- For handler/router/auth changes, update related tests in `src/*_test.go`.
- For data/domain changes, update domain/model/db tests as needed.
- Before opening a PR, run at least `make test`; run `make build` when code paths affect compilation or startup.

## Documentation Expectations

- If behavior or configuration changes, update relevant docs in `README.md`.
- Keep documentation concise and consistent with actual Make targets and runtime behavior.

## Performance and Environment

- Assume constrained build/runtime resources unless told otherwise.
- Prefer efficient tooling and avoid unnecessary heavy operations in routine development tasks.

## Copilot Integration

Repository-scoped Copilot guidance lives in `.github/copilot-instructions.md`.
Use that file for coding-agent behavior details and this file for canonical repository workflow context.
