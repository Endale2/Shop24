<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-6 rounded shadow w-full max-w-lg">
      <h2 class="text-2xl font-bold mb-4 text-center">Select a Shop</h2>

      <!-- If no shops -->
      <div v-if="shops.length === 0" class="text-center mb-4">
        <p class="text-gray-600">No shops found. Create your first shop below:</p>
      </div>

      <!-- List of existing shops -->
      <ul v-if="shops.length" class="space-y-2 mb-6">
        <li
          v-for="shop in shops"
          :key="shop.id"
          class="flex justify-between items-center border p-3 rounded"
        >
          <span>{{ shop.name }}</span>
          <button
            class="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600"
            @click="selectShop(shop)"
          >
            Enter
          </button>
        </li>
      </ul>

      <!-- Create a new shop -->
      <form @submit.prevent="createNewShop" class="space-y-3">
        <h3 class="text-lg font-semibold">Create New Shop</h3>
        <input
          v-model="newShopName"
          placeholder="Shop Name"
          required
          class="w-full px-2 py-1 border rounded text-sm"
        />
        <button
          type="submit"
          :disabled="creating"
          class="w-full bg-green-500 text-white py-1 rounded hover:bg-green-600 disabled:opacity-50"
        >
          {{ creating ? 'Creating...' : 'Create Shop' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'ShopSelectionPage',
  data() {
    return {
      newShopName: '',
      creating: false,
    };
  },
  computed: {
    ...mapGetters('shops', ['allShops']),
    shops() {
      // alias for clarity
      return this.allShops;
    }
  },
  async created() {
    try {
      await this.fetchShops();
    } catch (err) {
      console.error('Failed to fetch shops:', err);
    }
  },
  methods: {
    ...mapActions('shops', ['fetchShops', 'createShop', 'setActiveShop']),
    async createNewShop() {
      if (!this.newShopName.trim()) return;
      this.creating = true;
      try {
        // createShop should return the newly created shop
        const newShop = await this.createShop({ name: this.newShopName });
        this.setActiveShop(newShop);
        this.$router.push({ name: 'Products' });
      } catch (err) {
        console.error('Failed to create shop:', err);
      } finally {
        this.creating = false;
      }
    },
    selectShop(shop) {
      this.setActiveShop(shop);
      this.$router.push({ name: 'Products' });
    }
  }
};
</script>
