import { createRouter, createWebHistory } from 'vue-router'
import StoreLayout from '../layouts/StoreLayout.vue'
import HomePage from '../pages/HomePage.vue'
import ProductsPage from '../pages/ProductsPage.vue'
import ProductDetail from '../pages/ProductDetail.vue'
import CollectionsPage from '../pages/CollectionsPage.vue'
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
    component: StoreLayout,
    children: [
      { path: '', name: 'Home', component: HomePage },
      { path: 'products', name: 'Products', component: ProductsPage },
      { path: 'products/:handle', name: 'ProductDetail', component: ProductDetail },
      { path: 'collections', name: 'Collections', component: CollectionsPage },
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
  next()
})

export default router
