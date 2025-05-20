// src/store/modules/shops.js
import shopService from '@/services/shop';

const state = () => ({
  list: [],
  active: null
});

const getters = {
  // All normalized shops
  allShops(state) {
    return state.list;
  },
  // The currently selected shop
  activeShop(state) {
    return state.active;
  }
};

const actions = {
  // Load shops from backend
  async fetchShops({ commit }) {
    const shops = await shopService.fetchShops();
    commit('setList', shops);
    return shops;
  },

  // Create then re-fetch
  async createShop({ dispatch }, payload) {
    const newShop = await shopService.createShop(payload);
    await dispatch('fetchShops');
    return newShop;
  },

  // Mark one active
  setActiveShop({ commit }, shop) {
    commit('setActive', shop);
  }
};

const mutations = {
  setList(state, shops) {
    state.list = shops;
  },
  setActive(state, shop) {
    state.active = shop;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
