<template>
  <div class="page-container">
    <h2 class="page-title">片段处理</h2>
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
            <h3 class="feature-title">{{ feature.title }}</h3>
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
              <div class="timeline-container">
                <el-slider
                  v-model="timeRange"
                  range
                  :max="duration"
                  :format-tooltip="formatTime"
                />
              </div>
            </div>
          </el-col>

          <el-col :span="8">
            <div class="feature-options">
              <!-- 剪辑选项 -->
              <template v-if="currentFeature?.type === 'cut'">
                <el-form label-position="top">
                  <el-form-item label="剪辑范围">
                    <div class="time-range">
                      <span>开始：{{ formatTime(timeRange[0]) }}</span>
                      <span>结束：{{ formatTime(timeRange[1]) }}</span>
                    </div>
                  </el-form-item>
                  <el-form-item label="输出格式">
                    <el-select v-model="editOptions.format">
                      <el-option label="MP4" value="mp4" />
                      <el-option label="WebM" value="webm" />
                    </el-select>
                  </el-form-item>
                </el-form>
              </template>

              <!-- 合并选项 -->
              <template v-if="currentFeature?.type === 'merge'">
                <el-upload
                  class="merge-upload"
                  action="/api/v1/video/upload"
                  :on-success="handleMergeFileSuccess"
                  multiple
                  list-type="text"
                >
                  <el-button type="primary">添加更多视频</el-button>
                </el-upload>
                <div class="merge-list">
                  <el-table :data="mergeFiles" style="width: 100%">
                    <el-table-column prop="name" label="文件名" />
                    <el-table-column width="100">
                      <template #default="scope">
                        <el-button
                          type="danger"
                          size="small"
                          @click="removeMergeFile(scope.$index)"
                        >
                          删除
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </template>

              <!-- 转场选项 -->
              <template v-if="currentFeature?.type === 'transition'">
                <el-form label-position="top">
                  <el-form-item label="转场效果">
                    <el-select v-model="editOptions.transition">
                      <el-option label="淡入淡出" value="fade" />
                      <el-option label="滑动" value="slide" />
                      <el-option label="溶解" value="dissolve" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="转场时长">
                    <el-slider
                      v-model="editOptions.duration"
                      :min="0.5"
                      :max="2"
                      :step="0.1"
                      :marks="{
                        0.5: '0.5s',
                        1: '1s',
                        2: '2s'
                      }"
                    />
                  </el-form-item>
                </el-form>
              </template>
            </div>
          </el-col>
        </el-row>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleProcess">
            开始处理
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import {
  Crop,
  Plus,
  Switch,
  UploadFilled
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const features = [
  {
    title: '视频剪辑',
    description: '精确剪辑视频片段，支持多种输出格式',
    icon: Crop,
    type: 'cut',
    tags: ['时间轴', '预览', '多格式'],
    uploadTip: '支持上传 MP4, AVI, MOV 等格式视频文件'
  },
  {
    title: '视频合并',
    description: '合并多个视频片段，自定义排序',
    icon: Plus,
    type: 'merge',
    tags: ['多文件', '排序', '预览'],
    uploadTip: '请上传相同格式的视频文件'
  },
  {
    title: '转场效果',
    description: '添加专业的转场效果，提升视频质量',
    icon: Switch,
    type: 'transition',
    tags: ['淡入淡出', '滑动', '溶解'],
    uploadTip: '建议上传高质量视频文件'
  }
]

const dialogVisible = ref(false)
const currentFeature = ref<any>(null)
const currentVideo = ref<{ url: string } | null>(null)
const duration = ref(100)
const timeRange = ref([0, duration.value])
const mergeFiles = ref<any[]>([])

const editOptions = reactive({
  format: 'mp4',
  transition: 'fade',
  duration: 1
})

const handleFeature = (type: string) => {
  currentFeature.value = features.find(f => f.type === type)
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

const handleMergeFileSuccess = (response: any, file: any) => {
  mergeFiles.value.push({
    name: file.name,
    url: URL.createObjectURL(file)
  })
}

const removeMergeFile = (index: number) => {
  mergeFiles.value.splice(index, 1)
}

const formatTime = (seconds: number) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = Math.floor(seconds % 60)
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

const handleProcess = () => {
  // 根据不同功能类型处理
  switch (currentFeature.value?.type) {
    case 'cut':
      ElMessage.success(`开始剪辑视频片段`)
      break
    case 'merge':
      ElMessage.success(`开始合并视频文件`)
      break
    case 'transition':
      ElMessage.success(`开始添加转场效果`)
      break
  }
  dialogVisible.value = false
}
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.merge-list {
  margin-top: 20px;
}

.time-range {
  display: flex;
  justify-content: space-between;
  color: #606266;
}

.timeline-container {
  margin-top: 20px;
  padding: 0 20px;
}
</style>
