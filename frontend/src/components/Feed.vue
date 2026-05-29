<script setup lang="ts">
import Card from "./ui/card/Card.vue";
import CardContent from "./ui/card/CardContent.vue";
import CardHeader from "./ui/card/CardHeader.vue";
import CardFooter from "./ui/card/CardFooter.vue";

import Pagination from "./Pagination.vue";
import FractalView from "./FractalView.vue";
</script>
<template>
    <section class="bg-primary-foreground w-screen z-40 p-0">
        <Card class="w-300 m-auto">
            <CardHeader> Works of others </CardHeader>
            <CardContent class="grid gap-4 p-4 grid-cols-[1fr_1fr_1fr_1fr_1fr_1fr] items-center justify-items-center">
                <Card class="border w-45" v-for="_ in Array(18)">
                    <CardContent class="grid gap-2 p-2 lg:grid-cols-[auto_1fr] lg:items-center">
                        <FractalView
                            class="self-center"
                            href="https://github.com/MiZentUi/fractal-flame/blob/coursework/images/result6.png?raw=true"
                        ></FractalView>
                    </CardContent>
                    <CardFooter class="py-2 px-4"> by: aaaa </CardFooter>
                </Card>
            </CardContent>
        </Card>
        <Pagination></Pagination>
    </section>
</template>
