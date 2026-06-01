import { usersApi } from "@/api";
import { ref } from "vue";
import { getJWTToken, removeJWTToken, setJWTToken } from "./useJWT";
import type { AuthRequest } from "@/api/generated";
import Cookies from 'js-cookie'

const isLoading = ref<boolean>(false);
const isLogedin = ref<boolean>(false);
const errors = ref<unknown[]>([]);

const createAuthRequest = (username: string, password: string): AuthRequest => ({
    username: username.trim(),
    password: password.trim(),
});

const login = async (username: string, password: string) => {
    isLoading.value = true;
    try {
        const resp = await usersApi.login(createAuthRequest(username, password));

        setJWTToken(resp.data.access_token);
        isLogedin.value = true;
        return resp.data.access_token;
    } catch (e) {
        isLogedin.value = false;
        errors.value.push(e);
        throw e;
    } finally {
        isLoading.value = false;
    }
};

const logout = async () => {
    Cookies.set("eee", "ddd")
    Cookies.remove('refresh_token')
    //TODO: add refresh token removal logic
    removeJWTToken();
    isLoading.value = false;
    isLogedin.value = false;
};

const useLogin = () => {
    //TODO: refactor deducing logic
    isLogedin.value = Boolean(Cookies.get("refresh_token")) || Boolean(getJWTToken());1
    return {
        isLoading,
        login,
        logout,
        isLogedin
    };
};

export { useLogin as useLogin };
