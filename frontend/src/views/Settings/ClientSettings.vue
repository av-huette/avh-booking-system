<template>
<SettingsLayout area="system">
  <h1 class="title">Client</h1>
  <h2 class="subtitle">Einstellungen dieses Geräts</h2>
  <div class="field">
    <label class="label">Client Name</label>
    <div class="control">
      <input
        class="input"
        type="text"
        v-model="clientName"
        placeholder="e.g. Bar, Kitchen"
      />
    </div>
    <p class="help">Identifies this client to the server. Changes take effect after saving and reloading.</p>
  </div>
  <div class="field">
    <div class="control">
      <button class="button is-primary" :disabled="!clientName.trim() || clientName.trim() == setting$.clientId" @click="save">
        Save &amp; Reload
      </button>
    </div>
  </div>
  <div class="field">
    <label class="label">Location</label>
    <div class="control">
      <div class="select">
        <select :value="setting$.locationId" @change="setLocation">
          <option v-for="loc in location$.all" :key="loc.id" :value="loc.id">{{ loc.name }}</option>
        </select>
      </div>
    </div>
    <p class="help">
      Only products visible at this location are shown. Card payments use the terminal of this location.
      <span v-if="setting$.currentLocation && !setting$.currentTerminalId" class="has-text-warning">No terminal configured for this location.</span>
    </p>
  </div>
</SettingsLayout>
</template>

<script lang="ts">
import SettingsLayout from '../../components/SettingsLayout.vue';
import { useSettingStore } from '../../store/SettingStore';
import { useLocationStore } from '../../store/LocationStore';

export default {
  components: {
    SettingsLayout,
  },
  data() {
    const setting$ = useSettingStore();
    return {
      clientName: setting$.clientId,
      setting$,
      location$: useLocationStore(),
    };
  },
  methods: {
    save() {
      const name = this.clientName.trim();
      if (!name) return;
      // Persisted by the SettingStore; the reload reconnects the socket with the new client id
      this.setting$.clientId = name;
      this.setting$.$persist(); // write to localStorage before reloading
      location.reload();
    },
    setLocation(e: Event) {
      // Takes effect immediately, no reload needed
      this.setting$.locationId = parseInt((e.target as HTMLSelectElement).value);
    },
  },
};
</script>
