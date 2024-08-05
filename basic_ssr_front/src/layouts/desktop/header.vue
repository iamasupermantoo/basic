<template>
  <div>
    <q-toolbar class="q-pr-xl" style="height: 60px;padding-left: 64px">
      <!-- 左侧 logo -->
      <q-toolbar-title @click="$router.push({ name: 'HomeIndex' })" shrink class="cursor-pointer text-black"
        style="margin-right: 88px;">
        <q-avatar>
          <img width="42px" height="42px" :src="imageSrc(config.logo)" alt="">
        </q-avatar>
        <span class="q-ml-sm">{{ config.name }}</span>
      </q-toolbar-title>

      <!-- 左侧tabBar菜单 -->
      <div class="row no-wrap items-center">
        <div v-for="(tabBar, tabBarIndex) in tabBarList" :key="tabBarIndex">
          <!-- 有子级 -->
          <q-btn-dropdown v-if="tabBar.children.length > 0" v-show="tabBar.data.isDesktop" :menu-offset="[80, 18]"
            class="text-grey-8 q-mr-sm q-py-xs q-px-sm" :label="$t(tabBar.name)" dense flat no-wrap no-caps>
            <q-list v-if="tabBar.children.length > 0" class="q-ma-sm">
              <q-item v-for="(children, childrenIndex) in tabBar.children" :key="childrenIndex" clickable
                aria-hidden="true" :to="children.route" class="rounded-borders">
                <q-item-section avatar style="min-width:auto">
                  <q-img no-spinner width="30px" height="30px" :src="imageSrc(children.icon ?? '')"></q-img>
                </q-item-section>
                <q-item-section class="text-grey-8">{{ $t(children.name) }}</q-item-section>
              </q-item>
            </q-list>
          </q-btn-dropdown>

          <!-- 没有子级 -->
          <q-btn v-else v-show="tabBar.data.isDesktop" class="text-grey-8 q-mr-sm q-py-xs q-px-md"
            :label="$t(tabBar.name)" dense flat no-wrap no-caps :to="tabBar.route"></q-btn>
        </div>
      </div>

      <!-- 左侧快捷菜单 -->
      <q-btn-dropdown :menu-offset="[50, 18]" class="text-grey-8 q-mr-sm q-py-xs q-px-sm" :label="$t('more')" dense flat
        no-caps>
        <q-list class="q-ma-sm">
          <q-item @click="$router.push(quickMenu.route)" v-for="(quickMenu, quickMenuIndex) in quickMenuList"
            class="rounded-borders" v-show="quickMenu.data.isDesktop" :key="quickMenuIndex" clickable v-close-popup
            aria-hidden="true">
            <q-item-section avatar style="min-width:auto">
              <q-img no-spinner width="30px" height="30px" :src="imageSrc(quickMenu.icon)"></q-img>
            </q-item-section>
            <q-item-section class="text-grey-8">{{ $t(quickMenu.name) }}</q-item-section>
          </q-item>
        </q-list>
      </q-btn-dropdown>

      <q-space />

      <!-- 右侧搜索 -->
      <q-input dense rounded outlined v-model="searchVal" :placeholder="$t('search')">
        <template v-slot:prepend>
          <q-icon style="color: #999999 !important;" name="o_search" />
        </template>
      </q-input>


      <!-- 右侧固定按钮 -->
      <div class="row items-center no-wrap">
        <!-- 登录状态 -->
        <div v-if="isLogin" class="row items-center no-wrap">
          <!-- 右侧Deposit -->
          <q-btn @click="$router.push({ name: 'WalletsDeposit', query: { mode: 1 } })" rounded flat dense no-wrap
            class="bg-primary text-white q-px-md q-mx-sm" no-caps :label="$t('deposit')"></q-btn>

          <!-- 头像 -->
          <q-btn class="q-mx-xs" round flat>
            <q-avatar size="28px">
              <q-img no-spinner :src="imageSrc(userInfo.avatar)"></q-img>
            </q-avatar>
            <q-menu :offset="[300, 15]" class="q-pa-sm">
              <q-list style="min-width: 218px;">
                <!-- 固定头部 -->
                <q-item aria-hidden="true">
                  <div class="row items-center">
                    <q-avatar size="40px" class="q-mt-sm cursor-pointer" @click="$router.push({ name: 'SettingIndex' })">
                      <q-img no-spinner :src="imageSrc(userInfo.avatar)"></q-img>
                    </q-avatar>
                    <div class="q-ml-sm">
                      <div class="row items-center">
                        <span class="q-mr-xs text-body1">{{ userInfo.username }}</span>
                        <span class="text-grey-7 text-caption">(ID:{{ userInfo.id }})</span>
                      </div>
                      <div class="row no-wrap q-mt-xs">
                        <!-- 会员等级 -->
                        <q-btn size="xs" rounded flat dense no-wrap class="q-px-sm q-mr-xs" no-caps
                          @click="$router.push({ name: 'UserLevel' })"
                          style="border: 1px solid #F7DEB6;color: #F7DEB6;background: #322B19;">
                          <q-img no-spinner width="13px" height="12px" src="/images/icons/vip-icon.png"></q-img>
                          <div class="q-ml-xs" style="font-size: 10px;">Lv{{ userInfo.level }}</div>
                        </q-btn>
                        <!-- 信用分 -->
                        <q-btn size="xs" rounded flat dense no-wrap class="q-px-sm q-mr-xs bg-grey-4 text-primary"
                          v-if="userInfo.score > 60">
                          <q-img no-spinner width="13px" height="13px" src="/images/icons/credit.png"></q-img>
                          <div class="q-ml-xs" style="font-size: 10px;">{{ $t('creditScore') + userInfo.score }}</div>
                        </q-btn>
                        <q-btn size="xs" rounded flat dense no-wrap class="q-px-sm q-mr-xs bg-grey-4 text-red" v-else>
                          <q-img no-spinner width="13px" height="13px" src="/images/icons/credit.png"></q-img>
                          <div class="q-ml-xs" style="font-size: 10px;">{{ $t('creditScore') + userInfo.score }}</div>
                        </q-btn>
                        <!-- 实名 -->
                        <q-btn size="xs" rounded flat dense no-wrap class="bg-grey-4 text-red q-px-sm" no-caps
                          v-if="userInfo.authStatus == 0" @click="$router.push({ name: 'UserRealAuth' })">
                          <div style="font-size: 10px">
                            {{ $t('alreadyRealName') }}
                          </div>
                        </q-btn>
                        <q-btn size="xs" rounded flat dense no-wrap class="bg-info text-white q-px-sm" no-caps
                          v-else-if="userInfo.authStatus == 10" @click="$router.push({ name: 'UserRealAuth' })">
                          <div style="font-size: 10px">
                            {{ $t('pendingRealName') }}
                          </div>
                        </q-btn>
                        <q-btn size="xs" rounded flat dense no-wrap class="bg-primary text-white q-px-sm" no-caps
                          v-else-if="userInfo.authStatus == 20">
                          <div style="font-size: 10px">
                            {{ $t('realNameFailed') }}
                          </div>
                        </q-btn>
                        <q-btn size="xs" rounded flat dense no-wrap class="bg-negative text-white q-px-sm" no-caps v-else
                          @click="$router.push({ name: 'UserRealAuth' })">
                          <div style="font-size: 10px">
                            {{ $t('notRealName') }}
                          </div>
                        </q-btn>
                      </div>
                    </div>
                  </div>
                </q-item>

                <q-separator inset class="q-mt-md q-mb-sm" />

                <!-- 用户列表 -->
                <q-item @click="$router.push(userMenu.route)" dense v-for="(userMenu, userMenuIndex) in homeMenuList"
                  :key="userMenuIndex" clickable v-close-popup aria-hidden="true" class="q-py-md rounded-borders">
                  <q-item-section avatar style="min-width:auto">
                    <q-img no-spinner width="24px" height="24px" :src="imageSrc(userMenu.icon)"></q-img>
                  </q-item-section>
                  <q-item-section class="text-grey-8 text-body1">
                    {{ $t(userMenu.name) }}
                  </q-item-section>
                </q-item>

                <q-separator inset class="q-my-sm" />

                <!-- 退出 -->
                <q-item @click="Logout()" clickable v-close-popup aria-hidden="true" class="rounded-borders">
                  <q-item-section avatar style="min-width:auto">
                    <q-icon size="24px" name="logout" />
                  </q-item-section>
                  <q-item-section class="text-grey-8 text-body1">{{ $t('logout') }}</q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>

          <!-- 钱包 -->
          <q-btn @click="$router.push({ name: 'WalletsIndex' })" round dense flat color="grey-8" size="16px"
            icon="o_account_balance_wallet" class="q-mx-xs"></q-btn>
        </div>

        <!-- 未登录状态 -->
        <div v-else>
          <q-btn @click="dialogOpenLogin(true)" rounded flat dense no-wrap class="bg-grey-4 text-grey-8 q-px-md q-ml-sm"
            no-caps :label="$t('login')"></q-btn>
          <q-btn @click="dialogOpenRegister(true)" rounded flat dense no-wrap
            class="bg-primary text-white q-px-md q-ml-sm" no-caps :label="$t('register')"></q-btn>
        </div>
        <q-btn v-if="config.settings.lang.showHome" round dense flat color="grey-8" size="16px" icon="o_language"
          class="q-mx-xs">
          <switchLanguage :offset="[0, 20]"></switchLanguage>
        </q-btn>
      </div>

    </q-toolbar>

    <LoginPages ref="LoginRef" @switchDialogFunc="switchDialogFunc"></LoginPages>
    <RegisterPages ref="RegisterRef" @switchDialogFunc="switchDialogFunc">
    </RegisterPages>
  </div>
