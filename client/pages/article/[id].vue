<template>
<div class="page">
  <PageHeader></PageHeader>

  <div class="page-body">
    <article class="screen-main article" v-if="detail">
      <div class="article-title">{{ detail.title }}</div>
      <div class="article-other">
        <span>{{ formatFromDateString(detail.created_at, "YYYY-MM-DD hh:mm:ss") }}</span>
      </div>
      <div class="article-content">
        <Markdown :content="detail.content"></Markdown>
      </div>
    </article>

    <div class="empty" v-if="!detail">
      <img class="empty-img" src="@/assets/images/empty.jpg">
      <div class="empty-msg">没有找到想要的资源</div>
    </div>

    <section class="article-comment">
      <CommentCreator @published="onPublished"></CommentCreator>

      <CommentList :stamp="listStamp" :article-id="artilceId"></CommentList>
    </section>
  </div>

  <PageFooter></PageFooter>
</div>
</template>

<script lang="ts" setup>
const { ARTICLE_API_PRIFIX } = useRuntimeConfig()
const { formatFromDateString } = useDay()
const route = useRoute()
const artilceId = computed(() => (route.params.id||"").toString())

const { data }: any = await useAsyncData(() => $fetch(`${ARTICLE_API_PRIFIX}/article/${artilceId.value}`))

const detail = ref(data.value ? data.value.article : null)
if (detail.value) {
  useHead({meta: [
    { name: "keyword", content: detail.value.keyword },
    { name: "description", content: detail.value.description },
  ]})
}


const listStamp = ref<string>(new Date().toString())
function onPublished() {
  listStamp.value = new Date().toString()
}
</script>

<style lang="scss">
.page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;

  .page-body {
    flex: 1 1 auto;
    padding-bottom: 48px;
  }

  .screen-main {
    max-width: 1024px;
    margin: 0 auto;
    padding: 0 16px;
    box-sizing: border-box;
  }
}

.empty {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  .empty-img {
    max-width: 320px;
  }

  .empty-msg {
    color: #888888;
  }
}

article.article {
  .article-title {
    line-height: 3;
    font-size: 32px;
    font-weight: bold;
  }

  .article-other {
    color: #888888;
  }
  .article-content {
    padding: 16px 0 64px;
  }
}

.article-comment {
  max-width: 1024px;
  margin: 0 auto;
}
</style>