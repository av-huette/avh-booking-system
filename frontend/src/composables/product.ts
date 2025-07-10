import type { Category } from "./category"
import type { ProductGroup } from "./productGroup"
import type { Unit } from "./unit"

export interface Product {
  id?: number
  name: string
  price: number
  group: ProductGroup
  size: number
  unit: Unit
  tax: number
  category: Category
}