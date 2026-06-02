import { usersApi } from "@/api";
import { ref } from "vue";
import type { AuthRequest } from "@/api/generated";


const isLoading = ref<boolean>(false);
const errors = ref<unknown[]>([]);

const createAuthRequest = (username: string, password: string): AuthRequest => ({
    username: username.trim(),
    password: password.trim(),
});

const register = async (username: string, password: string) => {
    isLoading.value = true;
    try {
        await usersApi.signUp(createAuthRequest(username, password));
    } catch (e) {
        errors.value.push(e);
        throw e;
    } finally {
        isLoading.value = false;
    }
};


const useRegister = () => {
    return {
        isLoading,
        register,
    };
};

export { useRegister };
