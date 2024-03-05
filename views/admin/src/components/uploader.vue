<template>
  <div>
    <q-uploader
      :url="baseURL + '/upload'"
      field-name="file"
      ref="uploaderRef"
      auto-upload
      style="box-shadow: none; width: auto"
      :multiple="multiple"
      :headers="[{ name: 'Authorization', value: 'Bearer ' + userToken }]"
      @uploaded="uploadedFunc"
      @start="uploaderStartFunc"
      @finish="uploaderFinishFunc"
      @failed="uploadFailedFunc"
    >
      <template v-slot:header="scope">
        <div class="q-my-xs q-mx-sm">
          <div class="row justify-between items-center">
            <div>{{ label }}</div>
            <div>
              <q-btn
                v-if="scope.canAddFiles"
                type="a"
                icon="add_box"
                @click="scope.pickFiles"
                round
                dense
                flat
              >
                <q-uploader-add-trigger />
                <q-tooltip>Pick Files</q-tooltip>
              </q-btn>
            </div>
          </div>
        </div>
      </template>

      <template v-slot:list="scope">
        <q-list separator v-if="multiple">
          <q-item v-for="(image, imageIndex) in modelValue" :key="imageIndex">
            <q-item-section thumbnail>
              <img :src="imageSrc(image)" alt="" class="q-ml-sm" />
            </q-item-section>
            <q-item-section></q-item-section>
            <q-item-section side>
              <q-btn
                class="gt-xs"
                size="12px"
                flat
                dense
                round
                icon="delete"
                @click="deleteUploadedEventFunc(imageIndex)"
              />
            </q-item-section>
          </q-item>
        </q-list>

        <div
          @click="scope.pickFiles"
          class="row justify-center no-padding"
          v-else
        >
          <q-uploader-add-trigger />
          <img
            :src="imageSrc(modelValue)"
            alt=""
            v-if="modelValue && modelValue != ''"
          />
        </div>
      </template>
    </q-uploader>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, ref } from 'vue';
import { Loading, QSpinnerBars } from 'quasar';
import { useInitStore } from 'src/stores/init';
import { imageSrc } from 'src/utils';
import { NotifyNegative } from 'src/utils/notify';

export default {
  name: 'UploaderComponent',
  props: {
    label: { type: String, default: '' },
    modelValue: { type: undefined, default: '' },
    multiple: { type: Boolean, default: false },
  },
  emits: ['uploaded'],
  setup(props: any, context: any) {
    const $initStore = useInitStore();
    const uploaderRef = ref(null) as any;
    const state = reactive({
      modelValue: props.modelValue,
      baseURL: process.env.baseURL,
      userToken: $initStore.userToken,
    });

    // 多图初始化
    if (props.multiple && props.modelValue == '') {
      state.modelValue = []
    }

    // 上传完成回调方法
    const uploadedFunc = (info: any) => {
      const imagePath = JSON.parse(info.xhr.response).data[0];
      props.multiple
        ? state.modelValue.push(imagePath)
        : (state.modelValue = imagePath);

      context.emit('uploaded', state.modelValue);
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

    // 删除上传的图片
    const deleteUploadedEventFunc = (index: any) => {
      state.modelValue.splice(index, 1);
      context.emit('uploaded', state.modelValue);
    };

    // 上传失败方法
    const uploadFailedFunc = () => {
      NotifyNegative('文件上传失败, 请检查文件是否符合上传...');
      Loading.hide();
    };

    // 激活选择文件
    const pickFilesFunc = () => {
      uploaderRef.value.pickFiles();
    };

    return {
      uploaderRef,
      imageSrc,
      uploaderStartFunc,
      uploaderFinishFunc,
      pickFilesFunc,
      uploadedFunc,
      uploadFailedFunc,
      deleteUploadedEventFunc,
      ...toRefs(state),
    };
  },
};
</script>
<style scoped></style>
