<template>
<div class="update-role form">
  <div class="flex id">
    <div class="prefix">ID</div>
    <div class="content"> {{ roleId ||  $t("ROLE.ID_EMPTY")}}</div>
  </div>

  <div class="flex">
    <div class="prefix" v-t="'ATTR.NAME'"></div>
    <div class="content">
      <YuumiInput v-model="name" clearable></YuumiInput>
    </div>
  </div>

  <div class="flex">
    <div class="prefix" v-t="'ATTR.TYPE'"></div>
    <div class="content">
      <YuumiInput v-model="type" clearable></YuumiInput>
    </div>
  </div>

  <div class="flex">
    <div class="prefix" v-t="'ATTR.PARENT'"></div>
    <div class="content">
      <YuumiSelect v-model="parentId" :options="options" clearable></YuumiSelect>
    </div>
  </div>

  <div class="submit">
    <template v-if="roleId">
      <YuumiButton v-t="'HANDLE.DELETE'" theme="error" @click="onDelete"></YuumiButton>
      <YuumiButton v-t="'HANDLE.CANCEL'" @click="onCancel"></YuumiButton>
      <YuumiButton v-t="'HANDLE.SAVE'" theme="primary" @click="onUpdate"></YuumiButton>
    </template>
    <template v-else>
      <YuumiButton v-t="'HANDLE.RESET'" @click="onReset"></YuumiButton>
      <YuumiButton v-t="'HANDLE.CREATE'" theme="primary" @click="onCreate"></YuumiButton>
    </template>
  </div>
</div>
</template>

<script lang="ts" setup>
import { RoleAPI } from "@/api"
import { $helper } from "@/common"
import { defineProps, defineEmits, ref, watch, computed } from "vue"
import { useI18n } from "vue-i18n"
import { createAlert, createMessage } from "yuumi-ui-vue"
const { t } = useI18n()

const props = defineProps({
  roles: { type: Array },
  role: { type: Object }
})

const emit = defineEmits(["loading", "completed", "created", "deleted", "updated"])

const roleId = ref("")
const name = ref("")
const type = ref("")
const parentId = ref("0")
const options = computed(() => {
  const value = [{ label: "无", value: "0" }]

  return (props.roles || []).reduce((acc: any, item: any) => {
    if (item.id === roleId.value) return acc
    return acc.concat({ label: item.name, value: item.id })
  }, value)
})

watch(() => props.role, (role) => {
  roleId.value = $helper.getValueByPath<string>(role, "id", "")
  name.value = $helper.getValueByPath<string>(role, "name", "")
  type.value = $helper.getValueByPath<string>(role, "type", "")
  parentId.value = $helper.getValueByPath<string>(role, "parent_id", "0")
})

function onDelete() {
  createAlert({
    title: t("ROLE.DELETE_CONFIRM_TITLE"),
    content: t("ROLE.DELETE_CONFIRM_MSG", { name: name.value }),
    onConfirm: () => {
      emit("loading", true)
      const id = roleId.value
      RoleAPI.deleteWithId(id)
        .finally(() => emit("loading", false))
        .then(() => {
          emit("completed")
          emit("deleted", id)
        })
    }
  })
}

function getFormData() {
  if (!name.value) {
    createMessage({ message: t("ROLE.NAME_REQUIRED_MSG"), theme: "warn" })
    return
  }

  if (!type.value) {
    createMessage({ message: t("ROLE.TYPE_REQUIRED_MSG"), theme: "warn" })
    return
  }

  const v: any = {
    name: name.value,
    type: type.value
  }

  if (parentId.value !== "0") {
    v.parent_id = parentId.value
  }

  return v
}

function onCreate() {
  const v = getFormData()
  if (!v) return

  emit("loading", true)
  return RoleAPI.create(v)
    .finally(() => emit("loading", false))
    .then(({ role }: any) => {
      emit("completed")
      emit("created", role)
    })
}

function onReset() {
  name.value = ""
  type.value = ""
  parentId.value = "0"
}

function onCancel() {
  emit("completed")
}

function onUpdate() {
  emit("loading", true)
  return RoleAPI.updateWithId(roleId.value, getFormData())
    .finally(() => emit("loading", false))
    .then(({ role }: any) => {
      emit("completed")
      emit("updated", role)
    })
}
</script>

<style lang="scss" scoped>
.form {
  max-width: 800px;
  padding-left: map-get($--space, "md");
  margin: 0 auto;
  .flex  {
    margin-bottom: map-get($--space, "sm");
    .prefix {
      padding-right: map-get($--space, "xm");
      flex-basis: 100px;
      text-align: right;
    }
    .content {
      display: flex;

      :deep(.yuumi-input), :deep(.yuumi-select) {
        width: 100%;
      }
    }

    &.id .content {
      color: map-get($--color, "placeholder");
    }
  }

  .submit {
    text-align: center;

    button:not(:last-child) {
      margin-right: map-get($--space, "sm");
    }
  }
}
</style>
