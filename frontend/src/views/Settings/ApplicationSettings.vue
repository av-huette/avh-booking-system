<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import SettingsLayout from '../../components/SettingsLayout.vue';
import { useSettingStore } from '../../store/SettingStore';
import Buttons from '../../composables/elements/Buttons.vue';
import Button from '../../composables/elements/Button.vue';

const route = useRoute();
const setting$ = useSettingStore();

const section = computed(() => route.params.section as string);

type DraftKey = 'compTitle' | 'compSlogan' | 'compLogo';

const draft = ref({ compTitle: '', compSlogan: '', compLogo: null as string | null });
const saveStatus = ref<'idle' | 'success'>('idle');

function storeValue(key: DraftKey): string | null {
  const s = setting$.get(key);
  return s == -1 ? null : s.value as string;
}

function valuesFromStore() {
  return {
    compTitle: storeValue('compTitle') ?? '',
    compSlogan: storeValue('compSlogan') ?? '',
    compLogo: storeValue('compLogo'),
  };
}

function loadFromStore() {
  draft.value = valuesFromStore();
}

const changedKeys = computed(() => {
  const stored = valuesFromStore();
  return (Object.keys(draft.value) as DraftKey[]).filter(key => draft.value[key] !== stored[key]);
});
const isDirty = computed(() => changedKeys.value.length > 0);

onMounted(() => {
  loadFromStore();
});

function save() {
  changedKeys.value.forEach(key => setting$.set(key, draft.value[key]));
  saveStatus.value = 'success';
  window.setTimeout(() => { saveStatus.value = 'idle' }, 1500);
}

function reset() {
  loadFromStore();
}

function changeLogo(e){
  const reader = new FileReader();
  reader.addEventListener("load", () => {
    const newValue = reader.result as string;
    if(newValue.length / 1024 > 200) {
      console.error("New Image uploaded with size (KB)", newValue.length / 1024, "This is too much. Reduce image Size so that is is below 200 KB.");
      return;
    }
    draft.value.compLogo = newValue;
  })

  reader.readAsDataURL(e.target.files[0]);
}

const companyIcon = ref("");
function changeIcon(e){
  const reader = new FileReader();
  reader.addEventListener("load", () => {
    companyIcon.value = reader.result;
    console.log(companyIcon.value.length);
  })

  reader.readAsDataURL(e.target.files[0]);
}
</script>

<template>

<SettingsLayout area="system">

<template v-if="section === 'company'">
  <h1 class="title">Company Settings</h1>

  <div class="columns">
    <div class="column is-3">
      Firmenname
    </div>
    <div class="column">
      <p class="control has-icons-left">
          <input type="text" v-model="draft.compTitle" class="input" placeholder="SOS Children's Villages">
          <span class="icon is-small is-left">
            <icon :icon="['fas', 'id-card']" />
          </span>
        </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Slogan
    </div>
    <div class="column">
      <p class="control has-icons-left">
          <input type="text" v-model="draft.compSlogan" class="input" placeholder="Every child a home!">
          <span class="icon is-small is-left">
            <icon :icon="['fas', 'microphone']" />
          </span>
        </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3 has-start-align">
      Logo (.png/.svg below 200kb)
    </div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="file" class="input" accept="image/png, .svg" @change="changeLogo">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'image']" />
        </span>
      </p>
      <!-- Logo Preview -->
      <img width="150px" v-if="draft.compLogo" :src="draft.compLogo">

    </div>
  </div>

  <div class="columns">
    <div class="column is-3 has-start-align">
      Icon
    </div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="file" class="input" accept="image/png, .svg" @change="changeIcon">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'image']" />
        </span>
      </p>
      <!-- Icon Preview -->
      <img width="50px ":src="companyIcon.valueOf()">
    </div>
  </div>

  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <div class="is-flex is-align-items-center">
        <Buttons>
          <Button :fa-icon="['fas', 'times']" icon-position="left" @click="reset" :disabled="!isDirty">
            Zurücksetzen
          </Button>

          <Button class="is-primary" @click="save" :fa-icon="['fas', 'save']" icon-position="right" :disabled="!isDirty">
            Speichern
          </Button>
        </Buttons>
        <span v-if="saveStatus === 'success'" class="ml-3 icon-text has-text-success">
          <span class="icon"><icon :icon="['fas', 'check']" /></span>
          <span>Saved</span>
        </span>
      </div>
    </div>
  </div>

</template>

<template v-if="section === 'look'">
  <h1 class="title">Application Look</h1>

  <div class="columns">
    <div class="column is-3">
      Application Color Scheme (Dark)
    </div>
    <div class="column">
      <div class="select">
        <select>
          <option value="">Not yet implemented</option>
        </select>
      </div>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Application Color Scheme (Light)
    </div>
    <div class="column">
      <div class="select">
        <select>
          <option value="">Not yet implemented</option>
        </select>
      </div>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Primary Color
    </div>
    <div class="column is-2">
      <input type="color" class="input">
    </div>
  </div>

</template>

<template v-if="section === 'localisation'">
  <h1 class="title">Localisation</h1>

  <div class="columns">
    <div class="column is-3">
      {{ $t('settings.language')}}
    </div>
    <div class="column is-2">
      <div class="select">
        <select v-model="$i18n.locale">
          <option v-for="l of $i18n.availableLocales" :key="`locale-${l}`" :value="l">{{ $t(`settings.${l}`) }}</option>
        </select>
      </div>
    </div>
  </div>

  <!-- ToDo Allow seperate changing of currency -->


</template>

</SettingsLayout>



</template>

<style scoped>
.columns{
  align-items: center;
}
.has-start-align {
  align-self: start;
  margin-top:.5em;
}
</style>