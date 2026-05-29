import { inject, reactive, readonly } from "vue"
const stateSymbol = Symbol("GLOBAL_STATE")
const createStore = () => {
    const store = reactive({
        test: "GGG"
    })
    const update = () => {
        console.log(store)
    }
    return {
        store: readonly(store),
        update
    }
}
const useStore = () => inject<{d:{num:1}}>(stateSymbol) as {d:{num:1}}

export {stateSymbol, createStore, useStore}

