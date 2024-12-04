<template>
  <div class="page-container">
    <el-row :gutter="24">
      <el-col :span="8" v-for="tool in tools" :key="tool.path">
        <router-link :to="tool.path" class="tool-link">
          <el-card class="feature-card" shadow="hover">
            <div class="feature-content">
              <el-icon class="feature-icon">
                <component :is="tool.icon" />
              </el-icon>
              <h3 class="feature-title">{{ tool.name }}</h3>
              <p class="feature-desc">{{ tool.description }}</p>
            </div>
          </el-card>
        </router-link>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import {
  Connection,
  Crop,
  Monitor
} from '@element-plus/icons-vue'

const tools = [
  {
    name: '端到端转换',
    description: '一键完成视频格式转换，支持多种格式，快速导出',
    icon: Connection,
    path: '/convert'
  },
  {
    name: '片段处理',
    description: '视频剪辑、合并、提取音频等片段处理功能',
    icon: Crop,
    path: '/edit'
  },
  {
    name: 'AI处理',
    description: '智能增强、场景识别、字幕生成等AI功能',
    icon: Monitor,
    path: '/ai'
  }
]
</script>

<style scoped>
.tool-link {
  text-decoration: none;
  color: inherit;
  display: block;
}
</style>
