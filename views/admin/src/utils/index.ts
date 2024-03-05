// 图片处理方法
export const imageSrc = (url: string) => {
  if (url === '') {
    url = '/images/logo.png';
  }
  if (url.indexOf('http') > -1) {
    return url;
  }

  const baseURL = new URL(<string>process.env.baseURL);
  return baseURL.origin + url;
};
