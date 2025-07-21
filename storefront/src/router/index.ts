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
import { useCartStore } from '../stores/cart'
import { useWishlistStore } from '../stores/wishlist'
import api from '../services/api'
import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { getCurrentShopSlug, clearPersistedShop, setCurrentShopSlug } from '../services/shop'

const routes = [
  {
    path: '/:shopSlug',
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
  // Shop selection page is now the root
  {
    path: '/',
    component: ShopSelection
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Global navigation guard: enforce shop slug presence in path
router.beforeEach(async (to, from, next) => {
  // Allow /select-shop and root
  if (to.path === '/select-shop' || to.path === '/') return next();
  const shopSlug = getCurrentShopSlug();
  const toShopSlug = to.params.shopSlug as string | undefined;
  if (toShopSlug && toShopSlug !== shopSlug) {
    // User is navigating to a different shop, clear previous shop context
    clearPersistedShop();
    setCurrentShopSlug(toShopSlug);

    // Clear all localStorage
    localStorage.clear();

    // Call backend logout to clear cookies
    try { await api.post('/auth/customer/logout'); } catch {}

    // Reset Pinia stores
    useAuthStore().$reset();
    useCartStore().$reset();
    useWishlistStore().$reset();
  }
  if (!shopSlug && !toShopSlug) {
    return next('/select-shop');
  }
  next();
});

export default router
