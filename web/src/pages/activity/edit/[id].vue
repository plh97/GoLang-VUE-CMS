<script setup>
import { defineRouteMeta, useRoute, useRouter } from '@fesjs/fes'
import { FButton, FDivider, FMessage } from '@fesjs/fes-design'
import { useRequest } from 'vue-hooks-plus'
import { request } from '@/api'
import activityForm from '@/components/activityForm.vue'

defineRouteMeta({
  name: 'activity/edit',
  title: '编辑活动',
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
function onSubmit(data) {
  request('/activity/update', data, {
    method: 'POST',
  }).then(() => {
    FMessage.success({
      content: '修改成功',
    })
    handleBack()
  })
}
</script>

<template>
  <nav>
    <h1>
      编辑
    </h1>
    <FButton type="primary" @click="handleBack">
      返回
    </FButton>
  </nav>
  <FDivider />
  <activityForm :on-submit="onSubmit" :on-cancel="handleBack" :data="data?.item" :loading="loading" />
</template>

<style scoped>
nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
</style>
