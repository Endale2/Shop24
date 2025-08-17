package services

import (
	"encoding/json"
	"sync"
	"time"
)

// CacheItem represents a cached item with expiration
type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// CacheService provides in-memory caching functionality
type CacheService struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

// Global cache instance
var cache *CacheService
var cacheOnce sync.Once

// GetCacheService returns the singleton cache service instance
func GetCacheService() *CacheService {
	cacheOnce.Do(func() {
		cache = &CacheService{
			items: make(map[string]CacheItem),
		}
		// Start cleanup goroutine
		go cache.cleanup()
	})
	return cache
}

// Set stores a value in the cache with expiration
func (c *CacheService) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration),
	}
}

// Get retrieves a value from the cache
func (c *CacheService) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	item, exists := c.items[key]
	if !exists {
		return nil, false
	}
	
	// Check if expired
	if time.Now().After(item.Expiration) {
		// Remove expired item
		delete(c.items, key)
		return nil, false
	}
	
	return item.Value, true
}

// Delete removes a value from the cache
func (c *CacheService) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	delete(c.items, key)
}

// Clear removes all items from the cache
func (c *CacheService) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.items = make(map[string]CacheItem)
}

// GetJSON retrieves and unmarshals JSON data from cache
func (c *CacheService) GetJSON(key string, target interface{}) bool {
	value, exists := c.Get(key)
	if !exists {
		return false
	}
	
	jsonData, ok := value.([]byte)
	if !ok {
		return false
	}
	
	err := json.Unmarshal(jsonData, target)
	return err == nil
}

// SetJSON marshals and stores JSON data in cache
func (c *CacheService) SetJSON(key string, value interface{}, duration time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	
	c.Set(key, jsonData, duration)
	return nil
}

// cleanup removes expired items periodically
func (c *CacheService) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			now := time.Now()
			for key, item := range c.items {
				if now.After(item.Expiration) {
					delete(c.items, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}

// Stats returns cache statistics
func (c *CacheService) Stats() map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	totalItems := len(c.items)
	expiredItems := 0
	now := time.Now()
	
	for _, item := range c.items {
		if now.After(item.Expiration) {
			expiredItems++
		}
	}
	
	return map[string]interface{}{
		"totalItems":   totalItems,
		"activeItems":  totalItems - expiredItems,
		"expiredItems": expiredItems,
	}
}

// Cache duration constants
const (
	CacheShort  = 5 * time.Minute   // For frequently changing data
	CacheMedium = 30 * time.Minute  // For moderately stable data
	CacheLong   = 2 * time.Hour     // For stable data
	CacheDay    = 24 * time.Hour    // For very stable data
)

// Cache key generators
func ThemeCacheKey(shopID string) string {
	return "theme:" + shopID
}

func ComponentCacheKey(shopID, componentType string) string {
	return "components:" + shopID + ":" + componentType
}

func LayoutCacheKey(shopID, pageType string) string {
	return "layout:" + shopID + ":" + pageType
}

func CSSCacheKey(shopID string) string {
	return "css:" + shopID
}

func ShopCacheKey(shopSlug string) string {
	return "shop:" + shopSlug
}

// Helper functions for common caching patterns

// CacheTheme caches theme data
func CacheTheme(shopID string, theme interface{}) error {
	cache := GetCacheService()
	return cache.SetJSON(ThemeCacheKey(shopID), theme, CacheMedium)
}

// GetCachedTheme retrieves cached theme data
func GetCachedTheme(shopID string, target interface{}) bool {
	cache := GetCacheService()
	return cache.GetJSON(ThemeCacheKey(shopID), target)
}

// InvalidateThemeCache removes theme from cache
func InvalidateThemeCache(shopID string) {
	cache := GetCacheService()
	cache.Delete(ThemeCacheKey(shopID))
	cache.Delete(CSSCacheKey(shopID)) // Also invalidate CSS cache
}

// CacheCSS caches compiled CSS
func CacheCSS(shopID string, css string) {
	cache := GetCacheService()
	cache.Set(CSSCacheKey(shopID), css, CacheLong)
}

// GetCachedCSS retrieves cached CSS
func GetCachedCSS(shopID string) (string, bool) {
	cache := GetCacheService()
	value, exists := cache.Get(CSSCacheKey(shopID))
	if !exists {
		return "", false
	}
	
	css, ok := value.(string)
	return css, ok
}

// CacheShop caches shop data
func CacheShop(shopSlug string, shop interface{}) error {
	cache := GetCacheService()
	return cache.SetJSON(ShopCacheKey(shopSlug), shop, CacheMedium)
}

// GetCachedShop retrieves cached shop data
func GetCachedShop(shopSlug string, target interface{}) bool {
	cache := GetCacheService()
	return cache.GetJSON(ShopCacheKey(shopSlug), target)
}

// InvalidateShopCache removes shop from cache
func InvalidateShopCache(shopSlug string) {
	cache := GetCacheService()
	cache.Delete(ShopCacheKey(shopSlug))
}

// Performance monitoring
type PerformanceMetrics struct {
	CacheHits   int64
	CacheMisses int64
	mutex       sync.RWMutex
}

var metrics = &PerformanceMetrics{}

// RecordCacheHit increments cache hit counter
func RecordCacheHit() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.CacheHits++
}

// RecordCacheMiss increments cache miss counter
func RecordCacheMiss() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.CacheMisses++
}

// GetCacheMetrics returns cache performance metrics
func GetCacheMetrics() map[string]interface{} {
	metrics.mutex.RLock()
	defer metrics.mutex.RUnlock()
	
	total := metrics.CacheHits + metrics.CacheMisses
	hitRate := float64(0)
	if total > 0 {
		hitRate = float64(metrics.CacheHits) / float64(total) * 100
	}
	
	return map[string]interface{}{
		"hits":    metrics.CacheHits,
		"misses":  metrics.CacheMisses,
		"total":   total,
		"hitRate": hitRate,
	}
}
