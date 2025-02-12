import httpClient from "@/utils/httpclient";

const postsService = {
  async getPostsByUserId(userId: number) {
    const posts = await httpClient.get(`/posts?userId=${userId}`);
    return posts.data;
  },

  async deletePostById(id: number) {
    const result = await httpClient.delete(`/posts/${id}`);
    return result.data;
  },

  async createPost(data: { title: string, body: string, userId: number }) {
    const post = await httpClient.post("/posts", {
      author_id: Number(data.userId),
      title: data.title,
      body: data.body,
    });
    return post.data;
  }
}

export default postsService;