<template>
  <div class="page-container">
    <h2 class="page-title">AI处理</h2>
    <el-row :gutter="24">
      <el-col :span="8" v-for="feature in aiFeatures" :key="feature.type">
        <el-card
          class="feature-card"
          shadow="hover"
          @click="handleFeature(feature.type)"
        >
          <div class="feature-content">
            <el-icon class="feature-icon">
              <component :is="feature.icon" />
            </el-icon>
            <h3>{{ feature.title }}</h3>
            <p class="feature-desc">{{ feature.description }}</p>
            <div class="feature-tags">
              <el-tag
                v-for="tag in feature.tags"
                :key="tag"
                size="small"
                class="feature-tag"
              >
                {{ tag }}
              </el-tag>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 功能对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="currentFeature?.title"
      width="70%"
      destroy-on-close
      class="feature-dialog"
    >
      <!-- 对话框内容保持不变 -->
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import {
  Monitor,
  VideoCamera,
  ChatLineRound,
  Connection,
  UploadFilled
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const aiFeatures = [
  {
    title: '超分辨率',
    description: '使用AI技术提升视频分辨率，支持2K/4K/8K输出',
    icon: Monitor,
    type: 'upscale',
    tags: ['超分辨率', '画质增强', '细节优化'],
    uploadTip: '建议上传720p以上视频文件'
  },
  {
    title: 'AI降噪',
    description: '智能去除视频噪点，提升画面清晰度',
    icon: VideoCamera,
    type: 'denoise',
    tags: ['降噪', '画质优化', '细节保护'],
    uploadTip: '支持任何分辨率的视频文件'
  },
  {
    title: 'AI补帧',
    description: '智能补充视频帧，实现流畅播放效果',
    icon: Connection,
    type: 'interpolate',
    tags: ['帧率提升', '动画优化', '平滑播放'],
    uploadTip: '建议上传30fps以上视频文件'
  },
  {
    title: '智能字幕',
    description: '自动识别语音并生成多语言字幕',
    icon: ChatLineRound,
    type: 'subtitle',
    tags: ['语音识别', '多语言', '字幕同步'],
    uploadTip: '请确保视频包含清晰的语音'
  }
]

const dialogVisible = ref(false)
const currentFeature = ref<any>(null)
const currentVideo = ref<{ url: string } | null>(null)

const aiOptions = reactive({
  resolution: '4k',
  model: 'standard',
  strength: 50,
  preserveDetail: true,
  targetFps: '60',
  algorithm: 'neural',
  language: ['zh'],
  style: 'standard'
})

const handleFeature = (type: string) => {
  currentFeature.value = aiFeatures.find(f => f.type === type)
  dialogVisible.value = true
}

const handleUploadSuccess = (response: any, file: any) => {
  currentVideo.value = {
    url: URL.createObjectURL(file)
  }
  ElMessage.success('上传成功')
}

const handleUploadError = () => {
  ElMessage.error('上传失败')
}

const handleProcess = () => {
  ElMessage.success(`开始${currentFeature.value.title}处理`)
  dialogVisible.value = false
}
</script>

<style scoped>
.ai-view {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.page-title {
  margin-bottom: 30px;
  color: #303133;
  font-size: 24px;
  text-align: center;
}

.feature-card {
  height: 280px;
  margin-bottom: 20px;
  transition: all 0.3s ease;
  cursor: pointer;
  display: flex;
  flex-direction: column;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 16px rgba(0,0,0,0.1);
  border-color: #409EFF;
}

.feature-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  padding: 20px;
  text-align: center;
}

.feature-icon {
  font-size: 48px;
  color: #409EFF;
  margin-bottom: 20px;
}

h3 {
  margin: 0 0 15px 0;
  color: #303133;
  font-size: 20px;
}

.feature-desc {
  color: #606266;
  margin: 0 0 20px 0;
  line-height: 1.6;
  flex-grow: 1;
}

.feature-tags {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 8px;
}

.feature-tag {
  background-color: #ecf5ff;
  border-color: #409EFF;
  color: #409EFF;
}

.video-preview {
  margin-top: 20px;
}

.preview-player {
  width: 100%;
  border-radius: 4px;
}

.feature-options {
  padding: 20px;
  background: #f8f9fa;
  border-radius: 4px;
}

:deep(.el-upload-dragger) {
  width: 100%;
  height: 200px;
}

/* 响应式布局 */
@media (max-width: 768px) {
  .el-col {
    width: 100% !important;
  }
}
</style>
