<template>
  <div>
    <h2 class="text-2xl mb-4">Staffs</h2>
    <ul>
      <li v-for="c in staffList" :key="c.id">
        <router-link :to="`/staffs/${c.id}`">
          {{ c.firstName }} {{ c.lastName }}
        </router-link>
      </li>
    </ul>
    <div v-if="staffList.length === 0" class="text-gray-600 mt-4">
      No staffList available.
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'StaffList',
  data() {
    return {
      staffList: []
    }
  },
  async created() {
    try {
      const res = await axios.get('/admin/staffs/', {
        withCredentials: true
      });
      // Normalize uppercase JSON keys (from Go structs) to lowercase
      this.staffList = res.data.map(c => ({
        id:        c.ID          || c.id,
        firstName: c.FirstName   || c.firstName,
        lastName:  c.LastName    || c.lastName,
        email:     c.Email       || c.email,
        phone:     c.Phone       || c.phone,
        address:   c.Address     || c.address,
        city:      c.City        || c.city,
        state:     c.State       || c.state,
        postalCode:c.PostalCode  || c.postalCode,
        country:   c.Country     || c.country,
        createdAt: c.CreatedAt   || c.createdAt,
      }));
    } catch (err) {
      console.error('Error fetching staffList:', err);
    }
  }
}
</script>
