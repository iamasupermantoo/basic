<template>
  <div class="window-height">home</div>
  <!-- 客服图标 -->
  <q-avatar v-if="config.settings.online.showRegister"
    class="fixed-bottom-right q-mb-md q-mr-md cursor-pointer shadow-1 z-top">
    <img :src="imageSrc(config.onlineIcon)" alt="">
  </q-avatar>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';
export default {
  name: 'homeIndex',
  setup() {
    const $initStore = InitStore();

    const state = reactive({

      // 登录配置
      config: $initStore.config,

      // 客服图标
      onlineIcon: $initStore.config.onlineIcon,

      // 是否登录状态
      isLogin: $initStore.userToken != '',
      // cardList
      cardList: [
        { title: 'Membership <br> Organisations', url: 'team' },
        { title: 'National <br> Associations', url: 'National' },
        { title: 'Clubs And <br> Groups', url: 'Groups' },
      ],
      FAQ: [
        { title: 'How to change my account', content: 'Donec a eros justo. Fusce egestas tristique ultrices. Nam tempor, augue nec tincidunt molestie, massa nunc varius arcu, at scelerisque elit erat a magna. Donec quis erat at libero ultrices mollis. In hac habitasse platea dictumst.' },
        { title: 'How do I order a new card', content: 'Donec a eros justo. Fusce egestas tristique ultrices. Nam tempor, augue nec tincidunt molestie, massa nunc varius arcu, at scelerisque elit erat a magna. Donec quis erat at libero ultrices mollis. In hac habitasse platea dictumst.' },
        { title: 'How to change my account', content: 'Donec a eros justo. Fusce egestas tristique ultrices. Nam tempor, augue nec tincidunt molestie, massa nunc varius arcu, at scelerisque elit erat a magna. Donec quis erat at libero ultrices mollis. In hac habitasse platea dictumst.' },
        { title: 'How to change my account', content: 'Donec a eros justo. Fusce egestas tristique ultrices. Nam tempor, augue nec tincidunt molestie, massa nunc varius arcu, at scelerisque elit erat a magna. Donec quis erat at libero ultrices mollis. In hac habitasse platea dictumst.' },
        { title: 'How to change my account', content: 'Donec a eros justo. Fusce egestas tristique ultrices. Nam tempor, augue nec tincidunt molestie, massa nunc varius arcu, at scelerisque elit erat a magna. Donec quis erat at libero ultrices mollis. In hac habitasse platea dictumst.' },
      ],
    });

    return {
      imageSrc,
      ...toRefs(state),
    }
  }
}
</script>

<style scoped lang="scss">
:deep(.q-expansion-item .q-item) {
  padding: 0;
  margin-bottom: 16px;
}

:deep(.q-card .q-card__section--vert) {
  padding: 0;
  color: #666666 !important;
}

.size36 {
  font-size: 36px;
  line-height: 44px;
}

.size28 {
  font-size: 28px;
  line-height: 36px;
}

.homeBack {
  height: 600px;
  background: url('/public/images/pc/home/background.png');
  background-size: cover;
  background-repeat: no-repeat;
}
</style>
