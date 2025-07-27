<template>
  
  <!-- Category Filter -->
  <div class="tags is-flex is-justify-content-space-around">
    <button class="tag is-hoverable" :class="selectedCategory == 0 ? 'is-primary' : ''" @click="selectCategory(0)">All</button>
    <button v-for="category in category$.categorys" class="tag is-hoverable" :class="selectedCategory == category.id ? 'is-primary' : ''" @click="selectCategory(category.id)">{{ category.title }}</button>
  </div>
  <hr>
  <!-- BUTTON VIEW -->
  <div v-if="view=='button'" class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-around is-gap-1">
    <div v-for="account in accounts" @click="account$.select(account)" >
      <button class="button is-fullwidth" :class="account$.selected.includes(account) ? 'is-primary' : ''" title="select account">{{ account.firstName }}</button>
    </div>
  </div>

  <!-- LIST VIEW -->
  <div v-if="view=='list'" class="table-container">
    <table class="table is-striped is-hoverable">
      <tbody>
        <tr>
          <th>ID</th>
          <th>First Name</th>
          <th>Nick Name</th>
          <th>Last Name</th>
          <th>E-Mail</th>
          <th>Phone</th>
          <th>Balance</th>
          <th>MaxDebt</th>
          <th>Category</th>
          <th>Enabled</th>
          <th>Created At</th>
        </tr>
        <tr :class="account$.selected.includes(account) ? 'is-primary' : ''" v-for="account in accounts" @click="account$.select(account)">
          <td>{{ account.id }}</td>
          <td>{{ account.firstName }}</td>
          <td>{{ account.nickName }}</td>
          <td>{{ account.lastName }}</td>
          <td class="has-copy-btn">{{ account.email }} <span class="icon is-small" @click="copyText(account.email)"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-copy-btn">{{ account.phone }} <span class="icon is-small"><icon :icon="['fas', 'copy']" /></span></td>
          <td>{{ account.balance }}</td>
          <td>{{ account.maxDebt }}</td>
          <td>{{ account.category }}</td>
          <td>{{ account.enabled }}</td>
          <td>{{ account.createdAt }}</td>
        </tr>

      </tbody>
    </table>
  </div>
</template>

<style scoped>
  .has-copy-btn{
    position:relative;

    .icon {
      position:absolute;
      left:100%;
      z-index:10;
      padding:.8em;
      background-color:rgba(0,0,0,.3);
      border-radius:var(--bulma-radius);
      visibility:hidden;
      pointer-events: none;
      cursor:pointer;
    }
    &:hover .icon {
      visibility: visible;
      pointer-events: all;
    }
  }
</style>

<script lang="ts">
import { useAccountStore } from '../store/AccountStore';
import { useCategoryStore } from '../store/CategoryStore';

export default {
  props: {
    view: String
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
    },
    copyText(txt: string){
      navigator.clipboard.writeText(txt);
    }
  },
  computed: {
    accounts(){
      return this.account$.getByCategory(this.selectedCategory);
    }
  }
}
</script>