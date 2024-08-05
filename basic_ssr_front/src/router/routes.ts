import { RouteRecordRaw, Router } from 'vue-router';
import { TemplateRouteInterface } from 'src/router';
import { defaultRouter, DefaultTemplateName } from 'src/router/defaultRouter';

export const templateRoutes: any = new Map([[DefaultTemplateName, defaultRouter]]);

// componentPathList 获取文件路径文件
const componentPathList = Object.assign(
  import.meta.glob('src/layouts/**/*.vue'),
  import.meta.glob('src/pages/**/*.vue')
);
// 动态加载路由
export const dynamicRouterFunc = (
  router: Router, //  路由对象
  routerList: TemplateRouteInterface[], //  载入的路由
  isMobile: boolean, //  是否手机端
  parent = '', //  父级路由
) => {
  if (routerList && routerList.length > 0) {
    routerList.forEach((item) => {
      //  动态添加路由
      router.addRoute(parent, {
        path: item.route,
        name: item.name,
        component:
          componentPathList[
          isMobile ? item.componentMobile : item.componentDesktop
          ],
        meta: item.meta,
      });

      //  是否需要递归子级, 动态添加路由
      if (
        item.hasOwnProperty('children') &&
        item.children !== null &&
        item.children.length > 0
      ) {
        dynamicRouterFunc(router, item.children, isMobile, item.name);
      }
    });
  }
};

const routes: RouteRecordRaw[] = [
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/404.vue'),
  },
];

export default routes;
