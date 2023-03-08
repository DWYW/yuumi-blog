export const VISITOR_TOKEN = "VISITOR_TOKEN"
export const VISITOR_DATA = "VISITOR_DATA"
export const VISITOR_SIGNIN_SESSION = "VISITOR_SIGNIN_SESSION"

interface VisitorData {
  id: string
  name: string
}

export const useVisitorStore = defineStore('visitor', () => {
  const visitor = ref<VisitorData>({id: "", name: ""})
  const updateVisitor = (v: VisitorData) => {
    visitor.value = reactive(v)
    localStorage.setItem(VISITOR_DATA, JSON.stringify(v))
  }

  const signinSession = ref<string>("")
  const updateSigninSession = (v: string) => {
    signinSession.value = v
    localStorage.setItem(VISITOR_SIGNIN_SESSION, JSON.stringify(v))
  }

  const token = ref<string>("")
  const updateToken = (v: string) => {
    token.value = v
    localStorage.setItem(VISITOR_TOKEN, JSON.stringify(v))
  }

  const signout = () => {
    visitor.value = {id: "", name: ""}
    signinSession.value = ""
    token.value = ""
    localStorage.removeItem(VISITOR_DATA)
    localStorage.removeItem(VISITOR_SIGNIN_SESSION)
    localStorage.removeItem(VISITOR_TOKEN)
  }

  return {
    visitor,
    updateVisitor,
    signinSession,
    updateSigninSession,
    token,
    updateToken,
    signout
  }
})