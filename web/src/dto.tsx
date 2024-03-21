export type PageResponse<T = any> = {
  page: number;
  size: number;
  total: number;
  data: Array<T>;
};

export type PageParams = {
  page: number;
  size: number;
};

export type ErrorResponse = {
  message: string;
};

export type CreateVertexRequest = {
  name: string;
  type: string;
};

export type CreateVertexResponse = {
  id: number;
};

export type CreatePropertyRequest = {
  key: string;
  value: string;
};

export type CreateEdgeRequest = {
  from: number;
  to: number;
  type: string;
};

export type CreateEdgeResponse = {
  id: number;
};

export type Vertex = {
  id: number;
  name: string;
  type: string;
  properties: Array<Property>;
};

export type Property = {
  id: number;
  name: string;
  value: string;
};

export type Edge = {
  id: number;
  from: number;
  to: number;
  type: string;
};

export type SearchVerticesRequest = PageParams & { q?: string };

export type SearchVerticesResponse = PageResponse<Vertex>;

export type SearchEdgesRequest = PageParams & {
  from: Array<string>;
  to: Array<string>;
  type: Array<string>;
};

export type SearchEdgesResponse = PageResponse<Edge>;
