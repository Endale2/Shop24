import { defineStore } from 'pinia'
import { register, login, requestOTP, verifyOTP, getProfile, logout, refreshToken } from '@/services/auth'
import { useCartStore } from './cart'

let refreshTimeout: ReturnType<typeof setTimeout> | null = null;

function getAccessTokenFromCookie(): string | null {
  const match = document.cookie.match(/(?:^|; )access_token=([^;]*)/);
  return match ? decodeURIComponent(match[1]) : null;
}

function scheduleProactiveRefresh(store: any) {
  const token = getAccessTokenFromCookie();
  if (!token) return;
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    const exp = payload.exp * 1000;
    const now = Date.now();
    const refreshIn = exp - now - 60 * 1000;
    if (refreshTimeout) clearTimeout(refreshTimeout);
    if (refreshIn > 0) {
      refreshTimeout = setTimeout(async () => {
        try {
          await store.refreshSession();
        } catch {
          store.user = null;
        }
      }, refreshIn);
    }
  } catch {}
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
        await login(payload);
        scheduleProactiveRefresh(this);
        await this.fetchProfile();
      } catch (e: any) {
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
        const { data } = await verifyOTP(this.email, otp);
        this.user = data.profile;
        this.profileComplete = data.profileComplete;
        await this.fetchProfile();
        this.otpRequested = false;
        this.email = '';
        scheduleProactiveRefresh(this);
      } catch (e: any) {
        this.error = e.response?.data?.error || 'OTP verification failed';
        setTimeout(() => (this.error = null), 5000);
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async fetchProfile() {
      try {
        console.log('[fetchProfile] Calling /me...')
        // Log cookies for debugging
        console.log('[fetchProfile] document.cookie:', document.cookie)
        const { data } = await getProfile()
        console.log('[fetchProfile] Raw response:', data)
        // Try to extract the user profile from various possible structures
        let user = data
        if (data && typeof data === 'object') {
          if ('profile' in data) user = data.profile
          else if ('data' in data) user = data.data
        }
        // Accept either id or _id, and normalize to id
        if (user && (user.id || user._id)) {
          if (!user.id && user._id) user.id = user._id
          this.user = user
          console.log('[fetchProfile] User set:', user)
        } else {
          console.warn('[fetchProfile] Invalid user object:', user)
          this.user = null
          throw new Error('Invalid user profile')
        }
      } catch (e) {
        console.error('[fetchProfile] Failed:', e)
        this.user = null
        this.otpRequested = false
        this.email = ''
        throw e
      }
    },
    async refreshSession() {
      this.loading = true;
      this.error = null;
      try {
        await refreshToken();
        scheduleProactiveRefresh(this);
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Token refresh failed';
        this.user = null;
        setTimeout(() => (this.error = null), 5000);
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async verifySession() {
      this.sessionLoading = true;
      // Log cookies for debugging
      console.log('[verifySession] document.cookie:', document.cookie)
      try {
        console.log('[verifySession] Calling fetchProfile...')
        await this.fetchProfile();
        scheduleProactiveRefresh(this);
        console.log('[verifySession] Session verified, user:', this.user)
      } catch (e) {
        console.error('[verifySession] Session verification failed:', e)
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
        this.sessionLoading = false;
      }
    },
    async logout() {
      await logout();
      this.user = null;
      this.otpRequested = false;
      this.email = '';
      this.profileComplete = false;
      if (refreshTimeout) clearTimeout(refreshTimeout);
      // Clear localStorage and reset all relevant stores
      localStorage.clear();
      const cartStore = useCartStore();
      if (cartStore && typeof cartStore.$reset === 'function') cartStore.$reset();
      try {
        const { useWishlistStore } = await import('./wishlist');
        const wishlistStore = useWishlistStore();
        if (wishlistStore && typeof wishlistStore.$reset === 'function') wishlistStore.$reset();
      } catch {}
    },
  },
  persist: {
    paths: ['user', 'otpRequested', 'email', 'profileComplete']
  },
  // No persistence: always rely on backend cookies for session state
});