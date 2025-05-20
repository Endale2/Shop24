// src/store/index.js
import { createStore } from 'vuex';
import auth  from './auth';   // assume you have this
import shops from './shops';

export default createStore({
  modules: {
    auth,
    shops
  }
});
