// src/main.js
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

const app = createApp(App);

// Register global store and router
app.use(store);
app.use(router);

// Mount the application
app.mount('#app');