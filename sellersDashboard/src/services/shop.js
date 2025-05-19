// src/services/shop.js

let dummyShops = {
  'user_123': [
    { id: 'shop_1', name: 'My First Shop' },
    { id: 'shop_2', name: 'My Second Shop' }
  ]
};

export default {
  async fetchShops(userId) {
    // Simulate API delay
    await new Promise((resolve) => setTimeout(resolve, 300));
    return dummyShops[userId] || [];
  },

  async createShop(userId, shopData) {
    const newShop = {
      id: `shop_${Date.now()}`,
      name: shopData.name
    };

    if (!dummyShops[userId]) {
      dummyShops[userId] = [];
    }
    dummyShops[userId].push(newShop);

    // Simulate API delay
    await new Promise((resolve) => setTimeout(resolve, 300));
    return newShop;
  }
};
