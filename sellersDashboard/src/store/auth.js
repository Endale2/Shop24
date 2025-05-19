// src/store/auth.js
import authService from '../services/auth';

const state = () => ({
  token: localStorage.getItem('token') || null,
  user: JSON.parse(localStorage.getItem('user')) || null
});

const getters = {
  isAuthenticated: (state) => !!state.token,
  currentUser: (state) => state.user
};

const mutations = {
  setAuth(state, { token, user }) {
    state.token = token;
    state.user = user;
    localStorage.setItem('token', token);
    localStorage.setItem('user', JSON.stringify(user));
  },
  clearAuth(state) {
    state.token = null;
    state.user = null;
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  }
};

const actions = {
  async login({ commit }, credentials) {
    const response = await authService.login(credentials);
    commit('setAuth', { token: response.token, user: response.user });
  },
  async register({ commit }, data) {
    const response = await authService.register(data);
    commit('setAuth', { token: response.token, user: response.user });
  },
  logout({ commit }) {
    commit('clearAuth');
  }
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
