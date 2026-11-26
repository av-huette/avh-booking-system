# 0004. Traceable bookings

- Status: Accepted
- Date: 2026-01-24 (recorded 2026-11-21)

## Context

A point-of-sale system in Germany falls under the GoBD (principles for proper bookkeeping) and the KassenSichV (cash register security regulation): nothing that was booked may disappear or change without a trace, and every balance must be explainable from recorded events.
Orders, favorites and account options reference accounts.
This ADR documents the principle; the booking path that implements it will get its own ADR.

## Decision

- Accounts are never deleted.
  An account that is no longer used is disabled (`enabled = false`).
- A balance never changes through a general update of the account.
  It changes only together with a recorded transaction, such as an order or a payment.

## Considered options

- **Deleting accounts, with cascading deletes:** not chosen; it would remove the orders along with the account and cut the booking history off from whoever placed the orders.
- **A general account update that may also set the balance:** not chosen; a balance could then change without any booking or payment behind it.
- **Disabling instead of deleting, and balance changes only with a recorded transaction:** chosen.

## Consequences

- The backend rejects deleting accounts, and `AccountStore.Update` never writes the balance.
- The booking path must change a balance in the same database transaction as the order or payment that causes it.
