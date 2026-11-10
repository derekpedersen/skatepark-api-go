# Copilot Instructions for skatepark-api-go

Follow these rules when proposing or making changes in this repository.

## Scope and Style

- Make the smallest practical change that satisfies the request.
- Follow existing Go code patterns already used in `src/`.
- Avoid introducing new architectural patterns unless requested.
- Keep naming and file organization consistent with nearby code.

## Required Validation

- Run `make test` after code changes.
- Run `make build` when edits affect compile paths, startup wiring, or cross-file interfaces.
- If a command cannot be run, clearly state what was not executed and why.

## High-Risk Areas Requiring Targeted Tests

When touching these areas, update or add tests in related `src/*_test.go` files:

- Admin auth and OAuth flow (`admin.go`, `handler.go`)
- Routing behavior and endpoint wiring (`router.go`, handlers)
- Data access and model/domain contracts (`db.go`, `model.go`, `domain.go`)

## Safety and Consistency

- Do not change CI/CD, Docker, Helm, or Kubernetes files unless requested.
- Do not alter environment variable contracts without updating docs and tests.
- Keep docs aligned with real commands in `Makefile` and observed behavior.

## Source of Truth

- Use `AGENTS.md` as the canonical repo workflow and contribution contract.
- Use this file as the repo-scoped Copilot behavior layer.