</template>

<script lang="ts">
import LoginPages from 'src/pages/default/desktop/login.vue';
import RegisterPages from 'src/pages/default/desktop/register.vue';
import switchLanguage from 'src/components/switchLanguage.vue';
import { useRouter } from 'vue-router';
import { reactive, toRefs, ref, watch, onMounted } from 'vue';
import { imageSrc } from 'src/utils';
import { userInfoAPI } from 'src/apis/user';
import { InitStore } from 'src/stores/init';
import { UserStore } from 'src/stores/user';

export default {
  name: 'LayoutsHeader',
  components: { LoginPages, RegisterPages, switchLanguage },
  setup() {
    const $router = useRouter();
    const $initStore = InitStore()
    const $userStore = UserStore()

    const LoginRef = ref(null) as any;
    const RegisterRef = ref(null) as any;

    const state = reactive({
      userInfo: { avatar: '', username: '', email: '', level: 1, authStatus: 0 } as any,

      // 配置
      config: $initStore.config,

      //  是否登录
      isLogin: $initStore.userToken.length > 0,

      // 搜索框
      searchVal: '',

      //  左侧tabBar菜单
      tabBarList: $initStore.tabBars,

      // 左侧快捷菜单
      quickMenuList: $initStore.quickMenu,

      // 右侧头像菜单
      homeMenuList: $initStore.homeMenu,
    });
    state.userInfo = $userStore.userInfo

    onMounted(() => {
      switch ($router.currentRoute.value.name) {
        case 'UserLogin':
          dialogOpenLogin(true)
          break
        case 'UserRegister':
          dialogOpenRegister(true)
          break
      }

      //  如果有登录状态, 更新用户信息
      if (state.isLogin) {
        userInfoAPI().then((res: any) => {
          $userStore.updateUserInfo(res)
        })
      }
    })

    // dialogOpenFunc 打开登录
    const dialogOpenLogin = (showStatus: boolean) => {
      LoginRef.value?.openDialog(showStatus)
    }

    // dialogOpenFunc 打开注册
    const dialogOpenRegister = (showStatus: boolean) => {
      RegisterRef.value?.openDialog(showStatus)
    }

    // switchDialogFunc 切换注册登录dialog
    const switchDialogFunc = (LoginStatus: boolean, RegisterStatus: boolean) => {
      LoginRef.value?.openDialog(LoginStatus)
      RegisterRef.value?.openDialog(RegisterStatus)
    }

    // 退出登录
    const Logout = async () => {
      await $initStore.removeUserToken()
      void $router.push({ name: 'HomeIndex' })
    }

    // 监听用户Token
    watch(() => [$initStore.userToken, $userStore.userInfo], ([userToken, userInfo]) => {
      state.isLogin = userToken.length > 0
      state.userInfo = userInfo
    })

    return {
      imageSrc,
      ...toRefs(state),
      LoginRef,
      RegisterRef,
      dialogOpenLogin,
      dialogOpenRegister,
      switchDialogFunc,
      Logout,
    };
  },
};
</script>

<style lang="scss" scoped></style>
