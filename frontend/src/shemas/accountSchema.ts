import * as yup from "yup";

export type AccountFormValues = {
    username: string;
    password: string;
    image?: string;
};

const accountSchema = yup.object({
    username: yup
        .string()
        .required("Username is required")
        .max(50, "Username must be at most 50 characters")
        .min(3, "Username must be at least 3 characters")
        .matches(/^[\x00-\x7F]*$/, "Only ASCII characters are allowed"),
    password: yup
        .string()
        .transform((value) => (value === "" ? undefined : value))
        .notRequired()
        .max(50, "Password must be at most 50 characters")
        .min(6, "Password must be at least 6 characters")
        .matches(/^[\x00-\x7F]*$/, "Only ASCII characters are allowed"),
    image: yup.string().notRequired(),
});

export { accountSchema };
