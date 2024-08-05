import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Layouts',
    component: () => import('layouts/main.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('pages/index.vue'),
        meta: { requireAuth: true, keepAlive: true },
      },
      {
        path: '/views',
        name: 'Views',
        component: () => import('pages/views.vue'),
        meta: { requireAuth: true, keepAlive: true },
      },
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/login',
    name: 'Login',
    component: () => import('pages/login.vue'),
  },
  {
    path: '/:catchAll(.*)*',
    name: 'NotFound',
    component: () => import('src/pages/404.vue'),
  },
];

export default routes;
