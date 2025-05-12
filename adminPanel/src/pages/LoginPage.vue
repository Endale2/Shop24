<template>
    <div class="max-w-sm mx-auto p-6 border rounded">
      <h2 class="text-xl mb-4">Admin Login</h2>
      <form @submit.prevent="submit">
        <div class="mb-4">
          <input v-model="email" type="email" placeholder="Email" class="w-full p-2 border rounded" required />
        </div>
        <div class="mb-4">
          <input v-model="password" type="password" placeholder="Password" class="w-full p-2 border rounded" required />
        </div>
        <button type="submit" class="w-full px-4 py-2 bg-green-600 text-white rounded">Login</button>
        <p v-if="error" class="text-red-500 mt-2">{{ error }}</p>
      </form>
    </div>
  </template>
  
  <script>
  import auth from '../services/auth';
  export default {
    data() {
      return { email: '', password: '', error: null };
    },
    methods: {
      async submit() {
        this.error = null;
        try {
          await auth.login({ email: this.email, password: this.password });
          this.$router.push({ name: 'Dashboard' });
        } catch (e) {
          this.error = 'Invalid credentials';
        }
      }
    }
  };
  </script>