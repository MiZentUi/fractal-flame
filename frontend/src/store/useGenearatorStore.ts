import type { FractalRequest } from "@/api/generated"
import { inject, reactive, readonly } from "vue"
const stateSymbol = Symbol("GENERATOR_STORE")
const createStore = () => {
    const store = reactive<FractalRequest>({
        affine_params: [],
        functions: [],
        gamma: 0,
        height: 0,
        width: 0,
        iteration_count: 0,
        symmetry_level: 0
    })
    return store
}
const useStore = () => inject<FractalRequest>(stateSymbol)  as FractalRequest

export {stateSymbol, createStore, useStore}

