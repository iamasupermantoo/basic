<template>
  <div class="q-pb-xl">
    article
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue'

export default {
  name: 'ArticleIndex',
  setup() {
    const state = reactive({})
    return {
      ...toRefs(state),
    }
  }
}
</script>

<style lang="scss" scoped></style>
