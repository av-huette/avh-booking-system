import { defineStore } from 'pinia'
import { type CartContent } from '../composables/cartContent'
import { Product } from '../composables/product'
import { useAccountStore } from './AccountStore'
import type { BookingTotals } from '../composables/booking'
import type { Account } from '../composables/account'
import type { Category } from '../composables/category'

export interface CategoryTotals {
  category: Category | undefined
  accountCount: number
  totals: BookingTotals
}

function calcTotals(contents: CartContent[]): BookingTotals {
  let total = 0;
  let tax = 0;
  contents.forEach((cont) => {
    let subTotal = cont.product.price * cont.quantity;
    let subTax = subTotal*(cont.tax / 100);
    total += subTotal;
    tax += subTax;
  })
  return [total, tax] as BookingTotals;
}

export const useCartStore = defineStore('cart', {
  state: () => {
    return {
      cartContents: [] as CartContent[]
    }
  },
  getters:{
    totalsForAccount(): (account: Account) => BookingTotals {
      return (account: Account) => calcTotals(this.cartContents.filter((cont) => cont.account.id == account.id));
    },
    // Amount per account (of the first selected account)
    getTotals(): BookingTotals{
      const account$ = useAccountStore();
      if (account$.selected.length == 0) { return [0, 0] as BookingTotals }
      return this.totalsForAccount(account$.selected[0]);
    },
    // All accounts of one category have the same cart contents, so one account represents the category
    totalsByCategory(): CategoryTotals[] {
      const account$ = useAccountStore();
      const result: CategoryTotals[] = [];
      const seenCategories: (number | undefined)[] = [];
      account$.selected.forEach((acc) => {
        if (seenCategories.includes(acc.category)) { return }
        seenCategories.push(acc.category);
        result.push({
          category: acc.getCategory(),
          accountCount: account$.selected.filter((a) => a.category == acc.category).length,
          totals: this.totalsForAccount(acc),
        });
      })
      return result;
    },
    isOverdrawn(): Boolean{
      const account$ = useAccountStore();
      if (account$.selected.length > 1) {return false}
      return account$.selected[0].balance - this.totalsForAccount(account$.selected[0])[0] < (account$.selected[0].maxDebt * -1)
    }
  },
  actions: {
    addToCart(product: Product, accounts: Account[] = useAccountStore().selected){
      accounts.forEach((account) => {
        let alreadySelected = this.cartContents.filter((cont) => cont.product.id == product.id && cont.account.id == account.id);
        if (alreadySelected.length > 0){
          alreadySelected[0].quantity ++;
          return;
        }
        let newCartContent = {product: product, account: account, quantity: 1, price: product.price, tax: product.getVat()?.rate} as CartContent;
        this.cartContents.push(newCartContent);
      })
    },
    removeFromCart(product: Product){
      this.cartContents = this.cartContents.filter((cont) => {
        return cont.product != product
      })
    },
    setQuantity(product: Product, quantity: number){
      if (quantity <= 0) {
        this.removeFromCart(product);
        return;
      }
      this.cartContents.forEach((cont) => {
        if (cont.product == product) { cont.quantity = quantity }
      })
    },
    changeQuantity(product: Product, delta: number){
      const current = this.productCartQuantity(product);
      if (current == -1) { return }
      this.setQuantity(product, current + delta);
    },
    removeAccount(account: Account){
      this.cartContents = this.cartContents.filter((cont) => cont.account.id != account.id);
    },
    productCartQuantity(product: Product): number{
      let result = -1;
      this.cartContents.forEach((cont) => {
        if (cont.product != product){
          return;
        }
        result = cont.quantity;
      })
      return result;
    }
  }
})
