import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import NavigationBar from "./components/NavigationBar.tsx";
import AnonymousActions from "./components/AnonymousActions.tsx";
import userManager, { useAuthStore } from "./oidc.ts";
import { User } from "oidc-client-ts";
import AuthenticatedActions from "./components/AuthenticatedActions.tsx";
import Profile from "./Profile.tsx";
import Graph from "./Graph.tsx";
import ColorModeSwitch from "./components/ColorModeSwitch.tsx";
import useColorModeStore from "./stores/ColorModeStore.tsx";

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
    if (user instanceof User && !user.expired) {
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
  const { colorMode } = useColorModeStore();
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter basename={"/"}>
        <div className={`h-screen flex flex-col ${colorMode}`}>
          <header className={"w-full border-b shadow"}>
            <div
              className={
                "h-16 p-2 grid grid-cols-3 max-w-7xl items-center mx-auto"
              }
            >
              <h1 className={"font-bold"}>Broccoli</h1>
              <NavigationBar />
              <div className="flex items-center ml-auto space-x-10">
                <ColorModeSwitch />
                <UserActions />
              </div>
            </div>
          </header>
          <main className="flex-grow">
            <Routes>
              <Route path="/profile" element={<Profile />} />
              <Route path="/" element={<Graph />} />
            </Routes>
          </main>
        </div>
      </BrowserRouter>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}
