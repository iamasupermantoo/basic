<template>
  <!-- 语言切换 -->
  <q-header v-if="config.settings.lang.showLogin" class="bg-white">
    <q-toolbar>
      <q-space />
      <q-btn class="text-grey-8" rounded no-caps flat>
        <q-img no-spinner width="24px" height="24px" class="q-mr-sm" :src="imageSrc(currentLangInfo.icon)"></q-img>
        <div>{{ currentLangInfo.name }}</div>
        <switchLanguage></switchLanguage>
      </q-btn>
    </q-toolbar>
  </q-header>

  <div>
    <!-- logo -->
    <div class="row justify-center">
      <q-img no-spinner class="q-mt-lg q-mb-md" width="70px" height="70px" :src="imageSrc(config.logo)" />
    </div>
    <div class="row justify-center">
      <div class="text-weight-bold text-h6">{{ $t('loginSmall') }}</div>
    </div>

    <q-form class="q-mt-lg q-px-lg">
      <!-- 账号 -->
      <q-input class="q-mb-md" outlined v-model="params.username"
        :input-style="{ fontSize: '16px', color: '#999999!important' }" :placeholder="$t('username')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/username.png" />
        </template>
      </q-input>

      <!-- 密码 -->
      <q-input class="q-mb-md" v-model="params.password" outlined
        :input-style="{ fontSize: '16px', color: '#999999!important' }" :type="showPwd ? 'text' : 'password'"
        :placeholder="$t('password')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/password.png" />
        </template>
        <template v-slot:append>
          <q-icon style="color: #999999" :name="showPwd ? 'visibility' : 'visibility_off'" class="cursor-pointer"
            @click="showPwd = !showPwd" />
        </template>
      </q-input>

      <!-- 验证码 -->
      <q-input v-if="config.settings.login.showVerify" :input-style="{ fontSize: '16px', color: '#999999!important' }"
        class="q-mb-sm" outlined v-model="params.captchaVal" :placeholder="$t('code')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/code.png" />
        </template>
        <template v-slot:append>
          <q-img no-spinner v-if="params.captchaId !== ''" :src="baseURL + '/captcha/' + params.captchaId + '/150-50'"
            width="150px" height="50px" @click="refreshCaptchaFunc"></q-img>
        </template>
      </q-input>

      <!-- 忘记密码、登录、注册 -->
      <div class="text-right q-mb-lg text-grey-7 cursor-pointer">{{ $t('forgotPassword') }}</div>
      <q-btn @click="submitFunc()" class="full-width q-mb-lg" unelevated rounded no-caps size="lg" color="primary"
        :label="$t('login')" />
      <div @click="$router.push({ name: 'UserRegister' })" v-if="config.settings.login.showRegister"
        class="text-center text-primary q-mb-xl cursor-pointer">
        {{ $t('toRegister') }}
      </div>
    </q-form>
  </div>

  <!-- 客服图标 -->
  <q-avatar v-if="config.settings.online.showLogin" class="fixed-bottom-right q-mb-md q-mr-md shadow-1">
    <img :src="imageSrc(config.onlineIcon)" alt="">
  </q-avatar>
</template>

<script lang="ts">
import switchLanguage from 'src/components/switchLanguage.vue';
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { captchaAPI } from 'src/apis';
import { userLoginAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';
import { UserStore } from 'src/stores/user';

export default defineComponent({
  name: 'userLogin',
  components: {
    switchLanguage
  },
  setup() {
    const $router = useRouter();
    const $initStore = InitStore();
    const $userStore = UserStore();

    const state = reactive({
      baseURL: process.env.baseURL,

      // 当前语言信息
      currentLangInfo: $initStore.languageList.find((item: any) => item.alias == $initStore.userLang),

      // 初始化配置信息
      config: $initStore.config,

      // 是否显示密码
      showPwd: false,

      // 提交参数
      params: {
        username: '', //用户名
        password: '', //密码
        captchaId: '', //验证Id
        captchaVal: '', // 验证码
      },
    });

    onMounted(() => {
      refreshCaptchaFunc();
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
        .then((res: any) => {
          $userStore.updateUserInfo(res.userInfo)
          $initStore.updateUserToken(res.token);
          void $router.push({ name: 'HomeIndex' });
        })
        .catch(() => {
          refreshCaptchaFunc();
        });
    };

    return {
      imageSrc,
      ...toRefs(state),
      refreshCaptchaFunc,
      submitFunc,
    };
  },
});
</script>
<style lang="scss" scoped></style>
