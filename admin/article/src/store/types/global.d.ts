export interface User {
  id?: string
  name?: string
}

export interface GlobalState {
  token: string
  user: User|null
}