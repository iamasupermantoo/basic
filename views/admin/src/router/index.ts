import { route } from 'quasar/wrappers';
import { UserTokenKey, InitStoreState } from 'src/stores/init';
import { Cookies } from 'quasar';
import { api } from 'src/boot/axios';
import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router';

export interface MenuListDataInterface {
  icon: string;
  vueTmp: string;
  viewsUrl: string;
}

export interface MenuListInterface {
  name: string;
  route: string;
  children: MenuListInterface[];
  data: MenuListDataInterface;
}

import routes from 'src/router/routes';
const routePageList = import.meta.glob('../pages/**/*.vue');

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

  //  请求初始化数据
  const $cookies = ssrContext ? Cookies.parseSSR(ssrContext) : Cookies;
  store.state.value['init'] = JSON.parse(JSON.stringify(InitStoreState));
  store.state.value.init.userToken = <string>$cookies.get(UserTokenKey);
  if (store.state.value.init.userToken != '') {
    await api
      .get('/init', {
        headers: {
          Authorization: 'Bearer ' + store.state.value.init.userToken,
        },
      })
      .then((res: any) => {
        store.state.value.init.config = res.config;
        store.state.value.init.routerList = res.routerList;
        store.state.value.init.menuList = res.menuList;
        store.state.value.init.userInfo = res.adminInfo;
        dynamicLoadRouter(Router, store.state.value.init.menuList);
      })
      .catch(() => {
        Router.push({ name: 'Login' });
      });
  }

  // 路由前置守卫
  Router.beforeEach((to, form, next) => {
    if (
      (to.name === 'Login' || to.name === 'Register') &&
      store.state.value.init.userToken != null &&
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

// 动态添加路由
export const dynamicLoadRouter = (
  router: any,
  routerList: MenuListInterface[]
) => {
  dynamicRouter(router, routerList);
};

// 递归动态添加路由
const dynamicRouter = (router: any, routerList: MenuListInterface[]) => {
  for (let index = 0; index < routerList.length; index++) {
    const element = routerList[index];
    if (element.route !== '' && !router.hasRoute(element.route)) {
      router.addRoute('Layouts', {
        path: element.route,
        component: routePageList['../pages' + element.data.vueTmp + '.vue'],
        meta: {
          requireAuth: true,
          keepAlive: true,
          views: element.data.viewsUrl,
        },
      });
    }
    if (
      element.hasOwnProperty('children') &&
      element.children !== null &&
      element.children.length > 0
    ) {
      dynamicRouter(router, element.children);
    }
  }

  //  覆盖当前路由
  router.replace(router.currentRoute.value.fullPath);
};
