// src/store/modules/auth.js
import authService from '@/services/auth';

const state = () => ({
  user: null
});

const getters = {
  isAuthenticated: (s) => !!s.user,
  user: (s) => s.user
};

const actions = {
  // ← Add this!
  async register({ dispatch }, { username, email, password }) {
    // call your real backend:
    await authService.register({ username, email, password });
    // once registered, fetch the profile into the store
    return dispatch('fetchMe');
  },

  async login({ dispatch }, { email, password }) {
    await authService.login({ email, password });
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
  }
};

const mutations = {
  setUser(state, u) {
    state.user = u;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
