<template>
  <div class="column full-width">
    <div class="col page_bg q-pa-md full-width">

      <div style="height: 112px;background: linear-gradient(93deg, #10BE70 0%, #91DB82 100%);padding: 0 20px;"
        class="row justify-between rounded-borders q-mb-md">
        <div class="column justify-center">
          <div class="row items-center">
            <div class="text-white">{{ $t('totalAssets') }}</div>
            <q-icon @click="showMoney = !showMoney" :name="!showMoney ? 'o_visibility' : 'o_visibility_off'"
              class="q-ml-sm cursor-pointer text-white" size="18px"></q-icon>
          </div>
          <!-- 点击显示、隐藏金额 -->
          <div v-if="showMoney" class="text-white row items-center">
            <span class="q-mr-sm text-weight-bold" style="font-size: 22px;">{{ $t('currency') }}{{
              userAssetsInfo.moneyRateSum.toFixed(2) }}</span>
            <span>≈{{ $t('currency') }}{{ userAssetsInfo.moneyRateSum.toFixed(2) }} </span>
          </div>
          <div v-else class="text-white text-weight-bold " style="font-size: 22px;">**** </div>
        </div>
      </div>

      <!-- 充值提现 -->
      <div class="row justify-between q-mb-md btn">
        <q-btn @click="$router.push({ name: 'WalletsDeposit', query: { mode: 2 } })" style="width: 47%;"
          class="bg-white q-py-sm rounded-borders" no-caps unelevated>
          <div class="row justify-start items-center">
            <q-img no-spinner class="q-mr-sm" :src="imageSrc('/assets/icon/menu/deposit.png')" width="42px"
              height="42px" />
            <div>{{ $t('deposit') }}</div>
          </div>
        </q-btn>
        <q-btn @click="toWithdraw" style="width: 47%;" class="bg-white q-py-sm rounded-borders" no-caps unelevated>
          <div class="row justify-start items-center">
            <q-img no-spinner class="q-mr-sm" :src="imageSrc('/assets/icon/menu/withdraw.png')" width="42px"
              height="42px" />
            <div>{{ $t('withdraw') }}</div>
          </div>
        </q-btn>
      </div>

      <!-- eCharts -->
      <div class="bg-white q-pa-md rounded-borders" v-show="userAssetsInfo.userAssetsList.length > 0">
        <div class="text-weight-bold q-mb-lg">{{ $t('assetsBlock') }}
        </div>
        <div class="row justify-center">
          <div @click="chartSetOptions(i)" v-for="(item, i) in echartsTypeList" :key="i"
            :class="['q-mx-xs q-px-md q-py-xs', { 'text-white': i == echartsTypeIndex, 'bg-primary': i == echartsTypeIndex, 'text-grey-8': i != echartsTypeIndex, 'bg-grey-3': i != echartsTypeIndex, }]"
            style="border-radius: 18px;">
            {{ item.name }}
          </div>
        </div>

        <div class="q-mt-lg" v-show="echartsTypeIndex == 0">
          <!-- 饼状图 -->
          <div class="row justify-center q-mb-sm">
            <div :id="echartsPieId" style="height: 200px; width: 200px"></div>
          </div>

          <!-- 资产账户列表 -->
          <div>

            <div v-for="(assets, assetsIndex) in userAssetsInfo.userAssetsList" :key="assetsIndex"
              @click="$router.push({ name: 'WalletsAssetsDetails', query: { id: assets.walletAssetsId } })">
              <div class="row justify-between items-center justify-start">
                <div class="row items-center">
                  <q-separator vertical class="q-mr-xs q-mt-xs" style="width: 2px;
                    height: 12px;
                    border-radius: 2px;" :style="{ background: assets.itemStyle.color }" />
                  <div class="q-mr-md text-grey-7">{{ assets.name }}</div>
                  <div>- {{ (assets.moneyRate / userAssetsInfo.moneyRateSum * 100).toFixed(2) }}%</div>
                </div>
                <div class="text-weight-bold">{{ $t('currency') }}{{ assets.moneyRate.toFixed(2) }}</div>
              </div>
              <q-separator v-if="assetsIndex < userAssetsInfo.userAssetsList.length - 1" class="q-my-xs"
                style="background: #F4F5FD" />
            </div>
          </div>
        </div>

        <div v-show="echartsTypeIndex == 1">
          <div class="row justify-center q-mb-sm full-width">
            <div :id="echartsLineId" style="height: 260px; width: 100%"></div>
          </div>
        </div>

      </div>

      <!--  -->
      <div class="q-mt-md q-mb-sm text-weight-bold">{{ $t('assetsBlock') }}</div>
      <div @click="$router.push({ name: 'WalletsAssetsDetails', query: { id: assets.walletAssetsId } })"
        v-for="(assets, assetsIndex) in userAssetsInfo.userAssetsList" :key="assetsIndex"
        class="row justify-between items-center bg-white q-py-sm q-px-md q-mb-sm rounded-borders">
        <div class="row items-center">
          <q-img no-spinner class="q-mr-sm" width="26px" height="26px" :src="imageSrc(assets.icon)" />
          <div class="text-weight-bold">{{ assets.name }}</div>
        </div>
        <div>
          <div class="text-weight-bold text-right" style="font-size: 16px;">{{ assets.money
          }}</div>
          <div class="text-right text-grey-5" style="font-size: 12px;">≈{{ $t('currency') }}{{ assets.moneyRate.toFixed(2)
          }}</div>
        </div>
      </div>

      <div v-if="userAssetsInfo.userAssetsList.length <= 0" class="text-grey text-center q-py-lg">
        {{ $t('noData') }}
      </div>

    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { InitStore } from 'src/stores/init';
