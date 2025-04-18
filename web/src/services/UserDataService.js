import http from "../http-common";

const UserDataService = {
  getAll(params) {
    return http.get("/users", { params });
  },

  get(id) {
    return http.get(`/users/${id}`);
  },

  create(data) {
    return http.post("/users", data);
  },

  update(id, data) {
    return http.put(`/users/${id}`, data);
  },

  delete(id) {
    return http.delete(`/users/${id}`);
  },
};

export default UserDataService;
