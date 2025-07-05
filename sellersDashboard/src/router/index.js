// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'

import LandingPage        from '@/pages/LandingPage.vue'
import RegisterPage       from '@/pages/RegisterPage.vue'
import LoginPage          from '@/pages/LoginPage.vue'
import ProfileCompletion  from '@/pages/ProfileCompletion.vue'
import ShopSelectionPage  from '@/pages/ShopSelectionPage.vue'
import DashboardLayout    from '@/layouts/AppLayout.vue'
import DashboardPage      from '@/pages/dashboard/DashboardPage.vue'
import NotFound           from '@/pages/NotFound.vue'
import ProfilePage from '@/pages/ProfilePage.vue'
import AddProduct from '@/pages/products/AddProduct.vue'

const routes = [
  { path: '/',              name: 'Landing',       component: LandingPage,      meta: { public: true } },
  { path: '/register',      name: 'Register',      component: RegisterPage,     meta: { public: true } },
  { path: '/login',         name: 'Login',         component: LoginPage,        meta: { public: true } },
  { path: '/complete-profile',
    name: 'CompleteProfile',
    component: ProfileCompletion,
    meta: { requiresAuth: true }
  },
  {
    path: '/shops',
    name: 'ShopSelection',
    component: ShopSelectionPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/profile',
    name: 'ProfilePage',
    component: ProfilePage,
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { requiresAuth: true, requiresShop: true },
    children: [
      { path: '',                     name: 'Dashboard',      component: DashboardPage },
      { path: 'products',             name: 'Products',       component: () => import('@/pages/products/ProductsPage.vue') },
      { path: 'products/:productId',  name: 'ProductDetail',  component: () => import('@/pages/products/ProductDetailPage.vue'), props: true },
      { path: 'products/:productId/edit', name: 'EditProduct', component: () => import('@/pages/products/EditProductPage.vue'), props: true },
      { path: 'products/add',         name: 'AddProduct',      component: () => import('@/pages/products/AddProduct.vue'), props: true  },
      { path: 'collections',          name: 'Collections',    component: () => import('@/pages/collections/CollectionsPage.vue') },
      { path: 'collections/:collectionId', name: 'CollectionDetail', component: () => import('@/pages/collections/CollectionDetailPage.vue'), props: true },
      { path: 'customers',            name: 'Customers',      component: () => import('@/pages/customers/CustomersPage.vue') },
      { path: 'customers/:customerId',name: 'CustomerDetail', component: () => import('@/pages/customers/CustomerDetailPage.vue'), props: true },
      { path: 'orders',               name: 'Orders',         component: () => import('@/pages/orders/OrdersPage.vue') },
      { path: 'orders/:orderId',      name: 'OrderDetail',    component: () => import('@/pages/orders/OrderDetailPage.vue'), props: true },
      { path: 'discounts',            name: 'Discounts',      component: () => import('@/pages/discounts/DiscountsPage.vue') },
      { path: 'discounts/:discountId',name: 'DiscountDetail', component: () => import('@/pages/discounts/DiscountDetailPage.vue'), props: true },
      { path: 'analytics',            name: 'Analytics',      component: () => import('@/pages/analytics/AnalyticsPage.vue') },
      { path: 'settings',             name: 'Settings',       component: () => import('@/pages/settings/SettingsPage.vue') },
    ]
  },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  const auth  = useAuthStore()
  const shops = useShopStore()

  // 1) Ensure we have the user loaded if they're authenticated
  if (!auth.user && auth.$state.user === null && !to.meta.public) {
    try { await auth.fetchMe() }
    catch { /* ignore—will be redirected below if not authed */ }
  }

  const isPublic      = Boolean(to.meta.public)
  const requiresAuth  = Boolean(to.meta.requiresAuth || to.meta.requiresShop)
  const isLoggedIn    = auth.isAuthenticated
  const profileReady  = Boolean(auth.user?.first_name && auth.user?.last_name)
  const hasShop       = Boolean(shops.active)

  // 2) Unauthenticated => only public
  if (requiresAuth && !isLoggedIn) {
    return next({ name: 'Landing' })
  }

  // 3) Logged in users should not hit public routes
  if (isLoggedIn && isPublic) {
    // incomplete profile?
    if (!profileReady) {
      return next({ name: 'CompleteProfile' })
    }
    // complete profile but no shop yet?
    if (!hasShop) {
      return next({ name: 'ShopSelection' })
    }
    // fully set up => dashboard
    return next({ name: 'Dashboard' })
  }

  // 4) If logged in but profile incomplete, lock them to CompleteProfile
  if (isLoggedIn && !profileReady && to.name !== 'CompleteProfile') {
    return next({ name: 'CompleteProfile' })
  }

  // If profile just completed, redirect off of the completion page
  if (isLoggedIn && profileReady && to.name === 'CompleteProfile') {
    return next({ name: 'ShopSelection' })
  }

  // 5) If profile ready but no shop, and they're trying dashboard routes, force ShopSelection
  if (isLoggedIn && profileReady && !hasShop && to.meta.requiresShop) {
    return next({ name: 'ShopSelection' })
  }

  // otherwise allow
  next()
})

export default router