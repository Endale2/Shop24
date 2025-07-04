import { api, shopUrl } from './api'

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
    price: number
    display_price: number
    variant_count: number
  }>
}

export function fetchCollections(): Promise<Collection[]> {
  return api.get(shopUrl('/collections')).then(r => r.data)
}

export function fetchCollectionDetail(
  handle: string
): Promise<CollectionDetail> {
  return api
    .get(shopUrl(`/collections/${encodeURIComponent(handle)}`))
    .then(r => r.data)
}
