import type { FractalRequest } from "@/api/generated";
import * as yup from "yup";
import type { Shape } from "./common";



const generatorSchema = yup.object<Shape<FractalRequest>>({
    affine_params: yup.array(
        yup.object({
            color: yup.string(),
            a: yup.number(),
            b: yup.number(),
            c: yup.number(),
            d: yup.number(),
            e: yup.number(),
            f: yup.number(),
        }),
    ).min(1),
    functions: yup.array(
        yup.object({
            name: yup.string(),
            weight: yup.number().positive(),
        }),
    ).min(1),
    gamma: yup.number().positive(),
    height: yup.number().integer().positive().max(2000).min(100),
    width: yup.number().integer().positive().max(2000).min(100),
    iteration_count: yup.number().integer().positive().max(10000000),
    symmetry_level: yup.number().integer().positive().max(10),
});
export {generatorSchema as formSchema}
