<template>
  <div style="max-width: 280px" class="ellipsis">
    <div
      v-if="
        currentScope.col.type == ContentTypeList.InputEditText ||
        currentScope.col.type == ContentTypeList.InputEditNumber ||
        currentScope.col.type == ContentTypeList.InputEditTextArea
      "
    >
      {{ value === null || value === '' ? '- -' : value }}
      <q-popup-edit v-model="value">
        <template v-slot:default>
          <q-input
            v-model="value"
            dense
            autofocus
            counter
            v-if="currentScope.col.type === ContentTypeList.InputEditTextArea"
            type="textarea"
            @keyup.enter.stop
          />
          <q-input
            v-model="value"
            dense
            autofocus
            counter
            v-else-if="currentScope.col.type === ContentTypeList.InputEditText"
            @keyup.enter.stop
          />
          <q-input
            v-model.number="value"
            dense
            autofocus
            counter
            v-else-if="
              currentScope.col.type === ContentTypeList.InputEditNumber
            "
            type="number"
            @keyup.enter.stop
          />
          <div class="row justify-end">
            <div>
              <q-btn label="取消" color="primary" v-close-popup flat></q-btn>
            </div>
            <div>
              <q-btn
                label="确定"
                color="primary"
                v-close-popup
                flat
                @click="updatePopupEditFunc()"
              ></q-btn>
            </div>
          </div>
        </template>
      </q-popup-edit>
    </div>

    <!-- 开关按钮 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.InputEditToggle">
      <q-toggle
        v-model="value"
        :true-value="parseInt(currentScope.col.data[0].value)"
        :false-value="parseInt(currentScope.col.data[1].value)"
        @update:model-value="updatePopupEditFunc()"
      />
    </div>

    <!-- 显示图片 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Image">
      <q-img
        :src="imageSrc(value)"
        loading="lazy"
        spinner-color="white"
        @click="
          imageList = [value];
          imageDialog = true;
        "
        style="max-height: 50px; max-width: 50px"
      ></q-img>
    </div>

    <!-- 显示图片组 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Images">
      <q-img
        :src="imageSrc(JSON.parse(value)[0])"
        v-if="value !== ''"
        loading="lazy"
        spinner-color="white"
        @click="
          imageList = value;
          imageDialog = true;
        "
        style="max-height: 50px; max-width: 50px"
      >
        <q-popup-proxy transition-show="flip-up" transition-hide="flip-down">
          <q-card style="min-width: 30%">
            <q-card-section
              v-for="(image, imageIndex) in JSON.parse(value)"
              :key="imageIndex"
            >
              <q-img no-spinner :src="imageSrc(image)"></q-img>
            </q-card-section>
          </q-card>
        </q-popup-proxy>
      </q-img>
    </div>

    <!-- 显示时间 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.DatePicker">
      {{ date.formatDate(parseInt(value) * 1000, 'YYYY-MM-DD HH:mm:ss') }}
    </div>

    <!-- 显示选择框 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Select">
      <div
        v-for="(selected, selectedIndex) in currentScope.col.data"
        :key="selectedIndex"
      >
        <q-badge
          outline
          :color="quasarColorsList[selectedIndex % 25]"
          v-if="parseInt(selected.value) === parseInt(value)"
          :label="selected.label"
        />
      </div>
    </div>

    <!-- 多语言类型 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Translate">
      <q-btn
        flat
        label="点击修改"
        size="sm"
        color="primary"
        @click="editTranslateFunc"
      ></q-btn>
      <span class="text-caption">[{{ value }}]</span>
    </div>

    <!-- 正常显示文本 -->
    <div v-else class="ellipsis">
      {{ typeof value == 'number' ? value : value == '' ? '- -' : value }}
      <q-popup-proxy v-if="value">
        <q-card>
          <q-card-section>
            {{ value }}
          </q-card-section>
        </q-card>
      </q-popup-proxy>
    </div>

    <!-- 显示图dialog地方 -->
    <div
      v-if="
        currentScope.col.type == ContentTypeList.Image ||
        currentScope.col.type == ContentTypeList.Images
      "
    >
      <q-dialog v-model="imageDialog" full-width>
        <q-card class="q-pa-sm">
          <q-carousel
            swipeable
            animated
            v-model="imageSlide"
            thumbnails
            infinite
          >
            <q-carousel-slide
              :name="imageIndex"
              :img-src="imageSrc(image)"
              v-for="(image, imageIndex) in imageList"
              :key="imageIndex"
            />
          </q-carousel>
        </q-card>
      </q-dialog>
    </div>

    <q-dialog v-model="translateDialog">
      <q-card class="full-width">
        <q-card-section>
          <div class="text-h6">多语言翻译编辑</div>
        </q-card-section>

        <q-card-section>
          <div>
            <q-tabs
              v-model="translateTab"
              dense
              class="text-grey"
              active-color="primary"
              indicator-color="primary"
              align="justify"
              narrow-indicator
            >
              <q-tab
                :name="translate.id"
                :label="translate.name"
                v-for="(translate, translateIndex) in translateValues"
                :key="translateIndex"
              />
            </q-tabs>

            <q-separator />

            <q-tab-panels v-model="translateTab" animated>
              <q-tab-panel
                :name="translate.id"
                v-for="(translate, translateIndex) in translateValues"
                :key="translateIndex"
              >
                <input-editor-component
                  v-model="translate.value"
                  :label="translate.name"
                ></input-editor-component>
              </q-tab-panel>
            </q-tab-panels>
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="取消" color="grey" v-close-popup />
          <q-btn
            flat
            label="Google翻译"
            color="positive"
            @click="googleTranslate"
          />
          <q-btn flat label="保存" color="primary" @click="saveTranslateFunc" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { imageSrc } from 'src/utils';
