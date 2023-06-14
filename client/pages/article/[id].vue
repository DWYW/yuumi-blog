<template>
<div class="page">
  <PageHeader :class="{'_scrolling': isScrollToTop}"></PageHeader>

  <div class="page-body">
    <article class="screen-main article" v-if="isReady">
      <div class="article-cover">
        <RatioRect :ratio="5/2" :style="computedCoverStyle()"></RatioRect>
      </div>

      <div class="article-title">{{ detail.title }}</div>
      <div class="article-other">
        <span>{{ formatFromDateString(detail.created_at, "YYYY-MM-DD hh:mm:ss") }}</span>
      </div>
      <div class="article-content">
        <Markdown :content="detail.content"></Markdown>
      </div>
    </article>

    <div class="empty" v-else>
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
  useHead({
    title: detail.value.title,
    meta: [
      { name: "keywords", content: detail.value.keyword },
      { name: "description", content: detail.value.description }
    ]
  })
}

const isReady = computed(() => detail.value && (detail.value.status === "1" || route.query.preview === "1"))

function computedCoverStyle() {
  return {
    backgroundImage: `url(${detail.value.cover_url})`,
    backgroundSize: "cover",
    backgroundPosition: "center"
  }
}

const listStamp = ref<string>(new Date().toString())
function onPublished() {
  listStamp.value = new Date().toString()
}

const isScrollToTop = ref(false)
onMounted(() => {
  globalThis.onscroll = (e) => {
    const value = (globalThis.document.documentElement.scrollTop || globalThis.document.body.scrollTop) !== 0
    if (value !== isScrollToTop.value) {
      isScrollToTop.value = value
    }
  }
})
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


._scrolling {
  box-shadow: 0 1px 10px rgba(0, 0, 0, .2);
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
  .article-cover {
    background-color: #f5f5f5;
  }

  .article-title {
    font-size: 32px;
    font-weight: bold;
    padding: 48px 0;
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