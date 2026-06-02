// import type { FractalRequest } from "@/api/generated";
// import { inject, reactive, readonly } from "vue";

// interface FunctionInput {
//     name: string;
//     weight: number;
// }

// interface AffineParamsInput {
//     color: string;
//     a: string;
//     b: string;
//     c: string;
//     d: string;
//     e: string;
//     f: string;
// }

// interface GeneratorInputState {
//     affine_params: AffineParamsInput[];
//     functions: FunctionInput[];
//     gamma: string;
//     height: string;
//     width: string;
//     iteration_count: string;
//     symmetry_level: string;
// }

// interface GlobalAppState {
//     generatorInput: GeneratorInputState;
// }

// const stateSymbol = Symbol("GENERATOR_STORE");
// const createStore = () => {
//     const store = reactive<GlobalAppState>({
//         generatorInput: {
//             affine_params: [],
//             functions: [],
//             gamma: "",
//             height: "",
//             width: "",
//             iteration_count: "",
//             symmetry_level: "",
//         },
//     });
//     return store;
// };
// const useStore = () => inject<GlobalAppState>(stateSymbol) as GlobalAppState;

// export { stateSymbol, createStore, useStore };
// export type {GlobalAppState, GeneratorInputState, FunctionInput, AffineParamsInput}
