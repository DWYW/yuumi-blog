import { $helper } from "@/common"
import { getCurrentInstance, Ref, ref } from "vue"
import { useRoute, useRouter } from "vue-router"

export function useTableHelper(loadTableData: () => undefined|Promise<{list: any[], pagination: any}>) {
  if (!getCurrentInstance()) throw new Error("please use in `setup`")
  const route = useRoute()
  const router = useRouter()

  const list: Ref<any[]> = ref([])
  const pagination: Ref<{
    page: string
    pageTotal: string
    total: string
  }> = ref({
    page: $helper.getValueByPath<string>(route.query, "page", "1"),
    pageTotal: "0",
    total: "0"
  })

  function tableReload() {
    loadTableData()?.then((res) => {
      list.value = res.list
      pagination.value.page = res.pagination.page
      pagination.value.total = res.pagination.total
      pagination.value.pageTotal = res.pagination.page_total
      router.replace({ path: route.path, query: Object.assign({}, route.query, { page: pagination.value.page }) })
      return res
    })
  }

  function onPaginationChange(page: number) {
    if (page.toString() === pagination.value.page) return
    pagination.value.page = page.toString()
    tableReload()
  }

  return {
    list,
    pagination,
    tableReload,
    onPaginationChange
  }
}
