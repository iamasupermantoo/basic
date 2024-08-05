<template>
  <div style="padding: 48px 100px;">
    <div class="text-weight-bold text-h5">{{ $t('personalSetting') }}</div>

    <div class="q-mt-lg">
      <div class="row justify-between items-center q-mb-lg" v-for="(setting, settingIndex) in settingsList"
        :key="settingIndex">
        <div class="col-6">
          <div class="text-body1">{{ $t(setting.name) }}</div>
          <div class="text-grey text-caption">{{ $t(setting.desc) }}</div>
        </div>
        <div class="col">
          <div v-if="setting.type == 'password'">
            <div>******</div>
          </div>
          <div v-else-if="setting.type == 'avatar'">
            <q-avatar>
              <q-img no-spinner :src="imageSrc(setting.value)"></q-img>
            </q-avatar>
          </div>
          <div v-else-if="setting.type == 'datePicker'">
            {{ typeof (setting.value) == 'string' ? setting.value : date.formatDate(setting.value * 1000, 'YYYY/MM/DD') }}
          </div>
          <div v-else class="text-body2">
            <div v-if="setting.params == 'sex'">
              <div v-if="setting.value == 1">{{ $t('male') }}</div>
              <div v-else-if="setting.value == 2">{{ $t('female') }}</div>
              <div v-else>{{ $t('unknown') }}</div>
            </div>
            <div v-else>
              <div v-if="setting.value == ''">- -</div>
              <div v-else>{{ setting.value }}</div>
            </div>
          </div>
        </div>
        <div>
          <q-btn flat :label="$t('edit')" class="bg-grey-1 text-grey-8" @click="openUpdateDialog(setting)"></q-btn>
        </div>
      </div>
    </div>
    <q-dialog v-model="dialogShow">
      <q-card style="width: 300px">
        <q-card-section>
          <div class="text-h6">{{ $t('edit') }}{{ $t(currentSetting.name) }}</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <div class="relative row justify-center" v-if="currentSetting.params == 'avatar'">
            <uploader @uploaded="(url) => { params[currentSetting.params] = url; currentSetting.value = url }">
              <template v-slot:default>
                <q-uploader-add-trigger />
                <q-avatar style="width: 80px;height: 80px;">
                  <q-img no-spinner :src="imageSrc(currentSetting.value)" />
                  <q-badge floating class="bg-transparent" :style="{ top: '70%' }">
                    <q-img no-spinner src='/images/icons/edit.png' width="28px" height="28px"></q-img>
                  </q-badge>
                </q-avatar>
              </template>
            </uploader>

          </div>

          <div v-else-if="currentSetting.params == 'sex'">
            <q-select outlined v-model="params[currentSetting.params]" :options="sexList" option-value="value"
              option-label="name" map-options emit-value dropdown-icon="expand_more" />
          </div>

          <div v-else-if="currentSetting.params == 'birthday'">
            <q-input outlined v-model="params[currentSetting.params]" :label="$t(currentSetting.name)" mask="date">
              <template v-slot:append>
                <q-icon name="event" class="cursor-pointer">
                  <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                    <q-date v-model="params[currentSetting.params]">
                      <div class="row items-center justify-end">
                        <q-btn v-close-popup :label="$t('confirm')" color="primary" flat />
                      </div>
                    </q-date>
                  </q-popup-proxy>
                </q-icon>
              </template>
            </q-input>
          </div>

          <div v-else-if="currentSetting.params == 'password' || currentSetting.params == 'secretKey'">
            <q-input v-model="params.oldPassword" outlined class="q-mb-md" :type="showPwd.oldPwd ? 'text' : 'password'"
              :label="$t('oldPassword')">
              <template v-slot:append>
                <q-icon @click="showPwd.oldPwd = !showPwd.oldPwd"
                  :name="showPwd.oldPwd ? 'o_visibility' : 'o_visibility_off'"></q-icon>
              </template>
            </q-input>
            <q-input v-model="params.newPassword" outlined class="q-mb-md" :type="showPwd.newPwd ? 'text' : 'password'"
              :label="$t('newPassword')">
              <template v-slot:append>
                <q-icon @click="showPwd.newPwd = !showPwd.newPwd"
                  :name="showPwd.newPwd ? 'o_visibility' : 'o_visibility_off'"></q-icon>
              </template>
            </q-input>
            <q-input v-model="params.cmfPassword" outlined class="q-mb-md" :type="showPwd.cmfPwd ? 'text' : 'password'"
              :label="$t('cmfPassword')">
              <template v-slot:append>
                <q-icon @click="showPwd.cmfPwd = !showPwd.cmfPwd"
                  :name="showPwd.cmfPwd ? 'o_visibility' : 'o_visibility_off'"></q-icon>
              </template>
            </q-input>
          </div>

          <div v-else-if="currentSetting.params == 'telephone'">
            <div class="row no-wrap justify-between">
              <q-btn-dropdown class="col-4" color="grey" outline dropdown-icon="expand_more">
                <template v-slot:label>
                  <div class="row no-wrap items-center">
                    <q-img no-spinner :src="imageSrc(countryList[currentCountryIndex].icon)" width="18px" height="14px" />
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
                <q-input class="q-ml-xs" outlined v-model="params[currentSetting.params]" :label="$t('telephone')" />
              </div>
            </div>
          </div>

          <div v-else-if="currentSetting.params == 'desc'">
            <q-input outlined type="textarea" v-model="params[currentSetting.params]"
              :label="$t(currentSetting.name)"></q-input>
          </div>

          <div v-else>
            <q-input outlined v-model="params[currentSetting.params]" :label="$t(currentSetting.name)"></q-input>
          </div>
        </q-card-section>

        <q-card-actions align="right" class="q-mt-md">
          <q-btn @click="updateUserInfo" flat :label="$t('cancel')" color="grey" v-close-popup />
          <q-btn @click="submitFunc" flat :label="$t('submit')" color="primary" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { userInfoAPI, updateInfoAPI, updatePasswordAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { NotifyNegative, NotifyPositive } from 'src/utils/notify';
import { UserStore } from 'stores/user';
import { date } from 'quasar';
import uploader from 'src/components/uploader.vue';
import { useI18n } from 'vue-i18n';
import { InitStore } from 'stores/init';

export default defineComponent({
  name: 'SettingIndex',
  components: {
    uploader,
  },
  setup() {
    const $userStore = UserStore()
    const $initStore = InitStore()
    const { t } = useI18n()

    const state = reactive({
      // 性别列表
      sexList: [
        { name: t('unknown'), value: 0 },
        { name: t('male'), value: 1 },
        { name: t('female'), value: 2 },
      ],
      currentTelephone: [],

      // 是否显示密码
      showPwd: {
        oldPwd: false,
        newPwd: false,
        cmfPwd: false
      },

      // 弹窗数据
      countryList: $initStore.countryList,
      currentCountryIndex: 0 as any,
      dialogShow: false,
      currentSetting: {} as any,

      // 接口参数
      params: {} as any,
      settingsList: [] as any
    });

    onMounted(() => {
      const userInfo = $userStore.userInfo
      state.settingsList = [
        { name: 'avatar', params: 'avatar', type: 'avatar', desc: 'avatarSmall', value: userInfo.avatar },
        { name: 'nickname', params: 'nickname', type: 'text', desc: 'nicknameSmall', value: userInfo.nickname },
        { name: 'email', params: 'email', type: 'text', desc: 'emailSmall', value: userInfo.email },
        { name: 'password', params: 'password', type: 'password', desc: 'passwordSmall', value: '' },
        { name: 'secretKey', params: 'secretKey', type: 'password', desc: 'secretKeySmall', value: '' },
        { name: 'telephone', params: 'telephone', type: 'telephone', desc: 'telephoneSmall', value: userInfo.telephone },
        { name: 'sex', params: 'sex', type: 'toggle', desc: 'sexSmall', value: userInfo.sex },
        { name: 'birthday', params: 'birthday', type: 'datePicker', desc: 'birthdaySmall', value: userInfo.birthday },
        { name: 'personalSignature', params: 'desc', type: 'textarea', desc: 'personalSignatureSmall', value: userInfo.desc },
      ]
    })

    // 打开弹窗
    const openUpdateDialog = (setting: any) => {
      state.dialogShow = true
      state.currentSetting = setting
      state.params = {}
      state.params[state.currentSetting.params] = state.currentSetting.value
      if (state.currentSetting.params == 'birthday') {
        state.params[state.currentSetting.params] = ''
      }
      // 如果是手机, 那么默认选中
      if (state.currentSetting.params == 'telephone' && state.currentSetting.value != '') {
        state.currentTelephone = state.currentSetting.value.split('|')
        if (state.currentTelephone.length == 2) {
          for (const countryKey in state.countryList) {
            if (state.countryList[countryKey].code == state.currentTelephone[0]) {
              state.currentCountryIndex = countryKey
            }
          }
          state.params[state.currentSetting.params] = state.currentTelephone[1]
        }
      }
    }

    // 更新用户信息
    const submitFunc = () => {
      // 修改登录密码
      if (state.currentSetting.params == 'password' || state.currentSetting.params == 'secretKey') {
        if (state.params.newPassword != state.params.cmfPassword) {
          NotifyNegative(t('twoPasswordsAreDifferent'))
          return false
        }
        state.params.type = state.currentSetting.params == 'password' ? 1 : 2
        updatePasswordAPI(state.params).then(() => {
          updateUserInfo()
          state.dialogShow = false
          NotifyPositive(t('submittedSuccess'))
        })
        return
      }

      // 更新手机号码
      if (state.currentSetting.params == 'telephone') {
        state.params['telephone'] = state.countryList[state.currentCountryIndex].code + '|' + state.params['telephone']
      }
      // 更新生日
      if (state.currentSetting.params == 'birthday') {
        state.params['birthday'] = date.formatDate(state.params[state.currentSetting.params], 'X')
      }
      updateInfoAPI(state.params).then(() => {
        updateUserInfo()
        state.dialogShow = false
        state.currentSetting.value = state.params[state.currentSetting.params]

        NotifyPositive(t('submittedSuccess'))
      })
    }

    // 更新当前用户信息
    const updateUserInfo = () => {
      userInfoAPI().then((res: any) => {
        $userStore.updateUserInfo(res)
      })
    }

    return {
      imageSrc,
      date,
      ...toRefs(state),
      updateUserInfo,
      openUpdateDialog,
      submitFunc,
    }
  }
});
</script>
<style lang="scss" scoped></style>
