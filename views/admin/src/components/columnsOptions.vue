<template>
  <div></div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { ContentTypeList } from 'src/utils/define';

export default {
  name: 'ColumnsOptionsComponent',
  props: {
    type: { type: Number, default: ContentTypeList.Text },
    modelValue: { type: undefined, default: '' },
    options: { type: null, default: [] },
  },
  emits: ['done'],
  setup(props: any) {
    const state = reactive({
      value: props.modelValue,
    });

    return {
      ...toRefs(state),
    };
  },
};
</script>
