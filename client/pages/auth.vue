<template>
<div class="auth-page">
  {{ errMsg ? errMsg : "登录中，请稍后......" }}
</div>
</template>

<script lang="ts" setup>
import { useVisitorStore } from "~/stores/visitor"

const { PASSPORT_API_PRIFIX } = useRuntimeConfig()
const { token, signinSession, updateSigninSession, updateToken, updateVisitor } = useVisitorStore()
const route = useRoute()
const errMsg = ref("")

onMounted(() => {
  if (!signinSession) {
    code2session((route.query.code || "").toString())
  } else if (!token) {
    session2token(signinSession)
  } else {
    window.open((route.query.redirect||"").toString(), "_self")
  }
})

function code2session(code: string) {
  if (!code) return
  $fetch(`${PASSPORT_API_PRIFIX}/github/session/with-code?code=${code}&redirect=${encodeURIComponent(location.href)}`).then((res: any) => {
    updateSigninSession(res.session)
    session2token(res.session)
  }).catch((err) => {
    errMsg.value = err.message
  })
}

function session2token(session: string) {
  $fetch(`${PASSPORT_API_PRIFIX}/authorize/visitor/github-session`, {
    method: "POST",
    body: JSON.stringify({session})
  }).then((res: any) => {
    updateToken(res.token)
    updateVisitor({id: res.visitor.id, name: res.visitor.name})
    window.open((route.query.redirect||"").toString(), "_self")
  }).catch((err) => {
    errMsg.value = err.message
  })
}
</script>

<style lang="scss" scoped>
.auth-page {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>