import "../styles/BtnLoader.css";

export default function BtnLoader() {
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