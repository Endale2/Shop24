<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const auth = useAuthStore()
const router = useRouter()

const firstName    = ref('')
const lastName     = ref('')
const phone        = ref('')
const businessName = ref('')
const loading      = ref(false)
const error        = ref(null)

onMounted(() => {
  const u = auth.user || {}
  firstName.value    = u.first_name    || ''
  lastName.value     = u.last_name     || ''
  phone.value        = u.phone         || ''
  businessName.value = u.business_name || ''
})

async function submit() {
  loading.value = true
  error.value   = null

  // build a plain object payload
  const payload = {
    firstName:    firstName.value,
    lastName:     lastName.value,
    phone:        phone.value,
    businessName: businessName.value,
  }

  try {
    // use your authStore action which calls PATCH /auth/seller/update-profile
    await auth.updateProfile(payload)
    // redirect on success
    const dest = auth.user.shop_ids?.length ? 'Dashboard' : 'ShopSelection'
    router.push({ name: dest })
  } catch (err) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}
</script>
