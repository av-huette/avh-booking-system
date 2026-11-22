# AVHBS Frontend Conventions

Conventions for working in `frontend/` (Vue 3, TypeScript, Vite, Pinia, Vue Router, vue-i18n, Bulma with Sass), for developers and coding agents.
Cross-cutting rules are in the root `AGENTS.md`; the contribution process is in `CONTRIBUTING.md`.

## Commands (run from `frontend/`)

```bash
npm run dev      # Vite dev server (hot reload, no type check)
npm run build    # Type-check (vue-tsc), then build
```

## Checks

Before suggesting a commit: `npm run build` passes (type check, then build).

## Environment

No environment variables are read.
The WebSocket client ID is set in the app (client setup) and persisted.

## Where code goes

| New …               | Goes into                                                                                     |
|---------------------|-----------------------------------------------------------------------------------------------|
| Page                | `src/views/<Name>View.vue`, route in `src/router/router.ts`                                   |
| Settings page       | `src/views/Settings/<Entity>Settings.vue` (list), `<Entity>SettingsSingle.vue` (add and edit) |
| Component           | `src/components/<Feature>/<Name>.vue`                                                         |
| Reusable UI element | `src/composables/elements/<Name>.vue`                                                         |
| Domain type         | `src/composables/<entity>.ts`: interface and class of the same name                           |
| Store               | `src/store/<Entity>Store.ts`                                                                  |
| Query or mutation   | action in `src/store/socketStore.ts`                                                          |
| Handling a result   | case in `src/api/webSocketClient.ts`                                                          |
| Strings             | `src/locales/en.ts` and `src/locales/de.ts`                                                   |
| Global styles       | `src/assets/` (theme: `catppuccin.scss`)                                                      |

A new table needs, after the backend side (`backend/AGENTS.md`):

1. Domain type and store with a `patch<Entities>` action.
2. Query and mutation actions in `src/store/socketStore.ts`; the query also in `getAllFromDb()`.
3. Cases for the table in both the `queryResultList` and the `broadcast` switch of `src/api/webSocketClient.ts`; without the `broadcast` case, changes from other clients are dropped silently.
4. Settings pages and routes, if the table is editable.

## Rules

### Components

- New code: components use `<script setup>` (Composition API).
  Existing Options API components are not converted unless asked.

### State and WebSocket

- One store per entity in `src/store/<Entity>Store.ts`, written as an options store like the existing ones, with a `patch<Entities>` action that upserts by ID.
- Send queries and mutations only through actions in `src/store/socketStore.ts`; components never send on the socket themselves.

### Strings

- New code: user-facing strings go into `src/locales/en.ts` (the reference) and `src/locales/de.ts`, and are shown with `$t()` (why: `docs/adr/0006-internationalization.md`).
  Exception: receipts (`src/components/Receipt/`) are always German, whatever the UI language: a receipt is a document handed to customers and part of the fiscal record, so it looks the same on every till.

### Money

- Display money with `$n(cents / 100, 'currency')`; convert inputs back with `Math.round(value * 100)` (why: `docs/adr/0005-money-as-integer-cents.md`).
- The currency is EUR in every locale; a locale changes only texts and number formatting.
  Receipts always use German formatting (`toLocaleString('de-DE', …)`); see Strings.

### Styling

- Use Bulma; no other CSS framework.
- Colors come from the Catppuccin theme (`src/assets/catppuccin.scss`, flavor chosen via `themeStore`).
  Use Bulma's semantic classes (`is-primary`, `is-warning`, …) and theme variables, not hard-coded colors.
- The primary color is Catppuccin **Sky** in every flavor, not Mauve, Catppuccin's usual accent.
- Component styles are `<style scoped>`; global styles belong in `src/assets/`.
- Use the global `<icon>` component (FontAwesome solid icons) for icons.