import * as echarts from 'echarts'
import { walletsUserAssetsIndexAPI } from 'src/apis/wallets';
import { imageSrc } from 'src/utils';
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n';
import { NotifyNegative } from 'src/utils/notify';


export default defineComponent({
  name: 'WalletsAssetsIndex',
  setup(props: any, context: any) {
    const { t } = useI18n()
    context.emit('update', {
      title: t('myAssets'),
    })

    const $initStore = InitStore()
    const $router = useRouter()
    const echartsLineId = 'line'
    const echartsPieId = 'pie'
    const colorList = [
      '#3F82FE', '#14C9C9', '#F7BA1E',
      '#3F82FE', '#14C9C9', '#F7BA1E',
      '#3F82FE', '#14C9C9', '#F7BA1E',
      '#3F82FE', '#14C9C9', '#F7BA1E'
    ]

    const state = reactive({
      echarts: {} as any,
      userAssetsInfo: { moneySum: 0, moneyRateSum: 0, userAssetsList: [] as any } as any,
      showMoney: true,
      quickMenuList: $initStore.quickMenu as any,

      // 饼、折线图button
      echartsTypeList: [
        { name: t('pieChart'), value: echartsPieId },
        { name: t('lineChart'), value: echartsLineId },
      ],
      echartsTypeIndex: 0,

      // 饼图eCharts
      pieOption: {} as any,
      // 折线图eCharts
      lineOption: {} as any,

      myChart: '' as any,
    });

    onMounted(() => {
      // 获取用户资产列表
      walletsUserAssetsIndexAPI().then((res: any) => {
        state.userAssetsInfo = res
        state.echarts.series = state.userAssetsInfo.echarts
        state.userAssetsInfo.userAssetsList.forEach((item: any, i: any) => {
          item.value = item.moneyRate
          item.itemStyle = {
            color: colorList[i]
          }
        });

        // 初始化图形
        state.myChart = echarts.init(document.getElementById(echartsPieId));
        chartSetOptions(0)
      })
    })

    const chartSetOptions = (index: any) => {
      state.myChart.dispose()
      state.echartsTypeIndex = index
      if (index == 0) {
        chartSetPieOptions()
      } else {
        setTimeout(() => {
          chartSetLineOptions()
        }, 100)

      }
    }

    // 初始化饼状图形
    const chartSetPieOptions = () => {
      state.myChart = echarts.init(document.getElementById(echartsPieId));
      // 饼状图
      state.myChart.setOption(state.pieOption = {
        title: {
          text: t('totalAssets'),
          subtext: state.userAssetsInfo.moneyRateSum,
          left: 'center', // 标题居中
          top: '35%',
          textStyle: { // 标题样式
            color: '#4E5969', // 标题颜色
            fontSize: '12px',
            textDecoration: 'underline' // 标题装饰
          },
          subtextStyle: { // 子标题样式
            color: '#1D2129', // 子标题颜色
            fontStyle: 'bold', // 子标题字体样式
            fontSize: '12px',
          },
          padding: [10, 10], // 标题与内容间距
          itemGap: 8 // 同一级标签间距
        },
        series: {
          type: 'pie',
          radius: ['55%', '90%'],
          label: {
            show: false,
          },
          data: state.userAssetsInfo.userAssetsList
        },
      });
      window.addEventListener('resize', () => {
        setTimeout(state.myChart.resize, 300);
      });
    };

    // 初始化折线图形
    const chartSetLineOptions = () => {
      state.myChart = echarts.init(document.getElementById(echartsLineId));
      const legendList = [];

      for (let i = 0; i < state.echarts.series.length; i++) {
        legendList.push(state.echarts.series[i].name);
      }
      // 折线图
      state.myChart.setOption({
        tooltip: { trigger: 'axis' },
        legend: { data: legendList, bottom: '0' },
        grid: { left: '0', right: '0', bottom: '36px', containLabel: true },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: state.echarts.category,
        },
        yAxis: { type: 'value' },
        series: state.echarts.series.series as any,
      });
      window.addEventListener('resize', () => {
        setTimeout(state.myChart.resize, 300);
      });
    };

    // 点击提现
    const toWithdraw = () => {
      if (state.userAssetsInfo.userAssetsList && state.userAssetsInfo.userAssetsList.length <= 0) {
        NotifyNegative(t('notAssets'))
        return false
      }
      $router.push({ name: 'WalletsWithdraw', query: { mode: 12, assetsId: state.userAssetsInfo.userAssetsList[0].walletAssetsId } })
    }

    return {
      imageSrc,
      ...toRefs(state),
      chartSetOptions,
      colorList,
      echartsLineId,
      echartsPieId,
      toWithdraw,
    }
  },
});
</script>
<style scoped></style>
