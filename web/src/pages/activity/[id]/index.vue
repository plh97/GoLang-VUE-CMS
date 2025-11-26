<script setup>
import { defineRouteMeta, useRoute, useRouter } from '@fesjs/fes'
import { FButton, FDivider } from '@fesjs/fes-design'
import { useRequest } from 'vue-hooks-plus'
import { request } from '@/api'
import activityForm from '@/components/activityForm.vue'

defineRouteMeta({
  name: 'activity/detail',
  title: '活动管理',
})

const { replace } = useRouter()
const { params } = useRoute()
const { loading, data } = useRequest(() => request('/activity/detail', {
  activity_id: +params.id,
}, {
  method: 'POST',
}))
function handleBack() {
  replace('/activity')
}
</script>

<template>
  <nav>
    <h1>
      查看详情
    </h1>
    <FButton type="primary" @click="handleBack">
      返回
    </FButton>
  </nav>
  <FDivider />
  <activityForm :data="data?.item" readonly :loading="loading" />
</template>

<style scoped>
nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
</style>
