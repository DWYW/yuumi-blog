<template>
<header class="page-header">
  <div class="content">
    <div class="prefix">
      <NuxtLink to="/">首页</NuxtLink>
    </div>
    <div class="content">
      <slot></slot>
    </div>
    <div class="suffix">
      <input class="search" type="text" placeholder="搜索" v-model="keyword" @change="onChange">
      <span class="icon icon-search"></span>
    </div>
  </div>
</header>
</template>

<script lang="ts" setup>
const route = useRoute()
const keyword = ref(route.query.search || "")
function onChange() {
  window.open(`${location.origin}/article/list?search=${keyword.value}`, "_target")
}
</script>

<style lang="scss" scoped>
.page-header {
  position: sticky;
  top: 0;
  z-index: 999;
  background-color: #ffffff;
  transition: all 0.3s;
  border-bottom: 1px solid #f5f5f5;

  > .content {
    display: flex;
    align-items: center;
    max-width: 1024px;
    padding: 16px;
    margin: 0 auto;
    box-sizing: border-box;

    .content {
      flex: 1 1 auto;
    }
    .prefix, .suffix {
      flex: 0 0 auto;
      position: relative;
    }

    .prefix {
      font-size: 18px;
    }

    .search {
      outline: none;
      border-radius: 40px;
      border: 1px solid #d5d5d5;
      background-color: #ffffff;
      height: 32px;
      width: 160px;
      padding: 0 16px;
      transition: all 0.3s;

      &:focus {
        border-color: #888888;
        width: 280px;

        &+ .icon-search {
          color: #333333;
        }
      }
    }

    .icon-search {
      position: absolute;
      right: 0;
      top: 50%;
      transform: translate3d(-50%, -50%, 0);
      color: #888888;
      transition: color 0.3s;
    }
  }
}
</style>