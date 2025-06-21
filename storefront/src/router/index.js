// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useShopStore } from '@/stores/shop'
import { useAuthStore } from '@/stores/auth'

import HomePage from '@/pages/HomePage.vue'
import ProductsPage from '@/pages/ProductsPage.vue'
import ProductDetail from '@/pages/ProductDetail.vue'
import CollectionsPage from '@/pages/CollectionsPage.vue'
import CollectionDetailPage from '@/pages/CollectionDetailPage.vue'
import CartPage from '@/pages/CartPage.vue'
import AuthPage from '@/pages/AuthPage.vue'
import NotFoundPage from '@/pages/NotFoundPage.vue'

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/auth', name: 'Auth', component: AuthPage },

  // Public product routes
  { path: '/products', name: 'Products', component: ProductsPage },
  { 
    path: '/products/:productSlug', 
    name: 'Product', 
    component: ProductDetail, 
    props: route => ({
      shopSlug: window.location.host.split('.')[0],
      productSlug: route.params.productSlug
    })
  },
  { path: '/collections', name: 'Collections', component: CollectionsPage },
  { 
    path: '/collections/:handle', 
    name: 'CollectionDetail', 
    component: CollectionDetailPage, 
    props: true 
  },

  // Protect only cart (and future private pages)
  { 
    path: '/cart', 
    name: 'Cart', 
    component: CartPage,
    meta: { requiresAuth: true }
  },
  { path: '/account', name: 'Account', component: () => import('@/pages/AccountPage.vue'), meta: { requiresAuth: true } },


  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFoundPage }
]


const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior: (to, from, savedPosition) => savedPosition || { top: 0 }
})

router.beforeEach(async (to, from, next) => {
  const shopSlug = window.location.host.split('.')[0]
  const shop     = useShopStore()
  const auth     = useAuthStore()

  // Always ensure shop data is loaded from subdomain
  if (!shop.current || shop.current.slug !== shopSlug) {
    try {
      await shop.fetchShopAndProducts(shopSlug)
    } catch (e) {
      console.error('Error fetching shop data:', e)
    }
  }

  // If the route has `requiresAuth` and user isn't logged in, send to /auth
  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    return next({ name: 'Auth' })
  }

  // Otherwise proceed normally
  next()
})


export default router
