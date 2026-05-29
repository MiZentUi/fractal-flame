<script setup lang="ts">
import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import Input from "./ui/input/Input.vue";
</script>

<template>
    <Sidebar :side="'right'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-4 px-4">
                <Card>
                    <CardHeader> Genetator settings </CardHeader>
                    <CardContent class="flex flex-col gap-4">
                        <div class="flex align-baseline items-center gap-4">
                            <Input class="w-20" placeholder="Width" />
                            <span>X</span>
                            <Input class="w-20" placeholder="Height"/>
                        </div>
                        <Input placeholder="Iteration cpount" />
                        <Input placeholder="Symmetry level" />
                        <Input placeholder="Gamma" />
                    </CardContent>
                </Card>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
