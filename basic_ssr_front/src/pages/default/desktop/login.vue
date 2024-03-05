<template>
  <q-dialog v-model="LoginShow">
    <q-card style="width: 410px">
      <q-card-section class="q-pa-lg">
        <!--  -->
        <div class="text-center text-weight-bold text-primary text-h4">
          {{ $t('login') }}
        </div>
        <div class="text-center text-body1 text-weight-regular q-mt-sm text-grey">
          {{ $t('loginSmall') }}
        </div>

        <q-form class="q-mt-xl q-gutter-sm">
          <!-- 账号 -->
          <q-input :input-style="{ fontSize: '16px' }" outlined v-model="params.username" :label="$t('username')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/username.png" />
            </template>
          </q-input>

          <!-- 密码 -->
          <q-input :input-style="{ fontSize: '16px', }" outlined v-model="params.password"
            :type="isPwd ? 'text' : 'password'" :label="$t('password')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/password.png" />
            </template>
            <template v-slot:append>
              <q-icon style="color: #999999" :name="isPwd ? 'visibility' : 'visibility_off'" class="cursor-pointer"
                @click="isPwd = !isPwd" />
            </template>
          </q-input>

          <!-- 验证码 -->
          <q-input :input-style="{ fontSize: '16px' }" v-if="config.settings.login.showVerify" outlined
            v-model="params.captchaVal" :label="$t('code')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/code.png" />
            </template>
            <template v-slot:append>
              <q-img no-spinner v-if="params.captchaId !== ''" :src="baseURL + '/captcha/' + params.captchaId + '/150-50'"
                width="150px" height="50px" @click="refreshCaptchaFunc" class="cursor-pointer"></q-img>
            </template>
          </q-input>

          <!-- 忘记密码、登录、注册 -->
          <div class="row justify-end q-mb-lg">
            <div class="text-grey-7 cursor-pointer">{{ $t('forgotPassword') }}</div>
          </div>
          <q-btn @click="submitFunc()" class="full-width" unelevated rounded no-caps style="height: 44px" color="primary"
            :label="$t('login')" />
          <div @click="toRegister()" v-if="config.settings.login.showRegister"
            class="text-center text-primary q-mb-sm cursor-pointer">
            {{ $t('toRegister') }}
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import { useRouter } from 'vue-router';
import { captchaAPI } from 'src/apis';
import { userLoginAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';
import { UserStore } from 'stores/user';

export default defineComponent({
  name: 'userLogin',
  setup(props: any, context: any) {
    const $router = useRouter();
    const $initStore = InitStore();
    const $userStore = UserStore()

    const state = reactive({
      baseURL: process.env.baseURL,

      // 初始化配置信息
      config: $initStore.config,

      // 登录弹窗
      LoginShow: false,

      // 是否显示密码
      isPwd: false,

      params: {
        username: '', //用户名
        password: '', //密码
        captchaId: '', //验证Id
        captchaVal: '', // 验证码
      },
    });

    // 获取验证码
    const refreshCaptchaFunc = () => {
      captchaAPI().then((res: any) => {
        state.params.captchaId = res;
      });
    };

    // 提交登录
    const submitFunc = () => {
      userLoginAPI(state.params)
        .then(async (res: any) => {
          state.LoginShow = false
          $userStore.updateUserInfo(res.userInfo)
          await $initStore.updateUserToken(res.token);
          void $router.push({ name: 'HomeIndex' });
        })
        .catch(() => {
          refreshCaptchaFunc();
        });
    };

    // 打开登录弹窗
    const openDialog = (status: boolean) => {
      refreshCaptchaFunc();
      state.LoginShow = status
    };

    // 点击注册
    const toRegister = () => {
      context.emit('switchDialogFunc', false, true);
    }

    return {
      imageSrc,
      ...toRefs(state),
      refreshCaptchaFunc,
      submitFunc,
      openDialog,
      toRegister,
    };
  },
});
</script>

<style lang="scss" scoped></style>
