<template>
  <div class="column full-height full-width">
    <div class="col q-pa-md full-width column justify-between bg-white">
      <div class="text-body2 text-weight-medium q-pb-xs">{{ $t('depositAccount') }}</div>
      <!-- 卡类型选择 -->
      <q-scroll-area style="height: 80px; width: 100%;" :thumb-style="{ display: 'none' }" :visible="false">
        <div class="row no-wrap">
          <template v-for="(payment, paymentIndex) in paymentList" :key="paymentIndex">
            <div v-for="(children, childrenIndex) in payment.items" :key="childrenIndex" :style="{
              width: '180px', height: '50px', borderRadius: '8px',
              border: children.id == currentPaymentInfo.id ? '1px solid #01AC66' : '',
            }" class="q-pa-sm row cursor-pointer relative-position q-mr-md bg-grey-2"
              @click="switchPaymentFunc(children)">
              <q-img no-spinner class="q-mr-sm" :src="imageSrc(children.icon)" width="32px" height="32px" />
              <div class="self-center">{{ children.name }}</div>
              <q-img no-spinner v-if="children.id == currentPaymentInfo.id" class="absolute" src="/images/select.png"
                width="30PX" height="30px" style="bottom: 0;right: 0;"></q-img>
            </div>
          </template>
        </div>
      </q-scroll-area>

      <div>
        <div v-if="currentPaymentInfo.type == paymentTypeDigital">
          <div class="column items-center">
            <q-card flat style="border-radius: 10px;width: 172px;">
              <q-card-section>
                <img :src="currentQrcode" alt="">
              </q-card-section>
            </q-card>
            <div class="q-mt-md" style="width: 310px">
              <q-input outlined v-model="currentPaymentInfo.dataJson.number" readonly>
                <template v-slot:append>
                  <q-icon name="content_copy" class="cursor-pointer" size="xs"
                    @click="copyToClipboardFunc(currentPaymentInfo.dataJson.number)"></q-icon>
                </template>
              </q-input>
            </div>
          </div>
        </div>

        <div style="background: #F5F6FA;border-radius: 8px;" v-else-if="currentPaymentInfo.type == paymentTypeDebitCard"
          class="q-pa-md">
          <div class="q-pb-sm" style="border-bottom: 1px dashed #DDDDDD">{{ $t('depositAccountInfo') }}</div>
          <div class="row justify-between items-center q-mt-sm">
            <div>
              {{ currentPaymentInfo.dataJson.name }}
            </div>
            <div>
              <q-btn flat :label="$t('copy')" color="primary"
                @click="copyToClipboardFunc(currentPaymentInfo.dataJson.name)"></q-btn>
            </div>
          </div>
          <div class="row justify-between items-center">
            <div>
              {{ currentPaymentInfo.dataJson.realName }}
            </div>
            <div>
              <q-btn flat :label="$t('copy')" color="primary"
                @click="copyToClipboardFunc(currentPaymentInfo.dataJson.realName)"></q-btn>
            </div>
          </div>
          <div class="row justify-between items-center">
            <div>
              {{ currentPaymentInfo.dataJson.number }}
            </div>
            <div>
              <q-btn flat :label="$t('copy')" color="primary"
                @click="copyToClipboardFunc(currentPaymentInfo.dataJson.number)"></q-btn>
            </div>
          </div>
          <div class="row justify-between items-center">
            <div>
              {{ currentPaymentInfo.dataJson.code }}
            </div>
            <div>
              <q-btn flat :label="$t('copy')" color="primary"
                @click="copyToClipboardFunc(currentPaymentInfo.dataJson.code)"></q-btn>
            </div>
          </div>
        </div>
      </div>

      <!-- 充值信息 -->
      <div class="column full-width q-mt-md">
        <div class="q-my-sm">
          <div class="q-mb-sm">{{ $t('depositAmount') }}</div>
          <q-input type="number" :placeholder="$t('depositAmount')" outlined v-model="params.money" />
        </div>

        <div class="q-my-sm">
          <div class="q-mb-xs">
            {{ $t('depositProof') }}
          </div>
          <div>
            <uploader @uploaded="(url) => { params.voucher = url }">
              <template v-slot:default>
                <q-uploader-add-trigger />
                <q-card flat>
                  <div class="column items-center justify-center" style="height: 150px;border: grey 1px dashed">
                    <q-icon name="add" color="grey" size="30px" v-if="params.voucher == ''" />
                    <q-img no-spinner v-else :src="imageSrc(params.voucher)"></q-img>
                  </div>
                </q-card>
              </template>
            </uploader>
          </div>
        </div>

        <div class="q-mt-lg q-mb-md text-grey">
          <div v-html="config.depositTips"></div>
        </div>

        <q-btn unelevated rounded color="primary" :label="$t('submit')" class="q-my-md" no-caps @click="submitFunc"
          size="lg" />
      </div>

    </div>

  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted } from 'vue';
