<template>
  <q-layout view="hHh LpR fFf" class="bg-grey-1">
    <q-drawer v-model="leftDrawerOpen" side="left" behavior="mobile" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-drawer v-model="rightDrawerOpen" side="right" behavior="mobile" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container>
      <router-view :key="$route.fullPath" />
    </q-page-container>

    <q-footer reveal bordered class="bg-white text-black text-caption">
      <q-tabs v-model="currentTab" dense indicator-color="transparent" active-color="primary">
        <q-route-tab v-for="(tab, tabIndex) in tabBarList" :key="tabIndex" :name="tab.route" :icon="'img:' +
          imageSrc(tab.route == currentTab ? tab.activeIcon : tab.icon)
          " :ripple="false" :to="tab.route">
          <template v-slot:default>
            <div class="text-caption">{{ $t(tab.name) }}</div>
          </template>
        </q-route-tab>
      </q-tabs>
    </q-footer>
  </q-layout>
</template>

<script lang="ts">
import { toRefs, reactive, ref } from 'vue';
import { imageSrc } from 'src/utils';
import { useRouter } from 'vue-router';
import { onMounted } from 'vue';
import { InitStore } from 'src/stores/init';

export default {
  name: 'LayoutsTabBar',
  setup() {
    const $router = useRouter();
    const $initStore = InitStore();
    const leftDrawerOpen = ref(false);
    const rightDrawerOpen = ref(false);

    const state = reactive({
      currentTab: '/',
      tabBarList: $initStore.tabBars as any,
    });

    onMounted(() => {
      state.currentTab = $router.currentRoute.value.path;
    });

    return {
      leftDrawerOpen,
      rightDrawerOpen,
      imageSrc,
      ...toRefs(state),
    };
  },
};
</script>
