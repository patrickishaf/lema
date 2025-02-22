import "../styles/UsersPostsPage.css";
import { useState } from "react";
import UserPostCard from "./UserPostCard";
import useVector from "@/hooks/useVector";
import uuid from "react-uuid";
import NewPostBtn from "./NewPostBtn";
import { useNavigate, useParams } from "react-router-dom";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import userService from "@/services/user.service";
import Loader from "./Loader";
import postsService from "@/services/posts.service";

export default function UsersPostsPage() {
  const navigateTo = useNavigate();
  const { id: userId } = useParams();
  const queryClient = useQueryClient();

  const {data: userData, isLoading: isUserLoading, isError: isUserError, error: userError, refetch: userRefetch } = useQuery({
    queryFn: () => userService.getUserById(userId),
    queryKey: ["user", userId],
    enabled: !!userId,
  })
  const {data: postsData, isLoading: isPostsLoading, isError: isPostsError, error: postsError, refetch: postsRefetch } = useQuery({
    queryFn: () => postsService.getPostsByUserId(userId),
    queryKey: ["userposts", userId],
    enabled: !!userId,
  })
  const {mutateAsync: deletePost} = useMutation({
    mutationFn: postsService.deletePostById,
    onSuccess: queryClient.invalidateQueries(['user', userId])
  })
  const {prevBtn} = useVector();

  return (
    <div className="users-posts-page pt-44 pb-44">
      {
        isUserLoading ? (
          <Loader />
        ) : (
          <>
          <button className="flex items-center gap-3" onClick={() => {
            navigateTo(-1);
          }}>
            <img src={prevBtn} alt="back to users" />
            <p className="custom-pale-txt text-sm font-semibold">Back to Users</p>
          </button>
          {
            isUserError ? (
              <p>Failed to fetch user</p>
            ) : (
              <>
              <h1 className="page-title text-6xl font-medium my-4">{userData?.name}</h1>
              <p className="email text-sm custom-pale-txt mb-6">{userData?.email}
                {
                  !isPostsLoading && !isPostsError && postsData && <span className="font-medium">{` • ${postsData.length} posts`}</span>
                }
              </p>
              </>
            )
          }
          {
            isPostsLoading && <Loader />
          }
          {
            !isPostsError && !isPostsLoading && (
              <main className="cards-box w-full grid lg:grid-cols-3 md:grid-cols-2 gap-6">
                <NewPostBtn userId={userId} />
                {
                  postsData?.map((post) => (
                    <UserPostCard
                      key={uuid()}
                      post={post}
                      onDelete={async () => {
                        try {
                          await deletePost(post.id);
                        } catch (err) {
                          console.error("failed to delete post");
                          console.error(err);
                        }
                      }}
                    />
                  ))
                }
              </main>
            )
          }
          </>
        )
      }
    </div>
  )
}