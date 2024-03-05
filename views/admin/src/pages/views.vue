<template>
  <div class="q-pa-md">
    <!-- 头部数据查询 -->
    <div class="row q-gutter-sm items-center">
      <div
        v-for="(input, inputIndex) in config.search.inputList"
        :key="inputIndex"
      >
        <InputLayoutsComponent
          v-model="config.search.params[input.field]"
          :label="input.label"
          :type="input.type"
          :options="input.data"
        ></InputLayoutsComponent>
      </div>
      <div v-if="config.search.inputList.length > 0" style="margin-top: 0">
        <q-btn
          icon="search"
          color="primary"
          @click="requestTableFunc({ pagination: config.pagination })"
        ></q-btn>
      </div>
    </div>

    <!-- 数据表格 -->
    <q-card flat bordered class="q-mt-md">
      <!-- 工具栏 -->
      <q-card-section>
        <div class="row">
          <div
            v-for="(tool, toolIndex) in config.table.tools"
            :key="toolIndex"
            class="q-mr-sm q-mb-sm"
          >
            <q-btn
              :label="tool.label"
              :color="tool.color"
              :size="tool.size"
              v-if="initStore.hasRoute(tool.config.url)"
              @click="dialogOpenFunc(tool.config, null)"
            ></q-btn>
          </div>
        </div>
      </q-card-section>
      <q-card-section>
        <q-table
          flat
          :rows="rows"
          :columns="config.table.columns"
          :row-key="config.table.key"
          selection="multiple"
          @request="requestTableFunc"
          v-model:selected="checkboxList"
          v-model:pagination="config.pagination"
        >
          <template v-slot:top>
            <q-space />
            <q-btn
              color="secondary"
              icon-right="archive"
              label="下载CSV文件"
              no-caps
              @click="exportFileFunc"
            />
          </template>
          <template v-slot:body-cell="scope">
            <q-td>
              <ColumnsContent
                v-if="scope.col.type != 'options'"
                :url="config.table.updateUrl"
                :primary="config.table.key"
                v-model="scope.row[scope.col.field]"
                @done="dialogDoneFunc"
                :current-scope="scope"
              ></ColumnsContent>

              <!-- 额外的操作 -->
              <div v-if="scope.col.type == 'options'">
                <div class="row">
                  <div
                    v-for="(option, optionIndex) in config.table.options"
                    :key="optionIndex"
                    class="q-mr-sm q-mb-sm"
                  >
                    <q-btn
                      :label="option.label"
                      :color="option.color"
                      :size="option.size"
                      v-if="
                        initStore.hasRoute(option.config.url) &&
                        (option.eval == '' ||
                          initStore.executeEval(option.eval, {
                            scope: scope,
                          }))
                      "
                      @click="dialogOpenFunc(option.config, scope)"
                    ></q-btn>
                  </div>
                </div>
              </div>
            </q-td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>
  </div>

  <DialogComponent ref="domDialogRef" @done="dialogDoneFunc"></DialogComponent>
</template>

<script lang="ts">
import { reactive, toRefs, ref } from 'vue';
import { api } from 'src/boot/axios';
import { exportCSVFile } from 'src/utils/export';
import InputLayoutsComponent from 'src/components/inputLayouts.vue';
import DialogComponent from 'src/components/dialog.vue';
import ColumnsContent from 'src/components/columnsContent.vue';
import { useRouter } from 'vue-router';
import { useInitStore } from 'src/stores/init';
import { inputOperateSettingFun } from 'src/utils/input';

