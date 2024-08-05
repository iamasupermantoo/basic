// 递归结构体合并
export const assignRecursion = (target = {} as any, source = {} as any) => {
  const obj = target;
  if (typeof target != 'object' || typeof source != 'object') {
    return source;
  }

  for (const key in source) {
    if (target.hasOwnProperty(key)) {
      obj[key] = assignRecursion(target[key], source[key]);
    } else {
      obj[key] = source[key];
    }
  }
  return obj;
};
