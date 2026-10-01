<template>
  <tr class="cartProduct">
    <td class="cell productQuantity">
      <div class="field has-addons" v-if="allowEdit">
        <p class="control">
          <button class="button" @click="cart$.removeFromCart(content.product)"><icon :icon="['fas', 'trash']" /></button>
        </p>
        <p class="control has-icons-left has-icons-right">
          <input :value="content.quantity" @change="setQuant($event)" type="number" class="input"></input>
          <span class="icon is-left" @click="cart$.changeQuantity(content.product, -1)"><icon :icon="['fas', 'circle-minus']" /></span>
          <span class="icon is-right" @click="cart$.changeQuantity(content.product, 1)"><icon :icon="['fas', 'circle-plus']" /></span>
        </p>
      </div>
      <div class="field input" v-if="!allowEdit">{{ content.quantity }}</div>
    </td>
    <td class="cell productName">
      <span>{{ content.product.name }} ({{ content.product.size }} {{ content.product.getUnit().name }})</span>
      <span v-for="category in content.categories" class="tag category-tag">
        <span class="icon"><icon :icon="category.icon" /></span>
        <span>{{ category.title }}</span>
      </span>
    </td>
    <td class="cell productTax has-text-right"><span>{{ content.tax }}%</span></td>
    <td class="cell productPrice has-text-right"><span>{{ $n(content.price / 100, 'currency') }}</span></td>
    <td class="cell productAmount has-text-right"><span>{{ $n(content.price * content.quantity / 100, 'currency') }}</span></td>
  </tr>
</template>

<style scoped>
.cartProduct{
  margin-bottom:0;
  .grid{
    height:3em;
    gap:0;
    align-items: center;
  }
  td{
    vertical-align: middle;
  }
}

.category-tag{
  margin-left:.5em;
}

.productQuantity{
  .button{
    height:100%;
  }
  .icon{
    cursor:pointer;
    pointer-events: all;
  }
  input{
    text-align: center;
    -moz-appearance: textfield;
    max-width:15ch;
    min-width:12ch;

    &::-webkit-outer-spin-button,
    &::-webkit-inner-spin-button{
      -webkit-appearance: none;
      margin: 0;
    }
  }
}
</style>

<script lang="ts">
import { type CartRow } from '../../composables/cartContent';
import { useCartStore } from '../../store/CartStore';
import type { PropType } from 'vue';

export default{
  data() {
   return {
    cart$: useCartStore()
   }
  },
  props:{
    content: { type: Object as PropType<CartRow>, required: true },
    allowEdit: Boolean,
  },
  methods: {
    setQuant(e: Event){
      const quantity = parseInt((e.target as HTMLInputElement).value);
      this.cart$.setQuantity(this.content.product, isNaN(quantity) ? 0 : quantity);
    },
  }
}
</script>