<script setup lang="ts">
import { Card, CardContent, CardHeader } from "./ui/card";
import Pagination from "./Pagination.vue";
import FractalView from "./FractalView.vue";
import { createFractalImageUrl } from "@/api";
import type { FractalsResponse, UserResponse } from "@/api/generated";
import { ref } from "vue";
const model = defineModel<number>();

const { fractals, users } = defineProps<{
    fractals: FractalsResponse;
    users?: Map<number, UserResponse>;
    header: string;
}>();
</script>

<template>
    <section class="bg-primary-foreground flex flex-col w-full z-40 p-0">
        <Card class="w-300 m-auto">
            <CardHeader class="text-2xl"> {{ header }}</CardHeader>
            <CardContent class="grid gap-4 p-4 grid-cols-[1fr_1fr_1fr_1fr] items-center justify-items-center">
                <Card
                    class="flex items-center justify-center border h-36 w-full relative"
                    v-for="fractal in fractals.items"
                >
                    <CardContent class="flex items-center justify-center p-0 h-full">
                        <FractalView class="self-center" :href="createFractalImageUrl(fractal.image)"></FractalView>
                        <span v-if="users" class="absolute bottom-1 left-2">
                            By: {{ users.get(fractal.id)?.username || "REDACTED" }}
                        </span>
                    </CardContent>
                </Card>
            </CardContent>
        </Card>
        <Pagination :page-count="fractals.page_count" :items-per-page="12" v-model="model"></Pagination>
    </section>
</template>
