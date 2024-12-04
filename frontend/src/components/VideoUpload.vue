<template>
  <div class="video-upload">
    <el-upload
      class="upload-demo"
      drag
      action="/api/v1/video/upload"
      :on-success="handleSuccess"
      :on-error="handleError"
      :before-upload="beforeUpload"
      :on-progress="handleProgress"
      accept="video/*"
      :show-file-list="false"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        拖拽视频文件到此处或 <em>点击上传</em>
      </div>
      <div class="el-upload__tip">
        支持的格式：MP4, AVI, MOV, WebM 等视频文件
      </div>
    </el-upload>

    <el-progress
      v-if="uploadProgress > 0 && uploadProgress < 100"
      :percentage="uploadProgress"
      :status="uploadStatus"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { UploadFilled } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const emit = defineEmits(['upload-success'])
const uploadProgress = ref(0)
const uploadStatus = ref('active')

const handleSuccess = (response: any, file: File) => {
  uploadProgress.value = 100
  uploadStatus.value = 'success'
  emit('upload-success', { file, id: response.id })
}

const handleError = () => {
  uploadStatus.value = 'exception'
  ElMessage.error('上传失败')
}

const handleProgress = (event: any) => {
  uploadProgress.value = Math.round(event.percent)
}

const beforeUpload = (file: File) => {
  const isVideo = file.type.startsWith('video/')
  if (!isVideo) {
    ElMessage.error('只能上传视频文件!')
    return false
  }

  const isLt2G = file.size / 1024 / 1024 / 1024 < 2
  if (!isLt2G) {
    ElMessage.error('视频大小不能超过 2GB!')
    return false
  }

  uploadProgress.value = 0
  uploadStatus.value = 'active'
  return true
}
</script>

<style scoped>
.video-upload {
  width: 100%;
}

.el-upload {
  width: 100%;
}

.el-upload-dragger {
  width: 100%;
}

.el-upload__tip {
  margin-top: 10px;
  color: #666;
}

:deep(.el-upload-dragger) {
  width: 100%;
  height: 200px;
}

:deep(.el-progress) {
  margin-top: 20px;
}
</style>
