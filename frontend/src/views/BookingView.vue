<template>
  <section class="section">
    <!-- ToDo: Create own Component for Cart Info-->
    <div class="cart message is-info">
      <div class="message-header">
        Cart
        <button v-if="account$.selected.length > 0" class="delete" @click="account$.unselect()" title="discard cart and unselect account"></button>
      </div>
      <div class="message-body">
        <p v-if="account$.selected.length > 0">This is the cart of {{ account$.selected[0]?.firstName }}</p>
      </div>
    </div>

    <div class="tabs is-boxed is-centered">
      <ul>
        <li :class="visiblePart == 0 ? 'is-active' : ''">
          <a @click="setVisiblePart(0)">
            <span class="icon is-small"><icon :icon="['fas', 'user']" /></span>
            Accounts
          </a>
        </li>
        <li :class="visiblePart == 1 ? 'is-active' : ''">
          <a @click="setVisiblePart(1)">
            <span class="icon is-small"><icon :icon="['fas', 'th-large']" /></span>
            Items
          </a>
        </li>
      </ul>
    </div>

    <div v-if="visiblePart==0" class="accounts">
      <AccountSelector view="button"/>
    </div>

    <div v-if="visiblePart==1" class="accounts">
      <!-- ToDo: Develop Item Selector -->
      HERE ARE ITEMS
    </div>
  </section>
</template>

<style scoped>
.section{
  padding-block: 0;
}
</style>

<script lang="ts">
import AccountSelector from '../components/AccountSelector.vue';
import { useAccountStore } from '../store/AccountStore';

  export default {
    components: {
      AccountSelector
    },
    data() {
      return {
        visiblePart: 0 as number,
        account$: useAccountStore()
      }  
    },
    methods: {
      setVisiblePart(partNr: number){
        this.visiblePart = partNr;
      }
    }
  }
</script>