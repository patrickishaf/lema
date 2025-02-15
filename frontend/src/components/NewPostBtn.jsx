import "@/styles/NewPostBtn.css"
import addIc from "../assets/add.svg";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { useState } from "react";
import Loader from "./Loader";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import postsService from "@/services/posts.service";
import { useNavigate } from "react-router-dom";

export default function NewPostBtn({ userId }) {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [isOpen, setIsOpen] = useState(false);
  const navigateTo = useNavigate();

  const queryClient = useQueryClient();
  const {mutateAsync: createPost, isLoading} = useMutation({
    mutationFn: postsService.createPost,
    onSuccess: queryClient.invalidateQueries(['userposts', userId])
  })

  async function handlePublish() {
    try {
      if (title.length > 0 && body.length > 0) {
        await createPost({
          title,
          body,
          userId,
        });
        setTitle("");
        setBody("");
        setIsOpen(false);
      } else {
        window.alert("You have empty fields in your new post");
      }
    } catch (err) {
      console.error("failed to create post", err);
    }
  }

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen} className="h-full">
      <DialogTrigger asChild>
        <button className="new-post-btn border-2 rounded-lg p-6 border-dashed flex flex-col items-center justify-center w-full h-full" onClick={() => setIsOpen(true)}>
          <img src={addIc} alt="add new post" className="w-5 h-5" />
          <p className="new-post custom-pale-txt text-sm font-semibold mt-2">New Post</p>
        </button>
      </DialogTrigger>
      <DialogContent className="sm:min-w-max dialog-content">
        <DialogHeader className="sm:min-w-max dialog-header flex flex-col items-start">
          <DialogTitle className="text-4xl mb-6 font-medium">New Post</DialogTitle>
          <h6 className="label custom-pale-txt text-lg font-medium pb-2">Post Title</h6>
          <input type="text" name="title" className="custom-input-txt h-10 py-2 px-4 outline-none border border-solid rounded w-full" placeholder="Give your post a title" value={title} onChange={(e) => {
            setTitle(e.target.value);
          }} />
          <h6 className="label custom-pale-txt text-lg font-medium pb-2 pt-6">Post Content</h6>
          <textarea name="body" id="" placeholder="Write something mind-blowing" className="custom-input-txt py-2 px-4 outline-none border border-solid rounded w-full h-44 resize-none" value={body} onChange={(e) => {
            setBody(e.target.value);
          }}></textarea>
          <nav className="actions pt-6 flex flex-row gap-2 justify-end ml-auto">
            <DialogClose>
              <button className="custom-pale-txt py-3 px-4 border-solid rounded border text-sm font-normal">Cancel</button>
            </DialogClose>
            <button className="publish py-3 px-4 rounded text-sm font-normal flex flex-row items-center gap-2" onClick={async() => {
              if (title.length == 0 || body.length == 0) {
                window.alert("You have empty fields in your new post");
                return;
              }
              try {
                await createPost({
                  title,
                  body,
                  userId,
                });
                setTitle("");
                setBody("");
                setIsOpen(false);
              } catch (err) {
                console.error("failed to create post. error:", err);
                setTitle("");
                setBody("");
                setIsOpen(false);
              }
            }}>
              <p className="publish-btn-text text-sm font-normal">Publish</p>
              {isLoading && <Loader />}
            </button>
          </nav>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  )
}