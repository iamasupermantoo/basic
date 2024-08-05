<template>
  <div>
    <div>
      <div class="row justify-center q-py-lg">
        <div class="relative-position">
          <uploader @uploaded="(url) => { params.avatar = url }">
            <template v-slot:default>
              <q-uploader-add-trigger />
              <q-avatar size="80px">
                <q-img no-spinner :src="imageSrc(params.avatar)" />
                <q-badge floating class="bg-transparent" :style="{ top: '70%' }">
                  <q-img no-spinner src='/images/icons/edit.png' width="28px" height="28px"></q-img>
                </q-badge>
              </q-avatar>
            </template>
          </uploader>
        </div>
      </div>

      <div class="q-mt-lg q-px-lg">
        <div class="q-mb-md">
          <div class="text-weight-bold q-mb-sm">{{ $t('nickname') }}</div>
          <q-input color="green" :placeholder="$t('nickname')" outlined v-model="params.nickname" />
        </div>

        <div class="q-mb-md">
          <div class="text-weight-bold q-mb-sm">{{ $t('sex') }}</div>
          <q-select outlined v-model="params.sex" map-options :options="GenderList" option-value="value"
            option-label="name" emit-value dropdown-icon="expand_more" />
        </div>

        <div div class=" q-mb-md">
          <div class="text-weight-bold q-mb-sm">{{ $t('birthday') }}</div>
          <q-input @click="birthdayPopup = true" :placeholder="$t('birthday')" outlined v-model="params.birthdayStr"
            mask="date" class="q-mb-lg">
            <template v-slot:append>
              <q-popup-proxy v-model="birthdayPopup">
                <q-date v-model="params.birthdayStr">
                  <div class="row items-center justify-end">
                    <q-btn @click="birthdayPopup = false" no-caps :label="$t('confirm')" color="primary" flat />
                  </div>
                </q-date>
              </q-popup-proxy>
            </template>
          </q-input>
        </div>

        <q-btn @click="submitFunc" size="lg" class="full-width q-mb-xl row justify-center" unelevated rounded no-caps
          color="primary" :label="$t('submit')" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { imageSrc } from 'src/utils';
import { updateInfoAPI } from 'src/apis/user';
import { UserStore } from 'src/stores/user';
import uploader from 'src/components/uploader.vue';
import { date } from 'quasar';
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'SettingsUpdateInfo',
  components: {
    uploader,
  },
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', { title: t('settings') })

    const $userStore = UserStore();
    const $router = useRouter();
    const state = reactive({
      params: {} as any,
      GenderList: [
        { name: t('unknown'), value: 0 },
        { name: t('male'), value: 1 },
        { name: t('female'), value: 2 },
      ],
      //date选择器状态
      birthdayPopup: false,
    })

    onMounted(() => {
      state.params = $userStore.userInfo
      state.params.birthdayStr = date.formatDate($userStore.userInfo.birthday * 1000, 'YYYY-MM-DD')
    })

    // 执行接口
    const submitFunc = () => {
      updateInfoAPI(state.params).then(() => {
        $userStore.updateUserInfo(state.params)
        void $router.back()
      })
    }

    return {
      imageSrc,
      date,
      ...toRefs(state),
      submitFunc,
    }
  }
})
</script>

<style scoped lang="scss"></style>
