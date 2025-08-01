import { type Account } from './account'

export interface Order {
  id?: number
  account: Account
  createdAt?: string
}