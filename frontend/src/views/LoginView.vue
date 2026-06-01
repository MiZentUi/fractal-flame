<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { toast } from "vue-sonner";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { usersApi } from "@/api";
import { useLogin } from "@/utils/useLogin";
// import { ApiError } from "@/lib/api";
// import { useAuthStore } from "@/stores/auth";

// const auth = useAuthStore();
const router = useRouter();
const route = useRoute();

const {login, isLoading} = useLogin()

const form = reactive({
    name: "",
    email: "",
    password: "",
});
const errors = reactive<Record<"name" | "email" | "password", string>>({
    name: "",
    email: "",
    password: "",
});

function clearErrors() {
    errors.name = "";
    errors.email = "";
    errors.password = "";
}

function validateForm() {
    clearErrors();

    const name = form.name.trim();
    const email = form.email.trim();

    if (!name) errors.name = "Name is required.";
    else if (name.length > 255) errors.name = "Name is too long.";

    if (!form.password) errors.password = "Password is required.";
    else if (form.password.length < 6) errors.password = "Password must be at least 6 characters.";
    else if (form.password.length > 255) errors.password = "Password is too long.";

    return !errors.name && !errors.email && !errors.password;
}
async function submit() {
    if (!validateForm()){ console.error("E"); return;}

    try {
        console.log(await login(form.name, form.password));
        toast.success("Logged in successfully");


        await router.push(
          {name: "home"}
        );
    } catch (error) {
        const message = error instanceof Error ? error.message : "Authentication failed";
        // if (
        //   mode.value === "login" &&
        //   error instanceof ApiError &&
        //   error.status === 401
        // ) {
        //   errors.email = "Invalid email or password.";
        //   errors.password = "Invalid email or password.";
        //   toast.error("Invalid email or password");
        //   return;
        // }

        toast.error(message);
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
                            <Input
                                id="name"
                                v-model="form.name"
                                type="username"
                                autocomplete="username"
                                :aria-invalid="Boolean(errors.name)"
                            />
                            <p v-if="errors.name" class="text-sm text-red-600">
                                {{ errors.name }}
                            </p>
                        </div>
                        <div class="space-y-2">
                            <Label for="password">Password</Label>
                            <Input
                                id="password"
                                v-model="form.password"
                                type="password"
                                autocomplete="password"
                                :aria-invalid="Boolean(errors.password)"
                            />
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
