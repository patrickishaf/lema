export type User = {
  id: number;
  name: string;
  email: string;
  address: string;
}

export type Post = {
  id: number;
  title: string;
  body: string;
  authorId: number;
  createdAt: Date;
  updatedAt: Date;
}