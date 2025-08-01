<template>
  <div class="accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">
    <div class="dictionary" v-for="(dict, key) of productsInOrder" :key="key">
      <span class="title is-1">{{ key }}</span>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <div v-for="product in dict">
          <button class="button is-fullwidth" title="select product">{{ product.name }} {{ product.size }} {{ product.unit }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.accountList{
  height: var(--_height);
  overflow-y: scroll;
}
.dictionary{
  display:grid;
  grid-template-columns: 10% 90%;
  margin-bottom:1em;
  .title{
    justify-self: center;
  }
}
</style>

<script lang="ts">
import { Product } from '../../composables/product';
import { useProductStore } from '../../store/ProductStore';
import type { PropType } from 'vue';
import { useResizeObserver } from '@vueuse/core';

export default {
  data(){
    return {
      product$: useProductStore(),
      height: 0,
      resizeElement: {} as HTMLElement
    }
  },
  props: {
    products: {
      type: Array as PropType<Product[]>
    }
  },
  computed: {
    productsInOrder() {
      var dict: {[key: string]: Product[]} = {};
      this.products?.forEach(prod => {
      var char = prod.name[0].toUpperCase();
      var charCode = char.charCodeAt(0);
      if (charCode >= 65 && charCode <= 90) { // A-Z
      } else if (charCode >= 48 && charCode <= 57) { // 0-9
        char = "#";
      } else {
        char = "?";
      }
      if (dict[char] === undefined) {
        dict[char] = [prod]
      } else {
        dict[char].push(prod);
      }
    })
    return dict;
    }
  },
  methods: {
    onResize(){
      let y = window.innerHeight;
      let _y = this.resizeElement.getBoundingClientRect().top;
      let dy = y - _y;
      this.height = dy -15 ;
    }
  },
  mounted() {
    this.resizeElement = this.$refs.resizeBox as HTMLElement
    // Watch for resizing
    useResizeObserver(this.resizeElement, () => {
      this.onResize();
    })
  },
}
</script>