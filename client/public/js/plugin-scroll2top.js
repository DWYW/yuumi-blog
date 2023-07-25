(function(options){
  options = Object.assign({
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

  function styles2string(styles) {
    var str = ""
    for (var key in styles) {
      str += key.replace(/([A-Z])/g, "-$1").toLowerCase() + ":" + styles[key] + ";"
    }
    return str
  }

  function createBtn() {
    var arrow = document.createElement("div")
    arrow.setAttribute("style", styles2string({
      width: "30%",
      paddingTop: "30%",
      borderTop: "1px solid " + options.borderColor,
      borderLeft: "1px solid " + options.borderColor,
      transform: "rotate(45deg)",
      transformOrigin: "top left",
      position: "absolute",
      top: "20%",
      left: "50%"
    }))

    var text = document.createElement("p")
    text.innerText = options.text

    var btn = document.createElement("div")
    btn.setAttribute("style", styles2string({
      position: "fixed",
      right: options.right,
      bottom: options.bottom,
      border: "1px solid " + options.borderColor,
      borderRadius: options.borderRadius,
      width: options.size,
      height: options.size,
      background: options.background,
      cursor: "pointer",
      fontSize: options.fontSize,
      textAlign: "center",
      paddingTop: "24px",
      display: "none"
    }))

    btn.appendChild(arrow)
    btn.appendChild(text)
    return btn
  }

  function mounte(element) {
    document.body.appendChild(element)
  }

  function bindEvents(element) {
    var direction = ""
    var prevScrollTop = 0
    var timeout = null

    window.addEventListener("scroll", function() {
      if (timeout !== null) {
        clearTimeout(timeout)
      }

      timeout = setTimeout(function() {
        var scrollTop = window.document.documentElement.scrollTop || window.document.body.scrollTop
        direction = scrollTop > prevScrollTop ? "down" : "top"
        prevScrollTop = scrollTop
        timeout = null
        direction === "down" && scrollTop > options.threshold && showBtn(element)
        direction === "top" && scrollTop < options.threshold && hideBtn(element)
      }, 100)
    })

    element.addEventListener("click", function() {
      document.body.scroll({ top: 0, behavior: "smooth" })
      document.documentElement.scroll({ top: 0, behavior: "smooth" })
    })
  }


  var isVisible = false

  function showBtn(element) {
    if (isVisible) return
    isVisible = true
    element.style.setProperty("display", "block")

    setTimeout(function() {
      element.style.setProperty("opacity", "1")
      element.style.setProperty("transition", "opacity 0.3s ease")
    }, 16)
  }

  function hideBtn(element) {
    if (!isVisible) return
    isVisible = false
    element.style.setProperty("opacity", "0")

    setTimeout(function() {
      element.style.setProperty("display", "none")
      element.style.removeProperty("opacity")
      element.style.removeProperty("transition")
    }, 300)
  }

  var btnElement = createBtn()
  mounte(btnElement)
  bindEvents(btnElement)
})();