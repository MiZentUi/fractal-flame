import * as yup from "yup";
import type { Shape } from "./common";
import type { AuthRequest } from "@/api/generated";
const authSchema = yup.object<Shape<AuthRequest>>({
        password: yup.string().max(50).min(6),
        username: yup.string().max(50).min(3),
    })
export {authSchema}
