<template>
<section class="comment-list" v-if="list.length > 0">
  <div class="comment-list-title">文章评论</div>

  <div class="group" v-for="group,index in list" :key="index">
    <div class="comment-item" v-for="comment in group" :key="comment.id">
      <div class="comment-prefix">
        <div class="commenter-avatar">
          <div class="avatar" v-if="comment.creator" :style="{ backgroundImage: 'url('+comment.creator.avatar_url+')'}"></div>
        </div>
      </div>
      <div class="comment-content">
        <div class="content-top">
          <div class="commentor-name">
            <template v-if="comment.creator">{{ comment.creator.name }}</template>
          </div>

          <div class="created-date">{{ semanticFromDateString(comment.created_at, "YYYY-MM-DD") }}</div>
        </div>

        <div class="content-main">{{comment.content}}</div>
      </div>
    </div>
  </div>

  <section class="pagination" v-if="showMore">
    <button @click="getMore">查看更多评论</button>
  </section>
</section>
</template>

<script lang="ts" setup>
const { ARTICLE_API_PRIFIX } = useRuntimeConfig()
const { semanticFromDateString } = useDay()
const { catchMessage } = useMessage()

const props = defineProps({
  stamp: { type: String },
  articleId: { type: String, default: "" }
})

// 评论
const list = ref<any[]>([])
const pagination = ref<any>({})
const showMore = computed(() => Number(pagination.value.page) < Number(pagination.value.page_total))

function getArticleComments(page?: string) {
  $fetch(`${ARTICLE_API_PRIFIX}/comments/list`, { params: {
    article_id: props.articleId,
    preload_creator: 1,
    page: page || "1"
  }}).then((res: any) => {
    if (res.comments.length > 0) {
      res.pagination.page === "1" ? list.value = [res.comments] : list.value.push(res.comments)
    }
    pagination.value = res.pagination
  }).catch(catchMessage)
}

function getMore() {
  const next = Number(pagination.value.page) + 1
  getArticleComments(next.toString())
}

onMounted(() => {
  getArticleComments()
})

watch(() => props.stamp, () => getArticleComments())
</script>

<style lang="scss" scoped>
.comment-list-title {
  padding: 24px 16px;
  font-weight: bold;
  font-size: 1.41rem;
}
.comment-item {
  display: flex;
  margin-bottom: 32px;

  .comment-prefix {
    padding: 0 16px;
    .commenter-avatar, .avatar {
      width: 40px;
      height: 40px;
      border-radius: 50%;
      overflow: hidden;
      background-position: center;
      background-size: cover;
      background-repeat: no-repeat;
    }

    .commenter-avatar {
      position: relative;
    }

    .avatar {
      position: absolute;
    }
  }

  .comment-content {
    flex: 1 1 auto;
    padding-right: 16px;

    .content-top {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding-bottom: 4px;
    }

    .commentor-name {
      color: #000000;
      font-size: 18px;
    }

    .created-date {
      color: #515767;
      font-size: 0.8rem;
    }

    .content-main {
      white-space: pre-line;
      word-wrap: break-word;
      word-break: break-all;
      color: #888888;
    }
  }
}

.pagination {
  margin: auto;
  max-width: 480px;

  button {
    width: 100%;
    border-radius: 500px;
    transition: background 0.2s;
  }
}
</style>
