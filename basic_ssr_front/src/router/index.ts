import { route } from 'quasar/wrappers';
import { UserTokenKey, InitStoreState, UserLangKey } from 'src/stores/init';
import { Cookies, Platform } from 'quasar';
import { dynamicRouterFunc } from 'src/router/routes';
import { templateRoutes } from 'src/router/routes';
import { initAPI } from 'src/apis';

import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router';


// 路由接口
export interface TemplateRouteInterface {
  route: string; //  路由
  name: string; //  名称
  componentMobile: string; //  手机端文件
  componentDesktop: string; //  桌面端文件
  meta: any; //  meta数据
  children: TemplateRouteInterface[]; //  子级
}

import routes from 'src/router/routes';
/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(async function ({ store, ssrContext }) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === 'history'
      ? createWebHistory
      : createWebHashHistory;

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,
    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.VUE_ROUTER_BASE),
  });

  // 请求初始化数据
  const $platform = process.env.SERVER ? Platform.parseSSR(ssrContext) : Platform;

  if (process.env.SERVER) {
    const $cookies = Cookies.parseSSR(ssrContext);
    store.state.value['init'] = JSON.parse(JSON.stringify(InitStoreState));
    store.state.value.init.userToken = <string>$cookies.get(UserTokenKey) ?? '';
    store.state.value.init.userLang = <string>$cookies.get(UserLangKey) ?? 'zh-CN';

    //获取初始化数据
    await initAPI({ domain: ssrContext?.req.get('host'), lang: store.state.value.init.userLang }).then((res: any) => {
      //  初始化配置
      store.state.value.init.config = res.config

      // 初始化TabBar菜单
      store.state.value.init.tabBars = res.tabBars

      // 初始化用户菜单
      store.state.value.init.userMenu = res.userMenu

      // 初始化首页菜单
      store.state.value.init.homeMenu = res.homeMenu

      // 初始化快捷菜单
      store.state.value.init.quickMenu = res.quickMenu

      // 初始化翻译数据
      store.state.value.init.translate = res.translate

      // 初始化国家列表
      store.state.value.init.countryList = res.countryList

      // 初始化语言列表
      store.state.value.init.languageList = res.languageList

    });
  }

  //  动态加载路由
  dynamicRouterFunc(
    Router,
    templateRoutes.get(store.state.value.init.config.template),
    <boolean>$platform.is.mobile
  );

  //  路由守卫
  Router.beforeEach((to, form, next) => {
    if (
      (to.name === 'UserLogin' || to.name === 'UserRegister') &&
      store.state.value.init.userToken &&
      store.state.value.init.userToken.length > 0
    ) {
      next('/');
    } else {
      // 验证是否跳转到登录页面
      if (
        to.matched.some((record) => record.meta.requireAuth) &&
        (store.state.value.init.userToken == null ||
          store.state.value.init.userToken.length === 0)
      ) {
        next('/login');
      } else {
        next();
      }
    }
  });

  return Router;
});
