import { defineStore } from 'pinia'
import { register, login, requestOTP, verifyOTP, getProfile, logout, refreshToken } from '@/services/auth'
import { useCartStore } from './cart'
import router from '../router'
import { getCurrentShopSlug } from '@/services/shop'
import type { RegisterPayload, LoginPayload } from '@/services/auth';

interface User {
  id?: string;
  _id?: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  [key: string]: any;
}

let refreshTimeout: ReturnType<typeof setTimeout> | null = null;
let proactiveRefreshInterval: ReturnType<typeof setInterval> | null = null;

function getAccessTokenFromCookie(): string | null {
  const match = document.cookie.match(/(?:^|; )access_token=([^;]*)/);
  return match ? decodeURIComponent(match[1]) : null;
}

function getRefreshTokenFromCookie(): string | null {
  const match = document.cookie.match(/(?:^|; )refresh_token=([^;]*)/);
  return match ? decodeURIComponent(match[1]) : null;
}

function scheduleProactiveRefresh(store: any) {
  // Clear existing timeout
  if (refreshTimeout) {
    clearTimeout(refreshTimeout);
    refreshTimeout = null;
  }
  
  const token = getAccessTokenFromCookie();
  if (!token) return;
  
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    const exp = payload.exp * 1000;
    const now = Date.now();
    const refreshIn = exp - now - 60 * 1000; // Refresh 1 minute before expiry
    
    if (refreshIn > 0) {
      refreshTimeout = setTimeout(async () => {
        try {
          console.log('[AuthStore] Proactive token refresh triggered');
          await store.refreshSession();
        } catch (error) {
          console.error('[AuthStore] Proactive refresh failed:', error);
          store.user = null;
          store.sessionLoading = false;
        }
      }, refreshIn);
    }
  } catch (error) {
    console.error('[AuthStore] Error parsing token for proactive refresh:', error);
  }
}

function startProactiveRefreshInterval(store: any) {
  // Clear existing interval
  if (proactiveRefreshInterval) {
    clearInterval(proactiveRefreshInterval);
    proactiveRefreshInterval = null;
  }
  
  // Set up 5-minute interval for proactive refresh
  proactiveRefreshInterval = setInterval(async () => {
    if (store.user && getAccessTokenFromCookie()) {
      try {
        console.log('[AuthStore] 5-minute proactive refresh triggered');
        await store.refreshSession();
      } catch (error) {
        console.error('[AuthStore] 5-minute proactive refresh failed:', error);
        store.user = null;
        store.sessionLoading = false;
      }
    }
  }, 5 * 60 * 1000); // 5 minutes
}

function stopProactiveRefresh() {
  if (refreshTimeout) {
    clearTimeout(refreshTimeout);
    refreshTimeout = null;
  }
  if (proactiveRefreshInterval) {
    clearInterval(proactiveRefreshInterval);
    proactiveRefreshInterval = null;
  }
}

