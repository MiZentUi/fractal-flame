const set = (key: string, value: string) => {
    localStorage.setItem(key, value)
}
const get = (key : string) => {
    return localStorage.getItem(key)
}
const remove = (key : string) => {
    localStorage.removeItem(key)
}
const useLocalStorage = () => {
    return {set, get, remove}
}
export {useLocalStorage}
