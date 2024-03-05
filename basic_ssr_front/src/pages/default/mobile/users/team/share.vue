<template>
  <div class="column window-height" style=" background: linear-gradient(to right,  #14BF71,#82D880);">
    <div class="col column">
      <div class="rounded-borders bg-white column items-center q-py-md" style="width: 90%;margin: 36px auto;">

        <q-img no-spinner :src="imageSrc('')" class="q-mt-lg" width="60px" height="60px" />

        <div class="text-color-3 text-h6 q-mt-md">{{ $t('inviteFriends') }}</div>

        <img :src="inviteImage" alt="">

        <div class="text-grey-7 text-weight-medium" style="margin: 20px 0 10px 0;">Copy invitation link</div>

        <div class="text-weight-regular text-black ellipsis  q-pa-sm text-body1"
          style="background-color: #F4F5F8;width: 238px;border-radius: 10px;">
          {{ inviteUrl }}
        </div>

        <q-btn unelevated rounded style="margin: 30px 0 20px 0;width: 230px;height: 40px;" color="primary"
          :label="$t('copy')" no-caps @click="copyToClipboardFunc(inviteUrl)" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue';
import { copyToClipboard } from 'quasar';
import { useI18n } from 'vue-i18n';
import { inviteInfoAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { NotifyPositive } from 'src/utils/notify';
import QRCode from 'qrcode-svg-ts';

export default {
  name: 'TeamShareIndex',
  setup(props: any, context: any) {
    const { t } = useI18n()
    context.emit('update', {
      title: t('inviteFriends'),
    })

    const state = reactive({
      inviteUrl: '',
      inviteImage: '',
    });

    onMounted(() => {
      inviteInfoAPI().then((res: any) => {
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
    const copyToClipboardFunc = (str: string) => {
      copyToClipboard(str)
        .then(() => {
          NotifyPositive(t('copySuccess'));
        })
    };
    return {
      imageSrc,
      ...toRefs(state),
      copyToClipboardFunc,
    }
  }
};
</script>
