import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {ReactQueryDevtools} from "@tanstack/react-query-devtools";
import NewVertexForm from "./NewVertexForm.tsx";
import SearchVerticesView from "./SearchVerticesView.tsx";
import {useState} from "react";

const queryClient = new QueryClient();

export default function App() {

    const [i, setI] = useState(0);

    return (
        <QueryClientProvider client={queryClient}>
            <h1>Create Vertex</h1>
            <NewVertexForm/>
            <button onClick={() => setI(i + 1)}>Refresh</button>
            <SearchVerticesView key={i} />
            <ReactQueryDevtools initialIsOpen={false}/>
        </QueryClientProvider>
    );
}