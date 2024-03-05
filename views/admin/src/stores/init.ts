import { defineStore } from 'pinia';
import { Cookies } from 'quasar';
import { cookieOptions } from 'src/utils/define';

export const UserTokenKey = '_UserToken';
export const InitStoreState = {
  //  用户Cookies Token
  userToken: '',

  //  用户信息Cookies Info
  userInfo: {
    avatar: '',
    nickname: '',
    email: '',
  } as any,

  //  管理路由列表
  routerList: [] as any,

  //  管理菜单路由
  menuList: [] as any,

  //  配置文件
  config: {
    //  项目名称
    name: 'BaJie',

    //  项目Logo
    logo: '/images/logo.png',
  },
};

// 初始化数据
export const useInitStore = defineStore('init', {
  state: () => {
    return JSON.parse(JSON.stringify(InitStoreState));
  },

  getters: {},

  actions: {
    //  更新用户Token
    updateUserToken(token: string) {
      this.userToken = token;
      Cookies.set(UserTokenKey, token, cookieOptions());
    },

    //  判断是否存在当前路由
    hasRoute(url: string) {
      const baseURL = new URL(<string>process.env.baseURL);
      return (
        this.routerList.indexOf('*') > -1 || this.routerList.indexOf(baseURL.pathname + url) > -1
      );
    },

    // executeEval 执行eval代码
    executeEval(evalStr: string, data: { scope: any }) {
      if (evalStr == '') return true;
      const { scope } = data;
      console.log(scope.row)
      return eval(evalStr);
    },
  },
});
