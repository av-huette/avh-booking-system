import { type Category } from './category'
export interface Account {
  id?: number
  firstName: string
  lastName: string
  nickName?: string
  email: string
  phone?: string
  balance: number
  maxDebt: number
  category: Category
  enabled: boolean
  createdAt?: string
}