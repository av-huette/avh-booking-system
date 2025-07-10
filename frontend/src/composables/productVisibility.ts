import type { Category } from "./category"
import type { Product } from "./product"

export interface ProductVisibility {
  id?: number
  category: Category
  product: Product
}