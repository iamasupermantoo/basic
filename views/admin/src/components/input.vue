<template>
  <div class="q-mb-sm">
    <!-- 文本域 -->
    <div v-if="type == InputTypeList.TextArea">
      <q-input
        dense
        outlined
        rows="5"
        :readonly="readonly"
        v-model="value"
        :label="label"
        @update:model-value="$emit('refresh', value, field)"
        type="textarea"
      ></q-input>
    </div>

    <!-- 富文本编辑器 -->
    <div v-else-if="type == InputTypeList.Editor">
      <input-editor-component
        v-model="value"
        :label="label"
        @update:model-value="$emit('refresh', value, field)"
        :readonly="readonly"
      ></input-editor-component>
    </div>

    <!-- 选择框 -->
    <div v-else-if="type == InputTypeList.Select && options.length >= 1">
      <q-select
        dense
        outlined
        v-model="value"
        :label="label"
        :readonly="readonly"
        @update:model-value="$emit('refresh', value, field)"
        :options="data"
        emit-value
        map-options
      ></q-select>
    </div>

    <!-- 单选框 -->
    <div v-else-if="type == InputTypeList.Radio">
      <div class="text-body2 text-bold text-grey">{{ label }}</div>
      <div class="q-gutter-sm">
        <q-radio
          v-model="value"
          :val="radio.value"
          :label="radio.label"
          v-for="(radio, radioIndex) in options"
          :key="radioIndex"
        />
      </div>
    </div>

    <!-- 多选框 -->
    <div v-else-if="type == InputTypeList.Checkbox">
      <div class="text-body2 text-bold text-grey">{{ label }}</div>
      <q-checkbox
        v-for="(checkbox, checkboxIndex) in options"
        :readonly="readonly"
        v-model="value[checkbox.value]"
        :label="checkbox.label"
        class="no-margin"
        @update:model-value="$emit('refresh', value, field)"
        :key="checkboxIndex"
      >
      </q-checkbox>
    </div>

    <!-- 开关类型 -->
    <div v-else-if="type == InputTypeList.Toggle">
      <div>
        <div>
          <q-toggle
            v-model="value"
            :label="label"
            :readonly="readonly"
            @update:model-value="$emit('refresh', value, field)"
            :true-value="options[0].value"
            :false-value="options[1].value"
          ></q-toggle>
        </div>
      </div>
    </div>

    <!-- 时间类型 -->
    <div v-else-if="type == InputTypeList.DatePicker">
      <q-input
        dense
        outlined
        :label="label"
        :readonly="readonly"
        :model-value="
          typeof value === 'number'
            ? date.formatDate(value * 1000, 'YYYY/MM/DD')
            : value
        "
      >
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy
              ref="qDateProxy"
              cover
              transition-show="scale"
              transition-hide="scale"
            >
              <q-date
                v-model="value"
                @update:model-value="$emit('refresh', value, field)"
              >
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="关闭" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
    </div>

    <!-- 时间范围类型 -->
    <div v-else-if="type == InputTypeList.RangeDatePicker">
      <q-input
        dense
        outlined
        :label="label"
        :readonly="readonly"
        :model-value="
          value !== '' && value !== undefined && value !== null
            ? value.from + ' - ' + value.to
            : ''
        "
      >
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy
              ref="qDateProxy"
              cover
              transition-show="scale"
              transition-hide="scale"
            >
              <q-date
                range
                v-model="value"
                @update:model-value="$emit('refresh', value, field)"
              >
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="关闭" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
    </div>

    <!-- 文件上传 -->
    <div v-else-if="type == InputTypeList.File || type == InputTypeList.Image">
      <uploader-component
        :label="label"
        v-model="value"
        @uploaded="updateValueFunc"
      ></uploader-component>
    </div>

    <!-- 多文件, 多图上传 -->
    <div v-else-if="type == InputTypeList.Images">
      <uploader-component
        :label="label"
        v-model="value"
        :multiple="true"
        @uploaded="updateValueFunc"
      ></uploader-component>
    </div>

    <!-- 数字input -->
    <div v-else-if="type == InputTypeList.Number">
      <q-input
        dense
        outlined
        v-model.number="value"
        :label="label"
        type="number"
        :readonly="readonly"
        @change="$emit('refresh', value, field)"
      ></q-input>
    </div>

    <!-- 密码格式 -->
    <div v-else-if="type == InputTypeList.Password">
      <q-input
        dense
        outlined
        v-model="value"
        :label="label"
        type="password"
        :readonly="readonly"
        @change="$emit('refresh', value, field)"
      ></q-input>
    </div>

    <!-- 文本 -->
    <div v-else>
      <div v-if="type == InputTypeList.InputTranslate" class="text-caption text-grey">
        这边设置翻译键值, 设置具体值 系统配置 > 翻译配置
      </div>
      <q-input
        dense
        outlined
        v-model="value"
        :label="label"
        type="text"
        :readonly="readonly"
        @change="$emit('refresh', value, field)"
      ></q-input>
    </div>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs } from 'vue';
import { date } from 'quasar';
import InputEditorComponent from 'src/components/editor.vue';
import UploaderComponent from 'src/components/uploader.vue';
import { InputTypeList } from 'src/utils/define';

export default {
  name: 'InputComponent',
  components: {
    InputEditorComponent,
    UploaderComponent,
  },
  emits: ['refresh'],
  props: {
    type: { type: Number, default: InputTypeList.Text },
    field: { type: String, default: ''},
    label: { type: String, default: '' },
    modelValue: { type: undefined, default: '' },
    options: { type: null, default: [] },
    readonly: { type: Boolean, default: false },
  },
  setup(props: any, context: any) {
    const state = reactive({
      value: props.modelValue,
      data: props.options
    });

    //  复选框 - 默认值设置
    if (props.type == InputTypeList.Checkbox) {
      if (state.value == null || state.value == '') {
        state.value = {}
      }

      //  初始化设置
      for (let i = 0; i < props.options.length; i++) {
        if (!state.value.hasOwnProperty(props.options[i].value)) {
          state.value[props.options[i].value] = false
        }
      }
    }

    //  默认值设置 - 选择框设置
    if (props.type == InputTypeList.Select && props.options.length > 1) {
      if (typeof props.options[0].value == 'number') {
        state.value =
          !state.value || state.value == '' ? 0 : parseFloat(state.value);
      }

      //  是否需要自动添加全部选项
      state.data = [{label: '全部', value: typeof props.options[0].value == 'number' ? 0 : ''}].concat(props.options)
      for (let i = 0; i < props.options.length; i++) {
        if (props.options[i].value == '' || props.options[i].value == 0) {
          //  如果有设置 == '' 或者设置了 == 0 那么使用自带的options
          state.data = props.options
        }
      }
    }

    //  回调更新数据
    const updateValueFunc = (val: any) => {
      state.value = val;
      context.emit('refresh', val, props.field);
    };

    return {
      date,
      ...toRefs(state),
      InputTypeList,
      updateValueFunc,
    };
  },
};
</script>
