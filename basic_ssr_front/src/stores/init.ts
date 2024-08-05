import { defineStore } from 'pinia';
import { Cookies } from 'quasar';

export const UserTokenKey = '_UserToken';
export const UserLangKey = '_UserLang';

export const InitStoreState = {
  //  用户Token
  userToken: '' as any,

  //  用户语言
  userLang: '' as any,

  //  tabBars 菜单（桌面header左侧）
  tabBars: [] as menuInterface[],

  // 用户菜单(桌面左侧menu)
  userMenu: [] as menuInterface[],

  // home菜单（桌面header右侧头像）
  homeMenu: [] as menuInterface[],

  // 快捷菜单（桌面header左侧）
  quickMenu: [] as menuInterface[],

  //  翻译数据
  translate: [] as translateInterface[],

  //  国家列表
  countryList: [] as countryInterface[],

  //  语言列表
  languageList: [] as LanguageInterface[],

  //  配置文件
  config: {
    //  项目名称
    name: 'BaJie',
    //  项目Logo
    logo: '/images/logo.png',
    //  默认语言
    defaultLang: 'en-US',
    //  默认模版
    template: 'default',

    // 客服图标
    onlineIcon: '',

    // 客服链接
    onlineLink: '',

    // 充值注意事项
    depositTips: '',

    // 提现注意事项
    withdrawTips: '',

    //  配置设置
    settings: {
      //  注册配置
      register: {
        //  是否显示邮箱
        showEmail: true,
        //  是否显示确认密码
        showCmfPass: true,
        //  是否显示验证码
        showVerify: true,
        //  是否显示安全密钥
        showSecurityPass: true,
        //  是否显示手机号码
        showTelephone: true,
        //  是否显示邀请码
        isInvite: false,
      },
      //  登录配置
      login: {
        //  是否显示验证码
        showVerify: true,
        //  是否显示注册
        showRegister: true,
      },
      //  切换语言配置
      lang: {
        //  登录页面是否显示语言切换
        showLogin: true,
        //  注册页面是否显示语言切换
        showRegister: true,
        //  首页是否显示语言切换
        showHome: true,
      },
      //  在线客服配置
      online: {
        //  登录页面是否显示在线客服
        showLogin: true,
        //  注册页面是否显示在线客服
        showRegister: true,
        //  首页是否显示在线客服
        showHome: true,
        //  充值是否跳转到客服
        depositLink: false,
        //  提现是否跳转到客服
        withdrawLink: false,
      },
      wallets: {
        //  是否开启安全密钥, 如果开启了。 提现需要输入安全密钥，修改卡片｜删除卡片 需要输入安全密钥
        showSecurityPass: true,
        //  卡片管理是否开启删除
        showDelete: true,
        //  卡片管理是否开启更新
        showUpdate: true,
        //  卡片是否隐藏卡号
        showNumber: false,
      },
    },
  },
};

// 初始化数据
export const InitStore = defineStore('init', {
  state: () => {
    return JSON.parse(JSON.stringify(InitStoreState));
  },
  getters: {},
  actions: {
    //  更新用户Token和语言
    updateUserToken(token: string) {
      this.userToken = token;
      Cookies.set(UserTokenKey, token, { expires: '30d 3h 5m' });
    },

    //  更新用户语言
    updateUserLang(lang: string) {
      this.userLang = lang
      Cookies.set(UserLangKey, lang, { expires: '30d 3h 5m' });
    },

    //  删除用户Token
    removeUserToken() {
      this.userToken = ''
      Cookies.set(UserTokenKey, '')
    },
  },
});

interface menuInterface {
  // 列表名称
  name: string,

  // 列表路由
  routes: string,

  // 列表图标
  icon: string,

  // 列表高亮图标
  activeIcon: string,

  // 子级列表
  children: [],

  // 是否显示
  data: {
    // 手机是否显示
    isMobile: boolean,

    // 桌面是否显示
    isDesktop: boolean,
  }
}

interface countryInterface {
  //  国家ID
  id: number;

  //  国家名称
  name: string;

  //  国家图标
  icon: string;

  //  国家代码
  code: string;
}

interface translateInterface {
  // 建铭
  label: string;
  //  键值
  value: string;
}

interface LanguageInterface {
  //  语言ID
  id: number;

  //  语言名称
  name: string;

  //  语言别名
  alias: string;

  //  语言图标
  icon: string;
}
