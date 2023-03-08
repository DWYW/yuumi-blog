<template>
<section id="markdown" v-html="contentHTML"></section>
</template>

<script lang="ts" setup>
import MarkdownIt from 'markdown-it'
import hljs from 'highlight.js'
import { defineProps, computed, watch } from "vue"

const emit = defineEmits(["translated"])
const props = defineProps({
  content: { type: String, default: "" }
})

// 安全防护
const safetyContent = computed((): string => {
  return props.content
    .replace(/<script[^>]*?\>(\s|\S)*?<\/script>/ig, "")
    .replace(/<style[^>]*?\>(\s|\S)*?<\/style>/ig, "")
})

const contentHTML = computed(() => {
  const html = new MarkdownIt({
    html: true,
    langPrefix: 'hljs language-',
    highlight: (text: string, lang: string) => {
      return lang ? hljs.highlight(lang, text, true).value : hljs.highlightAuto(text).value
    }
  }).render(safetyContent.value)

  return html
})

watch(() => contentHTML.value, () => emit("translated"))
</script>

<style lang="scss">
@import "@/assets/styles/md.scss";
@import "@/assets/styles/hljs.scss";

@media screen and (max-width: 640px) {
  .markdown-content img {
    max-width: 100%;
  }
}
</style>
