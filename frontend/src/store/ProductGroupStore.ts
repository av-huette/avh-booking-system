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
        },
        // Depth-first ordered groups (excluding "no group" id 0); groups with a missing parent are treated as roots
        tree(): { group: ProductGroup, depth: number }[] {
            const groups = this.productGroups.filter(g => g.id !== 0);
            const ids = new Set(groups.map(g => g.id));
            const byName = (a: ProductGroup, b: ProductGroup) => a.name.localeCompare(b.name);
            const result: { group: ProductGroup, depth: number }[] = [];
            const visited = new Set<number>();

            const visit = (group: ProductGroup, depth: number) => {
                if (visited.has(group.id ?? -1)) return;
                visited.add(group.id ?? -1);
                result.push({ group, depth });
                groups.filter(g => g.parent_id === group.id).sort(byName).forEach(child => visit(child, depth + 1));
            };

            groups.filter(g => g.parent_id == null || !ids.has(g.parent_id)).sort(byName).forEach(root => visit(root, 0));
            return result;
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
        // Returns the IDs of all groups nested below the given group
        descendantIds(id: number): number[] {
            const result: number[] = [];
            const queue = [id];
            while (queue.length > 0) {
                const current = queue.shift();
                this.productGroups
                    .filter(g => g.parent_id === current && g.id !== undefined && !result.includes(g.id))
                    .forEach(g => {
                        result.push(g.id as number);
                        queue.push(g.id as number);
                    });
            }
            return result;
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
