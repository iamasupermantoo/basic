import { api } from 'src/boot/axios';

// 创建验证码
export const captchaAPI = () => {
  return api.get('/captcha/create');
};

// 初始化数据
export const initAPI = (params: any) => {
  return api.get(`/init?domain=${params.domain}&lang=${params.lang}`);
};

// 获取footer数据
export const footerInfoAPI = () => {
  return api.post('/footer');
};

// 获取文章详情
export const articleInfoAPI = (params: any) => {
  return api.post('/article/info', params);
};

// 获取下载信息
export const downloadInfoAPI = () => {
  return api.post('/download');
};

// 获取帮助列表
export const helpersInfoAPI = () => {
  return api.post('/helpers');
};

// 访问请求
export const accessAPI = (params: any) => {
  return api.post('/access', params)
}
