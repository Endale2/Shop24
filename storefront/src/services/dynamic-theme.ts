// Dynamic theme service for theme-driven storefront
import { ref, reactive, watch } from 'vue'
import type { Ref } from 'vue'

// =============== Types ===============

export interface DynamicTheme {
  id: string
  name: string
  version: string
  colors: ThemeColors
  fonts: ThemeFonts
  layout: ThemeLayout
  customCSS: string
  seo: ThemeSEO
  mobile: ThemeMobile
  dynamicCSS: string
  performance?: ThemePerformance
  lastModified: string
  cacheKey?: string
}

export interface ThemeColors {
  primary: string
  secondary: string
  background: string
  heading: string
  bodyText: string
  border?: string
  accent?: string
}

export interface ThemeFonts {
  heading: string
  body: string
  accent?: string
}

export interface ThemeLayout {
  containerWidth: string
  headerStyle: string
  footerStyle: string
  gridColumns: string
  sidebarPosition: string
  borderStyle: string
  spacing?: string
  navStyle?: string
  buttonStyle?: string
}

export interface ThemeSEO {
  metaTitle?: string
  metaDescription?: string
  keywords?: string
}

export interface ThemeMobile {
  enabled: string
  responsive: string
  touchOptimized: string
}

export interface ThemePerformance {
  loadTime?: number
  performanceScore?: number
  conversions?: number
  views?: number
}

export interface StorefrontConfig {
  shop: ShopInfo
  theme: DynamicTheme
  layout: LayoutConfig
  components: ComponentConfig
  navigation: NavigationConfig
  footer: FooterConfig
  seo: SEOConfig
  analytics: AnalyticsInfo
  scalable: boolean
  dynamic: boolean
}

export interface ShopInfo {
  id: string
  name: string
  slug: string
  description: string
  image: string
  banner: string
  email: string
  phone: string
  address: string
  currency: string
  status: string
  isVerified: boolean
  rating: number
  reviewCount: number
}

export interface LayoutConfig {
  containerWidth: string
  headerStyle: string
  footerStyle: string
  sidebarEnabled: boolean
  gridColumns: string
  spacing: string
  borderStyle: string
}

export interface ComponentConfig {
  productCard: {
    style: string
    showRating: boolean
    showBadges: boolean
    imageRatio: string
    hoverEffect: string
  }
  buttons: {
    style: string
    size: string
    animation: string
  }
  navigation: {
    style: string
    position: string
    transparent: boolean
  }
}

export interface NavigationConfig {
  items: NavigationItem[]
  style: string
  showIcons: boolean
  sticky: boolean
}

export interface NavigationItem {
  label: string
  path: string
  icon: string
}

export interface FooterConfig {
  style: string
  showSocial: boolean
  showNewsletter: boolean
  columns: number
}

export interface SEOConfig {
  title: string
  description: string
  keywords: string
  ogImage: string
  themeColor: string
}

export interface AnalyticsInfo {
  customerCount: number
  productCount: number
  revenue: number
  shopAge: string
}

// =============== Global State ===============

const currentStorefront = ref<StorefrontConfig | null>(null)
const currentTheme = ref<DynamicTheme | null>(null)
const isLoading = ref(false)
const error = ref<string | null>(null)

// Theme cache for performance
const themeCache = new Map<string, { data: DynamicTheme; timestamp: number }>()
const CACHE_DURATION = 5 * 60 * 1000 // 5 minutes

// =============== API Functions ===============

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080'

