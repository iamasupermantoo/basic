<template>
  <div>
    <q-dialog
      v-model="dialog"
      :full-width="config.fullWidth"
      :full-height="config.fullHeight"
    >
      <q-card
        :style="sizingList.get(config.sizing)"
        :class="config.fullHeight ? 'column full-height' : ''"
      >
        <q-card-section v-if="config.title != ''">
          <div class="text-h6">{{ config.title }}</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <div class="text-body1" v-if="config.content != ''">
            {{ config.content }}
          </div>

          <div class="q-mt-md">
            <div
              v-for="(input, inputIndex) in config.inputList"
              :key="inputIndex"
            >
              <InputLayoutsComponent
                v-model="config.params[input.field]"
                :label="input.label"
                :type="input.type"
                :options="input.data"
              ></InputLayoutsComponent>
            </div>
          </div>
        </q-card-section>

        <q-card-actions
          align="right"
          v-if="config.buttons.cannel || config.buttons.confirm"
        >
          <q-btn
            flat
            :color="
              config.buttons.cannel.color == ''
                ? 'grey'
                : config.buttons.cannel.color
            "
            :label="config.buttons.cannel.label"
            :size="config.buttons.cannel.size"
            v-close-popup
          />
          <q-btn
            flat
            :color="
              config.buttons.confirm.color == ''
                ? 'primary'
                : config.buttons.confirm.color
            "
            :label="config.buttons.confirm.label"
            :size="config.buttons.confirm.size"
            @click="submitFunc()"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { api } from 'src/boot/axios';
import InputLayoutsComponent from 'src/components/inputLayouts.vue';

export default {
  name: 'DialogComponent',
  components: { InputLayoutsComponent },
  emits: ['done'],
  setup(props: any, context: any) {
    const state = reactive({
      dialog: props.show,
      config: {
        id: 'none',
        url: '',
        title: 'Title', //  标题
        content: '', //  提示内容
        sizing: 'medium',
        fullWidth: false,
        fullHeight: false,
        params: {} as any, //  参数
        inputList: [] as any, //  input组
        buttons: [] as any,
      },

      //  大小 [small 小型] [medium 中型] [fullWidth 满屏宽度] [fullHeight 满屏高度]
      sizingList: new Map([
        ['small', 'width: 300px'],
        ['medium', 'width: 700px; max-width: 80vw'],
      ]),
    });

    // 设置配置
    const setConfig = (config: any) => {
      state.config = config;
    };

    //  打开弹窗
    const open = () => {
      state.dialog = true;
    };

    //  按钮提交
    const submitFunc = () => {
      //  需要请求数据
      api.post(state.config.url, state.config.params).then(() => {
        state.dialog = false;
        context.emit('done', state.config.id);
      });
    };

    return {
      submitFunc,
      open,
      setConfig,
      ...toRefs(state),
    };
  },
};
</script>
