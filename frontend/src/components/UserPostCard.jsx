import "@/styles/UserPostCard.css"
import delBtn from "../assets/delete.svg";

export default function UserPostCard({ post, onDelete }) {
  function handleDelete(id) {
    onDelete(id);
    console.log({ id });
  }

  return (
    <div className="user-post-card border-solid border rounded-lg p-6 relative">
      <button className="del-btn absolute h-3 w-3" onClick={() => handleDelete(post.id)}>
        <img src={delBtn} alt="delete post" className="h-full w-full" />
      </button>
      <h2 className="title custom-pale-txt text-lg font-medium mb-4">{post.title}</h2>
      <p className="body custom-pale-txt text-sm">{post.body}</p>
    </div>
  )
}