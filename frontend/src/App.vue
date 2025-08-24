<template>
  <DevModeBar v-if="dev"/>
  <MainNavigation />
  <!-- <div v-for="msg in store.notifications"> {{ msg }}</div>
  <button class="button" @click="sendTestMessage">Send Test Message</button> -->
  <RouterView />
</template>

<script lang="ts">
  import { connectWebSocket, closeConnection } from './api/ws';
  import MainNavigation from './components/MainNavigation.vue';
  import DevModeBar from './components/DevModeBar.vue';
  import { useSocketStore } from './store/socketStore';
  
  export default {
    components: {
      MainNavigation,
      DevModeBar
    },
    data() {
      return {
        store: useSocketStore(),
        dev: false
      }
    },
    created() {
      connectWebSocket();
    },
    unmounted() {
      closeConnection();
    },
    mounted() {
        this.dev = import.meta.env.DEV;
    },
    methods: {
      sendTestMessage(){
        this.store.sendMessage('Test Message from Store');
      }
    }
  }
</script>
