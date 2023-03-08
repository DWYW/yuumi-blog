<template>
<header>
  <div class="prefix">
    <slot name="prefix"></slot>
  </div>

  <div class="_expand"></div>

  <YuumiPopper placement="bottom-end">
    <template v-slot:trigger>
      <div class="user">
        <YuumiIcon icon="line-user"></YuumiIcon>
        <span class="username">{{username}}</span>
      </div>
    </template>
    <template v-slot:default>
      <div class="menus">
        <div class="menu-item" v-for="menu in menus" :key="menu.name" @click="menu.onClick">{{menu.name}}</div>
      </div>
    </template>
  </YuumiPopper>
</header>
</template>

<script lang="ts" setup>
import { computed } from "vue"
import { useStore } from "@/store"
import { useI18n } from "vue-i18n"
import { createAlert } from "yuumi-ui-vue"
const store = useStore()
const { t } = useI18n()
const username = computed(() => store.state.global.user?.name)

// 弹出的菜单
const menus = computed((): any[] => {
  return [{
    name: t("SIGNOUT.SIGNOUT"),
    onClick: signout
  }]
})

function signout() {
  createAlert({
    title: t("SIGNOUT.CONFIRM_TITLE"),
    content: t("SIGNOUT.CONFIRM_MSG"),
    alignCenter: true,
    onConfirm: () => {
      store.dispatch("application/signout")
    }
  })
}
</script>

<style lang="scss" scoped>
header {
  display: flex;
  align-items: center;
  padding: 0 map-get($--space, "sm");

  ._expand {
    flex: 1 1 auto;
  }
}

.user {
  display: flex;
  align-items: center;
  min-width: 60px;
  cursor: pointer;

  .yuumi-icon {
    font-size: map-get($--font-size, "md");
  }
  .username {
    padding-left: map-get($--space, "xm")*0.5;
  }
}

.menus {
  .menu-item {
    min-width: 120px;
    padding: map-get($--space, "xm") map-get($--space, "md");
    cursor: pointer;

    &:hover {
      color: map-get($--color, "primary");
    }
  }
}
</style>
