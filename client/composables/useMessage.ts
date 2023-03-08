import { render, Transition, withDirectives, vShow } from "vue"
import { useVisitorStore } from "~/stores/visitor"

interface ThemeMessageOption {
  message: string
  duration?: number
}

interface MessageOption extends ThemeMessageOption {
  theme?: "default"|"warn"|"success"|"error"
}

const MessagePanel = defineComponent({
  name: "MessagePanel",
  emits: ["complete"],
  props: {
    duration: { type: Number, default: 3000 },
    message: { type: String, default: "" },
    theme: { type: String, default: "success" }
  },
  data() {
    return {
      visible: false
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.visible = true
    })

    window.setTimeout(() => {
      this.visible = false
    }, this.duration)
  },
  render() {
    return h(Transition, {
      name: 'yuumi-message',
      appear: true,
      onAfterLeave: () => this.$emit("complete")
    }, {
      default: () => withDirectives(
        h("div", {
          class: ["message-panel", "theme__" + this.theme]
        }, this.message),
        [[vShow, this.visible]]
      )
    })
  }
})

export const useMessage =() => {
  function Message(option: MessageOption) {
    const el = document.createElement("div")
    el.classList.add("yuumi-message")
    document.body.appendChild(el)

    const conf = Object.assign({}, option)
    render(h(MessagePanel, {
      message: conf.message,
      duration: conf.duration,
      theme: conf.theme,
      onComplete: () => {
        el.parentElement?.removeChild(el)
      }
    }), el)
  }

  return {
    message: Message,
    warnMessage: (option: ThemeMessageOption) =>  Message(Object.assign({theme: "warn"}, option) as MessageOption),
    successMessage: (option: ThemeMessageOption) =>  Message(Object.assign({theme: "success"}, option) as MessageOption),
    errorMessage: (option: ThemeMessageOption) =>  Message(Object.assign({theme: "error"}, option) as MessageOption),
    catchMessage: (err: any) => {
      const { $pinia } = useNuxtApp()
      const { signout } = useVisitorStore($pinia)
      if (err.status === 401) {
        Message({theme: "default", message: "请登录后操作"})
        signout()
      } else {
        Message({theme: "error", message: err.data ? err.data.message : err.message })
      }
    }
  }
}