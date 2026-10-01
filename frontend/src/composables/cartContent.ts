import { Product } from "./product"
import type { Account } from "./account"
import type { Category } from "./category"

export interface CartContent {
  product: Product
  account: Account
  quantity: number
  price: number
  tax: number
}
//ToDo: Save Tax like Price individually

// Display row of the cart: all CartContents of one product, combined over the accounts
export interface CartRow {
  product: Product
  quantity: number
  price: number
  tax: number
  categories?: Category[] // only set if the row is not for all accounts
}

export function aggregateCartContents(contents: CartContent[], accounts: Account[]): CartRow[] {
  const rows: CartRow[] = [];
  const rowAccounts = new Map<CartRow, Account[]>();

  contents.forEach((cont) => {
    let row = rows.find((r) => r.product.id == cont.product.id);
    if (!row) {
      row = { product: cont.product, quantity: cont.quantity, price: cont.price, tax: cont.tax };
      rows.push(row);
      rowAccounts.set(row, []);
    }
    rowAccounts.get(row)!.push(cont.account);
  });

  rows.forEach((row) => {
    const accs = rowAccounts.get(row)!;
    const accIds = accs.map((acc) => acc.id);
    if (accounts.every((acc) => accIds.includes(acc.id))) { return }
    const categories: Category[] = [];
    accs.forEach((acc) => {
      const cat = acc.getCategory();
      if (cat && !categories.includes(cat)) { categories.push(cat) }
    });
    row.categories = categories;
  });

  return rows;
}