import { ContentTypeList, quasarColorsObject } from 'src/utils/define';
import { api } from 'src/boot/axios';
import { date } from 'quasar';
import InputEditorComponent from 'src/components/editor.vue';

export default {
  name: 'ColumnsContentComponent',
  components: { InputEditorComponent },
  props: {
    url: { type: String, default: '' },
    Primary: { type: String, default: 'id' },
    modelValue: { type: undefined, default: '' },
    currentScope: {
      type: Object,
      default: () => {
        return {};
      },
    },
  },
  emits: ['done'],
  setup(props: any, context: any) {
    const state = reactive({
      quasarColorsList: Object.values(quasarColorsObject),
      imageList: [] as any,
      imageSlide: 0,
      imageDialog: false,
      value: props.modelValue as any,
      translateTab: '',
      translateDialog: false,
      translateValues: [] as any,
    });

    //  编辑方法提交
    const updatePopupEditFunc = () => {
      const params = {} as any;
      params[props.Primary] = props.currentScope.row[props.Primary];
      params[props.currentScope.col.field] = state.value;

      api.post(props.url, params).then(() => {
        context.emit('done', props.currentScope);
      });
    };

    // editTranslateFunc 编辑多语言
    const editTranslateFunc = () => {
      //  请求接口
      api
        .post('/system/translate/fields', {
          field: state.value,
          adminId: props.currentScope.row.adminId,
        })
        .then((res: any) => {
          state.translateValues = res;
          state.translateTab = state.translateValues[0].id;
        });

      state.translateDialog = true;
    };

    // saveTranslateFunc 保存多语言内容
    const saveTranslateFunc = () => {
      api
        .post('/system/translate/update/fields', {
          data: state.translateValues,
          field: state.value,
          adminId: props.currentScope.row.adminId,
        })
        .then(() => {
          context.emit('done', props.currentScope);
          state.translateDialog = false;
        });
    };

    // googleTranslate google翻译内容
    const googleTranslate = () => {
      const currentIndex = currentTabIndex();
      if (currentIndex > -1) {
        const currentTranslate = state.translateValues[currentIndex];
        api
          .post('/system/translate/google', {
            field: state.value,
            alias: currentTranslate.alias,
            adminId: props.currentScope.row.adminId,
          })
          .then((res: any) => {
            currentTranslate.value = res;
          });
      }
    };

    const currentTabIndex = () => {
      for (let i = 0; i < state.translateValues.length; i++) {
        if (state.translateValues[i].id == state.translateTab) {
          return i;
        }
      }
      return -1;
    };

    return {
      date,
      imageSrc,
      ...toRefs(state),
      updatePopupEditFunc,
      saveTranslateFunc,
      editTranslateFunc,
      googleTranslate,
      ContentTypeList,
    };
  },
};
</script>
