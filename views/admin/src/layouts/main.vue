<template>
  <q-layout view="hHh lpR fFf">
    <q-header reveal bordered class="bg-primary text-white">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="toggleLeftDrawer" />
        <div
          @click="$router.push('/')"
          class="row justify-start items-center cursor-pointer"
        >
          <q-avatar>
            <q-img
              :src="imageSrc(adminConfig.logo)"
              no-spinner
              width="30px"
              height="30px"
            />
          </q-avatar>
          <div class="text-h6">{{ adminConfig.name }}</div>
        </div>
        <q-toolbar-title></q-toolbar-title>
        <q-btn
          dense
          flat
          :label="'剩余:' + remainingDays() + '天'"
          v-if="$q.screen.width > 428"
        ></q-btn>
        <q-btn-dropdown flat>
          <template v-slot:label>
            <q-avatar rounded size="30px">
              <q-img
                :src="imageSrc(adminInfo?.avatar ?? '')"
                width="30px"
                height="30px"
                no-spinner
              />
            </q-avatar>
            <div class="q-ml-xs">{{ adminInfo?.username }}</div>
          </template>
          <q-list class="bg-primary text-white text-body2">
            <q-item
              v-close-popup
              clickable
              @click="dialogOpenFunc(dialogUserInfoKey)"
            >
              <q-item-section>
                <div class="row q-gutter-sm items-center">
                  <q-icon name="sym_o_contact_mail" size="sm"></q-icon>
                  <div>用户信息</div>
                </div>
              </q-item-section>
            </q-item>
            <q-item
              v-close-popup
              clickable
              @click="dialogOpenFunc(dialogUpdatePassword)"
            >
              <q-item-section>
                <div class="row q-gutter-sm items-center">
                  <q-icon name="sym_o_lock" size="sm"></q-icon>
                  <div>登陆密码</div>
                </div>
              </q-item-section>
            </q-item>
            <q-item
              v-close-popup
              clickable
              @click="dialogOpenFunc(dialogUpdateSecurityPassword)"
            >
              <q-item-section>
                <div class="row q-gutter-sm items-center">
                  <q-icon name="sym_o_security" size="sm"></q-icon>
                  <div>安全密码</div>
                </div>
              </q-item-section>
            </q-item>
            <q-separator></q-separator>
            <q-item v-close-popup clickable @click="adminLogout">
              <q-item-section>
                <div class="row q-gutter-sm items-center">
                  <q-icon name="sym_o_logout" size="sm"></q-icon>
                  <div>退出登陆</div>
                </div>
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>
        <q-btn
          dense
          flat
          round
          icon="sym_o_fullscreen"
          v-if="$q.screen.width > 428"
          @click="$q.fullscreen.toggle()"
        ></q-btn>
      </q-toolbar>
    </q-header>

    <q-drawer
      show-if-above
      v-model="leftDrawerOpen"
      side="left"
      bordered
      :behavior="$q.platform.is.mobile ? 'mobile' : 'desktop'"
      :mini="miniState"
      :width="220"
    >
      <q-item
        clickable
        v-ripple
        @click="$router.push('/')"
        :active="$route.fullPath == '/'"
        class="q-mt-xs"
        :active-class="'bg-light-blue-1'"
        onselectstart="return false"
      >
        <q-item-section avatar style="min-width: 0px">
          <div class="text-light-blue">
            <q-icon name="dashboard" size="1.5rem"></q-icon>
          </div>
        </q-item-section>
        <q-item-section>
          <div class="text-body2 text-black">控制台</div>
        </q-item-section>
      </q-item>
      <MenuComponent :data="menuList"></MenuComponent>
      <div style="height: 100px"></div>
    </q-drawer>

    <q-drawer v-model="rightDrawerOpen" side="right" bordered></q-drawer>

    <q-page-container>
      <router-view :key="$router.currentRoute.value.fullPath" />
    </q-page-container>
  </q-layout>
  <DialogComponent ref="domDialogRef" @done="dialogDoneFunc"></DialogComponent>
</template>

<script lang="ts">
import { imageSrc } from 'src/utils';
import { useInitStore } from 'src/stores/init';
import { useRouter } from 'vue-router';
import { ref, reactive, toRefs, onMounted, watch } from 'vue';
import DialogComponent from 'src/components/dialog.vue';
import MenuComponent from 'src/components/menu.vue';
import { InputTypeList } from 'src/utils/define';
import { adminInfoAPI, adminAudioAPI } from 'src/apis';
import { NotifyNegative, WarningNotify } from 'src/utils/notify';
import { onUnmounted } from 'vue';

