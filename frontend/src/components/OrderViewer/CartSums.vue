<script lang="ts" setup>
import type { BookingTotals } from '../../composables/booking';
import type { CategoryTotals } from '../../store/CartStore';

const props = defineProps<{
  totals: BookingTotals,
  categoryTotals?: CategoryTotals[]
}>()
</script>


<template>
<template v-if="props.categoryTotals && props.categoryTotals.length > 1">
  <div class="categorySum" v-for="catTotals in props.categoryTotals">
    <p class="cartSum">
      <span class="tag category-tag">
        <span class="icon"><icon :icon="catTotals.category?.icon" /></span>
        <span>{{ catTotals.category?.title }}</span>
      </span>
      <span class="accountCount">{{ catTotals.accountCount }} × {{ $t('booking.accounts') }}</span>
      {{ $t('transaction.sum') }}:
      <span>{{ $n(catTotals.totals[0] / 100, 'currency') }}</span>
    </p>
    <p class="cartTax">
      {{ $t('transaction.partVat') }}:
      <span>{{ $n(catTotals.totals[1] / 100, 'currency') }}</span>
    </p>
  </div>
</template>
<template v-else>
<p class="cartSum">
  {{ $t('transaction.sum') }}:
  <span>{{ $n(totals[0] / 100, 'currency') }}</span>
</p>
<p class="cartTax">
  {{ $t('transaction.partVat') }}:
  <span>{{ $n(totals[1] / 100, 'currency') }}</span>
</p>
</template>
</template>

<style scoped>
.categorySum:not(:last-child){
  margin-bottom:.5rem;
}
.cartSum .category-tag,
.cartSum .category-tag span,
.cartSum .accountCount{
  font-size:.8rem;
  font-weight:400;
  margin-right:.5em;
}
.cartSum{
  font-size:1.3rem;
  text-align: right;
  color: hsl(var(--bulma-text-h), var(--bulma-text-s), var(--bulma-text-strong-l));
  span{
    font-weight:700;
    /* color: hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l)); */
  }
}
.cartTax{
  font-size:.8rem;
  text-align: right;
  /* opacity: 0.65; */
  span{
    font-weight:600;
  }
}
</style>

