<template>
  <div class="accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">

    <div v-if="selectedGroup" class="group-overlay">
      <div class="group-overlay-header">
        <button class="button is-small" @click="groupStack.pop()">
          <span class="icon is-small"><icon :icon="['fas', 'arrow-left']" /></span>
          <span>Zurück</span>
        </button>
        <span class="group-overlay-title">{{ groupStack.map(g => g.name).join(' › ') }}</span>
        <button class="delete is-medium" @click="groupStack = []" />
      </div>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <Button
          v-for="group in childGroupsForSelectedGroup"
          :key="'group-' + group.id"
          @click="onGroupClick(group)"
          title="select product group">
            {{ group.name }}
            <icon style="margin-left:.5em" :icon="['fas', 'list']" />
        </Button>
        <Button
          v-for="product in productsForSelectedGroup"
          :key="product.id"
          @click="cart$.addToCart(product)"
          title="select product">
            {{ product.name }} {{ product.size }} {{ product.getUnit()?.name }}
            <span class="cartHint" v-if="cart$.productCartQuantity(product) != -1">{{ cart$.productCartQuantity(product) }}</span>
        </Button>
      </div>
    </div>

    <div class="dictionary" v-for="(dict, key) of processedProducts.dict" :key="key">
      <span class="title is-1">{{ key }}</span>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <template v-for="product in dict">
          <Button
            v-if="product.constructor.name == 'Product'"
            @click="cart$.addToCart(product)"
            title="select product">
              {{ product.name }} {{ product.size }} {{ product.getUnit()?.name }}
              <span class="cartHint" v-if="cart$.productCartQuantity(product) != -1">{{ cart$.productCartQuantity(product) }}</span>
          </Button>

          <Button
            v-if="product.constructor.name == 'ProductGroup'"
            @click="onGroupClick(product)"
            title="select product group">
              {{ product.name }}
              <icon style="margin-left:.5em" :icon="['fas', 'list']" />
          </Button>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
.accountList{
  height: var(--_height);
  overflow-y: scroll;
  position: relative;
}
.dictionary{
  display:grid;
  grid-template-columns: 10% 90%;
  margin-bottom:1em;
  .title{
    justify-self: center;
    margin-bottom:0;
  }
}

.button:has(span){
  padding-right:1.8em;
}
.cartHint{
  position:absolute;
  right:.25em;
  top:50%;
  transform: translateY(-50%);
  font-size:.8em;
  background-color:rgba(255,255,255,.3);
  border-radius: 100vh;
  width:1.6em;
  aspect-ratio: 1;
}

@media screen and (min-width: 768px) {
  .button:has(span){
    padding: calc(var(--bulma-button-padding-vertical) - var(--bulma-button-border-width)) calc(var(--bulma-button-padding-horizontal) - var(--bulma-button-border-width));
  }
  .cartHint{
    display:none;
  }
}

.group-overlay {
  position: absolute;
  inset: 0;
  background-color: var(--bulma-scheme-main);
  overflow-y: auto;
  z-index: 10;
  padding: .5rem;
}
.group-overlay-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: .5rem;
  border-bottom: 1px solid var(--bulma-border);
  margin-bottom: .5rem;
}
.group-overlay-title {
  font-weight: 600;
  font-size: 1.1em;
}
</style>

<script lang="ts">
import { Product } from '../../composables/product';
import { ProductGroup } from '../../composables/productGroup.ts';
import { useProductStore } from '../../store/ProductStore';
import type { PropType } from 'vue';
import { useResizeObserver } from '@vueuse/core';
import { useCartStore } from '../../store/CartStore';
import { useAccountStore } from '../../store/AccountStore';
import { useProductVisibilityStore } from '../../store/ProductVisibilityStore';
import { useProductGroupStore } from '../../store/ProductGroupStore.ts';
import Button from '../../composables/elements/Button.vue';

