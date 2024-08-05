<template>
  <div style="padding: 48px 100px;">
    <!-- 总金额 -->
    <div class="row justify-between items-center q-pa-lg rounded-borders"
      style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);">
      <div class="row items-center">
        <q-avatar size="60px" class="bg-white">
          <q-img no-spinner :src="imageSrc(currentTeamInfo.avatar)"></q-img>
        </q-avatar>

        <div class="q-ml-md text-white">
          <div class="text-body1 text-weight-bold">
            {{ currentTeamInfo.username }}
            <span class="text-caption text-grey-1">[ID:{{ currentTeamInfo.id }}]</span>
            <q-icon name="keyboard_double_arrow_down" class="q-ml-xs"></q-icon>
            <span class="text-caption">{{ currentTeamInfo.depth }}</span>
          </div>
          <div class="row no-wrap items-center q-mt-sm">
            <div class="text-white text-subtitle2">
              {{ $t('teamEarnings') }}: <span class="text-body1">+{{ currentTeamInfo.earnings }}</span>
            </div>
          </div>
        </div>
      </div>
      <div>
        <div @click="$router.push({ name: 'TeamEarnings', query: { id: currentTeamInfo.id } })"
          class="cursor-pointer text-white text-body2">
          <span>{{ $t('desc') }}</span>
          <q-icon size="14px" class="text-white" name="arrow_forward_ios"></q-icon>
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
              <div class="text-caption text-grey">{{ $t('teamMember') }}</div>
              <div class="text-body2 text-bold">
                {{ children.username }}
                <span class="text-caption text-grey-7">(ID:{{ children.id }})</span>
              </div>
            </div>
          </div>
          <div class="col">
            <div class="text-caption text-grey">{{ $t('createdTime') }}</div>
            <div>
              {{ date.formatDate(children.createdAt * 1000, 'YYYY/MM/DD HH:mm:SS') }}
            </div>
          </div>
          <div class="col">
            <div class="text-primary text-h6 text-right q-mr-lg">
              +{{ children.earnings }}
            </div>
          </div>
          <div>
            <div @click="$router.push({ name: 'TeamIndex', query: { id: children.id } })"
              class="cursor-pointer text-grey text-body2">
              <span>{{ $t('views') }}</span>
              <q-icon size="14px" class="text-grey" name="arrow_forward_ios"></q-icon>
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
import { imageSrc } from 'src/utils'
import { teamIndexAPI } from 'src/apis/user';
import { UserStore } from 'src/stores/user';
import { date } from 'quasar';
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'TeamIndex',
  setup() {
    const $userStore = UserStore();
    const $router = useRouter()

    const state = reactive({
      currentUserId: $router.currentRoute.value.query.id ?? 0,
      currentTeamInfo: {} as any,
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
      date,
      ...toRefs(state),
    }
  }
});
</script>
<style lang="scss" scoped></style>
