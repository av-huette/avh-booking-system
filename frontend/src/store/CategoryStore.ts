import { defineStore } from 'pinia'
import { type Category, generateTestData } from '../composables/category'

export const useCategoryStore = defineStore('category', {
  state: () => {
    return {
      categorys: [] as Category[]
    }
  },
  actions: {
    generateTestData(){
      if (this.categorys.length == 0) {
        this.categorys.push(...generateTestData());
      }
    }
  }
})