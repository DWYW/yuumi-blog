<template>
<div class="home">
  <HeaderComponent></HeaderComponent>
  <AsideComponent>
    <template #prefix>
      <div class="logo">YUUMI</div>
    </template>
  </AsideComponent>

  <main>
    <div class="main-body">
      <div v-for="config in applicationsConfig" :key="config.name" :id="config.name" v-show="applicationIsShow(config)">
        <YuumiEmpty :description="$t('MSG.SS_LOADING')"></YuumiEmpty>
      </div>
    </div>
  </main>
</div>
</template>

<script lang="ts" setup>
import HeaderComponent from "@/components/header/index.vue"
import AsideComponent from "@/components/aside/index.vue"

import { computed, onMounted } from "vue"
import { APPLICATIONS_CONFIG, MicroApplicationConfig } from "@/constant/application/config"
import { registerApplications } from "@/constant/application/index"

function getActiveRule(mode: "history"|"hash", appName: string) {
  if (mode === "hash") {
    const prefix = `${process.env.VUE_APP_PUBLIC_PATH}micro/${appName}`
    return (location: Location) => location.pathname.startsWith(prefix)
  } else {
    const prefix = `#/micro/${appName}`
    return (location: Location) => location.hash.startsWith(prefix)
  }
}

/** 注册应用 */
const applicationsConfig = computed(() => APPLICATIONS_CONFIG.map((item) => {
  return Object.assign({
    activeRule: getActiveRule("history", item.name)
  }, item)
}))

onMounted(() => {
  registerApplications(applicationsConfig.value)
})

function applicationIsShow(config: MicroApplicationConfig) {
  return config.activeRule(location)
}
</script>

<style lang="scss" scoped>
$--header-height: 60px;
$--aside-width: 200px;

.home {
  width: 100vw;
  height: 100vh;
  overflow: hidden;

  header {
    position: fixed;
    top: 0;
    left: $--aside-width;
    right: 0;
    z-index: 9;
    height: $--header-height;
  }

  aside {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 8;

    width: $--aside-width;
    height: 100%;

    .logo {
      font-weight: bold;
      width: 100%;
      height: $--header-height;
      display: flex;
      justify-content: center;
      align-items: center;
      border-bottom: 1px solid #e5e5e5;
    }
  }

  main {
    width: 100%;
    height: 100%;

    padding-top: $--header-height;
    padding-left: $--aside-width;

    .main-body {
      width: 100%;
      height: 100%;
      background-color: #f5f5f5;

      >div {
        width: 100%;
        height: 100%;
      }

      :deep([id^="__qiankun_microapp_wrapper"]) {
        width: 100%;
        height: 100%;
        overflow: auto;
      }
    }
  }
}
</style>
