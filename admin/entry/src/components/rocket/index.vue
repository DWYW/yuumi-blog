<template>
<div class="rocket">
  <div class="rocket-body"></div>
  <div class="rocket-hole"></div>
  <div class="rocket-fins _left"></div>
  <div class="rocket-fins _right"></div>
  <div class="rocket-fire"></div>
</div>
</template>

<script lang="ts">
import { defineComponent } from "vue"
export default defineComponent({})
</script>

<style lang="scss" scoped>
.rocket {
  position: relative;
  width: 100px;
  height: 280px;

  .rocket-body {
    width: 100%;
    height: 100%;
    background-color: #ffffff;
    border-top-left-radius: 50%;
    border-top-right-radius: 50%;
    border-bottom-left-radius: 30%;
    border-bottom-right-radius: 30%;
    position: absolute;
    z-index: 2;
    box-shadow: 0px 10px 30px #888888 inset;
  }

  .rocket-hole {
    width: 50%;
    position: absolute;
    left: 50%;
    z-index: 3;
    transform: translate3d(-50%, 200%, 0);

    &::before {
      display: block;
      content: "";
      width: 100%;
      padding-top: 100%;

      border-radius: 50%;
      box-shadow: 0 0 0 8px #aaaaaa inset;
    }
  }

  .rocket-fins {
    position: absolute;
    bottom: 0;
    width: 50%;
    height: 50%;
    background-color: #ffffff;

    &._left {
      transform: skewY(-35deg) translate3d(-90%, 0 ,0);
      left: 0;

      border-top-left-radius: 50%;
      border-top-right-radius: 50%;
      border-bottom-left-radius: 20%;
      border-bottom-right-radius: 20%;
      box-shadow: -10px 0px 10px #aaaaaa inset;
    }

    &._right {
      transform: skewY(35deg) translate3d(90%, 0 ,0);
      right: 0;

      border-top-left-radius: 50%;
      border-top-right-radius: 50%;
      border-bottom-left-radius: 20%;
      border-bottom-right-radius: 20%;
      box-shadow: 10px 0px 10px #aaaaaa inset;
    }
  }

  .rocket-fire {
    width: 50%;
    position: absolute;
    left: 50%;
    bottom: 0;
    transform: translate3d(-50%, 110%, 0);

    &::before, &::after {
      display: inline-block;
      content: "";
      width: 50%;
      padding-top: 100%;
      background-color: #ffffff;
    }

    &::before {
      border-top-left-radius: 10%;
      border-bottom-left-radius: 100%;
      box-shadow: 10px 10px 20px #aaaaaa inset;
    }

    &::after {
      border-top-right-radius: 10%;
      border-bottom-right-radius: 100%;
      box-shadow: -10px 10px 20px #aaaaaa inset;
    }
  }
}
</style>
