import { createRouter, createWebHistory } from 'vue-router';
import LandingLayout from '../layouts/LandingLayout.vue';
import AdminLayout from '../layouts/AdminLayout.vue';
import LandingPage from '../pages/LandingPage.vue';
import LoginPage from '../pages/LoginPage.vue';
import DashboardPage from '../pages/admin/DashboardPage.vue';
import ProductsList from '../pages/admin/products/ProductsList.vue';
import ProductDetail from '../pages/admin/products/ProductDetail.vue';
import CustomersList from '../pages/admin/customers/CustomersList.vue';
import CustomerDetail from '../pages/admin/customers/CustomerDetail.vue';
import StoresList from '../pages/admin/stores/StoresList.vue';
import StoreDetail from '../pages/admin/stores/StoreDetail.vue';
import auth from '../services/auth';

const routes = [
  // Public
  {
    path: '/',
    component: LandingLayout,
    children: [
      { path: '', name: 'Landing', component: LandingPage },
      { path: 'login', name: 'Login', component: LoginPage }
    ]
  },
  // Protected Admin
  {
    path: '/',
    component: AdminLayout,
    meta: { requiresAuth: true },
    children: [
      { path: 'dashboard', name: 'Dashboard', component: DashboardPage },
      { path: 'products', name: 'ProductsList', component: ProductsList },
      { path: 'products/:id', name: 'ProductDetail', component: ProductDetail, props: true },
      { path: 'customers', name: 'CustomersList', component: CustomersList },
      { path: 'customers/:id', name: 'CustomerDetail', component: CustomerDetail, props: true },
      { path: 'stores', name: 'StoresList', component: StoresList },
      { path: 'stores/:id', name: 'StoreDetail', component: StoreDetail, props: true }
    ]
  },
  { path: '/:pathMatch(.*)*', redirect: '/' }
];

const router = createRouter({ history: createWebHistory(), routes });

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth) {
    const loggedIn = await auth.isAuthenticated();
    if (!loggedIn) next({ name: 'Login' });
    else next();
  } else next();
});

export default router;