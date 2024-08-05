<template>
  <div style="padding: 30px" class="text-white column items-center">
    <div class="row full-width" style="max-width: 1200px; padding: 60px 0">
      <div class="col-9">
        <div class="row justify-center">
          <div v-for="(footer, footerIndex) in footerList" :key="footerIndex" class="col-3">
            <div class="column">
              <div class="text-body1 text-bold">{{ footer.name }}</div>
              <div v-for="(children, childrenIndex) in footer.children" :key="childrenIndex" style="max-width: 100%;">
                <div class="row justify-start items-center q-mt-sm">
                  <div v-if="children.icon != ''" class="q-mr-xs">
                    <q-icon :name="children.icon" color="grey-4" size="20px"></q-icon>
                  </div>
                  <div class="ellipsis text-grey-4 text-body2 col cursor-pointer" @click="routerTo(children.link)">
                    {{ children.name }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="col-1"></div>
      <div class="col-2">
        <div class="text-body1 text-bold">{{ socialInfo.name }}</div>
        <div class="row items-center q-gutter-sm q-mt-xs">
          <div v-for="(social, socialIndex) in socialInfo.children" :key="socialIndex">
            <q-avatar color="white" rounded size="32px" @click="routerTo(social.link)">
              <q-img no-spinner :src="imageSrc(social.icon)" width="32px" height="32px" class="cursor-pointer"></q-img>
            </q-avatar>
          </div>
        </div>
      </div>
    </div>

    <div class="q-mt-xl text-center">
      <div class="text-grey">
        {{ config.name }}
        <q-icon name="copyright" size="xs"></q-icon>
        <span>2023</span>
        <span>Cookie settings</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { onMounted, reactive, toRefs } from 'vue'
import { InitStore } from 'src/stores/init';
import { imageSrc } from 'src/utils'
import { footerInfoAPI } from 'src/apis'
import { useRouter } from 'vue-router';

export default {
  name: 'LayoutsFooter',
  setup() {
    const $initStore = InitStore()
    const $router = useRouter()

    const state = reactive({
      config: $initStore.config,
      footerList: [] as any,
      socialInfo: { name: '', children: [] } as any,
    })

    onMounted(() => {
      footerInfoAPI().then((res: any) => {
        state.footerList = res.items
        state.socialInfo = res.socialInfo
      })
    })

    // 跳转路由
    const routerTo = (link: string) => {
      if (link == '') {
        return
      }

      if (link.indexOf('/') == 0) {
        $router.push(link)
      } else {
        window.location.href = link
      }
    }

    return {
      imageSrc,
      routerTo,
      ...toRefs(state)
    };
  },
};
</script>
