<template>
  <div class="column full-height full-width">
    <div class="col q-pa-md full-width column justify-between">
      <div class="col full-width">
        <q-card flat v-for="(account, accountIndex) in accountList" :key="accountIndex"
          :style="bgList.get(account.paymentType)" class="q-mb-sm">
          <q-card-section>
            <div class="row justify-between">
              <div class="row">
                <div>
                  <q-avatar>
                    <q-img no-spinner :src="imageSrc(account.icon)" width="48px" height="48px"></q-img>
                  </q-avatar>
                </div>
                <div class="q-ml-sm text-white">
                  <div class="q-mt-xs">
                    <div class="text-body1 text-bold">
                      {{ account.paymentName }}
                      <q-btn v-if="walletsSetting.showUpdate" rounded unelevated size="xs"
                        style="border: 1px solid #fff; padding: 0 10px;margin-left: 2px"
                        @click="$router.push({ name: 'WalletAccountUpdate', query: { id: account.id } })">
                        <div style="font-size: 12px">{{ $t('edit') }}</div>
                      </q-btn>
                    </div>
                    <div class="text-caption text-grey-2">{{ account.name }}</div>
                  </div>

                  <div class="q-mt-sm">
                    <div class="text-caption text-grey-2">{{ account.paymentType == 1 ? $t('bankNumber') :
                      $t('digitalAddress') }}</div>
                    <div v-if="walletsSetting.showNumber" class="text-body1 text-bold">{{ account.number }}</div>
                    <div v-else class="text-body1 text-bold">****{{ account.number.slice(-4) }}</div>
                  </div>
                </div>

              </div>
              <div class="text-right">
                <q-icon v-if="walletsSetting.showDelete" color="white" class="cursor-pointer" size="18px" name="cancel"
                  @click="deleteAccountFunc(account)"></q-icon>
              </div>
            </div>
          </q-card-section>
        </q-card>

        <!-- 添加按钮 -->
        <div style="border: 1px dashed #01AC66;height: 40px;background-color: rgba(1, 172, 102, 0.05);"
          class="rounded-borders column justify-center row q-mt-md"
          @click="$router.push({ name: 'WalletAccountUpdate', query: { type: 'add' } })">
          <div class="text-center text-primary text-weight-bold self-center row"> <q-icon size="20px" name="add"
              class="self-center" />Add Card
          </div>
        </div>
      </div>
    </div>

    <!-- 输入密码 -->
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
          <q-btn flat :label="$t('confirm')" @click="submitDeleteFunc" color="primary"></q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { useI18n } from 'vue-i18n';
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import { walletsAccountIndexAPI, walletsAccountDeleteAPI } from 'src/apis/wallets';
import { ConfirmPrompt, NotifyPositive } from 'src/utils/notify';
import { InitStore } from 'stores/init';

export default defineComponent({
  name: 'WalletsAccountIndex',
  emits: ['update'],
  setup(props: any, context: any) {
    const { t } = useI18n(); // 获取t函数进行翻译
    context.emit('update', {
      title: t('accountManage')
    })

    const $initStore = InitStore()

    const state = reactive({
      walletsSetting: $initStore.config.settings.wallets,
      accountList: [] as any,
      showSecurityKey: false,
      currentAccountInfo: {} as any,
      params: { id: 0, securityKey: '' },
      bgList: new Map([
        [1, { background: 'linear-gradient(90deg, #4CB8C4 0%, #3CD3AD 100%)', width: '100%', height: '132px', borderRadius: '8px' }],
        [11, { background: 'linear-gradient(90deg, #7474BF 0%, #348AC7 100%)', width: '100%', height: '132px', borderRadius: '8px' }]
      ])
    });

    onMounted(() => {
      getWalletsAccount()
    })

    // 删除账户
    const deleteAccountFunc = (account: any) => {
      ConfirmPrompt(t('deleteLabel'), t('deleteLabel') + '【' + account.name + '】?', () => {
        state.currentAccountInfo = account
        if (state.walletsSetting.showSecurityPass) {
          state.showSecurityKey = true
        } else {
          submitDeleteFunc()
        }
      }, { ok: { label: t('confirm') }, cancel: { label: t('cancel') } })
    }

    // 删除提交
    const submitDeleteFunc = () => {
      state.params.id = Number(state.currentAccountInfo.id)
      walletsAccountDeleteAPI(state.params).then(() => {
        NotifyPositive(t('submittedSuccess'))
        state.showSecurityKey = false
        getWalletsAccount()
      })
    }

    const getWalletsAccount = () => {
      walletsAccountIndexAPI({ modes: [11, 12] }).then((res: any) => {
        state.accountList = res
      })
    }

    return {
      imageSrc,
      ...toRefs(state),
      deleteAccountFunc,
      submitDeleteFunc,
    }
  }
});
</script>

<style lang="scss" scoped></style>
