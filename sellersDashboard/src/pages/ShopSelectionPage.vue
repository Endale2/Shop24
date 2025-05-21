<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 p-4">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-lg border border-gray-100">
      <h2 class="text-3xl font-extrabold mb-6 text-center text-gray-800">Select Your Shop</h2>

      <div v-if="shops.length === 0" class="text-center mb-6 p-4 bg-yellow-50 rounded-md border border-yellow-200">
        <p class="text-yellow-800 text-lg">No shops found. Let's create your first one!</p>
      </div>

      <div v-if="shops.length" class="mb-8">
        <h3 class="text-xl font-semibold mb-4 text-gray-700 border-b pb-2">Your Existing Shops</h3>
        <ul class="space-y-4">
          <li
            v-for="shop in shops"
            :key="shop.id"
            class="flex items-center justify-between p-4 bg-gray-50 rounded-lg border border-gray-200 transition duration-200 ease-in-out hover:shadow-md hover:border-green-300"
          >
            <span class="text-lg font-medium text-gray-800">{{ shop.name }}</span>
            <button
              class="bg-green-600 text-white px-5 py-2 rounded-full font-semibold hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-opacity-75 transition duration-300 ease-in-out transform hover:scale-105"
              @click="selectShop(shop)"
            >
              Enter Shop
            </button>
          </li>
        </ul>
      </div>

      <form @submit.prevent="createNewShop" class="space-y-5 p-6 bg-green-50 rounded-lg border border-green-200">
        <h3 class="text-xl font-semibold text-green-800">Create a New Shop</h3>
        <input
          v-model="newShopName"
          placeholder="Enter Shop Name"
          required
          class="w-full px-4 py-2 border border-green-300 rounded-md text-gray-800 placeholder-gray-500 focus:ring-2 focus:ring-green-400 focus:border-transparent transition duration-150 ease-in-out"
        />
        <button
          type="submit"
          :disabled="creating || !newShopName.trim()"
          class="w-full bg-green-700 text-white py-3 rounded-md font-bold text-lg hover:bg-green-800 disabled:opacity-60 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-opacity-75 transition duration-300 ease-in-out transform hover:scale-105"
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
      return this.allShops;
    }
  },
  async created() {
    try {
      await this.fetchShops();
    } catch (err) {
      console.error('Failed to fetch shops:', err);
      // Optionally add user-facing error message here
    }
  },
  methods: {
    ...mapActions('shops', ['fetchShops', 'createShop', 'setActiveShop']),
    async createNewShop() {
      if (!this.newShopName.trim()) return;
      this.creating = true;
      try {
        const newShop = await this.createShop({ name: this.newShopName });
        this.setActiveShop(newShop);
        this.$router.push({ name: 'Products' }); // Redirect to your main app page
      } catch (err) {
        console.error('Failed to create shop:', err);
        // Optionally add user-facing error message
      } finally {
        this.creating = false;
        this.newShopName = ''; // Clear input after attempt
      }
    },
    selectShop(shop) {
      this.setActiveShop(shop);
      this.$router.push({ name: 'Products' }); // Redirect to your main app page
    }
  }
};
</script>