export default {
  name: 'IndexViews',
  components: { InputLayoutsComponent, DialogComponent, ColumnsContent },
  setup() {
    const initStore = useInitStore();
    const domDialogRef = ref(null) as any;
    const $router = useRouter();
    const state = reactive({
      index: 0,
      config: {
        // 请求数据地址
        url: '',

        //  排序规则
        pagination: {
          sortBy: 'id',
          descending: true,
          page: 0,
          rowsPerPage: 20,
        } as any,

        //  查询条件
        search: {
          params: {} as any,
          inputList: [] as any,
        },

        //  数据表格
        table: {
          key: 'id',
          updateUrl: '',
          tools: [] as any,
          columns: [] as any,
          options: [] as any,
        },
      },

      //  当前表格数据
      rows: [] as any,
      //  选中框
      checkboxList: [] as any,
    });

    //  请求配置文件
    api
      .post(<string>$router.currentRoute.value.meta.views, {}, {
        showLoading: false,
      } as any)
      .then((config: any) => {
        state.config = config;

        if (state.config.table.options.length > 0) {
          state.config.table.columns.push({
            label: '操作栏',
            field: '_options',
            type: 'options',
            name: '_options',
            align: 'left',
          });
        }

        //  请求数据
        requestTableFunc({ pagination: state.config.pagination });
      });

    //  请求表格数据
    const requestTableFunc = (props: { pagination: any }) => {
      state.rows = [];
      state.config.pagination = props.pagination;
      state.config.search.params['pagination'] = state.config.pagination;
      api
        .post(state.config.url, state.config.search.params)
        .then((res: any) => {
          state.rows = res.items;
          state.config.pagination.rowsNumber = res.count;
        });
    };

    //  弹窗操作
    const dialogOpenFunc = (config: any, rawData: any) => {
      const currentConfig = JSON.parse(JSON.stringify(config));

      //  判断是否需要获取checkbox数据
      if (config.params.hasOwnProperty('operate')) {
        switch (config.params['operate']) {
          //  checkbox 操作
          case 'checkbox':
            if (state.checkboxList.length == 0) {
              return;
            }

            //  生成弹窗名称内容
            const content = [];
            const scanList = [];
            for (let i = 0; i < state.checkboxList.length; i++) {
              content.push(state.checkboxList[i][config.params['name']]);
              scanList.push(state.checkboxList[i][config.params['scan']]);
            }

            currentConfig.content = content.toString();
            currentConfig.params = {};
            currentConfig.params[config.params['field']] = scanList;
            break;

          //  设置操作
          case 'setting':
            const inputOperateInfo = inputOperateSettingFun(
              config.params['operate'],
              state.config,
              currentConfig,
              rawData
            );

            currentConfig.params = inputOperateInfo.params;
            currentConfig.inputList = inputOperateInfo.inputList;
            break;
          default:
            break;
        }
      } else {
        //  如果有原数据, 那么赋值操作
        if (rawData != null) {
          currentConfig.params[state.config.table.key] =
            rawData.row[state.config.table.key];
          for (const key in currentConfig.params) {
            if (typeof currentConfig.params[key] == 'object') {
              for (const children in currentConfig.params[key]) {
                if (
                  currentConfig.params[key][children] == null ||
                  currentConfig.params[key][children] == '' ||
                  currentConfig.params[key][children] == 0
                ) {
                  currentConfig.params[key][children] =
                    rawData.row[key][children];
                }
              }
            } else {
              if (
                currentConfig.params[key] == null ||
                currentConfig.params[key] == '' ||
                currentConfig.params[key] == 0
              ) {
                currentConfig.params[key] = rawData.row[key];
              }
            }
          }
        }
      }

      domDialogRef.value.setConfig(currentConfig);
      domDialogRef.value.open();
    };

    //  完成操作
    const dialogDoneFunc = () => {
      state.checkboxList = [];
      requestTableFunc({ pagination: state.config.pagination });
    };

    // 下载csv文件
    const exportFileFunc = () => {
      exportCSVFile(state.config.table.columns, state.rows);
    };

    return {
      ...toRefs(state),
      requestTableFunc,
      domDialogRef,
      dialogOpenFunc,
      dialogDoneFunc,
      exportFileFunc,
      initStore,
    };
  },
};
</script>
