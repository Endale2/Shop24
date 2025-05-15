<template>
  <div class="max-w-6xl mx-auto p-6 bg-white rounded shadow">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-2xl font-semibold">Stores</h2>
      <router-link
        to="/stores/create"
        class="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700"
      >
        + Create Store
      </router-link>
    </div>
    <table class="min-w-full bg-white border rounded">
      <thead class="bg-gray-100">
        <tr>
          <th class="px-4 py-2 border">ID</th>
          <th class="px-4 py-2 border">Name</th>
          <th class="px-4 py-2 border">Owner ID</th>
          <th class="px-4 py-2 border">Description</th>
          <th class="px-4 py-2 border">Created At</th>
          <th class="px-4 py-2 border">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="s in stores"
          :key="s.id"
          class="hover:bg-gray-50 border-b"
        >
          <td class="px-4 py-2 border text-sm text-gray-600">{{ s.id }}</td>
          <td class="px-4 py-2 border">
            <router-link :to="`/stores/${s.id}`" class="text-blue-600 hover:underline">
              {{ s.name || '—' }}
            </router-link>
          </td>
          <td class="px-4 py-2 border text-sm text-gray-600">{{ s.ownerId }}</td>
          <td class="px-4 py-2 border text-sm">{{ s.description || '—' }}</td>
          <td class="px-4 py-2 border text-sm text-gray-500">
            {{ formatDate(s.createdAt) }}
          </td>
          <td class="px-4 py-2 border flex space-x-2">
            <router-link
              :to="`/stores/${s.id}/edit`"
              class="px-2 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 text-sm"
            >
              Edit
            </router-link>
            <button
              @click="confirmDelete(s.id)"
              class="px-2 py-1 bg-red-500 text-white rounded hover:bg-red-600 text-sm"
            >
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="stores.length === 0" class="text-gray-600 mt-4">
      No stores available.
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showConfirm"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4"
    >
      <div class="bg-white rounded p-6 w-full max-w-sm">
        <h3 class="text-lg font-semibold mb-4">Delete Store?</h3>
        <p class="mb-6">Are you sure you want to delete this store? This action cannot be undone.</p>
        <div class="flex justify-end space-x-2">
          <button
            @click="showConfirm = false"
            class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
          >
            Cancel
          </button>
          <button
            @click="deleteStore"
            class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
          >
            Yes, Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'StoresList',
  data() {
    return {
      stores: [],
      showConfirm: false,
      toDeleteId: null,
    };
  },
  async created() {
    try {
      const res = await axios.get('/admin/stores/');
      this.stores = res.data.map(s => ({
        id:          s.ID || s.id,
        name:        s.Name || s.name,
        ownerId:     s.OwnerID || s.ownerId,
        description: s.Description || s.description,
        createdAt:   s.CreatedAt || s.createdAt,
      }));
    } catch (e) {
      console.error('Error fetching stores:', e);
    }
  },
  methods: {
    formatDate(dateStr) {
      return new Date(dateStr).toLocaleString();
    },
    confirmDelete(id) {
      this.toDeleteId = id;
      this.showConfirm = true;
    },
    async deleteStore() {
      try {
        await axios.delete(`/admin/stores/${this.toDeleteId}`);
        this.stores = this.stores.filter(s => s.id !== this.toDeleteId);
      } catch (e) {
        console.error('Delete failed:', e);
      } finally {
        this.showConfirm = false;
      }
    }
  }
};
</script>
