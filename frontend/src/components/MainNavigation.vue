<template>
<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link class="navbar-item" to="/">
      LOGO
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
      <router-link class="navbar-item" to="/">
        Home
      </router-link>

      <router-link class="navbar-item" to="/test">
        Test
      </router-link>

      <a class="navbar-item">Statistics</a>

      <div class="navbar-item has-dropdown is-hoverable">
        <a class="navbar-link">Settings</a>

        <div class="navbar-dropdown">
          <a class="navbar-item">Wareneinstellungen</a>
          <a class="navbar-item">Benutzereinstellungen</a>
          <a class="navbar-item">Zahlungseinstellungen</a>
          <a class="navbar-item">Systemeinstellungen</a>
        </div>
      </div>
    </div>

    <div class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
          <a class="button is-light">
            Log in
          </a>
        </div>
      </div>
    </div>
  </div>
</nav>
</template>

<script lang="ts">
  export default {
    data() {
      return {
        burgerActive: false as Boolean
      }
    },
    methods: {
      toggleBurger(){
        this.burgerActive = !this.burgerActive;
      }
    }
  }
</script>