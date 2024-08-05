<template>
  <q-list class="q-ml-xl q-mt-xl q-mr-md" style="min-width: 220px;width: 220px" padding>
    <div v-for="(menu, menuIndex) in menuList" :key="menuIndex">
      <!-- 含有子级的列表 -->
      <q-expansion-item accordion :header-style="{ borderRadius: '4px' }" v-if="menu.children.length > 0" default-opened>
        <!-- 父级 -->
        <template v-slot:header>
          <q-item-section avatar style="min-width: auto">
            <q-img no-spinner :src="imageSrc(menu.icon)" width="24px" height="24px"></q-img>
          </q-item-section>
          <q-item-section avatar>
            <div style="user-select: none;min-width: 120px;">{{ $t(menu.name) }}</div>
          </q-item-section>
        </template>

        <!-- 子级 -->
        <q-item class="rounded-borders" :to="children.route" v-for="(children, childrenIndex) in menu.children"
          :key="childrenIndex" :header-inset-level="1" :active="$route.path == children.route" clickable>
          <q-item-section avatar>
            <q-img no-spinner width="24px" height="24px" src=""></q-img>
          </q-item-section>
          <q-item-section>
            <div style="user-select: none">{{ $t(children.name) }}</div>
          </q-item-section>
        </q-item>
      </q-expansion-item>


      <!-- 不含子级的列表 -->
      <q-item v-else class="rounded-borders" :active="$route.path == menu.route" :to="menu.route">
        <q-item-section avatar style="min-width: auto">
          <q-img no-spinner :src="imageSrc(menu.icon)" width="24px" height="24px"></q-img>
        </q-item-section>
        <q-item-section>
          <div style="user-select: none">{{ $t(menu.name) }}</div>
        </q-item-section>
      </q-item>
    </div>
  </q-list>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import { InitStoreState, InitStore } from 'src/stores/init';

export default defineComponent({
  name: 'LayoutsMenu',
  setup() {
    const $initStore = InitStore()

    const state = reactive({
      // 项目logo
      config: InitStoreState.config,

      // 菜单列表
      menuList: [] as any,
    });

    state.menuList = $initStore.userMenu;

    return {
      imageSrc,
      ...toRefs(state),
    };
  },
});
</script>
<style lang="scss" scoped></style>
