<template>
  <el-container style="min-height: 100vh">
    <el-header height="60px" class="header">
      <el-menu
        :default-active="activeMenu"
        mode="horizontal"
        router
        background-color="#ffffff"
        text-color="#303133"
        active-text-color="#409EFF"
        class="nav-menu"
      >
        <el-menu-item index="/">User List</el-menu-item>
        <el-menu-item index="/add">Add User</el-menu-item>
      </el-menu>
    </el-header>

    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup>
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const activeMenu = computed(() => route.path);
</script>

<style>
body {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    "Helvetica Neue", Arial, sans-serif;
  font-size: 16px;
  background-color: #f8f9fa;
  color: #2c3e50;
}

.header {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08);
  display: flex;
  align-items: center;
  justify-content: center;
}

.nav-menu {
  width: 100%;
  max-width: 900px;
  font-size: 16px;
}
</style>
