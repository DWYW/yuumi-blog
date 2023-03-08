<template>
<YuumiDialog id="update"
  :title="$t('PERMISSION.UPDATE_TITLE')"
  :sync="false"
  :model-value="modelValue"
  @update:modelValue="onUpdateModelValue"
  @confirm="onConfirm"
  @cancel="onCancel"
  @close="onClose"
  @beforeEnter="onBeforeEnter"
>
  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.NAME'"></div>
    <div class="content">
      <YuumiInput v-model="name"></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.RPC_METHOD'"></div>
    <div class="content">
      <YuumiInput v-model="rpcMethod"></YuumiInput>
    </div>
  </div>
</YuumiDialog>
</template>

<script lang="ts" setup>
import { PermissionAPI } from "@/api"
import { ref, defineEmits, defineProps } from "vue"
import { useI18n } from "vue-i18n"
import { createMessage, createLoading, removeLoading } from "yuumi-ui-vue"

const emit = defineEmits(["update:modelValue", "success", "complete"])
const props = defineProps({
  modelValue: { type: Boolean, default: true },
  permission: { type: Object, default: () => ({}) }
})

function onUpdateModelValue(value: boolean) {
  emit("update:modelValue", value)
}

const { t } = useI18n()
const name = ref("")
const rpcMethod = ref("")

function onBeforeEnter() {
  name.value = props.permission.name
  rpcMethod.value = props.permission.rpc_method
}

function onConfirm() {
  if (!name.value) return createMessage({ message: t("PERMISSION.NAME_REQUIRED_MSG"), theme: "warn" })
  if (!rpcMethod.value) return createMessage({ message: t("PERMISSION.RPC_METHOD_REQUIRED_MSG"), theme: "warn" })
  const vnode = createLoading({ teleport: "#update .dialog-panel" })
  PermissionAPI.updateWithId(props.permission.id, {
    name: name.value,
    rpc_method: rpcMethod.value
  }).finally(() => {
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
  name.value = ""
  rpcMethod.value = ""
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

    :deep(.yuumi-input) {
      width: 100%;
    }
  }
}
</style>
