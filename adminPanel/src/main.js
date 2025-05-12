import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from 'axios';

// Set base URL for axios
axios.defaults.baseURL = 'http://localhost:8080';
// Include cookies on requests
axios.defaults.withCredentials = true;

const app = createApp(App);
app.use(router);
app.provide('axios', axios);
app.mount('#app');