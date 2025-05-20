<template>
  <div class="p-4">
    <h2 class="text-2xl font-semibold mb-4">Products</h2>

    <div v-if="loading" class="text-gray-600">Loading products…</div>
    <div v-else-if="error" class="text-red-500">{{ error }}</div>
    <div v-else>
      <table class="min-w-full bg-white border">
        <thead>
          <tr class="bg-gray-100">
            <th class="px-4 py-2 border">ID</th>
            <th class="px-4 py-2 border">Name</th>
            <th class="px-4 py-2 border">Category</th>
            <th class="px-4 py-2 border">Price</th>
            <th class="px-4 py-2 border">Created</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="prod in products"
            :key="prod.id"
            class="hover:bg-gray-50"
          >
            <td class="px-4 py-2 border text-sm break-all">{{ prod.id }}</td>
            <td class="px-4 py-2 border text-sm">{{ prod.name }}</td>
            <td class="px-4 py-2 border text-sm">{{ prod.category || '—' }}</td>
            <td class="px-4 py-2 border text-sm">
              <!-- only call toFixed if price is a number -->
              {{ typeof prod.price === 'number' ? `$${prod.price.toFixed(2)}` : '—' }}
            </td>
            <td class="px-4 py-2 border text-sm">
              {{ formatDate(prod.createdAt) }}
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="products.length === 0" class="text-gray-500 mt-4">
        No products found.
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import api from '@/services/api';

export default {
  name: 'ProductsPage',
  data() {
    return {
      products: [],
      loading: false,
      error: null
    };
  },
  computed: {
    ...mapGetters('shops', ['activeShop'])
  },
  async created() {
    if (!this.activeShop) {
      this.error = 'No shop selected.';
      return;
    }

    this.loading = true;
    try {
      const res = await api.get(
        `/seller/shops/${this.activeShop.id}/products`
      );

      this.products = res.data.map(p => ({
        id:        p.ID ?? p.id,
        name:      p.Name ?? p.name ?? '—',
        category:  p.Category ?? p.category,
        // default price to 0 if missing or not a number
        price:     isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
        createdAt: p.CreatedAt ?? p.createdAt
      }));
    } catch (err) {
      this.error = 'Failed to load products.';
      console.error(err);
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return dt ? new Date(dt).toLocaleString() : '—';
    }
  }
};
</script>

<style scoped>
/* ensure long IDs wrap */
.break-all {
  word-break: break-all;
}
</style>
