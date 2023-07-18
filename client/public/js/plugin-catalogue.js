(function (options) {
  function styles2string(styles) {
    var str = ""

    for (var key in styles) {
      str += key.replace(/([A-Z])/g, "-$1").toLowerCase() + ":" + styles[key] + ";"
    }

    return str
  }

  function getHeaderNodes(parentNode, tagNames) {
    tagNames = tagNames || ["h1", "h2", "h3", "h4", "h5", "h6"]
    return parentNode.querySelectorAll(tagNames.toString())
  }

  function createCatelogueNode(nodes) {
    var catelogue = document.createElement("ul")
    catelogue.setAttribute("id", "catelogue")
    if (options.listStyles) {
      catelogue.setAttribute("style", styles2string(options.listStyles))
    }

    var cache = [[-Infinity, 0 - options.space]]
    var prefix = "catelog_" + Date.now().toString(32)
    for (var i = 0; i < nodes.length; i++) {
      var level = Number(nodes[i].tagName.slice(1))
      var last
      while(cache.length > 0) {
        last = cache.pop()
        if (last[0] < level) {
          cache.push(last)
          break
        }
      }
      last = [level, last[1] + options.space]
      cache.push(last)

      // 绑定ID
      var nodeId = prefix + i
      nodes[i].setAttribute("id", nodeId)

      // 创建li
      var text = document.createElement("span")
      text.setAttribute("style", "cursor: pointer;")
      text.setAttribute("data-catelog", nodeId)
      text.innerText = nodes[i].innerText

      var li = document.createElement("li")
      if (options.listItemStyles) {
        li.setAttribute("style", styles2string(options.listItemStyles) + "padding-left:" + last[1] + "px;")
      }

      li.appendChild(text)
      catelogue.appendChild(li)
    }

    return catelogue
  }

  function render(node) {
    var siblingNode = document.body.querySelector(options.selector)
    siblingNode.parentNode.insertBefore(node, siblingNode)
  }

  function bindClickEvent(node, offset) {
    node.addEventListener("click", function (e) {
      if (e.target.tagName !== "SPAN") return
      var catelog = e.target.dataset.catelog
      var el = document.body.querySelector("#" + catelog)
      if (!el) return
      var top = el.getBoundingClientRect().top
      var scrollTop = document.body.scrollTop || document.documentElement.scrollTop
      document.body.scroll({ top: scrollTop + top - offset, behavior: "smooth" })
      document.documentElement.scroll({ top: scrollTop + top - offset, behavior: "smooth" })
    }, false)
  }

  var contentElement = options.selector ? document.body.querySelector(options.selector) : document.body
  var headerNodes = getHeaderNodes(contentElement)
  var catelogueNode = createCatelogueNode(headerNodes)
  render(catelogueNode)
  bindClickEvent(catelogueNode, document.body.querySelector(options.headerSelector).getBoundingClientRect().height)
})({
  selector: ".article-content",
  headerSelector: ".page-header",
  listStyles: {
    paddingTop: "24px"
  },
  listItemStyles: {
    listStyle: "inside",
    color: "#576b95",
    padding: "6px 0"
  },
  space: 14
});