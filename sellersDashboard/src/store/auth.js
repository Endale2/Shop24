// src/store/modules/auth.js
import { authService } from '@/services/auth';

const state = () => ({ user: null });
const getters = {
  isAuthenticated: s => !!s.user,
  user: s => s.user,
};

const actions = {
  async register({ dispatch }, payload) {
    await authService.register(payload);
    return dispatch('fetchMe');
  },
  async login({ dispatch }, payload) {
    await authService.login(payload);
    return dispatch('fetchMe');
  },
  async fetchMe({ commit }) {
    const user = await authService.me();
    commit('setUser', user);
    return user;
  },
  async logout({ commit }) {
    await authService.logout();
    commit('setUser', null);
  },
};
const mutations = {
  setUser(state, user) {
    state.user = user;
  },
};
export default { namespaced: true, state, getters, actions, mutations };
