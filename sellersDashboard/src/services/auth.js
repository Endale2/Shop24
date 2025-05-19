// src/services/auth.js

// Placeholder API simulation
export default {
  async login({ email, password }) {
    // Simulate API login
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve({
          token: 'dummy-token',
          user: { id: 1, email }
        });
      }, 500);
    });
  },

  async register({ email, password, name }) {
    // Simulate API registration
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve({
          token: 'dummy-token',
          user: { id: 1, email, name }
        });
      }, 500);
    });
  }
};
