// src/store/index.js
import { createStore } from 'vuex';
import auth from './auth';
import shops from './shops';

export default createStore({
  modules: {
    auth,
    shops
  }
});
