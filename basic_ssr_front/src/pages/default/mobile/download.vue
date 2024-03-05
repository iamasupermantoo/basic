<template>
  <div class="column bg-white window-height">
    <div class="col">
      <div class="column justify-between items-center">
        <div class="column justify-center shadow-1 q-mt-xl" style="width: 80px;height: 80px;border-radius: 15px;">
          <q-img no-spinner class="self-center" :src="imageSrc(config.logo)" width="56px" height="56px" />
        </div>
        <div class="text-h6 text-weight-bold q-mt-md">{{ config.name }}</div>
        <div class="q-mt-xl full-width">
          <div class="row no-wrap justify-center">
            <img style="width: 120px;" class="cursor-pointer q-mx-sm" src="/images/ios.svg" alt=""
              @click="downloadFunc(downloadInfo.ios)">
            <img style="width: 120px;" class="cursor-pointer q-mx-sm" src="/images/android.svg" alt=""
              @click="downloadFunc(downloadInfo.android)">
          </div>
        </div>
      </div>

    </div>
    <q-img no-spinner src="/images/download.png" width="240px" height="163px" class="self-center q-mb-xl" />
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue';
import { date } from 'quasar';
import { imageSrc } from 'src/utils';
import { useI18n } from 'vue-i18n';
import { InitStore } from 'src/stores/init';
import { downloadInfoAPI } from 'src/apis';

export default {
  name: 'DownloadIndex',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('download'),
    })

    const $initStore = InitStore();

    const state = reactive({
      config: $initStore.config,
      downloadInfo: {} as any,
    });

    onMounted(() => {
      downloadInfoAPI().then((res: any) => {
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
      date,
      ...toRefs(state),
      downloadFunc,
    }
  }
};
</script>

<style lang="scss" scoped></style>
