<template>
  <section class="relative overflow-hidden border-b bg-gradient-to-br from-slate-50 to-slate-200">
    <div class="absolute top-0 left-0 h-96 w-96 -translate-x-1/2 rounded-full bg-gradient-to-br from-green-200 to-cyan-200 opacity-50 blur-3xl"></div>
    <div class="absolute bottom-0 right-0 h-96 w-96 translate-x-1/2 rounded-full bg-gradient-to-br from-emerald-200 to-lime-200 opacity-40 blur-3xl"></div>

    <div class="relative z-10">
      <nav class="container mx-auto flex h-20 items-center justify-between px-4 sm:px-6 lg:px-8">
        <router-link to="/" class="group flex items-center gap-2.5 outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-4 rounded-sm">
          <div class="flex h-9 w-9 items-center justify-center rounded-lg bg-gradient-to-br from-green-500 to-emerald-500 text-white shadow-md transition-transform duration-300 group-hover:-rotate-12 group-hover:scale-110">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="h-5 w-5">
              <path d="M12.378 1.602a.75.75 0 00-.756 0L3.366 6.126A.75.75 0 003 6.82v10.36a.75.75 0 00.366.694l8.256 4.524a.75.75 0 00.756 0l8.256-4.524a.75.75 0 00.366-.694V6.82a.75.75 0 00-.366-.694L12.378 1.602zM12 7.5a.75.75 0 01.75.75v3.19l4.83-2.651a.75.75 0 11.74 1.342l-5.25 2.885a.75.75 0 01-.74 0l-5.25-2.885a.75.75 0 01.74-1.342L11.25 11.44V8.25A.75.75 0 0112 7.5zM11.25 14.53v3.19l-4.83-2.651a.75.75 0 01.74-1.342l4.09 2.247zm1.5 0l4.09-2.247a.75.75 0 11.74 1.342L12.75 17.72v-3.19z" />
            </svg>
          </div>
          <span class="text-xl font-bold tracking-tight text-slate-800">ShopFinder</span>
        </router-link>
        
        <div class="flex items-center gap-4">
          <slot>
            <a href="#" class="text-sm font-semibold text-slate-700 hover:text-green-600">Sign In</a>
            <a href="#" class="rounded-full bg-green-600 px-4 py-2 text-sm font-semibold text-white shadow-sm transition-colors hover:bg-green-700">Add your Shop</a>
          </slot>
        </div>
      </nav>

      <div class="container mx-auto px-4 sm:px-6 lg:px-8 py-20 text-center">
        <h1 class="text-4xl font-extrabold tracking-tight text-slate-900 sm:text-5xl lg:text-6xl">
          Discover Local Shops in Ethiopia
        </h1>
        <p class="mt-6 max-w-2xl mx-auto text-lg text-slate-600">
          Your gateway to the best local stores, boutiques, and services. Start exploring now.
        </p>

        <div class="mt-10 mx-auto max-w-xl">
          <div class="relative flex items-center rounded-full shadow-lg bg-white">
            <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-5">
              <svg class="h-5 w-5 text-slate-400" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M9 3.5a5.5 5.5 0 100 11 5.5 5.5 0 000-11zM2 9a7 7 0 1112.452 4.391l3.328 3.329a.75.75 0 11-1.06 1.06l-3.329-3.328A7 7 0 012 9z" clip-rule="evenodd" />
              </svg>
            </div>
            <input
              v-model="model"
              @keyup.enter="onSearch"
              type="search"
              placeholder="Enter a shop name or category..."
              class="w-full rounded-full border-0 bg-transparent py-3.5 pl-12 pr-32 text-slate-900 placeholder:text-slate-400 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-green-500"
            />
            <button @click="onSearch" class="absolute inset-y-1.5 right-1.5 flex items-center rounded-full bg-green-600 px-6 font-semibold text-white shadow-sm transition-colors hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2">
              Search
            </button>
          </div>
        </div>
        
        <div class="mt-12 text-sm">
          <span class="font-semibold text-slate-600">Featured Shops:</span>
          <div class="mt-2 inline-flex flex-wrap justify-center gap-x-6 gap-y-2">
            <a v-for="shop in featuredShops" :key="shop" href="#" class="text-slate-500 hover:text-green-600 transition-colors hover:underline">
              {{ shop }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
// We add featuredShops data to power the new section.
const featuredShops = ['Shega Crafts', 'BunaHub Roasters', 'Zemen Leather', 'Addis Books'];

const model = defineModel<string>()
const emit = defineEmits(['search'])

function onSearch() {
  emit('search', model.value)
}
</script>