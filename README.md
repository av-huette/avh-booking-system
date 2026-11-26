# AVH Booking System

A point-of-sale system for the bar of a student association: members pay with their own account, several tills stay in sync, and German fiscal requirements are part of the design.

> **Status:** version 3 is in development and not yet ready for production.

## What it is for

AVHBS is meant for associations that run a bar for their members.
Members have accounts with a balance and a credit limit, the tills book drinks and food onto these accounts, and the people running the bar manage products, prices and members.
It is built for associations in German-speaking countries, with an English and a German interface, and with the German rules for cash registers in mind.

AVHBS was developed for the bar of the [Akademischer Verein Hütte e.V.](https://av-huette.de/) (AV Hütte).

## Screenshots

Screenshots will follow once the interface of version 3 is stable.

## Installation

The installation with Docker Compose will be described here with the first usable release of version 3.

## Development

Requirements: Go 1.27 or newer, Node.js 22, and PostgreSQL 17 (or Docker to run it).

1. Start a database, for example with Docker:

   ```bash
   docker run -d --name avhbs-db -p 5432:5432 \
     -e POSTGRES_DB=avhbs -e POSTGRES_USER=avhbs -e POSTGRES_PASSWORD=avhbs \
     postgres:17-alpine
   ```

2. Configure and start the backend:

   ```bash
   cd backend
   cp .env.template .env      # set the DB_* values; AVHBS_FRONTEND_PATH can stay empty for development
   go run ./cmd/migrate up    # create the database schema
   go run ./cmd/server        # listens on port 8081
   ```

3. Start the frontend in a second terminal:

   ```bash
   cd frontend
   npm install
   npm run dev                # then open the address Vite prints, usually http://localhost:5173
   ```

How to contribute, including the checks to run before a pull request, is described in [CONTRIBUTING.md](CONTRIBUTING.md).

## Documentation

- [CONTRIBUTING.md](CONTRIBUTING.md): branches, commits, checks and pull requests.
- [docs/README.md](docs/README.md): where each kind of information belongs.
- [docs/architecture.md](docs/architecture.md): how the system works.
- [docs/glossary.md](docs/glossary.md): what the domain terms mean.
- [docs/adr/](docs/adr/README.md): why AVHBS is built the way it is.
- [AGENTS.md](AGENTS.md): instructions for coding agents; the conventions for backend and frontend are in [backend/AGENTS.md](backend/AGENTS.md) and [frontend/AGENTS.md](frontend/AGENTS.md).

## Support

Questions and bug reports go to [GitHub Issues](https://github.com/av-huette/avh-booking-system/issues).
Security problems: please report them privately, see [SECURITY.md](SECURITY.md).

## License

AVHBS is free software under the [GNU Affero General Public License v3.0 or later](LICENSE).
If you run a changed version for others, for example for the members of your association, the license asks you to offer them its source code.
