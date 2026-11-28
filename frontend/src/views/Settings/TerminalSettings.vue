<script setup lang="ts">
import { ref, computed } from 'vue';
import SettingsLayout from '../../components/SettingsLayout.vue';
import Buttons from '../../composables/elements/Buttons.vue';
import Button from '../../composables/elements/Button.vue';
import { useLocationStore } from '../../store/LocationStore';

interface StripeTerminal {
  id: string        // Stripe reader ID (tmr_...)
  label: string
  deviceType: string
  status: 'online' | 'offline'
}

const location$ = useLocationStore();

// Dummy: The terminals have to be read from Stripe by the backend.
// ToDo: query the terminals via WS once the message type is agreed on with the backend
const terminals = ref<StripeTerminal[]>([
  { id: 'tmr_DUMMY0001', label: 'Theke', deviceType: 'bbpos_wisepos_e', status: 'online' },
  { id: 'tmr_DUMMY0002', label: 'Mobil', deviceType: 'stripe_m2', status: 'offline' },
]);

// Assignment terminal ID -> location ID (null = not used). One terminal per location.
function assignmentsFromLocations(): Record<string, number | null> {
  const result: Record<string, number | null> = {};
  terminals.value.forEach((terminal) => {
    result[terminal.id] = location$.all.find((loc) => loc.stripeTerminalId == terminal.id)?.id ?? null;
  });
  return result;
}

const saved = ref(assignmentsFromLocations());
const draft = ref({ ...saved.value });
const saveStatus = ref<'idle' | 'pending' | 'success'>('idle');

const isDirty = computed(() =>
  terminals.value.some((terminal) => draft.value[terminal.id] !== saved.value[terminal.id])
);

// A location can only be assigned to one terminal
function isTakenByOtherTerminal(locationId: number, terminalId: string): boolean {
  return Object.entries(draft.value).some(([id, locId]) => id != terminalId && locId == locationId);
}

function save() {
  saveStatus.value = 'pending';
  // ToDo: send an update mutation for each changed location (stripeTerminalId) once the backend supports it
  window.setTimeout(() => {
    saved.value = { ...draft.value };
    saveStatus.value = 'success';
    window.setTimeout(() => { saveStatus.value = 'idle' }, 1500);
  }, 500);
}

function reset() {
  draft.value = { ...saved.value };
}
</script>

<template>
<SettingsLayout area="payment">
<h1 class="title">Terminals</h1>
<h2 class="subtitle">Stripe-Terminals den Standorten zuordnen</h2>

<div class="notification is-warning is-light">
  Die Terminals sind Beispieldaten. Das Auslesen aus Stripe und das Speichern folgen, sobald das Backend es unterstützt.
</div>

<div class="panel">
  <p class="panel-heading has-text-primary-dark">Terminals</p>
  <div class="panel-block">
    <table class="table is-fullwidth is-striped is-hoverable">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>Gerät</th>
          <th>Status</th>
          <th>Standort</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="terminal in terminals" :key="terminal.id">
          <td>{{ terminal.label }}</td>
          <td><code>{{ terminal.id }}</code></td>
          <td>{{ terminal.deviceType }}</td>
          <td>
            <span class="tag" :class="terminal.status == 'online' ? 'is-success' : 'is-warning'">{{ terminal.status }}</span>
          </td>
          <td>
            <div class="select is-small">
              <select v-model="draft[terminal.id]">
                <option :value="null">Nicht genutzt</option>
                <option
                  v-for="loc in location$.all"
                  :key="loc.id"
                  :value="loc.id"
                  :disabled="isTakenByOtherTerminal(loc.id!, terminal.id)">
                  {{ loc.name }}
                </option>
              </select>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>

<div class="is-flex is-align-items-center">
  <Buttons>
    <Button :fa-icon="['fas', 'times']" icon-position="left" @click="reset" :disabled="!isDirty || saveStatus === 'pending'">
      Zurücksetzen
    </Button>
    <Button class="is-primary" @click="save" :fa-icon="['fas', 'save']" icon-position="right" :disabled="!isDirty || saveStatus === 'pending'">
      Speichern
    </Button>
  </Buttons>
  <span v-if="saveStatus === 'pending'" class="ml-3 icon has-text-grey">
    <icon :icon="['fas', 'spinner']" :spin="true" />
  </span>
  <span v-else-if="saveStatus === 'success'" class="ml-3 icon-text has-text-success">
    <span class="icon"><icon :icon="['fas', 'check']" /></span>
    <span>Saved</span>
  </span>
</div>
</SettingsLayout>
</template>
