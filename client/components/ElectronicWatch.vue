<script setup lang="ts">
import { ref, computed, ComputedRef, onMounted } from "vue"

const year = ref()
const month = ref(0)
const date = ref(0)
const hours = ref(0)
const minutes = ref(0)
const seconds = ref(0)
const day = ref(0)

const dateString: ComputedRef<string> = computed(() => {
  return `${year.value} 年 ${("0" + month.value).slice(-2)} 月 ${("0" + date.value).slice(-2)} 日`
})

const hourHigher: ComputedRef<string> = computed(() => {
  return ("0" + hours.value).slice(-2, -1)
})
const hourLower: ComputedRef<string> = computed(() => {
  return ("0" + hours.value).slice(-1)
})

const minuteHigher: ComputedRef<string> = computed(() => {
  return ("0" + minutes.value).slice(-2, -1)
})
const minuteLower: ComputedRef<string> = computed(() => {
  return ("0" + minutes.value).slice(-1)
})

const secondHigher: ComputedRef<string> = computed(() => {
  return ("0" + seconds.value).slice(-2, -1)
})
const secondLower: ComputedRef<string> = computed(() => {
  return ("0" + seconds.value).slice(-1)
})

const dayString: ComputedRef<string> = computed(() => {
  const days = ["周日", "周一", "周二", "周三", "周四", "周五", "周六"]
  return days[day.value]
})

onMounted(() => {
  window.setInterval(update, 1000)
})

function update() {
  const now = new Date()
  year.value = now.getFullYear()
  month.value = now.getMonth() + 1
  date.value = now.getDate()
  hours.value = now.getHours()
  minutes.value = now.getMinutes()
  seconds.value = now.getSeconds()
  day.value = now.getDay()
}

update()
</script>

<template>
<div class="electronic-watch">
  <div class="watch-panel">
    <div class="numbers">
      <div class="hour">
        <ScoreBoard :value="hourHigher"></ScoreBoard>
        <ScoreBoard :value="hourLower"></ScoreBoard>
      </div>
      <div class="split"></div>
      <div class="minute">
        <ScoreBoard :value="minuteHigher"></ScoreBoard>
        <ScoreBoard :value="minuteLower"></ScoreBoard>
      </div>
      <div class="split"></div>
      <div class="second">
        <ScoreBoard :value="secondHigher"></ScoreBoard>
        <ScoreBoard :value="secondLower"></ScoreBoard>
      </div>
    </div>

    <div class="other">
      <div class="date">{{dateString}}</div>
      <div class="day">{{dayString}}</div>
    </div>
  </div>
</div>
</template>

<style lang="scss" scoped>
$background-color: #f0f0f0;
.electronic-watch {
  padding: 20px;
  .watch-panel {
    width: 280px;
    height: 120px;
    display: flex;
    flex-direction: column;
    border-radius: 16px;
    justify-content: center;
    align-items: center;
    background-color: $background-color;
    box-shadow:
      -6px -6px 12px rgba(0, 0, 0, 0.2),
      6px 6px 12px rgba(255, 255, 255, 1);
  }

  .numbers {    flex: 1 1 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 10px;

    .hour, .minute, .second {
      display: flex;
      background-color: $background-color;
    }

    .split {
      padding: 0 8px;
      &::before, &::after {
        content: " ";
        display: block;
        width: 4px;
        height: 4px;
        border-radius: 50%;
        background-color: #000000;
      }

      &::before {
        transform: translate3d(0, -5px, 0);
      }

      &::after {
        transform: translate3d(0, 5px, 0);
      }
    }
  }

  .other {
    flex: 0 0 auto;
    padding: 0 20px 10px;

    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
</style>
