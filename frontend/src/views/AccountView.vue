<script setup lang="ts">
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useLogin } from "@/utils/useLogin";
import { ImagePlus, KeyRound, Save, UserIcon } from "lucide-vue-next";
import MultiFractalView from "@/components/MultiFractalView.vue";
import type { ApiStatusResponse, FractalsResponse, UserRequest } from "@/api/generated";
import { computed, ref, watch } from "vue";
import { createUserImageUrl, fractalsApi, usersApi } from "@/api";
import { toast } from "vue-sonner";
import type { AxiosError } from "axios";
import { useForm } from "vee-validate";
import { accountSchema, type AccountFormValues } from "@/shemas/accountSchema";

const { user } = useLogin();

const fractals = ref<FractalsResponse>({
    items: [],
    page: 0,
    page_count: 0,
});

const page = ref<number>(1);
const imagePreview = ref<string>();
const imageInput = ref<HTMLInputElement | null>(null);
const isSaving = ref(false);
const { values, errors, defineField, handleSubmit, resetForm, setFieldError, setFieldValue } =
    useForm<AccountFormValues>({
        validationSchema: accountSchema,
        initialValues: {
            username: user.value?.username ?? "",
            password: "",
            image: undefined,
        },
    });

const [username] = defineField("username");
const [password] = defineField("password");

const fetchPage = async (num: number) => {
    if (!user.value?.id) {
        fractals.value = {
            items: [],
            page: 0,
            page_count: 0,
        };
        return;
    }

    const fratcalsResp = await fractalsApi.getFractals(num, 12, "created", "desc", user.value.id);
    fractals.value = fratcalsResp.data;
};

const avatarSrc = computed(() => {
    if (imagePreview.value) return imagePreview.value;
    if (user.value?.image) return createUserImageUrl(user.value.image);
    return undefined;
});

const hasChanges = computed(() => {
    return Boolean(
        user.value && (values.username.trim() !== user.value.username || values.password?.trim() || values.image),
    );
});

const resetAccountForm = () => {
    resetForm({
        values: {
            username: user.value?.username ?? "",
            password: "",
            image: undefined,
        },
    });
    imagePreview.value = undefined;
    if (imageInput.value) imageInput.value.value = "";
};

const handleImageChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    const file = target.files?.[0];
    setFieldError("image", undefined);

    if (!file) return;
    if (!file.type.startsWith("image/")) {
        setFieldError("image", "Please choose an image file");
        target.value = "";
        return;
    }

    const reader = new FileReader();
    reader.onload = () => {
        const result = String(reader.result);
        imagePreview.value = result;
        setFieldValue("image", result.split(",")[1] ?? result);
    };
    reader.onerror = () => {
        setFieldError("image", "Could not read selected image");
    };
    reader.readAsDataURL(file);
};

const handleSaveError = (e: unknown) => {
    const error = e as AxiosError;
    const response = error.response?.data as ApiStatusResponse | undefined;

    if (response?.message) {
        toast.error(response.message);
        return;
    }

    toast.error(error.message || "Could not update account");
};

const submit = handleSubmit(async (formValues) => {
    if (!user.value) return;
    const request: UserRequest = {};
    const nextUsername = formValues.username.trim();
    const nextPassword = formValues.password?.trim();

    if (nextUsername !== user.value.username) request.username = nextUsername;
    if (nextPassword) request.password = nextPassword;
    if (formValues.image) request.image = formValues.image;

    if (!Object.keys(request).length) return;

    isSaving.value = true;
    try {
        const response = await usersApi.patchUser(request, {
            withCredentials: true,
        });
        user.value = response.data;
        resetAccountForm();
        toast.success("Account updated");
    } catch (error) {
        handleSaveError(error);
    } finally {
        isSaving.value = false;
    }
});

watch(
    user,
    () => {
        resetAccountForm();
        fetchPage(page.value);
    },
    { immediate: true },
);
watch(page, fetchPage);
</script>

<template>
    <div class="pb-10 w-300 m-auto mt-20 space-y-4">
        <Card class="border">
            <CardContent class="grid gap-6 p-4 lg:grid-cols-[auto_1fr] lg:items-start w-100">
                <div class="space-y-3 flex flex-col">
                    <Avatar class="size-30 border bg-primary-foreground m-0" shape="square">
                        <AvatarFallback class="text-2xl">
                            <UserIcon />
                        </AvatarFallback>
                        <AvatarImage v-if="avatarSrc" :src="avatarSrc" class="p-1" />
                    </Avatar>
                    <input ref="imageInput" class="hidden" type="file" accept="image/*" @change="handleImageChange" />
                </div>

                <form class="flex flex-col w-full" novalidate @submit.prevent="submit">
                    <div class="">
                        <Label for="account-username">Username</Label>
                        <Input
                            id="account-username"
                            v-model="username"
                            autocomplete="username"
                            :error="Boolean(errors.username)"
                        />
                        <p v-if="errors.username" class="text-sm text-red-600">
                            {{ errors.username }}
                        </p>
                    </div>

                    <div class="">
                        <Label for="account-password">New password</Label>
                        <div class="relative">
                            <KeyRound
                                class="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-muted-foreground"
                            />
                            <Input
                                id="account-password"
                                v-model="password"
                                type="password"
                                autocomplete="new-password"
                                class="pl-9"
                                :error="Boolean(errors.password)"
                                placeholder="Leave empty to keep current"
                            />
                        </div>
                        <p v-if="errors.password" class="text-sm text-red-600">
                            {{ errors.password }}
                        </p>
                    </div>

                    <p v-if="errors.image" class="text-sm text-red-600 md:col-span-2">
                        {{ errors.image }}
                    </p>

                    <div class="flex gap-2 md:col-span-2"></div>
                </form>
                <Button type="button" variant="outline" class="w-full" @click="imageInput?.click()">
                        <ImagePlus />
                        Image
                    </Button>
                <div class="flex justify-around">

                    <Button type="submit" variant="outline" :disabled="isSaving || !hasChanges" @click="submit">
                        <Save />
                        {{ isSaving ? "Saving..." : "Save changes" }}
                    </Button>
                    <Button type="button" variant="ghost" :disabled="isSaving || !hasChanges" @click="resetAccountForm">
                        Cancel
                    </Button>
                </div>
            </CardContent>
        </Card>
        <MultiFractalView header="Your works" :fractals="fractals" v-model="page" />
    </div>
</template>