export default {
  name: 'MainLayout',
  components: { DialogComponent, MenuComponent },
  setup() {
    const $router = useRouter();
    const $initStore = useInitStore();
    const leftDrawerOpen = ref(false);
    const rightDrawerOpen = ref(false);
    const domDialogRef = ref(null) as any;
    const dialogUserInfoKey = 'userInfo';
    const dialogUpdatePassword = 'updatePassword';
    const dialogUpdateSecurityPassword = 'updateSecurity';

    const state = reactive({
      closeAudio: false,
      adminInfo: $initStore.userInfo,
      adminConfig: $initStore.config,
      menuList: $initStore.menuList as any,
      miniState: false,

      //  菜单弹窗事件
      dialogList: new Map([
        [
          dialogUserInfoKey,
          {
            id: dialogUserInfoKey,
            url: '/update',
            title: '更新管理信息',
            sizing: 'small',
            params: {
              avatar: $initStore.userInfo.avatar,
              nickname: $initStore.userInfo.nickname,
              email: $initStore.userInfo.email,
            } as any,
            inputList: [
              {
                label: '管理头像',
                field: 'avatar',
                type: InputTypeList.Image,
                default: '',
                readonly: false,
                data: [],
              },
              {
                label: '管理昵称',
                field: 'nickname',
                type: InputTypeList.Text,
                default: '',
                readonly: false,
                data: [],
              },
              {
                label: '管理邮箱',
                field: 'email',
                type: InputTypeList.Text,
                default: '',
                readonly: false,
                data: [],
              },
            ],
            buttons: {
              cannel: { label: '取消', color: 'grey', size: 'md' },
              confirm: { label: '提交', color: 'primary', size: 'md' },
            },
          },
        ],
        [
          dialogUpdatePassword,
          {
            id: dialogUpdatePassword,
            url: '/update/password',
            title: '修改登录密码',
            sizing: 'small',
            params: {
              type: 1,
              oldPassword: '',
              newPassword: '',
              cmfPassword: '',
            } as any,
            inputList: [
              {
                label: '旧密码',
                field: 'oldPassword',
                type: InputTypeList.Password,
                default: '',
                readonly: false,
                data: [],
              },
              {
                label: '新密码',
                field: 'newPassword',
                type: InputTypeList.Password,
                default: '',
                readonly: false,
                data: [],
              },
              {
                label: '确认密码',
                field: 'cmfPassword',
                type: InputTypeList.Password,
                default: '',
                readonly: false,
                data: [],
              },
            ],
            buttons: {
              cannel: { label: '取消', color: 'grey', size: 'md' },
              confirm: { label: '提交', color: 'primary', size: 'md' },
            },
          },
        ],
        [
          dialogUpdateSecurityPassword,
          {
            id: dialogUpdateSecurityPassword,
            url: '/update/password',
            title: '修改安全密码',
            sizing: 'small',
            params: {
              type: 2,
              oldPassword: '',
              newPassword: '',
              cmfPassword: '',
            } as any,
            inputList: [
              {
                label: '旧密码',
                field: 'oldPassword',
                type: InputTypeList.Password,
                default: '',
                readonly: false,
                data: [],
              },
              {
                label: '新密码',
                field: 'newPassword',
                type: InputTypeList.Password,
                default: '',
                readonly: false,
                data: [],
              },
              {
                label: '确认密码',
                field: 'cmfPassword',
                type: InputTypeList.Password,
                default: '',
                readonly: false,
                data: [],
              },
            ],
            buttons: {
              cannel: { label: '取消', color: 'grey', size: 'md' },
              confirm: { label: '提交', color: 'primary', size: 'md' },
            },
          },
        ],
      ]),
    });

    onMounted(() => {
      //  提示音递归请求
      const audioFunc = () => {
        adminAudioAPI().then((res: any) => {
          if (state.closeAudio) {
            return;
          }
          if (res && res.title !== '') {
            WarningNotify(res.title);
            const audioObj = new Audio(imageSrc(res.source));
            audioObj.play().catch(() => {
              NotifyNegative('页面未激活播放声音, 点击页面开启声音提示');
            });
          }
          setTimeout(audioFunc, 30000);
        });
      };
      audioFunc();
    });

    onUnmounted(() => {
      state.closeAudio = false;
    });

    //  请求管理信息
    const updateAdminInfo = () => {
      adminInfoAPI()
        .then((res: any) => {
          $initStore.userInfo = res;
          state.adminInfo = res;
        })
        .catch(() => {
          $router.push({ name: 'Home' });
        });
    };

    //  弹窗操作
    const dialogOpenFunc = (configKey: string) => {
      let configInfo = state.dialogList.get(configKey);
      if (configInfo && configKey == dialogUserInfoKey) {
        configInfo.params = {
          avatar: state.adminInfo.avatar,
          nickname: state.adminInfo.nickname,
          email: state.adminInfo.email,
        };
      }
      domDialogRef.value.setConfig(configInfo);
      domDialogRef.value.open();
    };

    //  回调信息
    const dialogDoneFunc = (dialogId: string) => {
      if (dialogId == dialogUserInfoKey) {
        updateAdminInfo();
      }
    };

    // 管理员退出
    const adminLogout = () => {
      $initStore.updateUserToken('');
      setTimeout(() => {
        $router.push({ name: 'Login' });
      }, 100);
    };

    // 计算剩余天数
    const remainingDays = () => {
      if (state.adminInfo) {
        if (state.adminInfo.expiredAt <= 0) {
          return 365;
        }

        const nowTime = Date.parse(String(new Date())) / 1000;
        return Math.floor((state.adminInfo.expiredAt - nowTime) / 86400);
      }
    };

    //  左侧开关
    const toggleLeftDrawer = () => {
      leftDrawerOpen.value = !leftDrawerOpen.value;
    };

    //  监听菜单是否变化
    watch(
      () => $initStore.menuList,
      (menuList) => {
        state.menuList = menuList;
      }
    );

    //  监听管理信息变化
    watch(
      () => $initStore.userInfo,
      (adminInfo: any) => {
        state.adminInfo = adminInfo;
      }
    );

    return {
      ...toRefs(state),
      DialogComponent,
      MenuComponent,
      imageSrc,
      leftDrawerOpen,
      rightDrawerOpen,
      domDialogRef,
      toggleLeftDrawer,
      dialogOpenFunc,
      dialogDoneFunc,
      dialogUserInfoKey,
      dialogUpdatePassword,
      dialogUpdateSecurityPassword,
      adminLogout,
      remainingDays,
    };
  },
};
</script>
