import axios, { isAxiosError, type AxiosRequestConfig } from "axios";
import { FractalsApi, UsersApi, type AccessToken } from "./generated";
import { getJWTToken, removeJWTToken, setJWTToken } from "@/utils/useJWT";

const API_BASE_URL = (import.meta as ImportMeta).env?.VITE_API_URL ?? "http://localhost:8080/api";
const REFRESH_COOKIE_PLACEHOLDER = "refresh_token";

type RetriableRequestConfig = AxiosRequestConfig & {
    _retry?: boolean;
};

type FailedRequest = {
    resolve: (token: string) => void;
    reject: (error: unknown) => void;
};

const axiosInstance = axios.create({
    withCredentials: true,
});

const refreshAxiosInstance = axios.create({
    withCredentials: true,
});


let isRefreshing = false;
let failedQueue: FailedRequest[] = [];

const processQueue = (error: unknown, token?: string) => {
    failedQueue.forEach((prom) => {
        if (error || !token) {
            prom.reject(error);
            return;
        }

        prom.resolve(token);
    });
    failedQueue = [];
};

const setAuthorizationHeader = (config: AxiosRequestConfig, token: string) => {
    config.headers = {
        ...config.headers,
        Authorization: `Bearer ${token}`,
    };
};

const refreshApi = new UsersApi(undefined, API_BASE_URL, refreshAxiosInstance);

const refreshAccessToken = async () => {
    const { data } = await refreshApi.refresh(REFRESH_COOKIE_PLACEHOLDER, {
        withCredentials: true,
    });

    const newToken = (data as AccessToken).access_token;
    setJWTToken(newToken);
    axiosInstance.defaults.headers.common.Authorization = `Bearer ${newToken}`;

    return newToken;
};

axiosInstance.interceptors.response.use(
    (response) => response,
    async (error) => {
        if (!isAxiosError(error)) {
            return Promise.reject(error);
        }

        const originalRequest = error.config as RetriableRequestConfig | undefined;

        if (error.response?.status !== 401 || !originalRequest || originalRequest._retry) {
            return Promise.reject(error);
        }

        if (isRefreshing) {
            return new Promise<string>((resolve, reject) => {
                failedQueue.push({ resolve, reject });
            }).then((token) => {
                setAuthorizationHeader(originalRequest, token);
                return axiosInstance(originalRequest);
            });
        }

        originalRequest._retry = true;
        isRefreshing = true;

        try {
            const newToken = await refreshAccessToken();

            processQueue(null, newToken);
            setAuthorizationHeader(originalRequest, newToken);
            return axiosInstance(originalRequest);
        } catch (refreshError) {
            processQueue(refreshError);
            removeJWTToken();
            return Promise.reject(refreshError);
        } finally {
            isRefreshing = false;
        }
    }
);

axiosInstance.interceptors.request.use((config) => {
    const token = getJWTToken();
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    } else {
        delete config.headers.Authorization;
    }

    return config;
});

const createApiUrl = (path: string) => `${API_BASE_URL}/${path.replace(/^\/+/, "")}`;
const createFractalImageUrl = (name: string) => createApiUrl(`/fractals/images/${name}`);
const createFractalEventsUrl = (id: number) => createApiUrl(`/fractals/gen/${id}/events`);
const createUserImageUrl = (name: string) => createApiUrl(`/users/images/${name}`);
const fractalsApi = new FractalsApi(undefined, API_BASE_URL, axiosInstance);
const usersApi = new UsersApi(undefined, API_BASE_URL, axiosInstance);

export {
    API_BASE_URL,
    axiosInstance,
    createApiUrl,
    createFractalEventsUrl,
    createFractalImageUrl,
    createUserImageUrl,
    fractalsApi,
    usersApi,
};
