<script setup lang="ts">
import { ref, computed } from 'vue'
import type { RouteLocationRaw } from 'vue-router'

type SettingsArea = 'product' | 'account' | 'payment' | 'system'

const props = defineProps<{ area: SettingsArea }>()

interface MenuLink {
  label: string
  icon: [string, string]
  to: RouteLocationRaw
}

const menus: Record<SettingsArea, { title: string, links: MenuLink[] }> = {
  product: {
    title: 'Wareneinstellungen',
    links: [
      { label: 'Produkte', icon: ['fas', 'cart-shopping'], to: { name: 'ProductSettings' } },
      { label: 'Kategorien', icon: ['fas', 'tags'], to: { name: 'CategorySettings', params: { type: 'product' } } },
      { label: 'Produktgruppen', icon: ['fas', 'list'], to: { name: 'ProductGroupSettings' } },
      { label: 'Einheiten', icon: ['fas', 'ruler'], to: { name: 'UnitSettings' } },
    ],
  },
  account: {
    title: 'Accounteinstellungen',
    links: [
      { label: 'Accounts', icon: ['fas', 'users'], to: { name: 'AccountSettings' } },
      { label: 'Kategorien', icon: ['fas', 'tags'], to: { name: 'CategorySettings', params: { type: 'account' } } },
    ],
  },
  payment: {
    title: 'Zahlungseinstellungen',
    links: [
      { label: 'Stripe', icon: ['fas', 'credit-card'], to: { name: 'PaymentSettings' } },
      { label: 'Steuersätze', icon: ['fas', 'percent'], to: { name: 'VatSettings' } },
    ],
  },
  system: {
    title: 'Systemeinstellungen',
    links: [
      { label: 'Firma', icon: ['fas', 'building'], to: { name: 'ApplicationSettings', params: { section: 'company' } } },
      { label: 'Darstellung', icon: ['fas', 'palette'], to: { name: 'ApplicationSettings', params: { section: 'look' } } },
      { label: 'Lokalisierung', icon: ['fas', 'language'], to: { name: 'ApplicationSettings', params: { section: 'localisation' } } },
    ],
  },
}

const menu = computed(() => menus[props.area])
const menuOpen = ref(false)
</script>

<template>
  <div>
  <!-- Wrapper keeps the parent's scoped `.columns` styles off the layout columns -->
  <div class="columns">
    <div class="column is-2">
      <button class="button is-fullwidth is-hidden-tablet mb-3" @click="menuOpen = !menuOpen">
        <span class="icon"><icon :icon="['fas', 'bars']" /></span>
        <span>{{ menu.title }}</span>
      </button>
      <aside class="menu" :class="{ 'is-hidden-mobile': !menuOpen }">
        <p class="menu-label is-hidden-mobile">{{ menu.title }}</p>
        <ul class="menu-list">
          <li v-for="link in menu.links" :key="link.label">
            <router-link :to="link.to" exact-active-class="is-active" @click="menuOpen = false">
              <span class="icon-text">
                <span class="icon"><icon :icon="link.icon" /></span>
                <span>{{ link.label }}</span>
              </span>
            </router-link>
          </li>
        </ul>
      </aside>
    </div>
    <div class="column">
      <slot />
    </div>
  </div>
  </div>
</template>
