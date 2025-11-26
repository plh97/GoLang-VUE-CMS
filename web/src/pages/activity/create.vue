<script setup>
import { defineRouteMeta, useRouter } from '@fesjs/fes'
import { FButton, FDivider, FMessage } from '@fesjs/fes-design'
import { request } from '@/api'
import activityForm from '@/components/activityForm.vue'

defineRouteMeta({
  name: 'activity/create',
  title: '创建活动',
})

const { replace } = useRouter()
function handleBack() {
  replace('/activity')
}
function onSubmit(data) {
  request('/activity/create', data, {
    method: 'POST',
  }).then(() => {
    FMessage.success({
      content: '创建成功',
    })
    handleBack()
  })
}
</script>

<template>
  <nav>
    <h1>
      新增
    </h1>
    <FButton type="primary" @click="handleBack">
      返回
    </FButton>
  </nav>
  <FDivider />
  <activityForm :on-submit="onSubmit" :on-cancel="handleBack" />
</template>

<style scoped>
nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
</style>
