# AVHBS Backend Conventions

Conventions for working in `backend/` (Go, PostgreSQL via pgx/v5, gorilla/websocket), for developers and coding agents.
Cross-cutting rules are in the root `AGENTS.md`; the contribution process is in `CONTRIBUTING.md`.

## Commands (run from `backend/`)

```bash
go run ./cmd/server           # Run the server
go run ./cmd/migrate status   # Show applied and pending migrations
go run ./cmd/migrate up       # Apply pending migrations (ask first: there are no down migrations)
go test ./...                 # Run the integration tests (need a database)
```

## Checks

Before suggesting a commit: `gofmt -l .` (must print nothing), `go vet ./...`, `go test ./...`.
The tests are integration tests and need a running Postgres configured in `internal/repo/.env`; without one, report them as skipped.

## Environment

The server reads `.env` (template: `.env.template`); the integration tests read `internal/repo/.env`.
Variables:

- `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`
- `AVHBS_HTTP_PORT` (default 8081), `AVHBS_LOG_LEVEL` (`DEBUG`, `INFO`, `WARN`, `ERROR`), `AVHBS_LOG_FORMAT` (`text` or `json`), `AVHBS_FRONTEND_PATH`

## Where code goes

| New …                 | Goes into                                                                                 |
|-----------------------|-------------------------------------------------------------------------------------------|
| Domain type           | `internal/models/<entity>.go`: struct and constructor; no DB access, no protocol          |
| Store                 | `internal/repo/<entity>.go` (`<Entity>Store{DB DBTx}`); table constant in `repo/query.go` |
| Store interface       | `internal/ws/stores.go`, plus a field in `Stores`                                         |
| Wiring                | `cmd/server/main.go`                                                                      |
| Config / env variable | `internal/config/config.go`, plus `.env.template`                                         |
| Message type          | `internal/ws/message.go`, `ReadPump` in `internal/ws/pump.go`, `schemas/`                 |
| Schema change         | `db/migrations/` (Up only; see `db/MIGRATIONS.md`)                                        |
| Test                  | `internal/repo/integration_test.go` (`package repo_test`)                                 |

A new table needs all of these, in this order:

1. Migration in `db/migrations/`; the integration tests build their schema from it.
2. Model, store and `TableName` constant.
3. Interface and `Stores` field; wiring in `cmd/server/main.go`.
4. Cases in `processQuery`, `processMutation` and `prepareBroadcast` (`internal/ws/processor.go`); composite-key tables use the `…NoReturn` helpers.
   Without `prepareBroadcast`, other clients never see changes.
5. The table name in `schemas/db/_tableNames.json`; without it, the schema rejects every query.
6. The frontend side (`frontend/AGENTS.md`).

## Rules

### Structure

- Wire everything by hand in `cmd/server/main.go`: no DI framework, no `init()`, no package-level state.
- Define interfaces where they are used (`internal/ws/stores.go`), with only the methods the consumer calls.

### Context

- Functions that do I/O take `ctx context.Context` first; use `context.Background()` only in `main()`, tests and one-off setup.
- Never use `r.Context()` for WS connections: it is canceled when `ServeHTTP` returns.
  Use the per-connection `Client.Ctx`, the only ctx stored in a struct.

### Errors and logging

- Log an error once, where it is finally handled (e.g. `sendError()`); never log and return it.
- A new `WSErrorCode` caused by bad client input goes into the `Warn` case of `logLevel()` (`internal/ws/error.go`); server-side failures fall through to `Error`.
- Log with `slog` only (before the logger exists, `main()` writes to stderr).
  Use structured fields with `snake_case` keys, never formatted strings; never log passwords, tokens or personal data.

### WebSocket

- Only `WritePump` writes to a connection; send through `Client.Send`.
- Build outgoing messages with the helpers in `internal/ws/conversion.go`; they validate against `schemas/message.json`.

### Repo layer and SQL

- No ORM (why: `docs/adr/0002-postgresql-with-pgx-and-raw-sql.md`).
  Values go into SQL only as `$N` placeholders, identifiers only via `pgx.Identifier{...}.Sanitize()` (see `buildSelectSQL`); never build SQL by string interpolation.
- Never import `database/sql`; map `pgx.ErrNoRows` to `database.ErrNoRecord`, the only "not found" error.
- `GetByID` delegates to `Get()` with a filter on the primary key.
- Nullable columns map to pointer types (`*int`), never to `COALESCE`.
- Model fields carry `json` tags in camelCase and `db` tags with the column name.
- Accounts are never deleted (disable them via `enabled`), and `AccountStore.Update` never changes `Balance`: every booking and every balance change must stay traceable for fiscalization (`docs/adr/0004-traceable-bookings.md`).

### Database changes

- Every schema change gets a goose migration with only a `-- +goose Up` section.
  While nothing is in production, it goes into `001_initial_schema.sql` and local databases are re-created (`db/MIGRATIONS.md`).
  There are no down migrations: we fix forward with new migrations.

### Testing

- Tests use testify and a real Postgres; no mocks, no mocking libraries, no testcontainers.
  Don't assume a database in CI.
