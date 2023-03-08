<template>
<YuumiDialog id="create"
  :title="$t('ADMINISTRATOR.CREATE_TITLE')"
  :sync="false"
  :model-value="modelValue"
  @update:modelValue="onUpdateModelValue"
  @confirm="onConfirm"
  @cancel="onCancel"
  @close="onClose"
>
  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.NAME'"></div>
    <div class="content">
      <YuumiInput v-model="name"></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.PASSWORD'"></div>
    <div class="content">
      <YuumiInput type="password" v-model="password"></YuumiInput>
    </div>
  </div>
</YuumiDialog>
</template>

<script lang="ts" setup>
import { AdministratorAPI } from "@/api"
import { ref, defineEmits, defineProps } from "vue"
import { useI18n } from "vue-i18n"
import { createMessage, createLoading, removeLoading } from "yuumi-ui-vue"

const emit = defineEmits(["update:modelValue", "success"])

defineProps({
  modelValue: { type: Boolean, default: true }
})

function onUpdateModelValue(value: boolean) {
  emit("update:modelValue", value)
}

const { t } = useI18n()
const name = ref("")
const password = ref("")

function onConfirm() {
  if (!name.value) return createMessage({ message: t("ADMINISTRATOR.NAME_REQUIRED"), theme: "warn" })
  if (!password.value) return createMessage({ message: t("ADMINISTRATOR.PASSWORD_REQUIRED"), theme: "warn" })
  const vnode = createLoading({ teleport: "#create .dialog-panel" })
  AdministratorAPI.create({
    name: name.value,
    password: password.value
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
  clearValue()
  onUpdateModelValue(false)
}

function clearValue() {
  name.value = ""
  password.value = ""
}
</script>

<style lang="scss" scoped>
.flex {
  margin-bottom: map-get($--space, "sm");
  .content {
    display: flex;

    :deep(.yuumi-input) {
      width: 100%;
    }
  }
}
</style>
