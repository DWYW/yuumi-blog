<template>
<section class="comment-creator">
  <div class="creator">
    <textarea v-model="content" rows="4" :placeholder="placeholderMsg" :disabled="!availabled"></textarea>
  </div>
  <div class="commit">
    <template v-if="availabled" >
      <button @click="publish">发表</button>
    </template>
    <template v-else>
      <button  @click="githubAuth">GITHUB 登录</button>
    </template>
  </div>
</section>
</template>

<script lang="ts" setup>
import { useVisitorStore } from "~/stores/visitor"

const emit = defineEmits(["published"])

const route = useRoute()
const { warnMessage, successMessage, catchMessage } = useMessage()
const { GITHUB_CLIENTID, ARTICLE_API_PRIFIX } = useRuntimeConfig()

const visitor = ref<any>(null)
const visitorToken = ref<string>("")
const availabled = ref(false)
const placeholderMsg = ref("")

watch(() => useVisitorStore().token, () => {nextTick(() => {
  const state = useVisitorStore()
  visitor.value = state.visitor
  visitorToken.value = state.token
  availabled.value = state.token != ""
  placeholderMsg.value = availabled.value ? "您已浏览完，分享一下吧 ◡‿◡" : "登录后分享一下吧 ◡‿◡"
})}, { immediate: true })


const content = ref("")

function githubAuth() {
  const target = encodeURIComponent(location.origin + location.pathname)
  const redirect = encodeURIComponent(`${location.origin}/auth?redirect=${target}`)
  const url = `https://github.com/login/oauth/authorize?client_id=${GITHUB_CLIENTID}&scope=public_repo,user&redirect_uri=${redirect}`
  window.open(url, "_self")
}

async function publish() {
  if (!content.value) return warnMessage({ message: "请输入您要分享的内容"})

  $fetch(`${ARTICLE_API_PRIFIX}/comment`, {
    method: "POST",
    body: JSON.stringify({
      content: content.value,
      article_id: route.params.id,
      creator_id: visitor.value.id
    }),
    headers: {
      "Authorization": visitorToken.value
    }
  }).then(() => {
    content.value = ""
    successMessage({message: "分享成功！"})
    emit("published")
  }).catch(catchMessage)
}
</script>

<style lang="scss" scoped>
.comment-creator {
  margin: 16px;
  padding: 10px;
  border: 1px solid #d5d5d5;
  border-radius: 2px;

  .creator {
    textarea {
      padding: 8px;
      background-color: #f5f5f5;
      border: none;
      border-radius: 2px;
      width: 100%;
      font-size: inherit;
      line-height: 1.5;
      resize: none;
    }
  }

  .commit {
    text-align: right;
    padding-top: 4px;

    .signout {
      margin-right: 24px;
      cursor: pointer;
    }
  }
}
</style>
