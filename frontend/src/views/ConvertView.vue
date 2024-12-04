<template>
  <div class="convert-view">
    <h2 class="page-title">端到端转换</h2>
    <el-row :gutter="24">
      <el-col :span="8" v-for="feature in features" :key="feature.type">
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
      <div v-if="dialogVisible" class="feature-dialog">
        <el-row :gutter="20">
          <el-col :span="16">
            <el-upload
              class="upload-area"
              drag
              action="/api/v1/video/upload"
              :on-success="handleUploadSuccess"
              :on-error="handleUploadError"
              accept="video/*"
            >
              <el-icon class="el-icon--upload"><upload-filled /></el-icon>
              <div class="el-upload__text">
                拖拽文件到此处或 <em>点击上传</em>
              </div>
              <template #tip>
                <div class="el-upload__tip">
                  {{ currentFeature?.uploadTip }}
                </div>
              </template>
            </el-upload>

            <div v-if="currentVideo" class="video-preview">
              <video
                :src="currentVideo.url"
                controls
                class="preview-player"
              ></video>
            </div>
          </el-col>

          <el-col :span="8">
            <div class="feature-options">
              <!-- 格式转换选项 -->
              <template v-if="currentFeature?.type === 'format'">
                <el-form label-position="top">
                  <el-form-item label="输出格式">
                    <el-select v-model="convertOptions.format">
                      <el-option label="MP4" value="mp4" />
                      <el-option label="WebM" value="webm" />
                      <el-option label="MKV" value="mkv" />
                      <el-option label="AVI" value="avi" />
                      <el-option label="MOV" value="mov" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="编码器">
                    <el-select v-model="convertOptions.codec">
                      <el-option label="H.264" value="h264" />
                      <el-option label="H.265" value="h265" />
                      <el-option label="VP9" value="vp9" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="视频质量">
                    <el-slider
                      v-model="convertOptions.quality"
                      :marks="{
                        0: '低',
                        50: '中',
                        100: '高'
                      }"
                    />
                  </el-form-item>
                </el-form>
              </template>

              <!-- 音频提取选项 -->
              <template v-if="currentFeature?.type === 'extract'">
                <el-form label-position="top">
                  <el-form-item label="音频格式">
                    <el-select v-model="extractOptions.format">
                      <el-option label="MP3" value="mp3" />
                      <el-option label="WAV" value="wav" />
                      <el-option label="AAC" value="aac" />
                      <el-option label="FLAC" value="flac" />
                      <el-option label="OGG" value="ogg" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="音频质量">
                    <el-select v-model="extractOptions.quality">
                      <el-option label="标准质量 (128kbps)" value="128" />
                      <el-option label="高质量 (256kbps)" value="256" />
                      <el-option label="超高质量 (320kbps)" value="320" />
                      <el-option label="无损" value="lossless" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="声道">
                    <el-radio-group v-model="extractOptions.channels">
                      <el-radio :value="'mono'">单声道</el-radio>
                      <el-radio :value="'stereo'">立体声</el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-form>
              </template>

              <!-- 压缩选项 -->
              <template v-if="currentFeature?.type === 'compress'">
                <el-form label-position="top">
                  <el-form-item label="压缩模式">
                    <el-radio-group v-model="compressOptions.mode">
                      <el-radio :value="'size'">目标大小</el-radio>
                      <el-radio :value="'quality'">目标质量</el-radio>
                    </el-radio-group>
                  </el-form-item>

                  <template v-if="compressOptions.mode === 'size'">
                    <el-form-item label="目标大小 (MB)">
                      <el-input-number
                        v-model="compressOptions.targetSize"
                        :min="1"
                        :max="1000"
                      />
                    </el-form-item>
                  </template>

                  <template v-if="compressOptions.mode === 'quality'">
                    <el-form-item label="压缩质量">
                      <el-slider
                        v-model="compressOptions.quality"
                        :marks="{
                          20: '高压缩',
                          50: '平衡',
                          80: '高质量'
                        }"
                      />
                    </el-form-item>
                  </template>

                  <el-form-item label="高级选项">
                    <el-checkbox v-model="compressOptions.keepMetadata">
                      保留元数据
                    </el-checkbox>
                    <el-checkbox v-model="compressOptions.fastMode">
                      快速模式
                    </el-checkbox>
                  </el-form-item>
                </el-form>
              </template>
            </div>

            <div class="process-info" v-if="currentVideo">
              <el-descriptions :column="1" border>
                <el-descriptions-item label="文件名">
                  {{ currentVideo.name }}
                </el-descriptions-item>
                <el-descriptions-item label="原始大小">
                  {{ formatSize(currentVideo.size) }}
                </el-descriptions-item>
                <el-descriptions-item label="预计时长">
                  {{ formatDuration(currentVideo.duration) }}
                </el-descriptions-item>
              </el-descriptions>
            </div>
          </el-col>
        </el-row>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleProcess" :loading="processing">
            开始处理
          </el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 处理进度对话框 -->
    <el-dialog
      v-model="progressVisible"
      title="处理进度"
      width="30%"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <el-progress
        :percentage="progress"
        :status="progressStatus"
      />
      <div class="progress-info" v-if="processing">
        <p>{{ progressInfo }}</p>
        <p class="time-remaining" v-if="timeRemaining">
          预计剩余时间: {{ timeRemaining }}
        </p>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelProcess" :disabled="!processing">
            取消
          </el-button>
          <el-button
            type="primary"
            @click="downloadResult"
            :disabled="!isCompleted"
          >
            下载
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import {
  Connection,
  Headset,
  VideoCamera,
  UploadFilled
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const features = [
  {
    title: '格式转换',
    description: '支持多种视频格式之间的转换，保持原有画质',
    icon: Connection,
    type: 'format',
    tags: ['MP4', 'AVI', 'MKV', 'WebM'],
    uploadTip: '支持上传 MP4, AVI, MOV, MKV 等格式视频文件',
    formats: [
      { label: 'MP4', value: 'mp4' },
      { label: 'WebM', value: 'webm' },
      { label: 'MKV', value: 'mkv' },
      { label: 'AVI', value: 'avi' }
    ]
  },
  {
    title: '音频提取',
    description: '从视频中提取音频，支持多种音频格式',
    icon: Headset,
    type: 'extract',
    tags: ['MP3', 'WAV', 'AAC'],
    uploadTip: '支持从任何包含音频的视频文件中提取'
  },
  {
    title: '视频压缩',
    description: '智能压缩视频文件大小，保持画质',
    icon: VideoCamera,
    type: 'compress',
    tags: ['体积优化', '画质平衡', '快速压缩'],
    uploadTip: '建议上传 5GB 以内的视频文件'
  }
]

const dialogVisible = ref(false)
const currentFeature = ref<any>(null)
const convertOptions = reactive({
  format: 'mp4',
  quality: 50
})
const extractOptions = reactive({
  format: 'mp3',
  quality: 'normal'
})
const compressOptions = reactive({
  mode: 'size',
  targetSize: 100,
  quality: 50,
  keepMetadata: true,
  fastMode: true
})
const processing = ref(false)
const progressVisible = ref(false)
const progress = ref(0)
const progressStatus = ref('')
const progressInfo = ref('')
const timeRemaining = ref('')
const isCompleted = ref(false)
const currentVideo = ref<any>(null)

const handleFeature = (type: string) => {
  currentFeature.value = features.find(f => f.type === type)
  dialogVisible.value = true
}

const handleUploadSuccess = (response: any) => {
  ElMessage.success('上传成功')
}

const handleUploadError = () => {
  ElMessage.error('上传失败')
}

const handleProcess = () => {
  // 根据不同功能类型处理
  switch (currentFeature.value?.type) {
    case 'format':
      ElMessage.success(`开始转换为 ${convertOptions.format} 格式`)
      break
    case 'extract':
      ElMessage.success(`开始提取 ${extractOptions.format} 音频`)
      break
    case 'compress':
      ElMessage.success('开始压缩视频')
      break
  }
  dialogVisible.value = false
}

const formatSize = (bytes: number) => {
  if (!bytes) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`
}

const formatDuration = (seconds: number) => {
  if (!seconds) return '00:00'
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = Math.floor(seconds % 60)
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

const cancelProcess = () => {
  // TODO: 实现取消处理的逻辑
  processing.value = false
  progressVisible.value = false
  ElMessage.info('已取消处理')
}

const downloadResult = () => {
  // TODO: 实现下载结果的逻辑
  ElMessage.success('开始下载处理结果')
}
</script>

<style scoped>
.convert-view {
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

.upload-area {
  width: 100%;
}

:deep(.el-upload-dragger) {
  width: 100%;
  height: 200px;
}

.convert-options,
.extract-options {
  margin-top: 20px;
}

.process-info {
  margin-top: 20px;
  padding: 15px;
  background: #f8f9fa;
  border-radius: 4px;
}

.progress-info {
  margin-top: 15px;
  color: #606266;
}

.time-remaining {
  margin-top: 5px;
  color: #909399;
  font-size: 14px;
}

.feature-options {
  margin-bottom: 20px;
}
</style>
