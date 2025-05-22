//src/store/modules/shops.js
import { shopService } from '@/services/shop';

const state = () => ({ list: [], active: null });
const getters = {
  allShops: s => s.list,
  activeShop: s => s.active,
};
const actions = {
  async fetchShops({ commit }) {
    const shops = await shopService.fetchShops();
    commit('setList', shops);
    return shops;
  },
  async createShop({ dispatch }, payload) {
    const newShop = await shopService.createShop(payload);
    await dispatch('fetchShops');
    return newShop;
  },
  setActiveShop({ commit }, shop) {
    commit('setActive', shop);
  },
};
const mutations = {
  setList(state, shops) {
    state.list = shops;
  },
  setActive(state, shop) {
    state.active = shop;
  },
};
export default { namespaced: true, state, getters, actions, mutations };
