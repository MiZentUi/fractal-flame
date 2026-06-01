import axios, { AxiosError, isAxiosError } from "axios";
import { FractalsApi, UsersApi } from "./generated";
import { getJWTToken, removeJWTToken, setJWTToken } from "@/utils/useJWT";

const BASE_URL = "http://192.168.0.109:8080/api/v1"
const axiosInstance = axios.create({});


// Track whether a refresh is already in progress to avoid parallel refresh calls
let isRefreshing = false;
let failedQueue : {resolve :(val : unknown) => void, reject: () => void}[] = [];






const processQueue = (error, token = null) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });
  failedQueue = [];
};

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        // Queue the request until the refresh completes
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then((token) => {
            // TODO
            originalRequest.headers["Authorization"] = `Bearer ${token}`;
            return axiosInstance(originalRequest);
          })
          .catch((err) => Promise.reject(err));
      }

      originalRequest._retry = true;
      isRefreshing = true;

      try {
        //TODO
        const { data } = await axiosInstance.post("/auth/refresh", {}, {
            withCredentials: true
        });

        const newToken = data.access_token;
        setJWTToken(newToken)
        //TODO
        axiosInstance.defaults.headers.common["Authorization"] = `Bearer ${newToken}`;

        processQueue(null, newToken);
        return axiosInstance(originalRequest);
      } catch (refreshError) {
        processQueue(refreshError, null);
        // Redirect to login or emit an event

        removeJWTToken()
        //TODO: perform redirect with router
        return Promise.reject(refreshError);
      } finally {
        isRefreshing = false;
      }
    }

    return Promise.reject(error);
  }
);

axiosInstance.interceptors.request.use((config) => {
    // const token =
    //     "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiIxMjMiLCJuYW1lIjoidXNlcjEyMyIsImlhdCI6MTc3OTQ1ODcxNiwiZXhwIjoxNzgyNzgxMjAwfQ.c9k2U8M360aC6wVyr_HNThnHNYqaHCUADX9YelwiEktbucgwR1XhnTsao9r76rfC";
    const token = getJWTToken()
    if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
    }

    return config;
});

axiosInstance.interceptors.response.use(undefined, (error: AxiosError) => {
    if (!isAxiosError(error)) {
        return;
    }

    if (error.response?.status === 401) {
        console.warn("401 Unauthorized");

    }

    return Promise.reject(error);
});

const fractalsApi = new FractalsApi(undefined, BASE_URL, axiosInstance);
const usersApi = new UsersApi(undefined, BASE_URL, axiosInstance)
export { fractalsApi, usersApi };
