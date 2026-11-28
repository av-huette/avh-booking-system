<script setup lang="ts">
import { computed } from 'vue'
import type { Product } from '../composables/product'
import type { Category } from '../composables/category'
import type { ProductGroup } from '../composables/productGroup'
import { useProductStore } from '../store/ProductStore'
import { useProductVisibilityStore } from '../store/ProductVisibilityStore'
import { useCategoryStore } from '../store/CategoryStore'
import { useProductGroupStore } from '../store/ProductGroupStore'
import { useSettingStore } from '../store/SettingStore'

const product$ = useProductStore()
const visibility$ = useProductVisibilityStore()
const category$ = useCategoryStore()
const productGroup$ = useProductGroupStore()

const setting$ = useSettingStore()

// A product is active if it has at least one visibility (at the current location)
const activeProductIds = computed(() => new Set(
  visibility$.visibilities
    .filter(vis => setting$.locationId == null || vis.locationId == setting$.locationId)
    .map(vis => vis.productId)
))

const activeProducts = computed(() => product$.products.filter(prod =>
  activeProductIds.value.has(prod.id ?? -1) && category$.byId(prod.category)?.enabled
))

const byName = (a: Product, b: Product) => a.name.localeCompare(b.name)

const sections = computed(() => {
  return [...category$.productCategories]
    .sort((a, b) => a.title.localeCompare(b.title))
    .map((category: Category) => {
      const products = activeProducts.value.filter(prod => prod.category == category.id)
      const ungrouped = products
        .filter(prod => prod.productGroup == 0 || productGroup$.byId(prod.productGroup) == undefined)
        .sort(byName)

      const groups = productGroup$.tree
        .map(({ group, depth }) => ({
          group,
          depth,
          products: products.filter(prod => prod.productGroup == group.id).sort(byName)
        }))
        // Keep groups that contain products themselves or in one of their subgroups
        .filter(({ group }) => {
          const ids = [group.id, ...productGroup$.descendantIds(group.id as number)]
          return products.some(prod => ids.includes(prod.productGroup))
        }) as { group: ProductGroup, depth: number, products: Product[] }[]

      return { category, ungrouped, groups }
    })
    .filter(section => section.ungrouped.length > 0 || section.groups.length > 0)
})
</script>

<template>
  <div class="menu-display">
    <section v-for="{ category, ungrouped, groups } in sections" :key="category.id" class="menu-category">
      <h1 class="title is-2">
        <span v-if="category.icon" class="icon mr-3"><icon :icon="category.icon" /></span>
        <span>{{ category.title }}</span>
      </h1>

      <div class="menu-columns">
        <div v-if="ungrouped.length > 0" class="menu-block">
          <div v-for="product in ungrouped" :key="product.id" class="menu-product">
            <span class="menu-product-name">{{ product.name }}</span>
            <span class="has-text-grey">{{ product.size }} {{ product.getUnit()?.name }}</span>
            <span class="menu-product-price has-text-weight-semibold">{{ $n(product.price / 100, 'currency') }}</span>
          </div>
        </div>

        <div
          v-for="{ group, depth, products } in groups"
          :key="group.id"
          class="menu-block"
          :style="{ paddingLeft: depth * 1 + 'em' }">
          <h2
            class="has-text-weight-bold has-text-primary"
            :style="{ fontSize: Math.max(1.5 - depth * 0.25, 1) + 'em' }">
            {{ group.name }}
          </h2>
          <div v-for="product in products" :key="product.id" class="menu-product">
            <span class="menu-product-name">{{ product.name }}</span>
            <span class="has-text-grey">{{ product.size }} {{ product.getUnit()?.name }}</span>
            <span class="menu-product-price has-text-weight-semibold">{{ $n(product.price / 100, 'currency') }}</span>
          </div>
        </div>
      </div>
    </section>

    <p v-if="sections.length === 0" class="title is-3 has-text-grey has-text-centered is-italic">
      Keine Produkte verfügbar
    </p>
  </div>
</template>

<style scoped>
.menu-display {
  min-height: 100vh;
  padding: 2rem 3rem;
  font-size: 1.5rem;
}

.menu-category:not(:last-child) {
  margin-bottom: 3rem;
}

.menu-columns {
  column-count: 1;
  column-gap: 4rem;
}

@media (min-width: 1024px) {
  .menu-columns {
    column-count: 2;
  }
}

@media (min-width: 1600px) {
  .menu-columns {
    column-count: 3;
  }
}

.menu-block {
  break-inside: avoid;
  margin-bottom: 1.5rem;
}

.menu-product {
  display: flex;
  align-items: baseline;
  gap: 0.75em;
  padding: 0.2em 0;
}

.menu-product-name {
  flex-shrink: 1;
}

.menu-product-price {
  margin-left: auto;
  white-space: nowrap;
}
</style>
