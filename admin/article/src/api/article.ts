import { ArticleServer, RequestOption } from "./instance"

export function create(data: {
  title: string
  description: string
  cover_url: string
  content: string
  keyword: string
}, options?: RequestOption) {
  return ArticleServer.post("/article", data, options)
}

export function updateWithId(id: string, data: {
  title?: string
  description?: string
  cover_url?: string
  content?: string
  keyword?: string
}, options?: RequestOption) {
  return ArticleServer.put(`/article/${id}`, data, options)
}

export function deleteWithId(id: string, options?: RequestOption) {
  return ArticleServer.delete(`/article/${id}`, {}, options)
}

export function getList(params?: { [key: string]: number|string }, options?: RequestOption) {
  return ArticleServer.get("/articles/list", params, options)
}

export function getInfoWithId(id: string, options?: RequestOption) {
  return ArticleServer.get(`/article/${id}`, {}, options)
}
