<script setup>
import { defineRouteMeta } from '@fesjs/fes'
import { FButton, FMessage } from '@fesjs/fes-design'
import html2canvas from 'html2canvas' // 需要安装
import jsQR from 'jsqr'
import { onMounted, onUnmounted, ref } from 'vue'
import { request } from '@/api'
import { SIGN_UP_STATUS } from '@/enums'

defineRouteMeta({
  name: 'scan',
  title: '扫码签到',
  layout: {
    navigation: null,
  },
})

const video = ref(null)
const canvas = ref(null)
let stream = null

// 初始化摄像头
async function initCamera() {
  try {
    stream = await navigator.mediaDevices.getUserMedia({
      video: { facingMode: 'environment' }, // 优先使用后置摄像头
    })
    video.value.srcObject = stream
    video.value.play()
  }
  catch (err) {
    console.error('无法访问摄像头: ', err)
    // TODO: 提示用户授权或设备不支持
    FMessage('需要摄像头权限才能使用扫码功能！')
  }
}

// 处理识别到的数据
function handleQRCodeData(data) {
  try {
    const json = JSON.parse(data)
    if (json.user_id && json.activity_id) {
      // 使用你的组件库提示
      // 可以在这里进行下一步操作，例如发送请求
      request('/sign_up/status_update', {
        activity_id: json.activity_id,
        user_id: json.user_id,
        status: SIGN_UP_STATUS.已签到,
      }, {
        method: 'post',
      }).then(() => {
        FMessage.success({
          content: '修改成功',
        })
      })
    }
  }
  catch (e) {
    console.error('数据解析失败:', e)
    FMessage.warning(`识别到内容但格式错误: ${data.substring(0, 30)}...`)
  }
}

onMounted(() => {
  initCamera()
})

onUnmounted(() => {
  if (stream) {
    stream.getTracks().forEach(track => track.stop())
  }
})

// 执行截图和识别
async function captureAndIdentify() {
  const canvasElement = await html2canvas(video.value, {
    useCORS: true, // 如果视频流有跨域问题可能需要
    allowTaint: true,
  })
  const context = canvasElement.getContext('2d')
  const imageData = context.getImageData(0, 0, canvasElement.width, canvasElement.height)
  const code = jsQR(imageData.data, imageData.width, imageData.height, {
    inversionAttempts: 'dontInvert',
  })
  if (code) {
    handleQRCodeData(code.data)
  }
  else {
    FMessage.error({ content: '截图未识别到二维码' })
    // 截图失败，恢复实时扫码循环
  }
}
</script>

<template>
  <div>
    <div class="scanner-container">
      <video id="video" ref="video" autoplay playsinline />
      <canvas id="canvas" ref="canvas" style="display: none;" />
      <div class="scan-area" />
    </div>
    <div class="capture-button">
      <FButton type="primary" size="large" @click="captureAndIdentify">
        按住截图识别
      </FButton>
      <!-- <FButton size="large" @click="handlePause">
        暂停
      </FButton> -->
    </div>
  </div>
</template>

<style scoped>
.scanner-container {
  position: relative;
  width: 100vw;
  overflow: hidden;
  background-color: #000;
}

#video {
  width: 100%;
  height: 35vh;
  object-fit: cover; /* 保证视频填满容器 */
  /* transform: scaleX(-1); */
}

.scan-area {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 70%;
  height: 25vh;
  border: 2px solid #4CAF50;
  box-shadow: 0 0 0 1000px rgba(0, 0, 0, 0.5); /* 遮罩效果 */
  pointer-events: none; /* 不会影响下面的视频点击 */
  z-index: 10;
}

.capture-button {
  position: fixed;
  bottom: 50px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  flex-direction: row;
  gap: 10px;
}
</style>
