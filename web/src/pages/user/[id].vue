<script setup>
import { defineRouteMeta, useRoute, useRouter } from '@fesjs/fes'
import { FButton, FDivider, FMessage } from '@fesjs/fes-design'
import { useRequest } from 'vue-hooks-plus'
import { request } from '@/api'
import { formatArray } from '@/common/utils'
import userFrom from '@/components/userForm.vue'

defineRouteMeta({
  name: 'user/detail',
  title: '编辑',
})

const { params } = useRoute()
const { loading, data } = useRequest(() => request('/user_profile/detail', {
  user_id: +params.id,
}, {
  method: 'post',
  transformData: (data) => {
    // 处理响应内容异常
    if (data?.code !== 0) {
      // Reject the promise with an error object containing code and message
      FMessage.error({
        content: data?.msg,
      })
    }
    return {
      ...data.data,
      images: formatArray(data.data.images),
    }
  },
}))
const { replace } = useRouter()
function handleBack() {
  replace('/user')
}
function onSubmit(data) {
  request('/user_profile/update', data, {
    method: 'POST',
  }).then(() => {
    FMessage.success({
      content: '修改成功',
    })
    replace('/user')
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
  <userFrom :on-submit="onSubmit" :on-cancel="handleBack" :data="data" :loading="loading" />
</template>

<style lang="less" scoped>
nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
</style>
