<template>
  <div style="padding: 48px 100px;">
    <!-- 钱包余额 -->
    <div class="row items-center justify-between rounded-borders q-pa-lg q-pr-xl"
      style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);">
      <div class="row">
        <q-img no-spinner class="q-mr-lg" :src="imageSrc(userAssetsInfo.icon)" width="66px" height="66px"></q-img>
        <div class="q-pt-sm">
          <div class="row items-center">
            <div class="text-white text-body1 text-bold q-mr-xs">{{ userAssetsInfo.name }}</div>
            <q-icon @click="showMoney = !showMoney" class="cursor-pointer" color="white" size="16px"
              :name="showMoney ? 'o_visibility' : 'o_visibility_off'"></q-icon>
          </div>
          <div class="text-body2 text-white q-mt-sm">
            <div v-if="showMoney" class="row items-center">
              <div>{{ $t('currency') }}{{ Number(userAssetsInfo.money).toFixed(2) }}</div>
              <div class="text-caption q-ml-xs">≈{{ $t('currency') }}{{ Number(userAssetsInfo.moneyRate).toFixed(2) }}
              </div>
            </div>
            <div v-else>
              ******
            </div>
          </div>
        </div>
      </div>

      <!-- 充值、提现 -->
      <div class="row">
        <q-btn @click="$router.push({ name: 'WalletsDeposit', query: { mode: 2 } })" unelevated
          class="text-primary bg-white" rounded no-caps style="width: 80px" :label="$t('deposit')"></q-btn>
        <q-btn @click="$router.push({ name: 'WalletsWithdraw', query: { mode: 12, assetsId: params.assetsId } })"
          unelevated class="text-primary bg-white q-ml-md" rounded style="width: 80px" no-caps
          :label="$t('withdraw')"></q-btn>
      </div>
    </div>


    <!-- 账单、订单部分 -->
    <q-card flat bordered class="q-mt-lg">
      <q-card-section>
        <div class="row justify-between">
          <q-tabs v-model="tab" narrow-indicator class="q-mb-lg" @update:model-value="changeTabFunc">
            <q-tab class="text-primary rounded-borders" no-caps name="tabTransactions" :label="$t('transactions')" />
            <q-tab class="text-primary rounded-borders" no-caps name="tabBillDetails" :label="$t('billDetails')" />
          </q-tabs>
          <!-- 右侧 -->
          <div v-if="tab == 'tabBillDetails'" class="row items-center">
            <!-- 选择 -->
            <q-btn-dropdown class="bg-grey-1" style="font-weight: 400;" :label="$t('filter')" unelevated no-caps rounded>
              <div class="row q-pa-sm" style="max-width: 420px">
                <div v-for="(billType, billTypeIndex) in billFilterParams.typeList" :key="billTypeIndex" class="col-4">
                  <div class="q-ma-xs">
                    <q-btn outline rounded
                      :class="params.types.indexOf(billType.value) > -1 ? 'bg-primary text-white' : 'text-grey'"
                      @click="selectBillTypeFunc(billType.value)" class="full-width ellipsis "
                      :label="billType.label"></q-btn>
                  </div>
                </div>
              </div>
              <div class="row justify-end q-pa-sm">
                <q-btn :label="$t('cancel')" color="primary" flat v-close-popup />
                <q-btn @click="walletBillFilterFunc" :label="$t('confirm')" color="primary" flat v-close-popup />
              </div>
            </q-btn-dropdown>

            <!-- 日期选择 -->
            <q-btn class="bg-grey-1 q-ml-md" unelevated no-caps rounded style="border: 1px solid #DDDDDD;height: 32px;">
              <div class="row items-center">
                <div class="q-mr-xs text-caption">{{ billFilterParams.dateTime.from }}</div>
                <q-icon class="q-mx-sm" style="color: #DDDDDD;" size="16px" name="trending_flat"></q-icon>
                <div class="q-mr-xs text-caption">{{ billFilterParams.dateTime.to }}</div>
                <q-icon class="q-ml-sm" size="15px" name="calendar_today"></q-icon>
              </div>
              <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                <q-date v-model="billFilterParams.dateTime" range>
                  <div class="row items-center justify-end q-gutter-sm">
                    <q-btn :label="$t('cancel')" color="primary" flat v-close-popup />
                    <q-btn @click="walletBillFilterFunc" :label="$t('confirm')" color="primary" flat v-close-popup />
                  </div>
                </q-date>
              </q-popup-proxy>
            </q-btn>
          </div>
        </div>
      </q-card-section>
      <q-card-section>
        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="tabTransactions" class="no-padding">
            <div v-for="(order, orderIndex) in dataList" :key="orderIndex" class="q-px-md">
              <q-expansion-item default-opened :expand-icon-class="order.status != -1 ? 'text-transparent' : ''"
                :header-style="{ borderRadius: '4px', height: '68px', lineHeight: '54px' }">
                <template v-slot:header>
                  <q-item-section>
                    <div class="col">
                      {{ date.formatDate(order.updatedAt * 1000, 'YYYY/MM/DD HH:mm:SS') }}
                    </div>
                  </q-item-section>
                  <q-item-section>
                    <div class="text-center">
                      {{ order.name }}
                    </div>
                  </q-item-section>
                  <q-item-section>
                    <div class="text-center text-body1 text-primary">
                      {{ (order.type == 1 || order.type == 101 ? '+' : '-') }}{{ order.money.toFixed(2) }}
                    </div>
                  </q-item-section>
                  <q-item-section>
                    <div class="text-right">
                      <div v-if="order.status == -1" class="text-red">
                        {{ $t('refuse') }}
                      </div>
                      <div v-if="order.status == 10" class="text-primary">
                        {{ $t('pending') }}
                      </div>
                      <div v-if="order.status == 20" class="text-grey">
                        {{ $t('complete') }}
                      </div>
                    </div>
                  </q-item-section>
                </template>
                <div v-if="order.status == -1" class="text-red text-caption q-pa-md bg-grey-1">
                  {{ order.data }}
                </div>
              </q-expansion-item>
              <q-separator />
            </div>

            <!-- 分页 -->
            <div class="row justify-center q-mt-md" v-if="dataList.length > 0">
              <q-pagination v-model="params.pagination.page" :max="pagination.countPage"
                @update:modelValue="changePagination" input input-class="primary" color="grey">
              </q-pagination>
            </div>
          </q-tab-panel>

          <q-tab-panel name="tabBillDetails">
            <q-list>
              <div v-for="(bill, billIndex) in dataList" :key="billIndex">
                <q-item class="rounded-borders" clickable style="height: 68px; line-height: 54px">
                  <q-item-section>
                    <div class="col">
                      {{ date.formatDate(bill.createdAt * 1000, 'YYYY/MM/DD HH:mm:SS') }}
                    </div>
                  </q-item-section>
                  <q-item-section>
                    <div class="text-center">
                      {{ bill.name }}
                    </div>
                  </q-item-section>
                  <q-item-section>
                    <div class="text-center text-body1 text-primary" v-if="bill.money > 0">
                      +{{ bill.money.toFixed(2) }}
                    </div>
                    <div class="text-center text-body1 text-red" v-else>
                      {{ bill.money.toFixed(2) }}
                    </div>
                  </q-item-section>
                </q-item>
                <q-separator />
              </div>
            </q-list>
            <!-- 分页 -->
            <div class="row justify-center q-mt-md" v-if="dataList.length > 0">
              <q-pagination v-model="params.pagination.page" :max="pagination.countPage"
                @update:modelValue="changePagination" input input-class="primary" color="grey">
              </q-pagination>
            </div>
          </q-tab-panel>
        </q-tab-panels>

        <div v-if="dataList.length <= 0">
          <div class="text-center q-my-lg text-grey">
            {{ $t('noData') }}
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import {
  walletsOrderIndexAPI,
  walletsBillOptionsAPI,
  walletsBillIndexAPI,
  walletsUserAssetsInfoAPI
} from 'src/apis/wallets';
import { date } from 'quasar'
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'WalletsAssetsDetails',
  setup() {
    const $router = useRouter()
    const WalletOrderTypeDeposit = 101
    const WalletOrderTypeWithdraw = 111
    const WalletBillAccountType = -2

    const initPagination = {
      rowsPerPage: 20, //  每页显示条数
      page: 1, //  当前页数
      descending: true,
      sortBy: 'created_at',
    }
    const state = reactive({
      userAssetsInfo: {} as any,
      showMoney: true,
      tab: 'tabTransactions',
      params: {
        assetsId: $router.currentRoute.value.query.id ? Number($router.currentRoute.value.query.id) : 0,
        types: [] as any,
        pagination: initPagination,
      } as any,

      // 数据列表
      dataList: [] as any,

      //筛选过滤
      billFilterParams: {
        typeList: [] as any,
        dateTime: {
          from: date.formatDate(Date.now(), 'YYYY/MM/DD'),
          to: date.formatDate(Date.now(), 'YYYY/MM/DD'),
        },
      },

      // 分页
      pagination: {
        total: 0,
        countPage: 1,
      },
    });

    onMounted(() => {
      getWalletsOrderList()

      // 获取钱包账单Options
      walletsBillOptionsAPI({ type: WalletBillAccountType }).then((res: any) => {
        state.billFilterParams.typeList = res
      })

      // 请求用户资产信息
      walletsUserAssetsInfoAPI({ assetsId: state.params.assetsId }).then((res: any) => {
        state.userAssetsInfo = res
      })
    })

    const changeTabFunc = () => {
      state.params.pagination = initPagination
      state.params.types = []
      if (state.tab == 'tabTransactions') {
        getWalletsOrderList()
      } else {
        getWalletsBillList()
      }
    }

    // 获取钱包订单
    const getWalletsOrderList = () => {
      state.params.types = [WalletOrderTypeDeposit, WalletOrderTypeWithdraw]
      walletsOrderIndexAPI(state.params).then((res: any) => {
        state.pagination.total = res.count
        state.dataList = res.items
        state.pagination.countPage = Math.ceil(state.pagination.total / state.params.pagination.rowsPerPage)
      })
    }

    // 获取用户账单列表
    const getWalletsBillList = () => {
      if (state.params.types.length == 0) {
        state.params.types = [WalletBillAccountType]
      }
      walletsBillIndexAPI(state.params).then((res: any) => {
        state.pagination.total = res.count
        state.dataList = res.items
        state.pagination.countPage = Math.ceil(state.pagination.total / state.params.pagination.rowsPerPage)
      })
    }

    // 监听加减页
    const changePagination = (val: number) => {
      state.params.pagination.page = Number(val)
      state.tab == 'tabTransactions' ? getWalletsOrderList() : getWalletsBillList()
    }

    // 选择账单类型
    const selectBillTypeFunc = (billType: number) => {
      const typesIndex = state.params.types.indexOf(billType)
      if (typesIndex > -1) {
        state.params.types.splice(typesIndex, 1)
        return
      }
      state.params.types.push(billType)
    }

    // 钱包账单筛选确定方法
    const walletBillFilterFunc = () => {
      state.params.pagination = initPagination
      state.params.createdAt = state.billFilterParams.dateTime
      getWalletsBillList()
    }

    return {
      imageSrc,
      date,
      ...toRefs(state),
      changePagination,
      changeTabFunc,
      selectBillTypeFunc,
      walletBillFilterFunc,
    }
  }
});
</script>
<style lang="scss" scoped></style>
