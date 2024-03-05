<template>
  <div>
    <div class="row items-center justify-center"
      :style="{ background: 'url(/images/label-bg.png)', height: '200px', backgroundSize: 'cover', backgroundRepeat: 'no-repeat' }">
      <div class="text-white text-h4">{{ $t('download') }}</div>
    </div>

    <div style="padding: 80px 0 260px">
      <div class="row justify-center items-center">
        <div class="col text-right">
          <q-img no-spinner class="q-mr-xl" width="300px" src="/images/download-bg.svg"></q-img>
        </div>
        <div class="col-1"></div>
        <div class="col">
          <div class="text-h3 text-bold">{{ config.name }}</div>
          <div class="text-h6 text-grey q-mt-sm">{{ $t('downloadSmall') }}</div>

          <div class="row no-wrap" style="margin-top: 60px">
            <div>
              <div class="text-body1 text-bold q-mb-xs">IOS</div>
              <img style="width: 150px;" class="cursor-pointer" src="/images/ios.svg" alt=""
                @click="downloadFunc(downloadInfo.ios)">
              <div class="text-body1 text-bold q-mt-lg q-mb-xs">Android</div>
              <img style="width: 150px;" class="cursor-pointer" src="/images/android.svg" alt=""
                @click="downloadFunc(downloadInfo.android)">
            </div>
            <div class="q-ml-xl row items-end">
              <div class="q-pa-lg rounded-borders" style="border: 1px solid rgba(0,20,42,0.12)">
                <img :src="downloadImage" alt="">
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
import { imageSrc } from 'src/utils';
import { downloadInfoAPI } from 'src/apis';
import { InitStore } from 'src/stores/init';
import QRCode from 'qrcode-svg-ts';

export default {
  name: 'DownloadIndex',
  setup() {
    const $initStore = InitStore();

    const state = reactive({
      config: $initStore.config,
      downloadInfo: { android: '', ios: '' } as any,
      downloadImage: '',
    });

    onMounted(() => {
      downloadInfoAPI().then((res: any) => {
        const qrCode = new QRCode({
          content: window.location.href,
          width: 130,
          height: 130,
          color: '#000000',
          background: '#ffffff',
          ecl: 'M',
        });
        state.downloadImage = qrCode.toDataURL()
        state.downloadInfo = res
      })
    })

    // 下载方法
    const downloadFunc = (url: string) => {
      if (url == '') {
        return
      }
      window.location.href = imageSrc(url)
    }

    return {
      imageSrc,
      downloadFunc,
      ...toRefs(state),
    }
  }
};
</script>
<style lang="scss" scoped></style>