// Fetch complete storefront configuration
export async function fetchStorefrontConfig(shopSlug: string): Promise<StorefrontConfig> {
  isLoading.value = true
  error.value = null
  
  try {
    const response = await fetch(`${API_BASE}/shops/${shopSlug}/storefront`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch storefront config: ${response.statusText}`)
    }
    
    const config: StorefrontConfig = await response.json()
    
    // Update global state
    currentStorefront.value = config
    currentTheme.value = config.theme
    
    // Apply dynamic theme
    await applyDynamicTheme(config.theme)
    
    // Update SEO
    updateSEO(config.seo)
    
    isLoading.value = false
    return config
    
  } catch (err) {
    const errorMessage = err instanceof Error ? err.message : 'Unknown error'
    error.value = errorMessage
    isLoading.value = false
    throw err
  }
}

// Fetch just theme data for real-time updates
export async function fetchThemeData(shopSlug: string, useCache = true): Promise<DynamicTheme> {
  // Check cache first
  if (useCache) {
    const cached = themeCache.get(shopSlug)
    if (cached && (Date.now() - cached.timestamp < CACHE_DURATION)) {
      return cached.data
    }
  }
  
  try {
    const response = await fetch(`${API_BASE}/shops/${shopSlug}/theme`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch theme data: ${response.statusText}`)
    }
    
    const theme: DynamicTheme = await response.json()
    
    // Update cache
    themeCache.set(shopSlug, { data: theme, timestamp: Date.now() })
    
    // Update global state
    currentTheme.value = theme
    
    // Apply theme updates
    await applyDynamicTheme(theme)
    
    return theme
    
  } catch (err) {
    console.error('Error fetching theme data:', err)
    throw err
  }
}

// =============== Theme Application ===============

// Apply dynamic theme to the DOM
export async function applyDynamicTheme(theme: DynamicTheme) {
  // Remove existing theme style
  const existingStyle = document.getElementById('dynamic-theme-style')
  if (existingStyle) {
    existingStyle.remove()
  }
  
  // Create new style element
  const styleElement = document.createElement('style')
  styleElement.id = 'dynamic-theme-style'
  styleElement.textContent = theme.dynamicCSS + (theme.customCSS || '')
  
  // Add to document head
  document.head.appendChild(styleElement)
  
  // Update CSS custom properties for immediate updates
  const root = document.documentElement
  
  // Colors
  if (theme.colors) {
    root.style.setProperty('--color-primary', theme.colors.primary)
    root.style.setProperty('--color-secondary', theme.colors.secondary)
    root.style.setProperty('--color-background', theme.colors.background)
    root.style.setProperty('--color-heading', theme.colors.heading)
    root.style.setProperty('--color-body-text', theme.colors.bodyText)
    root.style.setProperty('--color-border', theme.colors.border || '#E5E7EB')
  }
  
  // Fonts
  if (theme.fonts) {
    root.style.setProperty('--font-heading', `'${theme.fonts.heading}', sans-serif`)
    root.style.setProperty('--font-body', `'${theme.fonts.body}', sans-serif`)
  }
  
  // Layout
  if (theme.layout) {
    root.style.setProperty('--grid-columns', theme.layout.gridColumns)
    root.style.setProperty('--container-width', getContainerWidth(theme.layout.containerWidth))
    root.style.setProperty('--border-radius', getBorderRadius(theme.layout.borderStyle))
  }
  
  console.log('✅ Dynamic theme applied:', theme.name, 'v' + theme.version)
}

// Update SEO meta tags
function updateSEO(seo: SEOConfig) {
  // Update title
  if (seo.title) {
    document.title = seo.title
  }
  
  // Update meta description
  updateMetaTag('description', seo.description)
  
  // Update meta keywords
  if (seo.keywords) {
    updateMetaTag('keywords', seo.keywords)
  }
  
  // Update theme color
  updateMetaTag('theme-color', seo.themeColor)
  
  // Update Open Graph
  updateMetaTag('og:title', seo.title, 'property')
  updateMetaTag('og:description', seo.description, 'property')
  updateMetaTag('og:image', seo.ogImage, 'property')
}

function updateMetaTag(name: string, content: string, attribute = 'name') {
  if (!content) return
  
  let meta = document.querySelector(`meta[${attribute}="${name}"]`)
  if (!meta) {
    meta = document.createElement('meta')
    meta.setAttribute(attribute, name)
    document.head.appendChild(meta)
  }
  meta.setAttribute('content', content)
}

