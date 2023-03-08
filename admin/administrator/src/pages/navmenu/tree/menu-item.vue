<template>
 <div class="form">
  <div class="flex _mc id">
    <div class="prefix">ID</div>
    <div class="content">{{ menuId }}</div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.NAME'"></div>
    <div class="content">
      <YuumiInput v-model="name" clearable></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.LINK_URL'"></div>
    <div class="content">
      <YuumiInput v-model="linkUrl" clearable></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.ACTIVED_RULE'"></div>
    <div class="content">
      <YuumiInput v-model="activedRule" clearable></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.ICON'"></div>
    <div class="content">
      <YuumiInput v-model="icon" clearable></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.WEIGHT'"></div>
    <div class="content">
      <YuumiInput v-model="weight" clearable :limit="onlyNumber"></YuumiInput>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.PARENT'"></div>
    <div class="content">
      <YuumiSelect v-model="parentId" :options="menuOptions"></YuumiSelect>
    </div>
  </div>

  <div class="flex _mc">
    <div class="prefix" v-t="'ATTR.ACCESSIBLE_ROLE'"></div>
    <div class="content">
      <YuumiSelect v-model="roleId" :options="roleOptions"></YuumiSelect>
    </div>
  </div>

  <div class="submit">
    <YuumiButton v-if="menuId" v-t="'HANDLE.DELETE'" theme="error" @click="onDelete"></YuumiButton>
    <YuumiButton v-t="'HANDLE.CANCEL'" @click="onCancel"></YuumiButton>
    <YuumiButton v-if="menuId" v-t="'HANDLE.SAVE'" theme="primary" @click="onUpdate"></YuumiButton>
    <YuumiButton v-else v-t="'HANDLE.CREATE'" theme="primary" @click="onCreate"></YuumiButton>
  </div>
 </div>
</template>

<script lang="ts" setup>
import { NavMenuAPI } from "@/api"
import { MenuData } from "@/api/navmenu"
import { $helper } from "@/common"
import { computed, defineEmits, defineProps, ref, watch } from "vue"
import { createAlert } from "yuumi-ui-vue"
import { useI18n } from "vue-i18n"

const { t } = useI18n()
const emit = defineEmits(["loading", "completed", "updated", "created", "deleted"])
const props: {
  menus: any[]
  menu: any
  roles: any[]
  role: any
} = defineProps({
  menus: { type: Array, default: () => [] },
  menu: { type: Object, default: () => ({}) },
  roles: { type: Array, default: () => [] },
  role: { type: Object, default: () => ({}) }
})

const loading = ref(false)
watch(() => loading.value, (value) => {
  emit("loading", value)
})

const menuId = ref("")
const name = ref("")
const linkUrl = ref("")
const activedRule = ref("")
const icon = ref("")
const weight = ref("")
const parentId = ref("0")
const roleId = ref("0")

const onlyNumber = computed(() => /^\d+$/)
const menuOptions = computed(() => {
  const value: {[key:string]: string}[] = [{ label: "无", value: "0" }]

  return props.menus.reduce((acc, item) => {
    if (item.id === menuId.value) return acc
    return acc.concat({ label: item.name, value: item.id })
  }, value)
})

const roleOptions = computed(() => {
  const value: {[key:string]: string}[] = [{ label: "无", value: "0" }]

  return props.roles.reduce((acc, item) => {
    if (item.id === menuId.value) return acc
    return acc.concat({ label: item.name, value: item.id })
  }, value)
})

watch(() => props.menu, (data: any) => {
  menuId.value = $helper.getValueByPath<string>(data, "id", "")
  name.value = $helper.getValueByPath<string>(data, "name", "")
  linkUrl.value = $helper.getValueByPath<string>(data, "link_url", "")
  activedRule.value = $helper.getValueByPath<string>(data, "actived_rule", "")
  icon.value = $helper.getValueByPath<string>(data, "icon", "")
  weight.value = $helper.getValueByPath<string>(data, "weight", "")
  parentId.value = $helper.getValueByPath<string>(data, "parent_id", "0")
})

watch(() => props.role, (data: any) => {
  roleId.value = $helper.getValueByPath<string>(data, "id", "0")
})

function onDelete() {
  createAlert({
    title: t("NAVMENU.DELETE_CONFIRM_TITLE"),
    content: t("NAVMENU.DELETE_CONFIRM_MSG", { name: name.value }),
    onConfirm: () => {
      loading.value = true
      const _id = menuId.value
      const prevRoleId = $helper.getValueByPath<string>(props.role, "id", "0")

      const promise = prevRoleId !== "0" ? NavMenuAPI.unbindRoles(_id, [prevRoleId]) : Promise.resolve()
      promise.then(() => {
        return NavMenuAPI.deleteWithId(_id).then(() => {
          emit("completed")
          emit("deleted", { id: _id })
        })
      }).finally(() => {
        loading.value = false
      })
    }
  })
}

function onCancel() {
  emit("completed")
}

function onUpdate() {
  if (loading.value) return
  loading.value = true

  const promises = [
    NavMenuAPI.updateWithId(menuId.value, getFormData())
  ]

  const prevRoleId = $helper.getValueByPath<string>(props.role, "id", "0")
  const currRoleId = roleId.value
  if (prevRoleId !== currRoleId) {
    if (currRoleId !== "0") {
      promises.push(NavMenuAPI.bindRoles(menuId.value, [currRoleId]))
    }

    if (prevRoleId !== "0") {
      promises.push(NavMenuAPI.unbindRoles(menuId.value, [prevRoleId]))
    }
  }

  Promise.all(promises).finally(() => {
    loading.value = false
  }).then(([{ menu }]) => {
    if (currRoleId !== "0") {
      menu.roles = [props.roles.find((item) => item.id === currRoleId)]
    }

    emit("completed")
    emit("updated", menu)
  })
}

function onCreate() {
  if (loading.value) return
  loading.value = true

  NavMenuAPI.create(getFormData()).finally(() => {
    loading.value = false
  }).then(({ menu }) => {
    const currRoleId = roleId.value
    if (currRoleId !== "0") {
      return NavMenuAPI.bindRoles(menu.id, [currRoleId]).then(() => {
        menu.roles = [props.roles.find((item) => item.id === currRoleId)]
        return menu
      })
    }

    return menu
  }).then((menu) => {
    emit("completed")
    emit("created", menu)
  })
}

function getFormData() {
  const v: MenuData = {
    name: name.value,
    link_url: linkUrl.value,
    actived_rule: activedRule.value,
    icon: icon.value
  }

  if (weight.value) v.weight = weight.value
  if (parentId.value) v.parent_id = parentId.value
  return v
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
