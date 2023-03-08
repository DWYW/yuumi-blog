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
            <img style="width: 100px;" :src="$value">
          </template>
        </YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.DESCRIPTION')" prop="description" :width="200"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.CREATED_DATE')" prop="created_date" :width="160"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.UPDATED_DATE')" prop="updated_date" :width="160"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.HANDLE')" fixed="right">
          <template #default="scope">
            <YuumiButton theme="error" v-t="'HANDLE.DELETE'" @click="onDeleteHandler(scope)"></YuumiButton>
            <YuumiButton theme="primary" v-t="'HANDLE.EDIT'" @click="onUpdateHandler(scope)"></YuumiButton>
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
import { $day } from "@/common"
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
</script>

<style lang="scss" scoped>
</style>
