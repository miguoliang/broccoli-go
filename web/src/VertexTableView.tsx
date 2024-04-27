import { SearchVerticesParams, useSearchVertices } from "./api.ts";

const VertexListView = ({
  q = "",
  page = 0,
  size = 10,
}: Readonly<Partial<SearchVerticesParams>>) => {
  const { data } = useSearchVertices({ q, page, size });
  return (
    <ul>
      {data?.data?.map((vertex) => <li key={vertex.id}>{vertex.name}</li>)}
    </ul>
  );
};

export default VertexListView;
