<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useLogin } from "@/utils/useLogin";
import type { AxiosError } from "axios";
import type { ApiStatusResponse, AuthRequest } from "@/api/generated";
import { useForm } from "vee-validate";
import FormInput from "@/components/form/FormInput.vue";
import * as yup from "yup";
import { useRegister } from "@/utils/useRegister";
import { authSchema } from "@/shemas/authSchema";

const { register, isLoading } = useRegister();
const router = useRouter();
const { errors, values, setFieldError } = useForm<AuthRequest>({
    validationSchema: authSchema,
    initialValues: {
        password: "",
        username: "",
    },
});
const handleRegError = (e: unknown) => {
    const error = e as AxiosError;
    if (!error.response?.data) {
        toast.error(error.message);
        return;
    }

    const response = error.response?.data as ApiStatusResponse;

    switch (response.status) {
        case "ALREADY_EXISTS":
            setFieldError("username", "Name is already taken.");
            break;
        case "INVALID_ARGUMENT":
            setFieldError("password", "Password is too weak");
            break;
        default:
            toast.error(response.message);
            break;
    }
};

async function submit() {
    try {
        await register(values.username, values.password);
        toast.success("Account created");
        await router.push({ name: "home" });
    } catch (error) {
        handleRegError(error)
    }
}
</script>

<template>
    <div class="mx-auto flex max-w-5xl items-center justify-center w-full h-full">
        <div class="grid w-1/2 gap-8 lg:grid-cols-[1fr]">
            <Card class="border-primary bg-transparent">
                <CardHeader>
                    <CardTitle>CREATE ACCOUNT</CardTitle>
                    <CardDescription> Enter your credentials to create account </CardDescription>
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

                        <Button
                            variant="outline"
                            type="submit"
                            class="w-full"
                            :disabled="isLoading || errors.password || errors.username"
                        >
                            {{ isLoading ? "Please wait..." : "CREATE ACCOUNT" }}
                        </Button>
                        <RouterLink :to="{ name: 'login' }">
                            <Button type="button" variant="outline" class="w-full">
                                {{ "Already have an account? LOGIN" }}
                            </Button>
                        </RouterLink>
                    </form>
                </CardContent>
            </Card>
        </div>
    </div>
</template>
