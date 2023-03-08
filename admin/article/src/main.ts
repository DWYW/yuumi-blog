import "./public-path"
import { render } from "@/micro-application"

if (!window.__POWERED_BY_QIANKUN__) {
  render()
}

export { bootstrap, mount, unmount } from "@/micro-application"
