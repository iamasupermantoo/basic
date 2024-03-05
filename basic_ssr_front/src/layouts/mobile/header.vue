<template>
  <q-layout view="hhh LpR fFf" class="bg-grey-1">
    <q-header bordered class="bg-white text-black">
      <q-toolbar>
        <q-btn dense flat round :icon="data.leftBtn.icon" color="grey-8" @click="data.leftBtn.callback" />

        <q-toolbar-title class="text-center text-subtitle2">
          {{ data.title }}
        </q-toolbar-title>

        <q-icon v-if="data.rightBtn.icon" @click="data.rightBtn.callback" size="18px" :name="data.rightBtn.icon"></q-icon>
        <span v-else-if="data.rightBtn.text" @click="data.rightBtn.callback">{{ data.rightBtn.text }}</span>
        <q-btn v-else :ripple="false" flat color="grey-8" size="md"></q-btn>
      </q-toolbar>
    </q-header>
    <q-page-container>
      <router-view :key="$route.fullPath" @update="updateHeaderDataFunc" />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { useRouter } from 'vue-router';
import { assignRecursion } from 'src/utils/object';

export default {
  name: 'LayoutsHeader',
  setup() {
    const $router = useRouter();

    const state = reactive({
      data: {
        title: ' ',
        leftBtn: {
          icon: 'arrow_back',
          callback: () => {
            $router.back();
          },
        },
        rightBtn: {
          icon: '',
          callback: () => {
            return;
          },
        } as any,
      },
    });

    //  更新Header数据方法
    const updateHeaderDataFunc = (data: any) => {
      state.data = assignRecursion(state.data, data);
    };

    return {
      ...toRefs(state),
      updateHeaderDataFunc,
    };
  },
};
</script>
<style scoped></style>
