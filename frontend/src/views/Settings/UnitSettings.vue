<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUnitStore } from '../../store/UnitStore'
import SettingsLayout from '../../components/SettingsLayout.vue'

const router = useRouter()
const unit$ = useUnitStore()
</script>

<template>
<SettingsLayout area="product">
  <h1 class="title">Einheiten</h1>

  <div class="panel">
    <p class="panel-heading has-text-primary-dark">Einheiten</p>
    <div class="panel-block">
      <table class="table is-fullwidth is-striped is-hoverable">
        <thead>
          <tr>
            <th>Name</th>
            <th class="has-text-right">Edit</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="unit in unit$.all" :key="unit.id">
            <td>{{ unit.name }}</td>
            <td class="has-text-right">
              <button
                class="button is-small"
                @click="router.push({ name: 'UnitSettingsEdit', params: { unitId: unit.id } })">
                <span class="icon is-small"><icon :icon="['fas', 'pen']" /></span>
              </button>
            </td>
          </tr>
          <tr v-if="unit$.all.length === 0">
            <td colspan="2" class="has-text-grey has-text-centered is-italic">
              Keine Einheiten vorhanden
            </td>
          </tr>
          <tr>
            <td colspan="2" class="has-text-centered">
              <button class="button" @click="router.push({ name: 'UnitSettingsAdd' })">
                <span class="icon"><icon :icon="['fas', 'plus']" /></span>
                <span>Einheit hinzufügen</span>
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