// =============== Real-time Theme Updates ===============

// Watch for theme changes and auto-update
let themeWatchInterval: number | null = null

export function startThemeWatch(shopSlug: string, intervalMs = 30000) {
  if (themeWatchInterval) {
    clearInterval(themeWatchInterval)
  }
  
  themeWatchInterval = setInterval(async () => {
    try {
      const latestTheme = await fetchThemeData(shopSlug, false)
      
      // Check if theme has been updated
      if (currentTheme.value && latestTheme.version !== currentTheme.value.version) {
        console.log('🔄 Theme updated, applying changes...')
        await applyDynamicTheme(latestTheme)
        
        // Emit custom event for components to react
        window.dispatchEvent(new CustomEvent('theme-updated', { 
          detail: { theme: latestTheme, previousVersion: currentTheme.value.version } 
        }))
      }
    } catch (err) {
      console.warn('Theme watch error:', err)
    }
  }, intervalMs)
}

export function stopThemeWatch() {
  if (themeWatchInterval) {
    clearInterval(themeWatchInterval)
    themeWatchInterval = null
  }
}

// =============== Utility Functions ===============

export function getContainerWidth(type: string): string {
  switch (type) {
    case 'full': return '100%'
    case 'wide': return '1400px'
    case 'boxed': return '1200px'
    default: return '1200px'
  }
}

export function getBorderRadius(style: string): string {
  switch (style) {
    case 'sharp': return '0px'
    case 'rounded': return '8px'
    case 'pill': return '50px'
    default: return '8px'
  }
}

export function getGridColumns(count: string): string {
  return `repeat(${count}, 1fr)`
}

// Get current theme reactive reference
export function useCurrentTheme(): Ref<DynamicTheme | null> {
  return currentTheme
}

// Get current storefront reactive reference
export function useCurrentStorefront(): Ref<StorefrontConfig | null> {
  return currentStorefront
}

// Get loading state
export function useThemeLoading() {
  return { isLoading, error }
}

// Clear theme cache
export function clearThemeCache() {
  themeCache.clear()
}

// Get theme performance score based on metrics
export function getThemePerformanceScore(theme: DynamicTheme): number {
  if (!theme.performance) return 85 // Default score
  
  let score = 100
  
  // Deduct for slow load times
  if (theme.performance.loadTime && theme.performance.loadTime > 3) {
    score -= 20
  }
  if (theme.performance.loadTime && theme.performance.loadTime > 5) {
    score -= 30
  }
  
  // Bonus for good conversion rate
  if (theme.performance.conversions && theme.performance.views) {
    const conversionRate = (theme.performance.conversions / theme.performance.views) * 100
    if (conversionRate > 5) score += 5
    if (conversionRate > 10) score += 10
  }
  
  return Math.max(0, Math.min(100, score))
}

// =============== Component Helpers ===============

// Generate dynamic classes based on theme
export function getDynamicClasses(theme: DynamicTheme | null, component: string): string[] {
  if (!theme) return []
  
  const classes: string[] = []
  
  switch (component) {
    case 'container':
      classes.push('dynamic-container')
      break
    case 'grid':
      classes.push('dynamic-grid')
      break
    case 'card':
      classes.push('product-card')
      break
    case 'button-primary':
      classes.push('btn-primary')
      break
    case 'button-secondary':
      classes.push('btn-secondary')
      break
    case 'header':
      classes.push('dynamic-header')
      break
  }
  
  return classes
}

export default {
  fetchStorefrontConfig,
  fetchThemeData,
  applyDynamicTheme,
  startThemeWatch,
  stopThemeWatch,
  useCurrentTheme,
  useCurrentStorefront,
  useThemeLoading,
  clearThemeCache,
  getThemePerformanceScore,
  getDynamicClasses,
}
