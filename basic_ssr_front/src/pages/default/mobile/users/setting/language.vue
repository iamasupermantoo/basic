<template>
  <div>
    <q-list bordered class="q-mb-md rounded-borders no-border">
      <div @click="switchLang(item)" v-for="(item, i) in languageList" :key="i" class="bg-white">
        <q-item v-ripple class="q-pa-md rounded-borders" clickable>
          <q-item-section avatar class="q-mr-sm" style="min-width: 0;">
            <q-img no-spinner :src="imageSrc(item.icon)" width="32px" height="21px" />
          </q-item-section>

          <q-item-section>
            <q-item-label class="text-weight-bold">{{ item.name }}</q-item-label>
          </q-item-section>

          <q-item-section side>
            <q-icon v-if="locale == item.alias" name="check_circle" color="primary" size="20px"></q-icon>
          </q-item-section>
        </q-item>
        <q-separator style="background: #F4F5FD;" inset />
      </div>
    </q-list>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import { InitStore } from 'src/stores/init';
import { imageSrc } from 'src/utils';
import { useI18n } from 'vue-i18n'

// 列表
export default defineComponent({
  name: 'SettingsLanguageIndex',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('language'),
    })

    const { locale } = useI18n({ useScope: 'global' })
    const $initStore = InitStore();
    const state = reactive({
      languageList: $initStore.languageList as any,
      locale,
    })

    const switchLang = async (language: any) => {
      await $initStore.updateUserLang(language.alias)
      setTimeout(() => {
        location.reload()
      }, 200)
    }

    return {
      ...toRefs(state),
      switchLang,
      imageSrc
    }
  }
})
</script>

<style scoped></style>
