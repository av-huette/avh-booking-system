# Documentation

Where each kind of information about AVHBS belongs, for developers and coding agents. Before writing something
down, find its place here; before searching for something, look where it should be.

## Where things go

| File | Holds | Does not hold |
|---|---|---|
| `README.md` | What AVHBS is, status, quick start, links | rules, architecture details |
| `CONTRIBUTING.md` | Process: branches, commits, checks, pull requests, reviews | how the code works |
| `.github/CODEOWNERS` | Who reviews which paths | anything else |
| `docs/architecture.md` | How the system works: message protocol, layers, data flow | the why (ADRs) |
| `docs/glossary.md` | What the domain terms mean (account, user, booking, ...) | implementation details |
| `docs/adr/NNNN-*.md` | Why a decision was made, the options, the consequences | how-tos; edits once accepted |
| `AGENTS.md` (root) | One-line rules every agent must not miss, plus links | knowledge, process details |
| `backend/AGENTS.md`, `frontend/AGENTS.md` | Commands, checks, environment, where code goes, conventions; for developers and agents | rules that span both stacks (root) |
| Code comments | Why non-obvious code is the way it is | what the code does; history |
| Commit messages, PR descriptions | What changed and why, for this change | lasting knowledge (docs, ADRs) |

## In this folder

- [Architecture decision records](adr/README.md): why AVHBS is built the way it is.

## Deciding where something goes

Apply these questions in order:

1. Can it be read from the code, the schemas or the git history? Then write nothing.
2. Why did we do it this way? An ADR.
3. How does it work, or how do I do X? A document in `docs/`.
4. How do we work together? `CONTRIBUTING.md`.
5. What is this, and how do I start? `README.md`.
6. Would a coding agent get it wrong without noticing? Additionally a one-line rule in the matching `AGENTS.md`
   that links to the place from questions 2 to 5.
