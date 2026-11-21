# 0003. Docker Compose with Caddy

- Status: Accepted
- Date: 2026-04-01 (recorded 2026-11-21)

## Context

AVHBS is meant to be installed by other associations, usually by a member who is comfortable with a terminal and Docker but is not a developer of the project.
The installation has to be simple and the same everywhere.
This ADR documents a decision made when the deployment setup was added.

## Decision

AVHBS is deployed with Docker Compose (`docker-compose.yml`) as three services:

- `avhbs`: one image with the Go server, which also serves the built frontend;
- `db`: PostgreSQL;
- `caddy`: Caddy as reverse proxy in front of `avhbs` (`Caddyfile` for a domain, `Caddyfile.local` for local use).

## Considered options

- **Docker Compose:** chosen, because it installs everything with one command, gives the same setup everywhere and includes the database.
- **One image, with the Go server serving the frontend:** chosen, because there is only one container to deploy, the page and the WebSocket share one origin (so no CORS), and it is simpler than a separate web server for static files.
- **Caddy as reverse proxy:** the reason for choosing Caddy over alternatives was not recorded.
  The setup relies on two of its properties: automatic HTTPS for a domain, and proxying the WebSocket with a two-line configuration.

## Consequences

- An adopting association needs Docker and, for HTTPS, a domain pointing to the server.
- The frontend is built into the backend image; frontend and backend are released together.
- The frontend must reach the WebSocket through the page's own origin, so that it works behind the proxy.