import { NotifyPositive } from 'src/utils/notify';
import { imageSrc } from 'src/utils';
import uploader from 'src/components/uploader.vue';
import { walletsPaymentIndexAPI, walletsDepositCreateAPI } from 'src/apis/wallets';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { InitStore } from 'src/stores/init';
import QRCode from 'qrcode-svg-ts';
import { copyToClipboard } from 'quasar';

export default {
  name: 'WalletsDepositIndex',
  components: { uploader },
  setup(props: any, context: any) {
    const { t } = useI18n()
    context.emit('update', {
      title: t('deposit')
    })

    const $router = useRouter()
    const $initStore = InitStore()
    const paymentTypeDigital = 11
    const paymentTypeDebitCard = 1

    const state = reactive({
      mode: $router.currentRoute.value.query.mode ? Number($router.currentRoute.value.query.mode) : 0,
      params: {
        voucher: '',
      } as any,
      currentQrcode: '',
      paymentList: [] as any,
      config: $initStore.config,
      currentPaymentInfo: {} as any,
    });

    onMounted(() => {
      walletsPaymentIndexAPI({ modes: [state.mode] }).then((res: any) => {
        state.paymentList = res
        for (let index: any = 0; index < state.paymentList.length; index++) {
          if (state.paymentList[index].items.length > 0) {
            switchPaymentFunc(state.paymentList[index].items[0])
            break;
          }
        }
      })
    })

    // 充值
    const submitFunc = () => {
      if (state.config.settings.online.depositLink) {
        window.location.href = state.config.onlineLink
        return
      }

      state.params.paymentId = state.currentPaymentInfo.id
      state.params.money = Number(state.params.money)
      walletsDepositCreateAPI(state.params).then(() => {
        NotifyPositive(t('submittedSuccess'))

        // 如果是余额充值跳转到钱包列表, 资产充值跳转到资产列表
        if (state.mode == 1) {
          $router.push({ name: 'WalletsIndex' })
        } else {
          $router.push({ name: 'WalletsAssetsIndex' })
        }
      })
    }

    // 切换支付类型
    const switchPaymentFunc = (paymentInfo: any) => {
      state.currentPaymentInfo = paymentInfo

      //  如果是数字货币, 那么生成二维码
      if (state.currentPaymentInfo.type == paymentTypeDigital) {
        const qrCode = new QRCode({
          content: state.currentPaymentInfo.dataJson.number,
          width: 138,
          height: 138,
          color: '#000000',
          background: '#ffffff',
          ecl: 'M',
        });
        state.currentQrcode = qrCode.toDataURL()
      }
    }

    // 复制方法
    const copyToClipboardFunc = (str: string) => {
      copyToClipboard(str)
        .then(() => {
          NotifyPositive(t('copySuccess'))
        })
    };

    return {
      imageSrc,
      paymentTypeDebitCard,
      paymentTypeDigital,
      ...toRefs(state),
      switchPaymentFunc,
      submitFunc,
      copyToClipboardFunc,
    }
  }
};
</script>

<style lang="scss" scoped></style>
