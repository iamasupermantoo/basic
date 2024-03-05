<template>
  <div class="column full-height">
    <div class="col page_bg q-pa-md ">
      <div class="rounded-borders q-px-md q-py-lg row"
        style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);">
        <q-avatar size="50px" class="bg-white">
          <q-img no-spinner :src="imageSrc(currentTeamInfo.currentInfo.avatar)"></q-img>
        </q-avatar>
        <div class="q-ml-md text-subtitle1">
          <div class="text-white text-weight-medium">
            {{ currentTeamInfo.currentInfo.username }}
            <span class="text-caption text-grey-1">[ID:{{ currentTeamInfo.currentInfo.id }}]</span>
            <q-icon name="keyboard_double_arrow_down" class="q-ml-xs"></q-icon>
            <span class="text-caption">{{ currentTeamInfo.currentInfo.depth }}</span>
          </div>
          <div class="text-white text-weight-medium">
            {{ $t('teamEarnings') }}: <span class="text-body1">+{{ currentTeamInfo.currentInfo.earnings }}</span>
          </div>
        </div>
      </div>

      <div class="rounded-borders bg-white q-px-md q-py-lg  q-mt-md">
        <div class="text-black text-weight-bold">Beneficial Data</div>
        <div class="row ">
          <div class="col-4 q-pt-md">
            <div class="text-weight-bold text-h6 text-center">{{ currentTeamInfo.todayNums }}</div>
            <div class="text-grey-8  text-subtitle2  text-center text-weight-regular ellipsis">
              {{ $t('todayNums') }}</div>
          </div>
          <div class="col-4 q-pt-md">
            <div class="text-weight-bold text-h6 text-center">{{ currentTeamInfo.todayAmount }}</div>
            <div class="text-grey-8  text-subtitle2  text-center text-weight-regular ellipsis">
              {{ $t('todayAmount') }}</div>
          </div>
          <div class="col-4 q-pt-md">
            <div class="text-weight-bold text-h6 text-center">{{ currentTeamInfo.todayEarnings }}</div>
            <div class="text-grey-8  text-subtitle2  text-center text-weight-regular ellipsis">
              {{ $t('todayEarnings') }}</div>
          </div>
          <div class="col-4 q-pt-md">
            <div class="text-weight-bold text-h6 text-center">{{ currentTeamInfo.inviteNums }}</div>
            <div class="text-grey-8  text-subtitle2  text-center text-weight-regular ellipsis">
              {{ $t('inviteNums') }}</div>
          </div>
          <div class="col-4 q-pt-md">
            <div class="text-weight-bold text-h6 text-center">{{ currentTeamInfo.buyAmount }}</div>
            <div class="text-grey-8  text-subtitle2  text-center text-weight-regular ellipsis">
              {{ $t('buyAmount') }}</div>
          </div>
          <div class="col-4 q-pt-md">
            <div class="text-weight-bold text-h6 text-center">{{ currentTeamInfo.teamEarnings }}</div>
            <div class="text-grey-8  text-subtitle2  text-center text-weight-regular ellipsis">
              {{ $t('teamEarnings') }}</div>
          </div>
        </div>
      </div>

      <div class="q-py-md row">
        <div class="text-subtitle1 text-weight-bolder column col-3">
          <div class="text-center">{{ $t('transactions') }}</div>
          <q-separator style="height: 4px;width: 20px;border-radius: 2px;" class="bg-primary self-center" />
        </div>
      </div>

      <div v-for="(children, childrenIndex) in currentTeamInfo.children" :key="childrenIndex"
        class="row justify-between items-center bg-white rounded-borders q-pa-md q-mb-md">

        <div class="row">
          <q-avatar size="40px">
            <q-img no-spinner :src="imageSrc(children.avatar)"></q-img>
          </q-avatar>
          <div class="q-ml-sm">
            <div class="text-subtitle2 text-weight-medium">
              {{ children.username }}
              <span class="text-caption text-grey-7">(ID:{{ children.id }})</span>
            </div>
            <div class="text-grey-8 text-caption text-weight-regular text-weight-regular">{{
              date.formatDate(children.createdAt
                *
                1000, 'YYYY/MM/DD HH:mm:ss') }}</div>
          </div>
        </div>
        <div>
          <div class="text-primary text-right text-subtitle1 text-weight-medium">+{{ children.money }}</div>
          <div class="text-caption text-right">{{ children.name }}</div>
        </div>
      </div>

      <div v-if="currentTeamInfo.children && currentTeamInfo.children.length <= 0" class="text-grey text-center q-py-lg">
        {{ $t('noData') }}
      </div>


    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { imageSrc } from 'src/utils';
import { teamDetailsAPI } from 'src/apis/user';
import { UserStore } from 'src/stores/user';
import { useI18n } from 'vue-i18n';
import { useRoute } from 'vue-router';
import { date } from 'quasar';

export default defineComponent({
  name: 'TeamEarnings',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('teamEarnings'),
    })

    const $userStore = UserStore();
    const $route = useRoute();

    const state = reactive({
      currentUserId: $route.query.id ?? 0,
      currentTeamInfo: {
        currentInfo: {} as any,
        children: [],
      } as any,
    });


    onMounted(() => {
      if (state.currentUserId == 0) {
        state.currentUserId = $userStore.userInfo.id
      }
      teamDetailsAPI({ id: Number(state.currentUserId) }).then((res: any) => {
        state.currentTeamInfo = res
      })
    })

    return {
      imageSrc,
      date,
      ...toRefs(state),
    }
  }
});
</script>


