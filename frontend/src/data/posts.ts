import { Post } from "./models";

const body = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dol...';

export const posts: Post[] = [
  {
    id: 1,
    title: 'Title 1',
    body: body,
    user_id: 1,
    created_at: new Date(),
    updatedAt: new Date(),
  },
  {
    id: 2,
    title: 'Title 2',
    body: body,
    user_id: 1,
    created_at: new Date(),
    updatedAt: new Date(),
  },
  {
    id: 3,
    title: 'Title 3',
    body: body,
    user_id: 1,
    created_at: new Date(),
    updatedAt: new Date(),
  },
  {
    id: 4,
    title: 'Title 4',
    body: body,
    user_id: 1,
    created_at: new Date(),
    updatedAt: new Date(),
  },
  {
    id: 5,
    title: 'Title 5',
    body: body,
    user_id: 1,
    created_at: new Date(),
    updatedAt: new Date(),
  },
  {
    id: 6,
    title: 'Title 6',
    body: body,
    user_id: 1,
    created_at: new Date(),
    updatedAt: new Date(),
  },
]

export function getPosts() {
  return posts;
}