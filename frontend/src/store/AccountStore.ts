import { defineStore } from 'pinia'
import { Account, generateTestData } from '../composables/account'

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
    selectAdd(acc: Account){
      if(this.selected.includes(acc)){ return; }
      this.selected.push(acc);
    },
    selectSubstract(acc: Account){
      this.selected = this.selected.filter((a) => a != acc);
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
    },
    getBySearchAndCategory(searchString: string, categoryId: number): Account[]{
      let search = searchString.toLowerCase();
      let byCategory = this.getByCategory(categoryId);
      let searchResults = byCategory.filter((acc) => {
        return (
          acc.firstName.toLowerCase().includes(search) ||
          acc.nickName?.toLowerCase().includes(search) ||
          acc.lastName.toLowerCase().includes(search)
        )
      })

      return searchResults
    }
  }
})