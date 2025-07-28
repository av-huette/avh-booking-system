<template>
  <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-around is-gap-1">
    <div v-for="account in accounts" @mousemove="changeSelectMode" @click="selectAccount($event, account)" >
      <button class="button is-fullwidth" :class="account$.selected.includes(account) ? 'is-primary' : ''" title="select account">{{ account.firstName }}</button>
    </div>
  </div>
</template>

<script lang="ts">
import type { Account } from '../../composables/account';
import { useAccountStore } from '../../store/AccountStore';
import type { PropType } from 'vue';

export default {
  data(){
    return {
      account$: useAccountStore(),
      selectMode: "single"
    }
  },
  props: {
    accounts: {
      type: Array as PropType<Account[]>
    }
  },
  methods: {
    selectAccount(e: MouseEvent, account: Account) {
      this.changeSelectMode(e);
      if(this.selectMode == "single") {
        this.account$.select(account);
        return;
      }
      if(this.selectMode == "multiple") {
        this.account$.selectAdd(account);
      }
    },
    changeSelectMode(e: MouseEvent){
      if(e.ctrlKey){
        this.selectMode = "multiple";
        (e.target as HTMLElement).style.cursor = "copy";
        return;
      }
      this.selectMode = "single";
      (e.target as HTMLElement).style.cursor = "";
    }
  }
}
</script>