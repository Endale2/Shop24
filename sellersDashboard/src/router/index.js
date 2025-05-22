// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import store from '@/store';
import LandingPage from '@/pages/LandingPage.vue';
import RegisterPage from '@/pages/RegisterPage.vue';
import LoginPage from '@/pages/LoginPage.vue';
import ShopSelectionPage from '@/pages/ShopSelectionPage.vue';
import DashboardLayout from '@/layouts/AppLayout.vue';
import NotFound from '@/pages/NotFound.vue';
import DashboardPage from '@/pages/dashboard/DashboardPage.vue';
const routes = [
  {
    path: '/',
    name: 'Landing',
    component: LandingPage,
    meta: { requiresAuth: false }
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterPage,
    meta: { requiresAuth: false }
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
    meta: { requiresAuth: false }
  },
  {
    path: '/shops',
    name: 'ShopSelection',
    component: ShopSelectionPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { requiresAuth: true, requiresShop: true },
    children: [
       {
        path: '',               
        name: 'Dashboard',      
        component: DashboardPage
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('../pages/products/ProductsPage.vue')
      },
      {
  path: 'customers',
  name: 'Customers',
  component: () => import('../pages/customers/CustomersPage.vue')
}

    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound
  }
];

const router = createRouter({ history: createWebHistory(), routes });

router.beforeEach(async (to, from, next) => {
  const isAuth = store.getters['auth/isAuthenticated'];
  if (to.meta.requiresAuth) {
    if (!isAuth) return next({ name: 'Login' });
    if (!store.state.auth.user) await store.dispatch('auth/fetchMe');
  }
  if (!to.meta.requiresAuth && isAuth) {
    return next({ name: store.getters['shops/activeShop'] ? 'Products' : 'ShopSelection' });
  }
  if (to.meta.requiresShop && !store.getters['shops/activeShop']) {
    return next({ name: 'ShopSelection' });
  }
  next();
});

export default router;
