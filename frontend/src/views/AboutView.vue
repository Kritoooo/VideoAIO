<template>
  <div class="about-view">
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card class="about-card">
          <template #header>
            <div class="card-header">
              <h2>关于 VideoAIO</h2>
            </div>
          </template>
          <div class="about-content">
            <div class="section">
              <h3>
                <el-icon><InfoFilled /></el-icon>
                项目介绍
              </h3>
              <p>VideoAIO 是一个功能强大的在线视频处理工具，提供视频转换、剪辑、AI 增强等功能。无需安装任何软件，通过浏览器即可完成专业级的视频处理工作。</p>
            </div>

            <div class="section">
              <h3>
                <el-icon><List /></el-icon>
                主要功能
              </h3>
              <ul>
                <li>视频格式转换：支持多种格式之间的转换</li>
                <li>视频剪辑：快速剪辑、合并视频片段</li>
                <li>AI 增强：智能提升视频质量</li>
                <li>字幕处理：自动生成、编辑字幕</li>
                <li>视频压缩：智能压缩，保持画质</li>
              </ul>
            </div>

            <div class="section">
              <h3>
                <el-icon><Monitor /></el-icon>
                技术栈
              </h3>
              <div class="tech-stack">
                <el-tag>Vue 3</el-tag>
                <el-tag>TypeScript</el-tag>
                <el-tag>Element Plus</el-tag>
                <el-tag>Go</el-tag>
                <el-tag>FFmpeg</el-tag>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="8">
        <el-card class="contact-card">
          <template #header>
            <div class="card-header">
              <h3>联系我们</h3>
            </div>
          </template>
          <div class="contact-content">
            <div class="contact-item">
              <el-icon><Message /></el-icon>
              <span>反馈建议</span>
            </div>
            <div class="contact-item">
              <el-icon><Link /></el-icon>
              <a href="https://github.com/yourusername/videoaio" target="_blank">GitHub</a>
            </div>
            <div class="contact-item">
              <el-icon><Document /></el-icon>
              <a href="#" target="_blank">使用文档</a>
            </div>
          </div>
        </el-card>

        <el-card class="version-card">
          <template #header>
            <div class="card-header">
              <h3>版本信息</h3>
            </div>
          </template>
          <div class="version-content">
            <p>当前版本：v1.0.0</p>
            <p>更新日期：2024-02-20</p>
            <el-button type="primary" plain size="small" class="update-btn">
              检查更新
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { InfoFilled, List, Monitor, Message, Link, Document } from '@element-plus/icons-vue'
</script>

<style scoped>
.about-view {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.about-card {
  margin-bottom: 20px;
}

.section {
  margin-bottom: 30px;
}

.section h3 {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #409EFF;
  margin-bottom: 15px;
}

.section p {
  color: #606266;
  line-height: 1.6;
  margin: 10px 0;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  color: #606266;
  margin: 8px 0;
  padding-left: 20px;
  position: relative;
}

li::before {
  content: "•";
  color: #409EFF;
  position: absolute;
  left: 0;
}

.tech-stack {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.contact-card, .version-card {
  margin-bottom: 20px;
}

.contact-content {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #606266;
}

.contact-item a {
  color: #409EFF;
  text-decoration: none;
}

.contact-item a:hover {
  text-decoration: underline;
}

.version-content {
  color: #606266;
}

.version-content p {
  margin: 5px 0;
}

.update-btn {
  margin-top: 15px;
}

:deep(.el-tag) {
  margin: 5px;
}
</style>
