// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useShopStore }                from '@/stores/shop'
import HomePage                        from '@/pages/HomePage.vue'
import ProductsPage                    from '@/pages/ProductsPage.vue'
import ProductDetail                   from '@/pages/ProductDetail.vue'
import CollectionsPage                 from '@/pages/CollectionsPage.vue'
import CollectionDetailPage            from '@/pages/CollectionDetailPage.vue'
import NotFoundPage                    from '@/pages/NotFoundPage.vue'

const routes = [
  { path: '/',         name: 'Home',        component: HomePage },
  { path: '/products', name: 'Products',    component: ProductsPage },
  {
    path: '/products/:productSlug',
    name: 'Product',
    component: ProductDetail,
    props: route => ({
      shopSlug:    window.location.host.split('.')[0],
      productSlug: route.params.productSlug
    })
  },

  // ← COLLECTIONS listing:
   {
    path: '/collections',
    name: 'Collections',
    component: CollectionsPage
  },
  {
    path: '/collections/:handle',
    name: 'CollectionDetail',
    component: CollectionDetailPage,
    props: true
  },

  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFoundPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior: (to, from, saved) => saved || { top: 0 }
})

router.beforeEach(async (to, from, next) => {
  const shopSlug = window.location.host.split('.')[0]
  const shop     = useShopStore()

  if (!shop.current || shop.current.slug !== shopSlug) {
    try {
      await shop.fetchShopAndProducts(shopSlug)
    } catch (e) {
      console.error('Error fetching shop data:', e)
    }
  }

  next()
})

export default router
