import { useLocalStorage } from "./useLocalStorage";

const TOKEN_KEY = "access_token";
const { get, set, remove } = useLocalStorage();

const getJWTToken = () => {
    return get(TOKEN_KEY);
};

const setJWTToken = (token: string) => {
    return set(TOKEN_KEY, token);
};

const removeJWTToken = () => {
    return remove(TOKEN_KEY);
};

export {getJWTToken, setJWTToken, removeJWTToken}
