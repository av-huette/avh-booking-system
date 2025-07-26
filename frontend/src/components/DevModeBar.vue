<template>
<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link class="navbar-item" to="/">
      DEVELOPEMENT
    </router-link>
    
    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" @click="toggleBurger">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
    </a>
  </div>

  <div class="navbar-menu" :class="{'is-active': burgerActive}">
    <div class="navbar-start">

      <a class="navbar-item" @click="generateTestData()">Generate Test Data</a>
    </div>

    <div class="navbar-end">
      <span>This is not visible in production</span>
    </div>
  </div>
</nav>
</template>

<script lang="ts">
import { useAccountStore } from '../store/AccountStore';

  export default {
    data() {
      return {
        account$: useAccountStore(),
        burgerActive: false as Boolean
      }
    },
    methods: {
      toggleBurger(){
        this.burgerActive = !this.burgerActive;
      },
      generateTestData(){
        this.account$.generateTestData();
      }
    }
  }
</script>