<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useVatStore } from '../../store/VatStore'
import SettingsLayout from '../../components/SettingsLayout.vue'

const router = useRouter()
const vat$ = useVatStore()
</script>

<template>
<SettingsLayout area="payment">
  <h1 class="title">Steuersätze</h1>

  <div class="panel">
    <p class="panel-heading has-text-primary-dark">Steuersätze</p>
    <div class="panel-block">
      <table class="table is-fullwidth is-striped is-hoverable">
        <thead>
          <tr>
            <th>Steuersatz</th>
            <th class="has-text-right">Edit</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="vat in vat$.all" :key="vat.id">
            <td>{{ vat.rate }} %</td>
            <td class="has-text-right">
              <button
                class="button is-small"
                @click="router.push({ name: 'VatSettingsEdit', params: { vatId: vat.id } })">
                <span class="icon is-small"><icon :icon="['fas', 'pen']" /></span>
              </button>
            </td>
          </tr>
          <tr v-if="vat$.all.length === 0">
            <td colspan="2" class="has-text-grey has-text-centered is-italic">
              Keine Steuersätze vorhanden
            </td>
          </tr>
          <tr>
            <td colspan="2" class="has-text-centered">
              <button class="button" @click="router.push({ name: 'VatSettingsAdd' })">
                <span class="icon"><icon :icon="['fas', 'plus']" /></span>
                <span>Steuersatz hinzufügen</span>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</SettingsLayout>
</template>

<style scoped>
td {
  vertical-align: middle;
}
</style>
