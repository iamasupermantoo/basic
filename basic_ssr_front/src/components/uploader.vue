<template>
  <div>
    <q-uploader flat auto-upload :url="baseURL + '/upload'" :field-name="name" :style="uploaderStyle" :accept="accept"
      style="width: 100%;background: transparent;" :multiple="false" @uploaded="uploadedFunc" @start="uploaderStartFunc"
      @finish="uploaderFinishFunc" @rejected="uploaderRejectedFunc"
      :headers="[{ name: 'Authorization', value: 'Bearer ' + userToken }]">
      <template v-slot:header></template>
      <template v-slot:list="scope">
        <div @click="scope.pickFiles">
          <div v-if="!$slots.default" class="text-center">
            <q-uploader-add-trigger />
            <img :src="imageSrc(respValue)" alt="">
          </div>
          <slot name="default"></slot>
        </div>
      </template>
    </q-uploader>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import { Loading, QSpinnerBars } from 'quasar';
import { InitStore } from 'src/stores/init';
import { NotifyNegative } from 'src/utils/notify';

export default {
  name: 'UploaderComponents',
  props: {
    // 上传名称
    name: { type: String, default: 'file' },

    // 图片路径
    respValue: { type: String, default: '' },

    // uploader 自定义样式
    uploaderStyle: { type: Object, default: null },

    // 上传文件的类型
    accept: { type: String, default: '.jpg,.png,.gif,image/*' },
  },
  setup(props: any, context: any) {
    const $initStore = InitStore()
    const state = reactive({
      // 用户token
      userToken: $initStore.userToken,

      // API路由
      baseURL: process.env.baseURL,

      // 返回数据
      respValue: props.respValue,
    });

    // 上传完成回调方法
    const uploadedFunc = (info: any) => {
      state.respValue = JSON.parse(info.xhr.response).data[0];
      context.emit('uploaded', state.respValue);
    };

    // 开始上传
    const uploaderStartFunc = () => {
      Loading.show({
        spinner: QSpinnerBars,
        spinnerColor: 'secondary',
        spinnerSize: 50,
        message: 'Some important process is in progress. Hang on...',
      });
    };

    // 上传完成
    const uploaderFinishFunc = () => {
      Loading.hide();
    };

    // 检查格式是否正确
    const uploaderRejectedFunc = () => {
      NotifyNegative(props.accept)
    };

    return {
      imageSrc,
      uploadedFunc,
      uploaderStartFunc,
      uploaderFinishFunc,
      uploaderRejectedFunc,
      ...toRefs(state),
    };
  },
};
</script>

<style scoped lang="scss">
// :deep(.q-uploader .q-uploader__list) {
//   padding: 0 !important;
//   min-height: 0 !important;
// }
</style>
