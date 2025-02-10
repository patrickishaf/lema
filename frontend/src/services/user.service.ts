import httpClient from "@/utils/httpclient";

const userService = {
  async getUsers() {
    const data = await httpClient.get("/users");
    return data;
  },

  async getUserById(id: number) {
    const data = await httpClient.get(`/users/${id}`);
    return data;
  }
}

export default userService;