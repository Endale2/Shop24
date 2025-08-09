import api, { getShopUrl, getShopSlugFromSubdomain } from './api'

export interface Shop {
  id: string
  name: string
  slug: string
  image: string
  banner?: string
  description: string
}

export function setCurrentShopSlug(slug: string) {
  localStorage.setItem('currentShopSlug', slug);
}

export function getPersistedShopSlug(): string | null {
  return localStorage.getItem('currentShopSlug');
}

export function setCurrentShopData(shop: Shop) {
  localStorage.setItem('currentShopData', JSON.stringify(shop));
}

export function getPersistedShopData(): Shop | null {
  const data = localStorage.getItem('currentShopData');
  if (!data) return null;
  try {
    return JSON.parse(data);
  } catch {
    return null;
  }
}

export function getCurrentShopSlug(): string | null {
  // Prefer subdomain when available
  const fromSubdomain = getShopSlugFromSubdomain();
  if (fromSubdomain) return fromSubdomain;
  // Fallback to persisted value
  const persisted = getPersistedShopSlug();
  if (persisted) return persisted;
  // Dev convenience: allow an env-provided default when running without subdomain
  const envDefault = (import.meta as any).env?.VITE_DEFAULT_SHOP_SLUG as string | undefined
  if (envDefault) return envDefault;
  return null;
}

export async function fetchShop(shopSlug?: string): Promise<Shop | null> {
  const slug = shopSlug || getShopSlugFromSubdomain() || getCurrentShopSlug();
  if (!slug) return null;
  try {
    const response = await api.get(getShopUrl(slug, ''))
    return response.data
  } catch (error) {
    console.error('Failed to fetch shop:', error)
    return null
  }
}

export function clearPersistedShop() {
  localStorage.removeItem('currentShopSlug');
  localStorage.removeItem('currentShopData');
}
