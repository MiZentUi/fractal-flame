import { fractalsApi } from "@/api";
import type { FractalsResponse } from "@/api/generated";
import { inject, reactive } from "vue";
interface FracatalStore {
    data: FractalsResponse;
    readonly fetch: (
        page?: number,
        count?: number,
        sort?: string,
        order?: "asc" | "desc",
        userid?: number,
    ) => Promise<FractalsResponse>;
}

const storeSymbol = Symbol("FRACTAL_STORE");
const createFractalsStore = () => {
    const fetchData = (page?: number, count?: number, sort?: string, order?: "asc" | "desc", userid?: number) => {
        return new Promise<FractalsResponse>(async (res, rej) => {
            try {
                const response = (await fractalsApi.getFractals(page, count, sort, order, userid)).data;
                store.data = response;
                res(store.data);
            } catch (e) {
                rej(e);
            }
        });
    };
    const store = reactive<FracatalStore>({
        data: { items: [], page: 0, page_count: 0 },
        fetch: fetchData,
    });
    return store;
};
const useFractalStore = () => inject<FracatalStore>(storeSymbol) as FracatalStore;

export { storeSymbol, createFractalsStore, useFractalStore };
export type { FracatalStore };
