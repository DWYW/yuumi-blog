<template>
<footer class="page-footer">
  <div class="content">
    <div class="prefix">
      <div class="copyright">@COPYRIGHT 2019</div>
    </div>
    <div class="content"></div>
    <div class="suffix">
      <a class="github" :href="githubUrl" target="_blank">
        <span class="icon icon-github"></span>
      </a>
    </div>
  </div>
</footer>
</template>

<script lang="ts" setup>
const { GITHUB_HOME } = useRuntimeConfig()
const githubUrl = computed(() => GITHUB_HOME)
</script>

<style lang="scss" scoped>
.page-footer {
  color: #888888;
  background-color: #f5f5f5;
  > .content {
    display: flex;
    padding: 24px 16px;
    max-width: 1024px;
    margin: 0 auto;
    align-items: center;

    .content {
      flex: 1 1 auto;
    }
    .prefix {
      flex: 0 0 auto;
      font-size: 14px;
    }
    .suffix {
      flex: 0 0 auto;
    }
  }
}

.github {
  font-size: 24px;
}
</style>