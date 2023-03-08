/**
 * 根据路径获取属性值
 */
export function getValueByPath<T>(data: any, path: string, resolve?: T): T {
  if (!path) return resolve !== undefined ? resolve : undefined as any

  const attrs = path.match(/\w+/g) || []
  let index = 0

  while (index < attrs.length && data !== null && data !== undefined) {
    data = data[attrs[index++]]
  }

  data = index && index === attrs.length ? data : undefined
  return (data === null || data === undefined) && resolve !== undefined ? resolve : data
}

export function getTreeNode(validator: (node: any) => boolean, treeData: any[]): any {
  if (!treeData.length) return

  let target
  for (let i = 0; i < treeData.length; i++) {
    const item = treeData[i]
    if (validator(item)) {
      target = item
    } else if (item.children) {
      target = getTreeNode(validator, item.children)
    }

    if (target) break
  }

  return target
}

/**
 * 深拷贝
 */
export function deepCopy(data: any) {
  if (!data || !(data instanceof Object) || data instanceof Function) return data

  const copy = data instanceof Array ? [] : {}
  const copyed = [{
    origin: data,
    copy: copy
  }]

  const foreach = (stackItem: any) => {
    const { origin, copy } = stackItem

    // 依次访问本层属性
    for (const key in origin) {
      if (!Object.prototype.hasOwnProperty.call(origin, key)) continue
      const item = origin[key]

      if (item instanceof Date || item instanceof Function) {
        copy[key] = item
        continue
      }

      if (!(item instanceof Object)) {
        copy[key] = item
        continue
      }

      // 检测当前值是否被拷贝过
      const _copyed = copyed.find(t => t.origin === item)
      if (_copyed) {
        copy[key] = _copyed.copy
        continue
      }

      copy[key] = item instanceof Array ? [] : {}

      // 用于下一次判断，防止循环引用造成的溢出
      copyed.push({
        origin: item,
        copy: copy[key]
      })

      foreach({
        origin: item,
        copy: copy[key]
      })
    }
  }

  foreach({
    origin: data,
    copy: copy
  })

  return copy
}

/**
 * 将普通数组转化为树形结构数组
 */
export function list2tree(data: any[] = [], options?: {[x: string]:string}): any[] {
  const keys = Object.assign({
    id: "id",
    parentId: "parent_id",
    parent: "parent",
    children: "children"
  }, options)

  const list: any[] = deepCopy(data)
  let i = -1
  while (++i < list.length) {
    const parentId = list[i][keys.parentId]
    if (!parentId) continue

    const parent = list.find((item) => item[keys.id] === parentId)
    if (!parent) continue
    list[i][keys.parent] = parent

    const children: any[] = parent[keys.children] || (parent[keys.children] = [])
    children.push(list[i])
  }

  return list.filter(item => item[keys.parentId] === "0" || !item[keys.parentId])
}

export function treeNodeDiff(oldNodes: any[], newNodes: any[]) {
  oldNodes = oldNodes || []
  newNodes = newNodes || []
  let oldStartIndex = 0
  let oldEndIndex = oldNodes.length - 1
  let newStartIndex = 0
  let newEndIndex = newNodes.length - 1
  let oldStartNode = oldNodes[oldStartIndex]
  let oldEndNode = oldNodes[oldEndIndex]
  let newStartNode = newNodes[newStartIndex]
  let newEndNode = newNodes[newEndIndex]

  const add = []
  const remove = []

  while (oldStartIndex <= oldEndIndex && newStartIndex <= newEndIndex) {
    if (oldStartNode.id === newStartNode.id) {
      oldStartNode = oldNodes[++oldStartIndex]
      newStartNode = newNodes[++newStartIndex]
    } else if (oldEndNode.id === newEndNode.id) {
      oldEndNode = oldNodes[--oldEndIndex]
      newEndNode = newNodes[--newEndIndex]
    } else if (oldStartNode.id === newEndNode.id) {
      oldStartNode = oldNodes[++oldStartIndex]
      newEndNode = newNodes[--newEndIndex]
    } else if (oldEndNode.id === newStartNode.id) {
      oldEndNode = oldNodes[--oldEndIndex]
      newStartNode = newNodes[++newStartIndex]
    } else {
      add.push(newStartNode)
      newStartNode = newNodes[++newStartIndex]
    }
  }

  while (oldStartIndex <= oldEndIndex) {
    remove.push(oldStartNode)
    oldStartNode = oldNodes[++oldStartIndex]
  }

  while (newStartIndex <= newEndIndex) {
    add.push(newStartNode)
    newStartNode = newNodes[++newStartIndex]
  }

  return {
    remove,
    add
  }
}
