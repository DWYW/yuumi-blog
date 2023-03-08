<template>
<YuumiDialog id="update"
  :title="$t('PERMISSION.UPDATE_ROLE_TITLE')"
  :sync="false"
  :model-value="modelValue"
  @update:modelValue="onUpdateModelValue"
  @confirm="onConfirm"
  @cancel="onCancel"
  @close="onClose"
  @beforeEnter="onBeforeEnter"
>
  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.ROLE'"></div>
    <div class="content">
      <YuumiSelect v-model="roleId" :options="roles" :option-key="optionKey"></YuumiSelect>
    </div>
  </div>
</YuumiDialog>
</template>

<script lang="ts" setup>
import { RoleAPI } from "@/api"
import { ref, defineEmits, defineProps, Ref, computed } from "vue"
import { createLoading, removeLoading } from "yuumi-ui-vue"

const emit = defineEmits(["update:modelValue", "success", "complete"])
const props = defineProps({
  modelValue: { type: Boolean, default: true },
  permission: { type: Object, default: () => ({}) },
  role: { type: Object, default: () => ({}) }
})

function onUpdateModelValue(value: boolean) {
  emit("update:modelValue", value)
}

const roles: Ref<any[]> = ref([])
const optionKey = computed(() => ({ label: "name", value: "id" }))
const roleId = ref("")

function onBeforeEnter() {
  if (roles.value.length > 0) {
    roleId.value = props.role.id
    return
  }

  RoleAPI.getRoles().then((res) => {
    roles.value = res.roles
    roleId.value = props.role.id
  })
}

function onConfirm() {
  if (roleId.value === props.role.id) {
    return onClose()
  }

  const vnode = createLoading({ teleport: "#update .dialog-panel" })
  const promises = [
    RoleAPI.appendPermissions(roleId.value, [props.permission.id])
  ]

  if (props.role.id) {
    promises.push(RoleAPI.deletePermissions(props.role.id, [props.permission.id]))
  }

  Promise.all(promises).finally(() => {
    removeLoading(vnode)
  }).then(() => {
    completed()
    emit("success")
  })
}

function onCancel() {
  completed()
}

function onClose() {
  completed()
}

function completed() {
  emit("complete")
  clearValue()
  onUpdateModelValue(false)
}

function clearValue() {
  roleId.value = ""
}
</script>

<style lang="scss" scoped>
.flex {
  margin-bottom: map-get($--space, "sm");
  .prefix {
    flex-basis: 60px;
    text-align: right;
  }
  .content {
    display: flex;

    :deep(.yuumi-input), :deep(.yuumi-select) {
      width: 100%;
    }
  }
}
</style>
