import { defineStore } from 'pinia'
import type { Account } from '../composables/account'

export const useAccountStore = defineStore('account', {
  state: () => {
    return {
      accounts: [] as Account[],
      selected: {} as Account
    }
  },
  getters: {
    fullName(): string {
      //ToDo: Implement Nickname Logic
      return `${this.selected.firstName} ${this.selected.lastName}`;
    }
  }
})