<script setup lang="ts">
import { ref, computed } from 'vue';
import Buttons from '../../composables/elements/Buttons.vue';
import Button from '../../composables/elements/Button.vue';

// Dummy: values only live in this component until the backend table exists
const saved = ref({ stripeEnabled: false, stripeApiKey: '' });
const draft = ref({ ...saved.value });
const saveStatus = ref<'idle' | 'pending' | 'success'>('idle');

const isDirty = computed(() =>
  draft.value.stripeEnabled !== saved.value.stripeEnabled ||
  draft.value.stripeApiKey !== saved.value.stripeApiKey
);

function save() {
  saveStatus.value = 'pending';
  // ToDo: send mutation to the WS once the backend table exists
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
<h1 class="title">Payment Settings</h1>
<h2 class="subtitle">Zahlungsdienstleister - Stripe</h2>
<div class="columns">
  <div class="column is-3">Aktivieren</div>
  <div class="column">
    <label class="checkbox">
      <input type="checkbox" v-model="draft.stripeEnabled">
      Stripe aktivieren
    </label>
  </div>
</div>

<div class="columns">
  <div class="column is-3">Stripe API Key</div>
  <div class="column">
    <p class="control has-icons-left">
      <input type="password" class="input" v-model="draft.stripeApiKey">
      <span class="icon is-small is-left">
        <icon :icon="['fas', 'key']" />
      </span>
    </p>
  </div>
</div>

<div class="columns">
  <div class="column is-3"></div>
  <div class="column">
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
  </div>
</div>
</template>

<style scoped>
.columns{
  align-items: center;
}
</style>
