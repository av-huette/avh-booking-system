# Architecture decision records

An architecture decision record (ADR) explains why the software is built the way it is: the context, the decision,
the options that were considered and the consequences. The rules that follow from a decision live in the
`AGENTS.md` files; the ADR keeps the reasons.

## Index

| ADR | Title | Status |
|---|---|---|
| [0001](0001-websocket-only-transport.md) | WebSocket-only transport | Accepted |
| [0002](0002-postgresql-with-pgx-and-raw-sql.md) | PostgreSQL with pgx and raw SQL, no ORM | Accepted |
| [0003](0003-docker-compose-with-caddy.md) | Docker Compose with Caddy | Accepted |
| [0004](0004-traceable-bookings.md) | Traceable bookings | Accepted |
| [0005](0005-money-as-integer-cents.md) | Money as integer cents | Accepted |
| [0006](0006-internationalization.md) | Internationalization | Accepted |

## When to write an ADR

Write an ADR when a decision shapes how code is written or structured, and it is either hard to reverse or likely
to be questioned by someone new to the project. A reason that fits into half a sentence stays as a clause in the
`AGENTS.md` rule instead.

## How to write one

- Copy [`template.md`](template.md) to `NNNN-kebab-case-title.md`, using the next free number, and add it to the
  index above.
- Status is `Proposed`, `Accepted`, or `Superseded by NNNN`. An accepted ADR is not edited afterwards; a changed
  decision gets a new ADR that supersedes the old one.
- A decision made before it was recorded gets an ADR too; its context says that it documents an existing decision.
