<script setup lang="ts">
import { ref } from "vue"

const props = defineProps<{
  value: string
}>()

const emits = defineEmits(["compelete"])
const current = ref(props.value)

function onChange () {
  current.value = props.value
  emits("compelete")
}
</script>

<template>
<div :class="['score-board', { '_changing': current !== value }]">
  <div class="upper">
    <div class="front">{{value}}</div>
    <div class="back">{{current}}</div>
  </div>
  <div class="lower">
    <div class="front">{{current}}</div>
    <div class="back" @animationend="onChange">{{value}}</div>
  </div>
</div>
</template>

<style lang="scss" scoped>
$--duration: 0.8s;

.score-board {
  width: 30px;
  height: 60px;
  position: relative;
  font-size: 48px;
  background-color: inherit;

  &._changing {
    .upper .back {
      animation: flip-upper $--duration;
    }

    .lower .back {
      animation: flip-lower $--duration;
    }
  }
}

.upper, .lower {
  width: 100%;
  height: 50%;
  overflow: hidden;
  position: absolute;
  left: 0;
  background-color: inherit;
  transform-style: preserve-3d;

  .front, .back {
    perspective: 800px;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 200%;
    display: flex;
    justify-content: center;
    align-items: center;
    backface-visibility: hidden;
    background-color: inherit;
    transform-origin: center;
  }
}

.upper {
  top: 0;
}

.lower {
  bottom: 0;

  .front, .back {
    top: -100%;
  }

  .back {
    transform: rotateX(0.5turn);
  }
}

@keyframes flip-upper {
  from {
    transform: rotateX(0turn);
  }

  to {
    transform: rotateX(0.5turn);
  }
}

@keyframes flip-lower {
  from {
    transform: rotateX(0.5turn);
  }

  to {
    transform: rotateX(0turn);
  }
}
</style>
