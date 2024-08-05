<template>
  <div>
    <div>Home</div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';

export default {
  name: 'HomeIndex',
  setup() {
    const state = reactive({});

    return {
      ...toRefs(state),
    };
  },
};
</script>
