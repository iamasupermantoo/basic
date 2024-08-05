<template>
  <div class="row items-center justify-center"
    :style="{ background: 'url(/images/label-bg.png)', height: '200px', backgroundSize: 'cover', backgroundRepeat: 'no-repeat' }">
    <div class="text-white text-h4">{{ $t('memberVip') }}</div>
  </div>

  <div style="margin: 80px 0 120px 0">
    <div class="text-center q-mb-xl">
      <div class="text-h5" style="font-size: 32px;">{{ $t('memberVipTitle') }}</div>
      <div class="text-body1 text-grey q-mt-md">{{ $t('memberVipSmall') }}</div>
    </div>

    <div class="column items-center">
      <q-scroll-area style="height: 440px; width: 80%" :thumb-style="{ height: '0' }">
        <div class="row no-wrap q-pt-lg q-gutter-sm">
          <q-card v-for="(level, levelIndex) in levelList" :key="levelIndex" :style="{
            width: '288px', height: '400px', borderRadius: '8px',
            boxShadow: '0px 4px 10px 0px rgba(192,192,192,0.3)',
            border: currentLevelIndex == levelIndex ? '2px solid #01AC66' : '',
          }" class="cursor-pointer" @click="currentLevelIndex = levelIndex">
            <div :style="{
              height: '160px',
              background: currentLevelIndex == levelIndex ? 'white' : 'linear-gradient(180deg, #10BE70 0%, #91DB82 100%)',
              color: currentLevelIndex == levelIndex ? 'black' : 'white'
            }" class="q-pb-md">
              <q-img no-spinner class="absolute" :src="imageSrc(level.icon)" width="54px" height="54px"
                style="top: 0; left: 50%; transform: translate(-50%, -36%);z-index: 999;"></q-img>
              <div style="height: 64px;"></div>
              <div class="text-center text-body1">{{ level.name }}</div>
              <div class="text-center">
                <span class="text-caption">{{ $t('currency') }}</span>
                <span class="text-h5 text-bold">{{ level.money }}</span>
              </div>
            </div>
            <q-card-section>
              <div class="q-mt-md">
                <q-btn :disable="level.level <= userInfo.level" @click="submitFunc(level)" no-caps outline rounded
                  color="primary" class="full-width"
                  :label="level.level <= userInfo.level ? $t('purchased') : $t('buy')"></q-btn>
              </div>
              <div v-html="level.desc" class="q-mt-lg"></div>
            </q-card-section>
          </q-card>
        </div>
      </q-scroll-area>
    </div>
  </div>
</template>

<script lang="ts">
import { onMounted, reactive, toRefs } from 'vue';
import { levelIndexAPI, levelCreateAPI } from 'src/apis/user';
import { imageSrc } from 'src/utils';
import { UserStore } from 'src/stores/user';
import { ConfirmPrompt } from 'src/utils/notify';
import { useI18n } from 'vue-i18n';

export default {
  name: 'UserLevel',
  setup() {
    const { t } = useI18n()
    const $userStore = UserStore()

    const state = reactive({
      userInfo: {} as any,
      currentLevelIndex: 0,
      levelList: [
        {} as any
      ] as any,
    });

    onMounted(() => {
      state.userInfo = $userStore.userInfo
      levelIndexAPI().then((res: any) => {
        state.levelList = res
      })
    })

    const submitFunc = (level: any) => {
      ConfirmPrompt(t('isExecute'), t('isBuy') + '【' + level.name + '】?', () => {
        levelCreateAPI({ id: level.id }).then(() => {
          state.userInfo.level = level.level
          $userStore.updateUserInfo(state.userInfo)
        })
      }, { ok: { label: t('submit') }, cancel: { label: t('cancel') } })
    }

    return {
      imageSrc,
      ...toRefs(state),
      submitFunc,
    }
  }
};
</script>
<style scoped></style>
