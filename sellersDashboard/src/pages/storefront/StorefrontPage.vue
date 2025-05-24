<script setup>
import { ref, onMounted } from 'vue'
import { useRoute }      from 'vue-router'
import { useShopStore }  from '@/store/shops'
import { productService }from '@/services/product'

const route    = useRoute()
const shops    = useShopStore()
const shopId   = window.location.hostname.split('.')[0]
const shop     = ref(null)
const products = ref([])
const loading  = ref(true)
const error    = ref(null)

onMounted(async () => {
  loading.value = true
  try {
    shop.value = await shops.fetchPublicShop(shopId)
    products.value = await productService.fetchByShop(shopId)
  } catch (e) {
    console.error(e)
    error.value   = 'Could not load this storefront.'
  } finally {
    loading.value = false
  }
})
</script>
