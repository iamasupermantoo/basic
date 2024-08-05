<template>
  <q-layout view="hhh LpR fff" style="min-width: 1200px">
    <q-header bordered class="bg-white">
      <HeaderLayouts></HeaderLayouts>
    </q-header>

    <q-page-container>
      <div class="row no-wrap" style="margin-bottom: 200px;">
        <MenuLayouts></MenuLayouts>
        <div style="width: calc(100% - 284px);min-height: 90vh;">
          <router-view :key="$route.fullPath" />
        </div>
      </div>
    </q-page-container>

    <q-footer style="background: #021736">
      <FooterLayouts class="full-width"></FooterLayouts>
    </q-footer>

  </q-layout>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import HeaderLayouts from 'src/layouts/desktop/header.vue';
import FooterLayouts from 'src/layouts/desktop/footer.vue';
import MenuLayouts from 'src/layouts/desktop/menu.vue';

export default {
  components: { HeaderLayouts, FooterLayouts, MenuLayouts },
  name: 'LayoutsSetting',
  setup() {
    const state = reactive({
      drawer: true
    });
    return {
      ...toRefs(state),
    };
  },
};
</script>

<style lang="scss" scoped></style>
