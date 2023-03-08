<template>
  <router-view></router-view>
</template>

<script lang="ts" setup>
import { onBeforeMount } from "@vue/runtime-core"
import { useStore } from "./store"

const store = useStore()

onBeforeMount(() => {
  store.dispatch("global/syncFromLocal")
})
</script>

<style lang="scss">
@import "./styles/index.scss";
body {
  font-size: map-get($--font-size, "sm");
  color: map-get($--text-color, "primary");
}
</style>
