<template>
  <div class="row justify-center full-width">
    <div class="col">
      <!-- 背景 -->
      <div class="row items-center justify-center" :style="{background: 'url(/images/label-bg.png)', height: '200px', backgroundSize: 'cover', backgroundRepeat: 'no-repeat'}">
        <div class="text-white text-h4">{{ $t('helpers') }}</div>
      </div>

      <div style="padding: 20px 10%">
        <div>
          <q-card flat bordered v-for="( helper, helperIndex ) in  articleList " :key="helperIndex" class="q-mb-md">
            <q-card-section>
              <div class="text-h6">{{helper.name}}</div>
            </q-card-section>
            <q-card-section>
              <div v-html="helper.content"></div>
            </q-card-section>
            <q-card-actions align="right">
              <div class="text-caption text-grey">{{date.formatDate(helper.createdAt * 1000, 'YYYY/MM/DD HH:mm:ss')}}</div>
            </q-card-actions>
          </q-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue'
import { helpersInfoAPI } from 'src/apis'
import { imageSrc } from 'src/utils'
import {date} from 'quasar'

export default {
  name: 'HelpersCenter',
  setup() {
    const state = reactive({
      // 帮助文章
      articleList: [] as any,

      // 帮助列表
      socialList: [] as any,
    })

    onMounted(() => {
      helpersInfoAPI().then((res: any) => {
        state.articleList = res.articleList
        state.socialList = res.socialList
      })
    })

    return {
      imageSrc,
      date,
      ...toRefs(state),
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
