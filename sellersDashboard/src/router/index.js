// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import store from '../store';

// Page components
import LandingPage from '../pages/LandingPage.vue';
import RegisterPage from '../pages/RegisterPage.vue';
import LoginPage from '../pages/LoginPage.vue';
import ShopSelectionPage from '../pages/ShopSelectionPage.vue';
import NotFound from '../pages/NotFound.vue';

// Dashboard layout and its child pages
import DashboardLayout from '../layouts/DashboardLayout.vue';
import Sidebar from '@/Sidebar.vue';

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
        // component: ProductsPage
        component: () => import('../pages/ProductsPage.vue')
      },
    //   {
    //     path: 'orders',
    //     name: 'Orders',
    //     component: () => import('../views/OrdersPage.vue')
    //   },
    //   {
    //     path: 'customers',
    //     name: 'Customers',
    //     component: () => import('../views/CustomersPage.vue')
    //   },
    //   {
    //     path: 'discounts',
    //     name: 'Discounts',
    //     component: () => import('../views/DiscountsPage.vue')
    //   },
    //   {
    //     path: 'analytics',
    //     name: 'Analytics',
    //     component: () => import('../views/AnalyticsPage.vue')
    //   },
    //   {
    //     path: 'settings',
    //     name: 'Settings',
    //     component: () => import('../views/SettingsPage.vue')
    //   }
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

// Global navigation guard
router.beforeEach((to, from, next) => {
  const isAuthenticated = store.getters['auth/isAuthenticated'];
  const activeShop = store.getters['shops/activeShop'];

  // Redirect unauthenticated users away from protected routes
  if (to.meta.requiresAuth && !isAuthenticated) {
    return next({ name: 'Login' });
  }

  // After login/registration, prevent access to auth pages
  if (!to.meta.requiresAuth && isAuthenticated) {
    // If user has an active shop, send to dashboard
    if (activeShop) {
      return next({ name: 'Products' });
    }
    // Otherwise send to shop selection
    return next({ name: 'ShopSelection' });
  }

  // Ensure a shop is selected for routes that require one
  if (to.meta.requiresShop && !activeShop) {
    return next({ name: 'ShopSelection' });
  }

  next();
});

export default router;