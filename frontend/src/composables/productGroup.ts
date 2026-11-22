export interface ProductGroup {
    id?: number
    name: string
    parentId?: number | null
}

export class ProductGroup implements ProductGroup {

    constructor(group: ProductGroup) {
        this.id = group.id;
        this.name = group.name;
        this.parentId = group.parentId ?? null;
    }
}
