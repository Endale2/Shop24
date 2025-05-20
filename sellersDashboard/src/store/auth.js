// store/modules/auth.js
import authService from '@/services/auth';

export default {
  namespaced: true,
  state: { user: null },
  getters: {
    isAuthenticated: (s) => !!s.user,
    user: (s) => s.user,
  },
  actions: {
    async register({ dispatch }, creds) {
      const data = await authService.register(creds);
      // After register, fetch profile
      await dispatch('fetchMe');
      return data;
    },
    async login({ dispatch }, creds) {
      await authService.login(creds);
      return dispatch('fetchMe');
    },
    async logout({ commit }) {
      await authService.logout();
      commit('setUser', null);
    },
    async fetchMe({ commit }) {
      const user = await authService.me();
      commit('setUser', user);
      return user;
    },
  },
  mutations: {
    setUser(state, user) {
      state.user = user;
    },
  },
};
