import { useSSRContext } from 'vue';
import { Cookies } from 'quasar';

//  获取cookies
const cookies = process.env.SERVER
  ? Cookies.parseSSR(useSSRContext())
  : Cookies;

export default cookies;
