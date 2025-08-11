import { createRouter, createWebHistory } from 'vue-router'
import DynamicStoreLayout from '../layouts/DynamicStoreLayout.vue'
import DynamicHomePage from '../pages/DynamicHomePage.vue'
import DynamicProductsPage from '../pages/DynamicProductsPage.vue'
import DynamicCollectionsPage from '../pages/DynamicCollectionsPage.vue'
import ProductDetail from '../pages/ProductDetail.vue'
import CollectionDetail from '../pages/CollectionDetail.vue'
import AccountPage from '../pages/AccountPage.vue'
import LoginPage from '../pages/LoginPage.vue'

import CartPage from '../pages/CartPage.vue'
import OrderConfirmation from '../pages/OrderConfirmation.vue'
import WishlistPage from '../pages/WishlistPage.vue'
import MyOrdersPage from '../pages/MyOrdersPage.vue'
import { getCurrentShopSlug } from '../services/shop'

const routes = [
  {
    path: '/',
    component: DynamicStoreLayout,
    children: [
      { path: '', name: 'Home', component: DynamicHomePage },
      { path: 'products', name: 'Products', component: DynamicProductsPage },
      { path: 'products/:handle', name: 'ProductDetail', component: ProductDetail },
      { path: 'collections', name: 'Collections', component: DynamicCollectionsPage },
      { path: 'collections/:handle', name: 'CollectionDetail', component: CollectionDetail },
      { path: 'account', name: 'Account', component: AccountPage, meta: { requiresAuth: true } },
      { path: 'login', name: 'Login', component: LoginPage },
      { path: 'cart', name: 'Cart', component: CartPage, meta: { requiresAuth: true } },
      { path: 'order-confirmation/:orderId', name: 'OrderConfirmation', component: OrderConfirmation, meta: { requiresAuth: true } },
      { path: 'wishlist', name: 'Wishlist', component: WishlistPage, meta: { requiresAuth: true } },
      { path: 'orders', name: 'MyOrders', component: MyOrdersPage, meta: { requiresAuth: true } },
      { path: 'search', name: 'SearchResults', component: () => import('../pages/SearchResults.vue') },
    ],
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
 
// Ensure a valid subdomain shop is present; otherwise redirect to a friendly landing or 404
router.beforeEach((to, _from, next) => {
  const slug = getCurrentShopSlug()
  // If there is no slug at all (no subdomain and no configured default),
  // only allow Home and Login to render a friendly experience
  if (!slug && to.name !== 'Home' && to.name !== 'Login') {
    return next({ name: 'Home' })
  }
  // If route requires auth and user is not authenticated (based on cookie absence), push to login
  if (to.meta?.requiresAuth) {
    const hasAccessCookie = /(?:^|; )access_token=/.test(document.cookie)
    if (!hasAccessCookie) {
      return next({ name: 'Login' })
    }
  }
  next()
})

export default router
