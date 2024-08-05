<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useMeta, LocalStorage } from 'quasar';
import { InitStore } from 'src/stores/init';
import { imageSrc } from 'src/utils';
import { useI18n } from 'vue-i18n';
import { setLanguageFunc } from 'boot/i18n';
import { UserStore, UserInfoKey } from 'stores/user';
import {accessAPI} from 'src/apis';

export default defineComponent({
  name: 'App',
  setup() {
    onMounted(() => {
      accessAPI({name: 'init', type: 1})
    })
  },
  created: () => {
    const $initStore = InitStore();
    const $userStore = UserStore();
    const $i18n = useI18n();

    // 初始化国际化语言
    setLanguageFunc($i18n, $initStore.translate, $initStore.userLang)

    // 初始化用户信息
    if ($initStore.userToken.length > 0) {
      const userInfoStr = LocalStorage.getItem(UserInfoKey)
      $userStore.userInfo = userInfoStr ? JSON.parse(<string>userInfoStr) : {}
    }

    //  设置客户端Meta信息
    const metaData = {
      title: $initStore.config.name,
      meta: {
        description: {
          name: 'description',
          content: $initStore.config.name,
        },
      },
      link: {
        material: {
          rel: 'icon',
          href: imageSrc($initStore.config.logo),
          type: 'image/png',
        },
      },
    };
    useMeta(metaData);
  },
});
</script>
