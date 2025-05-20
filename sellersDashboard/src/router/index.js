import { createRouter, createWebHistory } from 'vue-router';
import store from '../store';

// Pages
import LandingPage from '../pages/LandingPage.vue';
import RegisterPage from '../pages/RegisterPage.vue';
import LoginPage from '../pages/LoginPage.vue';
import ShopSelectionPage from '../pages/ShopSelectionPage.vue';
import NotFound from '../pages/NotFound.vue';

// Dashboard
import DashboardLayout from '../layouts/DashboardLayout.vue';

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
    path: '/app',
    component: DashboardLayout,
    meta: { requiresAuth: true, requiresShop: true },
    children: [
      {
        path: '',
        redirect: { name: 'Products' }
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('../pages/products/ProductsPage.vue')
      },
      // Future routes can be uncommented when needed
      // {
      //   path: 'orders',
      //   name: 'Orders',
      //   component: () => import('../pages/OrdersPage.vue')
      // },
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// Global auth/shop guard
router.beforeEach(async (to, from, next) => {
  const isAuthenticated = store.getters['auth/isAuthenticated'];
  const activeShop = store.getters['shops/activeShop'];

  // If route requires auth and user is not logged in
  if (to.meta.requiresAuth && !isAuthenticated) {
    return next({ name: 'Login' });
  }

  // If already logged in, prevent navigating to auth pages
  if (!to.meta.requiresAuth && isAuthenticated) {
    return next(activeShop ? { name: 'Products' } : { name: 'ShopSelection' });
  }

  // If route requires a shop and none is selected
  if (to.meta.requiresShop && !activeShop) {
    return next({ name: 'ShopSelection' });
  }

  next();
});

export default router;
