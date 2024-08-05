<template>
  <div class="column full-height">
    <div class="col page_bg column ">
      <div class="rounded-borders q-px-md q-py-lg row justify-between q-ma-md "
        style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);">
        <div class="row">
          <q-avatar size="50px" class="bg-white">
            <q-img no-spinner :src="imageSrc(currentTeamInfo.avatar)"></q-img>
          </q-avatar>
          <div class="q-ml-md text-subtitle1">
            <div class="text-white">
              <div>
                <span class="text-weight-medium">{{ currentTeamInfo.username }}</span>
                <span class="text-caption text-grey-1">[ID:{{ currentTeamInfo.id }}]</span>
                <q-icon name="keyboard_double_arrow_down" class="q-ml-xs"></q-icon>
                <span class="text-caption">{{ currentTeamInfo.depth }}</span>
              </div>
            </div>
            <div class="text-white text-weight-medium">
              {{ $t('teamEarnings') }}: <span class="text-body1">+{{ currentTeamInfo.earnings }}</span>
            </div>
          </div>
        </div>
        <div @click="$router.push({ name: 'TeamEarnings', query: { id: currentTeamInfo.id } })" class="text-white">{{
          $t('desc') }} <q-icon name="chevron_right" size="16px" /></div>
      </div>
      <div class="bg-white q-pa-md" v-if="currentTeamInfo.children.length > 0">
        <div v-for="(children, childrenIndex) in currentTeamInfo.children" :key="childrenIndex">
          <div @click="$router.push({ name: 'TeamIndex', query: { id: children.id } })"
            class="row justify-between bg-white q-py-md">
            <div class="row">
              <q-avatar size="32px">
                <q-img no-spinner :src="children.avatar ? imageSrc(children.avatar) : imageSrc('')"></q-img>
              </q-avatar>
              <div class="q-ml-md">
                <div class="text-color-3 text-subtitle2 text-weight-medium">
                  {{ children.username }}
                  (ID:{{ children.id }})
                </div>
                <div class="text-grey-6 text-caption text-weight-regular text-weight-regular">{{
                  formatDate(children.createdAt) }}</div>
              </div>
            </div>
            <div class="row justify-end">
              <div class="text-primary self-center text-subtitle1 text-weight-medium">+{{ children.earnings }}</div>
              <q-icon class="self-center" name="chevron_right" size="22px" style="color: #999999;" />
            </div>
          </div>
          <q-separator style="height: 1px;background: #F4F5FD;" />
        </div>
      </div>
      <div v-else class="q-py-lg text-center text-grey-6">
        {{ $t('noData') }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { onMounted, reactive, toRefs } from 'vue';
import { useI18n } from 'vue-i18n';
import { UserStore } from 'src/stores/user';
import { imageSrc, formatDate } from 'src/utils';
import { teamIndexAPI } from 'src/apis/user';
import { useRouter } from 'vue-router';


export default {
  name: 'TeamIndex',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('myTeam'),
    })

    const $userStore = UserStore();
    const $router = useRouter();

    const state = reactive({
      currentUserId: $router.currentRoute.value.query.id ?? 0,
      currentTeamInfo: {
        children: [],
      } as any,
    });

    onMounted(() => {
      if (state.currentUserId == 0) {
        state.currentUserId = $userStore.userInfo.id
      }
      teamIndexAPI({ id: Number(state.currentUserId) }).then((res: any) => {
        state.currentTeamInfo = res
      })
    })

    return {
      imageSrc,
      formatDate,
      ...toRefs(state),
    }
  }
};
</script>
