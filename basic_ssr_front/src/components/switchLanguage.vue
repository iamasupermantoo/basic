<template>
  <q-menu>
    <q-list class="q-ma-sm">
      <q-item @click="switchLang(lang)" :active="initStore.userLang == lang.alias" class="rounded-borders"
        v-for="(lang, langIndex) in initStore.languageList" :key="langIndex" clickable v-close-popup aria-hidden="true">
        <q-item-section avatar>
          <q-img no-spinner width="24px" height="24px" :src="imageSrc(lang.icon)" />
        </q-item-section>
        <q-item-section>{{ lang.name }}</q-item-section>
      </q-item>
    </q-list>
  </q-menu>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive } from 'vue';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';

export default defineComponent({
  name: 'SwitchLanguage',
  setup() {
    const $initStore = InitStore()
    const state = reactive({})

    const switchLang = async (lang: any) => {
      await $initStore.updateUserLang(lang.alias)
      location.reload()
    }

    return {
      imageSrc,
      ...toRefs(state),
      initStore: $initStore,
      switchLang
    }
  }
});
</script>

<style lang="scss" scoped>
:deep(.q-item .q-item__section) {
  min-width: auto;
  padding-right: 10px;
}
</style>
