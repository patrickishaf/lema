import AppRouter from "./navigation/router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {ReactQueryDevtools} from "@tanstack/react-query-devtools"
import ErrorBoundary from "./components/ErrorBoundary";
import { useEffect } from "react";
import { pageStore } from "./utils/pagestore";

const queryClient = new QueryClient();

function App() {
  useEffect(() => {
    return pageStore.clear();
  }, []);
  
  return (
    <ErrorBoundary>
      <QueryClientProvider client={queryClient}>
        <AppRouter />
        <ReactQueryDevtools />
      </QueryClientProvider>
    </ErrorBoundary>
  )
}

export default App
