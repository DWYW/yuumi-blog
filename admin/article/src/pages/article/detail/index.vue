<template>
<div class="page" v-loading="loading">
  <div class="page-header">
  </div>
  <div class="page-body">
    <div class="content">
      <section class="content-main">
        <YuumiScrollbar>
          <div class="placeholder" v-t="'ARTICLE.PLACEHOLDER_MSG'" v-if="!content"></div>
          <div ref="mdElement" id="md" contenteditable="true" @input="onChange"></div>
        </YuumiScrollbar>
      </section>
      <section class="content-aside flex _col">
        <div class="content form">
          <div class="flex _mc">
            <div class="prefix" v-t="'ATTR.TITLE'"></div>
            <div class="content">
              <YuumiInput v-model="title"></YuumiInput>
            </div>
          </div>
          <div class="flex _mc">
            <div class="prefix" v-t="'ATTR.DESCRIPTION'"></div>
            <div class="content">
              <YuumiInput v-model="description"></YuumiInput>
            </div>
          </div>
          <div class="flex _mc">
            <div class="prefix" v-t="'ATTR.KEYWORD'"></div>
            <div class="content">
              <YuumiInput v-model="keyword" :placeholder="$t('ARTICLE.KEYWORD_PLACEHOLDER_MSG')"></YuumiInput>
            </div>
          </div>
          <div class="flex _mc">
            <div class="prefix" v-t="'ATTR.COVER_URL'"></div>
            <div class="content">
              <YuumiInput v-model="coverUrl"></YuumiInput>
            </div>
          </div>
          <div class="flex">
            <div class="prefix"></div>
            <div class="content">
              <img :src="coverUrl" alt="" style="width: 100%;">
            </div>
          </div>
        </div>

        <div class="suffix submit">
          <YuumiButton size="lg" v-t="'HANDLE.CANCEL'" @click="onCancel"></YuumiButton>
          <YuumiButton size="lg" theme="primary" v-t="'HANDLE.SAVE'" @click="onSave"></YuumiButton>
        </div>
      </section>
    </div>
  </div>
  <div class="page-footer"></div>
</div>
</template>

<script lang="ts" setup>
import { ArticleAPI } from "@/api"
import { $helper } from "@/common"
import { computed, onBeforeUnmount, onMounted, ref } from "vue"
import { useI18n } from "vue-i18n"
import { useRoute, useRouter } from "vue-router"
import { createMessage } from "yuumi-ui-vue"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const articleId = computed(() => $helper.getValueByPath(route, "params.id", "0"))

const loading = ref(false)
const title = ref("")
const description = ref("")
const coverUrl = ref("")
const keyword = ref("")
const mdElement = ref()
const content = ref("")

onMounted(() => {
  if (articleId.value !== "0") {
    getInfo()
  }

  document.addEventListener("keydown", bindShortcutKeys, false)
})

onBeforeUnmount(() => {
  document.removeEventListener("keydown", bindShortcutKeys, false)
})

function bindShortcutKeys(e: any) {
  console.log(e.ctrlKey, e.key)
  if (e.ctrlKey && e.key === "s") {
    onSave()
    e.preventDefault()
    e.stopPropagation()
  }
}

function getInfo() {
  if (loading.value) return
  loading.value = true

  ArticleAPI.getInfoWithId(articleId.value).finally(() => {
    loading.value = false
  }).then(({ article }) => {
    title.value = article.title
    description.value = article.description
    coverUrl.value = article.cover_url
    content.value = article.content
    keyword.value = article.keyword

    if (article.content) {
      mdElement.value.innerText = article.content
    }
  })
}

function onChange(e: Event) {
  const { target }: any = e
  content.value = target.innerText
}

function onCancel() {
  router.back()
}

function onSave() {
  if (!title.value) return createMessage({ message: t("ARTICLE.TITLE_REQUIRED_MSG"), theme: "warn" })
  if (!description.value) return createMessage({ message: t("ARTICLE.DESCRIPTION_REQUIRED_MSG"), theme: "warn" })
  if (!keyword.value) return createMessage({ message: t("ARTICLE.KEYWORD_REQUIRED_MSG"), theme: "warn" })
  if (!coverUrl.value) return createMessage({ message: t("ARTICLE.COVER_URL_REQUIRED_MSG"), theme: "warn" })

  const v:{[key: string]: string} = {
    title: title.value,
    description: description.value,
    keyword: keyword.value,
    cover_url: coverUrl.value,
    content: content.value
  }
  articleId.value === "0" ? createArticle(v) : updateArticle(v)
}

function createArticle(v: any) {
  if (loading.value) return
  loading.value = true

  return ArticleAPI.create(v).finally(() => {
    loading.value = false
  }).then(({ article }) => {
    createMessage({ message: t("MSG.CREATED_SUCCESS"), theme: "success" })
    router.replace(`/system/article/detail/${article.id}`)
  })
}

function updateArticle(v: any) {
  if (loading.value) return
  loading.value = true

  return ArticleAPI.updateWithId(articleId.value, v).finally(() => {
    loading.value = false
  }).then(() => {
    createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
  })
}
</script>

<style lang="scss" scoped>
.page-body > .content {
  display: flex;

  .content-main {
    flex: 1 1 66.66%;

    .placeholder {
      color: map-get($--color, "placeholder");
      position: absolute;
      left: 0;
      top: 0;
    }

    #md {
      position: relative;
      width: 100%;
      min-height: 100%;
      outline: none;
    }
  }

  .content-aside {
    border-left: 1px solid  map-get($--color, "border");
    flex: 1 1 33.33%;
    min-width: 320px;

  }
}
.form > .flex {
  margin-bottom: map-get($--space, "xm");
  > .prefix {
    width: 80px;
    text-align: right;
  }

  > .content {
    display: flex;

    :deep(.yuumi-input) {
      width: 100%;
    }
  }
}

.submit {
  display: flex;
  justify-content: center;
  margin-top: map-get($--space, "lg");

  :deep(.yuumi-button) {
    min-width: 100px;
  }

  > :not(:last-child) {
    margin-right: map-get($--space, "sm");
  }
}
</style>
