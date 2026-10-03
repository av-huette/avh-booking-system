<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useVatStore } from '../../store/VatStore'
import { useSocketStore } from '../../store/socketStore'
import Buttons from '../../composables/elements/Buttons.vue'
import Button from '../../composables/elements/Button.vue'
import ErrorModal from '../../components/ErrorModal.vue'

const route = useRoute()
const router = useRouter()
const vat$ = useVatStore()
const socket$ = useSocketStore()

const isEdit = computed(() => !!route.params.vatId)
const editId = computed(() => isEdit.value ? parseInt(route.params.vatId as string) : null)

// v-model.number yields '' for an empty input
const rate = ref<number | ''>('')
const rateIsValid = computed(() => typeof rate.value === 'number' && Number.isInteger(rate.value) && rate.value >= 0)

const saveStatus = ref<'idle' | 'pending' | 'success' | 'error'>('idle')
const errorModalVisible = ref(false)
const currentError = ref<{ code: string, message: string, details?: string } | null>(null)

let saveStartTime = 0
let saveTimeoutId: number | null = null
let mutationHandler: ((res: any) => void) | null = null
let wsErrorHandler: ((err: any) => void) | null = null

onMounted(() => {
  if (isEdit.value && editId.value !== null) {
    const vat = vat$.byId(editId.value)
    if (vat) rate.value = vat.rate
  }
})

onBeforeUnmount(() => {
  cleanupSaveListeners()
})

function cleanupSaveListeners() {
  if (mutationHandler) {
    socket$.wsClient.off('mutationResult', mutationHandler)
    mutationHandler = null
  }
  if (wsErrorHandler) {
    socket$.wsClient.off('wsError', wsErrorHandler)
    wsErrorHandler = null
  }
  if (saveTimeoutId !== null) {
    window.clearTimeout(saveTimeoutId)
    saveTimeoutId = null
  }
}

function save() {
  if (!rateIsValid.value) return
  const newRate = rate.value as number

  saveStatus.value = 'pending'
  saveStartTime = Date.now()

  mutationHandler = (res: any) => {
    if (res.table !== 'vat') return
    if (isEdit.value && res.id !== editId.value) return
    cleanupSaveListeners()

    const elapsed = Date.now() - saveStartTime
    window.setTimeout(() => {
      saveStatus.value = 'success'
      window.setTimeout(() => {
        router.push({ name: 'VatSettings' })
      }, 500)
    }, Math.max(0, 500 - elapsed))
  }

  wsErrorHandler = (err: any) => {
    cleanupSaveListeners()
    saveStatus.value = 'idle'
    currentError.value = err
    errorModalVisible.value = true
  }

  socket$.wsClient.on('mutationResult', mutationHandler)
  socket$.wsClient.on('wsError', wsErrorHandler)

  saveTimeoutId = window.setTimeout(() => {
    cleanupSaveListeners()
    saveStatus.value = 'error'
    window.setTimeout(() => { saveStatus.value = 'idle' }, 500)
  }, 5000)

  if (isEdit.value && editId.value !== null) {
    socket$.updateVat(editId.value, newRate)
  } else {
    socket$.addVat(newRate)
  }
}
</script>

<template>
  <h1 class="title" v-if="isEdit">Steuersatz bearbeiten</h1>
  <h1 class="title" v-else>Neuen Steuersatz erstellen</h1>

  <div class="columns">
    <div class="column is-3">Steuersatz (%):</div>
    <div class="column is-3">
      <p class="control has-icons-left">
        <input type="number" class="input" v-model.number="rate" step="1" min="0" placeholder="z.B. 19">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'percent']" />
        </span>
      </p>
    </div>
  </div>

  <hr class="divider">
  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <div class="is-flex is-align-items-center">
        <Buttons>
          <Button :fa-icon="['fas', 'times']" icon-position="left" @click="router.push({ name: 'VatSettings' })">
            Cancel
          </Button>
          <Button
            class="is-primary"
            @click="save"
            :fa-icon="['fas', 'save']"
            icon-position="right"
            :disabled="saveStatus === 'pending' || !rateIsValid"
          >
            {{ isEdit ? 'Speichern' : 'Erstellen' }}
          </Button>
        </Buttons>
        <span v-if="saveStatus === 'pending'" class="ml-3 icon has-text-grey">
          <icon :icon="['fas', 'spinner']" :spin="true" />
        </span>
        <span v-else-if="saveStatus === 'success'" class="ml-3 icon-text has-text-success">
          <span class="icon"><icon :icon="['fas', 'check']" /></span>
          <span>Saved</span>
        </span>
        <span v-else-if="saveStatus === 'error'" class="ml-3 icon-text has-text-danger">
          <span class="icon"><icon :icon="['fas', 'times']" /></span>
          <span>Saving failed</span>
        </span>
      </div>
    </div>
  </div>

  <ErrorModal v-model="errorModalVisible" :error="currentError" />
</template>

<style scoped>
.columns {
  align-items: center;
}
</style>
