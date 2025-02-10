import "@/styles/Loader.css";

export default function Loader() {
  return (
    <div className="loader w-full flex justify-center items-center">
      <div className="lds-ellipsis">
        <div></div>
        <div></div>
        <div></div>
        <div></div>
      </div>
    </div>
  )
}