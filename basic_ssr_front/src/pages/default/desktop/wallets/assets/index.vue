<template>
  <div style="padding: 48px 100px;">
    <div class="row">
      <div style="width: 300px;">
        <q-card flat class="full-width" style="background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%)">
          <q-card-section style="padding: 16px 16px 8px">
            <div class="row items-center">
              <div class="text-body1 text-bold text-white q-mr-xs">{{ $t('totalAssets') }}</div>
              <q-icon color="white" @click="showMoney = !showMoney"
                :name="showMoney ? 'o_visibility' : 'o_visibility_off'" class="cursor-pointer" size="16px"></q-icon>
            </div>
          </q-card-section>
          <q-card-section style="padding: 8px 16px">
            <div class="text-white">
              <div class="text-h6">{{ $t('currency') }}{{ userAssetsInfo.moneyRateSum.toFixed(2) }}</div>
              <div class="text-caption">≈{{ $t('currency') }}{{ userAssetsInfo.moneyRateSum.toFixed(2) }}</div>
            </div>
          </q-card-section>
          <q-card-actions align="right">
            <q-btn @click="$router.push({ name: 'WalletsDeposit', query: { mode: 2 } })" flat rounded
              class="bg-white text-primary" size="sm" style="padding: 6px 18px">
              <div>{{ $t('deposit') }}</div>
            </q-btn>
            <q-btn @click="toWithdraw" flat rounded class="bg-white text-primary" size="sm" style="padding: 6px 18px">
              <div>{{ $t('withdraw') }}</div>
            </q-btn>
          </q-card-actions>
        </q-card>
      </div>
      <div class="col q-ml-sm">
        <q-scroll-area style="height: 162px;width: 100%;" :thumb-style="{ display: 'none' }">
          <div class="row no-wrap q-gutter-sm">
            <q-card bordered flat
              style="width: 278px; background: linear-gradient(180deg, rgba(3,179,107,0.14) 0%, rgba(255,255,255,0) 100%)"
              v-for="(assets, assetsIndex) in userAssetsInfo.userAssetsList" :key="assetsIndex">
              <q-card-section style="padding: 16px 16px 0">
                <div class="row justify-between items-center">
                  <div class="text-h6">{{ assets.name }}</div>
                  <div>
                    <q-icon :name="'img:' + imageSrc(assets.icon)" size="42px"></q-icon>
                  </div>
                </div>
              </q-card-section>
              <q-card-section style="padding: 4px 16px 0">
                <div class="text-black">
                  <div class="text-h6">{{ $t('currency') }}{{ assets.money.toFixed(2) }}</div>
                  <div class="text-caption text-grey">≈{{ $t('currency') }}{{ assets.moneyRate.toFixed(2) }}</div>
                </div>
              </q-card-section>

              <q-card-actions align="right">
                <q-btn flat rounded class="bg-green-1 text-primary" size="sm" style="padding: 6px 18px"
                  @click="$router.push({ name: 'WalletsAssetsDetails', query: { id: assets.walletAssetsId } })">
                  <div>{{ $t('views') }}</div>
                </q-btn>
              </q-card-actions>
            </q-card>
          </div>
        </q-scroll-area>
      </div>
    </div>

    <!--    折线图-->
    <div class="q-mt-xl">
      <div :id="echartsDomId" style="height: 400px; width: 100%"></div>
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
                      {{ order.money.toFixed(2) }}
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
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { UserStore } from 'src/stores/user';
import * as echarts from 'echarts'
import { walletsUserAssetsIndexAPI, walletsBillIndexAPI, walletsOrderIndexAPI, walletsBillOptionsAPI } from 'src/apis/wallets';
import { imageSrc } from 'src/utils';
import { NotifyNegative } from 'src/utils/notify';
import { date } from 'quasar'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

export default defineComponent({
  name: 'WalletsAssetsIndex',
  setup() {
    const WalletBillAccountType = -2
    const WalletOrderTypeAssetsDeposit = 2
    const WalletOrderTypeAssetsWithdraw = 12

    const $userStore = UserStore()
    const $router = useRouter()
    const { t } = useI18n()
    const initPagination = {
      rowsPerPage: 20, //  每页显示条数
      page: 1, //  当前页数
      descending: true,
      sortBy: 'created_at',
    }

    const echartsDomId = 'echartsId'
    const state = reactive({
      echarts: {} as any,
      userAssetsInfo: { moneySum: 0, moneyRateSum: 0, userAssetsList: [] as any } as any,
      userInfo: {} as any,
      showMoney: true,
      tab: 'tabTransactions',
      params: {
        types: [] as any,
        pagination: initPagination,
      } as any,

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

      // 数据列表
      dataList: [] as any,
    });

    onMounted(() => {
      state.userInfo = $userStore.userInfo
      getWalletsOrderList()

      // 获取钱包账单Options
      walletsBillOptionsAPI({ type: WalletBillAccountType }).then((res: any) => {
        state.billFilterParams.typeList = res
      })

      // 获取用户资产列表
      walletsUserAssetsIndexAPI().then((res: any) => {
        state.userAssetsInfo = res
        state.echarts = state.userAssetsInfo.echarts

        // 初始化图形
        chartSetOptions()
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
      state.params.types = [WalletOrderTypeAssetsDeposit, WalletOrderTypeAssetsWithdraw]
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

    const chartSetOptions = () => {
      const chartDom = document.getElementById(echartsDomId) as HTMLElement;
      const myChart = echarts.init(chartDom);
      let option: echarts.EChartsOption;
      const legendList = [];
      for (let i = 0; i < state.echarts.series.length; i++) {
        legendList.push(state.echarts.series[i].name);
      }
      option = {
        tooltip: { trigger: 'axis' },
        legend: { data: legendList, bottom: '0' },
        grid: { left: '0', right: '0', bottom: '36px', containLabel: true },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: state.echarts.category,
        },
        yAxis: { type: 'value' },
        series: state.echarts.series as any,
      };
      myChart.setOption(option);
      window.addEventListener('resize', () => {
        setTimeout(myChart.resize, 300);
      });
    };

    // 点击提现
    const toWithdraw = () => {
      if (state.userAssetsInfo.userAssetsList && state.userAssetsInfo.userAssetsList.length <= 0) {
        NotifyNegative(t('notAssets'))
      } else {
        $router.push({ name: 'WalletsWithdraw', query: { mode: 12, assetsId: state.userAssetsInfo.userAssetsList[0].walletAssetsId } })
      }
    }

    return {
      imageSrc,
      date,
      echartsDomId,
      ...toRefs(state),
      changePagination,
      changeTabFunc,
      selectBillTypeFunc,
      walletBillFilterFunc,
      toWithdraw,
    }
  }
});
</script>
<style lang="scss" scoped></style>
