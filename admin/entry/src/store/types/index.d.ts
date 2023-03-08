export interface User {
  id?: number
  name?: string
}

export interface GlobalState {
  token: string
  user: User|null
}

export interface ApplicationState {
  navsLoaded: boolean
  navs: any[]
}

export interface RootState {
  global: GlobalState
  application: ApplicationState
}