<template>
<div class="not-found">
  <div class="number">
    <div class="number-item" num="4">4</div>
    <div class="number-item" num="0">0</div>
    <div class="number-item" num="4">4</div>
  </div>

  <div class="desc" v-t="'MSG.PAGE_NOT_FOUND'"></div>

  <div class="btn">
    <YuumiButton v-t="'MSG.GO_HOME'" @click="goHome"></YuumiButton>
  </div>
</div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router"
const router = useRouter()

function goHome() {
  router.replace("/")
}
</script>

<style lang="scss">
.not-found {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;

  .number {
    font-size: 108px;
    font-weight: bold;

    .number-item {
      position: relative;
      transform-style: preserve-3d;
      perspective: 100px;
      display: inline-block;

      &::before, &::after {
        content: attr(num);
        position: absolute;
        left: 0;
        top: 0;
        color: #aaaaaa;
      }

      &::before {
        transform: translate3d(4px, 0, -4px);
      }

      &::after  {
        transform: translate3d(8px, 0, -8px);
      }
    }
  }

  .btn {
    margin-top: 40px;
  }
}
</style>
