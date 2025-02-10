import { createBrowserRouter, RouterProvider } from "react-router-dom";
import routeNames from "./routenames";
import UsersPage from "../components/UsersPage";
import UsersPostsPage from "@/components/UsersPostsPage";

export default function AppRouter() {
  const router = createBrowserRouter([
    {
      path: routeNames.index,
      element: <UsersPage />
    },
    {
      path: routeNames.users,
      element: <UsersPage />
    },
    {
      path: `${routeNames.posts}/:id`,
      element: <UsersPostsPage />
    }
  ]);

  return <RouterProvider router={router} />
}
