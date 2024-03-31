import axios, { AxiosError, AxiosRequestConfig } from "axios";
import qs from "qs";
import { useAuthStore } from "./oidc.ts";

const customAxios = axios.create({
  paramsSerializer: (params) => {
    // Serialize query parameters as JSON
    return qs.stringify(params, { arrayFormat: "repeat" });
  },
});

customAxios.interceptors.request.use(
  (config) => {
    console.log("intercepting request");
    const user = useAuthStore.getState().user;
    if (user) {
      // Add the id token to the Authorization header
      console.log("Adding id token to request: ", user.id_token);
      config.headers.Authorization = `Bearer ${user.id_token}`;
    }
    return config;
  },
  (error) => {
    // Handle errors during request creation
    return Promise.reject(error);
  },
);

// add a second `options` argument here if you want to pass extra options to each generated query
export const customInstance = <T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig,
): Promise<T> => {
  const source = axios.CancelToken.source();
  const promise = customAxios({
    ...config,
    ...options,
    cancelToken: source.token,
  }).then(({ data }) => data);

  // @ts-ignore
  promise.cancel = () => {
    source.cancel("Query was cancelled");
  };

  return promise;
};

// In some case with react-query and swr you want to be able to override the return error type, so you can also do it here like this
export type ErrorType<Error> = AxiosError<Error>;

export type BodyType<BodyData> = BodyData;
