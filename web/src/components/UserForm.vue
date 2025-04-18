<template>
  <el-row justify="center" class="form-container">
    <el-col class="form-wrapper">
      <h1>{{ isEditMode ? "Update User" : "Create User" }}</h1>
      <el-form ref="userForm" label-position="top">
        <el-form-item
          label="Email"
          :error="submitCount > 0 ? emailError : ''"
          status-icon
        >
          <el-input v-model="email" placeholder="Enter email" />
        </el-form-item>

        <el-form-item
          label="Age"
          :error="submitCount > 0 ? ageError : ''"
          status-icon
        >
          <el-input-number v-model="age" placeholder="Enter age" />
        </el-form-item>
        <el-form-item class="form-buttons">
          <el-button type="primary" @click="submit">
            {{ isEditMode ? "Update" : "Create" }}
          </el-button>
          <el-button @click="handleCancel">Cancel</el-button>
        </el-form-item>
        <div v-if="apiError" class="api-error">{{ apiError }}</div>
      </el-form>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import UserDataService from "@/services/UserDataService";
import { useField, useForm } from "vee-validate";
import * as yup from "yup";

const route = useRoute();
const router = useRouter();

const isEditMode = ref(!!route.params.id);

const form = ref({});

const apiError = ref("");

const { handleSubmit, submitCount } = useForm({
  validationSchema: yup.object({
    email: yup.string().required("Email is required").email("Invalid email"),
    age: yup
      .number()
      .required("Age is required")
      .min(18, "Age must be greater than 18"),
  }),
  validateOnMount: false,
});

const { value: email, errorMessage: emailError } = useField("email");

const { value: age, errorMessage: ageError } = useField("age");

const fetchUser = async () => {
  if (isEditMode.value) {
    try {
      const res = await UserDataService.get(route.params.id);
      email.value = res.data.email;
      age.value = res.data.age;
    } catch (error) {
      apiError.value = error.response?.data?.message || "Failed to fetch user.";
    }
  }
};

const submit = handleSubmit(async () => {
  try {
    const userPayload = {
      email: email.value,
      age: age.value,
    };

    if (isEditMode.value) {
      await UserDataService.update(route.params.id, userPayload);
      alert("User updated successfully!");
    } else {
      await UserDataService.create(userPayload);
      alert("User created successfully!");
    }
    router.push("/");
  } catch (error) {
    apiError.value = error.response?.data?.message || "Failed to submit form.";
  }
});

const handleCancel = () => {
  router.push("/");
};

watch(
  () => route.fullPath,
  () => {
    isEditMode.value = !!route.params.id;
    apiError.value = "";

    email.value = "";
    age.value = null;

    if (isEditMode.value) {
      fetchUser();
    }
  }
);

onMounted(fetchUser);
</script>

<style scoped>
.el-row {
  margin-top: 2rem;
}

.form-container {
  background: #fff;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  max-width: 600px;
  margin: 0 auto;
}

.form-buttons {
  padding-top: 1rem;
}

h1 {
  margin-bottom: 1.5rem;
}

.api-error {
  color: red;
  margin-top: 1rem;
  font-size: 1rem;
}
</style>
