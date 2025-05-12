<template>
  <div>
    <h2 class="text-2xl mb-4">Stores</h2>
    <table class="min-w-full bg-white border">
      <thead>
        <tr>
          <th class="px-4 py-2 border">ID</th>
          <th class="px-4 py-2 border">Name</th>
          <th class="px-4 py-2 border">Owner ID</th>
          <th class="px-4 py-2 border">Description</th>
          <th class="px-4 py-2 border">Created At</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="s in stores" :key="s.id" class="hover:bg-gray-100">
          <td class="px-4 py-2 border">{{ s.id }}</td>
          <td class="px-4 py-2 border">
            <router-link :to="`/stores/${s.id}`" class="text-blue-600 hover:underline">
              {{ s.name || '—' }}
            </router-link>
          </td>
          <td class="px-4 py-2 border">{{ s.ownerId }}</td>
          <td class="px-4 py-2 border">{{ s.description || '—' }}</td>
          <td class="px-4 py-2 border">{{ formatDate(s.createdAt) }}</td>
        </tr>
      </tbody>
    </table>
    <div v-if="stores.length === 0" class="text-gray-600 mt-4">No stores available.</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'StoresList',
  data() {
    return {
      stores: []
    }
  },
  async created() {
    try {
      const res = await axios.get('/admin/stores/', {
        withCredentials: true
      });
      // if your API emits uppercase keys, map them to lowercase here:
      this.stores = res.data.map(s => ({
        id:          s.ID || s.id,
        name:        s.Name || s.name,
        ownerId:     s.OwnerID || s.ownerId,
        description: s.Description || s.description,
        createdAt:   s.CreatedAt || s.createdAt,
      }));
    } catch (err) {
      console.error('Error fetching stores:', err);
    }
  },
  methods: {
    formatDate(dateStr) {
      return new Date(dateStr).toLocaleString();
    }
  }
}
</script>
