<template>
  <h1 class="text-2xl mb-4">Collections</h1>
  <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-4">
    <router-link
      v-for="c in collections"
      :key="c.id"
      :to="`/collections/${c.handle}`"
      class="block border rounded overflow-hidden hover:shadow"
    >
      <img :src="c.image" class="w-full h-40 object-cover" alt="" />
      <div class="p-2">
        <h2 class="font-semibold">{{ c.title }}</h2>
      </div>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchCollections } from '@/services/collections'
import type { Collection } from '@/services/collections'

const collections = ref<Collection[]>([])

onMounted(async () => {
  collections.value = await fetchCollections()
})
</script>
