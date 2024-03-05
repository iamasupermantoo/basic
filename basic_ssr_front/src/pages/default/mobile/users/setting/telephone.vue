<template>
  <div>
    <div>
      <div class="row justify-center q-mt-lg q-mb-sm">
        <div class="text-h6 text-weight-bold">{{ $t('bindTelephone') }}</div>
      </div>
      <div class="row justify-center q-px-lg">
        {{ $t('bindTelephoneSmall') }}
      </div>
      <div class="q-mt-lg q-px-lg">
        <div class="row no-wrap justify-between">
          <q-btn-dropdown class="col-4 text-weight-regular rounded-borders" unelevated flat no-caps
            dropdown-icon="expand_more" style="height: 56px;border: 1px solid rgba(0, 0, 0, 0.24);">
            <template v-slot:label>
              <div class="row no-wrap items-center">
                <q-img no-spinner :src="imageSrc(countryList[countryIndex].icon ? countryList[countryIndex].icon : '')"
                  width="24px" height="16px" />
                <div class="q-ml-sm">+{{ countryList[countryIndex].code }}</div>
              </div>
            </template>
            <!-- 下拉 -->
            <q-list style="min-width: 268px" class="q-py-sm">
              <q-item @click="countryIndex = i" v-for="(item, i) in countryList" :key="i" clickable v-close-popup
                class="row no-wrap items-center">
                <q-img no-spinner class="q-mr-sm" :src="imageSrc(item.icon)" width="38px" height="38px" />
                <div>
                  <div style="font-size: 16px;">{{ item.name }}</div>
                </div>
              </q-item>
            </q-list>
          </q-btn-dropdown>
          <q-input style="width: 64%;" :placeholder="$t('telephone')" class="q-mb-lg" outlined
            v-model="params.telephone" />
        </div>
        <q-btn @click="submitFunc" class="full-width q-mb-xl" unelevated rounded no-caps size="lg" color="primary"
          :label="$t('submit')" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import { InitStore } from 'src/stores/init';
import { useRouter } from 'vue-router';
import { UserStore } from 'src/stores/user';
import { updateInfoAPI } from 'src/apis/user';
import { useI18n } from 'vue-i18n';

// 列表
export default defineComponent({
  name: 'SettingsPhoneIndex',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('telephone'),
    })

    const $router = useRouter();
    const $initStore = InitStore();
    const $userStore = UserStore();

    const state = reactive({
      countryList: $initStore.countryList as any,
      countryIndex: 0,
      params: {} as any,
      userInfo: {} as any,
    })

    onMounted(() => {
      state.userInfo = $userStore.userInfo
      state.params.telephone = $userStore.userInfo.telephone.split('|')[1]
    })

    // 执行接口
    const submitFunc = () => {
      const telephone = state.countryList[state.countryIndex].code + '|' + state.params.telephone
      updateInfoAPI({ telephone }).then(() => {
        state.userInfo.telephone = telephone
        $userStore.updateUserInfo(state.userInfo)
        $router.back()
      })
    }

    return {
      imageSrc,
      ...toRefs(state),
      submitFunc,
    }
  }
})
</script>

<style scoped></style>
