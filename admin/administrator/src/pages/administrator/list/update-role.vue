<template>
<YuumiDrawer width="40%"
  :model-value="modelValue"
  @update:modelValue="onUpdateModelValue"
  @before-open="onBeforeOpen"
>
  <div class="content">
    <YuumiTree :data="treeData" :option-key="optionKey" @checked="onNodeChecked" ref="tree"></YuumiTree>
  </div>
</YuumiDrawer>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, ref, Ref, computed } from "vue"
import { AdministratorAPI, RoleAPI } from "@/api"
import { $helper, $validator } from "@/common"
import { createMessage } from "yuumi-ui-vue"
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const emit = defineEmits(["update:modelValue"])
const props = defineProps({
  modelValue: { type: Boolean },
  administratorId: { type: String, default: "" }
})

function onUpdateModelValue(value: boolean) {
  emit("update:modelValue", value)
}

const roles: Ref<any[]> = ref([])
const optionKey = computed(() => ({ label: "name", value: "id" }))
const treeData = computed(() => $helper.list2tree(roles.value))
const owner: Ref<any[]> = ref([])

function onBeforeOpen() {
  const promises = [
    RoleAPI.getRolesWithAdministratorId(props.administratorId).then((res) => {
      owner.value = res.roles
    })
  ]

  if (roles.value.length === 0) {
    promises.push(RoleAPI.getRoles().then((res) => {
      roles.value = res.roles
    }))
  }

  Promise.all(promises).then(() => {
    roles.value.forEach((role) => {
      role.checked = $validator.isDefine(owner.value.find((item) => item.id === role.id))
    })
  })
}

const tree: Ref = ref()

function onNodeChecked() {
  const checkedNodes = tree.value.getCheckedNodes(true)

  const { add, remove } = $helper.treeNodeDiff(owner.value, checkedNodes)
  const promises = []
  if (add.length) {
    promises.push(AdministratorAPI.appendRoles(props.administratorId, add.map((item) => item.id)))
  }
  if (remove.length) {
    promises.push(AdministratorAPI.removeRoles(props.administratorId, remove.map((item) => item.id)))
  }
  Promise.all(promises).then(() => {
    createMessage({ message: t("MSG.UPDATED_SUCCESS"), theme: "success" })
    owner.value = checkedNodes
  })
}
</script>

<style lang="scss" scoped>
.content {
  padding: map-get($--space, "sm");
}
</style>
