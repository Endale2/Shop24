import api, { getShopUrl } from './api'

export interface Collection {
  handle: string
  id: string
  title: string
  image: string
}

export interface CollectionDetail extends Collection {
  description: string
  created_at: string
  filters: any
  products: Array<{
    id: string
    slug: string
    name: string
    main_image: string
    description: string
    price?: number
    display_price?: number
    variant_count?: number
    discounts?: Array<{
      id: string
      name: string
      type: string
      value: number
      active: boolean
    }>
    variants?: Array<any>
  }>
}

export function fetchCollections(shopSlug: string): Promise<Collection[]> {
  return api.get(getShopUrl(shopSlug, '/collections')).then(r => r.data)
}

export function fetchCollectionDetail(
  shopSlug: string,
  handle: string
): Promise<CollectionDetail> {
  return api
    .get(getShopUrl(shopSlug, `/collections/${encodeURIComponent(handle)}`))
    .then(r => r.data)
}
