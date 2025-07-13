<script lang="ts">
  import { mapState } from 'pinia';
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
    computed: {
      ...mapState(useSocketStore, ['socket', 'notifications'])
    },
    methods: {
      sendTestMessage(){
        this.store.sendMessage('Test Message from Store');
        this.store.socket?.send('Test Message from Socket in Store');
        this.socket?.send('Test Message from Computed Store variable ');
      }
    }
  }
</script>

<template>
  <MainNavigation />
  <div v-for="msg in notifications"> {{ msg }}</div>
  <button class="button" @click="sendTestMessage">Send Test Message</button>
  <RouterView />
</template>

<style>

</style>
