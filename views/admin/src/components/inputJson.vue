<template>
  <div>
    <div class="text-body2 text-grey q-mb-xs">{{ label }}</div>
    <div v-for="(inputRow, inputRowIndex) in options" :key="inputRowIndex">
      <div class="row items-center q-gutter-xs">
        <div
          v-for="(inputCol, inputColIndex) in inputRow"
          :key="inputColIndex"
          class="col"
        >
          <InputComponent
            :type="inputCol.type"
            :field="inputCol.field"
            v-model="value[inputCol.field]"
            :label="inputCol.label"
            :options="inputCol.data"
            :readonly="inputCol.readonly"
            @refresh="componentsRefreshValue"
          ></InputComponent>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import InputComponent from 'src/components/input.vue';
import { InputTypeList } from 'src/utils/define';

export default {
  name: 'InputJsonComponent',
  props: {
    type: { type: Number, default: InputTypeList.Text },
    label: { type: String, default: '' },
    modelValue: { type: undefined, default: '' },
    options: { type: null, default: [] },
    readonly: { type: Boolean, default: false },
  },
  emits: ['refresh'],
  components: { InputComponent },
  setup(props: any, context: any) {
    const state = reactive({
      value: props.modelValue,
    });

    //  组件刷新值
    const componentsRefreshValue = (value: any, field: string) => {
      state.value[field] = value;
      context.emit('refresh', state.value);
    };

    return {
      componentsRefreshValue,
      ...toRefs(state),
    };
  },
};
</script>
