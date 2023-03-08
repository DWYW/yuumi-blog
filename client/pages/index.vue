<template>
<div class="screen">
  <header class="header">
    <div class="logo">
      <div class="letter" v-for="item in 'DWYW'" :key="item">
        <span>{{item}}</span>
      </div>
    </div>

    <div class="slogn"> NO MATTER WHO YOU ARE AND YOU CAN DO WATHEVER YOU WANT</div>
  </header>
  <main class="body">
    <div class="entries">
      <NuxtLink to="/article/list">随笔</NuxtLink>
    </div>

    <div class="dest">
      <div class="ornament">
        <ElectronicWatch class="watch"></ElectronicWatch>
        <img class="plant" src="@/assets/images/flower1.png">
        <img class="plant" src="@/assets/images/flower3.png">
      </div>
    </div>
  </main>
  <footer class="footer"></footer>
</div>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped>

.screen {
  width: 100%;
  min-height: 100vh;

  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 32px 0;
  box-sizing: border-box;
  background-color: #f0f0f0;

  .header, .footer, .body {
    flex: 0 0 auto;
  }
}

.logo {
  max-width: 600px;
  margin: 0 auto;
  display: flex;
  justify-content: space-around;

  .letter {
    font-size: 100px;
    font-weight: bold;
    transform: scale(1, 1.2);
    position: relative;

    span {
      position: relative;

      text-shadow: -2px -2px 6px rgb(255, 255, 255), 4px 4px 6px rgb(0, 0, 0);
    }
  }
}

.slogn {
  margin: 0 auto;
  text-align: center;
  font-weight: bold;
  padding: 10px 0;
}

.entries {
  padding: 16px 0 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.dest {
  max-width: 768px;
  margin: 0 auto;
  position: relative;
  overflow: hidden;

  .ornament {
    position: relative;
    z-index: 1;
    display: flex;
    align-items: flex-end;
    justify-content: space-between;

    .plant {
      width: 160px;
    }
  }
}

@media screen and (max-width: 600px) {
  .logo {
    width: 86%;
    .letter {
      font-size: 64px;
    }
  }

  .slogn {
    width: 86%;
  }
  .dest .ornament {
    flex-wrap: wrap;
    justify-content: center;

    .watch {
      margin: 63px 0 0;
    }
  }
}
</style>