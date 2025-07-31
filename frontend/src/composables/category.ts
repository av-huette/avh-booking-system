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
  id: number
}
export class Category {
  public title: string;
  public type: CategoryType;
  public visibility: CategoryVisibility;
  public icon?: Icon;
  public id: number;

  constructor(cat: Category){
    this.title = cat.title;
    this.type = cat.type;
    this.visibility = cat.visibility;
    this.icon = cat.icon;
    this.id = cat.id;
  }
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
    id: 1,
    title: "Rooms", 
    icon: ['fas', 'house'],
    type: 0,
    visibility: 1
  }
  c1 = new Category(c1);

  let c2: Category = {
    id: 2,
    title: "Guests", 
    icon: ['fas', 'beer'],
    type: 0,
    visibility: 1
  }
  c2 = new Category(c2);
  
  return [c1, c2];
}

export function generateTestData(): Category[]{
  return getTestCategorys();
}
