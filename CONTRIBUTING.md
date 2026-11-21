# Contributing to AVHBS

Thanks for helping with the AVH Booking System.
This guide describes how changes get from an idea to the main branch.
It applies to everyone, whether you write the code yourself or with a coding agent.

## Before you start

- Report bugs and ask questions in [GitHub Issues](https://github.com/av-huette/avh-booking-system/issues).
- For the development setup, see the quick start in the [README](README.md).
- Documentation, code comments, commit messages and pull requests are written in English.

## Branches

- Start from `v3/main`, the development branch of version 3.
- Name the branch `v3/<area>/<topic>` for work on one side (e.g. `v3/backend/monetary-values`), or `v3/<topic>` when it spans backend and frontend.
- Topics are lowercase kebab-case, a short noun phrase, without suffixes such as `_2` or `-bkp`.
- The `v3/` prefix applies while version 3 is in development; it changes once version 3 is merged into `master`.

## Commits

- Keep commits small and focused: one independent change per commit, and every commit builds.
- Write a title of under 50 characters in imperative mood ("Add code owners", not "Added code owners"), followed by a description of one or two sentences that says what changed and why.
  Like the Markdown files, the description has one sentence per line and is not wrapped otherwise.
- The commit message records what changed; lasting knowledge belongs in the documentation (see [`docs/README.md`](docs/README.md)).

## Checks

Before opening a pull request, run the checks of every area you changed and say in the pull request which ran and which you had to skip.
The commands are in the `AGENTS.md` of each area:

- Backend: [`backend/AGENTS.md`](backend/AGENTS.md), section "Checks".
- Frontend: [`frontend/AGENTS.md`](frontend/AGENTS.md), section "Checks".

## Pull requests

- One logical change per pull request, usually a few commits.
- Before opening it, review the whole branch: are the commits consistent, are the conventions followed, and is documentation that mentions the change still correct?
- The description gives a short summary and a test plan: how a reviewer can check that the change works.
  For a point-of-sale system that usually means walking through booking, payment and order views, not only tests.
- Every pull request is reviewed by one of the code owners (see [`.github/CODEOWNERS`](.github/CODEOWNERS)).

## AI-assisted contributions

Coding agents are welcome.
They follow the [`AGENTS.md`](AGENTS.md) files, which also forbid them to commit or push.

- A human reviews every AI-generated change, commits it, and is responsible for it like for any other change.
- Crediting the agent with a `Co-Authored-By` trailer in the commit message is optional.

## Code conventions and documentation

- The conventions for the code, the commands and where new code goes are in [`backend/AGENTS.md`](backend/AGENTS.md) and [`frontend/AGENTS.md`](frontend/AGENTS.md).
  Despite the name, they are written for developers and agents alike.
- Where each kind of information belongs (README, documentation, decision records, comments) is in [`docs/README.md`](docs/README.md).
- Markdown files are written with one sentence per line and no other line breaks inside a paragraph.
  A change then touches only the sentences it changes, which keeps diffs readable.
- A decision that shapes how the code is written and is hard to reverse gets an architecture decision record, see [`docs/adr/README.md`](docs/adr/README.md).
