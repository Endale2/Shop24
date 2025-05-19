// src/store/shops.js
import shopService from '../services/shop';

const state = () => ({
  list: [],
  activeShop: JSON.parse(localStorage.getItem('activeShop')) || null
});

const getters = {
  shopsList: (state) => state.list,
  activeShop: (state) => state.activeShop
};

const mutations = {
  setShops(state, shops) {
    state.list = shops;
  },
  setActiveShop(state, shop) {
    state.activeShop = shop;
    localStorage.setItem('activeShop', JSON.stringify(shop));
  },
  clearShops(state) {
    state.list = [];
    state.activeShop = null;
    localStorage.removeItem('activeShop');
  }
};

const actions = {
  async fetchShops({ commit }) {
    const shops = await shopService.getShops();
    commit('setShops', shops);
  },
  async createShop({ commit }, shopData) {
    const shop = await shopService.createShop(shopData);
    commit('setActiveShop', shop);
  },
  selectShop({ commit }, shop) {
    commit('setActiveShop', shop);
  },
  clear({ commit }) {
    commit('clearShops');
  }
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
