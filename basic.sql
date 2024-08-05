/*
 Navicat Premium Data Transfer

 Source Server         : localhostMysql
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : basic

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 02/03/2024 20:24:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(120) NOT NULL COMMENT '名称',
  `ip` int unsigned NOT NULL COMMENT 'IP4地址',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型 1访问量',
  `headers` text COMMENT '头信息',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包资产管理表';

-- ----------------------------
-- Records of access
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_logs
-- ----------------------------
DROP TABLE IF EXISTS `admin_logs`;
CREATE TABLE `admin_logs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `name` varchar(120) NOT NULL COMMENT '名称',
  `ip` int unsigned NOT NULL COMMENT 'IP4',
  `headers` text COMMENT '头信息',
  `route` varchar(60) NOT NULL COMMENT '路由',
  `body` text COMMENT '内容',
  `data` text COMMENT '数据',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='日志表';

-- ----------------------------
-- Records of admin_logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `parent_id` int unsigned NOT NULL COMMENT '父级ID',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `type` tinyint NOT NULL COMMENT '类型 1后台菜单 11前台菜单 12用户菜单 13快捷菜单',
  `route` varchar(50) NOT NULL COMMENT '路由',
  `sort` tinyint NOT NULL COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启 ',
  `data` text COMMENT '数据',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=436 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理菜单表';

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (1, 1, 0, '后台管理', 1, '', 90, 10, '{\"icon\":\"manage_accounts\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (2, 1, 1, '管理列表', 1, '/manage/index', 0, 10, '{\"icon\":\"manage_accounts\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/manage/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (3, 1, 1, '角色配置', 1, '/manage/role/index', 1, 10, '{\"icon\":\"6_ft_apart\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/manage/role/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (4, 1, 1, '菜单管理', 1, '/manage/menu/index', 2, 10, '{\"icon\":\"menu_open\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/manage/menu/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (5, 1, 1, '操作日志', 1, '/manage/logs/index', 3, 10, '{\"icon\":\"event_note\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/manage/logs/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (20, 1, 0, '配置管理', 1, '', 91, 10, '{\"icon\":\"settings\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (21, 1, 20, '参数配置', 1, '/setting/index', 0, 10, '{\"icon\":\"settings_ethernet\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/setting/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (22, 1, 20, '菜单配置', 1, '/setting/menu/index', 1, 10, '{\"icon\":\"menu_open\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/setting/menu/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (23, 1, 20, '文章配置', 1, '/setting/article/index', 2, 10, '{\"icon\":\"article\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/setting/article/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (40, 1, 0, '系统配置', 1, '', 92, 10, '{\"icon\":\"widgets\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (41, 1, 40, '等级配置', 1, '/system/level', 0, 10, '{\"icon\":\"diamond\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/system/level/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (42, 1, 40, '国家配置', 1, '/system/country', 1, 10, '{\"icon\":\"flag\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/system/country/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (43, 1, 40, '语言配置', 1, '/system/lang', 2, 10, '{\"icon\":\"language\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/system/lang/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (44, 1, 40, '翻译配置', 1, '/system/translate', 3, 10, '{\"icon\":\"translate\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/system/translate/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (45, 1, 40, '通知配置', 1, '/system/notify', 4, 10, '{\"icon\":\"notifications\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/system/notify/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (100, 1, 0, '用户管理', 1, '', 1, 10, '{\"icon\":\"manage_accounts\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (101, 1, 100, '用户列表', 1, '/user/index', 0, 10, '{\"icon\":\"people\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/user/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (102, 1, 100, '用户资产', 1, '/user/assets/index', 1, 10, '{\"icon\":\"card_giftcard\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/user/assets/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (103, 1, 100, '会员列表', 1, '/user/level/index', 2, 10, '{\"icon\":\"diamond\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/user/level/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (104, 1, 100, '认证申请', 1, '/user/auth/index', 3, 10, '{\"icon\":\"badge\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/user/auth/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (105, 1, 100, '邀请管理', 1, '/user/invite/index', 4, 10, '{\"icon\":\"share\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/user/invite/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (106, 1, 100, '用户关系', 1, '/user/relation/index', 5, 10, '{\"icon\":\"diversity_1\",\"activeIcon\":\"\",\"vueTmp\":\"/user/relation\",\"viewsUrl\":\"/user/relation/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (120, 1, 0, '钱包管理', 1, '', 2, 10, '{\"icon\":\"account_balance_wallet\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (121, 1, 120, '充值订单', 1, '/wallets/deposit/index', 0, 10, '{\"icon\":\"format_textdirection_r_to_l\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/wallets/deposit/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (122, 1, 120, '提现订单', 1, '/wallets/withdraw/index', 1, 10, '{\"icon\":\"format_textdirection_l_to_r\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/wallets/withdraw/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (123, 1, 120, '账单明细', 1, '/wallets/bill/index', 2, 10, '{\"icon\":\"assignment\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/wallets/bill/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (124, 1, 120, '资产管理', 1, '/wallets/assets/index', 3, 10, '{\"icon\":\"card_giftcard\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/wallets/assets/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (125, 1, 120, '支付管理', 1, '/wallets/payment/index', 4, 10, '{\"icon\":\"currency_exchange\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/wallets/payment/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (126, 1, 120, '卡片管理', 1, '/wallets/account/index', 5, 10, '{\"icon\":\"credit_card\",\"activeIcon\":\"\",\"vueTmp\":\"/views\",\"viewsUrl\":\"/wallets/account/views\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (200, 1, 0, 'walletManage', 12, '', 20, 10, '{\"icon\":\"/assets/icon/menu/wallets-manage.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (201, 1, 200, 'accountManage', 12, '/wallets/account/index', 1, 10, '{\"icon\":\"/assets/icon/menu/account.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (202, 1, 200, 'myWallet', 12, '/wallets/index', 2, 10, '{\"icon\":\"/assets/icon/menu/wallets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (203, 1, 200, 'myAssets', 12, '/wallets/assets/index', 3, 10, '{\"icon\":\"/assets/icon/menu/assets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (220, 1, 0, 'teamManage', 12, '', 21, 10, '{\"icon\":\"/assets/icon/menu/team-manage.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (221, 1, 220, 'inviteFriends', 12, '/team/share', 1, 10, '{\"icon\":\"/assets/icon/menu/share.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (222, 1, 220, 'myTeam', 12, '/team/index', 2, 10, '{\"icon\":\"/assets/icon/menu/team.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (223, 1, 220, 'teamEarnings', 12, '/team/earnings', 3, 10, '{\"icon\":\"/assets/icon/menu/earnings.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (230, 1, 0, 'serviceManage', 12, '', 22, 10, '{\"icon\":\"/assets/icon/menu/service-manage.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (231, 1, 230, 'memberVip', 12, '/user/level', 1, 10, '{\"icon\":\"/assets/icon/menu/level.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (232, 1, 230, 'realAuth', 12, '/user/auth', 2, 10, '{\"icon\":\"/assets/icon/menu/realauth.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (233, 1, 230, 'helpers', 12, '/helpers', 3, 10, '{\"icon\":\"/assets/icon/menu/helpers.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (234, 1, 230, 'settings', 12, '/user/settings', 4, 10, '{\"icon\":\"/assets/icon/menu/settings.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (235, 1, 230, 'download', 12, '/download', 5, 10, '{\"icon\":\"/assets/icon/menu/download.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (260, 1, 0, 'myWallet', 11, '/wallets/index', 0, 10, '{\"icon\":\"/assets/icon/menu/wallets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (261, 1, 0, 'myAssets', 11, '/wallets/assets/index', 0, 10, '{\"icon\":\"/assets/icon/menu/assets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (262, 1, 0, 'myTeam', 11, '/team/index', 0, 10, '{\"icon\":\"/assets/icon/menu/team.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (263, 1, 0, 'inviteFriends', 11, '/team/share', 0, 10, '{\"icon\":\"/assets/icon/menu/share.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (300, 1, 0, 'home', 15, '/', 1, 10, '{\"icon\":\"/assets/icon/tabbar/home.png\",\"activeIcon\":\"/assets/icon/tabbar/home_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (301, 1, 0, 'markets', 15, '/markets', 2, 10, '{\"icon\":\"/assets/icon/tabbar/markets.png\",\"activeIcon\":\"/assets/icon/tabbar/markets_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (302, 1, 0, 'contract', 15, '/contract', 3, 10, '{\"icon\":\"/assets/icon/tabbar/contract.png\",\"activeIcon\":\"/assets/icon/tabbar/contract_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (303, 1, 0, 'futures', 15, '/futures', 4, 10, '{\"icon\":\"/assets/icon/tabbar/futures.png\",\"activeIcon\":\"/assets/icon/tabbar/futures_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (304, 1, 0, 'user', 15, '/user', 5, 10, '{\"icon\":\"/assets/icon/tabbar/user.png\",\"activeIcon\":\"/assets/icon/tabbar/user_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (305, 1, 301, 'digitalMarkets', 15, '/markets/digital', 0, 10, '{\"icon\":\"/assets/icon/tabbar/markets_active.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (306, 1, 301, 'forexMarkets', 15, '/markets/forex', 1, 10, '{\"icon\":\"/assets/icon/tabbar/contract_active.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (307, 1, 301, 'futuresMarkets', 15, '/markets/futures', 2, 10, '{\"icon\":\"/assets/icon/tabbar/futures_active.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (400, 1, 0, 'deposit', 13, '/wallets/deposit?mode=1', 1, 10, '{\"icon\":\"/assets/icon/menu/deposit.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (401, 1, 0, 'withdraw', 13, '/wallets/withdraw?mode=11', 2, 10, '{\"icon\":\"/assets/icon/menu/withdraw.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (402, 1, 0, 'helpers', 13, '/helpers', 3, 10, '{\"icon\":\"/assets/icon/menu/chatbot.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (403, 1, 0, 'download', 13, '/download', 4, 10, '{\"icon\":\"/assets/icon/menu/download.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (404, 1, 0, 'memberVip', 13, '/user/level', 4, 10, '{\"icon\":\"/assets/icon/menu/level.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (405, 2, 0, 'walletManage', 12, '', 20, 10, '{\"icon\":\"/assets/icon/menu/wallets-manage.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (406, 2, 405, 'accountManage', 12, '/wallets/account/index', 1, 10, '{\"icon\":\"/assets/icon/menu/account.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (407, 2, 405, 'myWallet', 12, '/wallets/index', 2, 10, '{\"icon\":\"/assets/icon/menu/wallets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (408, 2, 405, 'myAssets', 12, '/wallets/assets/index', 3, 10, '{\"icon\":\"/assets/icon/menu/assets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (409, 2, 0, 'teamManage', 12, '', 21, 10, '{\"icon\":\"/assets/icon/menu/team-manage.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (410, 2, 409, 'inviteFriends', 12, '/team/share', 1, 10, '{\"icon\":\"/assets/icon/menu/share.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (411, 2, 409, 'myTeam', 12, '/team/index', 2, 10, '{\"icon\":\"/assets/icon/menu/team.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (412, 2, 409, 'teamEarnings', 12, '/team/earnings', 3, 10, '{\"icon\":\"/assets/icon/menu/earnings.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (413, 2, 0, 'serviceManage', 12, '', 22, 10, '{\"icon\":\"/assets/icon/menu/service-manage.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (414, 2, 413, 'memberVip', 12, '/user/level', 1, 10, '{\"icon\":\"/assets/icon/menu/level.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (415, 2, 413, 'realAuth', 12, '/user/auth', 2, 10, '{\"icon\":\"/assets/icon/menu/realauth.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (416, 2, 413, 'helpers', 12, '/helpers', 3, 10, '{\"icon\":\"/assets/icon/menu/helpers.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (417, 2, 413, 'settings', 12, '/user/settings', 4, 10, '{\"icon\":\"/assets/icon/menu/settings.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (418, 2, 413, 'download', 12, '/download', 5, 10, '{\"icon\":\"/assets/icon/menu/download.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (419, 2, 0, 'myWallet', 11, '/wallets/index', 0, 10, '{\"icon\":\"/assets/icon/menu/wallets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (420, 2, 0, 'myAssets', 11, '/wallets/assets/index', 0, 10, '{\"icon\":\"/assets/icon/menu/assets.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (421, 2, 0, 'myTeam', 11, '/team/index', 0, 10, '{\"icon\":\"/assets/icon/menu/team.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (422, 2, 0, 'inviteFriends', 11, '/team/share', 0, 10, '{\"icon\":\"/assets/icon/menu/share.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (423, 2, 0, 'home', 15, '/', 1, 10, '{\"icon\":\"/assets/icon/tabbar/home.png\",\"activeIcon\":\"/assets/icon/tabbar/home_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (424, 2, 0, 'markets', 15, '/markets', 2, 10, '{\"icon\":\"/assets/icon/tabbar/markets.png\",\"activeIcon\":\"/assets/icon/tabbar/markets_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (425, 2, 424, 'digitalMarkets', 15, '/markets/digital', 0, 10, '{\"icon\":\"/assets/icon/tabbar/markets_active.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (426, 2, 424, 'forexMarkets', 15, '/markets/forex', 1, 10, '{\"icon\":\"/assets/icon/tabbar/contract_active.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (427, 2, 424, 'futuresMarkets', 15, '/markets/futures', 2, 10, '{\"icon\":\"/assets/icon/tabbar/futures_active.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (428, 2, 0, 'contract', 15, '/contract', 3, 10, '{\"icon\":\"/assets/icon/tabbar/contract.png\",\"activeIcon\":\"/assets/icon/tabbar/contract_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (429, 2, 0, 'futures', 15, '/futures', 4, 10, '{\"icon\":\"/assets/icon/tabbar/futures.png\",\"activeIcon\":\"/assets/icon/tabbar/futures_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (430, 2, 0, 'user', 15, '/user', 5, 10, '{\"icon\":\"/assets/icon/tabbar/user.png\",\"activeIcon\":\"/assets/icon/tabbar/user_active.png\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":false,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (431, 2, 0, 'deposit', 13, '/wallets/deposit?mode=1', 1, 10, '{\"icon\":\"/assets/icon/menu/deposit.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (432, 2, 0, 'withdraw', 13, '/wallets/withdraw?mode=11', 2, 10, '{\"icon\":\"/assets/icon/menu/withdraw.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":true}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (433, 2, 0, 'helpers', 13, '/helpers', 3, 10, '{\"icon\":\"/assets/icon/menu/chatbot.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (434, 2, 0, 'download', 13, '/download', 4, 10, '{\"icon\":\"/assets/icon/menu/download.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
INSERT INTO `admin_menu` (`id`, `admin_id`, `parent_id`, `name`, `type`, `route`, `sort`, `status`, `data`, `created_at`) VALUES (435, 2, 0, 'memberVip', 13, '/user/level', 4, 10, '{\"icon\":\"/assets/icon/menu/level.png\",\"activeIcon\":\"\",\"vueTmp\":\"\",\"viewsUrl\":\"\",\"isPc\":true,\"isMobile\":false}', 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for admin_setting
-- ----------------------------
DROP TABLE IF EXISTS `admin_setting`;
CREATE TABLE `admin_setting` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `group_id` int unsigned NOT NULL COMMENT '分组ID',
  `name` varchar(60) NOT NULL COMMENT '设置名称',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT 'input类型 1文本类型 2文本域类型 3富文本类型 4数字类型 5密码类型 10选择类型 11单选类型 12多选类型 13开关类型 21文件类型 22图片类型 23图片组类型 31时间类型 32时间范围类型 41对象类型 42子集对象类型 52编辑文本类型 53编辑富文本类型 54编辑开关类型 61翻译类型',
  `field` varchar(60) NOT NULL COMMENT '建铭',
  `value` text COMMENT '键值',
  `data` text COMMENT 'input配置',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理设置表';

-- ----------------------------
-- Records of admin_setting
-- ----------------------------
BEGIN;
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (1, 1, 1, '站点名称', 1, 'siteName', 'BaJie', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (2, 1, 1, '站点Logo', 22, 'siteLogo', '', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (3, 1, 1, '站点介绍', 61, 'siteIntroduce', 'introduce', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (4, 1, 1, '站点公告', 61, 'siteAnnouncement', 'announcement', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (5, 1, 1, '轮播图', 23, 'siteBanners', '[\"/assets/banner/banner1.png\", \"/assets/banner/banner2.png\", \"/assets/banner/banner3.png\", \"/assets/banner/banner4.png\"]', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (6, 1, 1, '客服链接', 1, 'onlineLink', '', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (7, 1, 1, '客服图标', 22, 'onlineImage', '/assets/icon/online.png', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (8, 1, 1, '社区账号', 42, 'socialList', '[{\"icon\": \"/assets/icon/telegram.png\", \"name\": \"Telegram\", \"link\": \"https://t.me/bajie119\", \"status\": 10}, {\"icon\": \"/assets/icon/facebook.png\", \"name\": \"Facebook\", \"link\": \"\", \"status\": 10}, {\"icon\": \"/assets/icon/whatsapp.png\", \"name\": \"Whatsapp\", \"link\": \"\", \"status\": 10}, {\"icon\": \"/assets/icon/line.png\", \"name\": \"Line\", \"link\": \"\", \"status\": 10}, {\"icon\": \"/assets/icon/twitter.png\", \"name\": \"Twitter\", \"link\": \"\", \"status\": 10}]', '[[{\"label\": \"图标\", \"field\": \"icon\", \"type\": 22}],[{\"label\": \"名称\", \"field\": \"name\", \"type\": 1}, {\"label\": \"链接\", \"field\": \"link\", \"type\": 1}, {\"label\": \"状态\", \"field\": \"status\", \"type\": 10, \"data\": [{\"label\": \"开启\", \"value\": 10}, {\"label\": \"关闭\", \"value\": -1}]}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (9, 1, 1, '联系我们', 42, 'contactList', '[{\"icon\": \"place\", \"name\": \"607-699 14th Pl NE, Washington, DC 20002 USA\", \"link\": \"https://maps.app.goo.gl/U2gNxExtu16Tbasb8\", \"status\": 10}, {\"icon\": \"phone_iphone\", \"name\": \"+12023883500\", \"link\": \"tel:12023883500\", \"status\": 10}, {\"icon\": \"forward_to_inbox\", \"name\": \"muiprosperyls15@gmail.com\", \"link\": \"mailto:muiprosperyls15@gmail.com\", \"status\": 10}]', '[[{\"label\": \"图标\", \"field\": \"icon\", \"type\": 22}],[{\"label\": \"名称\", \"field\": \"name\", \"type\": 1}, {\"label\": \"链接\", \"field\": \"link\", \"type\": 1}, {\"label\": \"状态\", \"field\": \"status\", \"type\": 10, \"data\": [{\"label\": \"开启\", \"value\": 10}, {\"label\": \"关闭\", \"value\": -1}]}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (10, 1, 1, '软件下载', 41, 'download', '{\"android\": \"\", \"ios\": \"\"}', '[[{\"label\": \"安卓包地址\", \"field\": \"android\", \"type\": 21}, {\"label\": \"苹果包地址\", \"field\": \"ios\", \"type\": 21}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (11, 1, 1, '等级购买模式', 10, 'buyLevelMode', '1', '[{\"label\": \"补价购买\", \"value\": 1}, {\"label\": \"等额购买\", \"value\": 2}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (12, 1, 1, '提示音设置', 12, 'audioSetting', '{\"deposit\": false, \"withdraw\": false}', '[{\"label\": \"充值\", \"value\": \"deposit\"}, {\"label\": \"提现\", \"value\": \"withdraw\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (13, 1, 2, '绑定卡片类型数量', 4, 'walletAccountNums', '5', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (14, 1, 2, '充值金额范围', 41, 'walletDepositAmountBetween', '{\"min\": 100, \"max\": 999999}', '[[{\"label\": \"最小金额\", \"field\": \"min\", \"type\": 4}, {\"label\": \"最大金额\", \"field\": \"max\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (15, 1, 2, '充值注意事项提示', 61, 'walletDepositDesc', 'depositDesc', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (16, 1, 2, '提现金额范围', 41, 'walletWithdrawAmountBetween', '{\"min\": 100, \"max\": 999999}', '[[{\"label\": \"最小金额\", \"field\": \"min\", \"type\": 4}, {\"label\": \"最大金额\", \"field\": \"max\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (17, 1, 2, '提现配置', 41, 'walletWithdrawSetting', '{\"day\": 7, \"nums\": 7, \"fee\": 0}', '[[{\"label\": \"天数\", \"field\": \"day\", \"type\": 4}, {\"label\": \"次数\", \"field\": \"nums\", \"type\": 4}, {\"label\": \"手续费\", \"field\": \"fee\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (18, 1, 2, '提现注意事项提示', 61, 'walletWithdrawDesc', 'withdrawDesc', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (19, 1, 3, '注册邀请奖励', 41, 'registerAward', '{\"register\": 0, \"share\": 0}', '[[{\"label\": \"注册奖励\", \"field\": \"register\", \"type\": 4}, {\"label\": \"邀请奖励\", \"field\": \"share\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (20, 1, 3, '分销配置(第一条数据为 1Lv ...类推)', 42, 'earningsSetting', '[{\"options\": {}, \"rate\": 50}, {\"options\": {}, \"rate\": 30}, {\"options\": {\"walletBillTypeBuyProduct\": false, \"walletBillTypeAssetsBuyProduct\": false, \"walletBillTypeEarnings\": false, \"walletBillTypeAssetsEarnings\": false}, \"rate\": 10}]', '[[{\"label\": \"操作\", \"field\": \"options\", \"type\": 12, \"data\": [{\"label\": \"账户购买产品\", \"value\": \"walletBillTypeBuyProduct\"}, {\"label\": \"资产购买产品\", \"value\": \"walletBillTypeAssetsBuyProduct\"}, {\"label\": \"账户收益\", \"value\": \"walletBillTypeEarnings\"}, {\"label\": \"资产收益\", \"value\": \"walletBillTypeAssetsEarnings\"}]}, {\"label\": \"收益率\", \"field\": \"rate\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (21, 1, 4, '基础配置', 12, 'templateBasic', '{\"isFreezeOrder\": true, \"isFreezeWithdraw\": true}', '[{\"label\": \"禁用账户禁止下单\", \"value\": \"isFreezeOrder\"}, {\"label\": \"禁用账户禁止提现\", \"value\": \"isFreezeWithdraw\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (22, 1, 4, '登录配置', 12, 'templateLogin', '{\"showVerify\": true, \"showRegister\": true}', '[{\"label\": \"是否显示验证码\", \"value\": \"showVerify\"}, {\"label\": \"是否显示注册按钮\", \"value\": \"showRegister\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (23, 1, 4, '注册配置', 12, 'templateRegister', '{\"showVerify\": true, \"showCmfPass\": false, \"showSecurityPass\": false, \"showTelephone\": false, \"showEmail\": false, \"isInvite\": false, \"closeRegister\": false}', '[{\"label\": \"是否显示验证码\", \"value\": \"showVerify\"}, {\"label\": \"是否显示确认密码\", \"value\": \"showCmfPassword\"}, {\"label\": \"是否显示安全密码\", \"value\": \"showSecurityPassword\"}, {\"label\": \"是否显示手机号码\", \"value\": \"showTelephone\"}, {\"label\": \"是否显示电子邮箱\", \"value\": \"showEmail\"}, {\"label\": \"是否必须邀请码\", \"value\": \"isInvite\"}, {\"label\": \"是否关闭注册\", \"value\": \"closeRegister\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (24, 1, 4, '切换语言配置', 12, 'templateLang', '{\"showLogin\": true, \"showRegister\": true, \"showHome\": true}', '[{\"label\": \"登录是否显示\", \"value\": \"showLogin\"}, {\"label\": \"注册是否显示\", \"value\": \"showRegister\"}, {\"label\": \"是否显示主页\", \"value\": \"showHome\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (25, 1, 4, '在线客服配置', 12, 'templateOnline', '{\"showLogin\": true, \"showRegister\": true, \"showHome\": true, \"depositLink\": false, \"withdrawLink\": false}', '[{\"label\": \"登录是否显示\", \"value\": \"showLogin\"}, {\"label\": \"注册是否显示\", \"value\": \"showRegister\"}, {\"label\": \"首页是否显示\", \"value\": \"showHome\"}, {\"label\": \"充值跳转客服\", \"value\": \"linkDeposit\"}, {\"label\": \"提现跳转客服\", \"value\": \"linkWithdraw\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (26, 1, 4, '钱包配置', 12, 'templateWallet', '{\"showSecurityPass\": true, \"showDelete\": true, \"showUpdate\": true, \"showNumber\": true}', '[{\"label\": \"是否输入安全密码\", \"value\": \"showSecurityPass\"}, {\"label\": \"是否显示更新\", \"value\": \"showUpdate\"}, {\"label\": \"是否显示删除\", \"value\": \"showDelete\"}, {\"label\": \"是否显示卡号\", \"value\": \"showNumber\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (27, 2, 1, '站点名称', 1, 'siteName', 'BaJie', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (28, 2, 1, '站点Logo', 22, 'siteLogo', '', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (29, 2, 1, '站点介绍', 61, 'siteIntroduce', 'introduce', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (30, 2, 1, '站点公告', 61, 'siteAnnouncement', 'announcement', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (31, 2, 1, '轮播图', 23, 'siteBanners', '[\"/assets/banner/banner1.png\", \"/assets/banner/banner2.png\", \"/assets/banner/banner3.png\", \"/assets/banner/banner4.png\"]', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (32, 2, 1, '客服链接', 1, 'onlineLink', '', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (33, 2, 1, '客服图标', 22, 'onlineImage', '/assets/icon/online.png', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (34, 2, 1, '社区账号', 42, 'socialList', '[{\"icon\": \"/assets/icon/telegram.png\", \"name\": \"Telegram\", \"link\": \"https://t.me/bajie119\", \"status\": 10}, {\"icon\": \"/assets/icon/facebook.png\", \"name\": \"Facebook\", \"link\": \"\", \"status\": 10}, {\"icon\": \"/assets/icon/whatsapp.png\", \"name\": \"Whatsapp\", \"link\": \"\", \"status\": 10}, {\"icon\": \"/assets/icon/line.png\", \"name\": \"Line\", \"link\": \"\", \"status\": 10}, {\"icon\": \"/assets/icon/twitter.png\", \"name\": \"Twitter\", \"link\": \"\", \"status\": 10}]', '[[{\"label\": \"图标\", \"field\": \"icon\", \"type\": 22}],[{\"label\": \"名称\", \"field\": \"name\", \"type\": 1}, {\"label\": \"链接\", \"field\": \"link\", \"type\": 1}, {\"label\": \"状态\", \"field\": \"status\", \"type\": 10, \"data\": [{\"label\": \"开启\", \"value\": 10}, {\"label\": \"关闭\", \"value\": -1}]}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (35, 2, 1, '联系我们', 42, 'contactList', '[{\"icon\": \"place\", \"name\": \"607-699 14th Pl NE, Washington, DC 20002 USA\", \"link\": \"https://maps.app.goo.gl/U2gNxExtu16Tbasb8\", \"status\": 10}, {\"icon\": \"phone_iphone\", \"name\": \"+12023883500\", \"link\": \"tel:12023883500\", \"status\": 10}, {\"icon\": \"forward_to_inbox\", \"name\": \"muiprosperyls15@gmail.com\", \"link\": \"mailto:muiprosperyls15@gmail.com\", \"status\": 10}]', '[[{\"label\": \"图标\", \"field\": \"icon\", \"type\": 22}],[{\"label\": \"名称\", \"field\": \"name\", \"type\": 1}, {\"label\": \"链接\", \"field\": \"link\", \"type\": 1}, {\"label\": \"状态\", \"field\": \"status\", \"type\": 10, \"data\": [{\"label\": \"开启\", \"value\": 10}, {\"label\": \"关闭\", \"value\": -1}]}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (36, 2, 1, '软件下载', 41, 'download', '{\"android\": \"\", \"ios\": \"\"}', '[[{\"label\": \"安卓包地址\", \"field\": \"android\", \"type\": 21}, {\"label\": \"苹果包地址\", \"field\": \"ios\", \"type\": 21}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (37, 2, 1, '等级购买模式', 10, 'buyLevelMode', '1', '[{\"label\": \"补价购买\", \"value\": 1}, {\"label\": \"等额购买\", \"value\": 2}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (38, 2, 1, '提示音设置', 12, 'audioSetting', '{\"deposit\": false, \"withdraw\": false}', '[{\"label\": \"充值\", \"value\": \"deposit\"}, {\"label\": \"提现\", \"value\": \"withdraw\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (39, 2, 2, '绑定卡片类型数量', 4, 'walletAccountNums', '5', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (40, 2, 2, '充值金额范围', 41, 'walletDepositAmountBetween', '{\"min\": 100, \"max\": 999999}', '[[{\"label\": \"最小金额\", \"field\": \"min\", \"type\": 4}, {\"label\": \"最大金额\", \"field\": \"max\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (41, 2, 2, '充值注意事项提示', 61, 'walletDepositDesc', 'depositDesc', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (42, 2, 2, '提现金额范围', 41, 'walletWithdrawAmountBetween', '{\"min\": 100, \"max\": 999999}', '[[{\"label\": \"最小金额\", \"field\": \"min\", \"type\": 4}, {\"label\": \"最大金额\", \"field\": \"max\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (43, 2, 2, '提现配置', 41, 'walletWithdrawSetting', '{\"day\": 7, \"nums\": 7, \"fee\": 0}', '[[{\"label\": \"天数\", \"field\": \"day\", \"type\": 4}, {\"label\": \"次数\", \"field\": \"nums\", \"type\": 4}, {\"label\": \"手续费\", \"field\": \"fee\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (44, 2, 2, '提现注意事项提示', 61, 'walletWithdrawDesc', 'withdrawDesc', '', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (45, 2, 3, '注册邀请奖励', 41, 'registerAward', '{\"register\": 0, \"share\": 0}', '[[{\"label\": \"注册奖励\", \"field\": \"register\", \"type\": 4}, {\"label\": \"邀请奖励\", \"field\": \"share\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (46, 2, 3, '分销配置(第一条数据为 1Lv ...类推)', 42, 'earningsSetting', '[{\"options\": {}, \"rate\": 50}, {\"options\": {}, \"rate\": 30}, {\"options\": {\"walletBillTypeBuyProduct\": false, \"walletBillTypeAssetsBuyProduct\": false, \"walletBillTypeEarnings\": false, \"walletBillTypeAssetsEarnings\": false}, \"rate\": 10}]', '[[{\"label\": \"操作\", \"field\": \"options\", \"type\": 12, \"data\": [{\"label\": \"账户购买产品\", \"value\": \"walletBillTypeBuyProduct\"}, {\"label\": \"资产购买产品\", \"value\": \"walletBillTypeAssetsBuyProduct\"}, {\"label\": \"账户收益\", \"value\": \"walletBillTypeEarnings\"}, {\"label\": \"资产收益\", \"value\": \"walletBillTypeAssetsEarnings\"}]}, {\"label\": \"收益率\", \"field\": \"rate\", \"type\": 4}]]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (47, 2, 4, '基础配置', 12, 'templateBasic', '{\"isFreezeOrder\": true, \"isFreezeWithdraw\": true}', '[{\"label\": \"禁用账户禁止下单\", \"value\": \"isFreezeOrder\"}, {\"label\": \"禁用账户禁止提现\", \"value\": \"isFreezeWithdraw\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (48, 2, 4, '登录配置', 12, 'templateLogin', '{\"showVerify\": true, \"showRegister\": true}', '[{\"label\": \"是否显示验证码\", \"value\": \"showVerify\"}, {\"label\": \"是否显示注册按钮\", \"value\": \"showRegister\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (49, 2, 4, '注册配置', 12, 'templateRegister', '{\"showVerify\": true, \"showCmfPass\": false, \"showSecurityPass\": false, \"showTelephone\": false, \"showEmail\": false, \"isInvite\": false, \"closeRegister\": false}', '[{\"label\": \"是否显示验证码\", \"value\": \"showVerify\"}, {\"label\": \"是否显示确认密码\", \"value\": \"showCmfPassword\"}, {\"label\": \"是否显示安全密码\", \"value\": \"showSecurityPassword\"}, {\"label\": \"是否显示手机号码\", \"value\": \"showTelephone\"}, {\"label\": \"是否显示电子邮箱\", \"value\": \"showEmail\"}, {\"label\": \"是否必须邀请码\", \"value\": \"isInvite\"}, {\"label\": \"是否关闭注册\", \"value\": \"closeRegister\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (50, 2, 4, '切换语言配置', 12, 'templateLang', '{\"showLogin\": true, \"showRegister\": true, \"showHome\": true}', '[{\"label\": \"登录是否显示\", \"value\": \"showLogin\"}, {\"label\": \"注册是否显示\", \"value\": \"showRegister\"}, {\"label\": \"是否显示主页\", \"value\": \"showHome\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (51, 2, 4, '在线客服配置', 12, 'templateOnline', '{\"showLogin\": true, \"showRegister\": true, \"showHome\": true, \"depositLink\": false, \"withdrawLink\": false}', '[{\"label\": \"登录是否显示\", \"value\": \"showLogin\"}, {\"label\": \"注册是否显示\", \"value\": \"showRegister\"}, {\"label\": \"首页是否显示\", \"value\": \"showHome\"}, {\"label\": \"充值跳转客服\", \"value\": \"linkDeposit\"}, {\"label\": \"提现跳转客服\", \"value\": \"linkWithdraw\"}]', 1709021277, 1709021277);
INSERT INTO `admin_setting` (`id`, `admin_id`, `group_id`, `name`, `type`, `field`, `value`, `data`, `updated_at`, `created_at`) VALUES (52, 2, 4, '钱包配置', 12, 'templateWallet', '{\"showSecurityPass\": true, \"showDelete\": true, \"showUpdate\": true, \"showNumber\": true}', '[{\"label\": \"是否输入安全密码\", \"value\": \"showSecurityPass\"}, {\"label\": \"是否显示更新\", \"value\": \"showUpdate\"}, {\"label\": \"是否显示删除\", \"value\": \"showDelete\"}, {\"label\": \"是否显示卡号\", \"value\": \"showNumber\"}]', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int unsigned NOT NULL COMMENT '上级ID',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `nickname` varchar(60) NOT NULL COMMENT '昵称',
  `email` varchar(60) NOT NULL COMMENT '邮箱',
  `avatar` varchar(120) NOT NULL COMMENT '头像',
  `password` varchar(120) NOT NULL COMMENT '密码',
  `security_key` varchar(120) NOT NULL COMMENT '密钥',
  `money` decimal(12,2) NOT NULL COMMENT '金额',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 10激活 -1冻结 -2删除',
  `data` text COMMENT '数据',
  `domains` varchar(1020) NOT NULL COMMENT '绑定域名',
  `expired_at` int unsigned NOT NULL COMMENT '过期时间',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理表';

-- ----------------------------
-- Records of admin_user
-- ----------------------------
BEGIN;
INSERT INTO `admin_user` (`id`, `parent_id`, `username`, `nickname`, `email`, `avatar`, `password`, `security_key`, `money`, `status`, `data`, `domains`, `expired_at`, `updated_at`, `created_at`) VALUES (1, 0, 'admin', '八戒网络科技', 'muiprosperyls15@gmail.com', '', 'f32761253f92b0c52d32a4bd0718276a', 'f32761253f92b0c52d32a4bd0718276a', 0.00, 10, '{\"rsaList\":{\"admin\":{\"priKey\":\"-----BEGIN PRIVATE KEY-----\\nMIIEogIBAAKCAQEAuWN/IJu8gVbHCdGE/jMBr0JmiGtpMU/ZXMnnb/J15oRb4exa\\nxE7JEHjJ01KBeLZIsk0LyVl5Vv7HBa5FwN/t0Wh8/QOv50lxMuvj9C/AeZGLBzRI\\nOmWH/x6wwqoHxl2GeVBQo3c8LMJiOCsabKq2I79oHpzZaOcIOW3UMkoPbGdkxMOt\\nmq77ZGHFYZfusAvcn5XYl+dkFtMKiu18zb1tMjM5BL3j6udwt/WiB5B2FV0ApmMO\\nK1AZYzQ+vesbid+7gS6EYRQ8YP36tTOejn2euSySDNWKPahF6pykEyvfFFqzA9PT\\ngjRF4Stcn6BxDT7MHCeu9yqKlYioXvfSmxa54QIDAQABAoIBABGZBT9GKT2pzBj9\\nf39VvioxxmPatlCNANgS/lodDN4F4Gbwtaj+xY0/ugDKSyn3O61ZLO4/BQqiAqhi\\nY5Ksbvm6zmqCCBePXCR2SdwrED05+JWFe/m3G5K3ChZgZ36H8DThz9XEgrzI8uPc\\nC5UXG1UHXU3Sm7yeeaRh7YuchhxHBuRObTjfZZOJVJKYjCXuurmzCdsFu/ZlDWaQ\\nfahm/sGdSj/G9qVntqmF168Ao8G8w06lFRTNXUEwVAcKeeu/1ABIiwmnQPvW2pqP\\nvQVTMEUCds+IgcxWJCb87f0l1NT03WBd7UH0j929v31GYg+BHA46xNLc/jtf4vAe\\nWL2H+4UCgYEA4EdIOnzU8fchBmsJu6YWNZKvi2GQQsjWhA3xf5qnmhHpsZ2EJWB0\\nVwqFg2FmdkC6neqaD+xOwglW5mVrzjb03WxWILuwUwMJo6J6BuNf8gjMqSOameqf\\nAN2VIB6vzkWkddAhQH6rEA7u0bFtCyzhB0KF4kojfrTAtIqQvTRXC2MCgYEA05wV\\n3g464z3PSc810VOqe92LauEsvsiyqva2mXQ90SAz+LBiGJBcTMb40I0Mdn6Q+hxu\\nwqoPyGk7KQxOYAnNdHE/dYnjsz/rwzNifjwKeUiFSEIFWLsRetRxhW7qg/9n6kX0\\nwqvkU2VzBF2vFXjdWt9b9tRgF9Z9C33EqqktgusCgYBAbrolc2+KBEGMonutWU3Q\\nHlAobuMPDLv0PD1BN6Em9jZ5PJOWWVuTFga9c+IH3xi9/YQ9RtppjF1W25RZLhiy\\n9EjaJpHFh33hcPA1wmTyF+0UOpJT3b/Ic+A/1hET5ZYV8rFa4gkrF98shxiYuU/8\\n0fO8yyffYvZp59UlRkArmQKBgACvPJLfUOlzRbxjYzUuJBsKeGz8FXz4gTt0WNre\\nOWT3ybNAPtD9ho7pBd0G18d6WVW1ydXvXuWzEXHsjERbQ8Lgqufibk4iIs0a2XK3\\nFwGVbnjxXbsPv9q53TQlS7TpmphzaXtHuxFZ/qlA9FQJdA3bMxz1SupI01a6LWyv\\n52nJAoGAcxNfT9rP8IKmmhWx5agNKYbh5L3L9+2iNRbNOkFOGQvFdzw0MsJvEd5r\\n8KRLrJLRJ6EdtazKAZQASXJgW/gus99xtarVS8/DxdwXtO7PseRO4zJJZYmIFFul\\nOTyyGWW07tqgahyM281ccv3NTjO9vw6kG+TKRwwMsWxOE5v+3yM=\\n-----END PRIVATE KEY-----\\n\",\"pubKey\":\"-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuWN/IJu8gVbHCdGE/jMB\\nr0JmiGtpMU/ZXMnnb/J15oRb4exaxE7JEHjJ01KBeLZIsk0LyVl5Vv7HBa5FwN/t\\n0Wh8/QOv50lxMuvj9C/AeZGLBzRIOmWH/x6wwqoHxl2GeVBQo3c8LMJiOCsabKq2\\nI79oHpzZaOcIOW3UMkoPbGdkxMOtmq77ZGHFYZfusAvcn5XYl+dkFtMKiu18zb1t\\nMjM5BL3j6udwt/WiB5B2FV0ApmMOK1AZYzQ+vesbid+7gS6EYRQ8YP36tTOejn2e\\nuSySDNWKPahF6pykEyvfFFqzA9PTgjRF4Stcn6BxDT7MHCeu9yqKlYioXvfSmxa5\\n4QIDAQAB\\n-----END PUBLIC KEY-----\\n\"},\"home\":{\"priKey\":\"-----BEGIN PRIVATE KEY-----\\nMIIEogIBAAKCAQEA73ZtAlS3y+mXpHXqbZ5xfPGGeIl5VKcoTQOyGzD+tgODd9m1\\ndlq1Qz0HSRUYlAaxg5c3hb12Ty0L3tLEhu02+LIRrknnqazVS2+XXT2Y14WblYTd\\nOOrnoEPD4Kfa3Lb74IoMJ50F/bSkHY6csjNmJEaRSIun5tW+pTZhkKDYSBu2kNev\\nHFV5k20iDNlsf2MlXhalXIG2IXzB5Z2JtYhqEX9tQJ6INACnevD2N4AfFSX3tbWR\\nsVHMC0v/wDdA/cgW/LP44mUEvT1cwt5QfEAgyCEChp8+UFMsDec3gtHlyze1ZPIt\\nsg5Xq7bzYpKR1R7DQwSIIhY0+qkfpmdVTMVX5QIDAQABAoIBAFo8TieqtPfqNnKK\\n3KQiKLHkcb/KTiZQNyeOVPdaJyF/gXMQXwkSdWu4+53WjUR6oTntKccD9ikv2GFH\\nGzec/DILKA59WwbdUiLzEh7Yr1fHyTE2uAZSvqXt40os9pRlf9TFMH9c9Hz+LTnW\\nG8YrCDpzPRwRRFv//69SJsHRWq8Qx7BUDwQwbb5yGaaueuo1v4G/Mk2co8F7vr+I\\nadujmMVqf4c+enmRWYE4zdVDuAbT7G4w6DahQmdlNMeLofNYUSiTT9hsIW1Be2F5\\n8FFLtV2IF4rwE+WjuQ1bvVOHRiuvS+f1agrjmINHvsQe9T+Pbepr3WXUMB1Iq/Vj\\n7lOO1wECgYEA/7G5CL0LcLmlr4LxKAXLMnCtb/Q8PxlLbUBlzqu9zvOmolWKA4PK\\noemjSRiv2yVmgVgGqhtgKvP0/CHoE9D/L6v20HE4jQC8m6BPk3EkjRU/Yrev+QtI\\nG/62cUq/NnRK6sU28g+pj6xTkTjyTRS2rwR1sB64PLkER8FWCP9Mxe0CgYEA77+7\\n45IhS/OSSg/AnHRx8y5wto9Xcj3aQN/bVX3q5vxPVDrLhgPDYgGQhr6ZugALOaPz\\ngMJg01bIv6aZDNfs7kWpv5T1IqMt/Dt2wOhRqP+d0bMoVmBp3QwSj35ueBQcZYq8\\nnzptCe/1+ssxtfbXj59mYGmvLZBq8kk/CTgFmtkCgYAlBYWGB2EtrCOaOvpR0izu\\nm7Pw/sruU+pA4k8bUnCEE3EwFfSKt71SHjL/NWzY7RxfY/BrFtWgwnvZOcuRevRH\\n4b01xv7qI44rdlWvQnWJW+c2kuQOyhxhuUqPMsRmzQW/4lgnSi9B1zCuWTF0Cai9\\nxIaJvpjsadl9zjd3zAdArQKBgGjRribv48jeJA5nVrHQo2VL35Ghl/zll/+XH8D/\\n/Wyh0VklH4hnKsw1nOece397t2yrBrI7ybN8lOZdwzp/SSJfqLiPOqG7MEbABMqQ\\nh+tYXrqpFrC1FHPFbHP6Nfgf6s5mWtNO6w9WL3hH0GMbGeG8Mjli22kTY/6sEXhJ\\nWwH5AoGADMqgNomT8bnNL/1luH5e2h+ug0XcujiigI05LCf0UyYHsswjJeew+T2G\\nOWtG2PLwMN0fkZJuu+ToQeYDWtUe6EOm9MTFoNJDwTGBktRhMlXGyLp2JK4hRm5Q\\n86Ko3fFFpLGePu7ZvVo1QZ8mS3bF40lURLxonY5NPjlBmhzVcdI=\\n-----END PRIVATE KEY-----\\n\",\"pubKey\":\"-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA73ZtAlS3y+mXpHXqbZ5x\\nfPGGeIl5VKcoTQOyGzD+tgODd9m1dlq1Qz0HSRUYlAaxg5c3hb12Ty0L3tLEhu02\\n+LIRrknnqazVS2+XXT2Y14WblYTdOOrnoEPD4Kfa3Lb74IoMJ50F/bSkHY6csjNm\\nJEaRSIun5tW+pTZhkKDYSBu2kNevHFV5k20iDNlsf2MlXhalXIG2IXzB5Z2JtYhq\\nEX9tQJ6INACnevD2N4AfFSX3tbWRsVHMC0v/wDdA/cgW/LP44mUEvT1cwt5QfEAg\\nyCEChp8+UFMsDec3gtHlyze1ZPItsg5Xq7bzYpKR1R7DQwSIIhY0+qkfpmdVTMVX\\n5QIDAQAB\\n-----END PUBLIC KEY-----\\n\"}}}', '', 0, 1709021277, 1709021277);
INSERT INTO `admin_user` (`id`, `parent_id`, `username`, `nickname`, `email`, `avatar`, `password`, `security_key`, `money`, `status`, `data`, `domains`, `expired_at`, `updated_at`, `created_at`) VALUES (2, 1, 'merchant', '八戒网络科技', 'muiprosperyls15@gmail.com', '', '57f231b1ec41dc6641270cb09a56f897', '57f231b1ec41dc6641270cb09a56f897', 0.00, 10, '{\"template\":\"default\",\"nums\":5}', '127.0.0.1:8089,localhost:8089,basic.ainn.us', 0, 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `type` smallint unsigned NOT NULL DEFAULT '10' COMMENT '类型 1帮助中心类型 2底部类型 10基础文章',
  `image` varchar(255) NOT NULL COMMENT '图片',
  `name` varchar(60) NOT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章表';

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 1, 2, '', 'aboutUsLabel', 'aboutUsDesc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 1, 2, '', 'serviceAgreementLabel', 'serviceAgreementDesc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (3, 1, 2, '', 'privacyAgreementLabel', 'privacyAgreementDesc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (4, 1, 3, '', 'service1Label', 'service1Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (5, 1, 3, '', 'service2Label', 'service2Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (6, 1, 3, '', 'service3Label', 'service3Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (7, 1, 1, '', 'helpers1Label', 'helpers1Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (8, 1, 1, '', 'helpers2Label', 'helpers2Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (9, 1, 1, '', 'helpers3Label', 'helpers3Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (10, 1, 1, '', 'helpers4Label', 'helpers4Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (11, 1, 1, '', 'helpers5Label', 'helpers5Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (12, 1, 1, '', 'helpers6Label', 'helpers6Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (13, 1, 1, '', 'helpers7Label', 'helpers7Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (14, 2, 2, '', 'aboutUsLabel', 'aboutUsDesc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (15, 2, 2, '', 'serviceAgreementLabel', 'serviceAgreementDesc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (16, 2, 2, '', 'privacyAgreementLabel', 'privacyAgreementDesc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (17, 2, 3, '', 'service1Label', 'service1Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (18, 2, 3, '', 'service2Label', 'service2Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (19, 2, 3, '', 'service3Label', 'service3Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (20, 2, 1, '', 'helpers1Label', 'helpers1Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (21, 2, 1, '', 'helpers2Label', 'helpers2Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (22, 2, 1, '', 'helpers3Label', 'helpers3Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (23, 2, 1, '', 'helpers4Label', 'helpers4Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (24, 2, 1, '', 'helpers5Label', 'helpers5Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (25, 2, 1, '', 'helpers6Label', 'helpers6Desc', 10, '', 1709021277, 1709021277);
INSERT INTO `article` (`id`, `admin_id`, `type`, `image`, `name`, `content`, `status`, `data`, `updated_at`, `created_at`) VALUES (26, 2, 1, '', 'helpers7Label', 'helpers7Desc', 10, '', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for auth_assignment
-- ----------------------------
DROP TABLE IF EXISTS `auth_assignment`;
CREATE TABLE `auth_assignment` (
  `name` varchar(50) NOT NULL COMMENT '名称',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  KEY `idx_admin_name` (`name`,`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限分配表';

-- ----------------------------
-- Records of auth_assignment
-- ----------------------------
BEGIN;
INSERT INTO `auth_assignment` (`name`, `admin_id`, `created_at`) VALUES ('超级管理员', 1, 1709021277);
INSERT INTO `auth_assignment` (`name`, `admin_id`, `created_at`) VALUES ('商户管理员', 2, 1709021277);
INSERT INTO `auth_assignment` (`name`, `admin_id`, `created_at`) VALUES ('代理管理员', 3, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for auth_child
-- ----------------------------
DROP TABLE IF EXISTS `auth_child`;
CREATE TABLE `auth_child` (
  `parent` varchar(50) NOT NULL COMMENT '父级',
  `child` varchar(50) NOT NULL COMMENT '子级',
  `type` tinyint NOT NULL COMMENT '类型 2路由名称对应路由 3角色对应路由名',
  KEY `idx_parent_child` (`parent`,`child`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限目录子集表';

-- ----------------------------
-- Records of auth_child
-- ----------------------------
BEGIN;
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('所有权限', '*', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('超级管理员', '所有权限', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('超级管理员', '商户管理员', 1);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '代理管理员', 1);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('首页信息', '/api/v1/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '首页信息', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '首页信息', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('初始化数据', '/api/v1/init', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '初始化数据', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '初始化数据', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('管理信息', '/api/v1/info', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '管理信息', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '管理信息', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('管理提示音', '/api/v1/audio', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '管理提示音', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '管理提示音', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新管理信息', '/api/v1/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新管理信息', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '更新管理信息', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('修改管理密码', '/api/v1/update/password', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '修改管理密码', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '修改管理密码', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('上传文件', '/api/v1/upload', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '上传文件', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '上传文件', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增用户', '/api/v1/user/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增用户', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '新增用户', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户列表', '/api/v1/user/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户更新', '/api/v1/user/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除用户', '/api/v1/user/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除用户', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '删除用户', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户余额变动', '/api/v1/user/balance', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户余额变动', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户余额变动', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户Json模版', '/api/v1/user/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户关系图', '/api/v1/user/relation', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户关系图', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户关系图', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('邀请列表', '/api/v1/user/invite/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '邀请列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '邀请列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增邀请', '/api/v1/user/invite/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增邀请', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '新增邀请', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新邀请', '/api/v1/user/invite/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新邀请', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '更新邀请', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除邀请', '/api/v1/user/invite/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除邀请', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '删除邀请', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('邀请Json模版', '/api/v1/user/invite/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '邀请Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '邀请Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('会员列表', '/api/v1/user/level/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '会员列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '会员列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增会员', '/api/v1/user/level/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增会员', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '新增会员', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新会员', '/api/v1/user/level/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新会员', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '更新会员', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除会员', '/api/v1/user/level/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除会员', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '删除会员', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('会员Json模版', '/api/v1/user/level/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '会员Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '会员Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('认证列表', '/api/v1/user/auth/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '认证列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '认证列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增认证', '/api/v1/user/auth/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增认证', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '新增认证', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新认证', '/api/v1/user/auth/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新认证', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '更新认证', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除认证', '/api/v1/user/auth/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除认证', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '删除认证', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('认证Json模版', '/api/v1/user/auth/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '认证Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '认证Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增国家配置', '/api/v1/system/country/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增国家配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('查询国家配置', '/api/v1/system/country/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '查询国家配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除国家配置', '/api/v1/system/country/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除国家配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新国家配置', '/api/v1/system/country/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新国家配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('国家配置Json模版', '/api/v1/system/country/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '国家配置Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增等级配置', '/api/v1/system/level/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增等级配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('查询等级配置', '/api/v1/system/level/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '查询等级配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除等级配置', '/api/v1/system/level/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除等级配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新等级配置', '/api/v1/system/level/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新等级配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('等级配置Json模版', '/api/v1/system/level/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '等级配置Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增语言配置', '/api/v1/system/lang/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增语言配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('查询语言配置', '/api/v1/system/lang/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '查询语言配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除语言配置', '/api/v1/system/lang/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除语言配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新语言配置', '/api/v1/system/lang/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新语言配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('语言配置Json模版', '/api/v1/system/lang/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '语言配置Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增翻译配置', '/api/v1/system/translate/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增翻译配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('查询翻译配置', '/api/v1/system/translate/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '查询翻译配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除翻译配置', '/api/v1/system/translate/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除翻译配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新翻译配置', '/api/v1/system/translate/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新翻译配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('查询字段语言', '/api/v1/system/translate/fields', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '查询字段语言', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '查询字段语言', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('Google翻译', '/api/v1/system/translate/google', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', 'Google翻译', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', 'Google翻译', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新字段语言', '/api/v1/system/translate/update/fields', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新字段语言', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '更新字段语言', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('翻译配置Json模版', '/api/v1/system/translate/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '翻译配置Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增通知配置', '/api/v1/system/notify/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增通知配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('查询通知配置', '/api/v1/system/notify/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '查询通知配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除通知配置', '/api/v1/system/notify/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除通知配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新通知配置', '/api/v1/system/notify/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新通知配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('通知配置Json模版', '/api/v1/system/notify/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '通知配置Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户充值审核', '/api/v1/wallets/deposit/status', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户充值审核', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户充值审核', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户充值新增', '/api/v1/wallets/deposit/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户充值新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户充值新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户充值删除', '/api/v1/wallets/deposit/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户充值删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户充值删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户充值更新', '/api/v1/wallets/deposit/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户充值更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户充值更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户充值列表', '/api/v1/wallets/deposit/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户充值列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户充值列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户充值Json模版', '/api/v1/wallets/deposit/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户充值Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户充值Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户账单列表', '/api/v1/wallets/bill/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户账单列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户账单列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户账单Json模版', '/api/v1/wallets/bill/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户账单Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户账单Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现账户新增', '/api/v1/wallets/account/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现账户新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现账户新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现账户删除', '/api/v1/wallets/account/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现账户删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现账户删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现账户更新', '/api/v1/wallets/account/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现账户更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现账户更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现账户列表', '/api/v1/wallets/account/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现账户列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现账户列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现账户Json模版', '/api/v1/wallets/account/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现账户Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现账户Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('支付创建', '/api/v1/wallets/payment/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '支付创建', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('支付删除', '/api/v1/wallets/payment/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '支付删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('支付更新', '/api/v1/wallets/payment/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '支付更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('支付更新数据', '/api/v1/wallets/payment/desc', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '支付更新数据', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('支付列表', '/api/v1/wallets/payment/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '支付列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('支付Json模版', '/api/v1/wallets/payment/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '支付Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现新增', '/api/v1/wallets/withdraw/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现删除', '/api/v1/wallets/withdraw/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现更新', '/api/v1/wallets/withdraw/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现列表', '/api/v1/wallets/withdraw/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现审核', '/api/v1/wallets/withdraw/status', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现审核', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现审核', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户提现Json模版', '/api/v1/wallets/withdraw/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户提现Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户提现Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('资产创建', '/api/v1/wallets/assets/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '资产创建', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('资产删除', '/api/v1/wallets/assets/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '资产删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('资产更新', '/api/v1/wallets/assets/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '资产更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('资产列表', '/api/v1/wallets/assets/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '资产列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('资产Json模版', '/api/v1/wallets/assets/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '资产Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户资产修改', '/api/v1/user/assets/balance', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户资产修改', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户资产修改', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户资产删除', '/api/v1/user/assets/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户资产删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户资产删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户资产列表', '/api/v1/user/assets/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户资产列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户资产列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户资产更新', '/api/v1/user/assets/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户资产更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户资产更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('用户资产json模版', '/api/v1/user/assets/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '用户资产json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('代理管理员', '用户资产json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('文章创建', '/api/v1/setting/article/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '文章创建', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('文章删除', '/api/v1/setting/article/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '文章删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('文章修改', '/api/v1/setting/article/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '文章修改', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('文章列表', '/api/v1/setting/article/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '文章列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('文章json模版', '/api/v1/setting/article/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '文章json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('设置创建', '/api/v1/setting/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '设置创建', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('设置删除', '/api/v1/setting/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '设置删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('设置修改', '/api/v1/setting/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '设置修改', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('设置修改详情', '/api/v1/setting/update/desc', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '设置修改详情', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('设置列表', '/api/v1/setting/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '设置列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('设置json模版', '/api/v1/setting/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '设置json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('菜单列表', '/api/v1/setting/menu/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '菜单列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('菜单新增', '/api/v1/setting/menu/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '菜单新增', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('菜单更新', '/api/v1/setting/menu/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '菜单更新', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('菜单更新详情', '/api/v1/setting/menu/update/desc', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '菜单更新详情', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('菜单删除', '/api/v1/setting/menu/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '菜单删除', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('菜单Json模板', '/api/v1/setting/menu/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '菜单Json模板', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('操作日志列表', '/api/v1/manage/logs/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '操作日志列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('操作日志Json模板', '/api/v1/manage/logs/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '操作日志Json模板', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增角色', '/api/v1/manage/role/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除角色', '/api/v1/manage/role/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('角色列表', '/api/v1/manage/role/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新角色信息', '/api/v1/manage/role/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('角色Json模版', '/api/v1/manage/role/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '角色Json模版', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('管理列表', '/api/v1/manage/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '管理列表', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('新增管理', '/api/v1/manage/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '新增管理', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('删除管理', '/api/v1/manage/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '删除管理', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新管理', '/api/v1/manage/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新管理', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('更新管理配置', '/api/v1/manage/update/desc', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '更新管理配置', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('重置商户数据', '/api/v1/manage/reset', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '重置商户数据', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('同步商户数据', '/api/v1/manage/sync', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '同步商户数据', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('管理Json模板', '/api/v1/manage/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '管理Json模板', 3);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('后台菜单列表', '/api/v1/manage/menu/index', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('后台菜单新增', '/api/v1/manage/menu/create', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('后台菜单更新', '/api/v1/manage/menu/update', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('后台菜单更新详情', '/api/v1/manage/menu/update/desc', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('后台菜单删除', '/api/v1/manage/menu/delete', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('后台菜单Json模板', '/api/v1/manage/menu/views', 2);
INSERT INTO `auth_child` (`parent`, `child`, `type`) VALUES ('商户管理员', '后台菜单Json模板', 3);
COMMIT;

-- ----------------------------
-- Table structure for auth_item
-- ----------------------------
DROP TABLE IF EXISTS `auth_item`;
CREATE TABLE `auth_item` (
  `name` varchar(50) NOT NULL COMMENT '名称',
  `type` tinyint NOT NULL COMMENT '类型 1权限角色 2路由权限 3路由名称',
  `desc` varchar(255) NOT NULL COMMENT '详情',
  `rule` varchar(255) NOT NULL COMMENT '规则',
  `data` varchar(255) NOT NULL COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限目录表';

-- ----------------------------
-- Records of auth_item
-- ----------------------------
BEGIN;
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('*', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/audio', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/info', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/init', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/logs/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/logs/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/menu/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/menu/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/menu/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/menu/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/menu/update/desc', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/menu/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/reset', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/role/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/role/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/role/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/role/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/role/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/sync', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/update/desc', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/manage/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/article/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/article/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/article/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/article/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/article/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/menu/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/menu/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/menu/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/menu/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/menu/update/desc', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/menu/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/update/desc', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/setting/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/country/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/country/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/country/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/country/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/country/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/lang/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/lang/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/lang/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/lang/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/lang/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/level/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/level/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/level/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/level/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/level/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/notify/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/notify/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/notify/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/notify/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/notify/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/fields', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/google', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/update/fields', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/system/translate/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/update/password', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/upload', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/assets/balance', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/assets/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/assets/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/assets/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/assets/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/auth/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/auth/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/auth/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/auth/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/auth/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/balance', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/invite/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/invite/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/invite/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/invite/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/invite/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/level/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/level/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/level/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/level/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/level/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/relation', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/user/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/account/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/account/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/account/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/account/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/account/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/assets/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/assets/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/assets/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/assets/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/assets/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/bill/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/bill/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/deposit/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/deposit/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/deposit/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/deposit/status', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/deposit/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/deposit/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/payment/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/payment/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/payment/desc', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/payment/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/payment/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/payment/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/withdraw/create', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/withdraw/delete', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/withdraw/index', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/withdraw/status', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/withdraw/update', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('/api/v1/wallets/withdraw/views', 2, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('Google翻译', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('上传文件', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('代理管理员', 1, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('会员Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('会员列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('修改管理密码', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('初始化数据', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除会员', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除国家配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除用户', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除等级配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除管理', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除翻译配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除角色', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除认证', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除语言配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除通知配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('删除邀请', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('同步商户数据', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('后台菜单Json模板', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('后台菜单列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('后台菜单删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('后台菜单新增', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('后台菜单更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('后台菜单更新详情', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('商户管理员', 1, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('国家配置Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('所有权限', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('操作日志Json模板', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('操作日志列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('支付Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('支付列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('支付创建', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('支付删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('支付更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('支付更新数据', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('文章json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('文章修改', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('文章列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('文章创建', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('文章删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增会员', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增国家配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增用户', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增等级配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增管理', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增翻译配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增角色', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增认证', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增语言配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增通知配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('新增邀请', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新会员', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新国家配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新字段语言', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新等级配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新管理', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新管理信息', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新管理配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新翻译配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新角色信息', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新认证', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新语言配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新通知配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('更新邀请', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('查询国家配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('查询字段语言', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('查询等级配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('查询翻译配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('查询语言配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('查询通知配置', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户余额变动', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户充值Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户充值列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户充值删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户充值审核', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户充值新增', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户充值更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户关系图', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现审核', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现新增', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现账户Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现账户列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现账户删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现账户新增', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户提现账户更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户账单Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户账单列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户资产json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户资产修改', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户资产列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户资产删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('用户资产更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('等级配置Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('管理Json模板', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('管理信息', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('管理列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('管理提示音', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('翻译配置Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('菜单Json模板', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('菜单列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('菜单删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('菜单新增', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('菜单更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('菜单更新详情', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('角色Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('角色列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('认证Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('认证列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('设置json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('设置修改', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('设置修改详情', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('设置列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('设置创建', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('设置删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('语言配置Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('资产Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('资产列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('资产创建', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('资产删除', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('资产更新', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('超级管理员', 1, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('通知配置Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('邀请Json模版', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('邀请列表', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('重置商户数据', 3, '', '', '', 1709021277, 1709021277);
INSERT INTO `auth_item` (`name`, `type`, `desc`, `rule`, `data`, `updated_at`, `created_at`) VALUES ('首页信息', 3, '', '', '', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for country
-- ----------------------------
DROP TABLE IF EXISTS `country`;
CREATE TABLE `country` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `alias` varchar(60) NOT NULL COMMENT '别名',
  `icon` varchar(60) NOT NULL COMMENT '图标',
  `iso1` varchar(60) NOT NULL COMMENT '二位代码',
  `sort` tinyint NOT NULL DEFAULT '99' COMMENT '排序',
  `code` varchar(50) NOT NULL COMMENT '区号',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=447 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统国家表';

-- ----------------------------
-- Records of country
-- ----------------------------
BEGIN;
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 1, '中国', 'China', '/assets/country/china.png', 'CN', 99, '86', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 1, '台湾', 'Taiwan', '/assets/country/taiwan.png', 'TW', 99, '886', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (3, 1, '香港', 'Hong kong', '/assets/country/hongkong.png', 'HK', 99, '852', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (4, 1, '澳门', 'Macao', '/assets/country/macao.png', 'MO', 99, '853', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (5, 1, '蒙古', 'Mongolia', '/assets/country/mongolia.png', 'MN', 99, '976', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (6, 1, '朝鲜', 'North Korea', '/assets/country/north_korea.png', 'KP', 99, '850', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (7, 1, '韩国', 'South Korea', '/assets/country/south_korea.png', 'KR', 99, '82', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (8, 1, '日本', 'Japan', '/assets/country/japan.png', 'JP', 99, '81', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (9, 1, '菲律宾', 'Philippines', '/assets/country/philippines.png', 'PH', 99, '63', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (10, 1, '越南', 'Vietnam', '/assets/country/vietnam.png', 'VN', 99, '84', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (11, 1, '老挝', 'Laos', '/assets/country/laos.png', 'LA', 99, '856', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (12, 1, '柬埔寨', 'Cambodia', '/assets/country/cambodia.png', 'kH', 99, '855', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (13, 1, '缅甸', 'Myanmar', '/assets/country/myanmar.png', 'MM', 99, '95', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (14, 1, '泰国', 'Thailand', '/assets/country/thailand.png', 'TH', 99, '66', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (15, 1, '马来西亚', 'Malaysia', '/assets/country/malaysia.png', 'MY', 99, '60', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (16, 1, '文莱', 'Brunei', '/assets/country/brunei.png', 'BN', 99, '673', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (17, 1, '新加坡', 'Singapore', '/assets/country/singapore.png', 'SG', 99, '65', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (18, 1, '印度尼西亚', 'Indonesia', '/assets/country/indonesia.png', 'ID', 99, '62', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (19, 1, '东帝汶', 'Timor-Leste', '/assets/country/east_timor.png', 'TL', 99, '670', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (20, 1, '尼泊尔', 'Nepal', '/assets/country/nepal.png', 'NP', 99, '977', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (21, 1, '不丹', 'Bhutan', '/assets/country/bhutan.png', 'BT', 99, '975', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (22, 1, '孟加拉国', 'Bangladesh', '/assets/country/bangladesh.png', 'BD', 99, '880', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (23, 1, '印度', 'India', '/assets/country/india.png', 'IN', 99, '91', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (24, 1, '巴基斯坦', 'Pakistan', '/assets/country/pakistan.png', 'PK', 99, '92', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (25, 1, '斯里兰卡', 'SriLanka', '/assets/country/sri_lanka.png', 'LK', 99, '94', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (26, 1, '马尔代夫', 'Maldives', '/assets/country/maldives.png', 'MV', 99, '960', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (27, 1, '哈萨克斯坦', 'Kazakhstan', '/assets/country/kazakhstan.png', 'KZ', 99, '7', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (28, 1, '吉尔吉斯斯坦', 'Kyrgyzstan', '/assets/country/kyrgyzstan.png', 'KG', 99, '996', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (29, 1, '塔吉克斯坦', 'Tajikistan', '/assets/country/tajikistan.png', 'TJ', 99, '992', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (30, 1, '乌兹别克斯坦', 'Uzbekistan', '/assets/country/uzbekistan.png', 'UZ', 99, '998', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (31, 1, '土库曼斯坦', 'Turkmenistan', '/assets/country/turkmenistan.png', 'TM', 99, '993', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (32, 1, '阿富汗', 'Afghanistan', '/assets/country/afghanistan.png', 'AF', 99, '93', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (33, 1, '伊拉克', 'Iraq', '/assets/country/iraq.png', 'IQ', 99, '964', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (34, 1, '伊朗', 'Iran', '/assets/country/iran.png', 'IR', 99, '98', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (35, 1, '叙利亚', 'Syria', '/assets/country/syria.png', 'SY', 99, '963', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (36, 1, '约旦', 'Jordan', '/assets/country/jordan.png', 'JO', 99, '962', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (37, 1, '黎巴嫩', 'Lebanon', '/assets/country/lebanon.png', 'LB', 99, '961', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (38, 1, '以色列', 'Israel', '/assets/country/israel.png', 'IL', 99, '972', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (39, 1, '巴勒斯坦', 'Palestine', '/assets/country/palestine.png', 'PS', 99, '970', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (40, 1, '沙特阿拉伯', 'SaudiArabia', '/assets/country/saudi_arabia.png', 'SA', 99, '966', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (41, 1, '巴林', 'Bahrain', '/assets/country/bahrain.png', 'BH', 99, '973', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (42, 1, '卡塔尔', 'Qatar', '/assets/country/qatar.png', 'QA', 99, '974', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (43, 1, '科威特', 'Kuwait', '/assets/country/kuwait.png', 'KW', 99, '965', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (44, 1, '阿拉伯联合酋长国', 'United Arab Emirates', '/assets/country/united_arab_emirates.png', 'AE', 99, '971', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (45, 1, '阿曼', 'Oman', '/assets/country/oman.png', 'OM', 99, '968', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (46, 1, '也门', 'Yemen', '/assets/country/yemen.png', 'YE', 99, '967', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (47, 1, '格鲁吉亚', 'Georgia', '/assets/country/georgia.png', 'GE', 99, '995', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (48, 1, '亚美尼亚', 'Armenia', '/assets/country/armenia.png', 'AM', 99, '374', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (49, 1, '阿塞拜疆', 'Azerbaijan', '/assets/country/azerbaijan.png', 'AZ', 99, '994', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (50, 1, '土耳其', 'Turkey', '/assets/country/turkey.png', 'TR', 99, '90', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (51, 1, '塞浦路斯', 'Cyprus', '/assets/country/cyprus.png', 'CY', 99, '357', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (52, 1, '英国', 'England', '/assets/country/england.png', 'GB', 99, '44', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (53, 1, '爱尔兰', 'Ireland', '/assets/country/ireland.png', 'IE', 99, '353', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (54, 1, '法国', 'France', '/assets/country/france.png', 'FR', 99, '33', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (55, 1, '荷兰', 'Netherlands', '/assets/country/netherlands.png', 'NL', 99, '31', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (56, 1, '比利时', 'Belgium', '/assets/country/belgium.png', 'BE', 99, '32', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (57, 1, '卢森堡', 'Luxembourg', '/assets/country/luxembourg.png', 'LU', 99, '352', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (58, 1, '摩纳哥', 'Monaco', '/assets/country/monaco.png', 'MC', 99, '377', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (59, 1, '挪威', 'Norway', '/assets/country/norway.png', 'NO', 99, '47', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (60, 1, '瑞典', 'Sweden', '/assets/country/sweden.png', 'SE', 99, '46', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (61, 1, '芬兰', 'Finland', '/assets/country/finland.png', 'FI', 99, '358', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (62, 1, '丹麦', 'Denmark', '/assets/country/denmark.png', 'DK', 99, '45', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (63, 1, '冰岛', 'Iceland', '/assets/country/iceland.png', 'IS', 99, '354', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (64, 1, '法罗群岛', 'Faroe Islands', '/assets/country/faroese.png', 'FO', 99, '298', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (65, 1, '德国', 'Germany', '/assets/country/germany.png', 'DE', 99, '49', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (66, 1, '波兰', 'Poland', '/assets/country/poland.png', 'PL', 99, '48', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (67, 1, '捷克', 'Czechia', '/assets/country/czech.png', 'CZ', 99, '420', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (68, 1, '斯洛伐克', 'Slovakia', '/assets/country/slovakia.png', 'SK', 99, '421', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (69, 1, '匈牙利', 'Hungary', '/assets/country/hungary.png', 'HU', 99, '36', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (70, 1, '奥地利', 'Austria', '/assets/country/austria.png', 'AT', 99, '43', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (71, 1, '列支敦士登', 'Liechtenstein', '/assets/country/liechtenstein.png', 'LI', 99, '423', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (72, 1, '瑞士', 'Switzerland', '/assets/country/switzerland.png', 'CH', 99, '41', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (73, 1, '俄罗斯', 'Russia', '/assets/country/russia.png', 'RU', 99, '7', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (74, 1, '爱沙尼亚', 'Estonia', '/assets/country/estonia.png', 'EE', 99, '372', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (75, 1, '拉脱维亚', 'Latvia', '/assets/country/latvia.png', 'LV', 99, '371', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (76, 1, '立陶宛', 'Lithuania', '/assets/country/lithuania.png', 'LT', 99, '370', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (77, 1, '白俄罗斯', 'Belarus', '/assets/country/belarus.png', 'BY', 99, '375', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (78, 1, '乌克兰', 'Ukraine', '/assets/country/ukraine.png', 'UA', 99, '380', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (79, 1, '摩尔多瓦', 'Moldova', '/assets/country/moldova.png', 'MD', 99, '373', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (80, 1, '罗马尼亚', 'Romania', '/assets/country/romania.png', 'RO', 99, '40', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (81, 1, '保加利亚', 'Bulgaria', '/assets/country/bulgaria.png', 'BG', 99, '359', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (82, 1, '塞尔维亚', 'Serbia', '/assets/country/serbia.png', 'RS', 99, '338', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (83, 1, '斯洛文尼亚', 'Slovenia', '/assets/country/slovenia.png', 'SI', 99, '386', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (84, 1, '克罗地亚', 'Croatia', '/assets/country/croatia.png', 'HR', 99, '385', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (85, 1, '波斯尼亚和黑塞哥维那', 'Bosnia And Herzegovina', '/assets/country/bosnia_and_herzegovina.png', 'BA', 99, '387', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (86, 1, '马其顿', 'Macedonia', '/assets/country/macedonia.png', 'MK', 99, '389', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (87, 1, '阿尔巴尼亚', 'Albania', '/assets/country/albania.png', 'AL', 99, '355', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (88, 1, '希腊', 'Greece', '/assets/country/greece.png', 'GR', 99, '30', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (89, 1, '意大利', 'Italy', '/assets/country/italy.png', 'IT', 99, '39', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (90, 1, '圣马力诺', 'San Marino', '/assets/country/san_marino.png', 'SM', 99, '378', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (91, 1, '梵蒂冈', 'Vatican', '/assets/country/vatican.png', 'VA', 99, '379', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (92, 1, '葡萄牙', 'Portugal', '/assets/country/portugal.png', 'PT', 99, '351', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (93, 1, '西班牙', 'Spain', '/assets/country/spain.png', 'ES', 99, '34', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (94, 1, '安道尔', 'Andorra', '/assets/country/andorra.png', 'AD', 99, '376', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (95, 1, '马耳他', 'Malta', '/assets/country/malta.png', 'MT', 99, '356', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (96, 1, '埃及', 'Egypt', '/assets/country/egypt.png', 'EG', 99, '20', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (97, 1, '利比亚', 'Libya', '/assets/country/libya.png', 'LY', 99, '218', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (98, 1, '苏丹', 'Sudan', '/assets/country/sudan.png', 'SD', 99, '249', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (99, 1, '突尼斯', 'Tunisia', '/assets/country/tunisia.png', 'TN', 99, '216', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (100, 1, '阿尔及利亚', 'Algeria', '/assets/country/algeria.png', 'DZ', 99, '213', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (101, 1, '摩洛哥', 'Morocco', '/assets/country/morocco.png', 'MA', 99, '212', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (102, 1, '埃塞俄比亚', 'Ethiopia', '/assets/country/ethiopia.png', 'ET', 99, '251', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (103, 1, '厄立特里亚', 'Eritrea', '/assets/country/eritrea.png', 'ER', 99, '291', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (104, 1, '索马里', 'Somalia', '/assets/country/somalia.png', 'SO', 99, '252', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (105, 1, '吉布提', 'Djibouti', '/assets/country/djibouti.png', 'DJ', 99, '253', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (106, 1, '肯尼亚', 'Kenya', '/assets/country/kenya.png', 'KE', 99, '254', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (107, 1, '坦桑尼亚', 'Tanzania', '/assets/country/tasania.png', 'TZ', 99, '255', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (108, 1, '乌干达', 'Uganda', '/assets/country/uganda.png', 'UG', 99, '256', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (109, 1, '卢旺达', 'Rwanda', '/assets/country/rwanda.png', 'RW', 99, '250', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (110, 1, '布隆迪', 'Burundi', '/assets/country/burundi.png', 'BI', 99, '257', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (111, 1, '塞舌尔', 'Seychelles', '/assets/country/seychelles.png', 'SC', 99, '248', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (112, 1, '乍得', 'Chad', '/assets/country/chad.png', 'TD', 99, '235', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (113, 1, '中非', 'Central Africa', '/assets/country/central_africa.png', 'CF', 99, '236', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (114, 1, '喀麦隆', 'Cameroon', '/assets/country/cameroon.png', 'CM', 99, '237', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (115, 1, '赤道几内亚', 'Equatorial Guinea', '/assets/country/equatorial_guinea.png', 'GQ', 99, '240', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (116, 1, '加蓬', 'Gabon', '/assets/country/gabon.png', 'GA', 99, '241', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (117, 1, '刚果', 'Congo', '/assets/country/congo.png', 'CG', 99, '242', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (118, 1, '圣多美和普林西比', 'Sao Tome And Principe', '/assets/country/sao_tome_and_principe.png', 'ST', 99, '239', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (119, 1, '毛里塔尼亚', 'Mauritania', '/assets/country/mauritania.png', 'MR', 99, '222', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (120, 1, '塞内加尔', 'Senegal', '/assets/country/senegal.png', 'SN', 99, '221', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (121, 1, '冈比亚', 'Gambia', '/assets/country/gambia.png', 'GM', 99, '220', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (122, 1, '马里', 'Mali', '/assets/country/mali.png', 'ML', 99, '223', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (123, 1, '布基纳法索', 'Burkina Faso', '/assets/country/burkina_faso.png', 'BF', 99, '226', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (124, 1, '几内亚', 'Guinea', '/assets/country/guinea.png', 'GN', 99, '224', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (125, 1, '几内亚比绍', 'Guinea Bissau', '/assets/country/guinea_bissau.png', 'GW', 99, '245', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (126, 1, '佛得角', 'Cape Verde', '/assets/country/cape_verde.png', 'CV', 99, '238', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (127, 1, '塞拉利昂', 'Sierra Leone', '/assets/country/sierra_leone.png', 'SL', 99, '232', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (128, 1, '利比里亚', 'Liberia', '/assets/country/liberia.png', 'LR', 99, '231', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (129, 1, '科特迪瓦', 'Ivory Coast', '/assets/country/ivory_coast.png', 'CI', 99, '225', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (130, 1, '加纳', 'Ghana', '/assets/country/ghana.png', 'GH', 99, '233', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (131, 1, '多哥', 'Togo', '/assets/country/togo.png', 'TG', 99, '228', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (132, 1, '贝宁', 'Benin', '/assets/country/benin.png', 'BJ', 99, '229', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (133, 1, '尼日尔', 'Niger', '/assets/country/niger.png', 'NE', 99, '227', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (134, 1, '尼日利亚', 'Nigeria', '/assets/country/nigeria.png', 'NG', 99, '234', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (135, 1, '西撒哈拉', 'WesternSahara', '/assets/country/western_sahara.png', 'EH', 99, '210', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (136, 1, '赞比亚', 'Zambia', '/assets/country/zambia.png', 'ZM', 99, '260', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (137, 1, '安哥拉', 'Angola', '/assets/country/angola.png', 'AO', 99, '244', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (138, 1, '津巴布韦', 'Zimbabwe', '/assets/country/zimbabwe.png', 'ZW', 99, '263', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (139, 1, '马拉维', 'Malawi', '/assets/country/malawi.png', 'MW', 99, '265', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (140, 1, '莫桑比克', 'Mozambique', '/assets/country/mozambique.png', 'MZ', 99, '258', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (141, 1, '博茨瓦纳', 'Botswana', '/assets/country/botswana.png', 'BW', 99, '267', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (142, 1, '纳米比亚', 'Namibia', '/assets/country/namibia.png', 'NA', 99, '264', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (143, 1, '南非', 'SouthAfrica', '/assets/country/south_africa.png', 'ZA', 99, '27', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (144, 1, '斯威士兰', 'Eswatini', '/assets/country/swaziland.png', 'SZ', 99, '268', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (145, 1, '莱索托', 'Lesotho', '/assets/country/Lesotho.png', 'LS', 99, '266', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (146, 1, '马达加斯加', 'Madagascar', '/assets/country/madagascar.png', 'MG', 99, '261', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (147, 1, '科摩罗', 'Comoros', '/assets/country/comoros.png', 'KM', 99, '269', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (148, 1, '毛里求斯', 'Mauritius', '/assets/country/Mauritius.png', 'MU', 99, '230', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (149, 1, '圣赫勒拿岛', 'Saint Helena', '/assets/country/st_helena.png', 'SH', 99, '290', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (150, 1, '马约特', 'Mayotte', '/assets/country/mayotte.png', 'YT', 99, '262', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (151, 1, '加拿大', 'Canada', '/assets/country/canada.png', 'CA', 99, '1', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (152, 1, '美国', 'United States', '/assets/country/usa.png', 'US', 99, '1', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (153, 1, '墨西哥', 'Mexico', '/assets/country/mexico.png', 'MX', 99, '52', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (154, 1, '格陵兰', 'Greenland', '/assets/country/greenland.png', 'GL', 99, '299', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (155, 1, '圣皮埃尔和密克隆', 'Saint Pierre and Miquelon', '/assets/country/saint_pierre_and_miquelon.png', 'PM', 99, '508', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (156, 1, '百慕大', 'Bermuda', '/assets/country/bermuda.png', 'BM', 99, '441', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (157, 1, '危地马拉', 'Guatemala', '/assets/country/guatemala.png', 'GT', 99, '502', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (158, 1, '伯利兹', 'Belize', '/assets/country/belize.png', 'BZ', 99, '501', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (159, 1, '萨尔瓦多', 'Salvador', '/assets/country/salvador.png', 'SV', 99, '503', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (160, 1, '洪都拉斯', 'Honduras', '/assets/country/honduras.png', 'HN', 99, '504', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (161, 1, '尼加拉瓜', 'Nicaragua', '/assets/country/nicaragua.png', 'NI', 99, '505', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (162, 1, '哥斯达黎加', 'CostaRica', '/assets/country/costa_rica.png', 'CR', 99, '506', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (163, 1, '巴拿马', 'Panama', '/assets/country/panama.png', 'PA', 99, '507', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (164, 1, '巴哈马', 'Bahamas', '/assets/country/bahamas.png', 'BS', 99, '1242', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (165, 1, '古巴', 'Cuba', '/assets/country/cuba.png', 'CU', 99, '53', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (166, 1, '牙买加', 'Jamaica', '/assets/country/jamaica.png', 'JM', 99, '1876', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (167, 1, '海地', 'Haiti', '/assets/country/haiti.png', 'HT', 99, '509', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (168, 1, '多米尼加共和国', 'Dominican Republic', '/assets/country/dominica_republic.png', 'DO', 99, '1809', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (169, 1, '安提瓜和巴布达', 'Antigua And Barbuda', '/assets/country/antigua_and_barbuda.png', 'AG', 99, '1268', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (170, 1, '圣基茨和尼维斯', 'Saint Kitts And Nevis', '/assets/country/saint_kitts_and_nevis.png', 'KN', 99, '1869', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (171, 1, '多米尼克', 'Dominica', '/assets/country/dominica.png', 'DM', 99, '1767', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (172, 1, '圣卢西亚', 'SaintLucia', '/assets/country/saint_lucia.png', 'LC', 99, '1758', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (173, 1, '圣文森特和格林纳丁斯', 'Saint Vincent And The Grenadines', '/assets/country/saint_vincent_and_the_grenadines.png', 'VC', 99, '784', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (174, 1, '格林纳达', 'Grenada', '/assets/country/grenada.png', 'GD', 99, '1473', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (175, 1, '巴巴多斯', 'Barbados', '/assets/country/barbados.png', 'BB', 99, '1246', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (176, 1, '特立尼达和多巴哥', 'Trinidad And Tobago', '/assets/country/trinidad_and_tobago.png', 'TT', 99, '1868', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (177, 1, '波多黎各', 'Puerto Rico', '/assets/country/puerto_rico.png', 'PR', 99, '1787', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (178, 1, '英属维尔京群岛', 'The British Virgin Islands', '/assets/country/the_british_virgin_islands.png', 'VG', 99, '1284', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (179, 1, '美属维尔京群岛', 'United States Virgin Islands', '/assets/country/united_states_virgin_islands.png', 'VI', 99, '1340', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (180, 1, '安圭拉', 'Anguilla', '/assets/country/anguilla.png', 'AI', 99, '264', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (181, 1, '蒙特塞拉特', 'Montserrat', '/assets/country/montserrat.png', 'MS', 99, '1664', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (182, 1, '瓜德罗普', 'Guadeloupe', '/assets/country/guadeloupe.png', 'GP', 99, '590', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (183, 1, '马提尼克', 'Martinique', '/assets/country/martinique.png', 'MQ', 99, '596', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (184, 1, '阿鲁巴', 'Aruba', '/assets/country/aruba.png', 'AW', 99, '297', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (185, 1, '法属圣马丁', 'Saint Martin', '/assets/country/st_martin.png', 'MF', 99, '1721', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (186, 1, '特克斯和凯科斯群岛', 'Turks And Caicos Islands', '/assets/country/turks_and_caicos_islands.png', 'TC', 99, '1649', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (187, 1, '开曼群岛', 'Cayman Islands', '/assets/country/cayman_islands.png', 'KY', 99, '1345', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (188, 1, '库拉索', 'Curacao', '/assets/country/curacao.png', 'CW', 99, '5999', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (189, 1, '哥伦比亚', 'Colombia', '/assets/country/colombia.png', 'CO', 99, '57', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (190, 1, '委内瑞拉', 'Venezuela', '/assets/country/venezuela.png', 'VE', 99, '58', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (191, 1, '圭亚那', 'Guyana', '/assets/country/guyana.png', 'GY', 99, '592', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (192, 1, '苏里南', 'Suriname', '/assets/country/suriname.png', 'SR', 99, '597', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (193, 1, '法属圭亚那', 'French Guiana', '/assets/country/french_guiana.png', 'GF', 99, '594', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (194, 1, '厄瓜多尔', 'Ecuador', '/assets/country/ecuador.png', 'EC', 99, '593', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (195, 1, '秘鲁', 'Peru', '/assets/country/peru.png', 'PE', 99, '51', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (196, 1, '玻利维亚', 'Bolivia', '/assets/country/bolivia.png', 'BO', 99, '591', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (197, 1, '巴西', 'Brazil', '/assets/country/brazil.png', 'BR', 99, '55', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (198, 1, '智利', 'Chile', '/assets/country/chile.png', 'CL', 99, '56', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (199, 1, '阿根廷', 'Argentina', '/assets/country/argentina.png', 'AR', 99, '54', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (200, 1, '乌拉圭', 'Uruguay', '/assets/country/uruguay.png', 'UY', 99, '598', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (201, 1, '巴拉圭', 'Paraguay', '/assets/country/paraguay.png', 'PY', 99, '595', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (202, 1, '澳大利亚', 'Australia', '/assets/country/australia.png', 'AU', 99, '61', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (203, 1, '新西兰', 'New Zealand', '/assets/country/new_zealand.png', 'NZ', 99, '64', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (204, 1, '帕劳', 'Palau', '/assets/country/palau.png', 'PW', 99, '680', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (205, 1, '密克罗尼西亚联邦', 'Federated States Of Micronesia', '/assets/country/federated_states_of_micronesia.png', 'FM', 99, '691', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (206, 1, '马绍尔群岛', 'Marshall Islands', '/assets/country/marshall_islands.png', 'MH', 99, '692', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (207, 1, '基里巴斯', 'Kiribati', '/assets/country/kiribati.png', 'KI', 99, '686', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (208, 1, '瑙鲁', 'Nauru', '/assets/country/nauru.png', 'NR', 99, '674', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (209, 1, '北马里亚纳', 'Northern Mariana', '/assets/country/northern_mariana.png', 'MP', 99, '670', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (210, 1, '关岛', 'Guam', '/assets/country/guam.png', 'GU', 99, '1671', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (211, 1, '巴布亚新几内亚', 'Papua New Guinea', '/assets/country/papua_new_guinea.png', 'PG', 99, '675', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (212, 1, '所罗门群岛', 'SolomonIslands', '/assets/country/solomon_islands.png', 'SB', 99, '677', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (213, 1, '瓦努阿图', 'Vanuatu', '/assets/country/vanuatu.png', 'VU', 99, '678', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (214, 1, '斐济', 'Fiji', '/assets/country/fiji.png', 'FJ', 99, '679', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (215, 1, '图瓦卢', 'Tuvalu', '/assets/country/tuvalu.png', 'TV', 99, '688', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (216, 1, '萨摩亚', 'Samoa', '/assets/country/samoa.png', 'WS', 99, '685', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (217, 1, '汤加', 'Tonga', '/assets/country/tonga.png', 'TO', 99, '676', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (218, 1, '库克群岛', 'Cook Islands', '/assets/country/island.png', 'CK', 99, '682', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (219, 1, '纽埃', 'Niue', '/assets/country/niue.png', 'NU', 99, '683', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (220, 1, '托克劳', 'Tokelau', '/assets/country/tokelau.png', 'TK', 99, '690', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (221, 1, '瓦利斯和富图纳', 'Wallis And Futuna', '/assets/country/wallis_and_futuna.png', 'WF', 99, '681', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (222, 1, '皮特凯恩群岛', 'Pitcairn Islands', '/assets/country/pitcairn_islands.png', 'PN', 99, '649', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (223, 1, '美属萨摩亚', 'American Samoa', '/assets/country/american_samoa.png', 'AS', 99, '684', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (224, 2, '中国', 'China', '/assets/country/china.png', 'CN', 99, '86', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (225, 2, '台湾', 'Taiwan', '/assets/country/taiwan.png', 'TW', 99, '886', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (226, 2, '香港', 'Hong kong', '/assets/country/hongkong.png', 'HK', 99, '852', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (227, 2, '澳门', 'Macao', '/assets/country/macao.png', 'MO', 99, '853', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (228, 2, '蒙古', 'Mongolia', '/assets/country/mongolia.png', 'MN', 99, '976', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (229, 2, '朝鲜', 'North Korea', '/assets/country/north_korea.png', 'KP', 99, '850', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (230, 2, '韩国', 'South Korea', '/assets/country/south_korea.png', 'KR', 99, '82', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (231, 2, '日本', 'Japan', '/assets/country/japan.png', 'JP', 99, '81', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (232, 2, '菲律宾', 'Philippines', '/assets/country/philippines.png', 'PH', 99, '63', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (233, 2, '越南', 'Vietnam', '/assets/country/vietnam.png', 'VN', 99, '84', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (234, 2, '老挝', 'Laos', '/assets/country/laos.png', 'LA', 99, '856', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (235, 2, '柬埔寨', 'Cambodia', '/assets/country/cambodia.png', 'kH', 99, '855', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (236, 2, '缅甸', 'Myanmar', '/assets/country/myanmar.png', 'MM', 99, '95', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (237, 2, '泰国', 'Thailand', '/assets/country/thailand.png', 'TH', 99, '66', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (238, 2, '马来西亚', 'Malaysia', '/assets/country/malaysia.png', 'MY', 99, '60', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (239, 2, '文莱', 'Brunei', '/assets/country/brunei.png', 'BN', 99, '673', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (240, 2, '新加坡', 'Singapore', '/assets/country/singapore.png', 'SG', 99, '65', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (241, 2, '印度尼西亚', 'Indonesia', '/assets/country/indonesia.png', 'ID', 99, '62', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (242, 2, '东帝汶', 'Timor-Leste', '/assets/country/east_timor.png', 'TL', 99, '670', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (243, 2, '尼泊尔', 'Nepal', '/assets/country/nepal.png', 'NP', 99, '977', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (244, 2, '不丹', 'Bhutan', '/assets/country/bhutan.png', 'BT', 99, '975', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (245, 2, '孟加拉国', 'Bangladesh', '/assets/country/bangladesh.png', 'BD', 99, '880', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (246, 2, '印度', 'India', '/assets/country/india.png', 'IN', 99, '91', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (247, 2, '巴基斯坦', 'Pakistan', '/assets/country/pakistan.png', 'PK', 99, '92', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (248, 2, '斯里兰卡', 'SriLanka', '/assets/country/sri_lanka.png', 'LK', 99, '94', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (249, 2, '马尔代夫', 'Maldives', '/assets/country/maldives.png', 'MV', 99, '960', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (250, 2, '哈萨克斯坦', 'Kazakhstan', '/assets/country/kazakhstan.png', 'KZ', 99, '7', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (251, 2, '吉尔吉斯斯坦', 'Kyrgyzstan', '/assets/country/kyrgyzstan.png', 'KG', 99, '996', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (252, 2, '塔吉克斯坦', 'Tajikistan', '/assets/country/tajikistan.png', 'TJ', 99, '992', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (253, 2, '乌兹别克斯坦', 'Uzbekistan', '/assets/country/uzbekistan.png', 'UZ', 99, '998', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (254, 2, '土库曼斯坦', 'Turkmenistan', '/assets/country/turkmenistan.png', 'TM', 99, '993', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (255, 2, '阿富汗', 'Afghanistan', '/assets/country/afghanistan.png', 'AF', 99, '93', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (256, 2, '伊拉克', 'Iraq', '/assets/country/iraq.png', 'IQ', 99, '964', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (257, 2, '伊朗', 'Iran', '/assets/country/iran.png', 'IR', 99, '98', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (258, 2, '叙利亚', 'Syria', '/assets/country/syria.png', 'SY', 99, '963', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (259, 2, '约旦', 'Jordan', '/assets/country/jordan.png', 'JO', 99, '962', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (260, 2, '黎巴嫩', 'Lebanon', '/assets/country/lebanon.png', 'LB', 99, '961', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (261, 2, '以色列', 'Israel', '/assets/country/israel.png', 'IL', 99, '972', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (262, 2, '巴勒斯坦', 'Palestine', '/assets/country/palestine.png', 'PS', 99, '970', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (263, 2, '沙特阿拉伯', 'SaudiArabia', '/assets/country/saudi_arabia.png', 'SA', 99, '966', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (264, 2, '巴林', 'Bahrain', '/assets/country/bahrain.png', 'BH', 99, '973', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (265, 2, '卡塔尔', 'Qatar', '/assets/country/qatar.png', 'QA', 99, '974', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (266, 2, '科威特', 'Kuwait', '/assets/country/kuwait.png', 'KW', 99, '965', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (267, 2, '阿拉伯联合酋长国', 'United Arab Emirates', '/assets/country/united_arab_emirates.png', 'AE', 99, '971', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (268, 2, '阿曼', 'Oman', '/assets/country/oman.png', 'OM', 99, '968', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (269, 2, '也门', 'Yemen', '/assets/country/yemen.png', 'YE', 99, '967', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (270, 2, '格鲁吉亚', 'Georgia', '/assets/country/georgia.png', 'GE', 99, '995', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (271, 2, '亚美尼亚', 'Armenia', '/assets/country/armenia.png', 'AM', 99, '374', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (272, 2, '阿塞拜疆', 'Azerbaijan', '/assets/country/azerbaijan.png', 'AZ', 99, '994', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (273, 2, '土耳其', 'Turkey', '/assets/country/turkey.png', 'TR', 99, '90', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (274, 2, '塞浦路斯', 'Cyprus', '/assets/country/cyprus.png', 'CY', 99, '357', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (275, 2, '英国', 'England', '/assets/country/england.png', 'GB', 99, '44', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (276, 2, '爱尔兰', 'Ireland', '/assets/country/ireland.png', 'IE', 99, '353', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (277, 2, '法国', 'France', '/assets/country/france.png', 'FR', 99, '33', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (278, 2, '荷兰', 'Netherlands', '/assets/country/netherlands.png', 'NL', 99, '31', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (279, 2, '比利时', 'Belgium', '/assets/country/belgium.png', 'BE', 99, '32', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (280, 2, '卢森堡', 'Luxembourg', '/assets/country/luxembourg.png', 'LU', 99, '352', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (281, 2, '摩纳哥', 'Monaco', '/assets/country/monaco.png', 'MC', 99, '377', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (282, 2, '挪威', 'Norway', '/assets/country/norway.png', 'NO', 99, '47', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (283, 2, '瑞典', 'Sweden', '/assets/country/sweden.png', 'SE', 99, '46', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (284, 2, '芬兰', 'Finland', '/assets/country/finland.png', 'FI', 99, '358', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (285, 2, '丹麦', 'Denmark', '/assets/country/denmark.png', 'DK', 99, '45', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (286, 2, '冰岛', 'Iceland', '/assets/country/iceland.png', 'IS', 99, '354', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (287, 2, '法罗群岛', 'Faroe Islands', '/assets/country/faroese.png', 'FO', 99, '298', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (288, 2, '德国', 'Germany', '/assets/country/germany.png', 'DE', 99, '49', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (289, 2, '波兰', 'Poland', '/assets/country/poland.png', 'PL', 99, '48', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (290, 2, '捷克', 'Czechia', '/assets/country/czech.png', 'CZ', 99, '420', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (291, 2, '斯洛伐克', 'Slovakia', '/assets/country/slovakia.png', 'SK', 99, '421', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (292, 2, '匈牙利', 'Hungary', '/assets/country/hungary.png', 'HU', 99, '36', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (293, 2, '奥地利', 'Austria', '/assets/country/austria.png', 'AT', 99, '43', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (294, 2, '列支敦士登', 'Liechtenstein', '/assets/country/liechtenstein.png', 'LI', 99, '423', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (295, 2, '瑞士', 'Switzerland', '/assets/country/switzerland.png', 'CH', 99, '41', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (296, 2, '俄罗斯', 'Russia', '/assets/country/russia.png', 'RU', 99, '7', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (297, 2, '爱沙尼亚', 'Estonia', '/assets/country/estonia.png', 'EE', 99, '372', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (298, 2, '拉脱维亚', 'Latvia', '/assets/country/latvia.png', 'LV', 99, '371', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (299, 2, '立陶宛', 'Lithuania', '/assets/country/lithuania.png', 'LT', 99, '370', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (300, 2, '白俄罗斯', 'Belarus', '/assets/country/belarus.png', 'BY', 99, '375', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (301, 2, '乌克兰', 'Ukraine', '/assets/country/ukraine.png', 'UA', 99, '380', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (302, 2, '摩尔多瓦', 'Moldova', '/assets/country/moldova.png', 'MD', 99, '373', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (303, 2, '罗马尼亚', 'Romania', '/assets/country/romania.png', 'RO', 99, '40', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (304, 2, '保加利亚', 'Bulgaria', '/assets/country/bulgaria.png', 'BG', 99, '359', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (305, 2, '塞尔维亚', 'Serbia', '/assets/country/serbia.png', 'RS', 99, '338', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (306, 2, '斯洛文尼亚', 'Slovenia', '/assets/country/slovenia.png', 'SI', 99, '386', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (307, 2, '克罗地亚', 'Croatia', '/assets/country/croatia.png', 'HR', 99, '385', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (308, 2, '波斯尼亚和黑塞哥维那', 'Bosnia And Herzegovina', '/assets/country/bosnia_and_herzegovina.png', 'BA', 99, '387', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (309, 2, '马其顿', 'Macedonia', '/assets/country/macedonia.png', 'MK', 99, '389', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (310, 2, '阿尔巴尼亚', 'Albania', '/assets/country/albania.png', 'AL', 99, '355', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (311, 2, '希腊', 'Greece', '/assets/country/greece.png', 'GR', 99, '30', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (312, 2, '意大利', 'Italy', '/assets/country/italy.png', 'IT', 99, '39', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (313, 2, '圣马力诺', 'San Marino', '/assets/country/san_marino.png', 'SM', 99, '378', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (314, 2, '梵蒂冈', 'Vatican', '/assets/country/vatican.png', 'VA', 99, '379', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (315, 2, '葡萄牙', 'Portugal', '/assets/country/portugal.png', 'PT', 99, '351', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (316, 2, '西班牙', 'Spain', '/assets/country/spain.png', 'ES', 99, '34', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (317, 2, '安道尔', 'Andorra', '/assets/country/andorra.png', 'AD', 99, '376', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (318, 2, '马耳他', 'Malta', '/assets/country/malta.png', 'MT', 99, '356', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (319, 2, '埃及', 'Egypt', '/assets/country/egypt.png', 'EG', 99, '20', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (320, 2, '利比亚', 'Libya', '/assets/country/libya.png', 'LY', 99, '218', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (321, 2, '苏丹', 'Sudan', '/assets/country/sudan.png', 'SD', 99, '249', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (322, 2, '突尼斯', 'Tunisia', '/assets/country/tunisia.png', 'TN', 99, '216', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (323, 2, '阿尔及利亚', 'Algeria', '/assets/country/algeria.png', 'DZ', 99, '213', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (324, 2, '摩洛哥', 'Morocco', '/assets/country/morocco.png', 'MA', 99, '212', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (325, 2, '埃塞俄比亚', 'Ethiopia', '/assets/country/ethiopia.png', 'ET', 99, '251', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (326, 2, '厄立特里亚', 'Eritrea', '/assets/country/eritrea.png', 'ER', 99, '291', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (327, 2, '索马里', 'Somalia', '/assets/country/somalia.png', 'SO', 99, '252', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (328, 2, '吉布提', 'Djibouti', '/assets/country/djibouti.png', 'DJ', 99, '253', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (329, 2, '肯尼亚', 'Kenya', '/assets/country/kenya.png', 'KE', 99, '254', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (330, 2, '坦桑尼亚', 'Tanzania', '/assets/country/tasania.png', 'TZ', 99, '255', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (331, 2, '乌干达', 'Uganda', '/assets/country/uganda.png', 'UG', 99, '256', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (332, 2, '卢旺达', 'Rwanda', '/assets/country/rwanda.png', 'RW', 99, '250', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (333, 2, '布隆迪', 'Burundi', '/assets/country/burundi.png', 'BI', 99, '257', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (334, 2, '塞舌尔', 'Seychelles', '/assets/country/seychelles.png', 'SC', 99, '248', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (335, 2, '乍得', 'Chad', '/assets/country/chad.png', 'TD', 99, '235', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (336, 2, '中非', 'Central Africa', '/assets/country/central_africa.png', 'CF', 99, '236', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (337, 2, '喀麦隆', 'Cameroon', '/assets/country/cameroon.png', 'CM', 99, '237', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (338, 2, '赤道几内亚', 'Equatorial Guinea', '/assets/country/equatorial_guinea.png', 'GQ', 99, '240', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (339, 2, '加蓬', 'Gabon', '/assets/country/gabon.png', 'GA', 99, '241', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (340, 2, '刚果', 'Congo', '/assets/country/congo.png', 'CG', 99, '242', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (341, 2, '圣多美和普林西比', 'Sao Tome And Principe', '/assets/country/sao_tome_and_principe.png', 'ST', 99, '239', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (342, 2, '毛里塔尼亚', 'Mauritania', '/assets/country/mauritania.png', 'MR', 99, '222', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (343, 2, '塞内加尔', 'Senegal', '/assets/country/senegal.png', 'SN', 99, '221', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (344, 2, '冈比亚', 'Gambia', '/assets/country/gambia.png', 'GM', 99, '220', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (345, 2, '马里', 'Mali', '/assets/country/mali.png', 'ML', 99, '223', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (346, 2, '布基纳法索', 'Burkina Faso', '/assets/country/burkina_faso.png', 'BF', 99, '226', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (347, 2, '几内亚', 'Guinea', '/assets/country/guinea.png', 'GN', 99, '224', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (348, 2, '几内亚比绍', 'Guinea Bissau', '/assets/country/guinea_bissau.png', 'GW', 99, '245', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (349, 2, '佛得角', 'Cape Verde', '/assets/country/cape_verde.png', 'CV', 99, '238', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (350, 2, '塞拉利昂', 'Sierra Leone', '/assets/country/sierra_leone.png', 'SL', 99, '232', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (351, 2, '利比里亚', 'Liberia', '/assets/country/liberia.png', 'LR', 99, '231', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (352, 2, '科特迪瓦', 'Ivory Coast', '/assets/country/ivory_coast.png', 'CI', 99, '225', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (353, 2, '加纳', 'Ghana', '/assets/country/ghana.png', 'GH', 99, '233', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (354, 2, '多哥', 'Togo', '/assets/country/togo.png', 'TG', 99, '228', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (355, 2, '贝宁', 'Benin', '/assets/country/benin.png', 'BJ', 99, '229', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (356, 2, '尼日尔', 'Niger', '/assets/country/niger.png', 'NE', 99, '227', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (357, 2, '尼日利亚', 'Nigeria', '/assets/country/nigeria.png', 'NG', 99, '234', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (358, 2, '西撒哈拉', 'WesternSahara', '/assets/country/western_sahara.png', 'EH', 99, '210', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (359, 2, '赞比亚', 'Zambia', '/assets/country/zambia.png', 'ZM', 99, '260', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (360, 2, '安哥拉', 'Angola', '/assets/country/angola.png', 'AO', 99, '244', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (361, 2, '津巴布韦', 'Zimbabwe', '/assets/country/zimbabwe.png', 'ZW', 99, '263', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (362, 2, '马拉维', 'Malawi', '/assets/country/malawi.png', 'MW', 99, '265', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (363, 2, '莫桑比克', 'Mozambique', '/assets/country/mozambique.png', 'MZ', 99, '258', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (364, 2, '博茨瓦纳', 'Botswana', '/assets/country/botswana.png', 'BW', 99, '267', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (365, 2, '纳米比亚', 'Namibia', '/assets/country/namibia.png', 'NA', 99, '264', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (366, 2, '南非', 'SouthAfrica', '/assets/country/south_africa.png', 'ZA', 99, '27', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (367, 2, '斯威士兰', 'Eswatini', '/assets/country/swaziland.png', 'SZ', 99, '268', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (368, 2, '莱索托', 'Lesotho', '/assets/country/Lesotho.png', 'LS', 99, '266', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (369, 2, '马达加斯加', 'Madagascar', '/assets/country/madagascar.png', 'MG', 99, '261', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (370, 2, '科摩罗', 'Comoros', '/assets/country/comoros.png', 'KM', 99, '269', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (371, 2, '毛里求斯', 'Mauritius', '/assets/country/Mauritius.png', 'MU', 99, '230', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (372, 2, '圣赫勒拿岛', 'Saint Helena', '/assets/country/st_helena.png', 'SH', 99, '290', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (373, 2, '马约特', 'Mayotte', '/assets/country/mayotte.png', 'YT', 99, '262', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (374, 2, '加拿大', 'Canada', '/assets/country/canada.png', 'CA', 99, '1', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (375, 2, '美国', 'United States', '/assets/country/usa.png', 'US', 99, '1', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (376, 2, '墨西哥', 'Mexico', '/assets/country/mexico.png', 'MX', 99, '52', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (377, 2, '格陵兰', 'Greenland', '/assets/country/greenland.png', 'GL', 99, '299', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (378, 2, '圣皮埃尔和密克隆', 'Saint Pierre and Miquelon', '/assets/country/saint_pierre_and_miquelon.png', 'PM', 99, '508', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (379, 2, '百慕大', 'Bermuda', '/assets/country/bermuda.png', 'BM', 99, '441', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (380, 2, '危地马拉', 'Guatemala', '/assets/country/guatemala.png', 'GT', 99, '502', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (381, 2, '伯利兹', 'Belize', '/assets/country/belize.png', 'BZ', 99, '501', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (382, 2, '萨尔瓦多', 'Salvador', '/assets/country/salvador.png', 'SV', 99, '503', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (383, 2, '洪都拉斯', 'Honduras', '/assets/country/honduras.png', 'HN', 99, '504', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (384, 2, '尼加拉瓜', 'Nicaragua', '/assets/country/nicaragua.png', 'NI', 99, '505', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (385, 2, '哥斯达黎加', 'CostaRica', '/assets/country/costa_rica.png', 'CR', 99, '506', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (386, 2, '巴拿马', 'Panama', '/assets/country/panama.png', 'PA', 99, '507', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (387, 2, '巴哈马', 'Bahamas', '/assets/country/bahamas.png', 'BS', 99, '1242', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (388, 2, '古巴', 'Cuba', '/assets/country/cuba.png', 'CU', 99, '53', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (389, 2, '牙买加', 'Jamaica', '/assets/country/jamaica.png', 'JM', 99, '1876', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (390, 2, '海地', 'Haiti', '/assets/country/haiti.png', 'HT', 99, '509', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (391, 2, '多米尼加共和国', 'Dominican Republic', '/assets/country/dominica_republic.png', 'DO', 99, '1809', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (392, 2, '安提瓜和巴布达', 'Antigua And Barbuda', '/assets/country/antigua_and_barbuda.png', 'AG', 99, '1268', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (393, 2, '圣基茨和尼维斯', 'Saint Kitts And Nevis', '/assets/country/saint_kitts_and_nevis.png', 'KN', 99, '1869', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (394, 2, '多米尼克', 'Dominica', '/assets/country/dominica.png', 'DM', 99, '1767', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (395, 2, '圣卢西亚', 'SaintLucia', '/assets/country/saint_lucia.png', 'LC', 99, '1758', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (396, 2, '圣文森特和格林纳丁斯', 'Saint Vincent And The Grenadines', '/assets/country/saint_vincent_and_the_grenadines.png', 'VC', 99, '784', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (397, 2, '格林纳达', 'Grenada', '/assets/country/grenada.png', 'GD', 99, '1473', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (398, 2, '巴巴多斯', 'Barbados', '/assets/country/barbados.png', 'BB', 99, '1246', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (399, 2, '特立尼达和多巴哥', 'Trinidad And Tobago', '/assets/country/trinidad_and_tobago.png', 'TT', 99, '1868', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (400, 2, '波多黎各', 'Puerto Rico', '/assets/country/puerto_rico.png', 'PR', 99, '1787', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (401, 2, '英属维尔京群岛', 'The British Virgin Islands', '/assets/country/the_british_virgin_islands.png', 'VG', 99, '1284', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (402, 2, '美属维尔京群岛', 'United States Virgin Islands', '/assets/country/united_states_virgin_islands.png', 'VI', 99, '1340', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (403, 2, '安圭拉', 'Anguilla', '/assets/country/anguilla.png', 'AI', 99, '264', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (404, 2, '蒙特塞拉特', 'Montserrat', '/assets/country/montserrat.png', 'MS', 99, '1664', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (405, 2, '瓜德罗普', 'Guadeloupe', '/assets/country/guadeloupe.png', 'GP', 99, '590', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (406, 2, '马提尼克', 'Martinique', '/assets/country/martinique.png', 'MQ', 99, '596', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (407, 2, '阿鲁巴', 'Aruba', '/assets/country/aruba.png', 'AW', 99, '297', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (408, 2, '法属圣马丁', 'Saint Martin', '/assets/country/st_martin.png', 'MF', 99, '1721', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (409, 2, '特克斯和凯科斯群岛', 'Turks And Caicos Islands', '/assets/country/turks_and_caicos_islands.png', 'TC', 99, '1649', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (410, 2, '开曼群岛', 'Cayman Islands', '/assets/country/cayman_islands.png', 'KY', 99, '1345', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (411, 2, '库拉索', 'Curacao', '/assets/country/curacao.png', 'CW', 99, '5999', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (412, 2, '哥伦比亚', 'Colombia', '/assets/country/colombia.png', 'CO', 99, '57', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (413, 2, '委内瑞拉', 'Venezuela', '/assets/country/venezuela.png', 'VE', 99, '58', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (414, 2, '圭亚那', 'Guyana', '/assets/country/guyana.png', 'GY', 99, '592', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (415, 2, '苏里南', 'Suriname', '/assets/country/suriname.png', 'SR', 99, '597', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (416, 2, '法属圭亚那', 'French Guiana', '/assets/country/french_guiana.png', 'GF', 99, '594', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (417, 2, '厄瓜多尔', 'Ecuador', '/assets/country/ecuador.png', 'EC', 99, '593', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (418, 2, '秘鲁', 'Peru', '/assets/country/peru.png', 'PE', 99, '51', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (419, 2, '玻利维亚', 'Bolivia', '/assets/country/bolivia.png', 'BO', 99, '591', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (420, 2, '巴西', 'Brazil', '/assets/country/brazil.png', 'BR', 99, '55', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (421, 2, '智利', 'Chile', '/assets/country/chile.png', 'CL', 99, '56', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (422, 2, '阿根廷', 'Argentina', '/assets/country/argentina.png', 'AR', 99, '54', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (423, 2, '乌拉圭', 'Uruguay', '/assets/country/uruguay.png', 'UY', 99, '598', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (424, 2, '巴拉圭', 'Paraguay', '/assets/country/paraguay.png', 'PY', 99, '595', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (425, 2, '澳大利亚', 'Australia', '/assets/country/australia.png', 'AU', 99, '61', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (426, 2, '新西兰', 'New Zealand', '/assets/country/new_zealand.png', 'NZ', 99, '64', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (427, 2, '帕劳', 'Palau', '/assets/country/palau.png', 'PW', 99, '680', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (428, 2, '密克罗尼西亚联邦', 'Federated States Of Micronesia', '/assets/country/federated_states_of_micronesia.png', 'FM', 99, '691', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (429, 2, '马绍尔群岛', 'Marshall Islands', '/assets/country/marshall_islands.png', 'MH', 99, '692', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (430, 2, '基里巴斯', 'Kiribati', '/assets/country/kiribati.png', 'KI', 99, '686', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (431, 2, '瑙鲁', 'Nauru', '/assets/country/nauru.png', 'NR', 99, '674', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (432, 2, '北马里亚纳', 'Northern Mariana', '/assets/country/northern_mariana.png', 'MP', 99, '670', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (433, 2, '关岛', 'Guam', '/assets/country/guam.png', 'GU', 99, '1671', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (434, 2, '巴布亚新几内亚', 'Papua New Guinea', '/assets/country/papua_new_guinea.png', 'PG', 99, '675', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (435, 2, '所罗门群岛', 'SolomonIslands', '/assets/country/solomon_islands.png', 'SB', 99, '677', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (436, 2, '瓦努阿图', 'Vanuatu', '/assets/country/vanuatu.png', 'VU', 99, '678', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (437, 2, '斐济', 'Fiji', '/assets/country/fiji.png', 'FJ', 99, '679', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (438, 2, '图瓦卢', 'Tuvalu', '/assets/country/tuvalu.png', 'TV', 99, '688', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (439, 2, '萨摩亚', 'Samoa', '/assets/country/samoa.png', 'WS', 99, '685', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (440, 2, '汤加', 'Tonga', '/assets/country/tonga.png', 'TO', 99, '676', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (441, 2, '库克群岛', 'Cook Islands', '/assets/country/island.png', 'CK', 99, '682', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (442, 2, '纽埃', 'Niue', '/assets/country/niue.png', 'NU', 99, '683', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (443, 2, '托克劳', 'Tokelau', '/assets/country/tokelau.png', 'TK', 99, '690', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (444, 2, '瓦利斯和富图纳', 'Wallis And Futuna', '/assets/country/wallis_and_futuna.png', 'WF', 99, '681', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (445, 2, '皮特凯恩群岛', 'Pitcairn Islands', '/assets/country/pitcairn_islands.png', 'PN', 99, '649', 10, '', 1709021277, 1709021277);
INSERT INTO `country` (`id`, `admin_id`, `name`, `alias`, `icon`, `iso1`, `sort`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (446, 2, '美属萨摩亚', 'American Samoa', '/assets/country/american_samoa.png', 'AS', 99, '684', 10, '', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for invite
-- ----------------------------
DROP TABLE IF EXISTS `invite`;
CREATE TABLE `invite` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `code` varchar(50) NOT NULL COMMENT '邀请码',
  `status` tinyint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_invite_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='邀请表';

-- ----------------------------
-- Records of invite
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for lang
-- ----------------------------
DROP TABLE IF EXISTS `lang`;
CREATE TABLE `lang` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `alias` varchar(60) NOT NULL COMMENT '别名',
  `icon` varchar(60) NOT NULL COMMENT '图标',
  `sort` tinyint NOT NULL COMMENT '排序',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统语言表';

-- ----------------------------
-- Records of lang
-- ----------------------------
BEGIN;
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 1, '简体中文', 'zh-CN', '/assets/country/china.png', 1, 10, '简体中文', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 1, '繁體中文', 'zh-TW', '/assets/country/taiwan.png', 99, -1, '繁体中文', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (3, 1, 'English', 'en-US', '/assets/country/usa.png', 99, 10, '英格兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (4, 1, 'عربي', 'ar-AE', '/assets/country/united_arab_emirates.png', 99, -1, '阿拉伯语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (5, 1, 'беларускі', 'be-BY', '/assets/country/belarus.png', 99, -1, '白俄罗斯语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (6, 1, 'български', 'bg-BG', '/assets/country/bulgaria.png', 99, -1, '保加利亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (7, 1, 'čeština', 'cs-CZ', '/assets/country/czech.png', 99, -1, '捷克语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (8, 1, 'dansk', 'da-DK', '/assets/country/denmark.png', 99, -1, '丹麦语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (9, 1, 'Deutsch', 'de-DE', '/assets/country/germany.png', 99, -1, '德语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (10, 1, 'Ελληνικά', 'el-GR', '/assets/country/greece.png', 99, -1, '希腊语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (11, 1, 'español', 'es-ES', '/assets/country/spain.png', 99, -1, '西班牙语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (12, 1, 'eesti keel', 'et-EE', '/assets/country/estonia.png', 99, -1, '爱沙尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (13, 1, 'Suomalainen', 'fi-FI', '/assets/country/finland.png', 99, -1, '芬兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (14, 1, 'Français', 'fr-FR', '/assets/country/france.png', 99, -1, '法语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (15, 1, 'Hrvatski', 'hr-HR', '/assets/country/croatia.png', 99, -1, '克罗地亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (16, 1, 'Magyar', 'hu-HU', '/assets/country/hungary.png', 99, -1, '匈牙利语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (17, 1, 'íslenskur', 'is-IS', '/assets/country/iceland.png', 99, -1, '冰岛语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (18, 1, 'italiano', 'it-IT', '/assets/country/italy.png', 99, -1, '意大利语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (19, 1, '日本', 'ja-JP', '/assets/country/japan.png', 99, -1, '日语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (20, 1, 'Melayu', 'ms-MY', '/assets/country/malaysia.png', 99, -1, '马来语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (21, 1, 'Tiếng Việt', 'vi-VN', '/assets/country/vietnam.png', 99, -1, '越南语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (22, 1, '한국인', 'ko-KR', '/assets/country/north_korea.png', 99, -1, '朝鲜语(韩语)', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (23, 1, 'lietuvių', 'lt-LT', '/assets/country/lithuania.png', 99, -1, '立陶宛语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (24, 1, 'македонски', 'mk-MK', '/assets/country/macedonia.png', 99, -1, '马其顿语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (25, 1, 'Nederlands', 'nl-NL', '/assets/country/netherlands.png', 99, -1, '荷兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (26, 1, 'norsk', 'no-NO', '/assets/country/norway.png', 99, -1, '挪威语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (27, 1, 'Polski', 'pl-PL', '/assets/country/poland.png', 99, -1, '波兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (28, 1, 'Português', 'pt-PT', '/assets/country/portugal.png', 99, -1, '葡萄牙语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (29, 1, 'Română', 'ro-RO', '/assets/country/romania.png', 99, -1, '罗马尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (30, 1, 'Русский', 'ru-RU', '/assets/country/russia.png', 99, -1, '俄语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (31, 1, 'Hrvatski', 'sh-YU', '/assets/country/croatia.png', 99, -1, '克罗地亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (32, 1, 'slovenský', 'sk-SK', '/assets/country/slovakia.png', 99, -1, '斯洛伐克语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (33, 1, 'Slovenščina', 'sl-SI', '/assets/country/slovenia.png', 99, -1, '斯洛文尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (34, 1, 'shqiptare', 'sq-AL', '/assets/country/albania.png', 99, -1, '阿尔巴尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (35, 1, 'svenska', 'sv-SE', '/assets/country/sweden.png', 99, -1, '瑞典语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (36, 1, 'แบบไทย', 'th-TH', '/assets/country/thailand.png', 99, -1, '泰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (37, 1, 'Türkçe', 'tr-TR', '/assets/country/turkey.png', 99, -1, '土耳其语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (38, 1, 'українська', 'uk-UA', '/assets/country/ukraine.png', 99, -1, '乌克兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (39, 1, 'Српски', 'sr-YU', '/assets/country/serbia.png', 99, -1, '塞尔维亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (40, 1, 'עִברִית', 'iw-IL', '/assets/country/israel.png', 99, -1, '希伯来语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (41, 1, 'हिंदी', 'hi-IN', '/assets/country/india.png', 99, -1, '印地语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (42, 1, 'Indonesia', 'id-ID', '/assets/country/indonesia.png', 99, -1, '印尼语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (43, 2, '简体中文', 'zh-CN', '/assets/country/china.png', 1, 10, '简体中文', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (44, 2, '繁體中文', 'zh-TW', '/assets/country/taiwan.png', 99, -1, '繁体中文', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (45, 2, 'English', 'en-US', '/assets/country/usa.png', 99, 10, '英格兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (46, 2, 'عربي', 'ar-AE', '/assets/country/united_arab_emirates.png', 99, -1, '阿拉伯语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (47, 2, 'беларускі', 'be-BY', '/assets/country/belarus.png', 99, -1, '白俄罗斯语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (48, 2, 'български', 'bg-BG', '/assets/country/bulgaria.png', 99, -1, '保加利亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (49, 2, 'čeština', 'cs-CZ', '/assets/country/czech.png', 99, -1, '捷克语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (50, 2, 'dansk', 'da-DK', '/assets/country/denmark.png', 99, -1, '丹麦语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (51, 2, 'Deutsch', 'de-DE', '/assets/country/germany.png', 99, -1, '德语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (52, 2, 'Ελληνικά', 'el-GR', '/assets/country/greece.png', 99, -1, '希腊语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (53, 2, 'español', 'es-ES', '/assets/country/spain.png', 99, -1, '西班牙语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (54, 2, 'eesti keel', 'et-EE', '/assets/country/estonia.png', 99, -1, '爱沙尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (55, 2, 'Suomalainen', 'fi-FI', '/assets/country/finland.png', 99, -1, '芬兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (56, 2, 'Français', 'fr-FR', '/assets/country/france.png', 99, -1, '法语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (57, 2, 'Hrvatski', 'hr-HR', '/assets/country/croatia.png', 99, -1, '克罗地亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (58, 2, 'Magyar', 'hu-HU', '/assets/country/hungary.png', 99, -1, '匈牙利语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (59, 2, 'íslenskur', 'is-IS', '/assets/country/iceland.png', 99, -1, '冰岛语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (60, 2, 'italiano', 'it-IT', '/assets/country/italy.png', 99, -1, '意大利语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (61, 2, '日本', 'ja-JP', '/assets/country/japan.png', 99, -1, '日语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (62, 2, 'Melayu', 'ms-MY', '/assets/country/malaysia.png', 99, -1, '马来语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (63, 2, 'Tiếng Việt', 'vi-VN', '/assets/country/vietnam.png', 99, -1, '越南语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (64, 2, '한국인', 'ko-KR', '/assets/country/north_korea.png', 99, -1, '朝鲜语(韩语)', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (65, 2, 'lietuvių', 'lt-LT', '/assets/country/lithuania.png', 99, -1, '立陶宛语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (66, 2, 'македонски', 'mk-MK', '/assets/country/macedonia.png', 99, -1, '马其顿语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (67, 2, 'Nederlands', 'nl-NL', '/assets/country/netherlands.png', 99, -1, '荷兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (68, 2, 'norsk', 'no-NO', '/assets/country/norway.png', 99, -1, '挪威语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (69, 2, 'Polski', 'pl-PL', '/assets/country/poland.png', 99, -1, '波兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (70, 2, 'Português', 'pt-PT', '/assets/country/portugal.png', 99, -1, '葡萄牙语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (71, 2, 'Română', 'ro-RO', '/assets/country/romania.png', 99, -1, '罗马尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (72, 2, 'Русский', 'ru-RU', '/assets/country/russia.png', 99, -1, '俄语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (73, 2, 'Hrvatski', 'sh-YU', '/assets/country/croatia.png', 99, -1, '克罗地亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (74, 2, 'slovenský', 'sk-SK', '/assets/country/slovakia.png', 99, -1, '斯洛伐克语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (75, 2, 'Slovenščina', 'sl-SI', '/assets/country/slovenia.png', 99, -1, '斯洛文尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (76, 2, 'shqiptare', 'sq-AL', '/assets/country/albania.png', 99, -1, '阿尔巴尼亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (77, 2, 'svenska', 'sv-SE', '/assets/country/sweden.png', 99, -1, '瑞典语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (78, 2, 'แบบไทย', 'th-TH', '/assets/country/thailand.png', 99, -1, '泰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (79, 2, 'Türkçe', 'tr-TR', '/assets/country/turkey.png', 99, -1, '土耳其语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (80, 2, 'українська', 'uk-UA', '/assets/country/ukraine.png', 99, -1, '乌克兰语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (81, 2, 'Српски', 'sr-YU', '/assets/country/serbia.png', 99, -1, '塞尔维亚语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (82, 2, 'עִברִית', 'iw-IL', '/assets/country/israel.png', 99, -1, '希伯来语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (83, 2, 'हिंदी', 'hi-IN', '/assets/country/india.png', 99, -1, '印地语', 1709021277, 1709021277);
INSERT INTO `lang` (`id`, `admin_id`, `name`, `alias`, `icon`, `sort`, `status`, `data`, `updated_at`, `created_at`) VALUES (84, 2, 'Indonesia', 'id-ID', '/assets/country/indonesia.png', 99, -1, '印尼语', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for level
-- ----------------------------
DROP TABLE IF EXISTS `level`;
CREATE TABLE `level` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `icon` varchar(60) NOT NULL COMMENT '图标',
  `level` tinyint NOT NULL COMMENT '等级',
  `money` decimal(12,2) NOT NULL COMMENT '金额',
  `days` tinyint NOT NULL COMMENT '天数',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `desc` text COMMENT '详情',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统等级表';

-- ----------------------------
-- Records of level
-- ----------------------------
BEGIN;
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (1, 1, 'level1Label', '/assets/level/diamond.png', 1, 120.00, -1, 10, '', 'level1Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (2, 1, 'level2Label', '/assets/level/diamond.png', 2, 220.00, -1, 10, '', 'level2Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (3, 1, 'level3Label', '/assets/level/diamond.png', 3, 320.00, -1, 10, '', 'level3Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (4, 1, 'level4Label', '/assets/level/diamond.png', 4, 580.00, -1, 10, '', 'level4Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (5, 1, 'level5Label', '/assets/level/diamond.png', 5, 888.00, -1, 10, '', 'level5Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (6, 2, 'level1Label', '/assets/level/diamond.png', 1, 120.00, -1, 10, '', 'level1Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (7, 2, 'level2Label', '/assets/level/diamond.png', 2, 220.00, -1, 10, '', 'level2Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (8, 2, 'level3Label', '/assets/level/diamond.png', 3, 320.00, -1, 10, '', 'level3Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (9, 2, 'level4Label', '/assets/level/diamond.png', 4, 580.00, -1, 10, '', 'level4Desc', 1709021277, 1709021277);
INSERT INTO `level` (`id`, `admin_id`, `name`, `icon`, `level`, `money`, `days`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (10, 2, 'level5Label', '/assets/level/diamond.png', 5, 888.00, -1, 10, '', 'level5Desc', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for level_order
-- ----------------------------
DROP TABLE IF EXISTS `level_order`;
CREATE TABLE `level_order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `level_id` int unsigned NOT NULL COMMENT '等级ID',
  `status` tinyint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `expired_at` int unsigned NOT NULL COMMENT '过期时间',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户等级订单表';

-- ----------------------------
-- Records of level_order
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for notify
-- ----------------------------
DROP TABLE IF EXISTS `notify`;
CREATE TABLE `notify` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `type` smallint NOT NULL DEFAULT '11' COMMENT '类型 11前台通知',
  `name` varchar(60) NOT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统通知表';

-- ----------------------------
-- Records of notify
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for real_name_auth
-- ----------------------------
DROP TABLE IF EXISTS `real_name_auth`;
CREATE TABLE `real_name_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `real_name` varchar(50) NOT NULL COMMENT '真实姓名',
  `number` varchar(50) NOT NULL COMMENT '卡号',
  `photo1` varchar(120) NOT NULL COMMENT '证件照1',
  `photo2` varchar(120) NOT NULL COMMENT '证件照2',
  `photo3` varchar(120) NOT NULL COMMENT '证件照3',
  `telephone` varchar(50) NOT NULL COMMENT '手机号码',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '详细地址',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型 1身份证',
  `status` tinyint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1拒绝 10审核 20完成',
  `data` text COMMENT 'input配置',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户实名认证表';

-- ----------------------------
-- Records of real_name_auth
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '管理类型一样',
  `field` varchar(50) NOT NULL COMMENT '建铭',
  `value` text COMMENT '键值',
  `data` text COMMENT 'input配置',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_setting_field` (`field`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户设置表';

-- ----------------------------
-- Records of setting
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for translate
-- ----------------------------
DROP TABLE IF EXISTS `translate`;
CREATE TABLE `translate` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `lang` varchar(60) NOT NULL COMMENT '语言标识',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '类型 1系统翻译 2前台翻译',
  `field` varchar(60) NOT NULL COMMENT '建铭',
  `value` text COMMENT '键值',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=511 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统语言字典表';

-- ----------------------------
-- Records of translate
-- ----------------------------
BEGIN;
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (1, 1, 'zh-CN', '验证参数不能为空', 1, 'required', '参数不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (2, 1, 'zh-CN', '验证参数min', 1, 'min', '参数最小值为', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (3, 1, 'zh-CN', '验证参数max', 1, 'max', '参数最大值为', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (4, 1, 'zh-CN', '验证参数oneof', 1, 'oneof', '范围不匹配', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (5, 1, 'zh-CN', '验证参数大于', 1, 'gt', '必须大于', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (6, 1, 'zh-CN', '验证参数大于等于', 1, 'gte', '必须大于等于', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (7, 1, 'zh-CN', '未设置安全密钥', 1, 'notSecurityKey', '未设置安全密钥, 请前往设置...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (8, 1, 'zh-CN', '参数不正确', 1, 'incorrectFormat', '格式不正确, 检查当前格式~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (9, 1, 'zh-CN', '验证码不正确', 1, 'incorrectCode', '验证码不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (10, 1, 'zh-CN', '旧密码不正确', 1, 'TheOldPasswordIsIncorrect', '旧密码不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (11, 1, 'zh-CN', '旧安全密钥不正确', 1, 'TheOldSecurityKeyIsIncorrect', '旧安全密钥不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (12, 1, 'zh-CN', '账号或密码不正确', 1, 'accountOrPasswordIncorrect', '账号或密码不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (13, 1, 'zh-CN', '安全密钥不能为空', 1, 'securityBeEmpty', '安全密钥不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (14, 1, 'zh-CN', '安全密钥不正确', 1, 'securityKeyIncorrect', '安全密钥不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (15, 1, 'zh-CN', '手机号码不能为空', 1, 'telephoneBeEmpty', '手机号码不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (16, 1, 'zh-CN', '邮箱不能为空', 1, 'emailBeEmpty', '邮箱不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (17, 1, 'zh-CN', '邀请码不存在', 1, 'inviteBeEmpty', '邀请码不存在', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (18, 1, 'zh-CN', '暂时不能注册', 1, 'notRegister', '暂时不能注册', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (19, 1, 'zh-CN', '账单类型不正确', 1, 'wrongBillType', '账单类型不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (20, 1, 'zh-CN', '用户名不存在', 1, 'usernameNotFound', '用户名不存在', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (21, 1, 'zh-CN', '支付方式不存在', 1, 'paymentNotFound', '支付方式不存在', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (22, 1, 'zh-CN', '余额不足', 1, 'insufficientBalance', '余额不足', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (23, 1, 'zh-CN', '未绑定卡片', 1, 'accountBeEmpty', '未绑定卡片', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (24, 1, 'zh-CN', '超出绑定限制', 1, 'cardAdditionLimitExceeded', '超出绑定限制', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (25, 1, 'zh-CN', '超出提现次数', 1, 'exceededNumberOfWithdrawals', '提款次数超出', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (26, 1, 'zh-CN', '金额范围不匹配', 1, 'amountBetweenMismatch', '金额范围不匹配', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (27, 1, 'zh-CN', '账单类型充值', 1, 'walletBillTypeDeposit', '余额充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (28, 1, 'zh-CN', '账单类型提现', 1, 'walletBillTypeWithdraw', '余额提现', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (29, 1, 'zh-CN', '账单类型资产充值', 1, 'walletBillTypeAssetsDeposit', '资产充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (30, 1, 'zh-CN', '账单类型资产提现', 1, 'walletBillTypeAssetsWithdraw', '资产提现', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (31, 1, 'zh-CN', '账单类型购买产品', 1, 'walletBillTypeBuyProduct', '购买产品', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (32, 1, 'zh-CN', '账单类型购买等级', 1, 'walletBillTypeBuyLevel', '购买等级', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (33, 1, 'zh-CN', '账单类型收益', 1, 'walletBillTypeEarnings', '收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (34, 1, 'zh-CN', '账单类型奖励', 1, 'walletBillTypeAward', '奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (35, 1, 'zh-CN', '账单类型资产购买产品', 1, 'walletBillTypeAssetsBuyProduct', '资产购买产品', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (36, 1, 'zh-CN', '账单类型资产收益', 1, 'walletBillTypeAssetsEarnings', '资产收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (37, 1, 'zh-CN', '账单类型资产奖励', 1, 'walletBillTypeAssetsAward', '资产奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (38, 1, 'zh-CN', '账单类型注册奖励', 1, 'walletBillTypeRegisterAward', '注册奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (39, 1, 'zh-CN', '账单类型邀请奖励', 1, 'walletBillTypeShareAward', '邀请奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (40, 1, 'zh-CN', '账单类型系统充值', 1, 'walletBillTypeSystemDeposit', '余额系统充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (41, 1, 'zh-CN', '账单类型资产系统充值', 1, 'walletBillTypeSystemAssetsDeposit', '资产系统充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (42, 1, 'zh-CN', '账单类型系统扣款', 1, 'walletBillTypeSystemWithdraw', '余额系统扣款', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (43, 1, 'zh-CN', '账单类型资产系统扣款', 1, 'walletBillTypeSystemAssetsWithdraw', '资产系统扣款', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (44, 1, 'zh-CN', '账单类型团队收益', 1, 'walletBillTypeTeamEarnings', '团队收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (45, 1, 'zh-CN', '账单类型团队资产收益', 1, 'walletBillTypeTeamAssetsEarnings', '团队资产收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (46, 1, 'zh-CN', '账单类型资金提现拒绝', 1, 'walletBillTypeWithdrawRefuse', '资金提现拒绝', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (47, 1, 'zh-CN', '账单类型资产提现拒绝', 1, 'walletBillTypeAssetsWithdrawRefuse', '资产提现拒绝', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (48, 1, 'zh-CN', 'level1标题', 1, 'level1Label', 'VIP1', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (49, 1, 'zh-CN', 'level2标题', 1, 'level2Label', 'VIP2', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (50, 1, 'zh-CN', 'level3标题', 1, 'level3Label', 'VIP3', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (51, 1, 'zh-CN', 'level4标题', 1, 'level4Label', 'VIP4', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (52, 1, 'zh-CN', 'level5标题', 1, 'level5Label', 'VIP5', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (53, 1, 'zh-CN', 'level1详情', 1, 'level1Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (54, 1, 'zh-CN', 'level2详情', 1, 'level2Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (55, 1, 'zh-CN', 'level3详情', 1, 'level3Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (56, 1, 'zh-CN', 'level4详情', 1, 'level4Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (57, 1, 'zh-CN', 'level5详情', 1, 'level5Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (58, 1, 'zh-CN', '储蓄卡', 1, 'DebitCard', '储蓄卡', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (59, 1, 'zh-CN', '美国银行', 1, 'BankOfAmerica', '美国银行', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (60, 1, 'zh-CN', '储蓄卡(资金提现)', 1, 'bankFundsWithdraw', '储蓄卡[资金提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (61, 1, 'zh-CN', '储蓄卡(资产提现)', 1, 'bankAssetsWithdraw', '储蓄卡[资产提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (62, 1, 'zh-CN', '数字货币(资金提现)', 1, 'digitalFundsWithdraw', '数字货币[资金提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (63, 1, 'zh-CN', '数字货币(资产提现)', 1, 'digitalAssetsWithdraw', '数字货币[资产提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (64, 1, 'zh-CN', '储蓄卡(资金充值)', 1, 'bankFundsDeposit', '储蓄卡[资金充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (65, 1, 'zh-CN', '储蓄卡(资产充值)', 1, 'bankAssetsDeposit', '储蓄卡[资产充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (66, 1, 'zh-CN', '数字货币(资金充值)', 1, 'digitalFundsDeposit', '数字货币[资金充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (67, 1, 'zh-CN', '数字货币(资产充值)', 1, 'digitalAssetsDeposit', '数字货币[资产充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (68, 1, 'zh-CN', '充值提示详情', 1, 'depositDesc', '充值注意事项', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (69, 1, 'zh-CN', '提现提示详情', 1, 'withdrawDesc', '提现注意事项', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (70, 1, 'zh-CN', '站点介绍', 1, 'siteIntroduce', '站点介绍', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (71, 1, 'zh-CN', '站点通知', 1, 'siteAnnouncement', '站点通知', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (72, 1, 'zh-CN', '关于我们', 1, 'aboutUs', '关于我们', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (73, 1, 'zh-CN', '平台服务', 1, 'siteService', '平台服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (74, 1, 'zh-CN', '联系我们', 1, 'contactUs', '联系我们', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (75, 1, 'zh-CN', '社区', 1, 'social', '社区', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (76, 1, 'zh-CN', '关于我们标题', 1, 'aboutUsLabel', '关于我们', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (77, 1, 'zh-CN', '关于我们详情', 1, 'aboutUsDesc', '<pre><span style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">关于我们</font></font></span><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">欢迎来到创新与卓越相遇的地方！ 我们是一家先锋科技企业，致力于突破技术界限并为客户提供卓越的解决方案。 我们的承诺不仅仅是创新； 这是关于了解客户需求并承担社会责任。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的使命和核心价值观</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">在[公司名称]，我们的使命是通过创新为未来铺平道路。 我们的核心价值观决定了我们的成功：</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">技术创新：不断探索和采用前沿技术，提供最先进的解决方案和服务。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">以客户为中心的方法：通过提供满足客户独特需求的高质量定制服务来优先考虑客户的成功。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">协作团队合作：我们的团队由顶级专业人士组成，我们相信协作是成功的关键，他们不断协作以实现共同成功。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">社会责任：作为企业责任的一部分，我们积极支持可持续发展努力，为社会进步做出贡献。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的服务范围</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们提供跨多个领域的全面服务：</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">软件开发与 工程：利用最新的软件开发技术，我们提供高效、可靠和定制的软件解决方案。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">人工智能与 数据分析：专注于人工智能和数据分析，利用洞察和智能技术提供定制化解决方案。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">物联网 (IoT) 和 未来技术：我们致力于推进物联网和未来技术，创造智能解决方案和创新产品。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">数字化转型服务：协助企业数字化转型，利用技术提高效率、降低成本、创造更多商业价值。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">客户关系</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">与客户建立持久的合作伙伴关系是我们的首要任务。 我们用心倾听他们的需求，提供个性化的解决方案，并提供持续的支持以推动他们的持续增长。 积极寻求并纳入客户反馈可确保我们提供最优质的服务和技术支持。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">社会责任承诺</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们不仅仅是一家科技公司； 我们是社会责任的管家。 参与社区活动、支持环境保护和支持社会事业是我们致力于社会进步和可持续发展的重要组成部分。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">与我们合作</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">无论您的业务阶段或需求如何，我们都致力于成为您值得信赖的合作伙伴，提供最佳的技术解决方案。 如果您想进一步探索我们的服务或与我们合作，请随时与我们联系。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的创新方法</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">创新不仅仅是一个流行词，而是一个新词。 这是支撑我们所做一切的理念。 这是我们运营的基石。 对创新的不懈追求促使我们不断探索新技术和新方法，确保我们保持领先优势。 我们大力投资研发，培育鼓励创造力、实验和创新思维的文化。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们成功背后的团队</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的成功归功于一支多元化且才华横溢的团队。 他们由行业专家、远见卓识者和具有前瞻性思维的个人组成，他们在挑战中不断成长，并无缝协作，提供突破性的解决方案。 我们重视多样性、包容性和持续学习，创造一个促进增长和创新的环境。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">行业认可和成就</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们对卓越和创新的承诺赢得了业界的认可。 我们对我们的成就感到自豪，不仅将它们视为里程碑，而且将其视为进一步前进的灵感。 我们的业绩记录意义重大，反映出我们致力于提供超出预期的成果。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">道德实践和透明度</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">透明度和道德实践构成了我们运营的基石。 我们与客户、合作伙伴和利益相关者保持开放的沟通，确保每次互动中的信任和可靠性。 我们对诚信的承诺指导着我们所有的业务往来。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">可持续发展和环境意识</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">作为一个负责任的企业实体，我们在运营的各个方面都优先考虑可持续发展。 我们将环保实践融入我们的流程中，努力最大限度地减少我们的环境足迹。 此外，我们积极支持促进环境保护和生态意识的举措。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">拥抱变革，共创美好未来</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们将变革视为进步的催化剂。 我们迅速适应不断发展的技术格局，预见趋势，并为我们自己和我们的客户做好迎接未来挑战和机遇的准备。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"text-align: left; font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">加入我们的创新之旅</font></font></div></pre>', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (78, 1, 'zh-CN', '服务协议标题', 1, 'serviceAgreementLabel', '服务协议', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (79, 1, 'zh-CN', '服务协议详情', 1, 'serviceAgreementDesc', '服务描述\n平台同意按照本协议的条款和条件向用户提供以下服务：\n*   账户管理 - 用户可以在平台上创建和管理个人交易账户，包括存储、兑换和交易数字货币。\n*   交易服务 - 平台提供数字货币交易服务，包括买卖订单管理、市场信息提供和交易执行。\n*   安全与合规 - 平台致力于提供安全的交易环境，并遵守相关的法律法规。\n*   客户支持 - 用户可以通过客户支持获取交易相关的帮助和咨询。\n用户责任\n*   账户安全 - 用户负责保护其账户信息的安全，包括用户名和密码。\n*   合规操作 - 用户同意遵守所有适用的法律、法规和平台的规则。\n*   风险认知 - 用户理解并同意，数字货币交易存在市场风险，包括但不限于价格波动、技术问题和法律风险。\n风险披露\n用户明白，数字货币交易涉及高风险，可能导致部分或全部资金的损失。用户应根据自己的财务状况和风险承受能力进行交易。\n知识产权\n平台的所有内容和技术，包括但不限于软件、界面设计、商标和版权均属于平台或其授权人所有。\n免责声明\n平台不对由于系统故障、网络问题或其他不可抗力因素导致的服务中断或交易问题承担责任。\n修改和终止\n平台保留随时修改本协议或服务内容的权利，并会提前通知用户。用户继续使用服务即表示接受修改后的协议。\n适用法律和争端解决\n本协议的解释和执行适用 。任何由本协议引起的争端应通过友好协商解决，如协商不成，可提交法院诉讼解决。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (80, 1, 'zh-CN', '隐私协议标题', 1, 'privacyAgreementLabel', '隐私协议', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (81, 1, 'zh-CN', '隐私协议详情', 1, 'privacyAgreementDesc', '<div style=\"text-align: left;\"><span style=\"font-family: Roboto, -apple-system, Avenir, BlinkMacSystemFont, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">尊敬的用户欢迎使用 我们的产品和服务。 我们深知您对个人信息安全和隐私的重视，我们将尽力保护您个人信息的安全，并遵循以下隐私政策： 收集信息的目的：我们收集您的个人信息 为了更好地为您提供服务，提升用户体验，保证产品和服务的安全。 我们只会收集必要的信息，绝不会超出声明目的使用您的个人信息。 收集的信息类型：我们可能收集的信息类型包括但不限于：姓名、电子邮件地址、联系信息、IP地址、设备信息、浏览器类型等信息使用和共享：我们只会使用您的个人信息 在提供服务和提高产品质量的范围内。 除非获得您的同意或法律法规要求，我们不会向任何第三方出售、出租或分享您的个人信息。 信息安全：我们采取了一系列安全措施来保护您的个人信息，包括但不限于加密技术、访问控制、安全审计等。我们将尽力确保您的信息安全，不被未经授权的访问、使用 或披露。 用户权利：您有权查询、更正、删除个人信息。 如果您对我们收集的个人信息有任何疑问或需要行使相关权利，请随时联系我们的客服团队。 隐私政策的变更：随着我们的产品和服务的发展，我们的隐私政策可能会不时更新。 我们将在网站上发布任何政策变更，并根据适用法律的要求向您提供通知。 接受条款：使用我们的产品和服务，即表示您同意本隐私政策的条款和条件。 联系我们：如果您对我们的隐私政策有任何疑问或建议，请随时联系我们的隐私团队。 感谢您对我们的信任，我们将继续努力确保您的信息安全和隐私</font></font></span></div>', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (82, 1, 'zh-CN', '帮助中心1标题', 1, 'helpers1Label', '重置密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (83, 1, 'zh-CN', '帮助中心1详情', 1, 'helpers1Desc', '如果您忘记了密码，可以轻松重置。请在登录页面点击“忘记密码？”链接，然后输入您的注册电子邮件地址。我们将发送一封包含重置密码指令的电子邮件给您。遵循邮件中的指示，您就可以设置新密码。为了保证账户安全，请确保您的新密码足够复杂，且不与其他网站共享相同密码。如果在重置密码过程中遇到任何问题，欢迎联系我们的客户支持团队获取帮助。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (84, 1, 'zh-CN', '帮助中心2标题', 1, 'helpers2Label', '安全建议', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (85, 1, 'zh-CN', '帮助中心2详情', 1, 'helpers2Desc', '为了保护您的账户安全，我们建议您采取以下措施：首先，使用强密码，结合大写字母、小写字母、数字和特殊符号。其次，定期更换密码，避免使用相同的密码登录不同网站。此外，启用两因素认证（2FA）增加额外的安全层。请勿在公共计算机或不安全的网络环境中登录账户，并始终保持您的个人信息和联系方式更新。最后，警惕任何可疑的邮件或信息，避免点击不明链接或分享敏感信息。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (86, 1, 'zh-CN', '帮助中心3标题', 1, 'helpers3Label', '取消订阅', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (87, 1, 'zh-CN', '帮助中心3详情', 1, 'helpers3Desc', '取消订阅非常简单。登录您的账户后，前往账户设置 页面，在那里您会找到“订阅管理”选项。选择您想要取消的服务，并点击“取消订阅”按钮。根据您的订阅类型，可能会有一段提前通知期。请注意，取消订阅可能会影响到您享受的某些服务或优惠。如果您在取消订阅过程中遇到任何问题或需要进一步的帮助，欢迎随时联系我们的客户服务团队。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (88, 1, 'zh-CN', '帮助中心4标题', 1, 'helpers4Label', '支付方式', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (89, 1, 'zh-CN', '帮助中心4详情', 1, 'helpers4Desc', '为了方便用户，我们支持多种支付方式。您可以使用信用卡（Visa, MasterCard, American Express等）、借记卡、电子钱包（如PayPal, Apple Pay等）以及银行转账等方式进行支付。所有支付过程都是经过加密的，确保您的交易信息安全。请注意，根据您所在的地区，可用的支付方式可能会有所不同。如果您在支付过程中遇到任何问题，或者需要了解更多关于支付选项的信息，请联系我们的客户支持团队。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (90, 1, 'zh-CN', '帮助中心5标题', 1, 'helpers5Label', '退货政策', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (91, 1, 'zh-CN', '帮助中心5详情', 1, 'helpers5Desc', '我们承诺提供满意的退货政策。如果您对购买的产品不满意，可以在收到货物后30天内申请退货。要开始退货流程，请登录您的账户并前往“订单历史”页面，选择相应的订单并点击“申请退货”。请确保退回的商品处于未使用状态，并包含所有原包装和附件。退货处理完成后，我们将退还您的支付金额，退款通常在5至10个工作日内处理完毕。如果您在退货过程中需要任何帮助，请随时联系我们的客户服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (92, 1, 'zh-CN', '帮助中心6标题', 1, 'helpers6Label', '追踪订单', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (93, 1, 'zh-CN', '帮助中心6详情', 1, 'helpers6Desc', '在您的订单被发货后，您将通过电子邮件收到一个包含追踪链接的发货通知。点击该链接，您可以查看订单的最新运输状态和预计送达时间。此外，您也可以直接在我们的网站上通过“我的订单”部分追踪订单状态。如果您对订单的运输状态有任何疑问或需要帮助，请联系我们的客户服务团队。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (94, 1, 'zh-CN', '帮助中心7标题', 1, 'helpers7Label', '产品期限', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (95, 1, 'zh-CN', '帮助中心7详情', 1, 'helpers7Desc', '我们的产品享有一定期限的保修服务，确保您在使用过程中得到必要的支持。保修期限通常从购买日期起算，具体时长取决于产品类型。保修内容通常包括制造缺陷和材料问题。在保修期内，如果产品出现合格的保修问题，我们将提供免费修理或更换服务。请注意，保修不包括因误用、正常磨损或未经授权的维修造成的损坏。详细的保修条款和条件请参见我们的产品保修政策，或联系我们的客户服务团队获取更多信息。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (96, 1, 'zh-CN', '平台服务1标题', 1, 'service1Label', '产品服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (97, 1, 'zh-CN', '平台服务1详情', 1, 'service1Desc', '欢迎来到我们的交易所平台！我们是一家致力于为您提供安全、高效、多样化的交易体验的领先交易平台。无论您是初学者还是经验丰富的交易者，我们都为您提供了丰富的产品和服务，以满足您的需求。<br />我们的产品特点：<br />多样化的交易对: 我们支持多种数字资产交易对，涵盖主流加密货币以及新兴数字资产，让您可以灵活选择交易组合。<br />安全保障: 我们采取严格的安全措施，包括多重身份验证、冷热钱包存储体系等，保障您的资产安全。<br />高效的交易执行: 我们的平台配备高性能的交易引擎，保证交易执行速度和效率，让您能够快速响应市场变化。<br />用户友好的界面: 我们的界面简洁直观，适合各种水平的交易者使用，同时也提供了专业级别的图表和工具，满足有更高要求的用户。<br />教育资源和客户支持: 我们提供广泛的教育资源，包括文章、视频和在线课程，帮助您了解市场动态和交易策略。此外，我们的客户支持团队全天候为您提供支持。<br /><br />我们的服务承诺：<br />透明度与合规性: 我们遵循严格的法规标准，保障交易过程的透明度和合规性。<br />持续创新与发展: 我们致力于不断提升平台功能和服务，引入新的技术和工具，以满足市场不断变化的需求。<br />社区参与与奖励: 我们鼓励社区参与，通过奖励计划回馈我们的忠实用户。<br />无论您是寻求快速交易、投资多样化数字资产，还是寻找一个安全可靠的交易所平台，我们都将竭诚为您提供最优质的服务和支持。<br />加入我们，开始您的交易之旅吧！', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (98, 1, 'zh-CN', '平台服务2标题', 1, 'service2Label', '基础方案', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (99, 1, 'zh-CN', '平台服务2详情', 1, 'service2Desc', '注册和账户管理<br />浏览市场和实时数据<br />限制性的交易功能<br />专业方案 - $9.99/月：<br />全面的交易功能，包括买卖、存款和提款<br />高级图表和分析工具<br />客户支持优先服务<br />额外的教育资源和市场报告<br />企业方案 - 定制报价：<br />适用于机构和专业交易者<br />定制化的解决方案<br />安全审计和额外的安全措施<br />专属的客户经理和支持团队<br />价格策略说明：<br />弹性定价: 我们提供多种价格选择，满足不同用户群体的需求。用户可以根据其需求和预算选择适合自己的方案。<br />试用期: 对于新用户，我们提供免费试用期或者特定功能的试用，让用户可以在付费前体验我们的服务。<br />优惠和折扣: 定期推出促销活动、折扣和特惠套餐，以回馈现有用户并吸引新用户。<br />定期支付: 客户可以选择按月或按年支付，按年支付的用户可以享受一定的折扣优惠。<br />定价透明度与承诺：<br />我们承诺定价透明，没有隐藏费用或附加费用。所有价格均明确列出，让用户清楚了解他们所支付的服务内容。我们也欢迎用户反馈和建议，以不断优化我们的定价策略，以更好地满足用户需求。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (100, 1, 'zh-CN', '平台服务3标题', 1, 'service3Label', '新产品介绍', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (101, 1, 'zh-CN', '平台服务3详情', 1, 'service3Desc', '在数字资产市场的不断演变和增长中，我们引以为傲地推出了全新的「数字资产投资组合优化工具」，旨在帮助投资者最大化其投资回报并降低风险。<br />主要特点：<br />智能投资组合优化：<br />利用先进的算法和数据分析，我们的工具能够帮助用户构建最优的数字资产投资组合，考虑到风险和收益的平衡。<br />动态资产配置：<br />根据市场实时数据和用户偏好，工具能够动态调整投资组合，以应对市场波动和变化，最大程度地优化投资组合表现。<br />风险管理与模拟测试：<br />提供多种风险管理工具，让用户能够对不同的投资方案进行模拟测试，了解可能的风险和回报情况，做出更明智的投资决策。<br />个性化建议与报告：<br />提供个性化的投资建议和定期报告，帮助用户跟踪投资组合的表现，并根据实时数据做出调整。<br />如何使用我们的产品：<br />注册账户：<br />快速注册账户并登录我们的平台。<br />输入投资偏好：<br />根据您的投资目标、风险承受能力和偏好，输入相关信息。<br />获取优化建议：<br />根据您提供的信息，工具将生成最佳的数字资产投资组合建议。<br />跟踪和调整：<br />定期跟踪投资组合表现，并根据报告的建议调整您的投资策略。<br />我们的承诺：<br />我们的「数字资产投资组合优化工具」旨在为您提供更智能、更高效的投资解决方案。我们致力于提供卓越的用户体验，确保您的投资决策更加明智和可靠。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (102, 1, 'zh-CN', '未添加资产', 2, 'notAssets', '未添加资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (103, 1, 'zh-CN', '饼图', 2, 'pieChart', '饼图', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (104, 1, 'zh-CN', '折线图', 2, 'lineChart', '折线图', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (105, 1, 'zh-CN', '全部', 2, 'all', '全部', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (106, 1, 'zh-CN', '前往', 2, 'goto', '前往', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (107, 1, 'zh-CN', '暂无数据', 2, 'noData', '暂无数据', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (108, 1, 'zh-CN', '注册时间', 2, 'createdTime', '注册时间', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (109, 1, 'zh-CN', '确定操作标题', 2, 'isExecute', '您是否确定执行当前操作?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (110, 1, 'zh-CN', '确定购买', 2, 'isBuy', '您是否确定购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (111, 1, 'zh-CN', '货币单位', 2, 'currency', '$', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (112, 1, 'zh-CN', '充值', 2, 'deposit', '充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (113, 1, 'zh-CN', '提现', 2, 'withdraw', '提现', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (114, 1, 'zh-CN', '首页', 2, 'home', '首页', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (115, 1, 'zh-CN', '行情', 2, 'markets', '市场', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (116, 1, 'zh-CN', '合约', 2, 'contract', '合约', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (117, 1, 'zh-CN', '期货', 2, 'futures', '期货', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (118, 1, 'zh-CN', '我的', 2, 'user', '我的', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (119, 1, 'zh-CN', '数字货币', 2, 'digitalMarkets', '数字货币', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (120, 1, 'zh-CN', '外汇市场', 2, 'forexMarkets', '外汇市场', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (121, 1, 'zh-CN', '期货市场', 2, 'futuresMarkets', '期货市场', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (122, 1, 'zh-CN', '钱包管理', 2, 'walletManage', '钱包管理', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (123, 1, 'zh-CN', '团队管理', 2, 'teamManage', '团队管理', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (124, 1, 'zh-CN', '服务设置', 2, 'serviceManage', '服务设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (125, 1, 'zh-CN', '卡片管理', 2, 'accountManage', '卡片管理', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (126, 1, 'zh-CN', '我的钱包', 2, 'myWallet', '我的钱包', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (127, 1, 'zh-CN', '我的资产', 2, 'myAssets', '我的资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (128, 1, 'zh-CN', '邀请好友', 2, 'inviteFriends', '邀请好友', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (129, 1, 'zh-CN', '我的团队', 2, 'myTeam', '我的团队', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (130, 1, 'zh-CN', '团队收益', 2, 'teamEarnings', '团队收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (131, 1, 'zh-CN', '团队成员', 2, 'teamMember', '团队成员', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (132, 1, 'zh-CN', '会员中心', 2, 'memberVip', '会员中心', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (133, 1, 'zh-CN', '实名认证', 2, 'realAuth', '实名认证', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (134, 1, 'zh-CN', '帮助中心', 2, 'helpers', '帮助中心', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (135, 1, 'zh-CN', '个人设置', 2, 'settings', '个人设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (136, 1, 'zh-CN', '下载中心', 2, 'download', '下载中心', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (137, 1, 'zh-CN', '下载副标题', 2, 'downloadSmall', '使用手机版本能更好的体验服务~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (138, 1, 'zh-CN', '邀请好友副标题', 2, 'inviteFriendsSmall', '邀请您的好友一起Happy, 同时能领取到惊喜收益。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (139, 1, 'zh-CN', '会员中心副标题', 2, 'memberVipSmall', '升级成为尊贵的会员用户, 同时享受各种平台优惠服务~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (140, 1, 'zh-CN', '会员中心标题', 2, 'memberVipTitle', '选择您的会员计划', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (141, 1, 'zh-CN', '登录', 2, 'login', '登录', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (142, 1, 'zh-CN', '忘记密码', 2, 'forgotPassword', '忘记密码?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (143, 1, 'zh-CN', '登录副标题', 2, 'loginSmall', '登录您的账号, 开启旅程～', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (144, 1, 'zh-CN', '未注册去登录', 2, 'toRegister', '还没有账号, 去注册!', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (145, 1, 'zh-CN', '注册', 2, 'register', '注册', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (146, 1, 'zh-CN', '注册副标题', 2, 'registerSmall', '创建您的账户, 开启旅程~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (147, 1, 'zh-CN', '已有账户去登录', 2, 'toLogin', '已有账户, 现在去登录!', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (148, 1, 'zh-CN', '退出登录', 2, 'logout', '退出登录', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (149, 1, 'zh-CN', '退出登录副标题', 2, 'logoutSmall', '您确定要退出登录吗?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (150, 1, 'zh-CN', '有什么能帮上您的吗', 2, 'helperYou', '你好, 有什么能帮上您的吗?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (151, 1, 'zh-CN', '24小时在线服务', 2, '24hoursOnline', '24小时在线服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (152, 1, 'zh-CN', '更多', 2, 'more', '更多', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (153, 1, 'zh-CN', '查询', 2, 'search', '查询', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (154, 1, 'zh-CN', '提交', 2, 'submit', '提交', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (155, 1, 'zh-CN', '重新提交', 2, 'reSubmit', '重新提交', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (156, 1, 'zh-CN', '完成', 2, 'complete', '完成', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (157, 1, 'zh-CN', '审核', 2, 'pending', '审核', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (158, 1, 'zh-CN', '拒绝', 2, 'refuse', '拒绝', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (159, 1, 'zh-CN', '筛选', 2, 'filter', '筛选', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (160, 1, 'zh-CN', '类型', 2, 'type', '类型', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (161, 1, 'zh-CN', '时间范围', 2, 'betweenTime', '时间范围', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (162, 1, 'zh-CN', '余额', 2, 'balance', '余额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (163, 1, 'zh-CN', '充值金额', 2, 'depositAmount', '充值金额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (164, 1, 'zh-CN', '提现金额', 2, 'withdrawAmount', '提现金额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (165, 1, 'zh-CN', '可用金额', 2, 'availableAmount', '可用金额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (166, 1, 'zh-CN', '充值凭证', 2, 'depositProof', '充值凭证', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (167, 1, 'zh-CN', '提现账户', 2, 'withdrawAccount', '提现账户', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (168, 1, 'zh-CN', '未绑定提现账户', 2, 'notBindWithdrawAccount', '未绑定提现账户', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (169, 1, 'zh-CN', '充值帐户', 2, 'depositAccount', '充值帐户', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (170, 1, 'zh-CN', '账户信息', 2, 'depositAccountInfo', '账户信息', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (171, 1, 'zh-CN', '总资产', 2, 'totalAssets', '总资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (172, 1, 'zh-CN', '资产分布', 2, 'assetsBlock', '资产分布', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (173, 1, 'zh-CN', '资产', 2, 'assets', '资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (174, 1, 'zh-CN', '取消', 2, 'cancel', '取消', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (175, 1, 'zh-CN', '确定', 2, 'confirm', '确定', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (176, 1, 'zh-CN', '删除', 2, 'delete', '删除', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (177, 1, 'zh-CN', '删除确认标题', 2, 'deleteLabel', '删除确定', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (178, 1, 'zh-CN', '是否确定删除当前记录', 2, 'deleteRecord', '是否确定删除当前记录?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (179, 1, 'zh-CN', '新增', 2, 'create', '新增', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (180, 1, 'zh-CN', '编辑', 2, 'edit', '编辑', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (181, 1, 'zh-CN', '详情', 2, 'desc', '详情', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (182, 1, 'zh-CN', '查看', 2, 'views', '查看', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (183, 1, 'zh-CN', '复制', 2, 'copy', '复制', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (184, 1, 'zh-CN', '复制成功', 2, 'copySuccess', '复制成功', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (185, 1, 'zh-CN', '购买', 2, 'buy', '购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (186, 1, 'zh-CN', '立即购买', 2, 'nowBuy', '立即购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (187, 1, 'zh-CN', '已购买', 2, 'purchased', '已购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (188, 1, 'zh-CN', '钱包交易记录', 2, 'transactions', '交易记录', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (189, 1, 'zh-CN', '钱包账单明细', 2, 'billDetails', '账单明细', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (190, 1, 'zh-CN', '提交成功', 2, 'submittedSuccess', '提交成功', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (191, 1, 'zh-CN', '账号', 2, 'username', '账号', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (192, 1, 'zh-CN', '验证码', 2, 'code', '验证码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (193, 1, 'zh-CN', '旧密码', 2, 'oldPassword', '旧密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (194, 1, 'zh-CN', '新密码', 2, 'newPassword', '新密码(如果未设置请为空)', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (195, 1, 'zh-CN', '确认密码', 2, 'cmfPassword', '确认密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (196, 1, 'zh-CN', '确认安全密码', 2, 'cmfSecretKey', '确认安全密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (197, 1, 'zh-CN', '两次密码不一样', 2, 'twoPasswordsAreDifferent', '两次密码不一样', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (198, 1, 'zh-CN', '两次密钥不一样', 2, 'twoSecretKeyAreDifferent', '两次密钥不一样', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (199, 1, 'zh-CN', '请输入密钥', 2, 'enterSecretKey', '请输入密钥', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (200, 1, 'zh-CN', '邀请码', 2, 'inviteCode', '邀请码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (201, 1, 'zh-CN', '密码', 2, 'password', '密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (202, 1, 'zh-CN', '密码副标题', 2, 'passwordSmall', '编辑登录密码, 确保您的账户安全。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (203, 1, 'zh-CN', '安全密码', 2, 'secretKey', '安全密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (204, 1, 'zh-CN', '安全密码副标题', 2, 'secretKeySmall', '编辑安全密钥, 确保您的账户资金安全。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (205, 1, 'zh-CN', '邮箱', 2, 'email', '邮箱地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (206, 1, 'zh-CN', '邮箱副标题', 2, 'emailSmall', '绑定您的邮箱~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (207, 1, 'zh-CN', '手机号码', 2, 'telephone', '手机号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (208, 1, 'zh-CN', '手机号码副标题', 2, 'telephoneSmall', '绑定您的手机号码~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (209, 1, 'zh-CN', '头像', 2, 'avatar', '头像', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (210, 1, 'zh-CN', '头像副标题', 2, 'avatarSmall', '编辑您的头像, 更能体现您的气质~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (211, 1, 'zh-CN', '昵称', 2, 'nickname', '昵称', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (212, 1, 'zh-CN', '昵称副标题', 2, 'nicknameSmall', '编辑自己昵称, 使用个性化昵称~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (213, 1, 'zh-CN', '性别', 2, 'sex', '性别', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (214, 1, 'zh-CN', '男', 2, 'male', '男', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (215, 1, 'zh-CN', '女', 2, 'female', '女', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (216, 1, 'zh-CN', '未知', 2, 'unknown', '未知', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (217, 1, 'zh-CN', '性别副标题', 2, 'sexSmall', '编辑性别~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (218, 1, 'zh-CN', '生日', 2, 'birthday', '生日', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (219, 1, 'zh-CN', '生日副标题', 2, 'birthdaySmall', '设置您的生日, 让我们给你一份惊喜~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (220, 1, 'zh-CN', '个性签名', 2, 'personalSignature', '个性签名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (221, 1, 'zh-CN', '个性签名副标题', 2, 'personalSignatureSmall', '让更多用户知道您的兴趣爱好~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (222, 1, 'zh-CN', '证件名称', 2, 'idName', '证件名称', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (223, 1, 'zh-CN', '证件号码', 2, 'idNumber', '证件号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (224, 1, 'zh-CN', '证件照片', 2, 'idPhoto', '证件照片', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (225, 1, 'zh-CN', '证件正面', 2, 'idPhoto1', '证件照正面', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (226, 1, 'zh-CN', '证件反面', 2, 'idPhoto2', '证件照反面', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (227, 1, 'zh-CN', '手持证件', 2, 'idPhoto3', '手持证件照', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (228, 1, 'zh-CN', '详细地址', 2, 'address', '详细地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (229, 1, 'zh-CN', '银行名称', 2, 'bankName', '银行名称', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (230, 1, 'zh-CN', '归属姓名', 2, 'ownerName', '真实姓名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (231, 1, 'zh-CN', '银行卡号', 2, 'bankNumber', '银行卡号', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (232, 1, 'zh-CN', '银行地址', 2, 'bankAddress', '银行地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (233, 1, 'zh-CN', '网络名称', 2, 'digitalNetwork', '数字网络', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (234, 1, 'zh-CN', '数字地址', 2, 'digitalAddress', '数字地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (235, 1, 'zh-CN', '设置', 2, 'setting', '设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (236, 1, 'zh-CN', '个人设置', 2, 'personalSetting', '个人设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (237, 1, 'zh-CN', '修改登录密码', 2, 'updatePassword', '登录密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (238, 1, 'zh-CN', '修改安全密码', 2, 'updateSecretKey', '安全密钥', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (239, 1, 'zh-CN', '绑定邮箱', 2, 'bindEmail', '绑定邮箱', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (240, 1, 'zh-CN', '绑定手机号码', 2, 'bindTelephone', '绑定手机号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (241, 1, 'zh-CN', '绑定手机号码副标题', 2, 'bindTelephoneSmall', '请输入您的有效电话号码。 我们将向您发送 4 位代码来验证帐户。输入电话号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (242, 1, 'zh-CN', '语言', 2, 'language', '语言', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (243, 1, 'zh-CN', '关于', 2, 'about', '关于', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (244, 1, 'zh-CN', '新的通知', 2, 'noticeEnable', '新的通知', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (245, 1, 'zh-CN', '信用分', 2, 'creditScore', '信用分', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (246, 1, 'zh-CN', '未实名', 2, 'alreadyRealName', '未实名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (247, 1, 'zh-CN', '实名失败', 2, 'notRealName', '验证失败', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (248, 1, 'zh-CN', '实名审核中', 2, 'pendingRealName', '审核中', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (249, 1, 'zh-CN', '已实名', 2, 'realNameFailed', '已实名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (250, 1, 'zh-CN', '邀请总数', 2, 'inviteNums', '邀请总数', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (251, 1, 'zh-CN', '今日人数', 2, 'todayNums', '今日人数', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (252, 1, 'zh-CN', '购买总额', 2, 'buyAmount', '购买总额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (253, 1, 'zh-CN', '今日总额', 2, 'todayAmount', '今日总额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (254, 1, 'zh-CN', '团队收益', 2, 'teamEarnings', '团队收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (255, 1, 'zh-CN', '今日收益', 2, 'todayEarnings', '今日收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (256, 2, 'zh-CN', '验证参数不能为空', 1, 'required', '参数不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (257, 2, 'zh-CN', '验证参数min', 1, 'min', '参数最小值为', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (258, 2, 'zh-CN', '验证参数max', 1, 'max', '参数最大值为', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (259, 2, 'zh-CN', '验证参数oneof', 1, 'oneof', '范围不匹配', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (260, 2, 'zh-CN', '验证参数大于', 1, 'gt', '必须大于', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (261, 2, 'zh-CN', '验证参数大于等于', 1, 'gte', '必须大于等于', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (262, 2, 'zh-CN', '未设置安全密钥', 1, 'notSecurityKey', '未设置安全密钥, 请前往设置...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (263, 2, 'zh-CN', '参数不正确', 1, 'incorrectFormat', '格式不正确, 检查当前格式~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (264, 2, 'zh-CN', '验证码不正确', 1, 'incorrectCode', '验证码不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (265, 2, 'zh-CN', '旧密码不正确', 1, 'TheOldPasswordIsIncorrect', '旧密码不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (266, 2, 'zh-CN', '旧安全密钥不正确', 1, 'TheOldSecurityKeyIsIncorrect', '旧安全密钥不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (267, 2, 'zh-CN', '账号或密码不正确', 1, 'accountOrPasswordIncorrect', '账号或密码不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (268, 2, 'zh-CN', '安全密钥不能为空', 1, 'securityBeEmpty', '安全密钥不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (269, 2, 'zh-CN', '安全密钥不正确', 1, 'securityKeyIncorrect', '安全密钥不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (270, 2, 'zh-CN', '手机号码不能为空', 1, 'telephoneBeEmpty', '手机号码不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (271, 2, 'zh-CN', '邮箱不能为空', 1, 'emailBeEmpty', '邮箱不能为空', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (272, 2, 'zh-CN', '邀请码不存在', 1, 'inviteBeEmpty', '邀请码不存在', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (273, 2, 'zh-CN', '暂时不能注册', 1, 'notRegister', '暂时不能注册', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (274, 2, 'zh-CN', '账单类型不正确', 1, 'wrongBillType', '账单类型不正确', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (275, 2, 'zh-CN', '用户名不存在', 1, 'usernameNotFound', '用户名不存在', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (276, 2, 'zh-CN', '支付方式不存在', 1, 'paymentNotFound', '支付方式不存在', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (277, 2, 'zh-CN', '余额不足', 1, 'insufficientBalance', '余额不足', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (278, 2, 'zh-CN', '未绑定卡片', 1, 'accountBeEmpty', '未绑定卡片', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (279, 2, 'zh-CN', '超出绑定限制', 1, 'cardAdditionLimitExceeded', '超出绑定限制', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (280, 2, 'zh-CN', '超出提现次数', 1, 'exceededNumberOfWithdrawals', '提款次数超出', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (281, 2, 'zh-CN', '金额范围不匹配', 1, 'amountBetweenMismatch', '金额范围不匹配', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (282, 2, 'zh-CN', '账单类型充值', 1, 'walletBillTypeDeposit', '余额充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (283, 2, 'zh-CN', '账单类型提现', 1, 'walletBillTypeWithdraw', '余额提现', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (284, 2, 'zh-CN', '账单类型资产充值', 1, 'walletBillTypeAssetsDeposit', '资产充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (285, 2, 'zh-CN', '账单类型资产提现', 1, 'walletBillTypeAssetsWithdraw', '资产提现', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (286, 2, 'zh-CN', '账单类型购买产品', 1, 'walletBillTypeBuyProduct', '购买产品', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (287, 2, 'zh-CN', '账单类型购买等级', 1, 'walletBillTypeBuyLevel', '购买等级', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (288, 2, 'zh-CN', '账单类型收益', 1, 'walletBillTypeEarnings', '收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (289, 2, 'zh-CN', '账单类型奖励', 1, 'walletBillTypeAward', '奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (290, 2, 'zh-CN', '账单类型资产购买产品', 1, 'walletBillTypeAssetsBuyProduct', '资产购买产品', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (291, 2, 'zh-CN', '账单类型资产收益', 1, 'walletBillTypeAssetsEarnings', '资产收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (292, 2, 'zh-CN', '账单类型资产奖励', 1, 'walletBillTypeAssetsAward', '资产奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (293, 2, 'zh-CN', '账单类型注册奖励', 1, 'walletBillTypeRegisterAward', '注册奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (294, 2, 'zh-CN', '账单类型邀请奖励', 1, 'walletBillTypeShareAward', '邀请奖励', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (295, 2, 'zh-CN', '账单类型系统充值', 1, 'walletBillTypeSystemDeposit', '余额系统充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (296, 2, 'zh-CN', '账单类型资产系统充值', 1, 'walletBillTypeSystemAssetsDeposit', '资产系统充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (297, 2, 'zh-CN', '账单类型系统扣款', 1, 'walletBillTypeSystemWithdraw', '余额系统扣款', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (298, 2, 'zh-CN', '账单类型资产系统扣款', 1, 'walletBillTypeSystemAssetsWithdraw', '资产系统扣款', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (299, 2, 'zh-CN', '账单类型团队收益', 1, 'walletBillTypeTeamEarnings', '团队收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (300, 2, 'zh-CN', '账单类型团队资产收益', 1, 'walletBillTypeTeamAssetsEarnings', '团队资产收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (301, 2, 'zh-CN', '账单类型资金提现拒绝', 1, 'walletBillTypeWithdrawRefuse', '资金提现拒绝', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (302, 2, 'zh-CN', '账单类型资产提现拒绝', 1, 'walletBillTypeAssetsWithdrawRefuse', '资产提现拒绝', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (303, 2, 'zh-CN', 'level1标题', 1, 'level1Label', 'VIP1', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (304, 2, 'zh-CN', 'level2标题', 1, 'level2Label', 'VIP2', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (305, 2, 'zh-CN', 'level3标题', 1, 'level3Label', 'VIP3', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (306, 2, 'zh-CN', 'level4标题', 1, 'level4Label', 'VIP4', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (307, 2, 'zh-CN', 'level5标题', 1, 'level5Label', 'VIP5', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (308, 2, 'zh-CN', 'level1详情', 1, 'level1Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (309, 2, 'zh-CN', 'level2详情', 1, 'level2Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (310, 2, 'zh-CN', 'level3详情', 1, 'level3Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (311, 2, 'zh-CN', 'level4详情', 1, 'level4Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (312, 2, 'zh-CN', 'level5详情', 1, 'level5Desc', 'VIP等级介绍详情信息...', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (313, 2, 'zh-CN', '储蓄卡', 1, 'DebitCard', '储蓄卡', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (314, 2, 'zh-CN', '美国银行', 1, 'BankOfAmerica', '美国银行', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (315, 2, 'zh-CN', '储蓄卡(资金提现)', 1, 'bankFundsWithdraw', '储蓄卡[资金提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (316, 2, 'zh-CN', '储蓄卡(资产提现)', 1, 'bankAssetsWithdraw', '储蓄卡[资产提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (317, 2, 'zh-CN', '数字货币(资金提现)', 1, 'digitalFundsWithdraw', '数字货币[资金提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (318, 2, 'zh-CN', '数字货币(资产提现)', 1, 'digitalAssetsWithdraw', '数字货币[资产提现]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (319, 2, 'zh-CN', '储蓄卡(资金充值)', 1, 'bankFundsDeposit', '储蓄卡[资金充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (320, 2, 'zh-CN', '储蓄卡(资产充值)', 1, 'bankAssetsDeposit', '储蓄卡[资产充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (321, 2, 'zh-CN', '数字货币(资金充值)', 1, 'digitalFundsDeposit', '数字货币[资金充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (322, 2, 'zh-CN', '数字货币(资产充值)', 1, 'digitalAssetsDeposit', '数字货币[资产充值]', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (323, 2, 'zh-CN', '充值提示详情', 1, 'depositDesc', '充值注意事项', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (324, 2, 'zh-CN', '提现提示详情', 1, 'withdrawDesc', '提现注意事项', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (325, 2, 'zh-CN', '站点介绍', 1, 'siteIntroduce', '站点介绍', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (326, 2, 'zh-CN', '站点通知', 1, 'siteAnnouncement', '站点通知', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (327, 2, 'zh-CN', '关于我们', 1, 'aboutUs', '关于我们', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (328, 2, 'zh-CN', '平台服务', 1, 'siteService', '平台服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (329, 2, 'zh-CN', '联系我们', 1, 'contactUs', '联系我们', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (330, 2, 'zh-CN', '社区', 1, 'social', '社区', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (331, 2, 'zh-CN', '关于我们标题', 1, 'aboutUsLabel', '关于我们', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (332, 2, 'zh-CN', '关于我们详情', 1, 'aboutUsDesc', '<pre><span style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">关于我们</font></font></span><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">欢迎来到创新与卓越相遇的地方！ 我们是一家先锋科技企业，致力于突破技术界限并为客户提供卓越的解决方案。 我们的承诺不仅仅是创新； 这是关于了解客户需求并承担社会责任。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的使命和核心价值观</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">在[公司名称]，我们的使命是通过创新为未来铺平道路。 我们的核心价值观决定了我们的成功：</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">技术创新：不断探索和采用前沿技术，提供最先进的解决方案和服务。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">以客户为中心的方法：通过提供满足客户独特需求的高质量定制服务来优先考虑客户的成功。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">协作团队合作：我们的团队由顶级专业人士组成，我们相信协作是成功的关键，他们不断协作以实现共同成功。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">社会责任：作为企业责任的一部分，我们积极支持可持续发展努力，为社会进步做出贡献。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的服务范围</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们提供跨多个领域的全面服务：</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">软件开发与 工程：利用最新的软件开发技术，我们提供高效、可靠和定制的软件解决方案。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">人工智能与 数据分析：专注于人工智能和数据分析，利用洞察和智能技术提供定制化解决方案。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">物联网 (IoT) 和 未来技术：我们致力于推进物联网和未来技术，创造智能解决方案和创新产品。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">数字化转型服务：协助企业数字化转型，利用技术提高效率、降低成本、创造更多商业价值。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">客户关系</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">与客户建立持久的合作伙伴关系是我们的首要任务。 我们用心倾听他们的需求，提供个性化的解决方案，并提供持续的支持以推动他们的持续增长。 积极寻求并纳入客户反馈可确保我们提供最优质的服务和技术支持。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">社会责任承诺</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们不仅仅是一家科技公司； 我们是社会责任的管家。 参与社区活动、支持环境保护和支持社会事业是我们致力于社会进步和可持续发展的重要组成部分。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">与我们合作</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">无论您的业务阶段或需求如何，我们都致力于成为您值得信赖的合作伙伴，提供最佳的技术解决方案。 如果您想进一步探索我们的服务或与我们合作，请随时与我们联系。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的创新方法</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">创新不仅仅是一个流行词，而是一个新词。 这是支撑我们所做一切的理念。 这是我们运营的基石。 对创新的不懈追求促使我们不断探索新技术和新方法，确保我们保持领先优势。 我们大力投资研发，培育鼓励创造力、实验和创新思维的文化。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们成功背后的团队</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们的成功归功于一支多元化且才华横溢的团队。 他们由行业专家、远见卓识者和具有前瞻性思维的个人组成，他们在挑战中不断成长，并无缝协作，提供突破性的解决方案。 我们重视多样性、包容性和持续学习，创造一个促进增长和创新的环境。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">行业认可和成就</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们对卓越和创新的承诺赢得了业界的认可。 我们对我们的成就感到自豪，不仅将它们视为里程碑，而且将其视为进一步前进的灵感。 我们的业绩记录意义重大，反映出我们致力于提供超出预期的成果。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">道德实践和透明度</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">透明度和道德实践构成了我们运营的基石。 我们与客户、合作伙伴和利益相关者保持开放的沟通，确保每次互动中的信任和可靠性。 我们对诚信的承诺指导着我们所有的业务往来。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">可持续发展和环境意识</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">作为一个负责任的企业实体，我们在运营的各个方面都优先考虑可持续发展。 我们将环保实践融入我们的流程中，努力最大限度地减少我们的环境足迹。 此外，我们积极支持促进环境保护和生态意识的举措。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">拥抱变革，共创美好未来</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">我们将变革视为进步的催化剂。 我们迅速适应不断发展的技术格局，预见趋势，并为我们自己和我们的客户做好迎接未来挑战和机遇的准备。</font></font></div><div style=\"font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><br></div><div style=\"text-align: left; font-family: Roboto, -apple-system, Avenir, &quot;system-ui&quot;, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif; white-space-collapse: collapse;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">加入我们的创新之旅</font></font></div></pre>', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (333, 2, 'zh-CN', '服务协议标题', 1, 'serviceAgreementLabel', '服务协议', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (334, 2, 'zh-CN', '服务协议详情', 1, 'serviceAgreementDesc', '服务描述\n平台同意按照本协议的条款和条件向用户提供以下服务：\n*   账户管理 - 用户可以在平台上创建和管理个人交易账户，包括存储、兑换和交易数字货币。\n*   交易服务 - 平台提供数字货币交易服务，包括买卖订单管理、市场信息提供和交易执行。\n*   安全与合规 - 平台致力于提供安全的交易环境，并遵守相关的法律法规。\n*   客户支持 - 用户可以通过客户支持获取交易相关的帮助和咨询。\n用户责任\n*   账户安全 - 用户负责保护其账户信息的安全，包括用户名和密码。\n*   合规操作 - 用户同意遵守所有适用的法律、法规和平台的规则。\n*   风险认知 - 用户理解并同意，数字货币交易存在市场风险，包括但不限于价格波动、技术问题和法律风险。\n风险披露\n用户明白，数字货币交易涉及高风险，可能导致部分或全部资金的损失。用户应根据自己的财务状况和风险承受能力进行交易。\n知识产权\n平台的所有内容和技术，包括但不限于软件、界面设计、商标和版权均属于平台或其授权人所有。\n免责声明\n平台不对由于系统故障、网络问题或其他不可抗力因素导致的服务中断或交易问题承担责任。\n修改和终止\n平台保留随时修改本协议或服务内容的权利，并会提前通知用户。用户继续使用服务即表示接受修改后的协议。\n适用法律和争端解决\n本协议的解释和执行适用 。任何由本协议引起的争端应通过友好协商解决，如协商不成，可提交法院诉讼解决。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (335, 2, 'zh-CN', '隐私协议标题', 1, 'privacyAgreementLabel', '隐私协议', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (336, 2, 'zh-CN', '隐私协议详情', 1, 'privacyAgreementDesc', '<div style=\"text-align: left;\"><span style=\"font-family: Roboto, -apple-system, Avenir, BlinkMacSystemFont, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif;\"><font style=\"vertical-align: inherit;\"><font style=\"vertical-align: inherit;\">尊敬的用户欢迎使用 我们的产品和服务。 我们深知您对个人信息安全和隐私的重视，我们将尽力保护您个人信息的安全，并遵循以下隐私政策： 收集信息的目的：我们收集您的个人信息 为了更好地为您提供服务，提升用户体验，保证产品和服务的安全。 我们只会收集必要的信息，绝不会超出声明目的使用您的个人信息。 收集的信息类型：我们可能收集的信息类型包括但不限于：姓名、电子邮件地址、联系信息、IP地址、设备信息、浏览器类型等信息使用和共享：我们只会使用您的个人信息 在提供服务和提高产品质量的范围内。 除非获得您的同意或法律法规要求，我们不会向任何第三方出售、出租或分享您的个人信息。 信息安全：我们采取了一系列安全措施来保护您的个人信息，包括但不限于加密技术、访问控制、安全审计等。我们将尽力确保您的信息安全，不被未经授权的访问、使用 或披露。 用户权利：您有权查询、更正、删除个人信息。 如果您对我们收集的个人信息有任何疑问或需要行使相关权利，请随时联系我们的客服团队。 隐私政策的变更：随着我们的产品和服务的发展，我们的隐私政策可能会不时更新。 我们将在网站上发布任何政策变更，并根据适用法律的要求向您提供通知。 接受条款：使用我们的产品和服务，即表示您同意本隐私政策的条款和条件。 联系我们：如果您对我们的隐私政策有任何疑问或建议，请随时联系我们的隐私团队。 感谢您对我们的信任，我们将继续努力确保您的信息安全和隐私</font></font></span></div>', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (337, 2, 'zh-CN', '帮助中心1标题', 1, 'helpers1Label', '重置密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (338, 2, 'zh-CN', '帮助中心1详情', 1, 'helpers1Desc', '如果您忘记了密码，可以轻松重置。请在登录页面点击“忘记密码？”链接，然后输入您的注册电子邮件地址。我们将发送一封包含重置密码指令的电子邮件给您。遵循邮件中的指示，您就可以设置新密码。为了保证账户安全，请确保您的新密码足够复杂，且不与其他网站共享相同密码。如果在重置密码过程中遇到任何问题，欢迎联系我们的客户支持团队获取帮助。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (339, 2, 'zh-CN', '帮助中心2标题', 1, 'helpers2Label', '安全建议', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (340, 2, 'zh-CN', '帮助中心2详情', 1, 'helpers2Desc', '为了保护您的账户安全，我们建议您采取以下措施：首先，使用强密码，结合大写字母、小写字母、数字和特殊符号。其次，定期更换密码，避免使用相同的密码登录不同网站。此外，启用两因素认证（2FA）增加额外的安全层。请勿在公共计算机或不安全的网络环境中登录账户，并始终保持您的个人信息和联系方式更新。最后，警惕任何可疑的邮件或信息，避免点击不明链接或分享敏感信息。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (341, 2, 'zh-CN', '帮助中心3标题', 1, 'helpers3Label', '取消订阅', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (342, 2, 'zh-CN', '帮助中心3详情', 1, 'helpers3Desc', '取消订阅非常简单。登录您的账户后，前往账户设置 页面，在那里您会找到“订阅管理”选项。选择您想要取消的服务，并点击“取消订阅”按钮。根据您的订阅类型，可能会有一段提前通知期。请注意，取消订阅可能会影响到您享受的某些服务或优惠。如果您在取消订阅过程中遇到任何问题或需要进一步的帮助，欢迎随时联系我们的客户服务团队。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (343, 2, 'zh-CN', '帮助中心4标题', 1, 'helpers4Label', '支付方式', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (344, 2, 'zh-CN', '帮助中心4详情', 1, 'helpers4Desc', '为了方便用户，我们支持多种支付方式。您可以使用信用卡（Visa, MasterCard, American Express等）、借记卡、电子钱包（如PayPal, Apple Pay等）以及银行转账等方式进行支付。所有支付过程都是经过加密的，确保您的交易信息安全。请注意，根据您所在的地区，可用的支付方式可能会有所不同。如果您在支付过程中遇到任何问题，或者需要了解更多关于支付选项的信息，请联系我们的客户支持团队。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (345, 2, 'zh-CN', '帮助中心5标题', 1, 'helpers5Label', '退货政策', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (346, 2, 'zh-CN', '帮助中心5详情', 1, 'helpers5Desc', '我们承诺提供满意的退货政策。如果您对购买的产品不满意，可以在收到货物后30天内申请退货。要开始退货流程，请登录您的账户并前往“订单历史”页面，选择相应的订单并点击“申请退货”。请确保退回的商品处于未使用状态，并包含所有原包装和附件。退货处理完成后，我们将退还您的支付金额，退款通常在5至10个工作日内处理完毕。如果您在退货过程中需要任何帮助，请随时联系我们的客户服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (347, 2, 'zh-CN', '帮助中心6标题', 1, 'helpers6Label', '追踪订单', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (348, 2, 'zh-CN', '帮助中心6详情', 1, 'helpers6Desc', '在您的订单被发货后，您将通过电子邮件收到一个包含追踪链接的发货通知。点击该链接，您可以查看订单的最新运输状态和预计送达时间。此外，您也可以直接在我们的网站上通过“我的订单”部分追踪订单状态。如果您对订单的运输状态有任何疑问或需要帮助，请联系我们的客户服务团队。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (349, 2, 'zh-CN', '帮助中心7标题', 1, 'helpers7Label', '产品期限', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (350, 2, 'zh-CN', '帮助中心7详情', 1, 'helpers7Desc', '我们的产品享有一定期限的保修服务，确保您在使用过程中得到必要的支持。保修期限通常从购买日期起算，具体时长取决于产品类型。保修内容通常包括制造缺陷和材料问题。在保修期内，如果产品出现合格的保修问题，我们将提供免费修理或更换服务。请注意，保修不包括因误用、正常磨损或未经授权的维修造成的损坏。详细的保修条款和条件请参见我们的产品保修政策，或联系我们的客户服务团队获取更多信息。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (351, 2, 'zh-CN', '平台服务1标题', 1, 'service1Label', '产品服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (352, 2, 'zh-CN', '平台服务1详情', 1, 'service1Desc', '欢迎来到我们的交易所平台！我们是一家致力于为您提供安全、高效、多样化的交易体验的领先交易平台。无论您是初学者还是经验丰富的交易者，我们都为您提供了丰富的产品和服务，以满足您的需求。<br />我们的产品特点：<br />多样化的交易对: 我们支持多种数字资产交易对，涵盖主流加密货币以及新兴数字资产，让您可以灵活选择交易组合。<br />安全保障: 我们采取严格的安全措施，包括多重身份验证、冷热钱包存储体系等，保障您的资产安全。<br />高效的交易执行: 我们的平台配备高性能的交易引擎，保证交易执行速度和效率，让您能够快速响应市场变化。<br />用户友好的界面: 我们的界面简洁直观，适合各种水平的交易者使用，同时也提供了专业级别的图表和工具，满足有更高要求的用户。<br />教育资源和客户支持: 我们提供广泛的教育资源，包括文章、视频和在线课程，帮助您了解市场动态和交易策略。此外，我们的客户支持团队全天候为您提供支持。<br /><br />我们的服务承诺：<br />透明度与合规性: 我们遵循严格的法规标准，保障交易过程的透明度和合规性。<br />持续创新与发展: 我们致力于不断提升平台功能和服务，引入新的技术和工具，以满足市场不断变化的需求。<br />社区参与与奖励: 我们鼓励社区参与，通过奖励计划回馈我们的忠实用户。<br />无论您是寻求快速交易、投资多样化数字资产，还是寻找一个安全可靠的交易所平台，我们都将竭诚为您提供最优质的服务和支持。<br />加入我们，开始您的交易之旅吧！', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (353, 2, 'zh-CN', '平台服务2标题', 1, 'service2Label', '基础方案', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (354, 2, 'zh-CN', '平台服务2详情', 1, 'service2Desc', '注册和账户管理<br />浏览市场和实时数据<br />限制性的交易功能<br />专业方案 - $9.99/月：<br />全面的交易功能，包括买卖、存款和提款<br />高级图表和分析工具<br />客户支持优先服务<br />额外的教育资源和市场报告<br />企业方案 - 定制报价：<br />适用于机构和专业交易者<br />定制化的解决方案<br />安全审计和额外的安全措施<br />专属的客户经理和支持团队<br />价格策略说明：<br />弹性定价: 我们提供多种价格选择，满足不同用户群体的需求。用户可以根据其需求和预算选择适合自己的方案。<br />试用期: 对于新用户，我们提供免费试用期或者特定功能的试用，让用户可以在付费前体验我们的服务。<br />优惠和折扣: 定期推出促销活动、折扣和特惠套餐，以回馈现有用户并吸引新用户。<br />定期支付: 客户可以选择按月或按年支付，按年支付的用户可以享受一定的折扣优惠。<br />定价透明度与承诺：<br />我们承诺定价透明，没有隐藏费用或附加费用。所有价格均明确列出，让用户清楚了解他们所支付的服务内容。我们也欢迎用户反馈和建议，以不断优化我们的定价策略，以更好地满足用户需求。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (355, 2, 'zh-CN', '平台服务3标题', 1, 'service3Label', '新产品介绍', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (356, 2, 'zh-CN', '平台服务3详情', 1, 'service3Desc', '在数字资产市场的不断演变和增长中，我们引以为傲地推出了全新的「数字资产投资组合优化工具」，旨在帮助投资者最大化其投资回报并降低风险。<br />主要特点：<br />智能投资组合优化：<br />利用先进的算法和数据分析，我们的工具能够帮助用户构建最优的数字资产投资组合，考虑到风险和收益的平衡。<br />动态资产配置：<br />根据市场实时数据和用户偏好，工具能够动态调整投资组合，以应对市场波动和变化，最大程度地优化投资组合表现。<br />风险管理与模拟测试：<br />提供多种风险管理工具，让用户能够对不同的投资方案进行模拟测试，了解可能的风险和回报情况，做出更明智的投资决策。<br />个性化建议与报告：<br />提供个性化的投资建议和定期报告，帮助用户跟踪投资组合的表现，并根据实时数据做出调整。<br />如何使用我们的产品：<br />注册账户：<br />快速注册账户并登录我们的平台。<br />输入投资偏好：<br />根据您的投资目标、风险承受能力和偏好，输入相关信息。<br />获取优化建议：<br />根据您提供的信息，工具将生成最佳的数字资产投资组合建议。<br />跟踪和调整：<br />定期跟踪投资组合表现，并根据报告的建议调整您的投资策略。<br />我们的承诺：<br />我们的「数字资产投资组合优化工具」旨在为您提供更智能、更高效的投资解决方案。我们致力于提供卓越的用户体验，确保您的投资决策更加明智和可靠。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (357, 2, 'zh-CN', '未添加资产', 2, 'notAssets', '未添加资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (358, 2, 'zh-CN', '饼图', 2, 'pieChart', '饼图', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (359, 2, 'zh-CN', '折线图', 2, 'lineChart', '折线图', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (360, 2, 'zh-CN', '全部', 2, 'all', '全部', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (361, 2, 'zh-CN', '前往', 2, 'goto', '前往', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (362, 2, 'zh-CN', '暂无数据', 2, 'noData', '暂无数据', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (363, 2, 'zh-CN', '注册时间', 2, 'createdTime', '注册时间', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (364, 2, 'zh-CN', '确定操作标题', 2, 'isExecute', '您是否确定执行当前操作?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (365, 2, 'zh-CN', '确定购买', 2, 'isBuy', '您是否确定购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (366, 2, 'zh-CN', '货币单位', 2, 'currency', '$', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (367, 2, 'zh-CN', '充值', 2, 'deposit', '充值', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (368, 2, 'zh-CN', '提现', 2, 'withdraw', '提现', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (369, 2, 'zh-CN', '首页', 2, 'home', '首页', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (370, 2, 'zh-CN', '行情', 2, 'markets', '市场', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (371, 2, 'zh-CN', '合约', 2, 'contract', '合约', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (372, 2, 'zh-CN', '期货', 2, 'futures', '期货', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (373, 2, 'zh-CN', '我的', 2, 'user', '我的', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (374, 2, 'zh-CN', '数字货币', 2, 'digitalMarkets', '数字货币', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (375, 2, 'zh-CN', '外汇市场', 2, 'forexMarkets', '外汇市场', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (376, 2, 'zh-CN', '期货市场', 2, 'futuresMarkets', '期货市场', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (377, 2, 'zh-CN', '钱包管理', 2, 'walletManage', '钱包管理', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (378, 2, 'zh-CN', '团队管理', 2, 'teamManage', '团队管理', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (379, 2, 'zh-CN', '服务设置', 2, 'serviceManage', '服务设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (380, 2, 'zh-CN', '卡片管理', 2, 'accountManage', '卡片管理', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (381, 2, 'zh-CN', '我的钱包', 2, 'myWallet', '我的钱包', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (382, 2, 'zh-CN', '我的资产', 2, 'myAssets', '我的资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (383, 2, 'zh-CN', '邀请好友', 2, 'inviteFriends', '邀请好友', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (384, 2, 'zh-CN', '我的团队', 2, 'myTeam', '我的团队', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (385, 2, 'zh-CN', '团队收益', 2, 'teamEarnings', '团队收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (386, 2, 'zh-CN', '团队成员', 2, 'teamMember', '团队成员', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (387, 2, 'zh-CN', '会员中心', 2, 'memberVip', '会员中心', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (388, 2, 'zh-CN', '实名认证', 2, 'realAuth', '实名认证', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (389, 2, 'zh-CN', '帮助中心', 2, 'helpers', '帮助中心', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (390, 2, 'zh-CN', '个人设置', 2, 'settings', '个人设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (391, 2, 'zh-CN', '下载中心', 2, 'download', '下载中心', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (392, 2, 'zh-CN', '下载副标题', 2, 'downloadSmall', '使用手机版本能更好的体验服务~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (393, 2, 'zh-CN', '邀请好友副标题', 2, 'inviteFriendsSmall', '邀请您的好友一起Happy, 同时能领取到惊喜收益。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (394, 2, 'zh-CN', '会员中心副标题', 2, 'memberVipSmall', '升级成为尊贵的会员用户, 同时享受各种平台优惠服务~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (395, 2, 'zh-CN', '会员中心标题', 2, 'memberVipTitle', '选择您的会员计划', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (396, 2, 'zh-CN', '登录', 2, 'login', '登录', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (397, 2, 'zh-CN', '忘记密码', 2, 'forgotPassword', '忘记密码?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (398, 2, 'zh-CN', '登录副标题', 2, 'loginSmall', '登录您的账号, 开启旅程～', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (399, 2, 'zh-CN', '未注册去登录', 2, 'toRegister', '还没有账号, 去注册!', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (400, 2, 'zh-CN', '注册', 2, 'register', '注册', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (401, 2, 'zh-CN', '注册副标题', 2, 'registerSmall', '创建您的账户, 开启旅程~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (402, 2, 'zh-CN', '已有账户去登录', 2, 'toLogin', '已有账户, 现在去登录!', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (403, 2, 'zh-CN', '退出登录', 2, 'logout', '退出登录', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (404, 2, 'zh-CN', '退出登录副标题', 2, 'logoutSmall', '您确定要退出登录吗?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (405, 2, 'zh-CN', '有什么能帮上您的吗', 2, 'helperYou', '你好, 有什么能帮上您的吗?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (406, 2, 'zh-CN', '24小时在线服务', 2, '24hoursOnline', '24小时在线服务', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (407, 2, 'zh-CN', '更多', 2, 'more', '更多', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (408, 2, 'zh-CN', '查询', 2, 'search', '查询', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (409, 2, 'zh-CN', '提交', 2, 'submit', '提交', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (410, 2, 'zh-CN', '重新提交', 2, 'reSubmit', '重新提交', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (411, 2, 'zh-CN', '完成', 2, 'complete', '完成', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (412, 2, 'zh-CN', '审核', 2, 'pending', '审核', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (413, 2, 'zh-CN', '拒绝', 2, 'refuse', '拒绝', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (414, 2, 'zh-CN', '筛选', 2, 'filter', '筛选', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (415, 2, 'zh-CN', '类型', 2, 'type', '类型', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (416, 2, 'zh-CN', '时间范围', 2, 'betweenTime', '时间范围', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (417, 2, 'zh-CN', '余额', 2, 'balance', '余额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (418, 2, 'zh-CN', '充值金额', 2, 'depositAmount', '充值金额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (419, 2, 'zh-CN', '提现金额', 2, 'withdrawAmount', '提现金额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (420, 2, 'zh-CN', '可用金额', 2, 'availableAmount', '可用金额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (421, 2, 'zh-CN', '充值凭证', 2, 'depositProof', '充值凭证', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (422, 2, 'zh-CN', '提现账户', 2, 'withdrawAccount', '提现账户', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (423, 2, 'zh-CN', '未绑定提现账户', 2, 'notBindWithdrawAccount', '未绑定提现账户', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (424, 2, 'zh-CN', '充值帐户', 2, 'depositAccount', '充值帐户', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (425, 2, 'zh-CN', '账户信息', 2, 'depositAccountInfo', '账户信息', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (426, 2, 'zh-CN', '总资产', 2, 'totalAssets', '总资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (427, 2, 'zh-CN', '资产分布', 2, 'assetsBlock', '资产分布', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (428, 2, 'zh-CN', '资产', 2, 'assets', '资产', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (429, 2, 'zh-CN', '取消', 2, 'cancel', '取消', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (430, 2, 'zh-CN', '确定', 2, 'confirm', '确定', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (431, 2, 'zh-CN', '删除', 2, 'delete', '删除', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (432, 2, 'zh-CN', '删除确认标题', 2, 'deleteLabel', '删除确定', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (433, 2, 'zh-CN', '是否确定删除当前记录', 2, 'deleteRecord', '是否确定删除当前记录?', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (434, 2, 'zh-CN', '新增', 2, 'create', '新增', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (435, 2, 'zh-CN', '编辑', 2, 'edit', '编辑', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (436, 2, 'zh-CN', '详情', 2, 'desc', '详情', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (437, 2, 'zh-CN', '查看', 2, 'views', '查看', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (438, 2, 'zh-CN', '复制', 2, 'copy', '复制', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (439, 2, 'zh-CN', '复制成功', 2, 'copySuccess', '复制成功', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (440, 2, 'zh-CN', '购买', 2, 'buy', '购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (441, 2, 'zh-CN', '立即购买', 2, 'nowBuy', '立即购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (442, 2, 'zh-CN', '已购买', 2, 'purchased', '已购买', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (443, 2, 'zh-CN', '钱包交易记录', 2, 'transactions', '交易记录', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (444, 2, 'zh-CN', '钱包账单明细', 2, 'billDetails', '账单明细', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (445, 2, 'zh-CN', '提交成功', 2, 'submittedSuccess', '提交成功', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (446, 2, 'zh-CN', '账号', 2, 'username', '账号', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (447, 2, 'zh-CN', '验证码', 2, 'code', '验证码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (448, 2, 'zh-CN', '旧密码', 2, 'oldPassword', '旧密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (449, 2, 'zh-CN', '新密码', 2, 'newPassword', '新密码(如果未设置请为空)', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (450, 2, 'zh-CN', '确认密码', 2, 'cmfPassword', '确认密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (451, 2, 'zh-CN', '确认安全密码', 2, 'cmfSecretKey', '确认安全密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (452, 2, 'zh-CN', '两次密码不一样', 2, 'twoPasswordsAreDifferent', '两次密码不一样', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (453, 2, 'zh-CN', '两次密钥不一样', 2, 'twoSecretKeyAreDifferent', '两次密钥不一样', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (454, 2, 'zh-CN', '请输入密钥', 2, 'enterSecretKey', '请输入密钥', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (455, 2, 'zh-CN', '邀请码', 2, 'inviteCode', '邀请码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (456, 2, 'zh-CN', '密码', 2, 'password', '密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (457, 2, 'zh-CN', '密码副标题', 2, 'passwordSmall', '编辑登录密码, 确保您的账户安全。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (458, 2, 'zh-CN', '安全密码', 2, 'secretKey', '安全密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (459, 2, 'zh-CN', '安全密码副标题', 2, 'secretKeySmall', '编辑安全密钥, 确保您的账户资金安全。', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (460, 2, 'zh-CN', '邮箱', 2, 'email', '邮箱地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (461, 2, 'zh-CN', '邮箱副标题', 2, 'emailSmall', '绑定您的邮箱~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (462, 2, 'zh-CN', '手机号码', 2, 'telephone', '手机号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (463, 2, 'zh-CN', '手机号码副标题', 2, 'telephoneSmall', '绑定您的手机号码~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (464, 2, 'zh-CN', '头像', 2, 'avatar', '头像', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (465, 2, 'zh-CN', '头像副标题', 2, 'avatarSmall', '编辑您的头像, 更能体现您的气质~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (466, 2, 'zh-CN', '昵称', 2, 'nickname', '昵称', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (467, 2, 'zh-CN', '昵称副标题', 2, 'nicknameSmall', '编辑自己昵称, 使用个性化昵称~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (468, 2, 'zh-CN', '性别', 2, 'sex', '性别', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (469, 2, 'zh-CN', '男', 2, 'male', '男', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (470, 2, 'zh-CN', '女', 2, 'female', '女', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (471, 2, 'zh-CN', '未知', 2, 'unknown', '未知', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (472, 2, 'zh-CN', '性别副标题', 2, 'sexSmall', '编辑性别~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (473, 2, 'zh-CN', '生日', 2, 'birthday', '生日', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (474, 2, 'zh-CN', '生日副标题', 2, 'birthdaySmall', '设置您的生日, 让我们给你一份惊喜~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (475, 2, 'zh-CN', '个性签名', 2, 'personalSignature', '个性签名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (476, 2, 'zh-CN', '个性签名副标题', 2, 'personalSignatureSmall', '让更多用户知道您的兴趣爱好~', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (477, 2, 'zh-CN', '证件名称', 2, 'idName', '证件名称', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (478, 2, 'zh-CN', '证件号码', 2, 'idNumber', '证件号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (479, 2, 'zh-CN', '证件照片', 2, 'idPhoto', '证件照片', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (480, 2, 'zh-CN', '证件正面', 2, 'idPhoto1', '证件照正面', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (481, 2, 'zh-CN', '证件反面', 2, 'idPhoto2', '证件照反面', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (482, 2, 'zh-CN', '手持证件', 2, 'idPhoto3', '手持证件照', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (483, 2, 'zh-CN', '详细地址', 2, 'address', '详细地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (484, 2, 'zh-CN', '银行名称', 2, 'bankName', '银行名称', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (485, 2, 'zh-CN', '归属姓名', 2, 'ownerName', '真实姓名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (486, 2, 'zh-CN', '银行卡号', 2, 'bankNumber', '银行卡号', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (487, 2, 'zh-CN', '银行地址', 2, 'bankAddress', '银行地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (488, 2, 'zh-CN', '网络名称', 2, 'digitalNetwork', '数字网络', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (489, 2, 'zh-CN', '数字地址', 2, 'digitalAddress', '数字地址', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (490, 2, 'zh-CN', '设置', 2, 'setting', '设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (491, 2, 'zh-CN', '个人设置', 2, 'personalSetting', '个人设置', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (492, 2, 'zh-CN', '修改登录密码', 2, 'updatePassword', '登录密码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (493, 2, 'zh-CN', '修改安全密码', 2, 'updateSecretKey', '安全密钥', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (494, 2, 'zh-CN', '绑定邮箱', 2, 'bindEmail', '绑定邮箱', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (495, 2, 'zh-CN', '绑定手机号码', 2, 'bindTelephone', '绑定手机号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (496, 2, 'zh-CN', '绑定手机号码副标题', 2, 'bindTelephoneSmall', '请输入您的有效电话号码。 我们将向您发送 4 位代码来验证帐户。输入电话号码', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (497, 2, 'zh-CN', '语言', 2, 'language', '语言', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (498, 2, 'zh-CN', '关于', 2, 'about', '关于', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (499, 2, 'zh-CN', '新的通知', 2, 'noticeEnable', '新的通知', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (500, 2, 'zh-CN', '信用分', 2, 'creditScore', '信用分', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (501, 2, 'zh-CN', '未实名', 2, 'alreadyRealName', '未实名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (502, 2, 'zh-CN', '实名失败', 2, 'notRealName', '验证失败', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (503, 2, 'zh-CN', '实名审核中', 2, 'pendingRealName', '审核中', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (504, 2, 'zh-CN', '已实名', 2, 'realNameFailed', '已实名', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (505, 2, 'zh-CN', '邀请总数', 2, 'inviteNums', '邀请总数', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (506, 2, 'zh-CN', '今日人数', 2, 'todayNums', '今日人数', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (507, 2, 'zh-CN', '购买总额', 2, 'buyAmount', '购买总额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (508, 2, 'zh-CN', '今日总额', 2, 'todayAmount', '今日总额', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (509, 2, 'zh-CN', '团队收益', 2, 'teamEarnings', '团队收益', 1709021277);
INSERT INTO `translate` (`id`, `admin_id`, `lang`, `name`, `type`, `field`, `value`, `created_at`) VALUES (510, 2, 'zh-CN', '今日收益', 2, 'todayEarnings', '今日收益', 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL DEFAULT '1' COMMENT '管理ID',
  `parent_id` int unsigned NOT NULL COMMENT '父级ID',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `nickname` varchar(60) NOT NULL COMMENT '昵称',
  `email` varchar(60) DEFAULT NULL COMMENT '邮箱',
  `telephone` varchar(50) DEFAULT NULL COMMENT '手机号码',
  `avatar` varchar(120) NOT NULL COMMENT '头像',
  `score` tinyint NOT NULL DEFAULT '100' COMMENT '信用分',
  `sex` tinyint NOT NULL COMMENT '性别 1男 2女',
  `birthday` int unsigned NOT NULL COMMENT '生日',
  `password` varchar(120) NOT NULL COMMENT '密码',
  `security_key` varchar(120) NOT NULL COMMENT '密钥',
  `money` decimal(12,2) NOT NULL COMMENT '金额',
  `type` tinyint NOT NULL DEFAULT '11' COMMENT '类型 1虚拟用户 11默认用户 21会员用户',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1冻结 10激活',
  `data` text COMMENT '数据',
  `desc` text COMMENT '详情',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_user_name` (`username`),
  UNIQUE KEY `idx_user_email` (`email`),
  UNIQUE KEY `idx_user_telephone` (`telephone`)
) ENGINE=InnoDB AUTO_INCREMENT=24587 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `admin_id`, `parent_id`, `username`, `nickname`, `email`, `telephone`, `avatar`, `score`, `sex`, `birthday`, `password`, `security_key`, `money`, `type`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (1, 2, 0, 'ceshi', 'Jsona', 'muiprosperyls15@gmail.com', NULL, '', 100, 0, 0, '57f231b1ec41dc6641270cb09a56f897', '57f231b1ec41dc6641270cb09a56f897', 0.00, 11, 10, '', '', 1709021278, 1709021278);
INSERT INTO `user` (`id`, `admin_id`, `parent_id`, `username`, `nickname`, `email`, `telephone`, `avatar`, `score`, `sex`, `birthday`, `password`, `security_key`, `money`, `type`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (2, 2, 1, 'ceshi1', 'Jsona', NULL, NULL, '', 100, 0, 0, '57f231b1ec41dc6641270cb09a56f897', '57f231b1ec41dc6641270cb09a56f897', 100.00, 11, 10, '', '', 1709021278, 1709021278);
INSERT INTO `user` (`id`, `admin_id`, `parent_id`, `username`, `nickname`, `email`, `telephone`, `avatar`, `score`, `sex`, `birthday`, `password`, `security_key`, `money`, `type`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (3, 2, 1, 'ceshi2', 'Jsona', NULL, NULL, '', 100, 0, 0, '57f231b1ec41dc6641270cb09a56f897', '57f231b1ec41dc6641270cb09a56f897', 100.00, 11, 10, '', '', 1709021278, 1709021278);
INSERT INTO `user` (`id`, `admin_id`, `parent_id`, `username`, `nickname`, `email`, `telephone`, `avatar`, `score`, `sex`, `birthday`, `password`, `security_key`, `money`, `type`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (4, 2, 1, 'ceshi3', 'Jsona', NULL, NULL, '', 100, 0, 0, '57f231b1ec41dc6641270cb09a56f897', '57f231b1ec41dc6641270cb09a56f897', 100.00, 11, 10, '', '', 1709021278, 1709021278);
INSERT INTO `user` (`id`, `admin_id`, `parent_id`, `username`, `nickname`, `email`, `telephone`, `avatar`, `score`, `sex`, `birthday`, `password`, `security_key`, `money`, `type`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (5, 2, 1, 'ceshi4', 'Jsona', NULL, NULL, '', 100, 0, 0, '57f231b1ec41dc6641270cb09a56f897', '57f231b1ec41dc6641270cb09a56f897', 100.00, 11, 10, '', '', 1709021278, 1709021278);
COMMIT;

-- ----------------------------
-- Table structure for wallet_assets
-- ----------------------------
DROP TABLE IF EXISTS `wallet_assets`;
CREATE TABLE `wallet_assets` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `symbol` varchar(60) NOT NULL COMMENT '标识',
  `icon` varchar(60) NOT NULL COMMENT '图标',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型 1法币资产 11数字货币资产 21虚拟货币资产',
  `rate` decimal(12,2) NOT NULL COMMENT '汇率',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包资产管理表';

-- ----------------------------
-- Records of wallet_assets
-- ----------------------------
BEGIN;
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 1, 'USDT', 'USDT', '/assets/currency/usdt.png', 11, 1.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 1, 'USDC', 'USDC', '/assets/currency/usdc.png', 11, 1.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (3, 1, 'BTC', 'BTC', '/assets/currency/btc.png', 11, 38546.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (4, 1, 'ETH', 'ETH', '/assets/currency/eth.png', 11, 2090.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (5, 1, 'TRX', 'TRX', '/assets/currency/trx.png', 11, 0.10, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (6, 2, 'USDT', 'USDT', '/assets/currency/usdt.png', 11, 1.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (7, 2, 'USDC', 'USDC', '/assets/currency/usdc.png', 11, 1.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (8, 2, 'BTC', 'BTC', '/assets/currency/btc.png', 11, 38546.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (9, 2, 'ETH', 'ETH', '/assets/currency/eth.png', 11, 2090.00, 10, '', 1709021277, 1709021277);
INSERT INTO `wallet_assets` (`id`, `admin_id`, `name`, `symbol`, `icon`, `type`, `rate`, `status`, `data`, `updated_at`, `created_at`) VALUES (10, 2, 'TRX', 'TRX', '/assets/currency/trx.png', 11, 0.10, 10, '', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for wallet_bill
-- ----------------------------
DROP TABLE IF EXISTS `wallet_bill`;
CREATE TABLE `wallet_bill` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `assets_id` int unsigned NOT NULL COMMENT '资产ID',
  `source_id` int unsigned NOT NULL COMMENT '来源ID',
  `type` smallint NOT NULL DEFAULT '1' COMMENT '类型 ...',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `money` decimal(12,2) NOT NULL COMMENT '金额',
  `balance` decimal(12,2) NOT NULL COMMENT '余额',
  `data` text COMMENT '数据',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包账单表';

-- ----------------------------
-- Records of wallet_bill
-- ----------------------------
BEGIN;
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (1, 2, 1, 0, 0, 71, 'walletBillTypeTeamEarnings', 100.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (2, 2, 1, 0, 0, 11, 'walletBillTypeWithdraw', 100.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (3, 2, 1, 0, 0, 71, 'walletBillTypeTeamEarnings', 100.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (4, 2, 1, 6, 0, 171, 'walletBillTypeTeamAssetsEarnings', 200.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (5, 2, 1, 7, 0, 171, 'walletBillTypeTeamAssetsEarnings', 200.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (6, 2, 2, 0, 0, 71, 'walletBillTypeTeamEarnings', 100.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (7, 2, 3, 0, 0, 71, 'walletBillTypeTeamEarnings', 100.00, 0.00, '', 1709021278);
INSERT INTO `wallet_bill` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `name`, `money`, `balance`, `data`, `created_at`) VALUES (8, 2, 4, 0, 0, 71, 'walletBillTypeTeamEarnings', 100.00, 0.00, '', 1709021278);
COMMIT;

-- ----------------------------
-- Table structure for wallet_order
-- ----------------------------
DROP TABLE IF EXISTS `wallet_order`;
CREATE TABLE `wallet_order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `assets_id` bigint DEFAULT NULL COMMENT '资产ID',
  `source_id` int unsigned NOT NULL COMMENT '来源ID',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型 1充值类型 2资产充值类型 11提现类型 12资产提现类型',
  `order_sn` varchar(60) NOT NULL COMMENT '编号',
  `money` decimal(12,2) NOT NULL COMMENT '金额',
  `fee` decimal(5,2) NOT NULL COMMENT '手续费',
  `voucher` varchar(250) NOT NULL COMMENT '支付凭证',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态  -2删除 -1禁用 10审核 20完成',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_wallet_order_order_sn` (`order_sn`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包订单表';

-- ----------------------------
-- Records of wallet_order
-- ----------------------------
BEGIN;
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 2, 1, 0, 100, 1, 'E202402271607584222738960', 100.00, 0.00, '', 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 2, 1, 0, 100, 1, 'E202402271607583701226355', 100.00, 0.00, '', -1, '拒绝提现, 您还未达到标准提现', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (3, 2, 1, 0, 102, 1, 'E202402271607581617378751', 100.00, 0.00, '', 20, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (4, 2, 1, 0, 102, 1, 'E202402271607583001081213', 100.00, 0.00, '', 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (5, 2, 1, 0, 102, 1, 'E202402271607581380259463', 100.00, 0.00, '', 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (6, 2, 1, 6, 104, 101, 'E202402271607583388932485', 200.00, 0.00, '', 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (7, 2, 1, 6, 104, 101, 'E202402271607582445118234', 200.00, 0.00, '', -1, '拒绝提现, 您还未达到标准提现', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (8, 2, 1, 6, 106, 101, 'E202402271607583304411653', 200.00, 0.00, '', 20, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (9, 2, 1, 7, 106, 101, 'E20240227160758859641414', 200.00, 0.00, '', 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_order` (`id`, `admin_id`, `user_id`, `assets_id`, `source_id`, `type`, `order_sn`, `money`, `fee`, `voucher`, `status`, `data`, `updated_at`, `created_at`) VALUES (10, 2, 1, 7, 106, 101, 'E202402271607583358067421', 200.00, 0.00, '', 10, '', 1709021278, 1709021278);
COMMIT;

-- ----------------------------
-- Table structure for wallet_payment
-- ----------------------------
DROP TABLE IF EXISTS `wallet_payment`;
CREATE TABLE `wallet_payment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `assets_id` int unsigned NOT NULL COMMENT 'default:0',
  `name` varchar(60) NOT NULL COMMENT '名称',
  `icon` varchar(60) NOT NULL COMMENT '图标',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型 1银行卡类型 11数字货币类型 21第三方支付',
  `mode` tinyint NOT NULL DEFAULT '1' COMMENT '模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式',
  `level` tinyint NOT NULL DEFAULT '1' COMMENT '等级',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `desc` text COMMENT '详情',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包支付表';

-- ----------------------------
-- Records of wallet_payment
-- ----------------------------
BEGIN;
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (1, 0, 0, 'bankFundsWithdraw', '/assets/icon/bank.png', 1, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (2, 0, 0, 'bankAssetsWithdraw', '/assets/icon/bank.png', 1, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (3, 0, 0, 'digitalFundsWithdraw', '/assets/currency/usdt.png', 11, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (4, 0, 0, 'digitalAssetsWithdraw', '/assets/currency/usdt.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (5, 0, 0, 'bankFundsDeposit', '/assets/icon/bank.png', 1, 1, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (6, 0, 0, 'bankAssetsDeposit', '/assets/icon/bank.png', 1, 2, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (7, 0, 0, 'digitalFundsDeposit', '/assets/currency/usdt.png', 11, 1, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (8, 0, 0, 'digitalAssetsDeposit', '/assets/currency/usdt.png', 11, 2, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (100, 1, 0, 'DebitCard', '/assets/icon/bank.png', 1, 1, 1, 10, '{\"name\":\"建设银行\",\"realName\":\"隔壁老王\",\"number\":\"888866665555\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (101, 1, 0, 'BankOfAmerica', '/assets/icon/bank-of-america-logo.png', 1, 1, 1, 10, '{\"name\":\"建设银行\",\"realName\":\"隔壁老王\",\"number\":\"888866665555\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (102, 1, 0, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 1, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDT\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (103, 1, 0, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 1, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDC\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (104, 1, 1, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDT\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (105, 1, 2, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDC\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (106, 1, 3, 'ERC20-BTC', '/assets/currency/btc.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"BTC\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (107, 1, 4, 'ERC20-ETH', '/assets/currency/eth.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"ETH\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (108, 1, 5, 'ERC20-TRX', '/assets/currency/trx.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"TRX\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (109, 1, 0, 'DebitCard', '/assets/icon/bank.png', 1, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (110, 1, 0, 'BankOfAmerica', '/assets/icon/bank-of-america-logo.png', 1, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (111, 1, 0, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (112, 1, 0, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (113, 1, 1, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (114, 1, 2, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (115, 1, 3, 'ERC20-BTC', '/assets/currency/btc.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (116, 1, 4, 'ERC20-ETH', '/assets/currency/eth.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (117, 1, 5, 'ERC20-TRX', '/assets/currency/trx.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (118, 2, 0, 'DebitCard', '/assets/icon/bank.png', 1, 1, 1, 10, '{\"name\":\"建设银行\",\"realName\":\"隔壁老王\",\"number\":\"888866665555\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (119, 2, 0, 'BankOfAmerica', '/assets/icon/bank-of-america-logo.png', 1, 1, 1, 10, '{\"name\":\"建设银行\",\"realName\":\"隔壁老王\",\"number\":\"888866665555\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (120, 2, 0, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 1, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDT\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (121, 2, 0, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 1, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDC\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (122, 2, 6, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDT\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (123, 2, 7, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"USDC\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (124, 2, 8, 'ERC20-BTC', '/assets/currency/btc.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"BTC\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (125, 2, 9, 'ERC20-ETH', '/assets/currency/eth.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"ETH\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (126, 2, 10, 'ERC20-TRX', '/assets/currency/trx.png', 11, 2, 1, 10, '{\"name\":\"ETH\",\"realName\":\"TRX\",\"number\":\"0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c\",\"code\":\"xxx666\"}', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (127, 2, 0, 'DebitCard', '/assets/icon/bank.png', 1, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (128, 2, 0, 'BankOfAmerica', '/assets/icon/bank-of-america-logo.png', 1, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (129, 2, 0, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (130, 2, 0, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 11, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (131, 2, 6, 'ERC20-USDT', '/assets/currency/usdt.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (132, 2, 7, 'ERC20-USDC', '/assets/currency/usdc.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (133, 2, 8, 'ERC20-BTC', '/assets/currency/btc.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (134, 2, 9, 'ERC20-ETH', '/assets/currency/eth.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
INSERT INTO `wallet_payment` (`id`, `admin_id`, `assets_id`, `name`, `icon`, `type`, `mode`, `level`, `status`, `data`, `desc`, `updated_at`, `created_at`) VALUES (135, 2, 10, 'ERC20-TRX', '/assets/currency/trx.png', 11, 12, 1, 10, '', '', 1709021277, 1709021277);
COMMIT;

-- ----------------------------
-- Table structure for wallet_user_account
-- ----------------------------
DROP TABLE IF EXISTS `wallet_user_account`;
CREATE TABLE `wallet_user_account` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `payment_id` int unsigned NOT NULL COMMENT '支付ID',
  `name` varchar(50) NOT NULL COMMENT '银行名称｜Token',
  `real_name` varchar(50) NOT NULL COMMENT '真实姓名',
  `number` varchar(50) NOT NULL COMMENT '卡号|地址',
  `code` varchar(255) NOT NULL COMMENT '银行代码',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包卡片表';

-- ----------------------------
-- Records of wallet_user_account
-- ----------------------------
BEGIN;
INSERT INTO `wallet_user_account` (`id`, `admin_id`, `user_id`, `payment_id`, `name`, `real_name`, `number`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 2, 1, 36, '建设银行', '隔壁张三', '88888888', '建设银行分行', 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_user_account` (`id`, `admin_id`, `user_id`, `payment_id`, `name`, `real_name`, `number`, `code`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 2, 1, 40, '', '', '0x888skklxlxlckjlssdfsdfdsfd', '', 10, '', 1709021278, 1709021278);
COMMIT;

-- ----------------------------
-- Table structure for wallet_user_assets
-- ----------------------------
DROP TABLE IF EXISTS `wallet_user_assets`;
CREATE TABLE `wallet_user_assets` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL COMMENT '管理ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `wallet_assets_id` int unsigned NOT NULL COMMENT '钱包资产ID',
  `money` decimal(12,2) NOT NULL COMMENT '金额',
  `status` smallint NOT NULL DEFAULT '10' COMMENT '状态 -2删除 -1禁用 10开启',
  `data` text COMMENT '数据',
  `updated_at` int unsigned NOT NULL COMMENT '更新时间',
  `created_at` int unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包用户资产表';

-- ----------------------------
-- Records of wallet_user_assets
-- ----------------------------
BEGIN;
INSERT INTO `wallet_user_assets` (`id`, `admin_id`, `user_id`, `wallet_assets_id`, `money`, `status`, `data`, `updated_at`, `created_at`) VALUES (1, 2, 1, 6, 100.00, 10, '', 1709021278, 1709021278);
INSERT INTO `wallet_user_assets` (`id`, `admin_id`, `user_id`, `wallet_assets_id`, `money`, `status`, `data`, `updated_at`, `created_at`) VALUES (2, 2, 1, 7, 200.00, 10, '', 1709021278, 1709021278);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
