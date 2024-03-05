<template>
  <div class="q-pa-lg">
    <!-- 文章详情页面 -->
    <div class="text-h6 text-primary text-weight-bold q-mb-sm text-center">{{ name }}</div>
    <div class="content" v-html="content"></div>
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
