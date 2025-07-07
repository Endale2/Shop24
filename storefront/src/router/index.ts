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
    path: '/shops/:shopSlug',
    component: StoreLayout,
    props: true,
    children: [
      { path: '', component: HomePage },
      { path: 'products', component: ProductsPage },
      { path: 'products/:handle', component: ProductDetail, props: true },
      { path: 'collections', component: CollectionsPage },
      {
        path: 'collections/:handle',
        component: CollectionDetail,
        props: true,
      },
      { 
        path: 'account', 
        name: 'Account', 
        component: AccountPage,
        beforeEnter: (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
          const authStore = useAuthStore()
          if (!authStore.user) {
            next(`/shops/${to.params.shopSlug}/login`)
          } else {
            next()
          }
        }
      },
      { 
        path: 'orders', 
        name: 'MyOrders', 
        component: MyOrdersPage,
        beforeEnter: (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
          const authStore = useAuthStore()
          if (!authStore.user) {
            next(`/shops/${to.params.shopSlug}/login`)
          } else {
            next()
          }
        }
      },
      { path: 'login', name: 'Login', component: LoginPage },
      { path: 'register', name: 'Register', component: RegisterPage },
      { path: 'cart', name: 'Cart', component: CartPage },
      { path: 'order-confirmation/:orderId', name: 'OrderConfirmation', component: OrderConfirmation },
      { path: 'wishlist', name: 'Wishlist', component: WishlistPage },
    ],
  },
  // Shop selection page
  {
    path: '/',
    component: ShopSelection
  },
 
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