export default {
  data(){
    return {
      product$: useProductStore(),
      height: 0,
      resizeElement: {} as HTMLElement,
      cart$: useCartStore(),
      account$: useAccountStore(),
      visibility$: useProductVisibilityStore(),
      productGroup$: useProductGroupStore(),
      groupStack: [] as ProductGroup[],
    }
  },
  props: {
    products: {
      type: Array as PropType<Product[]>
    }
  },
  components: {
    Button
  },
  computed: {
    processedProducts() {
      const dict: {[key: string]: (Product | ProductGroup)[]} = {};
      const groupProductsMap: {[key: number]: Product[]} = {};
      const childGroupsMap: {[key: number]: ProductGroup[]} = {};

      this.products?.forEach(prod => {
        // Chain from the top-level group down to the product's own group
        const path = prod.productGroup != 0 ? this.productGroup$.path(prod.productGroup) : [];

        if (path.length > 0) {
          if (!groupProductsMap[prod.productGroup]) {
            groupProductsMap[prod.productGroup] = [];
          }
          groupProductsMap[prod.productGroup].push(prod);

          for (let i = 0; i < path.length - 1; i++) {
            const parentId = path[i].id ?? -1;
            if (!childGroupsMap[parentId]) {
              childGroupsMap[parentId] = [];
            }
            if (!childGroupsMap[parentId].includes(path[i + 1])) {
              childGroupsMap[parentId].push(path[i + 1]);
            }
          }
        }

        let char = prod.name[0].toUpperCase();
        let charCode = char.charCodeAt(0);

        if (path.length > 0) {
          char = path[0].name[0].toUpperCase();
          charCode = char.charCodeAt(0);
        }

        if (charCode >= 65 && charCode <= 90) { // A-Z
        } else if (charCode >= 48 && charCode <= 57) { // 0-9
          char = "#";
        } else {
          char = "?";
        }

        const toAdd: Product | ProductGroup = path.length > 0 ? path[0] : prod;

        if (dict[char] === undefined) {
          dict[char] = [toAdd];
        } else if (!dict[char].includes(toAdd)) {
          dict[char].push(toAdd);
        }
      });

      return { dict, groupProductsMap, childGroupsMap };
    },
    selectedGroup(): ProductGroup | null {
      return this.groupStack[this.groupStack.length - 1] ?? null;
    },
    productsForSelectedGroup() {
      if (!this.selectedGroup) return [];
      return this.processedProducts.groupProductsMap[this.selectedGroup.id ?? -1] ?? [];
    },
    childGroupsForSelectedGroup() {
      if (!this.selectedGroup) return [];
      return this.processedProducts.childGroupsMap[this.selectedGroup.id ?? -1] ?? [];
    }
  },
  watch: {
    products() {
      // this.selectedGroup = null;
    }
  },
  methods: {
    onGroupClick(group: ProductGroup){
      // Groups with visible subgroups always open, so the subgroups stay reachable
      if ((this.processedProducts.childGroupsMap[group.id ?? -1] ?? []).length > 0) {
        this.groupStack.push(group);
        return;
      }
      const groupProducts = this.processedProducts.groupProductsMap[group.id ?? -1] ?? [];
      const accounts = this.account$.selected;
      // Every account sees exactly one product of the group -> add the visible product per account
      const productPerAccount = accounts.map((acc) => {
        return groupProducts.filter((prod) => this.visibility$.categoryIsVisible(acc.category ?? -1, prod.id ?? -1));
      });
      if (productPerAccount.every((prods) => prods.length === 1)) {
        accounts.forEach((acc, i) => this.cart$.addToCart(productPerAccount[i][0], [acc]));
        return;
      }
      this.groupStack.push(group);
    },
    onResize(){
      let y = window.innerHeight;
      let _y = this.resizeElement.getBoundingClientRect().top;
      let dy = y - _y;
      this.height = dy - 15;
    }
  },
  mounted() {
    this.resizeElement = this.$refs.resizeBox as HTMLElement;
    useResizeObserver(this.resizeElement, () => {
      this.onResize();
    });
  },
}
</script>