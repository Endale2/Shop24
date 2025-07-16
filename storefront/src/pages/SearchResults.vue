<template>
  <div class="bg-gray-50 min-h-screen flex flex-col">
    <StoreFrontHeader v-model="searchQuery" @search="onSearch" />

    <main class="container mx-auto px-4 py-10 flex-1">
      <section class="mb-12">
        <h1 class="text-3xl md:text-4xl font-extrabold text-green-800">Search Results</h1>
        <p v-if="searchQuery" class="text-lg text-gray-600 mt-2">
          Showing results for: <span class="font-semibold text-gray-800">"{{ searchQuery }}"</span>
        </p>
      </section>
      
      <section>
        <div v-if="loading" class="text-center py-16">
          <svg class="mx-auto h-12 w-12 text-green-500 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
          <h3 class="mt-4 text-lg font-semibold text-gray-800">Searching...</h3>
        </div>

        <div v-else-if="shops.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
          <router-link
            v-for="shop in shops"
            :key="shop.id"
            :to="`/shops/${shop.slug}`"
            class="group block bg-white rounded-xl shadow-md overflow-hidden transition-all duration-300 hover:shadow-xl hover:-translate-y-1 border"
          >
            <div class="aspect-video bg-gray-50 flex items-center justify-center">
              <img v-if="shop.image" :src="shop.image" :alt="shop.name" class="w-full h-full object-cover transition-transform group-hover:scale-105" />
              <svg v-else class="w-12 h-12 text-gray-300" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M13.5 21v-7.5c0-.938-.56-1.75-1.5-1.75s-1.5.812-1.5 1.75V21m-4.5 0v-5.625c0-.938.56-1.75 1.5-1.75s1.5.812 1.5 1.75V21m-10.5 0v-3.75c0-.938.56-1.75 1.5-1.75S6 16.312 6 17.25V21m12-3.75v-1.125c0-.938-.56-1.75-1.5-1.75s-1.5.812-1.5 1.75v1.125m-10.5 0a.375.375 0 01.375-.375h11.25a.375.375 0 01.375.375m-12 0v-1.125c0-.938.56-1.75 1.5-1.75s1.5.812 1.5 1.75V21m-1.5-3.75a.375.375 0 01.375-.375H6a.375.375 0 01.375.375m0 0H21m-9.75-12.75H9.375c-.938 0-1.75.56-1.75 1.5v3.375c0 .938.812 1.75 1.75 1.75h2.25c.938 0 1.75-.812 1.75-1.75V6.375c0-.938-.812-1.75-1.75-1.75z" /></svg>
            </div>
            <div class="p-5">
              <h3 class="text-lg font-bold text-gray-800 truncate group-hover:text-green-600">{{ shop.name }}</h3>
              <p class="text-sm text-gray-600 mt-1 h-10 overflow-hidden">{{ shop.description }}</p>
              <div class="mt-4 text-sm font-semibold text-green-600 inline-flex items-center gap-1">
                Visit Shop
                <svg class="w-4 h-4 transition-transform group-hover:translate-x-1" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7"/></svg>
              </div>
            </div>
          </router-link>
        </div>
        
        <div v-else class="text-center py-16 px-6 bg-white rounded-lg shadow-sm">
          <h3 class="text-lg font-semibold text-gray-800">No Shops Found</h3>
          <p class="mt-1 text-sm text-gray-500">Your search did not match any shops. Please try a different query.</p>
        </div>
      </section>
    </main>
    <footer class="bg-white border-t py-6 mt-10 text-center text-gray-600 text-sm">
      &copy; {{ new Date().getFullYear() }} DRPS Storefront. All rights reserved.
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import StoreFrontHeader from '@/components/StoreFrontHeader.vue'
import { searchShops, type Shop } from '@/services/shop' // Assuming types can be imported

const route = useRoute()
const router = useRouter()
const shops = ref<Shop[]>([])
const loading = ref(true)
const searchQuery = ref(route.query.q as string || '')

async function doSearch() {
  if (!searchQuery.value) {
    shops.value = []
    return
  }
  loading.value = true
  try {
    const results = await searchShops({ q: searchQuery.value })
    shops.value = Array.isArray(results) ? results : []
  } catch(error) {
    console.error("Failed to perform search:", error);
    shops.value = []
  } finally {
    loading.value = false
  }
}

// Perform a new search by updating the URL query parameter
function onSearch(query: string) {
  router.push({ path: '/search-results', query: { q: query } })
}

// Watch for changes in the URL query to re-trigger the search
watch(() => route.query.q, (newQuery) => {
    searchQuery.value = newQuery as string || ''
    doSearch()
  },
  { immediate: true } // Use immediate to run the watcher on component mount
)
</script>