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
        <YuumiTableColumn :title="$t('ATTR.NAME')" prop="name" :width="100"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.RPC_METHOD')" prop="rpc_method" :width="200"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.ROLE')" prop="roles" :width="100">
          <template #default="{ $value }">
            <span v-for="item in $value" :key="item.id">{{ item.name }}</span>
          </template>
        </YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.CREATED_DATE')" prop="created_date" :width="200"></YuumiTableColumn>
        <YuumiTableColumn :title="$t('ATTR.HANDLE')" fixed="right">
          <template #default="scope">
            <YuumiButton v-t="'HANDLE.BIND_ROLE'" @click="updateRoleHandler(scope)" :width="200"></YuumiButton>
            <YuumiButton v-t="'HANDLE.EDIT'" @click="updateHandler(scope)"></YuumiButton>
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

  <CreatePermission v-model="createDialogVisible" @success="onCreateSuccess"></CreatePermission>

  <UpdatePermission v-model="updateDialogVisible"
    :permission="selected"
    @success="onUpdateSuccess"
    @complete="onUpdateComplete"
  ></UpdatePermission>

  <UpdateRole v-model="updateRoleDialogVisible"
    :permission="selected"
    :role="role"
    @success="onUpdateSuccess"
    @complete="onUpdateComplete"
  ></UpdateRole>
</div>
</template>

<script lang="ts" setup>
import { createAlert, createMessage } from "yuumi-ui-vue"
import { computed, onMounted, Ref, ref } from "vue"
import { useI18n } from "vue-i18n"
import { PermissionAPI } from "@/api"
import { $day } from "@/common"
import { useTableHelper } from "@/components/helper/table"
import CreatePermission from "./create.vue"
import UpdatePermission from "./update.vue"
import UpdateRole from "./update-role.vue"

const { t } = useI18n()
const loading: Ref<boolean> = ref(false)
const { list, pagination, tableReload, onPaginationChange } = useTableHelper(() => {
  if (loading.value) return
  loading.value = true

  return PermissionAPI.getList({
    page: pagination.value.page,
    preload_roles: "1"
  }).finally(() => {
    loading.value = false
  }).then(({ permissions, pagination }) => {
    permissions = permissions.map((item: any) => {
      item.created_date = item.created_date = $day.formatTimeFromDateString(item.created_at)
      return item
    })
    return { list: permissions, pagination }
  })
})

const selected: Ref<any> = ref({})
const role = computed(() => selected.value.roles ? selected.value.roles[0] : {})
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
function deleteHandler({ $attrs }: any) {
  selected.value = list.value[$attrs.rowIndex]

  createAlert({
    title: t("PERMISSION.DELETE_CONFIRM_TITLE"),
    content: t("PERMISSION.DELETE_CONFIRM_MSG", { name: selected.value.name }),
    onConfirm: () => {
      confirmDelete(selected.value.id)
    },
    onCancel: () => { selected.value = {} },
    onClose: () => { selected.value = {} }
  })
}

function confirmDelete(id: string) {
  if (loading.value) return
  loading.value = true
  return PermissionAPI.deleteWithId(id).finally(() => {
    loading.value = false
  }).then(() => {
    createMessage({ message: t("MSG.DELETED_SUCCESS"), theme: "success" })
    tableReload()
  })
}

// 创建
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
function onUpdateComplete() {
  selected.value = {}
}

function onUpdateSuccess() {
  createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
  tableReload()
}

const updateDialogVisible = ref(false)

function updateHandler({ $attrs }: any) {
  selected.value = list.value[$attrs.rowIndex]
  updateDialogVisible.value = true
}

// 更新角色
const updateRoleDialogVisible = ref(false)
function updateRoleHandler({ $attrs }: any) {
  selected.value = list.value[$attrs.rowIndex]
  updateRoleDialogVisible.value = true
}
</script>

<style lang="scss" scoped>
</style>
