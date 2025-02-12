import httpClient from "@/utils/httpclient";

const userService = {
  async getUsers() {
    const res = await httpClient.get("/users");
    return res.data;
  },

  async getUserById(id: number) {
    const res = await httpClient.get(`/users/${id}`);
    return res.data;
  }
}

export default userService;