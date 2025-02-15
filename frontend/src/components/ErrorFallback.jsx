export default function ErrorFallback({ errorName, info }) {
  return (
    <div className="err-fallback w-full flex flex-col items-center justify-center">
      <p className="text-3xl font-medium custom-pale-txt pt-6 pb-2">{errorName}</p>
      <p className="text-xl font-medium custom-pale-txt pb-5">{info}</p>
    </div>
  )
}