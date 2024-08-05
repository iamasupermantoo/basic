import { TemplateRouteInterface } from 'src/router';

import {
  LayoutsDesktopPath,
  LayoutsMobilePath,
  PagesTemplateDesktopPath,
  PagesTemplateMobilePath,
} from 'src/utils/router';

export const DefaultTemplateName = 'default';
export const defaultRouter: TemplateRouteInterface[] = [
  //  主体布局文件 TabBar + 桌面端头部尾部
  {
    name: 'Layouts',
    route: '/',
    componentMobile: LayoutsMobilePath('tabbar.vue'),
    componentDesktop: LayoutsDesktopPath('main.vue'),
    children: [
      {
        name: 'HomeIndex',
        route: '/',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'home.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'home.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
      {
        name: 'MarketIndex',
        route: '/markets',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'market.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'market.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
      {
        name: 'ContactIndex',
        route: '/contract',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'contact.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'contact.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
      {
        name: 'FuturesIndex',
        route: '/futures',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'futures.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'futures.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
      {
        name: 'UserIndex',
        route: '/user',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'users/index.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/setting/index.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
    ],
    meta: {},
  },


  //  全屏的布局文件
  {
    name: 'LayoutsFull',
    route: '/',
    componentMobile: LayoutsMobilePath('full.vue'),
    componentDesktop: LayoutsDesktopPath('main.vue'),
    children: [
      {
        name: 'UserLogin',
        route: '/login',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'login.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'home.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'UserRegister',
        route: '/register',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'register.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'home.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'MarketsDigital',
        route: '/markets/digital',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'markets/digital.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'markets/digital.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
      {
        name: 'MarketsForex',
        route: '/markets/forex',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'markets/forex.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'markets/forex.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
      {
        name: 'MarketsFutures',
        route: '/markets/futures',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'markets/futures.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'markets/futures.vue'),
        children: [],
        meta: { keepAlive: false, requireAuth: false },
      },
    ],
    meta: {},
  },
  {
    name: 'LayoutsMain',
    route: '/',
    componentMobile: LayoutsMobilePath('header.vue'),
    componentDesktop: LayoutsDesktopPath('main.vue'),
    children: [
      {
        name: 'WalletsDeposit',
        route: '/wallets/deposit',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/deposit.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/deposit.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'WalletsWithdraw',
        route: '/wallets/withdraw',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/withdraw.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/withdraw.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'WalletAccountUpdate',
        route: '/wallets/account/update',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/account/update.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/account/update.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingUpdateUserInfo',
        route: '/settings/update/info',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/setting/info.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, ''),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingUpdatePassword',
        route: '/settings/update/password',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/setting/password.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, ''),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingUpdateSecretKey',
        route: '/settings/update/secret-key',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/setting/secretKey.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, ''),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingBindTelephone',
        route: '/settings/update/telephone',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/setting/telephone.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, ''),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingUpdateEmail',
        route: '/settings/update/email',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/setting/email.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, ''),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingLanguage',
        route: '/settings/language',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/setting/language.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'home.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'ArticleDetails',
        route: '/article/details',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'article/details.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'article/details.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'UserShare',
        route: '/team/share',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/team/share.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/team/share.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'DownloadApp',
        route: '/download',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'download.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'download.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'HelpersCenter',
        route: '/helpers',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'helpers.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'helpers.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'UserLevel',
        route: '/user/level',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/level.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/level.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
    ],
    meta: {},
  },

  {
    name: 'LayoutsHeader',
    route: '/',
    componentMobile: LayoutsMobilePath('header.vue'),
    componentDesktop: LayoutsDesktopPath('setting.vue'),
    children: [
      {
        name: 'WalletsAccountIndex',
        route: '/wallets/account/index',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/account/index.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/account/index.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'WalletsIndex',
        route: '/wallets/index',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/index.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/index.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'WalletsAccountDetails',
        route: '/wallets/detail',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/details.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/index.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'WalletsAssetsIndex',
        route: '/wallets/assets/index',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/assets/index.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/assets/index.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'WalletsAssetsDetails',
        route: '/wallets/assets/details',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'wallets/assets/details.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'wallets/assets/details.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'TeamIndex',
        route: '/team/index',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/team/index.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/team/index.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },
      {
        name: 'TeamEarnings',
        route: '/team/earnings',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/team/earnings.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/team/earnings.vue'),
        children: [],
        meta: { requireAuth: false, keepAlive: false },
      },

      {
        name: 'UserRealAuth',
        route: '/user/auth',
        componentMobile: PagesTemplateMobilePath(
          DefaultTemplateName,
          'users/auth.vue'
        ),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/auth.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
      {
        name: 'SettingIndex',
        route: '/user/settings',
        componentMobile: PagesTemplateMobilePath(DefaultTemplateName, 'users/setting/index.vue'),
        componentDesktop: PagesTemplateDesktopPath(DefaultTemplateName, 'users/setting/index.vue'),
        children: [],
        meta: { requireAuth: true, keepAlive: false },
      },
    ],
    meta: {},
  },
];
