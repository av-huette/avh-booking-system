import {defineStore} from 'pinia'
import {ProductGroup} from '../composables/productGroup'

export const useProductGroupStore = defineStore('productGroup', {
    state: () => {
        return {
            productGroups: [] as ProductGroup[]
        }
    },
    getters: {
        all(): ProductGroup[] {
            return this.productGroups;
        }
    },
    actions: {
        byId(id: number | undefined): ProductGroup | undefined {
            return this.productGroups.find((productGroup) => productGroup.id == id)
        },
        // Returns the chain from the root group down to the given group
        path(id: number | undefined): ProductGroup[] {
            const path: ProductGroup[] = [];
            const visited = new Set<number>();
            let group = this.byId(id);
            while (group && !visited.has(group.id ?? -1)) {
                visited.add(group.id ?? -1);
                path.unshift(group);
                group = group.parent_id != null ? this.byId(group.parent_id) : undefined;
            }
            return path;
        },
        removeById(id: number) {
            this.$patch(state => {
                state.productGroups = state.productGroups.filter(g => g.id !== id)
            })
        },
        patchProductGroups(newGroups: ProductGroup[]) {
            this.$patch(state => {
                newGroups.forEach(newGroup => {
                    const newGroupObject = new ProductGroup(newGroup);
                    const existing = state.productGroups.find(a => a.id === newGroupObject.id);
                    if (existing) {
                        Object.assign(existing, newGroupObject);
                    } else {
                        state.productGroups.push(newGroupObject);
                    }
                })
            })
        }
    }
})
