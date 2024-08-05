<template>
  <q-dialog v-model="registerShow">
    <q-card style="width: 410px">
      <q-card-section class="q-pa-lg">
        <div class="text-center text-weight-bold text-primary text-h4">
          {{ $t('register') }}
        </div>
        <div class="text-center text-body1 text-weight-regular q-mt-sm text-grey">
          {{ $t('registerSmall') }}
        </div>

        <q-form class="q-mt-xl q-gutter-sm">
          <!-- 邮箱 -->
          <q-input v-if="config.settings.register.showEmail" :input-style="{ fontSize: '16px' }" outlined
            :label="$t('email')" v-model="params.email">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/email2.png" />
            </template>
          </q-input>

          <!-- 账号 -->
          <q-input outlined v-model="params.username" :input-style="{ fontSize: '16px' }" :label="$t('username')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/username.png" />
            </template>
          </q-input>

          <!-- 密码 -->
          <q-input v-model="params.password" outlined :input-style="{ fontSize: '16px' }"
            :type="showTextPassword.password ? 'text' : 'password'" :label="$t('password')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/password.png" />
            </template>
            <template v-slot:append>
              <q-icon class="text-grey-7 cursor-pointer"
                :name="showTextPassword.password ? 'visibility' : 'visibility_off'"
                @click="showTextPassword.password = !showTextPassword.password" />
            </template>
          </q-input>

          <!-- 确认密码 -->
          <q-input v-if="config.settings.register.showCmfPass" outlined :input-style="{ fontSize: '16px' }"
            v-model="params.cmfPassword" :type="showTextPassword.cmfPassword ? 'text' : 'password'"
            :label="$t('cmfPassword')">
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
          <q-input v-if="config.settings.register.showVerify" outlined :input-style="{ fontSize: '16px' }"
            v-model="params.captchaVal" :label="$t('code')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/code.png" />
            </template>
            <template v-slot:append>
              <q-img no-spinner v-if="params.captchaId !== ''" :src="baseURL + '/captcha/' + params.captchaId + '/150-50'"
                width="150px" height="50px" @click="refreshCaptchaFunc" class="cursor-pointer"></q-img>
            </template>
          </q-input>

          <!-- 安全秘钥 -->
          <q-input v-if="config.settings.register.showSecurityPass" :input-style="{ fontSize: '16px' }" outlined
            v-model="params.securityKey" :type="showTextPassword.securityKey ? 'text' : 'password'"
            :label="$t('enterSecretKey')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/key2.png" />
            </template>
            <template v-slot:append>
              <q-icon class="text-grey-7 cursor-pointer"
                :name="showTextPassword.securityKey ? 'visibility' : 'visibility_off'"
                @click="showTextPassword.securityKey = !showTextPassword.securityKey" />
            </template>
          </q-input>

          <!-- 确认安全秘钥 -->
          <q-input v-if="config.settings.register.showSecurityPass" :input-style="{ fontSize: '16px' }" outlined
            v-model="params.cmfSecurityKey" :type="showTextPassword.cmfSecurityKey ? 'text' : 'password'"
            :label="$t('enterSecretKey')">
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
          <q-input v-if="config.settings.register.isInvite" outlined :input-style="{ fontSize: '16px' }"
            v-model="params.code" :label="$t('inviteCode')">
            <template v-slot:prepend>
              <q-img no-spinner width="24px" height="24px" src="/images/icons/profile.png" />
            </template>
          </q-input>

          <!-- 手机号码 -->
          <div v-if="config.settings.register.showTelephone" class="row no-wrap justify-between">
            <q-btn-dropdown class="col-4" color="grey" outline no-caps dropdown-icon="expand_more" style="height: 56px;">
              <template v-slot:label>
                <div class="row no-wrap items-center">
                  <q-img no-spinner :src="imageSrc(countryList[currentCountryIndex].icon)" width="24px" height="16px" />
                  <div class="q-ml-sm">+{{ countryList[currentCountryIndex].code }}</div>
                </div>
              </template>
              <!-- 下拉 -->
              <q-list class="q-py-sm">
                <q-item @click="currentCountryIndex = countryIndex" v-for="(country, countryIndex) in countryList"
                  :key="countryIndex" clickable v-close-popup class="row no-wrap items-center">
                  <q-img no-spinner class="q-mr-sm" :src="imageSrc(country.icon)" width="38px" height="38px" />
                  <div>
                    <div style="font-size: 16px;">{{ country.name }}</div>
                  </div>
                </q-item>
              </q-list>
            </q-btn-dropdown>

            <div class="col-8">
              <q-input :input-style="{ fontSize: '16px' }" outlined class="q-ml-sm" v-model="params.telephone"
                :label="$t('telephone')" />
            </div>
          </div>

          <!-- 前往登录、点击注册 -->
          <q-btn @click="submitFunc()" class="full-width q-mt-xl" unelevated rounded no-caps style="height: 44px;"
            color="primary" :label="$t('register')" />
          <div @click="toLogin()" class="text-center q-pb-sm text-primary cursor-pointer">
            {{ $t('toLogin') }}
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { captchaAPI } from 'src/apis';
import { userRegisterAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';
import { UserStore } from 'stores/user';
import { NotifyNegative } from 'src/utils/notify';
import { useI18n } from 'vue-i18n';

export default defineComponent({
  name: 'registerDialog',
  setup(props: any, context: any) {
    const $router = useRouter();
    const $route = useRoute();
    const $initStore = InitStore();
    const $userStore = UserStore()
    const { t } = useI18n();

    const state = reactive({
      baseURL: process.env.baseURL,

      // 初始配置信息
      config: $initStore.config,

      // 地区选择
      currentCountryIndex: 0,
      countryList: $initStore.countryList as any,

      // 确认密码
      confirmPassword: '',

      // 是否显示密码
      showTextPassword: {
        // 登录密码
        password: false,
        cmfPassword: false,
        // 安全码
        securityKey: false,
        cmfSecurityKey: false,
      },

      //提交参数
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
        code: $route.query.code ?? '', //邀请码
      } as any,

      // 注册弹窗
      registerShow: false,
    });

    // 获取验证码
    const refreshCaptchaFunc = () => {
      captchaAPI().then((res: any) => {
        state.params.captchaId = res;
      })
    };

    // 注册弹窗
    const openDialog = (status: boolean) => {
      refreshCaptchaFunc();
      state.registerShow = status;
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
      //  用户注册
      userRegisterAPI(state.params).then(async (res: any) => {
        $userStore.updateUserInfo(res.userInfo)
        await $initStore.updateUserToken(res.token);
        state.registerShow = false;
        void $router.push({ name: 'HomeIndex' });
      }).catch(() => {
        refreshCaptchaFunc();
      });
    }

    // 点击注册
    const toLogin = () => {
      context.emit('switchDialogFunc', true, false);
    }

    return {
      ...toRefs(state),
      toLogin,
      openDialog,
      refreshCaptchaFunc,
      submitFunc,
      imageSrc,
    };
  },
});
</script>

<style lang="scss" scoped></style>
