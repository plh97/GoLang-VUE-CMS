<script setup>
import { defineRouteMeta, useRouter } from '@fesjs/fes'
import { FButton, FEllipsis, FForm, FFormItem, FMessage, FModal, FOption, FPagination, FSelect, FSpace, FTable, FTableColumn } from '@fesjs/fes-design'
import { LoadingOutlined } from '@fesjs/fes-design/icon'
import { computed, reactive, ref, watch } from 'vue'
import { useRequest } from 'vue-hooks-plus'
import { request } from '@/api'
import { useValidator } from '@/common/hooks'
import { formatTimestamp } from '@/common/utils'
import { ACTIVITY_STATUS } from '@/enums'

defineRouteMeta({
  name: 'activity/list',
  title: '活动管理',
})

const state = reactive({
  deleteIDs: [],
  statusModalVisible: false,
  status: 0,
  current_page: 1,
  page_size: 10,
})

const changeActivityStatusModelFormRef = ref()

const { loading, data, run: getActivityList } = useRequest(() => request('/activity/list', {
  page: {
    current_page: state.current_page,
    page_count: 0,
    page_size: state.page_size, // 获取全部活动
    total: 0,
  },
  sort: {
    field: 'create_time',
    type: 'asc',
  },
}, {
  method: 'POST',
}))

const loadingOnce = ref(loading.value)
watch(
  () => loading.value,
  val => loadingOnce.value = val,
  { once: true },
)
const router = useRouter()
// const action = [
//   {
//     label: '查看详情',
//     func: (row) => {
//       router.push(`/activity/${row.id}`)
//     },
//   },
//   {
//     label: '编辑',
//     func: (row) => {
//       router.push(`/activity/edit/${row.id}`)
//     },
//   },
// ]
function handleCreate() {
  router.push(`/activity/create`)
}
function handleDelete() {
  FModal.confirm({
    title: '删除',
    content: `确定删除你所选择的活动？`,
    okText: '确定',
    maskClosable: true,
    onOk() {
      request('/activity/delete', {
        ids: state.deleteIDs,
      }, {
        method: 'DELETE',
      }).then(() => {
        FMessage.success({
          content: '删除成功',
        })
        data.value.list = data.value.list.filter((item) => {
          return !state.deleteIDs.includes(item.id)
        })
        state.deleteIDs = []
      })
    },
  })
}

async function handleChangeStatus() {
  await changeActivityStatusModelFormRef.value?.validate()
  request('/activity/batch_update', {
    ids: state.deleteIDs,
    status: state.status,
  }, {
    method: 'post',
  }).then(() => {
    FMessage.success({
      content: '修改成功',
    })
    state.statusModalVisible = false
    data.value.list = data.value.list.map((item) => {
      if (state.deleteIDs.includes(item.id)) {
        return {
          ...item,
          status: state.status,
        }
      }
      return item
    })
    state.deleteIDs = []
  })
}
const validator = useValidator(state)

const calculateTableHeight = computed(() => {
  // Get viewport height
  const viewportHeight = window.innerHeight
  return viewportHeight - 315
})
function handleChange(page, pageSize) {
  state.page_size = pageSize
  state.current_page = page
  getActivityList()
}
</script>

