<template>
  <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-around is-align-content-flex-start is-gap-1 accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">
    <div v-for="account in accounts" @mousemove="changeSelectMode" @click="selectAccount($event, account)" >
      <button class="button is-fullwidth" :class="account$.selected.includes(account) ? 'is-primary' : ''" title="select account">{{ account.getShortName() }}</button>
    </div>
  </div>
</template>

<style scoped>
.accountList{
  height: var(--_height);
  overflow-y: scroll;
}
</style>

<script lang="ts">
import { Account } from '../../composables/account';
import { useAccountStore } from '../../store/AccountStore';
import type { PropType } from 'vue';
import { useResizeObserver } from '@vueuse/core';

export default {
  data(){
    return {
      account$: useAccountStore(),
      selectMode: "single",
      height: 0,
      resizeElement: {} as HTMLElement
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
    },
    onResize(){
      console.log("resize detected", this.resizeElement.getBoundingClientRect().top);
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