<template>
<div class="page" v-loading="loading">
  <div class="page-header">
    <div class="content flex">
      <div class="prefix"></div>
      <div class="content"></div>
      <div class="suffix">
        <YuumiButton theme="primary" v-t="'HANDLE.CREATE'" @click="onCreateHandler"></YuumiButton>
      </div>
    </div>
  </div>
  <div class="page-body">
    <div class="content" v-if="list.length">
      <YuumiTable :data="list" :row-class-name="rowClassName">
        <YuumiTableColumn :title="$t('ATTR.NAME')" prop="title"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.COVER')" prop="cover_url" :width="100">
          <template #default="{$value}">
            <div class="item-cover" :style="{backgroundImage: 'url('+$value+')'}"></div>
          </template>
        </YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.DESCRIPTION')" prop="status" :width="200">
          <template #default="{$value}">
            <span :class="{ '_published': $value === '1'}">{{ statusText($value) }}</span>
          </template>
        </YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.DESCRIPTION')" prop="description" :width="200"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.CREATED_DATE')" prop="created_date" :width="160"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.UPDATED_DATE')" prop="updated_date" :width="160"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.HANDLE')" fixed="right" :width="280">
          <template #default="scope">
            <YuumiButton v-t="'HANDLE.PREVIEW'" @click="onPreviewHandler(scope)"></YuumiButton>
            <YuumiButton theme="error" v-t="'HANDLE.DELETE'" @click="onDeleteHandler(scope)"></YuumiButton>
            <YuumiButton v-t="'HANDLE.EDIT'" @click="onUpdateHandler(scope)"></YuumiButton>
            <YuumiButton v-if="articleIsPublished(scope)" theme="error" v-t="'HANDLE.UNPUBLISH'" @click="onUnpublishHandler(scope)"></YuumiButton>
            <YuumiButton v-else v-t="'HANDLE.PUBLISH'" @click="onPublishHandler(scope)"></YuumiButton>
          </template>
        </YuumiTableColumn>
      </YuumiTable>
    </div>

    <YuumiEmpty :description="$t('MSG.EMPTY')"></YuumiEmpty>
  </div>
  <div class="page-footer">
    <div class="content">
      <YuumiPagination v-if="list.length"
        :page="pagination.page"
        :page-total="pagination.pageTotal"
        :total="pagination.total"
        @change="onPaginationChange"
      ></YuumiPagination>
    </div>
  </div>
</div>
</template>

<script lang="ts" setup>
import { createAlert, createMessage } from "yuumi-ui-vue"
import { onMounted, Ref, ref } from "vue"
import { ArticleAPI } from "@/api"
import { $day, $helper } from "@/common"
import { useTableHelper } from "@/components/helper/table"
import { useI18n } from "vue-i18n"
import { useRouter } from "vue-router"

const router = useRouter()
const { t } = useI18n()

const loading: Ref<boolean> = ref(false)
const { list, pagination, tableReload, onPaginationChange } = useTableHelper(() => {
  if (loading.value) return
  loading.value = true

  return ArticleAPI.getList({
    page: pagination.value.page
  }).finally(() => {
    loading.value = false
  }).then(({ articles, pagination }) => {
    articles = articles.map((item: any) => {
      item.created_date = $day.formatTimeFromDateString(item.created_at)
      item.updated_date = $day.formatTimeFromDateString(item.updated_at)
      return item
    })
    return { list: articles, pagination }
  })
})

const selected: Ref<any> = ref({})
function rowClassName({ rowIndex }: any) {
  const row = list.value[rowIndex]

  if (row.id === selected.value.id) {
    return "_selected"
  }
}

onMounted(() => {
  tableReload()
})

// 删除
function onDeleteHandler({ $attrs }: any) {
  selected.value = list.value[$attrs.rowIndex]
  const completedFun = () => {
    selected.value = {}
  }

  createAlert({
    title: t("ARTICLE.DELETE_CONFIRM_TITLE"),
    content: t("ARTICLE.DELETE_CONFIRM_MSG", { name: selected.value.title }),
    onConfirm: () => {
      deleteArticleWithId(selected.value.id)?.finally(completedFun)
    },
    onCancel: completedFun,
    onClose: completedFun
  })
}

function deleteArticleWithId(id: string) {
  if (loading.value) return
  loading.value = true

  return ArticleAPI.deleteWithId(id).finally(() => {
    loading.value = false
  }).then(() => {
    createMessage({ message: t("MSG.DELETED_SUCCESS"), theme: "success" })
    tableReload()
  })
}

// 创建
function onCreateHandler() {
  router.push("/system/article/detail/0")
}

// 修改
function onUpdateHandler({ $attrs }: any) {
  const { id } = list.value[$attrs.rowIndex]
  router.push(`/system/article/detail/${id}`)
}

// 状态相关
function statusText(status: string): string {
  const maps = {
    0: t("ARTICLE.UNPUBLISH"),
    1: t("ARTICLE.PUBLISHED")
  }

  return $helper.getValueByPath<string>(maps, status, "")
}

function articleIsPublished({ $attrs }: any): boolean {
  const { status } = list.value[$attrs.rowIndex]
  return status === "1"
}

function onUnpublishHandler({ $attrs }: any) {
  if (loading.value) return
  loading.value = true
  const { id } = list.value[$attrs.rowIndex]
  ArticleAPI.updateWithId(id, { status: "0" }).finally(() => {
    loading.value = false
  }).then(() => {
    createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
    tableReload()
  })
}

function onPublishHandler({ $attrs }: any) {
  if (loading.value) return
  loading.value = true
  const { id } = list.value[$attrs.rowIndex]
  ArticleAPI.updateWithId(id, { status: "1" }).finally(() => {
    loading.value = false
  }).then(() => {
    createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
    tableReload()
  })
}

// 预览
function onPreviewHandler({ $attrs }: any) {
  const { id } = list.value[$attrs.rowIndex]
  window.open(`${process.env.VUE_APP_SERVICE_PREVIEW}/article/${id}?preview=1`, "_blank")
}
</script>

<style lang="scss" scoped>
.item-cover {
  width: 60px;
  height: 60px;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  border-radius: map-get($--border-radius, "primary");
}

span._published {
  color: map-get($--color, "success");
}
</style>
