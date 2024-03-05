import { defineStore } from 'pinia';
import {LocalStorage} from 'quasar';

export const UserInfoKey = '_UserInfo'

// 初始化数据
export const UserStore = defineStore('user', {
  state: () => {
    return {
      userInfo: {} as any,
    }
  },
  getters: {},
  actions: {
    // 更新初始化数据
    updateUserInfo(info: any) {
      LocalStorage.set(UserInfoKey, JSON.stringify(info))
      this.userInfo = info
    },
  },
});
