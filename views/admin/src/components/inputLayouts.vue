<template>
  <div>
    <!-- Json对象 -->
    <JsonInputComponent
      v-if="type == InputTypeList.InputJson"
      :type="type"
      v-model="value"
      :label="label"
      :options="options"
      :readonly="readonly"
      @refresh="componentsRefreshValue"
    ></JsonInputComponent>

    <!-- 子级数组对象 -->
    <ChildrenInputComponent
      v-else-if="type == InputTypeList.InputChildren"
      :type="type"
      v-model="value"
      :label="label"
      :options="options"
      :readonly="readonly"
      @refresh="componentsRefreshValue"
    ></ChildrenInputComponent>

    <!-- 正常input对象 -->
    <InputComponent
      v-else
      :type="type"
      v-model="value"
      :label="label"
      :options="options"
      :readonly="readonly"
      @refresh="componentsRefreshValue"
    ></InputComponent>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { date } from 'quasar';
import { InputTypeList } from 'src/utils/define';
import InputComponent from 'src/components/input.vue';
import JsonInputComponent from 'src/components/inputJson.vue';
import ChildrenInputComponent from 'src/components/inputChildren.vue';

export default {
  name: 'InputLayoutsComponent',
  props: {
    type: { type: Number, default: InputTypeList.Text },
    label: { type: String, default: '' },
    modelValue: { type: undefined, default: '' },
    options: { type: null, default: [] },
    readonly: { type: Boolean, default: false },
  },
  components: { InputComponent, JsonInputComponent, ChildrenInputComponent },
  setup(props: any, context: any) {
    const state = reactive({
      value: props.modelValue,
    });

    //  默认值设置 - 时间选择
    if (props.type == InputTypeList.DatePicker) {
      state.value = date.formatDate(state.value * 1000, 'YYYY/MM/DD');
      context.emit('update:modelValue', state.value);
    }

    //  组件刷新值
    const componentsRefreshValue = (value: any) => {
      context.emit('update:modelValue', value);
    };

    return {
      InputTypeList,
      componentsRefreshValue,
      ...toRefs(state),
    };
  },
};
</script>