<template>
  <!-- mul-change activity status modal -->
  <FModal v-model:show="state.statusModalVisible" title="更改活动状态" display-directive="show" @ok="handleChangeStatus">
    <FForm
      ref="changeActivityStatusModelFormRef" :model="state" label-position="right" :label-width="100" :span="12"
      align="flex-start"
    >
      <FFormItem label="状态" prop="status" :rules="[{ required: true, type: 'string', message: '请选择状态', validator }]">
        <FSelect v-model="state.status">
          <FOption
            v-for="(id) in Object.keys(ACTIVITY_STATUS).filter(key => isNaN(+(ACTIVITY_STATUS[key])))" :key="id"
            :value="+id"
          >
            {{ ACTIVITY_STATUS[id] }}
          </FOption>
        </FSelect>
      </FFormItem>
    </FForm>
  </FModal>
  <!-- mul-change activity status modal -->
  <nav>
    <h1>
      活动管理
    </h1>
    <FSpace>
      <FButton type="primary" @click="handleCreate">
        新增活动
      </FButton>
      <FButton :disabled="!(state.deleteIDs?.length > 0)" @click="handleDelete">
        删除活动
      </FButton>
      <FButton :disabled="!(state.deleteIDs?.length > 0)" @click="state.statusModalVisible = true">
        更改状态
      </FButton>
    </FSpace>
  </nav>
  <div v-if="loading" class="loading">
    <LoadingOutlined class="icon" />
  </div>
  <FTable
    v-show="!loading" v-model:checked-keys="state.deleteIDs" always-scrollbar class="table"
    :height="calculateTableHeight" size="small" row-key="id" :data="data?.list"
  >
    <FTableColumn type="selection" :width="30" fixed="left" />
    <FTableColumn prop="title" label="活动标题" fixed="left">
      <template #default="{ row }">
        <FEllipsis :content="row.title" style="max-width: 100%" />
      </template>
    </FTableColumn>
    <FTableColumn prop="location" label="活动地点" fixed="left">
      <template #default="{ row }">
        <FEllipsis :content="row.location" style="max-width: 100%" />
      </template>
    </FTableColumn>
    <FTableColumn prop="start_time" label="活动开始时间" :width="170">
      <template #default="{ row }">
        {{ formatTimestamp(row.start_time * 1000) }}
      </template>
    </FTableColumn>
    <FTableColumn label="活动费用" :min-width="160">
      <template #default="{ row }">
        <div>
          日元：男{{ row.man_price_jp }} 女{{ row.woman_price_jp }}
        </div>
        <div>
          人民币：男{{ row.man_price }} 女{{ row.woman_price }}
        </div>
      </template>
    </FTableColumn>
    <FTableColumn label="人数上限">
      <template #default="{ row }">
        {{ row.man_sign_up_limit + row.woman_sign_up_limit }}
      </template>
    </FTableColumn>
    <FTableColumn label="虚拟人数">
      <template #default="{ row }">
        {{ row.man_sign_up_fake_count + row.woman_sign_up_fake_count }}
      </template>
    </FTableColumn>
    <FTableColumn prop="sign_up_count" label="预报名人数">
      <template #default="{ row }">
        <FButton type="link" @click="() => router.push(`/activity/${row.id}/signup`)">
          {{ row.sign_up_uncompleted_count }}
        </FButton>
      </template>
    </FTableColumn>
    <FTableColumn prop="sign_up_count" label="已报名人数">
      <template #default="{ row }">
        <FButton type="link" @click="() => router.push(`/activity/${row.id}/signup`)">
          {{ row.sign_up_completed_count }}
        </FButton>
      </template>
    </FTableColumn>
    <FTableColumn prop="sign_up_canceled_count" label="已取消人数" />
    <FTableColumn prop="status" label="活动状态">
      <template #default="{ row }">
        {{ ACTIVITY_STATUS[row.status] }}
      </template>
    </FTableColumn>
    <FTableColumn prop="start_time" label="活动创建时间" sortable :width="170">
      <template #default="{ row }">
        {{ formatTimestamp(row.create_time * 1000) }}
      </template>
    </FTableColumn>
    <FTableColumn prop="end_time" label="活动编辑时间" sortable :width="170">
      <template #default="{ row }">
        {{ formatTimestamp(row.update_time * 1000) }}
      </template>
    </FTableColumn>
    <FTableColumn label="操作" align="left" fixed="right" :width="170">
      <template #default="{ row }">
        <FButton type="link" @click="router.push(`/activity/${row.id}`)">
          查看详情
        </FButton>
        <FButton v-if="row.status === ACTIVITY_STATUS.招募中" type="link" @click="router.push(`/activity/edit/${row.id}`)">
          编辑
        </FButton>
      </template>
    </FTableColumn>
  </FTable>
  <FPagination
    v-if="!loadingOnce" show-total class="pagination" :current-page="state.current_page"
    :total-count="data?.page?.total" show-size-changer
    show-quick-jumper
    @change="handleChange"
  />
</template>

<style scoped>
.loading {
  width: 100%;
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;

  .icon {
    margin-top: 20px;
    font-size: 40px;
  }
}

.pagination {
  margin-top: 10px;
  align-self: center;
}

nav {
  margin-bottom: 20px;
}

.table {
  flex: 1;
}
</style>
