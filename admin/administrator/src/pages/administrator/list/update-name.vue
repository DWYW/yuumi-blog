<template>
<YuumiDialog id="update"
  :title="$t('ADMINISTRATOR.UPDATE_NAME_TITLE')"
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
</YuumiDialog>
</template>

<script lang="ts" setup>
import { AdministratorAPI } from "@/api"
import { ref, defineEmits, defineProps } from "vue"
import { useI18n } from "vue-i18n"
import { createMessage, createLoading, removeLoading } from "yuumi-ui-vue"

const emit = defineEmits(["update:modelValue", "success", "complete"])
const props = defineProps({
  administratorId: { type: String, default: "" },
  modelValue: { type: Boolean, default: true }
})

function onUpdateModelValue(value: boolean) {
  emit("update:modelValue", value)
}

const { t } = useI18n()
const name = ref("")

function onConfirm() {
  if (!name.value) return createMessage({ message: t("ADMINISTRATOR.NAME_REQUIRED"), theme: "warn" })
  const vnode = createLoading({ teleport: "#update .dialog-panel" })
  AdministratorAPI.updateName(props.administratorId, name.value).finally(() => {
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
}
</script>

<style lang="scss" scoped>
.flex {
  margin-top: map-get($--space, "sm");
  .content {
    display: flex;

    :deep(.yuumi-input) {
      width: 100%;
    }
  }
}
</style>
