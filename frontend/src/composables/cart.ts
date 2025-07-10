import { type Account } from './account'

export interface Cart {
  id?: number
  account: Account
  createdAt?: string
}