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
  category?: Category
  enabled: boolean
  createdAt?: string
}

export function generateTestData(): Account[]{
  let a1: Account = {
    firstName: "Nomen",
    lastName: "Omen",
    email: "test@test.de",
    balance: Math.random() * 100,
    maxDebt: 0,
    enabled: true
  }

  return [a1];
}