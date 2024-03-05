import { api } from 'src/boot/axios';

// 新的验证码
export const captchaCreateAPI = () => {
  return api.get('/captcha/create', { showLoading: false } as any);
};

// 管理登录
export const adminLoginAPI = (params: any) => {
  return api.post('/login', params);
};

// 获取管理路由菜单
export const adminInitAPI = () => {
  return api.get('/init', { showLoading: false } as any);
};

// 管理用户信息
export const adminInfoAPI = () => {
  return api.get('/info', { showLoading: false } as any);
};

// 管理提示音
export const adminAudioAPI = () => {
  return api.get('/audio', { showLoading: false } as any);
};

// 控制台信息
export const dashboardAPI = () => {
  return api.get('/index');
};
