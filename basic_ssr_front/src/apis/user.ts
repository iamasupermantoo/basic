import { api } from 'src/boot/axios';

// 用户登录
export const userLoginAPI = (params: any) => {
  return api.post('/login', params);
};

// 用户注册
export const userRegisterAPI = (params: any) => {
  return api.post('/register', params);
};

// 获取用户信息
export const userInfoAPI = () => {
  return api.post('/user/info');
};

// 用户团队成员
export const teamIndexAPI = (params: any) => {
  return api.post('/user/team/index', params);
};

// 用户团队详情数据
export const teamDetailsAPI = (params: any) => {
  return api.post('/user/team/details', params);
};

// 用户更改信息
export const updateInfoAPI = (params: any) => {
  return api.post('/user/update', params);
};

// 用户更改密码
export const updatePasswordAPI = (params: any) => {
  return api.post('/user/update/password', params);
};

// 会员等级信息
export const levelIndexAPI = () => {
  return api.post('/user/level/index');
};

// 用户购买等级
export const levelCreateAPI = (params: any) => {
  return api.post('/user/level/create', params);
};

// 用户实名认证
export const realAuthCreateAPI = (params: any) => {
  return api.post('/user/auth/create', params);
};

// 用户获取实名
export const realAuthInfoAPI = () => {
  return api.post('/user/auth/info');
};

// 获取用户邀请码
export const inviteInfoAPI = () => {
  return api.post('/user/invite/info');
};



