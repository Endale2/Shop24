import { createRouter, createWebHistory } from 'vue-router'
import StoreLayout from '../layouts/StoreLayout.vue'
import HomePage from '../pages/HomePage.vue'
import ProductsPage from '../pages/ProductsPage.vue'
import ProductDetail from '../pages/ProductDetail.vue'
import CollectionsPage from '../pages/CollectionsPage.vue'
import CollectionDetail from '../pages/CollectionDetail.vue'
import AccountPage from '../pages/AccountPage.vue'
import ShopSelection from '../pages/ShopSelection.vue'
import LoginPage from '../pages/LoginPage.vue'
import RegisterPage from '../pages/RegisterPage.vue'
import CartPage from '../pages/CartPage.vue'
import OrderConfirmation from '../pages/OrderConfirmation.vue'
import WishlistPage from '../pages/WishlistPage.vue'
import MyOrdersPage from '../pages/MyOrdersPage.vue'
import { useAuthStore } from '../stores/auth'
import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'

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
      { path: 'account', name: 'Account', component: AccountPage },
      { path: 'login', name: 'Login', component: LoginPage },
      { path: 'register', name: 'Register', component: RegisterPage },
      { path: 'cart', name: 'Cart', component: CartPage },
      { path: 'order-confirmation/:orderId', name: 'OrderConfirmation', component: OrderConfirmation },
      { path: 'wishlist', name: 'Wishlist', component: WishlistPage },
      { path: 'orders', name: 'MyOrders', component: MyOrdersPage },
    ],
  },
  // Shop selection page (for root domain only)
  {
    path: '/select-shop',
    component: ShopSelection
  },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
