import httpClient from "@/utils/httpclient";

const postsService = {
  async getPostsByUserId(userId: number) {
    const posts = await httpClient.get(`/posts?userId=${userId}`);
    return posts;
  },

  async deletePostById(id: number) {
    const result = await httpClient.delete(`/posts/${id}`);
    return result;
  },

  async createPost(data: { title: string, body: string, userId: number }) {
    const post = await httpClient.post("/posts", {
      author_id: Number(data.userId),
      title: data.title,
      body: data.body,
    });
    return post;
  }
}

export default postsService;