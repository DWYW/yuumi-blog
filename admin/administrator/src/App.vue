<template>
  <router-view/>
</template>

<script lang="ts" setup>
import { onBeforeMount } from "vue"
import { useStore } from "@/store"

const store = useStore()

onBeforeMount(() => {
  // 自运行时根据本地缓存同步数据到store
  if (!window.__POWERED_BY_QIANKUN__) {
    store.dispatch("global/syncFromLocal")
  }
})
</script>

<style lang="scss">
@import "./styles/base.scss";
@import "./styles/page.scss";
@import "./styles/layout.scss";

.micro-app {
  width: 100%;
  height: 100%;
}
</style>
