import { defineStore } from 'pinia'
import { Product, generateTestData } from '../composables/product'

export const useProductStore = defineStore('product', {
  state: () => {
    return {
      products: [] as Product[]
    }
  },
  actions: {
    generateTestData(){
      let testData = generateTestData();
      this.products.push(...testData);
      console.log(this.products);
    },
    getByCategory(categoryId: number, selectedAccountCategorys: number[]): Product[]{
      let visibleProducts = this.products.filter(prod =>
        prod.visibility.some(cat => selectedAccountCategorys.includes(cat))
      );

      console.log("Visible Products: ", visibleProducts);
      const currentCategoryProducts = visibleProducts.filter((prod) => prod.category == categoryId);
      if(categoryId == 0){
        return visibleProducts;
      }
      return currentCategoryProducts;
    },
    getBySearchAndCategory(searchString: string, categoryId: number, selectedAccountCategorys: number[]): Product[]{
      let search = searchString.toLowerCase();
      let byCategory = this.getByCategory(categoryId, selectedAccountCategorys);
      let searchResults = byCategory.filter((prod) => prod.name.toLowerCase().includes(search));
      return searchResults;
    }
  }
})