<template>
<div class="ratio-rect" :style="{paddingTop: paddingTop}">
  <div class="rect-body">
    <slot></slot>
  </div>
</div>
</template>

<script lang="ts" setup>
import { computed, defineProps } from "vue"

const props = defineProps({
  ratio: { type: Number, default: 1 }
})

const paddingTop = computed(() => {
  return `${props.ratio * 100}%`
})
</script>

<style lang="scss" scoped>
.ratio-rect {
  position: relative;

  .rect-body {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
  }
}
</style>
