<script setup lang="ts">
import Card from "./ui/card/Card.vue";
import CardContent from "./ui/card/CardContent.vue";
import CardHeader from "./ui/card/CardHeader.vue";
import Pagination from "./Pagination.vue";
import FractalView from "./FractalView.vue";
import { useFractalStore } from "@/store/useFractalStore.ts";
import { ref, watch } from "vue";
import { createFractalImageUrl } from "@/api";

const fractalStore = useFractalStore();

const page = ref<number>(1);
watch(page, (n) => {
    fractalStore.fetch(n, 12).then(console.log).catch(console.log);
});
fractalStore.fetch(1, 12).then(console.log).catch(console.log);
</script>
<template>
    <section class="bg-primary-foreground flex flex-col h-[calc(100vh-var(--footer-height)-var(--header-height))] w-full z-40 p-0">
        <Card class="w-300 m-auto">
            <CardHeader> Works of others </CardHeader>
            <CardContent class="grid gap-4 p-4 grid-cols-[1fr_1fr_1fr_1fr] items-center justify-items-center">
                <Card
                    class="flex items-center justify-center border h-36 w-full relative"
                    v-for="fractal in fractalStore.data.items"
                >
                    <CardContent class="flex items-center justify-center p-0 h-full">
                        <FractalView class="self-center" :href="createFractalImageUrl(fractal.image)"></FractalView>
                        <span class="absolute bottom-1 left-2"> By: {{ fractal.user_id || "REDACTED" }} </span>
                    </CardContent>
                </Card>
            </CardContent>
        </Card>
        <Pagination :page-count="fractalStore.data.page_count" :items-per-page="12" v-model="page"></Pagination>
    </section>
</template>
