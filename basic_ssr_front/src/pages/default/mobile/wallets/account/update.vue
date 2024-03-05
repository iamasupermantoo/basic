<template>
  <div class="column full-height full-width">
    <div class="col q-pa-md full-width column justify-between">
      <div class="col full-width row justify-center">
        <!-- 卡类型选择 -->
        <q-scroll-area style="height: 80px; width: 100%;" :thumb-style="{ display: 'none' }" :visible="false">
          <div class="row no-wrap">
            <div v-for="(payment, paymentIndex) in paymentList" :key="paymentIndex" :style="{
              width: '180px', height: '50px', borderRadius: '8px',
              border: paymentIndex == currentPaymentIndex ? '1px solid #01AC66' : '',
            }" class="q-pa-sm row justify-center cursor-pointer relative-position q-mr-md bg-grey-2"
              @click="switchPaymentFunc(payment, paymentIndex)">
              <q-img no-spinner class="q-mr-sm" :src="imageSrc(payment.icon)" width="32px" height="32px" />
              <div class="self-center">{{ payment.name }}</div>
              <q-img no-spinner v-if="paymentIndex == currentPaymentIndex" class="absolute" src="/images/select.png"
                width="30PX" height="30px" style="bottom: 0;right: 0;"></q-img>
            </div>
          </div>
        </q-scroll-area>


        <!-- 卡片信息 -->
        <div class="q-mt-lg full-width" v-if="paymentList[currentPaymentIndex]">
          <div class="column q-gutter-md">
            <div>
              <div class="q-mb-sm">{{ paymentList[currentPaymentIndex].type == 1 ? $t('bankName') : $t('digitalNetwork')
              }}
              </div>
              <q-select outlined v-model="currentPaymentInfo" :disable="params.id > 0"
                :options="paymentList[currentPaymentIndex].items" option-value="id" option-label="name">
                <template v-slot:selected>
                  <div class="row items-center q-gutter-sm">
                    <div>
                      <q-avatar size="md">
                        <q-img no-spinner :src="imageSrc(currentPaymentInfo.icon)"></q-img>
                      </q-avatar>
                    </div>
                    <div>{{ currentPaymentInfo.name }}</div>
                  </div>
                </template>

                <template v-slot:option="scope">
                  <q-item clickable v-close-popup @click="currentPaymentInfo = scope.opt">
                    <q-item-section>
                      <div class="row items-center q-gutter-sm cursor-pointer">
                        <div>
                          <q-avatar size="md">
                            <q-img no-spinner :src="imageSrc(scope.opt.icon)"></q-img>
                          </q-avatar>
                        </div>
                        <div>{{ scope.opt.name }}</div>
                      </div>
                    </q-item-section>
                  </q-item>
                </template>
              </q-select>
            </div>

            <div v-if="paymentList[currentPaymentIndex].type == 1">
              <div class="q-mb-sm">{{ $t('ownerName') }}</div>
              <q-input outlined v-model="params.realName" :placeholder="$t('ownerName')"></q-input>
            </div>

            <div>
              <div class="q-mb-sm">{{ paymentList[currentPaymentIndex].type == 1 ? $t('bankNumber') :
                $t('digitalAddress') }}</div>
              <q-input outlined v-model="params.number"
                :placeholder="paymentList[currentPaymentIndex].type == 1 ? $t('bankNumber') : $t('digitalAddress')"></q-input>
            </div>

            <div v-if="paymentList[currentPaymentIndex].type == 1">
              <div class="q-mb-sm">{{ $t('bankAddress') }}</div>
              <q-input outlined v-model="params.code" :placeholder="$t('bankAddress')"></q-input>
            </div>
          </div>

          <div class="q-mt-lg text-right">
            <q-btn rounded unelevated color="primary" size="lg" no-caps :label="$t('submit')" class="full-width"
              @click="submitFunc"></q-btn>
          </div>
        </div>
      </div>
    </div>

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
          <q-btn flat :label="$t('confirm')" @click="submitUpdateFunc" color="primary"></q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { useI18n } from 'vue-i18n';
import { reactive, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { imageSrc } from 'src/utils';
import { walletsPaymentIndexAPI, walletsAccountCreateAPI, walletsAccountUpdateAPI, walletsAccountInfoAPI } from 'src/apis/wallets';
import { NotifyPositive } from 'src/utils/notify';
import { InitStore } from 'stores/init';

export default {
  name: 'WalletsAccountUpdate',
  setup(props: any, context: any) {
    const { t } = useI18n()
    context.emit('update', {
      title: t('accountManage')
    })

    const $router = useRouter()
    const $initStore = InitStore()

    const state = reactive({
      walletSetting: $initStore.config.settings.wallets,
      paymentList: [] as any,
      currentPaymentIndex: 0,
      currentPaymentInfo: {} as any,
      showSecurityKey: false,

      // 提交参数
      params: {
        id: $router.currentRoute.value.query.id ?? 0,
        realName: '',
        number: '',
        code: '',
        paymentId: 0,
        securityKey: '',
      } as any
    });

    // 如果编辑卡片就要获取卡片详情
    onMounted(async () => {
      walletsPaymentIndexAPI({ modes: [] }).then((res: any) => {
        state.paymentList = res
        if (state.paymentList.length > 0) {
          switchPaymentFunc(state.paymentList[0], 0)
        }

        // 如果设置了Id, 那么请求卡片信息
        if (state.params.id > 0) {
          walletsAccountInfoAPI({ id: Number(state.params.id) }).then((res: any) => {
            state.params = res

            // 默认选中
            for (let i = 0; i < state.paymentList.length; i++) {
              for (let j = 0; j < state.paymentList[i].items.length; j++) {
                if (state.paymentList[i].items[j].id == state.params.paymentId) {
                  state.currentPaymentIndex = i
                  state.currentPaymentInfo = state.paymentList[i].items[j]
                }
              }
            }
          })

        }
      })
    })


    // 切换支付类型
    const switchPaymentFunc = (paymentInfo: any, paymentIndex: number) => {
      if (state.params.id > 0) {
        return
      }

      // 如果items长度=0，清空下拉框内容
      if (paymentInfo.items.length <= 0) {
        state.currentPaymentInfo = {}
      }

      if (paymentInfo.items.length > 0) {
        state.currentPaymentInfo = paymentInfo.items[0]
      }
      state.currentPaymentIndex = paymentIndex
    }


    // 提交信息
    const submitFunc = () => {
      state.params.paymentId = state.currentPaymentInfo.id

      // 新增提现账户
      if (state.params.id == 0) {
        walletsAccountCreateAPI(state.params).then(() => {
          NotifyPositive(t('submittedSuccess'))
          $router.push({ name: 'WalletsAccountIndex' })
        })
        return
      }

      // 编辑提现账户
      if (state.walletSetting.showSecurityPass) {
        state.showSecurityKey = true
      } else {
        submitUpdateFunc()
      }
    }

    // 提交更新提现账户
    const submitUpdateFunc = () => {
      walletsAccountUpdateAPI(state.params).then(() => {
        NotifyPositive(t('submittedSuccess'))
        state.showSecurityKey = false
        $router.push({ name: 'WalletsAccountIndex' })
      })
    }

    return {
      imageSrc,
      ...toRefs(state),
      switchPaymentFunc,
      submitUpdateFunc,
      submitFunc,
    }
  }
};
</script>

<style lang="scss" scoped></style>
