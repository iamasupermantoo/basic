import { boot } from 'quasar/wrappers';
import { createI18n } from 'vue-i18n';

// 解析配置语言
const parsePrefetchLocales = (locales: any, lang: string): any => {
  const messages = {} as any;
  messages[lang] = {} as any;
  for (let i = 0; i < locales.length; i++) {
    messages[lang][locales[i].label] = locales[i].value;
  }
  return messages;
};

// 设置当前语言
export const setLanguageFunc = (i18n: any, locales: any, lang: string) => {
  const messagesList = parsePrefetchLocales(locales, lang);
  i18n.setLocaleMessage(lang, messagesList[lang]);
  i18n.locale.value = lang;
};

export default boot(({ app }) => {
  const i18n = createI18n({
    legacy: false,
  });

  // Set i18n instance on app
  app.use(i18n);
});
