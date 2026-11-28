import {defineStore} from 'pinia'
import { ProductVisibility} from '../composables/productVisibility'
import { useSocketStore } from './socketStore';
import { useSettingStore } from './SettingStore';

export const useProductVisibilityStore = defineStore('product_visibility', {
    state: () => {
        return {
            visibilities: [] as ProductVisibility[]
        }
    },
    getters: {
        all(): ProductVisibility[] {
            return this.visibilities;
        }
    },
    actions: {
        byId(id: number | undefined): ProductVisibility | undefined {
            return this.visibilities.find((vis) => vis.id == id)
        },
        byProductId(id: number){
            let foundVis = this.visibilities.filter((vis) => {
                return vis.productId == id;
            })
            return foundVis;
        },
        // Without a locationId the location of this client is used
        categoryIsVisible(categoryId: number, productId: number, locationId?: number | null): boolean{
            const locId = locationId ?? useSettingStore().locationId;
            return this.byProductId(productId).some((vis) => vis.categoryId == categoryId && vis.locationId == locId);
        },
        // True if the product is visible for at least one category at the given location
        isVisibleAtLocation(productId: number, locationId: number | null): boolean{
            return this.byProductId(productId).some((vis) => vis.locationId == locationId);
        },
        // Categories for which the product is visible (independent of the location)
        categoriesOf(productId: number): number[] {
            return [...new Set(this.byProductId(productId).map((vis) => vis.categoryId))];
        },
        // Locations at which the product is visible (independent of the category)
        locationsOf(productId: number): number[] {
            return [...new Set(this.byProductId(productId).map((vis) => vis.locationId))];
        },
        // Stores the visibility of a product as every combination of the given categories and locations.
        // If one of both lists is empty, the product is not visible anywhere and all rows are removed.
        setVisibility(productId: number, categoryIds: number[], locationIds: number[]) {
            const socket$ = useSocketStore();
            const isWanted = (catId: number, locId: number) => categoryIds.includes(catId) && locationIds.includes(locId);
            const existing = this.byProductId(productId);

            existing
                .filter((vis) => !isWanted(vis.categoryId, vis.locationId))
                .forEach((vis) => socket$.removeVisibility(vis.id!));

            categoryIds.forEach((categoryId) => {
                locationIds.forEach((locationId) => {
                    if (existing.some((vis) => vis.categoryId == categoryId && vis.locationId == locationId)) { return }
                    socket$.addVisibility({categoryId, productId, locationId} as ProductVisibility);
                })
            })
        },
        patchVisibilities(newVisibilities: ProductVisibility[] | ProductVisibility) {
            this.$patch(state => {
                if(Array.isArray(newVisibilities)){
                    // If all visibilities are queried
                    state.visibilities = [];
                    newVisibilities.forEach((newVis) => {
                        const newVisObj = new ProductVisibility(newVis);
                        state.visibilities.push(newVisObj);
                    })
                } else {
                    // if only one visibility was queried
                    const newVisObj = new ProductVisibility(newVisibilities);
                    state.visibilities.push(newVisObj);
                }
            })
        }
    }
})