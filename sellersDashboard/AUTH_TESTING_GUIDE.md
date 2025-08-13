# Auth State Management Testing Guide

## Overview
This guide provides comprehensive testing procedures for the enhanced auth state management system in the sellers dashboard.

## Key Improvements Made

### 1. Enhanced Auth Store (`src/store/auth.js`)
- Added session loading state tracking
- Implemented session expiration detection
- Added persistence error handling
- Enhanced error recovery mechanisms
- Added session verification methods

### 2. Improved API Interceptor (`src/services/api.js`)
- Enhanced token refresh queue logic
- Added max refresh attempts protection
- Expanded refreshable paths coverage
- Better error handling and auth state cleanup
- Dynamic auth store import to avoid circular dependencies

### 3. Robust App Initialization (`src/main.js`)
- Proper session verification on app startup
- Graceful error handling for persistence failures
- Non-blocking initialization for better UX
- Comprehensive logging for debugging

### 4. Enhanced Router Guards (`src/router/index.js`)
- Improved auth state checking logic
- Better handling of race conditions
- Enhanced logging for debugging
- More reliable navigation flow

## Testing Scenarios

### 1. Basic Auth Flow Testing

#### Login Flow
```bash
# Test successful login
1. Navigate to landing page
2. Click login/register
3. Enter valid credentials
4. Verify redirect to profile completion or dashboard
5. Check browser storage for persisted auth state
```

#### Logout Flow
```bash
# Test successful logout
1. From authenticated state
2. Click logout
3. Verify redirect to landing page
4. Check auth state is cleared from storage
5. Verify subsequent API calls fail appropriately
```

### 2. Session Persistence Testing

#### Page Refresh
```bash
# Test auth state survives page refresh
1. Login successfully
2. Navigate to any dashboard page
3. Refresh the page (F5)
4. Verify user remains authenticated
5. Verify active shop is maintained
```

#### Browser Restart
```bash
# Test auth state survives browser restart
1. Login successfully
2. Close browser completely
3. Reopen browser and navigate to app
4. Verify user is still authenticated
5. Verify session verification occurs
```

### 3. Token Refresh Testing

#### Expired Access Token
```bash
# Test automatic token refresh
1. Login successfully
2. Wait for access token to expire (or manually expire it)
3. Make an API request
4. Verify token refresh occurs automatically
5. Verify original request succeeds after refresh
```

#### Failed Token Refresh
```bash
# Test handling of refresh token expiration
1. Login successfully
2. Manually expire both access and refresh tokens
3. Make an API request
4. Verify user is logged out automatically
5. Verify redirect to landing page
```

### 4. Production Environment Testing

#### Network Conditions
```bash
# Test under poor network conditions
1. Enable network throttling in dev tools
2. Perform login/logout operations
3. Test token refresh under slow network
4. Verify graceful handling of network errors
```

#### Multiple Tabs
```bash
# Test auth state synchronization across tabs
1. Open app in multiple browser tabs
2. Login in one tab
3. Verify other tabs update auth state
4. Logout in one tab
5. Verify all tabs are logged out
```

### 5. Error Recovery Testing

#### Storage Failures
```bash
# Test localStorage/sessionStorage failures
1. Disable localStorage in browser
2. Attempt to login
3. Verify app handles storage errors gracefully
4. Verify fallback mechanisms work
```

#### API Failures
```bash
# Test API endpoint failures
1. Block auth API endpoints
2. Attempt login/refresh operations
3. Verify appropriate error messages
4. Verify app doesn't crash
```

## Production Deployment Checklist

### Environment Variables
- [ ] `VITE_API_BASE` is set correctly for production
- [ ] API endpoints are accessible from production domain
- [ ] CORS is configured properly for production domain

### Cookie Configuration
- [ ] Backend sets secure cookies in production
- [ ] SameSite attribute is configured appropriately
- [ ] Cookie domain is set correctly for production

### HTTPS Configuration
- [ ] Production app is served over HTTPS
- [ ] API endpoints use HTTPS
- [ ] Mixed content warnings are resolved

### Performance Optimization
- [ ] Auth state persistence is optimized
- [ ] Token refresh logic doesn't cause excessive requests
- [ ] Session verification is efficient

### Monitoring and Logging
- [ ] Auth errors are logged appropriately
- [ ] Session metrics are tracked
- [ ] Failed login attempts are monitored

## Common Issues and Solutions

### Issue: Auth state lost on page refresh
**Solution**: Check browser storage persistence and session verification logic

### Issue: Infinite redirect loops
**Solution**: Review router guard logic and auth state conditions

### Issue: Token refresh failures
**Solution**: Verify refresh token endpoint and error handling

### Issue: Cross-tab auth synchronization
**Solution**: Implement storage event listeners for auth state changes

## Debug Commands

### Check Auth State
```javascript
// In browser console
const authStore = useAuthStore()
console.log('Auth State:', {
  user: authStore.user,
  isAuthenticated: authStore.isAuthenticated,
  sessionLoading: authStore.sessionLoading,
  lastAuthCheck: authStore.lastAuthCheck
})
```

### Verify Token Refresh
```javascript
// Force token refresh test
api.post('/auth/seller/refresh')
  .then(() => console.log('Refresh successful'))
  .catch(err => console.error('Refresh failed:', err))
```

### Clear Auth State
```javascript
// Reset auth state for testing
const authStore = useAuthStore()
authStore.$reset()
localStorage.clear()
sessionStorage.clear()
```

## Performance Metrics

Monitor these metrics in production:
- Session verification time
- Token refresh success rate
- Auth state persistence reliability
- Router navigation performance
- API request retry rates

## Security Considerations

- Never log sensitive auth data in production
- Implement proper CSRF protection
- Use secure cookie attributes
- Implement rate limiting for auth endpoints
- Monitor for suspicious auth patterns
