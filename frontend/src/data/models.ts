export type Address = {
  id: string;
  user_id: string;
  street: string;
  state: string;
  city: string;
  zipcode: string;
}

export type User = {
  id: string;
  name: string;
  username: string;
  email: string;
  address: Address
}

export type Post = {
  id: string;
  title: string;
  body: string;
  user_id: string;
  created_at: Date;
}