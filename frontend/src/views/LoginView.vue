<script setup lang="ts">
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { useLogin } from "@/utils/useLogin";
import { useForm } from "vee-validate";
import { authSchema } from "@/shemas/authSchema";
import type { AxiosError } from "axios";
import type { ApiStatusResponse, AuthRequest } from "@/api/generated";
import FormInput from "@/components/form/FormInput.vue";


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
        toast.success("Login sucsessful");
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
                    <form class="space-y-4 flex flex-col" novalidate @submit.prevent="submit">
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
