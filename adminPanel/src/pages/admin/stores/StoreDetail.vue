<template>
  <div class="max-w-2xl mx-auto p-6 bg-white rounded shadow">
    <h2 class="text-2xl font-bold mb-4">Store Detail</h2>

    <!-- Loading / Error -->
    <div v-if="loading" class="text-gray-500">Loading...</div>
    <div v-else-if="error" class="text-red-500">{{ error }}</div>

    <!-- Detail & Edit Form -->
    <div v-else>
      <form @submit.prevent="updateStore" class="space-y-4">
        <div>
          <label class="block text-sm font-medium">Name</label>
          <input
            v-model="form.name"
            type="text"
            class="mt-1 block w-full border rounded p-2"
            required
          />
        </div>

        <div>
          <label class="block text-sm font-medium">Description</label>
          <textarea
            v-model="form.description"
            class="mt-1 block w-full border rounded p-2"
            rows="3"
          ></textarea>
        </div>

        <div class="flex space-x-2">
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
            :disabled="updating"
          >
            {{ updating ? 'Saving...' : 'Save Changes' }}
          </button>

          <button
            type="button"
            @click="confirmDelete = true"
            class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
          >
            Delete Store
          </button>
        </div>
      </form>

      <p class="mt-6 text-sm text-gray-600">
        <strong>Owner ID:</strong> {{ store.ownerId }}<br>
        <strong>Created At:</strong> {{ formatDate(store.createdAt) }}<br>
        <strong>Updated At:</strong> {{ formatDate(store.updatedAt) }}
      </p>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="confirmDelete"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 p-4"
    >
      <div class="bg-white rounded p-6 space-y-4 max-w-sm w-full">
        <h3 class="text-lg font-semibold">Confirm Delete</h3>
        <p>Are you sure you want to delete this store? This action cannot be undone.</p>
        <div class="flex justify-end space-x-2">
          <button
            @click="confirmDelete = false"
            class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
          >
            Cancel
          </button>
          <button
            @click="deleteStore"
            class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
            :disabled="deleting"
          >
            {{ deleting ? 'Deleting...' : 'Yes, Delete' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'StoreDetail',
  data() {
    return {
      store: null,
      form: { name: '', description: '' },
      loading: true,
      error: null,
      updating: false,
      deleting: false,
      confirmDelete: false,
    };
  },
  async created() {
    try {
      const { data } = await axios.get(`/admin/stores/${this.$route.params.id}`);
      // normalize keys
      this.store = {
        id: data.ID,
        name: data.Name,
        ownerId: data.OwnerID,
        description: data.Description,
        createdAt: data.CreatedAt,
        updatedAt: data.UpdatedAt,
      };
      // populate form
      this.form.name = this.store.name;
      this.form.description = this.store.description;
    } catch (e) {
      this.error = 'Failed to load store';
    } finally {
      this.loading = false;
    }
  },
  methods: {
    formatDate(dt) {
      return new Date(dt).toLocaleString();
    },
    async updateStore() {
      this.updating = true;
      this.error = null;
      try {
        await axios.patch(
          `/admin/stores/update/${this.store.id}`,
          { name: this.form.name, description: this.form.description }
        );
        // update local store data
        this.store.name = this.form.name;
        this.store.description = this.form.description;
      } catch (e) {
        this.error = 'Failed to update store';
      } finally {
        this.updating = false;
      }
    },
    async deleteStore() {
      this.deleting = true;
      this.error = null;
      try {
        await axios.delete(`/admin/stores/delete/${this.store.id}`);
        this.$router.push({ name: 'StoresList' });
      } catch (e) {
        this.error = 'Failed to delete store';
      } finally {
        this.deleting = false;
        this.confirmDelete = false;
      }
    }
  }
};
</script>
