import { ref } from "vue";
import { getApi } from "../api/api";
type Icon = [string, string];

/**
 * Type definitions
 */

const categoryTypes = {
  ACCOUNT: 0,
  PRODUCT: 1
} as const;
export type CategoryType = (typeof categoryTypes)[keyof typeof categoryTypes];

const categoryVisibilitys = {
  HIDDEN: 0,
  SHOWN: 1
} as const;
export type CategoryVisibility = (typeof categoryVisibilitys)[keyof typeof categoryVisibilitys];

export interface Category {
  title: string
  type: CategoryType
  visibility: CategoryVisibility
  icon?: Icon
  id?: number
}

/**
 * Funcitons
 */

export async function getCategorys(){
  let {response: categorys, request} =  getApi<Category[]>(
    "/category"
  );
  
  let loaded = ref(false);
  if (loaded.value === false){
    await request();
    loaded.value = true;
  }

  return { categorys };
}

export function getTestCategorys(): Category[]{
  let c1: Category = {
    title: "Test Category", 
    icon: ['fas', 'house'],
    type: 0,
    visibility: 1
  }
  let c2: Category = {
    title: "Test Category 2", 
    icon: ['fas', 'beer'],
    type: 0,
    visibility: 1
  }

  return [c1, c2];
}

export function generateTestData(): Category[]{
  return getTestCategorys();
}
