<template>
<div class="page">
  <PageHeader :class="{'_scrolling': isScrollToTop}"></PageHeader>

  <div class="page-body">
    <ul class="screen-main">
      <template v-for="group,key in list" :key="key">
        <li :class="['list-item', { '_right': index % 2 === 1 } ]" v-for="item,index in group" :key="item.id">
          <NuxtLink class="content" :to="'/article/'+item.id" target="_blank">
            <div class="item-main">
              <div class="item-title">{{ item.title }}</div>
              <div class="item-desc">{{ item.description }}</div>
              <div class="other">
                <span class="item-date">{{ formatFromDateString(item.created_at, "YYYY-MM-DD hh:mm:ss") }}</span>
              </div>
            </div>
            <div class="item-cover">
              <RatioRect :ratio="5/3" :style="computedCoverStyle(item)"></RatioRect>
            </div>
          </NuxtLink>
        </li>
      </template>
    </ul>

    <div class="empty" v-if="list.length === 0">
      <img class="empty-img" src="@/assets/images/empty.jpg">
      <div class="empty-msg">没有找到想要的资源</div>
    </div>

    <section class="pagination" v-if="showMore">
      <button @click="getMore">阅读更多</button>
    </section>
  </div>

  <PageFooter></PageFooter>
</div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue"
const { formatFromDateString } = useDay()
const { catchMessage } = useMessage()
const route = useRoute()
const keyword = computed(() => route.query.search || "")

const { ARTICLE_API_PRIFIX } = useRuntimeConfig()
const listParams = computed(() => ({ keyword: keyword.value, status: "1" }))

const { data }: any = await useAsyncData(() => $fetch(`${ARTICLE_API_PRIFIX}/articles/list`, { params: listParams.value }))

const list = computed(() => data.value.articles.length ? [data.value.articles] : [])
const pagination = ref(data.value.pagination)
const showMore = computed(() => Number(pagination.value.page < pagination.value.page_total))

function onScroll () {
  const value = (window.document.documentElement.scrollTop || window.document.body.scrollTop) !== 0
  if (value !== isScrollToTop.value) {
    isScrollToTop.value = value
  }
}

const isScrollToTop = ref(false)
onMounted(() => {
  window.addEventListener("scroll", onScroll, false)
})

onBeforeUnmount(() => {
  window.removeEventListener("scroll", onScroll, false)
})

function computedCoverStyle(item: any) {
  return {
    backgroundImage: `url(${item.cover_url})`,
    backgroundSize: "cover",
    backgroundPosition: "center",
    borderRadius: "4px"
  }
}

function getMore() {
  $fetch(`${ARTICLE_API_PRIFIX}/articles/list`, { params: {
    ...listParams.value,
    page: Number(pagination.value.page) + 1
  }}).then((res: any) => {
    list.value.push(res.articles)
    pagination.value = res.pagination
  }).catch(catchMessage)
}
</script>

<style lang="scss" scoped>
$--space: 24px;

.page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;

  .page-body {
    flex: 1 1 1px;
    color: #333333;
  }

  .screen-main {
    max-width: 1024px;
    margin: 0 auto;
    padding: 0 8px;
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

.pagination {
  margin: auto;
  max-width: 480px;
  padding-bottom: 32px;

  button {
    width: 100%;
    border-radius: 500px;
    transition: background 0.2s;
  }
}

._scrolling {
  box-shadow: 0 1px 10px rgba(0, 0, 0, .2);
}

.list-item {
  overflow: hidden;

  &::after {
    content: " ";
    display: block;
    margin: $--space*2 auto;
  }

  &:first-child::before {
    content: " ";
    display: block;
    margin-top: $--space*2;
  }

  &:not(:last-child)::after {
    width: 90%;
    height: 1px;
    background-color: #d5d5d5;
  }

  &._right .content {
    flex-direction: row-reverse;
  }
  .content {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  .item-main {
    flex: 1 1 auto;
    padding: 0 $--space*0.5;
  }

  .item-cover {
    flex: 0 0 40%;
    padding: $--space*0.5;

    img {
      width: 100%;
    }
  }

  .item-title, .item-desc {
    white-space: pre-line;
    word-wrap: break-word;
    word-break: break-word;
  }
  .item-title {
    font-weight: bold;
    font-size: 24px;
  }

  .item-desc {
    padding: 16px 0;
    color: #888888;
  }

  .other {
    color: #888888;
  }
}

@media screen and (max-width: 600px) {
  .list-item {
    font-size: 12px;

    &::after {
      margin: $--space auto;
    }

    .content {
      position: relative;
    }

    .item-title {
      font-size: 16px;
    }

    .item-cover {
      width: 100px;
      position: absolute;
      right: 0;
      top: 50%;
      transform: translateY(-50%);
    }

    .item-desc {
      padding-right: 90px;
    }
  }
}
</style>