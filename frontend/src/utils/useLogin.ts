import { usersApi } from "@/api";
import { ref } from "vue";
import { getJWTToken, removeJWTToken, setJWTToken } from "./useJWT";
import type { AuthRequest, UserResponse } from "@/api/generated";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

const isLoading = ref<boolean>(false);
const isLogedin = ref<boolean>(false);
const user = ref<UserResponse>();
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
        const tokenData = jwtDecode(resp.data.access_token);
        if (tokenData.sub) {
            const userResponse = await usersApi.getUser(Number(tokenData.sub));
            user.value = userResponse.data;
        }
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
    Cookies.remove("refresh_token");
    //TODO: add refresh token removal logic
    removeJWTToken();
    isLoading.value = false;
    isLogedin.value = false;
    user.value = undefined;
};

const useLogin = () => {
    //TODO: refactor deducing logic
    const token = getJWTToken();
    isLogedin.value = Boolean(Cookies.get("refresh_token")) || Boolean(token);
    if (token) {
        const tokenData = jwtDecode(token);
        if (tokenData.sub) {
            usersApi.getUser(Number(tokenData.sub)).then((resp) => (user.value = resp.data));
        }
    }

    return {
        isLoading,
        login,
        logout,
        isLogedin,
        user,
    };
};

export { useLogin as useLogin };
