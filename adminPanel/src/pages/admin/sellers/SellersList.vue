<template>
  <div>
    <h2 class="text-2xl mb-4">Sellers</h2>
    <ul>
      <li v-for="seller in sellers" :key="seller.id">
        <router-link :to="`/sellers/${seller.id}`">
          {{ seller.name }} ({{ seller.email }})
        </router-link>
      </li>
    </ul>
    <div v-if="sellers.length === 0" class="text-gray-600 mt-4">
      No sellers available.
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SellersList',
  data() {
    return {
      sellers: []
    };
  },
  async created() {
    try {
      const res = await axios.get('/admin/sellers/', {
        withCredentials: true
      });
      this.sellers = res.data.map(s => ({
        id: s.id || s.ID,
        name: s.name || s.Name,
        email: s.email || s.Email,
        phone: s.phone || s.Phone,
        address: s.address || s.Address,
        businessName: s.business_name || s.BusinessName,
        ratings: s.ratings || s.Ratings,
        totalSales: s.total_sales || s.TotalSales,
        reviewsCount: s.reviews_count || s.ReviewsCount,
        createdAt: s.created_at || s.CreatedAt,
        updatedAt: s.updated_at || s.UpdatedAt
      }));
    } catch (err) {
      console.error('Error fetching sellers:', err);
    }
  }
};
</script>
