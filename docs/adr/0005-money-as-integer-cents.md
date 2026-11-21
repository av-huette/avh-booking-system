# 0005. Money as integer cents

- Status: Accepted
- Date: 2025-12-11 (recorded 2026-11-21)

## Context

Version 2 stored money as floating-point numbers (`float32`).
Version 3 first stored balances as PostgreSQL `NUMERIC`, which was sent over the WebSocket as a decimal number (`12.22`), while the frontend already worked with integer cents (#149).
This ADR documents the decision made in #150.

## Decision

Money is an integer number of cents everywhere: in the database, in the backend, in every message and in the frontend.
Only the display divides by 100.

## Considered options

- **Floating-point numbers, as in version 2:** not chosen.
  Floats can't represent most cent amounts exactly (`0.1 + 0.2 != 0.3`), so sums and balances drift by fractions of a cent, which a cash register can't tolerate.
- **A decimal type (PostgreSQL `NUMERIC`):** exact in the database, but sent over JSON it arrives in JavaScript as a float again, so the frontend would have to convert every amount (#149).
  Replaced.
- **Integer cents:** chosen; exact on every layer, and the same representation in Go, JSON and TypeScript.

## Consequences

- The frontend displays money with `$n(cents / 100, 'currency')` and converts inputs back with `Math.round(value * 100)`.
- The currency is EUR in every locale; a locale changes only the number format, never the currency.
