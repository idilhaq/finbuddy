# Changelog

## [Unreleased]

## [23-05-2025]
### Added
- Implemented filter and grouping support on `/expenses` (`?date`, `?month`, `?category`, `group_by=date`)
- Added `/plans` endpoints:
  - `POST /plans` to create or update monthly budget
  - `GET /plans/:month` to retrieve a specific plan
- Introduced `/dashboard` endpoint with summarized financial data
- Defined `Saving` model with UUID primary key
- Created Swagger annotations for `/plans` and `/dashboard`
- Defined `ErrorResponse` struct for consistent API error docs

### Changed
- Migrated all `user_id` fields from `int` to `uuid.UUID`
- Updated dashboard aggregation queries to use `COALESCE(SUM(...), 0)`
- Switched `swag init` to target `cmd/api/main.go` for correct doc generation

### Fixed
- Resolved Swagger error caused by `gin.H` in `@Failure` annotations
- Fixed SQL scan error on null SUM results in dashboard using `COALESCE` fallback

## [22-05-2025]
### Added
- Integrated `air` hot-reload for development
- Created `.air.toml` with build output to `tmp/main`
- Defined `docker-compose.override.yml` for dev mode with source code volume
- Extended `Makefile` with `dev-up`, `dev-down`, `dev-shell`, `lint`, and `fmt` targets

### Changed
- Updated dev Docker stage to use `golang:1.24-alpine`
- Switched `air` module path to `github.com/air-verse/air` to resolve Go module conflict

### Fixed
- Disabled Go VCS stamping via `-buildvcs=false` to fix exit status 128
- Fixed missing binary error (`tmp/main: not found`) during `air` run
- Confirmed hot reload works and API accessible at `localhost:3000`

---

## [21-05-2025]
### Added
- Created structured folder layout: `cmd/api`, `internal/db`, `internal/router`, `internal/handler`
- Implemented Gin router with `/` and `/healthz` endpoints
- Set up PostgreSQL 17 container with healthcheck
- Connected to DB using GORM and defined initial `Expense` model
- Added automatic DB migration on startup
- Created Dockerfile with multi-stage build (builder + runtime)
- Created `Makefile` with build, test, run, docker-up, docker-down, clean

### Fixed
- Resolved container exit issue by fixing database user and connection string
- Verified full system: browser accessible routes + containerized DB

---

## [20-05-2025]
### Added
- Initialized Go module and main server file
- Created `.env` with DB config
- Set up Docker container with PostgreSQL
- Verified DB connection from Go backend
- Added `/healthz` endpoint to test HTTP server
- Created `.gitignore` for Go, Docker, and env
- Initialized Git repo and pushed to GitHub
