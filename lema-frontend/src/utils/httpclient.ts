import axios from "axios";

const httpClient = {
  client: axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
  }),

  async get(path: string) {
    const res = await this.client.get(path);
    return res.data;
  },

  async post(path: string, data: object) {
    const res = await this.client.post(path, data);
    return res.data;
  },

  async delete(path: string) {
    const res = await this.client.delete(path);
    console.log({ res });
    return res.data;
  }
}

export default httpClient;