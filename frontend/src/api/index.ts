import axios, { AxiosError, isAxiosError } from "axios";
import { FractalsApi } from "./generated";

const axiosInstance = axios.create();

axiosInstance.interceptors.request.use((config) => {
    const token =
        "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiIxMjMiLCJuYW1lIjoidXNlcjEyMyIsImlhdCI6MTc3OTQ1ODcxNiwiZXhwIjoxNzgyNzgxMjAwfQ.c9k2U8M360aC6wVyr_HNThnHNYqaHCUADX9YelwiEktbucgwR1XhnTsao9r76rfC";

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

const fractalsApi = new FractalsApi(undefined, "http://localhost:8080/api/v1", axiosInstance);
export { fractalsApi };
