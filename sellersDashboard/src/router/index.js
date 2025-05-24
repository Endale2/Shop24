// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'

import LandingPage from '@/pages/LandingPage.vue'
import RegisterPage from '@/pages/RegisterPage.vue'
import LoginPage from '@/pages/LoginPage.vue'
import ShopSelectionPage from '@/pages/ShopSelectionPage.vue'
import DashboardLayout from '@/layouts/AppLayout.vue'
import DashboardPage from '@/pages/dashboard/DashboardPage.vue'
import NotFound from '@/pages/NotFound.vue'
import StorefrontPage from '@/pages/storefront/StorefrontPage.vue'
const routes = [
   // If a subdomain is set (we detect in main.js),
  // `/` will render the storefront:
   {
    path: '/',
    name: 'Storefront',
    component: StorefrontPage,
    meta: { requiresAuth: false }
  },

  
  {
    path: '/',
    name: 'Landing',
    component: LandingPage,
    meta: { requiresAuth: false },
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterPage,
    meta: { requiresAuth: false },
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
    meta: { requiresAuth: false },
  },
  {
    path: '/shops',
    name: 'ShopSelection',
    component: ShopSelectionPage,
    meta: { requiresAuth: true },
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { requiresAuth: true, requiresShop: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: DashboardPage,
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('@/pages/products/ProductsPage.vue'),
      },
      {
        path: 'products/:productId',
        name: 'ProductDetail',
        component: () => import('@/pages/products/ProductDetailPage.vue'),
        props: true,
      },
      {
        path: 'customers',
        name: 'Customers',
        component: () => import('@/pages/customers/CustomersPage.vue'),
      },
      {
        path: 'customers/:customerId',
        name: 'CustomerDetail',
        component: () => import('@/pages/customers/CustomerDetailPage.vue'),
        props: true,
      },
      {
        path: 'orders',
        name: 'Orders',
        component: () => import('@/pages/orders/OrdersPage.vue'),
      },
      {
        path: 'discounts',
        name: 'Discounts',
        component: () => import('@/pages/discounts/DiscountsPage.vue'),
      },
      {
        path: 'analytics',
        name: 'Analytics',
        component: () => import('@/pages/analytics/AnalyticsPage.vue'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  const auth = useAuthStore()
  const shops = useShopStore()

  // 1) If route requires auth, ensure we have a valid user
  if (to.meta.requiresAuth) {
    if (!auth.user) {
      try {
        await auth.fetchMe()
      } catch {
        // token invalid or expired — force logout and redirect
        await auth.logout()
        return next({ name: 'Login' })
      }
    }
    if (!auth.isAuthenticated) {
      return next({ name: 'Login' })
    }
  }

  // 2) Prevent logged-in users from visiting public pages
  if (!to.meta.requiresAuth && auth.isAuthenticated) {
    const dest = shops.activeShop ? 'Products' : 'ShopSelection'
    return next({ name: dest })
  }

  // 3) Protect dashboard routes until a shop is selected
  if (to.meta.requiresShop && !shops.activeShop) {
    return next({ name: 'ShopSelection' })
  }

  // 4) All clear
  next()
})

export default router
