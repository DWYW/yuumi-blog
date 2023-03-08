import { defineNuxtPlugin } from '#app'

export default defineNuxtPlugin(({ $config }) => {
  if (isExisted() && $config.public.UA) {
    registor($config.public.UA)
  }
})

function isExisted() {
  const n = document.querySelector("script#ga")
  return n !== null || n !== undefined
}

function registor(key: string) {
  const source = document.createElement("script")
  source.id = "ga"
  source.src = "https://www.googletagmanager.com/gtag/js?id="+key

  const runing = document.createElement("script")
  runing.innerHTML = "window.dataLayer = window.dataLayer || []; function gtag(){dataLayer.push(arguments);} gtag('js', new Date()); gtag('config', '"+key+"');"

  const fragment = document.createDocumentFragment()
  fragment.appendChild(source)
  fragment.appendChild(runing)
  document.body.appendChild(fragment)
}
