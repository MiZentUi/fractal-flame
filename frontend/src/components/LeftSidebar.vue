<script setup lang="ts">
import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import FunctionSelector from "@/components/FunctionSelector.vue";
import AffineSelector from "@/components/AffineSelector.vue";
import { Button } from "@/components/ui/button";

</script>

<template>
    <Sidebar :side="'left'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-6 px-4">
                <Card>
                    <CardHeader> Functions </CardHeader>
                    <CardContent class="flex flex-col gap-6 overflow-x-scroll h-44">
                        <FunctionSelector v-for="_ in Array(6)"></FunctionSelector>
                        <Button class="w-full" variant="outline">+</Button>
                    </CardContent>
                </Card>

                <Card class="w-80">
                    <CardHeader> Affine transforms </CardHeader>
                    <CardContent class="flex flex-col gap-6 overflow-x-scroll h-62">
                        <AffineSelector v-for="_ in Array(3)"></AffineSelector>
                        <Button class="w-full" variant="outline">+</Button>
                    </CardContent>
                </Card>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
