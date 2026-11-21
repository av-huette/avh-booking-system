# 0002. PostgreSQL with pgx and raw SQL, no ORM

- Status: Accepted
- Date: 2025-07-27 (recorded 2026-11-21)

## Context

Versions 1 and 2 used MySQL through Go's `database/sql`, with hand-written SQL. This ADR documents a decision made
when version 3 was started.

## Decision

AVHBS uses PostgreSQL, accessed through pgx's native API. SQL is written by hand, with a small query builder for
filtered selects (`buildSelectSQL` in `backend/internal/repo/query.go`). No ORM is used, and `database/sql` is not
used.

## Considered options

- **MySQL, as in versions 1 and 2:** replaced by PostgreSQL because of its popularity.
- **An ORM such as GORM:** not chosen. An ORM hides what is executed, and Go's standard library already provides a
  good data layer.
- **`database/sql` with the pgx driver:** not chosen. `database/sql` is a generic interface for any SQL database,
  so it can't expose PostgreSQL-specific behavior; pgx's native API gives PostgreSQL's features and performance
  without workarounds.
- **PostgreSQL with pgx's native API and hand-written SQL:** chosen.

## Consequences

- Every query is written by hand, so what runs against the database is always visible in the code.
- SQL safety is the code's responsibility: values only as `$N` placeholders, identifiers only through
  `pgx.Identifier`.
- `pgx.ErrNoRows` is mapped to the project's own "not found" error, since `database/sql`'s errors don't exist here.
