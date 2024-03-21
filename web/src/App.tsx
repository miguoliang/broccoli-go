import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import NewVertexForm from "./NewVertexForm.tsx";
import SearchVerticesView from "./SearchVerticesView.tsx";
import { Fragment, useState } from "react";
import { Menu, Transition } from "@headlessui/react";
import { ChevronDownIcon } from "@heroicons/react/20/solid";

const queryClient = new QueryClient();

export default function App() {
  const [i, setI] = useState(0);

  return (
    <QueryClientProvider client={queryClient}>
      <header className={"h-32 w-full flex justify-between p-2"}>
        <h1>Graph</h1>
        <Menu as={"div"} className={"relative inline-block text-left"}>
          <div>
            <Menu.Button className="inline-flex w-full justify-center rounded-md bg-black/20 px-4 py-2 text-sm font-medium text-white hover:bg-black/30 focus:outline-none focus-visible:ring-2 focus-visible:ring-white/75">
              Options
              <ChevronDownIcon
                className="-mr-1 ml-2 h-5 w-5 text-violet-200 hover:text-violet-100"
                aria-hidden="true"
              />
            </Menu.Button>
          </div>
          <Transition
            as={Fragment}
            enter="transition ease-out duration-100"
            enterFrom="transform opacity-0 scale-95"
            enterTo="transform opacity-100 scale-100"
            leave="transition ease-in duration-75"
            leaveFrom="transform opacity-100 scale-100"
            leaveTo="transform opacity-0 scale-95"
          >
            <Menu.Items>
              <Menu.Item>{() => <>Create Vertex</>}</Menu.Item>
            </Menu.Items>
          </Transition>
        </Menu>
      </header>

      <NewVertexForm />
      <button onClick={() => setI(i + 1)}>Refresh</button>
      <SearchVerticesView key={i} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}
