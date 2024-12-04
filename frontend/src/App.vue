<template>
  <el-config-provider>
    <el-container>
      <el-header>
        <div class="header-content">
          <div class="logo">VideoAIO</div>
          <el-menu
            mode="horizontal"
            router
            :default-active="$route.path"
          >
            <el-menu-item index="/home">
              <el-icon><HomeFilled /></el-icon>
              首页
            </el-menu-item>
            <el-menu-item index="/convert">
              <el-icon><Connection /></el-icon>
              端到端转换
            </el-menu-item>
            <el-menu-item index="/edit">
              <el-icon><Crop /></el-icon>
              片段处理
            </el-menu-item>
            <el-menu-item index="/ai">
              <el-icon><Monitor /></el-icon>
              AI处理
            </el-menu-item>
            <el-menu-item index="/about">
              <el-icon><InfoFilled /></el-icon>
              关于
            </el-menu-item>
          </el-menu>
        </div>
      </el-header>

      <el-main>
        <router-view v-slot="{ Component }">
          <component :is="Component" />
        </router-view>
      </el-main>

      <el-footer>
        <div class="footer-content">
          VideoAIO © 2024
        </div>
      </el-footer>
    </el-container>
  </el-config-provider>
</template>

<script setup lang="ts">
import { HomeFilled, Connection, Crop, Monitor, InfoFilled } from '@element-plus/icons-vue'
</script>

<style scoped>
.el-container {
  min-height: 100vh;
}

.el-header {
  background-color: #fff;
  border-bottom: 1px solid #dcdfe6;
  padding: 0;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  height: 60px;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: #409EFF;
  margin-right: 40px;
}

.el-menu {
  flex: 1;
  border-bottom: none;
}

.el-menu-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.el-menu-item .el-icon {
  margin-right: 4px;
}

.el-main {
  background-color: #f5f7fa;
  padding: 20px;
}

.el-footer {
  background-color: #fff;
  text-align: center;
  color: #909399;
  padding: 20px 0;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
}
</style>
