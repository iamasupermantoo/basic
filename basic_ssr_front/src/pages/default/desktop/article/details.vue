<template>
  <div style="padding: 80px 108px">
    <div class="text-h5 text-weight-bold text-center q-mb-lg">{{ name }}</div>
    <div v-html="content"></div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { articleInfoAPI } from 'src/apis'

export default {
  name: 'ArticleDetails',
  setup() {
    const $route = useRoute();

    const state = reactive({
      content: '',
      name: '',
    })

    onMounted(() => {
      articleInfoAPI({ id: Number($route.query.id) }).then((res: any) => {
        state.content = res.content
        state.name = res.name
      })
    })
    return {
      ...toRefs(state),
    }
  }
}
</script>

<style lang="scss" scoped></style>
