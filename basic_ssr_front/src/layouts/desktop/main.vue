<template>
  <q-layout view="hhh LpR fFf" style="min-width: 1200px">
    <q-header bordered class="bg-white">
      <HeaderLayouts></HeaderLayouts>
    </q-header>

    <q-page-container>
      <div style="min-height: 90vh;">
        <router-view :key="$route.fullPath" />
      </div>
    </q-page-container>

    <FooterLayouts style="background: #021736;"></FooterLayouts>
  </q-layout>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import HeaderLayouts from 'src/layouts/desktop/header.vue';
import FooterLayouts from 'src/layouts/desktop/footer.vue';

export default {
  components: { HeaderLayouts, FooterLayouts },
  name: 'LayoutsMain',
  setup() {
    const state = reactive({});
    return {
      ...toRefs(state),
    };
  },
};
</script>

<style lang="scss" scoped></style>
