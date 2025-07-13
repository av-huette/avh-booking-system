import { defineStore } from "pinia";
import { type Notification } from "../composables/notification";

export const useSocketStore = defineStore("notificationStore", {
  state: () => ({
    socket: null as WebSocket | null,
    notifications: [] as Notification[],
  }),
  actions: {
    addNotification(msg: any) {
      this.notifications.push(msg);
    },
    sendMessage(msg: any){
      this.socket?.send(msg);
      console.log("Notification send", msg);
    }
  }
})