// router.js
import { createWebHistory, createRouter } from "vue-router";

const routes = [
  {
    path: "/",
    alias: "/users",
    name: "user-list",
    component: () => import("./src/components/UserList.vue"),
  },
  {
    path: "/add",
    name: "user-create",
    component: () => import("./src/components/UserForm.vue"),
  },
  {
    path: "/edit/:id",
    name: "user-edit",
    component: () => import("./src/components/UserForm.vue"),
  },
  {
    path: "/delete/:id",
    name: "user-delete",
    component: () => import("./src/components/UserDelete.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
