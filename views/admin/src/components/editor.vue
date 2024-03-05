<template>
  <div>
    <q-editor
      min-height="5rem"
      :readonly="readonly"
      ref="editorRef"
      :placeholder="label"
      :toolbar="[
        ['token'],
        [
          {
            label: $q.lang.editor.align,
            icon: $q.iconSet.editor.align,
            fixedLabel: true,
            options: ['left', 'center', 'right', 'justify'],
          },
        ],
        ['bold', 'italic', 'strike', 'underline', 'subscript', 'superscript'],
        ['hr', 'link', 'custom_btn'],
        [
          {
            label: $q.lang.editor.fontSize,
            icon: $q.iconSet.editor.fontSize,
            fixedLabel: true,
            fixedIcon: true,
            list: 'no-icons',
            options: [
              'size-1',
              'size-2',
              'size-3',
              'size-4',
              'size-5',
              'size-6',
              'size-7',
            ],
          },
          {
            label: $q.lang.editor.defaultFont,
            icon: $q.iconSet.editor.font,
            fixedIcon: true,
            list: 'no-icons',
            options: [
              'default_font',
              'arial',
              'arial_black',
              'comic_sans',
              'courier_new',
              'impact',
              'lucida_grande',
              'times_new_roman',
              'verdana',
            ],
          },
          'removeFormat',
        ],
        ['viewsource'],
      ]"
      :fonts="{
        arial: 'Arial',
        arial_black: 'Arial Black',
        comic_sans: 'Comic Sans MS',
        courier_new: 'Courier New',
        impact: 'Impact',
        lucida_grande: 'Lucida Grande',
        times_new_roman: 'Times New Roman',
        verdana: 'Verdana',
      }"
      v-model="value"
      @update:model-value="$emit('update:modelValue', value)"
      :dense="$q.screen.lt.md"
    >
      <template v-slot:token>
        <q-btn-dropdown
          dense
          no-caps
          ref="editorTokenRef"
          no-wrap
          unelevated
          color="white"
          text-color="primary"
          label="图片｜颜色"
          size="sm"
        >
          <q-list dense>
            <q-item class="q-mb-md">
              <q-uploader
                flat
                auto-upload
                style="height: 36px"
                @uploaded="editorUploadedEventFunc"
                :headers="[
                  { name: 'Authorization', value: 'Bearer ' + userToken },
                ]"
                :url="baseURL + '/upload'"
                field-name="file"
              >
                <template v-slot:header></template>
                <template v-slot:list="scope">
                  <div @click="scope.pickFiles">
                    <q-uploader-add-trigger />
                    <div
                      class="text-body2 text-grey"
                      style="height: 36px; line-height: 36px"
                    >
                      选择需要上传图片...
                    </div>
                  </div>
                </template>
              </q-uploader>
            </q-item>
            <q-item
              tag="label"
              clickable
              @click="
                editorEditTextColorFunc('backColor', editorBackgroundColor)
              "
            >
              <q-item-section side>
                <q-icon name="highlight" />
              </q-item-section>
              <q-item-section>
                <q-color
                  v-model="editorBackgroundColor"
                  default-view="palette"
                  no-header
                  no-footer
                  :palette="[
                    '#ffccccaa',
                    '#ffe6ccaa',
                    '#ffffccaa',
                    '#ccffccaa',
                    '#ccffe6aa',
                    '#ccffffaa',
                    '#cce6ffaa',
                    '#ccccffaa',
                    '#e6ccffaa',
                    '#ffccffaa',
                    '#ff0000aa',
                    '#ff8000aa',
                    '#ffff00aa',
                    '#00ff00aa',
                    '#00ff80aa',
                    '#00ffffaa',
                    '#0080ffaa',
                    '#0000ffaa',
                    '#8000ffaa',
                    '#ff00ffaa',
                    '#ff00ffaa',
                  ]"
                />
              </q-item-section>
            </q-item>

            <q-item
              tag="label"
              clickable
              @click="editorEditTextColorFunc('foreColor', editorTextColor)"
            >
              <q-item-section side>
                <q-icon name="format_paint" />
              </q-item-section>
              <q-item-section>
                <q-color
                  v-model="editorTextColor"
                  default-view="palette"
                  no-header
                  no-footer
                  :palette="[
                    '#000000',
                    '#ff0000',
                    '#ff8000',
                    '#ffff00',
                    '#00ff00',
                    '#00ffff',
                    '#0080ff',
                    '#0000ff',
                    '#8000ff',
                    '#ff00ff',
                  ]"
                />
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>
      </template>
    </q-editor>
  </div>
</template>

<script lang="ts">
import { reactive, toRefs, ref, watch } from 'vue';
import { imageSrc } from 'src/utils';
import { Loading } from 'quasar';
import { useInitStore } from 'src/stores/init';

export default {
  name: 'EditorConent',
  props: {
    modelValue: { type: undefined, default: '' },
    label: { type: String, default: '' },
    readonly: { type: Boolean, default: false },
  },
  setup(props: any, context: any) {
    const $userStore = useInitStore();
    const editorRef = ref(null) as any;
    const editorTokenRef = ref(null) as any;
    const state = reactive({
      value: props.modelValue,
      userToken: $userStore.userToken,
      baseURL: process.env.baseURL,
      editorTextColor: '',
      editorBackgroundColor: '',
    });

    // 编辑文本颜色
    const editorEditTextColorFunc = (cmd: string, color: string) => {
      editorTokenRef.value.hide();
      editorRef.value.runCmd(cmd, color);
      editorRef.value.focus();
    };

    // 富文本编辑器上传图片
    const editorUploadedEventFunc = (info: any) => {
      const fileURL = JSON.parse(info.xhr.response).data[0];
      state.value += '<img src="' + imageSrc(fileURL) + '" />';
      editorTokenRef.value.hide();
      Loading.hide();
      context.emit('update:modelValue', state.value);
    };

    watch(
      () => props.modelValue,
      (value: any) => {
        state.value = value;
      }
    );

    return {
      ...toRefs(state),
      imageSrc,
      editorRef,
      editorTokenRef,
      editorEditTextColorFunc,
      editorUploadedEventFunc,
    };
  },
};
</script>
