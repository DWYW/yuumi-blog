(function (global) {
  function Scroll2Top(options) {
    this.options = Object.assign({
      threshold: 300,
      size: "48px",
      right: "16px",
      bottom: "84px",
      borderRadius: "8px",
      borderColor: "#cccccc",
      background: "rgba(255, 255, 255, .6)",
      fontSize: "12px",
      text: "顶部"
    }, options)


    this.$el = this.createBtn()
    this.onScroll = this.scrollEvent.bind(this)
    this.onClick = this.clickEvent.bind(this)
    document.body.appendChild(this.$el)
    window.addEventListener("scroll", this.onScroll, false)
    this.$el.addEventListener("click", this.onClick, false)
  }

  Scroll2Top.prototype.styles2string = function (styles) {
    var str = ""
    for (var key in styles) {
      str += key.replace(/([A-Z])/g, "-$1").toLowerCase() + ":" + styles[key] + ";"
    }
    return str
  }

  Scroll2Top.prototype.createBtn = function () {
    var arrow = document.createElement("div")
    arrow.setAttribute("style", this.styles2string({
      width: "30%",
      paddingTop: "30%",
      borderTop: "1px solid " + this.options.borderColor,
      borderLeft: "1px solid " + this.options.borderColor,
      transform: "rotate(45deg)",
      transformOrigin: "top left",
      position: "absolute",
      top: "20%",
      left: "50%"
    }))

    var text = document.createElement("p")
    text.innerText = this.options.text

    var btn = document.createElement("div")
    btn.setAttribute("style", this.styles2string({
      position: "fixed",
      right: this.options.right,
      bottom: this.options.bottom,
      border: "1px solid " + this.options.borderColor,
      borderRadius: this.options.borderRadius,
      width: this.options.size,
      height: this.options.size,
      background: this.options.background,
      cursor: "pointer",
      fontSize: this.options.fontSize,
      textAlign: "center",
      paddingTop: "24px",
      display: "none"
    }))

    btn.appendChild(arrow)
    btn.appendChild(text)
    return btn
  }

  Scroll2Top.prototype.isVisible = false
  Scroll2Top.prototype.showBtn = function () {
    if (this.isVisible) return
    this.isVisible = true
    this.$el.style.setProperty("display", "block")

    setTimeout(function () {
      this.$el.style.setProperty("opacity", "1")
      this.$el.style.setProperty("transition", "opacity 0.3s ease")
    }.bind(this), 16)
  }

  Scroll2Top.prototype.hideBtn = function () {
    if (!this.isVisible) return
    this.isVisible = false
    this.$el.style.setProperty("opacity", "0")

    setTimeout(function () {
      this.$el.style.setProperty("display", "none")
      this.$el.style.removeProperty("opacity")
      this.$el.style.removeProperty("transition")
    }.bind(this), 300)
  }

  Scroll2Top.prototype.direction = ""
  Scroll2Top.prototype.prevScrollTop = 0
  Scroll2Top.prototype.timeout = null

  Scroll2Top.prototype.scrollEvent = function () {
    if (this.timeout !== null) {
      clearTimeout(this.timeout)
    }

    this.timeout = setTimeout(function () {
      var scrollTop = window.document.documentElement.scrollTop || window.document.body.scrollTop
      this.direction = scrollTop > this.prevScrollTop ? "down" : "top"
      this.prevScrollTop = scrollTop
      this.timeout = null
      this.direction === "down" && scrollTop > this.options.threshold && this.showBtn()
      this.direction === "top" && scrollTop < this.options.threshold && this.hideBtn()
    }.bind(this), 100)
  }

  Scroll2Top.prototype.clickEvent = function () {
    document.body.scroll({ top: 0, behavior: "smooth" })
    document.documentElement.scroll({ top: 0, behavior: "smooth" })
  }

  Scroll2Top.prototype.destroy = function() {
    var ins = global.__scroll2top
    window.removeEventListener("scroll", ins.onScroll, false)
    document.body.removeChild(ins.$el)

    if (global.__scroll2top) {
      delete global.__scroll2top
    }
  }

  if (!global.__scroll2top) {
    global.__scroll2top = new Scroll2Top()
  }
})(window, {});