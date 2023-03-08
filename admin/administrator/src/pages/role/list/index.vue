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
        <RoleItem
          :roles="roles"
          :role="selected"
          @loading="onLoading"
          @completed="onCompleted"
          @created="onCreated"
          @deleted="onDeleted"
          @updated="onUpdated"
        ></RoleItem>
      </div>
    </div>
  </div>
  <div class="page-footer">
    <div class="content"></div>
  </div>
</div>
</template>

<script lang="ts" setup>
import RoleItem from "./role-item.vue"
import { RoleAPI } from "@/api"
import { $helper } from "@/common"
import { computed, ComputedRef, onMounted, Ref, ref } from "vue"
import { createMessage } from "yuumi-ui-vue"
import { useI18n } from "vue-i18n"
const { t } = useI18n()

onMounted(() => {
  getRoles()
})

const loading = ref(false)

const optionKey = computed(() => ({ value: "id", label: "name" }))
const roles: Ref<any[]> = ref([])
const treeData: ComputedRef<any[]> = computed(() => $helper.list2tree(roles.value))

function getRoles() {
  if (loading.value) return
  loading.value = true

  return RoleAPI.getRoles().finally(() => {
    loading.value = false
  }).then((res: any) => {
    roles.value = res.roles
  })
}

const selected: Ref<any> = ref({})

function onNodeClick({ instance }: any) {
  selected.value = instance.node
}

// 事件
function onLoading(value: boolean) {
  loading.value = value
}

function onCompleted() {
  selected.value = {}
}

function onCreated(role: any) {
  createMessage({ message: t("MSG.CREATED_SUCCESS"), theme: "success" })
  roles.value.unshift(role)
}

function onDeleted(id: string) {
  createMessage({ message: t("MSG.DELETED_SUCCESS"), theme: "success" })
  const index = roles.value.findIndex(item => item.id === id)
  if (index > -1) {
    roles.value.splice(index, 1)
  }
  selected.value = null
}

function onUpdated(role: any) {
  createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
  const index = roles.value.findIndex(item => item.id === role.id)
  if (index > -1) {
    roles.value[index] = role
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
