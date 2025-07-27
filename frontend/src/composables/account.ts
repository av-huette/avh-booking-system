export interface Account {
  id?: number
  firstName: string
  lastName: string
  nickName?: string
  email: string
  phone?: string
  balance: number
  maxDebt: number
  category?: number
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
    enabled: true,
    category: 1
  }

  let a2: Account = {
    firstName: "Peter",
    lastName: "Super",
    email: "test@test.de",
    balance: Math.random() * 100,
    maxDebt: 0,
    enabled: true,
    category: 2
  }

  return [a1, a2];
}