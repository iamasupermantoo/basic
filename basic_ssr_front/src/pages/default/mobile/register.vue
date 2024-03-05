<template>
  <!-- 语言切换 -->
  <q-header v-if="config.settings.lang.showRegister" class="bg-white">
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
      <q-img no-spinner class="q-mt-lg q-mb-md" width="70px" height="70px" :src="`${imageSrc(config.logo)}`" />
    </div>
    <div class="row justify-center">
      <div class="text-h6 text-weight-bold">{{ $t('registerSmall') }}</div>
    </div>

    <q-form class="q-mt-lg q-px-lg">
      <!-- 邮箱 -->
      <q-input v-if="config.settings.register.showEmail" :input-style="{ fontSize: '16px', color: '#999999!important' }"
        outlined class="q-mb-md" v-model="params.email" :placeholder="$t('email')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/email.png" />
        </template>
      </q-input>

      <!-- 账号 -->
      <q-input outlined class="q-mb-md" :input-style="{ fontSize: '16px', color: '#999999!important' }"
        v-model="params.username" :placeholder="$t('username')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/username.png" />
        </template>
      </q-input>

      <!-- 密码 -->
      <q-input class="q-mb-md" :input-style="{ fontSize: '16px', color: '#999999!important' }" v-model="params.password"
        outlined :type="showTextPassword.password ? 'text' : 'password'" :placeholder="$t('password')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/password.png" />
        </template>
        <template v-slot:append>
          <q-icon class="text-grey-7 cursor-pointer" :name="showTextPassword.password ? 'visibility' : 'visibility_off'"
            @click="showTextPassword.password = !showTextPassword.password" />
        </template>
      </q-input>

      <!-- 确认密码 -->
      <q-input v-if="config.settings.register.showCmfPass" :input-style="{ fontSize: '16px', color: '#999999!important' }"
        class="q-mb-md" v-model="params.cmfPassword" outlined :type="showTextPassword.cmfPassword ? 'text' : 'password'"
        :placeholder="$t('cmfPassword')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/password.png" />
        </template>
        <template v-slot:append>
          <q-icon class="text-grey-7 cursor-pointer"
            :name="showTextPassword.cmfPassword ? 'visibility' : 'visibility_off'"
            @click="showTextPassword.cmfPassword = !showTextPassword.cmfPassword" />
        </template>
      </q-input>

      <!-- 验证码 -->
      <q-input v-if="config.settings.register.showVerify" :input-style="{ fontSize: '16px', color: '#999999!important' }"
        class="q-mb-md" outlined v-model="params.captchaVal" :placeholder="$t('code')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/code.png" />
        </template>
        <template v-slot:append>
          <q-img no-spinner v-if="params.captchaId !== ''" :src="baseURL + '/captcha/' + params.captchaId + '/150-50'"
            width="150px" height="50px" @click="refreshCaptchaFunc"></q-img>
        </template>
      </q-input>

      <!-- 安全秘钥 -->
      <q-input v-if="config.settings.register.showSecurityPass"
        :input-style="{ fontSize: '16px', color: '#999999!important' }" class="q-mb-md" outlined
        v-model="params.securityKey" :type="showTextPassword.securityKey ? 'text' : 'password'"
        :placeholder="$t('enterSecretKey')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/key.png" />
        </template>
        <template v-slot:append>
          <q-icon class="text-grey-7 cursor-pointer"
            :name="showTextPassword.securityKey ? 'visibility' : 'visibility_off'"
            @click="showTextPassword.securityKey = !showTextPassword.securityKey" />
        </template>
      </q-input>

      <!-- 确认安全秘钥 -->
      <q-input v-if="config.settings.register.showSecurityPass"
        :input-style="{ fontSize: '16px', color: '#999999!important' }" class="q-mb-md" outlined
        v-model="params.cmfSecurityKey" :type="showTextPassword.cmfSecurityKey ? 'text' : 'password'"
        :placeholder="$t('enterSecretKey')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/key.png" />
        </template>
        <template v-slot:append>
          <q-icon class="text-grey-7 cursor-pointer"
            :name="showTextPassword.cmfSecurityKey ? 'visibility' : 'visibility_off'"
            @click="showTextPassword.cmfSecurityKey = !showTextPassword.cmfSecurityKey" />
        </template>
      </q-input>

      <!-- 邀请码 -->
      <q-input v-if="config.settings.register.isInvite" :input-style="{ fontSize: '16px', color: '#999999!important' }"
        class="q-mb-md" outlined v-model="params.code" :placeholder="$t('inviteCode')">
        <template v-slot:prepend>
          <q-img no-spinner width="24px" height="24px" src="/images/icons/profile.png" />
        </template>
      </q-input>

      <!-- 手机号码 -->
      <div v-if="config.settings.register.showTelephone" class="row no-wrap justify-between">
        <q-btn-dropdown class="col-4 text-weight-regular" unelevated flat no-caps dropdown-icon="expand_more"
          style="height: 50px;background: #f5f6fa;border-radius: 10px;color: #8F959E;">
          <template v-slot:label>
            <div class="row no-wrap items-center">
              <q-img no-spinner :src="imageSrc(countryList[currentCountryIndex].icon)" width="24px" height="16px" />
              <div class="q-ml-sm">+{{ countryList[currentCountryIndex].code }}</div>
            </div>
          </template>
          <!-- 下拉 -->
          <q-list style="min-width: 268px" class="q-py-sm">
            <q-item @click="currentCountryIndex = i" v-for="(item, i) in countryList" :key="i" clickable v-close-popup
              class="row no-wrap items-center">
              <q-img no-spinner class="q-mr-sm" :src="imageSrc(item.icon)" width="38px" height="38px" />
              <div>
                <div style="font-size: 16px;">{{ item.name }}</div>
              </div>
            </q-item>
          </q-list>
        </q-btn-dropdown>

        <q-input style="width: 64%;" :input-style="{ fontSize: '16px', color: '#999999!important' }" outlined
          v-model="params.telephone" :placeholder="$t('telephone')" />
      </div>

      <!-- 前往登录、点击注册 -->
      <q-btn @click="submitFunc()" class="full-width q-my-lg" unelevated rounded no-caps size="lg" color="primary"
        :label="$t('register')" />
      <div @click="$router.push('/login')" class="text-center q-pb-xl text-primary cursor-pointer">
        {{ $t('toLogin') }}
      </div>
    </q-form>
  </div>

  <!-- 客服图标 -->
  <q-avatar v-if="config.settings.online.showRegister" class="fixed-bottom-right q-mb-md q-mr-md shadow-1">
    <img :src="imageSrc(config.onlineIcon)" alt="">
  </q-avatar>
