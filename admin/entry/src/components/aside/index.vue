<template>
<aside>
  <div class="prefix">
    <slot name="prefix"></slot>
  </div>

  <div class="main">
    <YuumiScrollbar>
      <YuumiNavMenu :data="navs"
        :option-key="optionKey"
        :selected-node="selected"
        @node-click="nodeOnClick"
      ></YuumiNavMenu>
    </YuumiScrollbar>
  </div>

  <div class="suffix"></div>
</aside>
</template>

<script lang="ts" setup>
import { $helper, $storage } from "@/common"
import { REDIRECT_URI } from "@/constant/keys"
import { computed, onMounted } from "@vue/runtime-core"
import { useStore } from "@/store"
import { useRoute, useRouter } from "vue-router"
import { ref } from "vue"
const store = useStore()
const router = useRouter()
const route = useRoute()
const { application } = store.state

onMounted(() => {
  let promise:Promise<any> = Promise.resolve()
  if (!application.navsLoaded) {
    promise = store.dispatch("application/updateNavs")
  }

  promise.then(() => {
    const redirected = navigatorRedirect() || navigatorFistChild() || Promise.resolve()
    redirected.then(() => {
      const node = getActivedTreeNode()
      if (node) selected.value.id = node.id
    })
  })
})

function navigatorRedirect() {
  const target = $storage.session.getItem(REDIRECT_URI)
  if (target) {
    return router.replace(target)
  }
}

function navigatorFistChild() {
  const target = $helper.getTreeNode((node) => {
    return !node.children && node.link_url
  }, application.navs)

  if (route.path === "" || route.path === "/") {
    if (target && target.link_url) {
      return router.replace(target.link_url)
    }
  }

  if (!getActivedTreeNode()) {
    return router.replace(target)
  }
}

function getActivedTreeNode() {
  return $helper.getTreeNode((node) => {
    if (node.children || !node.link_url) return false
    if (node.actived_rule) {
      return new RegExp(node.actived_rule).test(route.path)
    }

    return route.path.startsWith(node.link_url)
  }, application.navs)
}

const navs = computed(() => application.navs)
const optionKey = computed(() => ({
  label: "name",
  value: "id"
}))

/** 设置选中的菜单 */
const selected = ref({ id: "" })

function nodeOnClick({ node }: any) {
  if (node.children) return
  selected.value.id = node.id

  if (node.link_url) {
    router.push(node.link_url)
  }
}
</script>

<style lang="scss" scoped>
aside {
  border-right: 1px solid map-get($--color, "divider");
  display: flex;
  flex-direction: column;

  .prefix, .suffix {
    flex: 0 0 auto;
  }

  .main {
    flex: 1 1 1px;
  }
}

:deep(.yuumi-navmenu)  {
  .menu-node .node-content {
    .prefix-icon {
      font-size: 1.2em;
    }
    .content__label {
      line-height: 2;
    }
  }
}
</style>
