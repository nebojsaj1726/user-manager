<template>
  <el-row
    justify="center"
    v-loading="isLoading"
    element-loading-text="Loading Users..."
  >
    <el-col :span="20">
      <div class="search-wrapper">
        <el-input
          v-model="search"
          placeholder="Search by email"
          clearable
          size="large"
          class="search-input"
        />
      </div>

      <el-table
        v-if="!isLoading && !error"
        :data="filteredUsers"
        highlight-current-row
        size="large"
        class="custom-table"
      >
        <el-table-column prop="email" label="Email" />
        <el-table-column prop="age" label="Age" />
        <el-table-column label="Actions" width="180">
          <template #default="scope">
            <router-link
              :to="`/edit/${scope.row.id}`"
              class="el-link el-link--primary"
              style="margin-right: 10px"
            >
              Edit
            </router-link>
            <router-link
              :to="`/delete/${scope.row.id}`"
              class="el-link el-link--danger"
            >
              Delete
            </router-link>
          </template>
        </el-table-column>
      </el-table>

      <div
        v-if="!isLoading && !error && filteredUsers.length === 0"
        class="empty-state"
      >
        No users found.
      </div>

      <el-alert
        v-if="error"
        title="Error"
        type="error"
        :description="error"
        show-icon
      />

      <el-pagination
        v-if="!isLoading && total > pageSize"
        background
        layout="prev, pager, next"
        :total="total"
        :page-size="pageSize"
        :current-page="currentPage"
        @current-change="handlePageChange"
        class="pagination"
      />
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import UserDataService from "@/services/UserDataService";

const users = ref([]);
const search = ref("");
const isLoading = ref(false);
const error = ref(null);

const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);

const fetchUsers = async (page = 1, limit = 10) => {
  isLoading.value = true;
  try {
    const params = { page, limit };
    const res = await UserDataService.getAll(params);
    users.value = res.data.users;
    total.value = res.data.total;
    error.value = null;
  } catch (e) {
    error.value = e.response?.data?.message || "Failed to load users";
  } finally {
    isLoading.value = false;
  }
};

const filteredUsers = computed(() => {
  return users.value.filter((user) =>
    user.email.toLowerCase().includes(search.value.toLowerCase())
  );
});

const handlePageChange = (page) => {
  currentPage.value = page;
  fetchUsers(page, pageSize.value);
};

onMounted(() => fetchUsers(currentPage.value, pageSize.value));
</script>

<style scoped>
.search-wrapper {
  display: flex;
  justify-content: center;
  margin: 2rem 0;
}

.search-input {
  width: 500px;
  max-width: 100%;
}

.custom-table {
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  background: white;
}

.pagination {
  margin-top: 1.5rem;
  justify-content: center;
}
</style>
