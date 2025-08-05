import type { Account } from "./account";
import type { Service } from "./service";

export interface ForeignServiceAccount {
  account: Account
  service: Service
  foreignAccountId: string
}