<script setup lang="ts">
import { reactive } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useLogin } from "@/utils/useLogin";
import { useForm } from "vee-validate";
import { authSchema } from "@/shemas/authSchema";
import type { AxiosError } from "axios";
import type { ApiStatusResponse, AuthRequest } from "@/api/generated";
import FormInput from "@/components/form/FormInput.vue";



// const form = reactive({
//     name: "",
//     email: "",
//     password: "",
// });
// const errors = reactive<Record<"name" | "email" | "password", string>>({
//     name: "",
//     email: "",
//     password: "",
// });

// function clearErrors() {
//     errors.name = "";
//     errors.email = "";
//     errors.password = "";
// }

// function validateForm() {
//     clearErrors();

//     const name = form.name.trim();
//     if (!name) errors.name = "Name is required.";
//     else if (name.length > 255) errors.name = "Name is too long.";

//     if (!form.password) errors.password = "Password is required.";
//     else if (form.password.length < 6) errors.password = "Password must be at least 6 characters.";
//     else if (form.password.length > 255) errors.password = "Password is too long.";

//     return !errors.name && !errors.email && !errors.password;
// }
// async function submit() {
//     if (!validateForm()) {
//         return;
//     }

//     try {
//         await login(form.name, form.password);
//         toast.success("Logged in successfully");

//         await router.push({ name: "home" });
//     } catch (error) {
//         const message = error instanceof Error ? error.message : "Authentication failed";
//         // if (
//         //   mode.value === "login" &&
//         //   error instanceof ApiError &&
//         //   error.status === 401
//         // ) {
//         //   errors.email = "Invalid email or password.";
//         //   errors.password = "Invalid email or password.";
//         //   toast.error("Invalid email or password");
//         //   return;
//         // }

//         toast.error(message);
//     }
// }


const { login, isLoading } = useLogin();
const router = useRouter();
const { errors, values, setFieldError } = useForm<AuthRequest>({
    validationSchema: authSchema,
    initialValues: {
        password: "",
        username: "",
    },
});
const handleLoninError = (e: unknown) => {
    const error = e as AxiosError;
    if (!error.response?.data) {
        toast.error(error.message);
        return;
    }

    const response = error.response?.data as ApiStatusResponse;

    switch (response.status) {
        case "NOT_FOUND":
        case "INVALID_ARGUMENT":
            setFieldError("username", "Invalid uesername or password");
            setFieldError("password", "Invalid uesername or password");;
            break;
        default:
            toast.error(response.message);
            break;
    }
};

async function submit() {
    try {
        await login(values.username, values.password);
        toast.success("Account created");
        await router.push({ name: "home" });
    } catch (error) {
        handleLoninError(error)
    }
}


</script>

<template>
    <div class="mx-auto flex max-w-5xl items-center justify-center w-full h-full">
        <div class="grid w-1/2 gap-8 lg:grid-cols-[1fr]">
            <Card class="border-primary bg-transparent">
                <CardHeader>
                    <CardTitle>LOGIN</CardTitle>
                    <CardDescription> Enter your credentials to continue. </CardDescription>
                </CardHeader>
                <CardContent>
                    <form class="space-y-4" novalidate @submit.prevent="submit">
                        <div class="space-y-2">
                            <Label for="name">Username</Label>
                            <FormInput field-name="username" id="name" autocomplete="username" />
                            <p v-if="errors.username" class="text-sm text-red-600">
                                {{ errors.username }}
                            </p>
                        </div>
                        <div class="space-y-2">
                            <Label for="password">Password</Label>
                            <FormInput field-name="password" id="password" type="password" autocomplete="password" />
                            <p v-if="errors.password" class="text-sm text-red-600">
                                {{ errors.password }}
                            </p>
                        </div>

                        <Button variant="outline" type="submit" class="w-full" :disabled="isLoading">
                            {{ isLoading ? "Please wait..." : "LOGIN" }}
                        </Button>
                        <RouterLink :to="{ name: 'register' }">
                            <Button type="button" variant="outline" class="w-full"> Need an account? REGISTER </Button>
                        </RouterLink>
                    </form>
                </CardContent>
            </Card>
        </div>
    </div>
</template>
