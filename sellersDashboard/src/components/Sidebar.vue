<template>
  <aside class="w-64 bg-gray-100 p-4 hidden md:block">
    <div class="mb-6">
      <h2 class="text-lg font-semibold">
        Shop: {{ activeShop ? activeShop.name : 'No Shop Selected' }}
      </h2>
    </div>
    <nav>
      <ul class="space-y-2">
        <li>
          <router-link
            to="/app/products"
            class="block px-2 py-1 rounded hover:bg-gray-200"
            active-class="bg-gray-200 font-medium"
          >Products</router-link>
        </li>
        <!-- add other links similarly -->
      </ul>
    </nav>
  </aside>
</template>

<script>
import { mapGetters } from 'vuex';
export default {
  name: 'Sidebar',
  computed: {
    ...mapGetters('shops', ['activeShop'])
  }
};
</script>
