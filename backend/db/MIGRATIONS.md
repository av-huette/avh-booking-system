# Database migrations

The database schema is managed with [goose](https://github.com/pressly/goose); the migrations are SQL files in `db/migrations/`.
Run all commands from `backend/`; they read the database settings from `backend/.env` (set up as in the "Development" quick start of the root `README.md`).

## Commands

```bash
go run ./cmd/migrate status    # applied and pending migrations
go run ./cmd/migrate up        # apply all pending migrations
go run ./cmd/migrate version   # current schema version
```

There are no down migrations: an applied migration can't be undone, so check it before running `up` on a database whose data you want to keep.

## Changing the schema

`001_initial_schema.sql` holds the whole schema.
While AVHBS is not in production, schema changes go into `001`, and local databases are dropped and re-created with `migrate up`.

Once a database has to be kept, every change becomes a new file `NNN_description.sql` with only a `-- +goose Up` section, for example `002_add_account_note.sql`:

```sql
-- +goose Up
ALTER TABLE account ADD COLUMN note VARCHAR(100);
```

We fix forward: a mistake is corrected by the next migration.

## Tests

The integration tests in `internal/repo` build their schema by running the migrations, then insert `internal/repo/testdata/insert_test_data.sql`.
A migration that fails therefore fails `go test ./...`.
If a migration ever inserts data, the test data may collide with it.
