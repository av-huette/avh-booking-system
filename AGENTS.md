# AVHBS Agent Instructions

AVH Booking System (AVHBS) is a point-of-sale system for a club. The Go backend lives in `backend/`, the Vue 3
frontend in `frontend/`; each has its own `AGENTS.md` with stack-specific rules and commands.

## Architecture

- WebSocket-only transport: there are no REST endpoints. Client and server exchange JSON `{type, payload}` messages,
  validated on both sides against the shared JSON Schemas in `schemas/`.
- Changing a message means changing the schema in `schemas/`, the backend (`internal/ws`) and the frontend
  (`webSocketClient.ts`, stores, composables) in the same change.
- Money is an integer number of cents on both sides, never a float.

## Rules

- Never run `git commit` or `git push`, and never open a PR. A human reviews every AI-generated diff and commits it.
- Keep changes minimal: touch only the files the task requires, and add no abstractions, config or error handling
  beyond what the task needs.
- Ask before any action git cannot undo: applying migrations (there are no down migrations), changing or deleting
  data in a database, deleting untracked or ignored files (e.g. `.env`), discarding uncommitted changes
  (`git reset --hard`, `git checkout --`, `git clean`), removing Docker volumes, and changes outside the repository.
- Before suggesting a commit, run the checks of every area you changed (see that area's `AGENTS.md`) and report
  which checks ran and which were skipped.
- After a change, suggest a commit message: a title under 50 characters in imperative mood, and a 1–2 sentence
  description.
- Before a PR, review the whole branch: consistency between commits, these rules, and docs that mention the change.
- A developer-defined workflow may exist at `.notes/WORKFLOW.md`. If it does, read it before starting a task and
  follow it.
