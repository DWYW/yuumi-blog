<template>
<div class="page" v-loading="loading">
  <div class="page-header">
    <div class="content flex">
      <div class="prefix"></div>
      <div class="content"></div>
      <div class="suffix">
        <YuumiButton theme="primary" v-t="'HANDLE.CREATE'" @click="createHandler"></YuumiButton>
      </div>
    </div>
  </div>
  <div class="page-body">
    <div class="content" v-if="list.length">
      <YuumiTable :data="list" :row-class-name="rowClassName">
        <YuumiTableColumn :title="$t('ATTR.NAME')" prop="name"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.CREATED_DATE')" prop="created_date" :width="200"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.HANDLE')" fixed="right">
          <template #default="scope">
            <YuumiButton v-t="'HANDLE.ROLE_UPDATE'" @click="updateRoleHandler(scope)"></YuumiButton>
            <YuumiButton v-t="'HANDLE.RESET_NAME'" @click="updateNameHandler(scope)"></YuumiButton>
            <YuumiButton v-t="'HANDLE.DELETE'" @click="deleteHandler(scope)"></YuumiButton>
          </template>
        </YuumiTableColumn>
      </YuumiTable>
    </div>
  </div>
  <div class="page-footer">
    <div class="content">
      <YuumiPagination
        :page="pagination.page"
        :page-total="pagination.pageTotal"
        :total="pagination.total"
        @change="onPaginationChange"
      ></YuumiPagination>
    </div>
  </div>

  <CreateAdministrator v-model="createDialogVisible" @success="onCreateSuccess"></CreateAdministrator>
  <UpdateAdministratorName v-model="updateNameDialogVisible"
    :administrator-id="selected.id"
    @success="onUpdateSuccess"
    @complete="onUpdateComplete"
  ></UpdateAdministratorName>

  <updateAdministratorRole v-model="updateRoleDrawerVisible"
    :administrator-id="selected.id"
  ></updateAdministratorRole>
</div>
</template>

<script lang="ts" setup>
import CreateAdministrator from "./create.vue"
import UpdateAdministratorName from "./update-name.vue"
import updateAdministratorRole from "./update-role.vue"
import { createAlert, createMessage } from "yuumi-ui-vue"
import { onMounted, Ref, ref, watch } from "vue"
import { AdministratorAPI, RoleAPI } from "@/api"
import { $day } from "@/common"
import { useTableHelper } from "@/components/helper/table"
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const loading: Ref<boolean> = ref(false)
const { list, pagination, tableReload, onPaginationChange } = useTableHelper(() => {
  if (loading.value) return
  loading.value = true

  return AdministratorAPI.getList({
    page: pagination.value.page
  }).finally(() => {
    loading.value = false
  }).then(({ administrators, pagination }) => {
    administrators = administrators.map((item: any) => {
      item.created_date = $day.formatTimeFromDateString(item.created_at)
      return item
    })
    return { list: administrators, pagination }
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

// 删除管理员
function deleteHandler({ $attrs }: any) {
  const item = list.value[$attrs.rowIndex]

  createAlert({
    title: t("ADMINISTRATOR.DELETE_CONFIRM_TITLE"),
    content: t("ADMINISTRATOR.DELETE_CONFIRM_MSG", { name: item.name }),
    onConfirm: () => {
      confirmDelete(item.id)
    }
  })
}

function confirmDelete(id: string) {
  if (loading.value) return
  loading.value = true

  // 首先删除关联的角色
  RoleAPI.getRolesWithAdministratorId(id).then(({ roles }) => {
    if (roles.length) {
      return AdministratorAPI.removeRoles(id, roles.map((item: any) => item.id))
    }
  }).then(() => {
    return AdministratorAPI.deleteWithId(id)
  }).then(() => {
    createMessage({ message: t("MSG.DELETED_SUCCESS"), theme: "success" })
    tableReload()
  }).finally(() => {
    loading.value = false
  })
}

// 创建管理员
const createDialogVisible = ref(false)

function createHandler() {
  createDialogVisible.value = true
}

function onCreateSuccess() {
  createMessage({ message: t("MSG.CREATED_SUCCESS"), theme: "success" })

  if (pagination.value.page === "1") {
    tableReload()
  } else {
    onPaginationChange(1)
  }
}

// 更新
function onUpdateSuccess() {
  createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })

  if (pagination.value.page === "1") {
    tableReload()
  } else {
    onPaginationChange(1)
  }
}

function onUpdateComplete() {
  selected.value = {}
}

// 修改Name
const updateNameDialogVisible = ref(false)

function updateNameHandler({ $attrs }: any) {
  selected.value = list.value[$attrs.rowIndex]
  updateNameDialogVisible.value = true
}

// 修改绑定的角色
const updateRoleDrawerVisible = ref(false)
watch(() => updateRoleDrawerVisible.value, (value: boolean) => {
  if (!value) {
    selected.value = {}
  }
})

function updateRoleHandler({ $attrs }: any) {
  selected.value = list.value[$attrs.rowIndex]
  updateRoleDrawerVisible.value = true
}
</script>

<style lang="scss" scoped>
</style>
