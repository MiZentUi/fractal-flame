import { usersApi } from "@/api";
import { ref } from "vue";
import { getJWTToken, removeJWTToken, setJWTToken } from "./useJWT";
import { toast } from "vue-sonner";
import axios from "axios";

const isLoading = ref<boolean>(false);
const isLogedin = ref<boolean>(false);
const errors = ref<unknown[]>([]);

const login = async (username: string, password: string) => {
    isLoading.value = true;
    let ret = undefined;
    try {
        // const resp = await usersApi.login({ username: username.trim(), password: password.trim() }, {
        //     withCredentials: true
        // });
        const resp = await axios.post('http://192.168.0.109:8080/api/v1/users/login',{ username: username.trim(), password: password.trim() }, {
            withCredentials: true
        });
        setJWTToken(resp.data.access_token);
        isLogedin.value = true;
        ret = resp.data.access_token;
    } catch (e) {
        isLogedin.value = false;
        errors.value.push(e);
        toast.info(`Could not login ${e}`);
    } finally {
        isLoading.value = false;
    }
    return ret;
};

const logout = () => {
    //TODO: add refresh token removal logic
    removeJWTToken();
    isLoading.value = false;
    isLogedin.value = false;
};

const useLogin = () => {
    // refactor deducing logic
    if(getJWTToken()) isLogedin.value = true;
    return {
        isLoading,
        login,
        logout,
        isLogedin,
    };
};

export { useLogin };
