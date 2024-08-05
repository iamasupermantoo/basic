<template>
  <div class="column full-height full-width">
    <div v-if="accountList.length <= 0" class="q-pa-md">
      <q-banner rounded class="bg-red text-white q-mt-sm">
        {{ $t('notBindWithdrawAccount') }}
        <template v-slot:action>
          <q-btn @click="$router.push({ name: 'WalletsAccountIndex' })" flat no-caps color="white"
            :label="$t('goto') + $t('accountManage')" />
        </template>
      </q-banner>
    </div>

    <div class="q-pa-md full-width column q-gutter-xl">
      <div class="col full-width">
        <div class="text-body2 text-weight-medium q-pb-xs">{{ $t('withdrawAccount') }}</div>
        <!-- 卡类型选择 -->
        <q-scroll-area style="height: 80px; width: 100%;" :thumb-style="{ display: 'none' }" :visible="false">
          <div class="row no-wrap">
            <div v-for="(account, accountIndex) in accountList" :key="accountIndex"
              @click="activeAccountIndex = accountIndex" :style="{
                width: '200px', height: '50px', borderRadius: '8px',
                border: accountIndex == activeAccountIndex ? '1px solid #01AC66' : '',
              }" class="row cursor-pointer items-center relative-position q-mr-md no-wrap bg-grey-2">
              <q-img no-spinner class="q-ml-sm" :src="imageSrc(account.icon)" width="32px" height="32px" />
              <div class="text-body1 q-ml-sm ellipsis" style="width: 168px">{{ account.paymentName
              }}({{ account.number.slice(-4) }})</div>
              <q-img no-spinner v-if="accountIndex == activeAccountIndex" class="absolute" src="/images/select.png"
                width="30PX" height="30px" style="bottom: 0;right: 0;"></q-img>
            </div>
          </div>
        </q-scroll-area>

        <div>
          <div class="text-body2 text-weight-medium q-pb-xs">{{ $t('withdrawAmount') }}</div>
          <q-input type="number" outlined v-model="params.money" class="q-mb-sm"
            :placeholder="$route.query.assetsId ? userAssetsInfo.name : $t('withdrawAmount')">
            <template v-slot:append>
              <div @click="$route.query.assetsId ? params.money = userAssetsInfo.money : params.money = userInfo.money"
                style="font-size: 14px" class="text-primary">{{ $t('all')
                }}
              </div>
            </template>
          </q-input>
          <div style="color: #F45E0C;">
            <span class="text-grey-7">{{ $t('availableAmount') }}:</span>

            <span v-if="!$route.query.assetsId">{{ $t('currency') }}</span>
            <span v-if="$route.query.assetsId">
              {{ Number(userAssetsInfo.money) }}
              <span class="text-caption">{{ userAssetsInfo.name }}</span>
            </span>
            <span v-else>
              {{ Number(userInfo.money).toFixed(2) }}
            </span>
          </div>
        </div>
      </div>

      <!-- 提现按钮 -->
      <q-btn unelevated rounded color="primary" :label="$t('withdraw')" class="full-width q-my-md q-mt-xl" size="lg"
        no-caps @click="submitFunc" />
    </div>

    <!-- 安全密码 -->
    <q-dialog v-model="showSecurityKey">
      <q-card style="width: 340px">
        <q-card-section>
          <div class="text-center text-body1">{{ $t('enterSecretKey') }}</div>
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
  name: 'WalletsWithdrawIndex',
  setup(props: any, context: any) {
    const { t } = useI18n()
    context.emit('update', {
      title: t('withdraw'),
    })

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
      activeAccountIndex: 0,

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
      state.params.accountId = state.accountList[state.activeAccountIndex].id
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
