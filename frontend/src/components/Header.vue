<script setup lang="ts">
import { useRouter, RouterLink } from "vue-router";
import { Button } from "./ui/button";
import { ref, watch } from "vue";
import { useScroll } from "@vueuse/core";
import { useLogin } from "@/utils/useLogin";
import { UserIcon } from "lucide-vue-next";
import { toast } from "vue-sonner";

const router = useRouter();

const opacity = ref(0);
const { isLogedin, logout } = useLogin();

const { y } = useScroll(window);

watch(y, (n) => {
    const scrollTop = window.scrollY;
    const docHeight = document.documentElement.scrollHeight;
    const winHeight = window.innerHeight;

    const scrollPercent = scrollTop / (docHeight - winHeight);
    opacity.value = scrollPercent;
});

const _logout = () => {
    logout().then(() => {
        router.push({ name: "home" });
        toast.info("Succsessful logout");
    });
};
</script>
<template>
    <header class="w-full h-(--header-height)" :style="{ backgroundColor: 'rgba(0,0,0,' + opacity + ')' }">
        <div class="mx-auto flex max-w-7xl items-center gap-4 px-4 py-4 sm:px-6 lg:px-8 justify-between">
            <RouterLink :to="{ name: 'home' }" class="shrink-0">
                <div class="flex items-center gap-3">
                    <div>
                        <p
                            class="text-lg font-semibold tracking-tight text-primary select-none font-ultrakill hover:text-accent"
                        >
                            FRACTAL FLAME
                        </p>
                    </div>
                </div>
            </RouterLink>

            <div class="items-center gap-3 md:flex">
                <RouterLink v-if="isLogedin" :to="{ name: 'account' }" class="hidden lg:flex">
                    <Button variant="outline" size="icon" class="">
                        <UserIcon></UserIcon>
                    </Button>
                </RouterLink>
                <RouterLink v-if="!isLogedin" :to="{ name: 'login' }">
                    <Button variant="ultrakill">LOGIN</Button>
                </RouterLink>
                <RouterLink v-if="!isLogedin" :to="{ name: 'register' }">
                    <Button variant="ultrakill">REGISTER</Button>
                </RouterLink>
                <Button v-else variant="ultrakill" @click="_logout">Logout</Button>
            </div>
        </div>
    </header>
</template>
