<script lang="ts" setup>
import { useRouter } from '@fesjs/fes'
import { FButton } from '@fesjs/fes-design'
import { request } from '@/api'
import { setToken } from '@/common/utils'

// import { useModel } from '@fesjs/fes'

// const initialState = useModel('@@initialState')
const router = useRouter()
async function handleLogout() {
  await request('/admin/logout', {}, { method: 'POST' })
  setToken('')
  router.replace('/login')
}
</script>

<template>
  <div class="right">
    <!-- {{ initialState.userName }} -->
    <FButton type="link" class="link" @click="handleLogout">
      退出登录
    </FButton>
  </div>
</template>

<style scope>
.link {
  color: rgb(255, 255, 255);
}
.right {
  text-align: right;
  padding: 0 20px;
}
</style>
