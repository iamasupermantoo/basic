<template>
  <!-- 背景 -->
  <div class="bg absolute full-width" style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);
  height: 190px;
  border-bottom-left-radius: 15%;
  border-bottom-right-radius: 15%;"></div>

  <div class="q-pa-md">
    <div class="row justify-between">
      <div class="col-6 z-top">
        <div class="text-body1 text-white text-weight-bold q-mb-sm">
          {{ $t('helperYou') }}
        </div>
        <div style="background: rgba(255, 255, 255, 0.12);padding: 7px 16px;max-width: 131px;border-radius: 19px;"
          class="text-center text-white ellipsis">{{ $t('24hoursOnline') }}</div>
      </div>
      <div class="column justify-end">
        <q-img no-spinner style="margin-bottom: -15px;margin-right: 20px;" src="/images/helpers.png" width="62px"
          height="79px" />
      </div>
    </div>

    <div class="row q-col-gutter-md justify-center q-mt-xs">

      <div class="col-4 z-top" v-for="(social, socialIndex ) in socialList" :key="socialIndex">
        <div class="rounded-borders bg-white column items-center shadow-1" style="padding: 20px 0;">
          <q-img no-spinner :src="imageSrc(social.icon)" width="41px" height="41px" />
          <div class="text-body2 text-weight-medium q-mt-sm">{{ social.name }}</div>
        </div>
      </div>
    </div>
  </div>

  <!-- 帮助中心列表 -->
  <div class="bg-grey-1 q-pb-xl">
    <q-list class="q-ma-md bg-white text-black rounded-borders">
      <q-card flat>
        <q-card-section>
          <div class="text-h6 text-weight-medium">FAQ</div>
        </q-card-section>
      </q-card>
      <q-expansion-item v-for="( helper, helperIndex ) in articleList " :key="helperIndex" :label="helper.name">
        <q-card flat>
          <q-card-section>
            <div v-html="helper.content"></div>
          </q-card-section>
        </q-card>
      </q-expansion-item>
    </q-list>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue';
import { imageSrc } from 'src/utils';
import { helpersInfoAPI } from 'src/apis'
import { useI18n } from 'vue-i18n'

export default {
  name: 'helpCenter',
  emits: ['update'],
  setup(props: any, context: any) {
    const { t } = useI18n();

    const state = reactive({
      // 帮助文章
      articleList: [] as any,

      // 帮助列表
      socialList: [] as any,
    });

    context.emit('update', {
      title: t('helpers'),
    })

    onMounted(() => {
      helpersInfoAPI().then((res: any) => {
        state.articleList = res.articleList
        state.socialList = res.socialList.splice(0, 3)
      })
    })

    return {
      imageSrc,
      ...toRefs(state),
    }
  }
};
</script>

<style lang="scss" scoped>
::v-deep(.q-expansion-item) {
  border-bottom: 1px solid #F4F5FD;
}
</style>
