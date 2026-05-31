import type { FractalRequest } from "@/api/generated"
import { inject, reactive, readonly } from "vue"
const stateSymbol = Symbol("GENERATOR_STORE")
const createStore = () => {
    const store = reactive<FractalRequest>({
        affine_params: [],
        functions: [],
        gamma: 2.2,
        height: 1080,
        width: 1920,
        iteration_count: 100000,
        symmetry_level: 1
    })
    return store
}
const useStore = () => inject<FractalRequest>(stateSymbol) as FractalRequest

export {stateSymbol, createStore, useStore}

