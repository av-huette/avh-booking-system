<template>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-4">
            <div class="box">
              <h1 class="title">Client Setup</h1>
              <p class="subtitle">Set a name for this terminal and choose its location.</p>
              <div class="field">
                <label class="label">Client Name</label>
                <div class="control">
                  <input
                    class="input"
                    type="text"
                    v-model="locationName"
                    placeholder="e.g. Bar, Kitchen"
                    @keyup.enter="save"
                  />
                </div>
              </div>
              <div class="field">
                <label class="label">Location</label>
                <div class="control">
                  <div class="select is-fullwidth">
                    <select v-model="locationId">
                      <option :value="null" disabled>Select a location</option>
                      <option v-for="loc in location$.all" :key="loc.id" :value="loc.id">{{ loc.name }}</option>
                    </select>
                  </div>
                </div>
              </div>
              <div class="field">
                <div class="control">
                  <button class="button is-primary" :disabled="!canSave" @click="save">
                    Save
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { useSettingStore } from '../store/SettingStore';
import { useLocationStore } from '../store/LocationStore';
import { useSocketStore } from '../store/socketStore';

export default {
  data() {
    const setting$ = useSettingStore();
    return {
      locationName: setting$.clientId,
      locationId: setting$.locationId as number | null,
      setting$,
      location$: useLocationStore(),
    };
  },
  computed: {
    canSave(): boolean {
      return this.locationName.trim() != '' && this.locationId != null;
    },
  },
  mounted() {
    // The socket connects without a client id until setup is done; the backend accepts this.
    useSocketStore().queryLocations();
  },
  methods: {
    save() {
      if (!this.canSave) return;
      this.setting$.clientId = this.locationName.trim();
      this.setting$.locationId = this.locationId;
      this.setting$.$persist(); // write to localStorage before reloading
      // Reload so that the socket reconnects with the new client id and all data is loaded
      location.reload();
    },
  },
};
</script>
