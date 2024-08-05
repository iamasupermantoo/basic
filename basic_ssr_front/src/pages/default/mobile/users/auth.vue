<template>
  <div class="column full-height">
    <div class="col bg-white q-pa-md">
      <div v-if="params.status == authStatus.refuse">
        <q-banner rounded class="bg-red text-white q-mt-sm">
          {{ params.data }}
        </q-banner>
      </div>
      <div class="text-weight-medium q-pb-sm q-mt-lg">{{ $t('idName') }}</div>
      <q-input :placeholder="$t('idName')" outlined
        :readonly="params.status == authStatus.pending || params.status == authStatus.complete" v-model="params.realName"
        class="q-mb-md" />
      <div class="text-weight-medium q-pb-sm">{{ $t('idNumber') }}</div>
      <q-input :placeholder="$t('idNumber')" outlined
        :readonly="params.status == authStatus.pending || params.status == authStatus.complete" v-model="params.number"
        class="q-mb-md" />
      <div class="text-weight-medium q-pb-sm">{{ $t('telephone') }}</div>
      <q-input :placeholder="$t('telephone')" outlined
        :readonly="params.status == authStatus.pending || params.status == authStatus.complete" v-model="params.telephone"
        class="q-mb-md" />
      <div class="text-weight-medium q-pb-sm">{{ $t('address') }}</div>
      <q-input :placeholder="$t('address')" outlined
        :readonly="params.status == authStatus.pending || params.status == authStatus.complete" v-model="params.address"
        class="q-mb-md" />

      <div class="text-weight-medium">{{ $t('idPhoto') }}</div>

      <div class="col">
        <uploader :respValue="params.photo1" @uploaded="(url) => { params.photo1 = url }">
          <template v-slot:default>
            <q-uploader-add-trigger v-if="params.status == authStatus.start || params.status == authStatus.refuse" />
            <q-card flat>
              <div class="column items-center justify-center" style="height: 150px;border: grey 1px dashed">
                <q-icon name="add" color="grey" size="30px" v-if="params.photo1 == ''" />
                <q-img no-spinner v-else :src="imageSrc(params.photo1)"></q-img>
              </div>
            </q-card>
          </template>
        </uploader>
        <div class="text-caption text-weight-medium text-center q-mt-sm">{{ $t('idPhoto2') }}</div>
      </div>

      <div class="col q-mt-lg">
        <uploader :respValue="params.photo2" @uploaded="(url) => { params.photo2 = url }">
          <template v-slot:default>
            <q-uploader-add-trigger v-if="params.status == authStatus.start || params.status == authStatus.refuse" />
            <q-card flat>
              <div class="column items-center justify-center" style="height: 150px;border: grey 1px dashed">
                <q-icon name="add" color="grey" size="30px" v-if="params.photo2 == ''" />
                <q-img no-spinner v-else :src="imageSrc(params.photo2)"></q-img>
              </div>
            </q-card>
          </template>
        </uploader>
        <div class="text-caption text-weight-medium text-center q-mt-sm">{{ $t('idPhoto2') }}</div>
      </div>

      <q-btn v-if="params.status == authStatus.start" unelevated rounded color="primary" size="lg" :label="$t('submit')"
        class="full-width q-my-xl" no-caps @click="submitFunc" />
      <q-btn v-if="params.status == authStatus.refuse" unelevated rounded color="primary" size="lg"
        :label="$t('reSubmit')" class="full-width q-my-xl" no-caps @click="submitFunc" />
    </div>
  </div>
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
  name: 'UserRealAuth',
  components: { uploader },
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('realAuth'),
    })

    const $userStore = UserStore()

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
      imageSrc,
      ...toRefs(state),
      submitFunc,
    }
  }
});
</script>

<style lang="scss" scoped></style>
