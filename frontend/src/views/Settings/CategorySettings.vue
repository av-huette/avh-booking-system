<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCategoryStore } from '../../store/CategoryStore'
import { CategoryType, CategoryTypeToString, StringToCategoryType, type Category } from '../../composables/category'
import { useSocketStore } from '../../store/socketStore'
import ToggleSwitch from '../../composables/elements/ToggleSwitch.vue'
import ErrorModal from '../../components/ErrorModal.vue'
import SettingsLayout from '../../components/SettingsLayout.vue'

const route = useRoute()
const router = useRouter()
const category$ = useCategoryStore()
const socket$ = useSocketStore()

const categoryType = computed(() => StringToCategoryType[route.params.type as string])
const isAccount = computed(() => categoryType.value === CategoryType.ACCOUNT)
const typeLabel = computed(() => isAccount.value ? 'Account' : 'Produkt')
const categories = computed(() => isAccount.value ? category$.accountCategories : category$.productCategories)

const pendingIds = ref<number[]>([])
const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

function editCategory(categoryId: number, type: string) {
  router.push({ name: 'CategorySettingsSingleEdit', params: { type, categoryId } })
}

function toggleEnabled(cat: Category) {
  if (pendingIds.value.includes(cat.id)) return
  pendingIds.value.push(cat.id)
  socket$.toggleCategoryEnabled(cat.id, !cat.enabled)
}

function onMutationResult(res: any) {
  if (res.table !== 'category') return
  const idx = pendingIds.value.indexOf(res.id)
  if (idx !== -1) pendingIds.value.splice(idx, 1)
}

function onWsError(err: any) {
  if (pendingIds.value.length === 0) return
  pendingIds.value = []
  currentError.value = err
  errorModalVisible.value = true
}

onMounted(() => {
  socket$.wsClient.on('mutationResult', onMutationResult)
  socket$.wsClient.on('wsError', onWsError)
})

onBeforeUnmount(() => {
  socket$.wsClient.off('mutationResult', onMutationResult)
  socket$.wsClient.off('wsError', onWsError)
})
</script>

<template>
<SettingsLayout :area="isAccount ? 'account' : 'product'">
  <h1 class="title">{{ typeLabel }}-Kategorien</h1>

  <div class="panel">
    <p class="panel-heading has-text-primary-dark">{{ typeLabel }}-Kategorien</p>

    <div class="panel-block">

      <table class="table is-fullwidth is-striped is-hoverable">
        <thead>
          <tr>
            <th>Aktiv</th>
            <th>Kategorie</th>
            <th class="has-text-right">Edit</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="cat in categories" :key="cat.id">
            <td>
              <div>
                <ToggleSwitch
                  :model-value="cat.enabled"
                  :disabled="pendingIds.includes(cat.id)"
                  @update:model-value="toggleEnabled(cat)"
                />
                <span v-if="pendingIds.includes(cat.id)" class="icon has-text-grey ml-2">
                  <icon :icon="['fas', 'spinner']" :spin="true" />
                </span>
              </div>
            </td>
            <td>
              <span class="tag is-medium">
                <span class="icon is-small"><icon :icon="cat.icon" /></span>
                <span>{{ cat.title }}</span>
              </span>
            </td>
            <td class="has-text-right">
              <button class="button is-small" @click="editCategory(cat.id, CategoryTypeToString[cat.type])">
                <span class="icon is-small"><icon :icon="['fas', 'pen']" /></span>
              </button>
            </td>
          </tr>
          <tr v-if="categories.length === 0">
            <td colspan="3" class="has-text-grey has-text-centered is-italic">
              Keine {{ typeLabel }}-Kategorien vorhanden
            </td>
          </tr>
          <tr>
            <td colspan="3" class="has-text-centered">
              <button class="button" @click="router.push({ name: 'CategorySettingsAdd', params: { type: CategoryTypeToString[categoryType] } })">
                <span class="icon"><icon :icon="['fas', 'plus']" /></span>
                <span>{{ typeLabel }} Kategorie hinzufügen</span>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

  <ErrorModal v-model="errorModalVisible" :error="currentError" />
</SettingsLayout>
</template>

<style scoped>
td {
  vertical-align: middle;
}
</style>
