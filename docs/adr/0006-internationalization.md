# 0006. Internationalization

- Status: Accepted
- Date: 2026-11-21

## Context

Version 3 started with German texts: the older components show them through vue-i18n (`$t()`), newer components
contain hard-coded German text. A language switcher exists in the application settings, and an English locale file
existed with an older, unrelated set of keys. AVHBS is meant to be adopted by other associations, not all of them
German-speaking.

## Decision

- New code shows every user-facing string through `$t()`.
- English (`src/locales/en.ts`) is the reference language and the fallback.
- German (`src/locales/de.ts`) is a full translation and the language most installations will use.
- The starting language is configurable, and the language a till has chosen is kept across reloads.
- Exception: receipts are always German.

## Considered options

- **German only, without internationalization:** simplest for today, but excludes associations that don't speak
  German, and the existing language switcher would switch nothing. Not chosen.
- **German as the reference language:** matches today's main use, but makes every further language a translation
  of German, which fewer contributors can read. Not chosen.
- **English as the reference, German as a full translation:** chosen. English is the common ground for further
  languages; developers write the English string together with the code, and every other language translates it;
  German stays fully supported because most installations will use it.

## Consequences

- A new language is a new locale file, translated from `en.ts`.
- Components with hard-coded text are migrated to `$t()` step by step.
- The currency stays EUR in every language (see ADR 0005).
- Receipts are a document handed to customers and part of the fiscal record, so they look the same on every till,
  whatever the interface language.