function handlePostLoginRedirect(profileComplete: boolean) {
  const shopSlug = getCurrentShopSlug();
  if (shopSlug) {
    if (profileComplete) {
      // Redirect to homepage if profile is complete
      console.log('[AuthStore] Profile complete, redirecting to homepage');
      router.push(`/`);
    } else {
      // Redirect to account page if profile is incomplete
      console.log('[AuthStore] Profile incomplete, redirecting to account page');
      router.push(`/account`);
    }
  } else {
    // Fallback to root
    console.log('[AuthStore] No shop slug, redirecting to root');
    router.push('/');
  }
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
    loading: false,
    error: null as string | null,
    sessionLoading: true,
    otpRequested: false,
    email: '',
    profileComplete: false,
  }),
  getters: {
    isAuthenticated: (state) => !!state.user,
    isUnauthenticated: (state) => !state.user && !state.sessionLoading,
  },
  actions: {
    async register(payload: RegisterPayload) {
      this.loading = true;
      try {
        await register(payload);
        await this.login({ email: payload.email, password: payload.password, shopId: payload.shopId });
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Registration failed';
        setTimeout(() => (this.error = null), 5000);
      } finally {
        this.loading = false;
      }
    },
    async login(payload: LoginPayload) {
      this.loading = true;
      this.error = null;
      try {
        console.log('[AuthStore] Starting login process...');
        await login(payload);
        console.log('[AuthStore] Login successful, fetching profile...');
        await this.fetchProfile();
        console.log('[AuthStore] Profile fetched, setting up refresh mechanisms...');
        scheduleProactiveRefresh(this);
        startProactiveRefreshInterval(this);
        console.log('[AuthStore] Login complete, handling redirect...');
        // Handle redirect after successful login
        handlePostLoginRedirect(this.profileComplete);
      } catch (e: any) {
        console.error('[AuthStore] Login failed:', e);
        this.error = e.response?.data?.error || 'Login failed';
        setTimeout(() => (this.error = null), 5000);
      } finally {
        this.loading = false;
      }
    },
    async requestOTP(email: string) {
      this.loading = true;
      this.error = null;
      try {
        await requestOTP(email);
        this.otpRequested = true;
        this.email = email;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to send OTP';
        setTimeout(() => (this.error = null), 5000);
      } finally {
        this.loading = false;
      }
    },
    async verifyOTP(otp: string) {
      this.loading = true;
      this.error = null;
      try {
        console.log('[AuthStore] Starting OTP verification...');
        const { data } = await verifyOTP(this.email, otp);
        this.user = data.profile;
        this.profileComplete = data.profileComplete;
        console.log('[AuthStore] OTP verification successful, fetching profile...');
        await this.fetchProfile();
        this.otpRequested = false;
        this.email = '';
        console.log('[AuthStore] Setting up refresh mechanisms...');
        scheduleProactiveRefresh(this);
        startProactiveRefreshInterval(this);
        console.log('[AuthStore] OTP verification complete, handling redirect...');
        // Handle redirect after successful OTP verification
        handlePostLoginRedirect(this.profileComplete);
      } catch (e: any) {
        console.error('[AuthStore] OTP verification failed:', e);
        this.error = e.response?.data?.error || 'OTP verification failed';
        setTimeout(() => (this.error = null), 5000);
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async fetchProfile() {
      try {
        console.log('[fetchProfile] Calling /auth/customer/me...')
        console.log('[fetchProfile] Cookies:', document.cookie)
        const { data } = await getProfile()
        console.log('[fetchProfile] Raw response:', data)
        
        // Try to extract the user profile from various possible structures
        let user = data
        let profileComplete = false
        if (data && typeof data === 'object') {
          if ('profile' in data) {
            user = data.profile
            profileComplete = data.profileComplete || false
          } else if ('data' in data) {
            user = data.data
            profileComplete = data.profileComplete || false
          }
        }
        
        // Accept either id or _id, and normalize to id
        if (user && (user.id || user._id)) {
          if (!user.id && user._id) user.id = user._id
          this.user = user
          this.profileComplete = profileComplete
          console.log('[fetchProfile] User set:', user)
          console.log('[fetchProfile] Profile complete:', profileComplete)
        } else {
          console.warn('[fetchProfile] Invalid user object:', user)
          this.user = null
          throw new Error('Invalid user profile')
        }
      } catch (e: any) {
        console.error('[fetchProfile] Failed:', e)
        this.user = null
        this.otpRequested = false
        this.email = ''
        // If it's a 401 error, ensure session loading is set to false
        if (e.response?.status === 401) {
          this.sessionLoading = false
        }
        throw e
      }
    },
    async refreshSession() {
      this.loading = true;
      this.error = null;
      try {
        console.log('[refreshSession] Attempting token refresh...')
        await refreshToken();
        console.log('[refreshSession] Token refresh successful')
        scheduleProactiveRefresh(this);
        startProactiveRefreshInterval(this);
      } catch (e: any) {
        console.error('[refreshSession] Token refresh failed:', e)
        this.error = e.response?.data?.error || 'Token refresh failed';
        this.user = null;
        // If refresh fails with 401, ensure session loading is set to false
        if (e.response?.status === 401) {
          this.sessionLoading = false;
        }
        setTimeout(() => (this.error = null), 5000);
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async verifySession() {
      this.sessionLoading = true;
      console.log('[verifySession] Starting session verification...')
      console.log('[verifySession] Cookies:', document.cookie)
      
      try {
        // First, try to get the current user profile
        console.log('[verifySession] Attempting to fetch profile...')
        await this.fetchProfile();
        
        // If successful, set up proactive refresh
        scheduleProactiveRefresh(this);
        startProactiveRefreshInterval(this);
        console.log('[verifySession] Session verified, user:', this.user)
        console.log('[verifySession] User is authenticated')
      } catch (e: any) {
        console.error('[verifySession] Profile fetch failed:', e)
        
        // If it's a 401 error, try to refresh the token
        if (e.response?.status === 401) {
          console.log('[verifySession] 401 error - attempting token refresh...')
          const refreshToken = getRefreshTokenFromCookie();
          
          if (refreshToken) {
            try {
              await this.refreshSession();
              // If refresh successful, try to fetch profile again
              await this.fetchProfile();
              scheduleProactiveRefresh(this);
              startProactiveRefreshInterval(this);
              console.log('[verifySession] Session restored via refresh')
              console.log('[verifySession] User is authenticated')
              return;
            } catch (refreshError) {
              console.error('[verifySession] Token refresh failed:', refreshError)
            }
          }
        }
        
        // If we get here, the user is definitely not authenticated
        console.log('[verifySession] Setting user as unauthenticated')
        this.user = null;
        
        // Clear localStorage and reset all relevant stores
        localStorage.clear();
        const cartStore = useCartStore();
        if (cartStore && typeof cartStore.$reset === 'function') cartStore.$reset();
        try {
          const { useWishlistStore } = await import('./wishlist');
          const wishlistStore = useWishlistStore();
          if (wishlistStore && typeof wishlistStore.$reset === 'function') wishlistStore.$reset();
        } catch {}
      } finally {
        // Always ensure session loading is false
        this.sessionLoading = false;
        console.log('[verifySession] Session loading set to false')
        console.log('[verifySession] Final auth state - User:', this.user ? 'Authenticated' : 'Unauthenticated')
      }
    },
    async logout() {
      console.log('[AuthStore] Starting logout process...');
      try {
        await logout();
        console.log('[AuthStore] Logout API call successful');
      } catch (error) {
        console.error('[logout] Logout API call failed:', error)
      }
      
      // Clear all auth state immediately
      console.log('[AuthStore] Clearing auth state...');
      this.user = null;
      this.otpRequested = false;
      this.email = '';
      this.profileComplete = false;
      this.sessionLoading = false;
      console.log('[AuthStore] User is now unauthenticated');
      
      // Stop all refresh mechanisms
      stopProactiveRefresh();
      
      // Clear localStorage and reset all relevant stores
      localStorage.clear();
      const cartStore = useCartStore();
      if (cartStore && typeof cartStore.$reset === 'function') cartStore.$reset();
      try {
        const { useWishlistStore } = await import('./wishlist');
        const wishlistStore = useWishlistStore();
        if (wishlistStore && typeof wishlistStore.$reset === 'function') wishlistStore.$reset();
      } catch {}
      
      // Redirect to login page
      const shopSlug = getCurrentShopSlug();
      if (shopSlug) {
        console.log('[AuthStore] Redirecting to login page');
        router.push(`/login`);
      } else {
        console.log('[AuthStore] Redirecting to root');
        router.push('/');
      }
    },
  },
  // No persistence: always rely on backend cookies for session state
});