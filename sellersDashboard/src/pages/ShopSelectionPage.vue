<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded shadow-md w-full max-w-xl">
      <h2 class="text-2xl font-bold mb-4 text-center">Select a Shop</h2>

      <div v-if="shops && shops.length === 0" class="text-center mb-4">
        <p class="text-gray-600">No shops found. Create your first shop below:</p>
      </div>

      <ul class="space-y-2 mb-6">
        <li v-for="shop in shops" :key="shop.id" class="flex justify-between items-center border p-3 rounded">
          <span>{{ shop.name }}</span>
          <button
            class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 transition"
            @click="selectShop(shop)"
          >
            Enter
          </button>
        </li>
      </ul>

      <form @submit.prevent="createNewShop" class="space-y-4">
        <h3 class="text-lg font-semibold">Create New Shop</h3>
        <input
          v-model="newShopName"
          placeholder="Shop Name"
          required
          class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring focus:border-blue-400"
        />
        <button
          type="submit"
          class="w-full bg-green-600 text-white py-2 rounded hover:bg-green-700 transition"
        >
          Create Shop
        </button>
      </form>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'ShopSelectorPage',
  data() {
    return {
      newShopName: ''
    };
  },
  computed: {
    ...mapState('shops', ['shops']),
    ...mapState('auth', ['user'])
  },
  async created() {
    if (this.user) {
      try {
        await this.fetchShops(this.user.id);
      } catch (err) {
        console.error('Failed to fetch shops:', err);
      }
    }
  },
  methods: {
    ...mapActions('shops', ['fetchShops', 'createShop', 'setActiveShop']),
    async createNewShop() {
      const newShop = await this.createShop({
        userId: this.user.id,
        shopData: { name: this.newShopName }
      });
      this.selectShop(newShop);
    },
    selectShop(shop) {
      this.setActiveShop(shop);
      this.$router.push({ name: 'DashboardHome' });
    }
  }
};
</script>
