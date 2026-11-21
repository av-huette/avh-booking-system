# Architecture

How AVHBS works: the parts of the system, the messages between them and the path a request takes.
Why it is built this way is in the [decision records](adr/README.md); the rules for writing code are in the `AGENTS.md` files.

## Overview

```mermaid
flowchart LR
    till1["Till 1 (browser)"] -- "HTTPS / WebSocket" --> caddy[Caddy]
    till2["Till 2 (browser)"] -- "HTTPS / WebSocket" --> caddy
    caddy --> server["Go server<br/>frontend files and /ws"]
    server --> db[(PostgreSQL)]
```

The Go server serves the built frontend and a single WebSocket endpoint, `/ws`; Caddy in front of it takes care of HTTPS.
Why: [ADR 0001](adr/0001-websocket-only-transport.md), [0002](adr/0002-postgresql-with-pgx-and-raw-sql.md), [0003](adr/0003-docker-compose-with-caddy.md).

## Message protocol

Every message is JSON of the form `{"type": …, "payload": …}`.
The schemas are in `schemas/`: `message.json` for the envelope, `sql/` for queries, mutations and their results, `db/` for the rows of single tables.

| Type | Direction | Payload |
|---|---|---|
| `ping`, `pong` | till → server, server → till | a timestamp; the till's heartbeat |
| `query` | till → server | a table, optionally filters, sorting and a limit |
| `queryResultList` | server → till | the table and the matching rows |
| `mutation` | till → server | an operation (`insert`, `update`, `delete`), a table, the key of the row (`where`) and the new values |
| `mutationResult` | server → till | the table, the operation and the ID of the affected row |
| `broadcast` | server → all tills | the table and the changed data |
| `error` | server → till | a code, a message and details |

An update sends only the fields that change; the server merges them into the stored row.

## Data flow

When the app starts, it queries all tables it needs (`getAllFromDb()` in `frontend/src/store/socketStore.ts`).

Loading the products:

```mermaid
sequenceDiagram
    participant UI as View
    participant Socket as socketStore / webSocketClient
    participant Server as Go server (ws)
    participant Repo as ProductStore (repo)
    participant DB as PostgreSQL
    UI->>Socket: queryProducts()
    Socket->>Server: query {table: product}
    Server->>Repo: Get(query)
    Repo->>DB: SELECT
    DB-->>Repo: rows
    Repo-->>Server: products
    Server-->>Socket: queryResultList {table: product, data}
    Socket->>UI: ProductStore.patchProducts(data)
```

Changing a product at one till:

```mermaid
sequenceDiagram
    participant Till as Till 1
    participant Server as Go server (ws)
    participant DB as PostgreSQL
    participant Others as All tills
    Till->>Server: mutation {operation: update, table: product, where, values}
    Server->>DB: read the row, merge the values, UPDATE
    Server-->>Till: mutationResult {table: product, operation: update, id}
    Server->>DB: read the changed row
    Server-->>Others: broadcast {table: product, data}
    Others->>Others: ProductStore.patchProducts([data])
```
