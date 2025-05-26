import { createRouter, createWebHistory } from 'vue-router'
import { useShopStore }         from '@/stores/shop'
import HomePage                 from '@/pages/HomePage.vue'
import ProductsPage             from '@/pages/ProductsPage.vue'
import ProductDetail            from '@/pages/ProductDetail.vue'
import NotFoundPage             from '@/pages/NotFoundPage.vue'

const routes = [
  { path: '/',                     name: 'Home',     component: HomePage },
  { path: '/products',             name: 'Products', component: ProductsPage },
  {
  path: '/products/:productSlug',
  name: 'ProductDetail',              // ← Name must be exactly 'Product'
  component: ProductDetail,
  props: route => ({
    shopSlug:    window.location.host.split('.')[0],
    productSlug: route.params.productSlug
  })
}
,
  { path: '/:pathMatch(.*)*',      name: 'NotFound', component: NotFoundPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { top: 0 }
  }
})

router.beforeEach(async (to, from, next) => {
  const host     = window.location.host.split(':')[0]
  const shopSlug = host.split('.')[0]

  const shop = useShopStore()

  if (!shop.current || shop.current.slug !== shopSlug) {
    try {
      await shop.fetchShop(shopSlug)
      await shop.fetchProducts(shopSlug)
    } catch (e) {
      console.error('Error fetching shop data:', e)
    }
  }

  next()
})

export default router