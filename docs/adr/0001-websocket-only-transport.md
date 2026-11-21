# 0001. WebSocket-only transport

- Status: Accepted
- Date: 2025-12-04 (recorded 2026-11-21)

## Context

AVHBS runs on several tills at once, all working on the same accounts, products and balances.
Version 2 used a REST API.
This ADR documents a decision made when version 3 was started.

## Decision

Client and server communicate only over a single WebSocket (`/ws`).
Every message is JSON of the form `{type, payload}` and is validated against the shared JSON Schemas in `schemas/`, by the backend and by the frontend.
The HTTP server provides no REST endpoints; besides `/ws` it only serves the frontend.

## Considered options

- **REST, as in version 2:** every till would have to poll to see changes made at another till, and requests and responses would have no shared, validated contract.
  Not chosen.
- **A single WebSocket:** chosen, for two reasons:
  - Live sync between tills: every change is broadcast to all connected clients, so no till has to poll.
  - One channel with one contract: both sides validate the same JSON Schemas, so the contract can't drift.

## Consequences

- Every feature goes through the message protocol: the schema in `schemas/`, the backend (`backend/internal/ws`) and the frontend (`webSocketClient.ts`, stores, composables) change together.
- Changes are broadcast to the other clients, so every table needs broadcast handling on both sides.
