<template>
<div class="ratio-rect" :style="styleValue">
  <div class="rect-body" v-bind="$attrs">
    <slot></slot>
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from "vue"

export default defineComponent({
  inheritAttrs: false,
  props: {
    ratio: { type: Number, default: 1 }
  },
  computed: {
    styleValue(): {[key:string]: string} {
      return {
        paddingTop: `${(1 / this.ratio * 100).toFixed(4)}%`
      }
    }
  }
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