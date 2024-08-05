<template>
  <div>
    <div class="row justify-between q-mb-xs">
      <div class="text-body2 text-grey">{{ label }}</div>
      <div class="row q-gutter-xs">
        <div>
          <q-btn
            dense
            flat
            outline
            class="bg-secondary text-white q-pa-xs"
            size="xs"
            @click="addChildrenFunc()"
            label="新增"
          ></q-btn>
        </div>
        <div>
          <q-btn
            dense
            flat
            outline
            class="bg-red text-white q-pa-xs"
            size="xs"
            @click="delChildrenFunc()"
            label="删除"
          ></q-btn>
        </div>
      </div>
    </div>

    <div v-for="(item, itemIndex) in value" :key="itemIndex">
      <div class="text-caption text-grey">
        #<span style="margin-left: 1px">{{ itemIndex }}</span>
      </div>
      <div v-for="(inputRow, inputRowIndex) in options" :key="inputRowIndex">
        <div class="row items-center q-gutter-xs">
          <div
            v-for="(inputCol, inputColIndex) in inputRow"
            :key="inputColIndex"
            class="col"
          >
            <InputComponent
              :type="inputCol.type"
              v-model="value[itemIndex][inputCol.field]"
              :label="inputCol.label"
              :options="inputCol.data"
              :readonly="inputCol.readonly"
              @refresh="componentsRefreshValue"
            ></InputComponent>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="value.length == 0"
      class="text-caption text-center text-grey q-my-md"
    >
      暂无数据
    </div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import InputComponent from 'src/components/input.vue';
import { InputTypeList } from 'src/utils/define';

export default {
  name: 'InputChildrenComponent',
  props: {
    type: { type: Number, default: InputTypeList.Text },
    label: { type: String, default: '' },
    modelValue: { type: null, default: [] as any },
    options: { type: null, default: [] },
    readonly: { type: Boolean, default: false },
  },
  emits: ['refresh'],
  components: { InputComponent },
  setup(props: any, context: any) {
    const state = reactive({
      value: props.modelValue,
    });

    if (state.value == '' || state.value == null) {
      state.value = [];
    }

    //  组件刷新值
    const componentsRefreshValue = (value: any) => {
      context.emit('refresh', value);
    };

    const addChildrenFunc = () => {
      const tmpValue = {} as any;
      for (let i = 0; i < props.options.length; i++) {
        for (let j = 0; j < props.options[i].length; j++) {
          tmpValue[props.options[i][j].field] = '';
        }
      }

      state.value.push(tmpValue);
      context.emit('refresh', state.value);
    };

    const delChildrenFunc = () => {
      state.value.pop();
      context.emit('refresh', state.value);
    };

    return {
      componentsRefreshValue,
      addChildrenFunc,
      delChildrenFunc,
      ...toRefs(state),
    };
  },
};
</script>
