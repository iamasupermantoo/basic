<template>
  <div>
    article
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue'
import { imageSrc } from 'src/utils'

export default {
  name: 'ArticleIndex',
  setup() {
    const state = reactive({})

    return {
      imageSrc,
      ...toRefs(state),
    }
  }
}
</script>

<style lang="scss" scoped></style>
