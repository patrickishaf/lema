export default function ErrorFallback({ errorName, info }) {
  return (
    <div className="err-fallback w-full h-screen flex items-center justify-center">
      <p className="text-6xl font-medium custom-pale-txt">{errorName}</p>
      <p className="text-3xl font-medium custom-pale-txt">{info}</p>
    </div>
  )
}