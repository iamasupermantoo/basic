<template>
  <div class="column window-height" style="padding: 48px 244px;background: #F8F9FC;">
    <q-banner v-if="accountList.length <= 0" rounded class="bg-red text-white">
      {{ $t('notBindWithdrawAccount') }}
      <template v-slot:action>
        <q-btn @click="$router.push({ name: 'WalletsAccountIndex' })" flat no-caps color="white"
          :label="$t('goto') + $t('accountManage')" />
      </template>
    </q-banner>

    <div class="col column justify-between bg-white rounded-borders q-mt-md">
      <!-- 标题 -->
      <div class="q-py-md q-px-lg row items-center no-wrap text-body1 text-weight-medium"
        style="background: linear-gradient(275deg, rgba(19,140,91,0.1) 0%, rgba(1,172,102,0.04) 100%);border-radius: 8px 8px 0 0;">
        <q-img no-spinner src="/images/withdraw.png" width="40PX" height="28px"></q-img>
        <div class="q-ml-md">{{ $t('withdraw') }}</div>
      </div>

      <div class="col full-width q-pa-lg">
        <div class="rounded-borders text-subtitle1 text-weight-medium q-py-xs q-px-md" style="background: #F8F9FC;">
          {{ $t('withdrawAccount') }}
        </div>

        <!-- 卡类型选择 -->
        <div class="row q-mt-md q-gutter-md">
          <div v-for="(account, accountIndex) in accountList" :key="accountIndex"
            @click="ActiveAccountIndex = accountIndex" :style="{
              width: '220px', height: '50px', borderRadius: '8px', background: '#F8F9FC',
              border: accountIndex == ActiveAccountIndex ? '1px solid #01AC66' : '',
            }" class="row cursor-pointer items-center relative-position">
            <q-img no-spinner class="q-ml-sm" :src="imageSrc(account.icon)" width="32px" height="32px" />
            <div class="text-body1 q-ml-sm ellipsis" style="width: 168px">{{ account.paymentName
            }}({{ account.number.slice(-4) }})</div>
            <q-img no-spinner v-if="accountIndex == ActiveAccountIndex" class="absolute" src="/images/select.png"
              width="30PX" height="30px" style="bottom: 0;right: 0;"></q-img>
          </div>
        </div>

        <!-- 提现 -->
        <div class="column q-gutter-md q-mt-lg" style="width: 60%;">
          <div>
            <div class="q-mb-xs text-grey">
              {{ $t('availableAmount') }}
            </div>
            <div class="text-bold text-body1" style="color: #F45E0C;">
              <span v-if="!$route.query.assetsId">{{ $t('currency') }}
              </span>
              <span v-if="$route.query.assetsId">
                {{ Number(userAssetsInfo.money) }}
                <span class="text-caption">{{ userAssetsInfo.name }}</span>
              </span>
              <span v-else>
                {{ Number(userInfo.money).toFixed(2) }}
              </span>

            </div>
          </div>

          <div>
            <div class="q-mb-xs text-grey">{{ $t('withdrawAmount') }}</div>
            <div>
              <q-input class="q-mr-sm" type="number" outlined v-model.number="params.money"
                :placeholder="$route.query.assetsId ? userAssetsInfo.name : $t('withdrawAmount')">
                <template v-slot:append>
                  <q-btn flat dense :label="$t('all')" color="primary"
                    @click="$route.query.assetsId ? params.money = userAssetsInfo.money : params.money = userInfo.money"></q-btn>
                </template>
              </q-input>
            </div>
          </div>

          <div class="q-mt-lg">
            <q-btn unelevated rounded color="primary" :label="$t('submit')" class="q-my-md" no-caps
              style="min-width: 100px;" @click="submitFunc" />
          </div>
        </div>
      </div>
    </div>


    <!-- 输入密码 -->
    <q-dialog v-model="showSecurityKey">
      <q-card style="width: 340px">
        <q-card-section>
          <div class="text-center text-h6">{{ $t('enterSecretKey') }}</div>
        </q-card-section>

        <q-card-section>
          <div class="q-mt-md">
            <q-input outlined v-model="params.securityKey" type="password" :label="$t('enterSecretKey')"></q-input>
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat :label="$t('cancel')" v-close-popup color="grey"></q-btn>
          <q-btn flat :label="$t('confirm')" @click="submitWithdrawFunc" color="primary"></q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue';
import { NotifyPositive } from 'src/utils/notify';
import { imageSrc } from 'src/utils';
import { walletsAccountIndexAPI, walletsUserAssetsInfoAPI, walletsWithdrawCreateAPI } from 'src/apis/wallets';
import { UserStore } from 'src/stores/user';
import { InitStore } from 'src/stores/init';
import { useI18n } from 'vue-i18n';
import { useRouter, useRoute } from 'vue-router';

export default {
  name: 'withdrawalIndex',
  setup() {
    const { t } = useI18n()
    const $router = useRouter()
    const $route = useRoute()
    const $userStore = UserStore()
    const $initStore = InitStore()

    const state = reactive({
      mode: $router.currentRoute.value.query.mode ? Number($router.currentRoute.value.query.mode) : 0,
      config: $initStore.config,
      showSecurityKey: false,
      userInfo: {} as any,
      params: {
        money: '',
        securityKey: '',
      } as any,

      // 选中卡片类型
      ActiveAccountIndex: 0,

      // 银行卡列表
      accountList: [] as any,

      // 资产页面点击提现后的余额
      userAssetsInfo: {} as any,
    });

    onMounted(() => {
      state.userInfo = $userStore.userInfo
      walletsAccountIndexAPI({ modes: [Number(state.mode)] }).then((res: any) => {
        state.accountList = res
      })

      // 如果assetsId存在，获取资产
      if ($route.query.assetsId) {
        walletsUserAssetsInfoAPI({ assetsId: Number($route.query.assetsId) }).then((res: any) => {
          state.userAssetsInfo = res
        })
      }
    })

    const submitFunc = () => {
      if (state.config.settings.online.depositLink) {
        window.location.href = state.config.onlineLink
        return
      }

      if (state.config.settings.wallets.showSecurityPass) {
        state.showSecurityKey = true
      } else {
        submitWithdrawFunc()
      }
    }

    // 提现
    const submitWithdrawFunc = () => {
      state.params.accountId = state.accountList[state.ActiveAccountIndex].id
      state.params.money = Number(state.params.money)

      // 关闭密码弹窗
      state.showSecurityKey = false
      walletsWithdrawCreateAPI(state.params).then(() => {
        NotifyPositive(t('submittedSuccess'))

        // 如果是余额充值跳转到钱包列表, 资产充值跳转到资产列表
        if (state.mode == 11) {
          $router.push({ name: 'WalletsIndex' })
        } else {
          $router.push({ name: 'WalletsAssetsIndex' })
        }
      })
    }

    return {
      imageSrc,
      ...toRefs(state),
      submitFunc,
      submitWithdrawFunc,
    }
  }
};
</script>

<style lang="scss" scoped></style>
