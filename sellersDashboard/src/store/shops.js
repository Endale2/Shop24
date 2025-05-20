// store/modules/shops.js
import shopService from '@/services/shop';

export default {
  namespaced: true,
  state: { list: [], active: null },
  getters: {
    all: (s) => s.list,
    activeShop: (s) => s.active,
  },
  actions: {
    async load({ commit }) {
      const shops = await shopService.fetchShops();
      commit('setList', shops);
      return shops;
    },
    async create({ dispatch }, payload) {
      await shopService.createShop(payload);
      return dispatch('load');
    },
    select({ commit }, shop) {
      commit('setActive', shop);
    },
  },
  mutations: {
    setList(state, shops) { state.list = shops; },
    setActive(state, shop) { state.active = shop; },
  },
};
