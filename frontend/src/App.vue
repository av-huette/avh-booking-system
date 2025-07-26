<template>
  <MainNavigation />
  <div v-for="msg in store.notifications"> {{ msg }}</div>
  <button class="button" @click="sendTestMessage">Send Test Message</button>
  <RouterView />
</template>

<script lang="ts">
  import { connectWebSocket, closeConnection } from './api/ws';
  import MainNavigation from './components/MainNavigation.vue';
  import { useSocketStore } from './store/socketStore';
  
  export default {
    components: {
      MainNavigation
    },
    data() {
      return {
        store: useSocketStore()
      }
    },
    created() {
      connectWebSocket();
    },
    unmounted() {
      closeConnection();
    },
    methods: {
      sendTestMessage(){
        this.store.sendMessage('Test Message from Store');
      }
    }
  }
</script>
