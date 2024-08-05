import { date } from 'quasar'

// 图片处理方法
export const imageSrc = (url: any) => {
  if (url === '') {
    url = '/images/logo.png';
  }

  if (url && url.indexOf('http') > -1) {
    return url;
  }

  const baseURL = new URL(<string>process.env.baseURL);
  return baseURL.origin + url;
};

// 图标方法
export const iconSrc = (icon: any) => {
  return icon.indexOf('/') > -1 ? 'img:' + icon : icon;
};

// 日期格式化
export const formatDate = (time: string) => {
  return date.formatDate(Number(time) * 1000, 'YYYY-MM-DD HH:mm:ss')
};
