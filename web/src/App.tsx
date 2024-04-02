import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import NewVertexForm from "./NewVertexForm.tsx";
import SearchVerticesView from "./SearchVerticesView.tsx";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import NavigationBar from "./components/NavigationBar.tsx";
import AnonymousActions from "./components/AnonymousActions.tsx";
import userManager, { useAuthStore } from "./oidc.ts";
import { User } from "oidc-client-ts";
import AuthenticatedActions from "./components/AuthenticatedActions.tsx";

const queryClient = new QueryClient();

const queryParams = new URLSearchParams(window.location.search);
const code = queryParams.get("code");
const state = queryParams.get("state");
if (code && state) {
  userManager
    .signinCallback()
    .then((user) => {
      if (user instanceof User) {
        useAuthStore.setState({ user });
        window.history.replaceState({}, document.title, "/");
      }
    })
    .catch((error) => {
      console.error("Error signing in", error);
    });
} else if (useAuthStore.getState().user === null) {
  userManager.getUser().then((user) => {
    if (user instanceof User) {
      useAuthStore.setState({ user });
    }
  });
}

const UserActions = () => {
  const { user } = useAuthStore();
  if (user) {
    return <AuthenticatedActions user={user} />;
  } else {
    return <AnonymousActions />;
  }
};

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter basename={"/"}>
        <header
          className={
            "h-16 w-full grid grid-cols-3 p-2 border-b shadow items-center"
          }
        >
          <h1 className={"font-bold"}>Broccoli</h1>
          <NavigationBar />
          <UserActions />
        </header>
        <main className={"p-5"}>
          <Routes>
            <Route path="/" element={<SearchVerticesView />} />
            <Route path="/create-vertex" element={<NewVertexForm />} />
          </Routes>
        </main>
      </BrowserRouter>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}
