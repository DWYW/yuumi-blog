<template>
<div class="page" v-loading="loading">
  <div class="page-header">
    <div class="content"></div>
  </div>
  <div class="page-body">
    <div class="content flex">
      <div class="prefix tree">
        <YuumiTree :data="treeData" :option-key="optionKey" :checkable="false" @node-click="onNodeClick">
          <template #default="{node}">
            <YuumiIcon :icon="node.icon" style="margin-right: 8px;"></YuumiIcon>
            <span>{{ node.name }}</span>
          </template>
        </YuumiTree>
      </div>

      <div class="content">
        <MenuItem :menu="selected" :menus="menus" :roles="roles" :role="selectedRole"
          @loading="onLoading"
          @completed="onCompleted"
          @created="onCreated"
          @updated="onUpdated"
          @deleted="onDeleted"
        ></MenuItem>
      </div>
    </div>
  </div>
  <div class="page-footer"></div>
</div>
</template>

<script lang="ts" setup>
import MenuItem from "./menu-item.vue"
import { computed, onMounted, Ref, ref } from "vue"
import { NavMenuAPI, RoleAPI } from "@/api"
import { $helper } from "@/common"
import { createMessage } from "yuumi-ui-vue"
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const loading: Ref<boolean> = ref(false)
const menus: Ref<any[]> = ref([])
const optionKey = computed(() => ({
  label: "name",
  value: "id"
}))
const treeData = computed(() => {
  return $helper.list2tree(menus.value)
})

function getNavMenus() {
  return NavMenuAPI.getNavMenus({ preload_roles: "1" }).then((res) => {
    menus.value = res.menus
  })
}

const roles = ref([])

function getRoles() {
  return RoleAPI.getRoles().then((res) => {
    roles.value = res.roles
  })
}

onMounted(() => {
  loading.value = true
  Promise.all([getNavMenus(), getRoles()]).finally(() => {
    loading.value = false
  })
})

const selected: Ref<any> = ref({})
const selectedRole = computed(() => $helper.getValueByPath<any>(selected.value, "roles[0]"))

function onNodeClick({ instance } : any) {
  const { node } = instance
  selected.value = node
}

function onLoading(value: boolean) {
  loading.value = value
}

function onCompleted() {
  selected.value = {}
}

function onDeleted({ id }: any) {
  createMessage({ message: t("MSG.DELETED_SUCCESS"), theme: "success" })
  const index = menus.value.findIndex((item) => item.id === id)
  if (index > -1) {
    menus.value.splice(index, 1)
  }
}

function onCreated(detail: any) {
  createMessage({ message: t("MSG.CREATED_SUCCESS"), theme: "success" })
  menus.value.push(detail)
}

function onUpdated(detail: any) {
  createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
  const index = menus.value.findIndex((item) => item.id === detail.id)
  if (index > -1) {
    menus.value[index] = detail
  }
}
</script>

<style lang="scss" scoped>
.tree {
  min-width: 400px;
  border-right: 1px solid map-get($--color, "border");
  padding-right: map-get($--space, "sm");

  :deep(.tree-node) {
    line-height: 2;
  }
}
</style>