</template>

<script lang="ts">
import switchLanguage from 'src/components/switchLanguage.vue';
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { captchaAPI } from 'src/apis';
import { userRegisterAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';
import { UserStore } from 'src/stores/user';
import { NotifyNegative } from 'src/utils/notify';
import { useI18n } from 'vue-i18n';

export default defineComponent({
  name: 'userRegister',
  components: {
    switchLanguage
  },
  setup() {
    const $router = useRouter();
    const $route = useRoute();
    const $initStore = InitStore();
    const $userStore = UserStore();
    const { t } = useI18n();

    const state = reactive({
      baseURL: process.env.baseURL,

      // 当前设置的语言信息
      currentLangInfo: $initStore.languageList.find((item: any) => item.alias == $initStore.userLang),

      // 初始配置信息
      config: $initStore.config,

      // 地区选择
      currentCountryIndex: 0,
      countryList: $initStore.countryList as any,

      // 是否显示密码
      showTextPassword: {
        // 登录密码
        password: false,
        cmfPassword: false,
        // 安全码
        securityKey: false,
        cmfSecurityKey: false,
      },

      // 提交参数
      params: {
        username: '', //用户名
        password: '', //密码
        cmfPassword: '',  // 确认密码
        captchaId: '', //验证id
        captchaVal: '', // 验证码
        email: '', //邮箱
        telephone: '', //电话
        securityKey: '', //安全秘钥
        cmfSecurityKey: '',//确认安全密钥
        code: $route.query.code ? $route.query.code as any : '', //邀请码
      }
    });

    onMounted(() => {
      refreshCaptchaFunc();
    })

    // 获取验证码
    const refreshCaptchaFunc = () => {
      captchaAPI().then((res: any) => {
        state.params.captchaId = res;
      })
    };

    // 提交注册
    const submitFunc = () => {
      // 判断两次登录密码是否一致
      if (state.config.settings.register.showCmfPass) {
        if (state.params.password !== state.params.cmfPassword) {
          NotifyNegative(t('twoPasswordsAreDifferent'));
          return false
        }
      }

      // 判断两次安全秘钥是否一致
      if (state.config.settings.register.showSecurityPass) {
        if (state.params.securityKey !== state.params.cmfSecurityKey) {
          NotifyNegative(t('twoSecretKeyAreDifferent'));
          return false
        }
      }

      //拼接手机区号 
      if (state.config.settings.register.showTelephone) {
        state.params.telephone = state.countryList[state.currentCountryIndex].code + '|' + state.params.telephone
      }
      userRegisterAPI(state.params).then((res: any) => {
        $userStore.updateUserInfo(res.userInfo)
        $initStore.updateUserToken(res.token);
        void $router.push({ name: 'HomeIndex' });
      }).catch(() => {
        refreshCaptchaFunc();
      });
    }

    return {
      ...toRefs(state),
      refreshCaptchaFunc,
      submitFunc,
      imageSrc,
    }
  }
});
</script>
<style lang="scss" scoped></style>
