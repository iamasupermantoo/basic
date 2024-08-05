<template>
  <div>
    <div class="row justify-between no-wrap q-pt-xl q-px-md q-pb-lg bg-white">
      <!-- 头像 -->
      <div class="row">
        <q-avatar @click="$router.push({ name: 'SettingUpdateUserInfo' })" class="q-mr-md avatar">
          <q-img no-spinner :src="imageSrc(userInfo.avatar ?? config.logo)" width="50px" height="50px" />
        </q-avatar>
        <div class="col-8">
          <div class="text-weight-bolder">
            {{ userInfo.username }}
          </div>
          <div>
            {{ userInfo.email }}
          </div>
          <div class="row no-wrap">
            <div @click="$router.push({ name: 'UserLevel' })">
              <q-chip class="text-white" size="sm" style="background: #322B19;border: 1px solid #F7DEB6;">
                <q-img no-spinner src="/images/icons/vip-icon.png" class="q-mr-xs" width="11px" height="11px" />
                <span style="color: #F7DEB6;">Lv{{ userInfo.level }}</span>
              </q-chip>
            </div>
            <q-chip class="text-grey-9 bg-white" size="sm" style="border: 1px solid #F1F1F1;">
              <q-img no-spinner src="/images/icons/credit.png" class="q-mr-xs" width="11px" height="11px" />
              {{ $t('creditScore') + userInfo.score }}
            </q-chip>
            <div @click="$router.push({ name: 'UserRealAuth' })">
              <q-chip v-if="userInfo.authStatus == 0" style="border: 1px solid red;" class="text-red bg-white" size="sm">
                {{ $t('alreadyRealName') }}
                <q-icon class="bg-white" name="keyboard_arrow_right" size="11px"></q-icon>
              </q-chip>
              <q-chip v-else-if="userInfo.authStatus == 10" class="bg-info text-white" size="sm">
                {{ $t('pendingRealName') }}
                <q-icon class="text-white" name="keyboard_arrow_right" size="11px"></q-icon>
              </q-chip>
              <q-chip v-else-if="userInfo.authStatus == 20" style="border: 1px solid #01AC66;"
                class="text-primary bg-white" size="sm">
                {{ $t('realNameFailed') }}
                <q-icon class="text-primary" name="keyboard_arrow_right" size="11px"></q-icon>
              </q-chip>
              <q-chip v-else class="bg-negative text-white" size="sm">
                {{ $t('notRealName') }}
                <q-icon class="text-white" name="keyboard_arrow_right" size="11px"></q-icon>
              </q-chip>
            </div>
          </div>
        </div>
      </div>
      <div @click="$router.push({ name: 'SettingUpdateUserInfo' })" class="row items-center">
        <q-icon name="keyboard_arrow_right" size="24px"></q-icon>
      </div>
    </div>

    <div class="bg-grey-1 q-px-md q-py-md full-width">
      <!-- 充值、提现 -->
      <div class="row justify-between q-mb-md btn">
        <q-btn @click="$router.push(quickMenu.route)" v-for="(quickMenu, quickMenuIndex) in quickMenuList"
          :key="quickMenuIndex" v-show="quickMenu.data.isMobile" style="width: 47%;"
          class="bg-white q-py-sm rounded-borders" no-caps unelevated>
          <div class="row justify-start items-center">
            <q-img no-spinner class="q-mr-sm" :src="imageSrc(quickMenu.icon)" width="42px" height="42px" />
            <div>{{ $t(quickMenu.name) }}</div>
          </div>
        </q-btn>
      </div>

      <!-- 用户列表 -->
      <q-list v-for="(item, i) in userList" :key="i" bordered class="q-mb-md rounded-borders no-border">
        <div v-for="(child, childKey) in item.children" :key="childKey" class="bg-white">
          <q-item @click="$router.push(child.route)" class="q-pa-md" clickable>
            <q-item-section avatar class="q-mr-sm" style="min-width: 0;">
              <q-img no-spinner :src="`${imageSrc(child.icon)}`" width="24px" height="24px" />
            </q-item-section>

            <q-item-section>
              <q-item-label class="text-weight-bold">{{ $t(child.name) }}</q-item-label>
            </q-item-section>

            <q-item-section side>
              <q-icon name="keyboard_arrow_right" size="24px"></q-icon>
            </q-item-section>
          </q-item>
          <q-separator style="background: #F4F5FD;" inset />
        </div>
      </q-list>
      <q-btn @click="dialog = true" class="full-width q-mb-lg q-mt-md" size="lg" unelevated rounded no-caps
        color="primary" :label="$t('logout')" />
    </div>

    <!-- 退出登录 -->
    <q-dialog full-width v-model="dialog" position="bottom">
      <q-card>
        <q-card-section>
          <div class="row justify-center q-mb-lg">
            <q-separator class="rounded-borders col-3" style="height: 4px" color="grey-4" />
          </div>
          <div class="row justify-center q-my-md">
            <div class="text-weight-bold text-h6">{{ $t('logout') }}</div>
          </div>
          <q-separator class="q-mt-md  q-mb-lg" color="grey-4" />
          <div class="row justify-center q-mb-xl">
            <div class="text-grey-10">{{ $t('logoutSmall') }}</div>
          </div>
          <div class="row justify-between no-wrap">
            <q-btn @click="dialog = false" size="lg" class="q-mr-md text-primary bg-white col-5" unelevated rounded
              no-caps style="border:1px solid #01AC66" :label="$t('cancel')" />
            <q-btn @click="logoutFunc" size="lg" class="col-5" unelevated rounded no-caps color="primary"
              :label="$t('logout')" />
          </div>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { useRouter } from 'vue-router';
import { InitStore } from 'src/stores/init';
import { UserStore } from 'src/stores/user';
import { imageSrc } from 'src/utils';
import { userInfoAPI } from 'src/apis/user';

export default defineComponent({
  name: 'UserIndex',
  setup() {
    const $router = useRouter();
    const $initStore = InitStore()
    const $userStore = UserStore()

    const state = reactive({
      userInfo: {} as any,

      config: $initStore.config as any,

      // 用户菜单
      userList: $initStore.userMenu as any,

      // 快捷菜单
      quickMenuList: $initStore.quickMenu as any,

      // 退出彈窗
      dialog: false,
    })

    onMounted(() => {
      userInfoAPI().then((res: any) => {
        state.userInfo = res
        $userStore.updateUserInfo(res)
      })
    })

    // 退出登录
    const logoutFunc = async () => {
      await $initStore.removeUserToken()
      void $router.push({ name: 'HomeIndex' })
    }

    return {
      imageSrc,
      ...toRefs(state),
      logoutFunc,
    }
  }
})
</script>

<style scoped></style>
