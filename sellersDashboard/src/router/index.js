// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'

import LandingPage        from '@/pages/LandingPage.vue'
import RegisterPage       from '@/pages/RegisterPage.vue'
import LoginPage          from '@/pages/LoginPage.vue'
import ShopSelectionPage  from '@/pages/ShopSelectionPage.vue'
import DashboardLayout    from '@/layouts/AppLayout.vue'
import DashboardPage      from '@/pages/dashboard/DashboardPage.vue'
import NotFound           from '@/pages/NotFound.vue'

const routes = [
  {
    path: '/',
    name: 'Landing',
    component: LandingPage,
    meta: { public: true }
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterPage,
    meta: { public: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
    meta: { public: true }
  },
  {
    path: '/shops',
    name: 'ShopSelection',
    component: ShopSelectionPage,
    // not marked public ⇒ requires auth
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { requiresShop: true },
    children: [
      { path: '', name: 'Dashboard', component: DashboardPage },
      { path: 'products', name: 'Products', component: () => import('@/pages/products/ProductsPage.vue') },
      { path: 'products/:productId', name: 'ProductDetail', component: () => import('@/pages/products/ProductDetailPage.vue'), props: true },
      { path: 'products/:productId/edit', name: 'EditProduct', component: () => import('@/pages/products/EditProductPage.vue'), props: true },
      { path: 'collections', name: 'Collections', component: () => import('@/pages/collections/CollectionsPage.vue') },
      { path: 'collections/:collectionId', name: 'CollectionDetail', component: () => import('@/pages/collections/CollectionDetailPage.vue'), props: true },
      { path: 'customers', name: 'Customers', component: () => import('@/pages/customers/CustomersPage.vue') },
      { path: 'customers/:customerId', name: 'CustomerDetail', component: () => import('@/pages/customers/CustomerDetailPage.vue'), props: true },
      { path: 'orders', name: 'Orders', component: () => import('@/pages/orders/OrdersPage.vue') },
      { path: 'discounts', name: 'Discounts', component: () => import('@/pages/discounts/DiscountsPage.vue') },
      { path: 'discounts/:discountId', name: 'DiscountDetail', component: () => import('@/pages/discounts/DiscountDetailPage.vue'), props: true },
      { path: 'analytics', name: 'Analytics', component: () => import('@/pages/analytics/AnalyticsPage.vue') },
      { path: 'settings', name: 'Settings', component: () => import('@/pages/settings/SettingsPage.vue') },
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  const auth  = useAuthStore()
  const shops = useShopStore()

  // 0) Try to bootstrap the auth user once, if we haven't yet
  if (!auth.user && !to.meta.public) {
    try {
      await auth.fetchMe()
    } catch {
      // leave auth.user null if fetch fails
    }
  }

  const isPublic      = !!to.meta.public
  const isLoggedIn    = auth.isAuthenticated
  const needsShop     = !!to.meta.requiresShop

  // 1) Unauthenticated users can ONLY see public pages
  if (!isLoggedIn && !isPublic) {
    return next({ name: 'Landing' })
  }

  // 2) Authenticated users should NOT see public pages
  if (isLoggedIn && isPublic) {
    const dest = shops.activeShop ? 'Dashboard' : 'ShopSelection'
    return next({ name: dest })
  }

  // 3) If this route needs an active shop, enforce it
  if (needsShop && !shops.activeShop) {
    return next({ name: 'ShopSelection' })
  }

  // 4) All good — proceed
  next()
})

export default router
