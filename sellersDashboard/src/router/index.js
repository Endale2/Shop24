// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'
import { useNavigationStore } from '@/store/navigation'

import LandingPage        from '@/pages/LandingPage.vue'
import ProfileCompletion  from '@/pages/ProfileCompletion.vue'
import ShopSelectionPage  from '@/pages/ShopSelectionPage.vue'
import CreateShopPage     from '@/pages/CreateShopPage.vue'
import DashboardLayout    from '@/layouts/AppLayout.vue'
import DashboardPage      from '@/pages/dashboard/DashboardPage.vue'
import NotFound           from '@/pages/NotFound.vue'
import ProfilePage from '@/pages/ProfilePage.vue'
import AddProduct from '@/pages/products/AddProduct.vue'

const routes = [
  { path: '/',              name: 'Landing',       component: LandingPage,      meta: { public: true } },
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
    path: '/create-shop',
    name: 'CreateShop',
    component: CreateShopPage,
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
      { path: 'collections/add',      name: 'AddCollection',  component: () => import('@/pages/collections/AddCollection.vue') },
      { path: 'collections/:collectionId', name: 'CollectionDetail', component: () => import('@/pages/collections/CollectionDetailPage.vue'), props: true },
      { path: 'collections/:collectionId/edit', name: 'EditCollection', component: () => import('@/pages/collections/EditCollectionPage.vue'), props: true },
      { path: 'customers',            name: 'Customers',      component: () => import('@/pages/customers/CustomersPage.vue') },
      { path: 'customers/:customerId',name: 'CustomerDetail', component: () => import('@/pages/customers/CustomerDetailPage.vue'), props: true },
      { path: 'customers/customer-segments',    name: 'CustomerSegments', component: () => import('@/pages/customers/CustomerSegmentsPage.vue') },
      { path: 'orders',               name: 'Orders',         component: () => import('@/pages/orders/OrdersPage.vue') },
      { path: 'orders/:orderId',      name: 'OrderDetail',    component: () => import('@/pages/orders/OrderDetailPage.vue'), props: true },
      { path: 'discounts',            name: 'Discounts',      component: () => import('@/pages/discounts/DiscountsPage.vue') },
      { path: 'discounts/create',     name: 'CreateDiscount',   component: () => import('@/pages/discounts/CreateDiscountPage.vue') },
      { path: 'discounts/:discountId',name: 'DiscountDetail', component: () => import('@/pages/discounts/DiscountDetailPage.vue'), props: true },
      { path: 'discounts/:discountId/edit', name: 'EditDiscount', component: () => import('@/pages/discounts/EditDiscountPage.vue'), props: true },
      { path: 'analytics',            name: 'Analytics',      component: () => import('@/pages/analytics/AnalyticsPage.vue') },
      { path: 'customization',        name: 'Customization',  component: () => import('@/pages/customization/CustomizationPage.vue') },
      { path: 'customization/theme',  name: 'ThemeCustomization', component: () => import('@/pages/customization/ThemePage.vue') },
      { path: 'customization/layout', name: 'LayoutCustomization', component: () => import('@/pages/customization/LayoutPage.vue') },
      { path: 'customization/typography', name: 'TypographyCustomization', component: () => import('@/pages/customization/TypographyPage.vue') },
      { path: 'customization/custom-css', name: 'CustomCSSCustomization', component: () => import('@/pages/customization/CustomCSSPage.vue') },
      { path: 'customization/mobile', name: 'MobileCustomization', component: () => import('@/pages/customization/MobilePage.vue') },
      { path: 'customization/seo',    name: 'SEOCustomization', component: () => import('@/pages/customization/SEOPage.vue') },
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
  const auth = useAuthStore()
  const shops = useShopStore()

  console.log(`[Router] Navigating from ${from.path} to ${to.path}`)

  // Skip auth checks for public routes unless we're already authenticated
  const isPublic = Boolean(to.meta.public)
  const requiresAuth = Boolean(to.meta.requiresAuth || to.meta.requiresShop)

  // 1) For non-public routes, ensure we have valid auth state
  if (!isPublic && requiresAuth) {
    // If we don't have a user or session is loading, try to fetch
    if (!auth.user && !auth.sessionLoading) {
      try {
        console.log('[Router] No user found, attempting to fetch...')
        await auth.fetchMe()
      } catch (error) {
        console.log('[Router] Auth fetch failed, redirecting to landing:', error.message)
        return next({ name: 'Landing' })
      }
    }

    // If still no user after fetch attempt, redirect to landing
    if (!auth.user) {
      console.log('[Router] No authenticated user, redirecting to landing')
      return next({ name: 'Landing' })
    }
  }

  // Get current state
  const isLoggedIn = auth.isAuthenticated
  const profileReady = Boolean(auth.user?.first_name && auth.user?.last_name)
  const hasShop = Boolean(shops.active)

  console.log(`[Router] Auth state - logged in: ${isLoggedIn}, profile ready: ${profileReady}, has shop: ${hasShop}`)

  // 2) Redirect authenticated users away from public routes
  if (isLoggedIn && isPublic) {
    if (!profileReady) {
      console.log('[Router] Authenticated but profile incomplete, redirecting to profile completion')
      return next({ name: 'CompleteProfile' })
    }
    if (!hasShop) {
      console.log('[Router] Profile complete but no shop, redirecting to shop selection')
      return next({ name: 'ShopSelection' })
    }
    console.log('[Router] Fully authenticated, redirecting to dashboard')
    return next({ name: 'Dashboard' })
  }

  // 3) Profile completion flow
  if (isLoggedIn && !profileReady && to.name !== 'CompleteProfile') {
    console.log('[Router] Profile incomplete, redirecting to profile completion')
    return next({ name: 'CompleteProfile' })
  }

  // 4) Redirect away from profile completion if already complete
  if (isLoggedIn && profileReady && to.name === 'CompleteProfile') {
    console.log('[Router] Profile already complete, redirecting to shop selection')
    return next({ name: 'ShopSelection' })
  }

  // 5) Shop selection flow
  if (isLoggedIn && profileReady && !hasShop && to.meta.requiresShop) {
    console.log('[Router] No shop selected, redirecting to shop selection')
    return next({ name: 'ShopSelection' })
  }

  // 6) All checks passed, allow navigation
  console.log('[Router] Navigation allowed')
  next()
})

// After navigation is complete, update navigation store
router.afterEach((to) => {
  const navigationStore = useNavigationStore()
  navigationStore.initializeFromRoute(to.path)
})

export default router