<template>
  <div>
    <div>Market</div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';

export default {
  name: 'MarketIndex',
  setup() {
    const state = reactive({});

    return {
      ...toRefs(state),
    };
  },
};
</script>
