import httpClient from "@/utils/httpclient";

const userService = {
  async getUsers(pageNumber: number = 0) {
    const res = await httpClient.get(`/users?pageNumber=${pageNumber}`);
    return res.data;
  },

  async getUserById(id: number) {
    const res = await httpClient.get(`/users/${id}`);
    return res.data;
  }
}

export default userService;