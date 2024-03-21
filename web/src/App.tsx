import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import NewVertexForm from "./NewVertexForm.tsx";
import SearchVerticesView from "./SearchVerticesView.tsx";
import { BrowserRouter, Route, Routes } from "react-router-dom";

const queryClient = new QueryClient();

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <header className={"h-16 w-full flex items-center p-2 border-b shadow"}>
        <h1 className={"font-bold mr-10"}>Graph</h1>
        <nav>
          <ul className={"flex space-x-2"}>
            <li className={"mr-5"}>
              <a href={"/"}>Search</a>
            </li>
            <li>
              <a href={"/create-vertex"}>Create Vertex</a>
            </li>
          </ul>
        </nav>
      </header>
      <main className={"p-5"}>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<SearchVerticesView />} />
            <Route path="/create-vertex" element={<NewVertexForm />} />
          </Routes>
        </BrowserRouter>
      </main>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}
