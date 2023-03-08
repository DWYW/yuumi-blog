import { defineNuxtPlugin } from '#app'
import { useVisitorStore, VISITOR_DATA, VISITOR_SIGNIN_SESSION, VISITOR_TOKEN } from '~/stores/visitor'

export default defineNuxtPlugin((nuxtApp) => {
  const { updateVisitor, updateSigninSession, updateToken } = useVisitorStore(nuxtApp.$pinia)
  const vistor = localStorage.getItem(VISITOR_DATA)
  if (vistor) {
    updateVisitor(JSON.parse(vistor))
  }

  const signinSession = localStorage.getItem(VISITOR_SIGNIN_SESSION)
  if (signinSession) {
    updateSigninSession(JSON.parse(signinSession))
  }

  const token = localStorage.getItem(VISITOR_TOKEN)
  if (token) {
    updateToken(JSON.parse(token))
  }
})