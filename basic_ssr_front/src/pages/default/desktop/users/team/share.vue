<template>
  <div>
    <div class="row items-center justify-center"
      :style="{ background: 'url(/images/label-bg.png)', height: '200px', backgroundSize: 'cover', backgroundRepeat: 'no-repeat' }">
      <div class="text-white text-h4">{{ $t('inviteFriends') }}</div>
    </div>

    <div style="padding: 80px 0 260px">
      <div class="row justify-center items-center">
        <div class="col text-right">
          <q-img no-spinner class="q-mr-xl" width="300px" src="/images/download-bg.svg"></q-img>
        </div>
        <div class="col-1"></div>
        <div class="col">
          <div style="width: 310px" class="text-center">
            <div class="text-h3 text-bold">{{ config.name }}</div>
            <div class="text-h6 text-grey q-mt-sm">{{ $t('inviteFriendsSmall') }}</div>

            <div class="column items-center q-mt-xl">
              <q-card bordered flat style="border-radius: 10px;width: 210px;">
                <q-card-section>
                  <img :src="inviteImage" alt="">
                </q-card-section>
              </q-card>
              <div class="q-mt-md" style="width: 310px">
                <q-input outlined dense v-model="inviteUrl" readonly>
                  <template v-slot:append>
                    <q-icon name="content_copy" class="cursor-pointer" size="xs" @click="copyToClipboardFunc"></q-icon>
                  </template>
                </q-input>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { onMounted, reactive, toRefs } from 'vue';
import { copyToClipboard } from 'quasar';
import { inviteInfoAPI } from 'src/apis/user';
import { NotifyPositive } from 'src/utils/notify';
import { InitStore } from 'stores/init';
import QRCode from 'qrcode-svg-ts';
import { useI18n } from 'vue-i18n';

export default {
  name: 'ShareIndex',
  setup() {
    const $initStore = InitStore()
    const { t } = useI18n()
    const state = reactive({
      config: $initStore.config,
      inviteCode: '',
      inviteUrl: '',
      inviteImage: '',
    });

    onMounted(() => {
      inviteInfoAPI().then((res: any) => {
        state.inviteCode = res.code
        state.inviteUrl = location.origin + `/register?code=${res.code}`
        const qrCode = new QRCode({
          content: state.inviteUrl,
          padding: 0,
          width: 175,
          height: 175,
          color: '#000000',
          background: '#ffffff',
          ecl: 'M',
        });
        state.inviteImage = qrCode.toDataURL()
      })
    })

    // 复制方法
    const copyToClipboardFunc = () => {
      copyToClipboard(state.inviteUrl)
        .then(() => {
          NotifyPositive(t('copySuccess'))
        })
    };
    return {
      ...toRefs(state),
      copyToClipboardFunc
    }
  }
};
</script>
<style scoped></style>
