<template>
  <div style="padding: 48px 100px;">
    <!-- 头像 -->
    <div class="rounded-borders q-pa-lg text-white" style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);">
      <div class="row items-center">
        <div class="row items-center">
          <q-avatar class="bg-white" size="60px">
            <q-img no-spinner :src="imageSrc(currentTeamInfo.currentInfo.avatar)"></q-img>
          </q-avatar>
          <div class="q-ml-md text-white">
            <div class="text-body1 text-weight-bold">
              {{ currentTeamInfo.currentInfo.username }}
              <span class="text-caption text-grey-1">[ID:{{ currentTeamInfo.currentInfo.id }}]</span>
              <q-icon name="keyboard_double_arrow_down" class="q-ml-xs"></q-icon>
              <span class="text-caption">{{ currentTeamInfo.currentInfo.depth }}</span>
            </div>
            <div class="row no-wrap items-center q-mt-sm">
              <div class="text-white text-subtitle2">
                {{ $t('teamEarnings') }}: <span class="text-body1">+{{ currentTeamInfo.currentInfo.earnings }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="col q-ml-lg">
          <q-scroll-area style="height: 50px; width: 100%" :thumb-style="{ display: 'none' }">
            <div class="row no-wrap items-center">
              <div style="width: 110px" class="row justify-end items-center">
                <q-separator color="white" vertical inset />
                <div class="col">
                  <div class="text-h5 text-center text-weight-medium">{{ currentTeamInfo.todayNums }}</div>
                  <div class="text-center text-caption">{{ $t('todayNums') }}</div>
                </div>
              </div>

              <div style="width: 110px" class="row justify-end items-center">
                <q-separator color="white" vertical inset />
                <div class="col">
                  <div class="text-h5 text-center text-weight-medium">{{ currentTeamInfo.todayAmount }}</div>
                  <div class="text-center text-caption">{{ $t('todayAmount') }}</div>
                </div>
              </div>

              <div style="width: 110px" class="row justify-end items-center">
                <q-separator color="white" vertical inset />
                <div class="col">
                  <div class="text-h5 text-center text-weight-medium">{{ currentTeamInfo.todayEarnings }}</div>
                  <div class="text-center text-caption">{{ $t('todayEarnings') }}</div>
                </div>
              </div>

              <div style="width: 110px" class="row justify-end items-center">
                <q-separator color="white" vertical inset />
                <div class="col">
                  <div class="text-h5 text-center text-weight-medium">{{ currentTeamInfo.inviteNums }}</div>
                  <div class="text-center text-caption">{{ $t('inviteNums') }}</div>
                </div>
              </div>

              <div style="width: 110px" class="row justify-end items-center">
                <q-separator color="white" vertical inset />
                <div class="col">
                  <div class="text-h5 text-center text-weight-medium">{{ currentTeamInfo.buyAmount }}</div>
                  <div class="text-center text-caption">{{ $t('buyAmount') }}</div>
                </div>
              </div>

              <div style="width: 110px" class="row justify-end items-center">
                <q-separator color="white" vertical inset />
                <div class="col">
                  <div class="text-h5 text-center text-weight-medium">{{ currentTeamInfo.teamEarnings }}</div>
                  <div class="text-center text-caption">{{ $t('teamEarnings') }}</div>
                </div>
              </div>
            </div>
          </q-scroll-area>
        </div>
      </div>
    </div>

    <div class="q-mt-xl">
      <div v-for="(children, childrenIndex) in currentTeamInfo.children" :key="childrenIndex">
        <div class="row justify-between items-center q-my-lg q-px-md">
          <div class="col row items-center">
            <q-avatar size="40px">
              <q-img no-spinner :src="imageSrc(children.avatar)"></q-img>
            </q-avatar>
            <div class="q-ml-sm">
              <div class="text-body2 text-bold">
                {{ children.username }}
                <span class="text-caption text-grey-7">(ID:{{ children.id }})</span>
              </div>
              <div class="text-grey text-caption">{{ date.formatDate(children.createdAt * 1000, 'YYYY/MM/DD HH:mm:SS') }}
              </div>
            </div>
          </div>
          <div class="col">
            <div class="text-center">{{ children.name }}</div>
          </div>
          <div class="col">
            <div class="text-primary text-h6 text-right q-mr-lg">
              +{{ children.money }}
            </div>
          </div>
        </div>
        <q-separator />
      </div>
      <div v-if="currentTeamInfo.children == null || currentTeamInfo.children.length <= 0">
        <div class="text-center q-my-lg text-body1 text-grey">
          {{ $t('noData') }}
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { imageSrc } from 'src/utils';
import { teamDetailsAPI } from 'src/apis/user';
import { UserStore } from 'src/stores/user';
import { useRoute } from 'vue-router';
import { date } from 'quasar';

export default defineComponent({
  name: 'TeamEarnings',
  setup() {
    const $userStore = UserStore();
    const $route = useRoute();

    const state = reactive({
      currentUserId: $route.query.id ?? 0,
      currentTeamInfo: {
        currentInfo: {} as any
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
<style lang="scss" scoped></style>
