export interface ProductGroup {
    id?: number
    name: string
    parent_id?: number | null
}

export class ProductGroup implements ProductGroup {

    constructor(group: ProductGroup) {
        this.id = group.id;
        this.name = group.name;
        this.parent_id = group.parent_id ?? null;
    }
}
