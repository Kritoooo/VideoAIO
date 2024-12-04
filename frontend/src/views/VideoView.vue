<template>
  <div class="video-view">
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card class="upload-card">
          <template #header>
            <div class="card-header">
              <h3>上传视频</h3>
            </div>
          </template>
          <VideoUpload @upload-success="handleUploadSuccess" />
        </el-card>
      </el-col>
      <el-col :span="12" v-if="currentVideo">
        <el-card class="preview-card">
          <template #header>
            <div class="card-header">
              <h3>视频预览</h3>
            </div>
          </template>
          <video
            v-if="currentVideo.url"
            :src="currentVideo.url"
            controls
            class="video-preview"
          ></video>
          <div class="video-info" v-if="currentVideo.name">
            <p>文件名：{{ currentVideo.name }}</p>
            <p>大小：{{ formatFileSize(currentVideo.size) }}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card class="process-card" v-if="currentVideo">
      <template #header>
        <div class="card-header">
          <h3>处理选项</h3>
        </div>
      </template>

      <el-form :model="processForm" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="输出格式">
              <el-select v-model="processForm.format" placeholder="选择输出格式">
                <el-option label="MP4" value="mp4" />
                <el-option label="WebM" value="webm" />
                <el-option label="GIF" value="gif" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="分辨率">
              <el-select v-model="processForm.resolution" placeholder="选择分辨率">
                <el-option label="原始分辨率" value="original" />
                <el-option label="1080P" value="1080p" />
                <el-option label="720P" value="720p" />
                <el-option label="480P" value="480p" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="视频压缩">
              <el-slider
                v-model="processForm.compression"
                :marks="{
                  0: '无压缩',
                  50: '中等',
                  100: '最大压缩'
                }"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="帧率">
              <el-select v-model="processForm.fps" placeholder="选择帧率">
                <el-option label="原始帧率" value="original" />
                <el-option label="60 FPS" value="60" />
                <el-option label="30 FPS" value="30" />
                <el-option label="24 FPS" value="24" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="高级选项">
              <el-checkbox-group v-model="processForm.options">
                <el-checkbox label="watermark">添加水印</el-checkbox>
                <el-checkbox label="trim">裁剪时长</el-checkbox>
                <el-checkbox label="subtitle">添加字幕</el-checkbox>
                <el-checkbox label="audio">提取音频</el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item>
          <el-button type="primary" @click="handleProcess" :loading="processing">
            开始处理
          </el-button>
          <el-button @click="resetForm">重置选项</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-dialog
      v-model="processDialog.visible"
      :title="processDialog.title"
      width="30%"
      :close-on-click-modal="false"
    >
      <el-progress
        :percentage="processDialog.progress"
        :status="processDialog.status"
      />
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="processDialog.visible = false" :disabled="processing">
            关闭
          </el-button>
          <el-button type="primary" @click="downloadResult" :disabled="!processDialog.completed">
            下载结果
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import VideoUpload from '@/components/VideoUpload.vue'
import { ElMessage } from 'element-plus'
import { processVideo } from '@/api/video'

const currentVideo = ref<{
  url?: string
  name?: string
  size?: number
  id?: string
} | null>(null)

const processForm = reactive({
  format: 'mp4',
  resolution: 'original',
  compression: 0,
  fps: 'original',
  options: []
})

const processing = ref(false)
const processDialog = reactive({
  visible: false,
  title: '处理进度',
  progress: 0,
  status: 'active',
  completed: false
})

const handleUploadSuccess = (response: any) => {
  currentVideo.value = {
    url: URL.createObjectURL(response.file),
    name: response.file.name,
    size: response.file.size,
    id: response.id
  }
  ElMessage.success('上传成功')
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const handleProcess = async () => {
  if (!currentVideo.value?.id) {
    ElMessage.warning('请先上传视频')
    return
  }

  processing.value = true
  processDialog.visible = true
  processDialog.progress = 0
  processDialog.status = 'active'
  processDialog.completed = false

  try {
    // 模拟进度
    const timer = setInterval(() => {
      if (processDialog.progress < 90) {
        processDialog.progress += 10
      }
    }, 1000)

    await processVideo(currentVideo.value.id)

    clearInterval(timer)
    processDialog.progress = 100
    processDialog.status = 'success'
    processDialog.completed = true
    ElMessage.success('处理完成')
  } catch (error) {
    processDialog.status = 'exception'
    ElMessage.error('处理失败')
  } finally {
    processing.value = false
  }
}

const resetForm = () => {
  processForm.format = 'mp4'
  processForm.resolution = 'original'
  processForm.compression = 0
  processForm.fps = 'original'
  processForm.options = []
}

const downloadResult = () => {
  // TODO: 实现下载逻辑
  ElMessage.success('开始下载')
}
</script>

<style scoped>
.video-view {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.upload-card, .preview-card, .process-card {
  margin-bottom: 20px;
}

.video-preview {
  width: 100%;
  max-height: 300px;
  object-fit: contain;
}

.video-info {
  margin-top: 10px;
  color: #666;
}

.process-card {
  margin-top: 20px;
}

:deep(.el-slider__marks-text) {
  font-size: 12px;
}
</style>
