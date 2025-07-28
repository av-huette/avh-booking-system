<template>
  
  <!-- Category Filter -->
  <div class="tags is-flex is-justify-content-space-around">
    <button class="tag is-hoverable" :class="selectedCategory == 0 ? 'is-primary' : ''" @click="selectCategory(0)">All</button>
    <button v-for="category in category$.categorys" class="tag is-hoverable" :class="selectedCategory == category.id ? 'is-primary' : ''" @click="selectCategory(category.id)">{{ category.title }}</button>
  </div>
  <hr>

  <ShowAccountButton v-if="show=='button'" :accounts="accounts" />
  <ShowAccountList v-if="show=='list'" :accounts="accounts" />
</template>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCategoryStore } from '../../store/CategoryStore';
import ShowAccountButton from './ShowAccountButton.vue';
import ShowAccountList from './ShowAccountList.vue';

export default {
  components: {
    ShowAccountButton,
    ShowAccountList
  },
  props: {
    show: String
  },
  data() {
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      selectedCategory: 0 as number,
    }
  },
  methods: {
    selectCategory(categoryId: number){
      this.selectedCategory = categoryId;
    }
  },
  computed: {
    accounts(){
      return this.account$.getByCategory(this.selectedCategory);
    }
  }
}
</script>