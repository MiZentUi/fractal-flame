<script setup lang="ts">
import { Mail } from "lucide-vue-next";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Card, CardContent, CardHeader } from "@/components/ui/card";

import CardFooter from "@/components/ui/card/CardFooter.vue";
import FractalView from "@/components/FractalView.vue";
import Pagination from "@/components/Pagination.vue";

</script>

<template>
    <div class="pb-10 w-300 m-auto mt-20 space-y-4">
        <Card class="border">
            <CardContent class="grid gap-4 p-4 lg:grid-cols-[auto_1fr] lg:items-center">
                <Avatar class="size-30 border bg-primary-foreground" shape="square">
                    <AvatarFallback class="text-2xl"> SK </AvatarFallback>
                    <AvatarImage src="https://avatars.githubusercontent.com/u/141739318?v=4" class="p-1"> </AvatarImage>
                </Avatar>

                <div class="space-y-2">
                    <div class="flex flex-wrap items-center gap-3">
                        <h2 class="text-2xl tracking-tight">skrr</h2>
                    </div>
                    <div class="flex flex-wrap gap-4 text-xl">
                        <span class="inline-flex items-center gap-2"><Mail class="size-4" /> kree@ba.sh</span>
                    </div>
                </div>
            </CardContent>
        </Card>
        <Card class="border">
            <CardHeader class="text-3xl p-4 gap-y-1"> Your Works </CardHeader>
            <CardContent class="grid gap-4 p-4 grid-cols-[1fr_1fr_1fr_1fr_1fr_1fr] items-center justify-items-center">
                <Card class="border w-45"  v-for="_ in Array(18)">
                    <CardContent class="grid gap-2 p-2 lg:grid-cols-[auto_1fr] lg:items-center">
                        <FractalView class="self-center" href="https://github.com/MiZentUi/fractal-flame/blob/coursework/images/result6.png?raw=true"></FractalView>
                    </CardContent>
                    <CardFooter class="py-2 px-4">
                        by: aaaa
                    </CardFooter>
                </Card>
            </CardContent>
        </Card>
        <Pagination></Pagination>
    </div>
</template>
