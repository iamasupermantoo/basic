<template>
  <div style="margin-top: 60px;">
    <div class="row justify-center">
      <div class="col-6">
        <div v-if="params.status == authStatus.refuse">
          <q-banner rounded class="bg-red text-white q-mt-sm">
            {{ params.data }}
          </q-banner>
        </div>
        <div class="text-h5 q-mt-md">{{ $t('realAuth') }}</div>
        <div class="q-gutter-md q-mt-md">
          <div>
            <div class="q-mb-xs">{{ $t('idName') }}</div>
            <q-input outlined v-model="params.realName"
              :readonly="params.status == authStatus.pending || params.status == authStatus.complete"
              :placeholder="$t('idName')" />
          </div>
          <div>
            <div class="q-mb-xs">{{ $t('idNumber') }}</div>
            <q-input outlined v-model="params.number"
              :readonly="params.status == authStatus.pending || params.status == authStatus.complete"
              :placeholder="$t('idNumber')" />
          </div>
          <div>
            <div class="q-mb-xs">{{ $t('telephone') }}</div>
            <q-input outlined v-model="params.telephone"
              :readonly="params.status == authStatus.pending || params.status == authStatus.complete"
              :placeholder="$t('telephone')" />
          </div>
          <div>
            <div class="q-mb-xs">{{ $t('address') }}</div>
            <q-input outlined v-model="params.address"
              :readonly="params.status == authStatus.pending || params.status == authStatus.complete"
              :placeholder="$t('address')" />
          </div>


          <div>
            <div>{{ $t('idPhoto') }}</div>

            <div class="row q-mt-sm">
              <div class="col q-mr-xs">
                <uploader :respValue="params.photo1" @uploaded="(url) => { params.photo1 = url }">
                  <template v-slot:default>
                    <q-uploader-add-trigger
                      v-if="params.status == authStatus.start || params.status == authStatus.refuse" />
                    <q-card flat>
                      <div class="column items-center justify-center" style="height: 150px;border: grey 1px dashed">
                        <q-icon name="add" color="grey" size="40px" v-if="params.photo1 == ''" />
                        <q-img no-spinner v-else :src="imageSrc(params.photo1)"></q-img>
                      </div>
                    </q-card>
                  </template>
                </uploader>
                <div class="text-caption text-weight-medium text-center q-mt-sm">{{ $t('idPhoto1') }}</div>
              </div>

              <!--  -->
              <div class="col q-ml-xs">
                <uploader :respValue="params.photo2" @uploaded="(url) => { params.photo2 = url }">
                  <template v-slot:default>
                    <q-uploader-add-trigger
                      v-if="params.status == authStatus.start || params.status == authStatus.refuse" />
                    <q-card flat>
                      <div class="column items-center justify-center" style="height: 150px;border: grey 1px dashed">
                        <q-icon name="add" color="grey" size="40px" v-if="params.photo2 == ''" />
                        <q-img no-spinner v-else :src="imageSrc(params.photo2)"></q-img>
                      </div>
                    </q-card>
                  </template>
                </uploader>
                <div class="text-caption text-weight-medium text-center q-mt-sm">{{ $t('idPhoto2') }}</div>
              </div>
            </div>
          </div>

          <div class="q-mt-xl row justify-end">
            <q-btn rounded unelevated color="primary" no-caps :label="$t('submit')"
              v-if="params.status == authStatus.start" style="min-width: 100px" size="md" @click="submitFunc"></q-btn>

            <q-btn rounded unelevated color="primary" no-caps :label="$t('reSubmit')"
              v-if="params.status == authStatus.refuse" style="min-width: 100px" size="md" @click="submitFunc"></q-btn>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- sub按钮 -->
</template>

<script lang="ts">
import uploader from 'components/uploader.vue';
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { NotifyPositive } from 'src/utils/notify';
import { realAuthCreateAPI, realAuthInfoAPI } from 'src/apis/user';
import { useI18n } from 'vue-i18n';
import { imageSrc } from 'src/utils';
import { UserStore } from 'stores/user';

export default defineComponent({
  methods: { imageSrc },
  components: {
    uploader
  },
  name: 'UserRealAuth',
  setup() {
    const $userStore = UserStore()
    const { t } = useI18n()
    const state = reactive({
      authStatus: { pending: 10, complete: 20, refuse: -1, start: 0 },
      params: {} as any,
    });

    onMounted(() => {
      realAuthInfoAPI().then((res: any) => {
        state.params = res
      })
    })

    const submitFunc = () => {
      realAuthCreateAPI(state.params).then(() => {
        state.params.status = state.authStatus.pending
        NotifyPositive(t('submittedSuccess'))

        //  更新用户信息状态
        const userInfo = $userStore.userInfo
        userInfo.authStatus = state.authStatus.pending
        $userStore.updateUserInfo(userInfo)
      })
    }

    return {
      ...toRefs(state),
      submitFunc,
    }
  }
});
</script>
<style lang="scss" scoped></style>
