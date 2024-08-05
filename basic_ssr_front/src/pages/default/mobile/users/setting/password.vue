<template>
  <div>
    <div>
      <div class="q-mt-lg q-px-lg">
        <q-form>
          <q-input class="q-mb-md" v-model="oldPassword" outlined :type="showPwd.oldPwd ? 'password' : 'text'"
            :label="$t('oldPassword')">
            <template v-slot:prepend>
              <q-img no-spinner class="iconLogo" src="/images/default/password.png" />
            </template>
            <template v-slot:append>
              <q-icon style="color: #999999;" :name="showPwd.oldPwd ? 'visibility_off' : 'visibility'"
                class="cursor-pointer" @click="showPwd.oldPwd = !showPwd.oldPwd" />
            </template>
          </q-input>
          <q-input class="q-mb-md" v-model="newPassword" outlined :type="showPwd.newPwd ? 'password' : 'text'"
            :label="$t('newPassword')">
            <template v-slot:prepend>
              <q-img no-spinner class="iconLogo" src="/images/default/password.png" />
            </template>
            <template v-slot:append>
              <q-icon style="color: #999999;" :name="showPwd.newPwd ? 'visibility_off' : 'visibility'"
                class="cursor-pointer" @click="showPwd.newPwd = !showPwd.newPwd" />
            </template>
          </q-input>
          <q-input class="q-mb-md" v-model="cmfPassword" outlined :type="showPwd.cmfPwd ? 'password' : 'text'"
            :label="$t('cmfPassword')">
            <template v-slot:prepend>
              <q-img no-spinner class="iconLogo" src="/images/default/password.png" />
            </template>
            <template v-slot:append>
              <q-icon style="color: #999999;" :name="showPwd.cmfPwd ? 'visibility_off' : 'visibility'"
                class="cursor-pointer" @click="showPwd.cmfPwd = !showPwd.cmfPwd" />
            </template>
          </q-input>
          <q-btn @click="submitFunc()" size="lg" class="full-width q-mb-xl" unelevated rounded no-caps color="primary"
            :label="$t('confirm')" />
        </q-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue';
import { useRouter } from 'vue-router'
import { updatePasswordAPI } from 'src/apis/user';
import { useI18n } from 'vue-i18n';
import { NotifyNegative } from 'src/utils/notify';

// 列表
export default defineComponent({
  name: 'SettingsPasswordIndex',
  setup(props: any, context: any) {
    const { t } = useI18n();
    context.emit('update', {
      title: t('updatePassword'),
    })

    const $router = useRouter();

    const state = reactive({
      showPwd: {
        oldPwd: true,
        newPwd: true,
        cmfPwd: true
      },

      oldPassword: '',
      newPassword: '',
      cmfPassword: '',
    })

    // 执行接口
    const submitFunc = () => {
      if (state.newPassword != state.cmfPassword) {
        NotifyNegative(t('twoPasswordsAreDifferent'))
        return false
      }
      const params = {
        type: 1,
        oldPassword: state.oldPassword,
        newPassword: state.newPassword,
      }
      updatePasswordAPI(params).then(() => {
        void $router.back()
      })
    }

    return {
      ...toRefs(state),
      submitFunc,
    }
  }
})
</script>

<style scoped></style>
