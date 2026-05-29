<script setup lang="ts">
import { useRoute, useRouter, RouterLink } from "vue-router";
import { Button } from "./ui/button";
import { ref, watch } from "vue";
import { useScroll, useWindowSize } from "@vueuse/core";
import { useStore } from "@/utils/useStore";

const route = useRoute();
const router = useRouter();

const opacity = ref(0);
const isLoggedIn = ref(false);
const { d } = useStore();

const {y} = useScroll(window)

watch(y, (n) => {
    const scrollTop = window.scrollY; // Current distance from top
    const docHeight = document.documentElement.scrollHeight; // Total document height
    const winHeight = window.innerHeight; // Visible viewport height

    const scrollPercent = (scrollTop / (docHeight - winHeight));
    opacity.value = scrollPercent;
    console.log(opacity.value);
});


</script>
<template>
    <header class="w-full h-18" :style="{ backgroundColor: 'rgba(0,0,0,' + opacity + ')' }">
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
                <RouterLink v-if="isLoggedIn" :to="{ name: 'home' }" class="hidden lg:flex">
                    <Button variant="ghost" size="icon" class="rounded-full">
                        <!-- <UserRound class="size-5" /> -->
                    </Button>
                </RouterLink>
                <RouterLink v-if="!isLoggedIn" :to="{ name: 'login' }">
                    <Button variant="ultrakill">LOGIN</Button>
                </RouterLink>
                <RouterLink v-if="!isLoggedIn" :to="{ name: 'register' }">
                    <Button variant="ultrakill">REGISTER</Button>
                </RouterLink>
                <Button v-else variant="ultrakill" @click="() => console.log('logout')">Logout</Button>
            </div>
        </div>
    </header>
</template>
