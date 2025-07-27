import { defineStore } from 'pinia'
import { type Account, generateTestData } from '../composables/account'

export const useAccountStore = defineStore('account', {
  state: () => {
    return {
      accounts: [] as Account[],
      selected: [] as Account[]
    }
  },
  actions: {
    generateTestData(){
      this.accounts.push(...generateTestData());
    },
    select(acc: Account){
      this.selected = [acc];
    },
    unselect(){
      this.selected = [];
    },
    getByCategory(categoryId: number): Account[]{
      const enabledUsers = this.accounts.filter((acc) => acc.enabled);
      if(categoryId == 0){
        return enabledUsers;
      }
      return enabledUsers.filter((acc) => acc.category == categoryId );
    }
  }
})