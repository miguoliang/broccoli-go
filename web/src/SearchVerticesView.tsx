import { SearchVerticesRequest, SearchVerticesResponse } from "./dto.tsx";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export default function SearchVerticesView({
  q = "",
  page = 0,
  size = 10,
}: Readonly<Partial<SearchVerticesRequest>>) {
  const { data } = useQuery({
    queryKey: ["searchVertices", q, page, size],
    queryFn: async ({ queryKey }) => {
      const [, q, page, size] = queryKey;
      const searchParams = new URLSearchParams();
      searchParams.set("q", q.toString());
      searchParams.set("page", page.toString());
      searchParams.set("size", size.toString());
      const url = "/api/vertex?" + searchParams.toString();
      const response = await axios.get<SearchVerticesResponse>(url);
      return response.data;
    },
    initialData: {
      page: 0,
      size: 10,
      total: 0,
      data: [],
    } as SearchVerticesResponse,
  });
  return (
    <ul>
      {data?.data?.map((vertex) => <li key={vertex.id}>{vertex.name}</li>)}
    </ul>
  );
}
