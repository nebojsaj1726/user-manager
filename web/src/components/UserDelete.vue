<template>
  <el-row justify="center" class="delete-container">
    <el-col class="delete-wrapper">
      <h1>Delete User</h1>
      <p>Are you sure you want to delete this user?</p>
      <div class="delete-buttons">
        <el-button type="danger" @click="deleteUser">Delete</el-button>
        <el-button @click="handleCancel">Cancel</el-button>
      </div>
      <div v-if="apiError" class="api-error">{{ apiError }}</div>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import UserDataService from "@/services/UserDataService";

const route = useRoute();
const router = useRouter();

const apiError = ref("");

const deleteUser = async () => {
  try {
    await UserDataService.delete(route.params.id);
    alert("User deleted successfully!");
    router.push("/");
  } catch (error) {
    apiError.value = error.response?.data?.message || "Failed to delete user.";
  }
};

const handleCancel = () => {
  router.push("/");
};
</script>

<style scoped>
.delete-container {
  margin-top: 2rem;
  background: #fff;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  max-width: 600px;
  margin: 0 auto;
  text-align: center;
}

.delete-wrapper h1 {
  margin-bottom: 1rem;
}

.delete-buttons {
  margin-top: 1.5rem;
}

.api-error {
  color: red;
  margin-top: 1rem;
  font-size: 1rem;
}
</style>
