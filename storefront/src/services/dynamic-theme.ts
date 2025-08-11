// Dynamic theme service for theme-driven storefront
import { ref} from 'vue'
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
  dynamicCSS?: string  // Backend compiled CSS with variables
  seo: ThemeSEO
  mobile: ThemeMobile
  performance?: ThemePerformance
  lastModified: string
  cacheKey?: string
  // Enhanced styling options
  gradients?: ThemeGradients
  shadows?: ThemeShadows
  animations?: ThemeAnimations
  spacing?: ThemeSpacing
  components?: ThemeComponents
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
  menuStyle?: string
  stickyHeader?: string
}

export interface ThemeGradients {
  hero: string
  button: string
  card: string
  background: string
}

export interface ThemeShadows {
  card: string
  button: string
  image: string
  hover: string
}

export interface ThemeAnimations {
  transition: string
  hover: string
  loading: string
  fadeIn: string
}

export interface ThemeSpacing {
  xs: string
  sm: string
  md: string
  lg: string
  xl: string
  '2xl': string
}

export interface ThemeComponents {
  productCard?: {
    style: string
    hoverEffect: string
    imageRatio: string
    showBadges: boolean
    showRating: boolean
  }
  buttons?: {
    style: string
    size: string
    animation: string
  }
  navigation?: {
    style: string
    position: string
    transparent: boolean
    sticky: boolean
  }
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

import api, { getShopUrl } from './api'

// Fetch complete storefront configuration
export async function fetchStorefrontConfig(shopSlug: string): Promise<StorefrontConfig> {
  isLoading.value = true
  error.value = null

  try {
    const response = await api.get(getShopUrl(shopSlug, '/storefront'))

    const config: StorefrontConfig = response.data

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
    const response = await api.get(getShopUrl(shopSlug, '/theme'))

    const theme: DynamicTheme = response.data
    
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

  // Combine dynamicCSS (backend compiled) and customCSS (user custom)
  const combinedCSS = (theme.dynamicCSS || '') + (theme.customCSS || '')
  styleElement.textContent = combinedCSS

  // Add to document head
  document.head.appendChild(styleElement)

  console.log('📝 Applied CSS:', combinedCSS.length, 'characters')
  
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

  // Enhanced styling options
  if (theme.gradients) {
    root.style.setProperty('--gradient-hero', theme.gradients.hero)
    root.style.setProperty('--gradient-button', theme.gradients.button)
    root.style.setProperty('--gradient-card', theme.gradients.card)
    root.style.setProperty('--gradient-background', theme.gradients.background)
  }

  if (theme.shadows) {
    root.style.setProperty('--shadow-card', theme.shadows.card)
    root.style.setProperty('--shadow-button', theme.shadows.button)
    root.style.setProperty('--shadow-image', theme.shadows.image)
    root.style.setProperty('--shadow-hover', theme.shadows.hover)
  }

  if (theme.animations) {
    root.style.setProperty('--transition-default', theme.animations.transition)
    root.style.setProperty('--transition-hover', theme.animations.hover)
    root.style.setProperty('--transition-loading', theme.animations.loading)
    root.style.setProperty('--transition-fade-in', theme.animations.fadeIn)
  }

  if (theme.spacing) {
    root.style.setProperty('--spacing-xs', theme.spacing.xs)
    root.style.setProperty('--spacing-sm', theme.spacing.sm)
    root.style.setProperty('--spacing-md', theme.spacing.md)
    root.style.setProperty('--spacing-lg', theme.spacing.lg)
    root.style.setProperty('--spacing-xl', theme.spacing.xl)
    root.style.setProperty('--spacing-2xl', theme.spacing['2xl'])
  }

  // Component-specific styling
  if (theme.components?.productCard) {
    root.style.setProperty('--product-card-style', theme.components.productCard.style)
    root.style.setProperty('--product-card-hover', theme.components.productCard.hoverEffect)
    root.style.setProperty('--product-card-ratio', theme.components.productCard.imageRatio)
  }

  if (theme.components?.buttons) {
    root.style.setProperty('--button-style', theme.components.buttons.style)
    root.style.setProperty('--button-size', theme.components.buttons.size)
    root.style.setProperty('--button-animation', theme.components.buttons.animation)
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
