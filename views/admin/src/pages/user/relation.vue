<template>
  <div class="q-ma-md">
    <q-card flat bordered>
      <q-card-section>
        <q-input
          ref="filterRef"
          dense
          outlined
          v-model="filter"
          label="搜索网站用户名..."
        >
          <template v-slot:append>
            <q-icon
              v-if="filter !== ''"
              name="clear"
              class="cursor-pointer"
              @click="resetFilterFunc"
            />
          </template>
        </q-input>
      </q-card-section>
      <q-card-section>
        <q-tree
          :nodes="userTree"
          node-key="id"
          :filter="filter"
          :filter-method="filterMethodFunc"
          ref="treeRef"
        >
          <template v-slot:header-root="prop">
            <q-item class="no-padding">
              <q-item-section avatar>
                <q-img
                  no-spinner
                  :src="imageSrc(prop.node.avatar)"
                  width="58px"
                  height="58px"
                ></q-img>
              </q-item-section>
              <q-item-section>
                <div class="text-weight-bold text-body1">
                  {{ prop.node.username }}
                  <span class="text-body2 text-grey"
                    >【{{ prop.node.nickname }}】</span
                  >
                  <span class="text-body2 text-grey"
                    >【{{ prop.node.email }}】</span
                  >
                  <span class="text-body2 text-grey"
                    >【{{ prop.node.telephone }}】</span
                  >
                </div>
                <div
                  class="row q-gutter-xs items-center text-grey text-caption"
                >
                  <div>下级人数 ({{ prop.node.sumPeople }}) |</div>
                  <div>总充值 ({{ prop.node.sumDeposit.toFixed(2) }}) |</div>
                  <div>总提现 ({{ prop.node.sumWithdraw.toFixed(2) }}) |</div>
                  <div>总收益 ({{ prop.node.sumEarnings.toFixed(2) }})</div>
                </div>
                <div
                  class="row q-gutter-xs items-center text-grey text-caption"
                >
                  <div>
                    更新时间 ({{
                      date.formatDate(
                        prop.node.updatedAt * 1000,
                        'YYYY/MM/DD HH:mm:ss'
                      )
                    }}) --
                  </div>
                  <div>
                    注册时间 ({{
                      date.formatDate(
                        prop.node.createdAt * 1000,
                        'YYYY/MM/DD HH:mm:ss'
                      )
                    }})
                  </div>
                </div>
              </q-item-section>
            </q-item>
          </template>
          <template v-slot:header-generic="prop">
            <q-item class="no-padding">
              <q-item-section avatar>
                <q-img
                  no-spinner
                  :src="imageSrc(prop.node.avatar)"
                  width="50px"
                  height="50px"
                ></q-img>
              </q-item-section>
              <q-item-section>
                <div class="text-weight-bold text-body1">
                  {{ prop.node.username }}
                  <span class="text-body2 text-grey"
                    >【{{ prop.node.nickname }}】</span
                  >
                  <span class="text-body2 text-grey"
                    >【{{ prop.node.email }}】</span
                  >
                  <span class="text-body2 text-grey"
                    >【{{ prop.node.telephone }}】</span
                  >
                </div>
                <div
                  class="row q-gutter-xs items-center text-grey text-caption"
                >
                  <div>下级人数 ({{ prop.node.sumPeople }}) |</div>
                  <div>总充值 ({{ prop.node.sumDeposit.toFixed(2) }}) |</div>
                  <div>总提现 ({{ prop.node.sumWithdraw.toFixed(2) }}) |</div>
                  <div>总收益 ({{ prop.node.sumEarnings.toFixed(2) }})</div>
                </div>
                <div
                  class="row q-gutter-xs items-center text-grey text-caption"
                >
                  <div>
                    更新时间 ({{
                      date.formatDate(
                        prop.node.updatedAt * 1000,
                        'YYYY/MM/DD HH:mm:ss'
                      )
                    }}) --
                  </div>
                  <div>
                    注册时间 ({{
                      date.formatDate(
                        prop.node.createdAt * 1000,
                        'YYYY/MM/DD HH:mm:ss'
                      )
                    }})
                  </div>
                </div>
              </q-item-section>
            </q-item>
          </template>
        </q-tree>
      </q-card-section>
    </q-card>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, onMounted, ref } from 'vue';
import { api } from 'src/boot/axios';
import { imageSrc } from 'src/utils';
import { date } from 'quasar';

export default {
  name: 'UserRelation',
  setup() {
    const filterRef = ref(null) as any;
    const treeRef = ref(null) as any;
    const filter = ref('') as any;

    const state = reactive({
      userTree: [] as any,
      queryURL: '/user/relation',
    });

    // 初始化请求数据
    onMounted(() => {
      queryTreeDataFunc();
    });

    // 请求数据
    const queryTreeDataFunc = () => {
      void api.post(state.queryURL).then((res: any) => {
        state.userTree = res;
        setTimeout(() => {
          treeRef.value.expandAll();
        }, 100);
      });
    };

    // 重置过滤器
    const resetFilterFunc = () => {
      filter.value = '';
      filterRef.value.focus();
    };

    // 自定义过滤器
    const filterMethodFunc = (node: any, filter: any) => {
      const filterLower = filter.toLowerCase();
      return (
        node.username && node.username.toLowerCase().indexOf(filterLower) > -1
      );
    };

    return {
      filter,
      filterRef,
      treeRef,
      imageSrc,
      date,
      filterMethodFunc,
      resetFilterFunc,
      ...toRefs(state),
    };
  },
};
</script>
