import axios from 'axios';
// You might want to import the router here if you handle redirecting to login
// directly within the interceptor on refresh failure, though handling in AdminLayout is fine.
// import router from '../router';

let isRefreshing = false;
let refreshSubscribers = [];

// Helper function to handle subscribers waiting for token refresh
function subscribeToRefresh(callback) {
  refreshSubscribers.push(callback);
}

// Helper function to notify subscribers once the token is refreshed
function onRefreshed(newToken) {
  refreshSubscribers.forEach(callback => callback(newToken));
  refreshSubscribers = [];
}

// Helper function to handle refresh failure and clear subscribers
function onRefreshFailed(error) {
    refreshSubscribers.forEach(callback => callback(null, error)); // Notify with error
    refreshSubscribers = [];
    // Optional: Redirect to login here if you want the interceptor to force logout on refresh failure
    // router.replace('/login');
}


// Axios response interceptor for token refresh
axios.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    const status = error.response ? error.response.status : null;

    // Check if the error is 401 Unauthorized and it's not a request that already failed and is being retried
    // Also, ensure it's not the refresh request itself that failed (though unlikely to get 401 on refresh endpoint itself unless refresh token is missing/invalid)
    if (status === 401 && originalRequest && !originalRequest._retry) {
      originalRequest._retry = true; // Mark the original request as retried

      // If a refresh is not already in progress
      if (!isRefreshing) {
        isRefreshing = true; // Set refresh in progress flag

        try {
          // Attempt to refresh the token by calling the refresh endpoint
          // The browser will automatically send the httpOnly refresh_token cookie
          console.log('Attempting token refresh...');
          const refreshResponse = await axios.post('/auth/admin/refresh', null, {
             // Ensure withCredentials is true for this request too, if needed explicitly
             withCredentials: true
          });

          isRefreshing = false; // Refresh is complete
          console.log('Token refresh successful.');
          // Notify all waiting subscribers that refresh is complete
          // Pass the new token if the backend returns it in the body (less common with httpOnly)
          // or null if tokens are only set via cookies (your case).
          onRefreshed(refreshResponse.data); // Pass potential data (like user info) or null

          // Retry the original request immediately for the current flow
          // This is for the *first* request that triggered the refresh
          return axios(originalRequest); // Return the promise of the retried request


        } catch (refreshError) {
          isRefreshing = false; // Refresh failed
          console.error('Token refresh failed:', refreshError);
          // Notify waiting subscribers about the failure
          onRefreshFailed(refreshError);

          // Redirect to login if the refresh itself fails (handles cases where refresh token is invalid)
          // Be careful with router logic inside interceptors; ensure the router is available.
          // This might be better handled in the component calling the API, after the promise is rejected.
           // router.replace('/login'); // Example direct redirect

          // Reject the original request's promise with the refresh error
          return Promise.reject(refreshError);
        }
      }

      // If a refresh is already in progress, queue the original request
      // When the active refresh finishes, the subscriber will retry this request
      return new Promise((resolve, reject) => {
          subscribeToRefresh((newToken, refreshError) => {
              if (refreshError) {
                  reject(refreshError); // Reject if refresh failed
              } else {
                  // Retry the original request with the new token (browser sends new cookies)
                  resolve(axios(originalRequest));
              }
          });
      });
    }

    // If the error is not a 401 or is already retried, propagate the error
    return Promise.reject(error);
  }
);

const auth = {
  // user state is kept here in memory, lost on page refresh
  // it should be re-fetched on page load
  user: null,
  // Add error state for fetching user on load
  fetchUserError: null,

  // Method to perform login
  async login(credentials) {
    // This sends credentials, backend sets httpOnly cookies on success
    const response = await axios.post('/auth/admin/login', credentials);
    // On successful login, you typically redirect the user.
    // The browser will automatically send the new httpOnly access_token cookie on the next request (e.g. to /me).
    return response;
  },

  // Method to check if the user is currently authenticated
  async isAuthenticated() {
    try {
      // Attempt to fetch user data using the access_token cookie
      // If access_token is expired, interceptor should attempt refresh
      // If refresh fails, this request will fail and throw
      const res = await axios.get('/auth/admin/me');
      this.user = res.data; // Update user state on successful fetch
      this.fetchUserError = null; // Clear any previous fetch errors
      return true;
    } catch (error) {
      // This catch block handles:
      // 1. Initial /me failing (no access_token or invalid)
      // 2. Interceptor caught 401 -> refresh attempted -> refresh failed
      console.error('Authentication check failed:', error);
      this.user = null; // Clear user state
      this.fetchUserError = error.message || 'Failed to fetch user info.'; // Set error message
      // Optionally log out user completely on severe failure
      // await this.logout(); // If logout cleans up server-side state
      return false;
    }
  },

   // Method to get user data, fetches if not already in state
   async getUser() {
     // If user is already in memory, return it immediately
     if (this.user) {
       return this.user;
     }

     // If user is not in memory, attempt to fetch it
     // This will use the access_token cookie.
     // If access_token is expired but refresh_token is valid, interceptor handles refresh.
     // If refresh also fails, the fetchUserDetails will throw, caught by the caller (e.g., AdminLayout).
     const isAuthenticated = await this.isAuthenticated();
     if (isAuthenticated) {
       return this.user;
     } else {
       // isAuthenticated already logged error and set this.user = null
       // The component calling getUser() (like AdminLayout) should handle the false return value
       // and redirect to login.
       throw new Error(this.fetchUserError || 'User not authenticated.'); // Propagate fetch error
     }
   },

   // Method to handle logout
   async logout() {
     try {
       // Call backend logout endpoint to invalidate tokens server-side
       // Backend should also clear httpOnly cookies (e.g., Max-Age=0)
       await axios.post('/auth/admin/logout', null, { withCredentials: true });
       console.log('Logged out server-side');
     } catch (error) {
       console.error('Logout failed server-side:', error);
       // Continue client-side cleanup even if server logout fails
     } finally {
       // Clear client-side state
       this.user = null;
       this.fetchUserError = null;
       // Redirect to login page
       // Ensure your router instance is accessible here if needed
       // router.replace('/login'); // Example
     }
   }
};

export default auth; // Export the auth object