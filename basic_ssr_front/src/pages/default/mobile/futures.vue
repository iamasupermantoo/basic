<template>
  <div>
    <div>Futures</div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';

export default {
  name: 'FuturesIndex',
  setup() {
    const state = reactive({});

    return {
      ...toRefs(state),
    };
  },
};
</script>
