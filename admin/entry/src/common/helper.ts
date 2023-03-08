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
 * 将普通数组转化为树形结构数组
 */
export function list2tree(list: any[] = [], options?: {[x: string]:string}): any[] {
  const keys = Object.assign({
    id: "id",
    parentId: "parent_id",
    parent: "parent",
    children: "children"
  }, options)

  let i = -1

  while (++i < list.length) {
    const parentId = list[i][keys.parentId]
    if (!parentId) continue

    const parent = list.find((item) => item[keys.id] === parentId)
    if (!parent) continue
    list[i][keys.parent] = parent

    const children: any[] = parent[keys.children] || (parent[keys.children] = [])

    if (!children.find((item) => item[keys.id] === list[i][keys.id])) {
      children.push(list[i])
    }
  }

  return list.filter(item => item[keys.parentId] === "0" || !item[keys.parentId])
}
