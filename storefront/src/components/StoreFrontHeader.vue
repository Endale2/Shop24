<template>
  <header class="sticky top-0 z-50 w-full border-b border-gray-200/70 bg-white/80 shadow-sm backdrop-blur-lg">
    <div class="container mx-auto flex items-center justify-between gap-4 px-4 py-3">
      
      <router-link to="/" class="group flex flex-shrink-0 items-center gap-3 outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2 rounded-lg">
        <div class="rounded-full bg-gradient-to-br from-green-500 to-green-600 p-2 shadow-md transition-all duration-300 group-hover:scale-110 group-hover:shadow-lg group-active:scale-100">
          <svg class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-3-3v6" />
          </svg>
        </div>
        <span class="text-2xl font-extrabold tracking-tight text-gray-800 transition-colors group-hover:text-green-700">DRPS</span>
      </router-link>

      <div class="flex flex-1 justify-center px-4">
        <div class="relative w-full max-w-lg">
          <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-4">
            <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
            </svg>
          </div>
          <input
            v-model="model"
            @keyup.enter="onSearch"
            type="search"
            placeholder="Search shops..."
            class="w-full rounded-full border border-gray-300 bg-gray-100/70 py-2.5 pl-11 pr-28 text-gray-900 transition focus:border-green-500 focus:bg-white focus:outline-none focus:ring-2 focus:ring-green-500"
          />
          <button @click="onSearch" class="absolute inset-y-1.5 right-1.5 flex items-center rounded-full bg-green-600 px-5 font-semibold text-white shadow-sm transition-colors hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2">
            Search
          </button>
        </div>
      </div>

      <div class="flex flex-shrink-0 items-center">
        <slot />
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
// This script remains the same. 
// defineModel provides a clean, modern way to handle v-model.
const model = defineModel<string>()
const emit = defineEmits(['search'])

function onSearch() {
  emit('search', model.value)
}
</script>