// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/pages/HomePage.vue'
import ProductsPage from '@/pages/ProductsPage.vue'
import ProductDetail from '@/pages/ProductDetail.vue'
import NotFoundPage from '@/pages/NotFoundPage.vue'

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/products', name: 'Products', component: ProductsPage },
  {
    // MODIFIED: Path now includes :shopSlug for clarity and to match API structure
    // This allows ProductGrid to link to this route with both slugs as params.
    path: '/products/:productSlug',
    name: 'Product',
    component: ProductDetail,
    props: true // This will pass shopSlug and productSlug as props to the component
  },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFoundPage }
]

export default createRouter({
  history: createWebHistory(),
  routes,
  // Optional: Scroll behavior for navigation
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0 };
    }
  }
})