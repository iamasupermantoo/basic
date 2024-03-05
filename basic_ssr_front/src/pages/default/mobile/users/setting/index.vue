<template>
  <div>
    <q-list bordered class="q-mb-md" style="border:0 !important;overflow: hidden;border-radius: 8px;">
      <div v-for="(item, i) in settingList" :key="i" class="bg-white">
        <q-item @click="$router.push(item.toggle ?? item.route)" v-ripple class="q-pa-md" clickable>
          <q-item-section avatar style="padding-right: 11px;min-width: 0;">
            <q-img no-spinner width="20px" height="20px" :src="item.icon"></q-img>
          </q-item-section>

          <q-item-section>
            <q-item-label class="text-weight-bold">{{ $t(item.name) }}</q-item-label>
          </q-item-section>

          <q-item-section side>
            <div class="row no-wrap items-center">
              <div v-if="item.version">{{ item.version }}</div>
              <q-toggle v-if="item.toggle" size="sm" v-model="toggle" val="lg" />
              <q-icon v-else name="chevron_right" size="20px" />
            </div>
          </q-item-section>
        </q-item>
        <q-separator style="background: #F4F5FD;" inset />
      </div>
    </q-list>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import { useI18n } from 'vue-i18n'

export default defineComponent({
  name: 'SettingIndex',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('settings'),
    })

    const store = reactive({
      settingList: [
        {
          name: 'personalSetting',
          icon: '/images/icons/user.png',
          route: '/settings/update/info'
        }, {
          name: 'updatePassword',
          icon: '/images/icons/pwd.png',
          route: '/settings/update/password',
        }, {
          name: 'updateSecretKey',
          icon: '/images/icons/key.png',
          route: '/settings/update/secret-key'
        }, {
          name: 'bindTelephone',
          icon: '/images/icons/mobile.png',
          route: '/settings/update/telephone'
        }, {
          name: 'email',
          icon: '/images/icons/email.png',
          route: '/settings/update/email'
        }, {
          name: 'language',
          icon: '/images/icons/language.png',
          route: '/settings/language'
        }, {
          name: 'noticeEnable',
          icon: '/images/icons/message.png',
          toggle: true
        },
      ] as any,
      toggle: false
    })

    return {
      ...toRefs(store),
    }
  }
})
</script>

<style scoped></style>
