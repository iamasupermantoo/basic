// 格式化金额
export const formatAmount = (amount: number) => {
  return new Intl.NumberFormat(undefined, {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(amount);
};

// 生成随机数字
export const randomInten = (
  rawData: Array<string>,
  min = 0 as number,
  max = 0 as number
) => {
  //  如果原数据没有， 那么直接返回随机数
  if (rawData.length === 0) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  return Math.floor(Math.random() * (rawData.length + 1));
};
