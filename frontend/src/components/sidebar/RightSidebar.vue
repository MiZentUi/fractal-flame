<script setup lang="ts">
import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import Label from "../ui/label/Label.vue";
import FormInput from "../form/FormInput.vue";
import Tip from "../Tip.vue";
</script>

<template>
    <Sidebar :side="'right'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-4 px-4 w-75">
                <Card>
                    <CardHeader> Genetator settings </CardHeader>
                    <CardContent class="flex flex-col gap-4">
                        <div>
                            <Tip>
                                <template #trigger> <Label>Resolution</Label> </template>
                                <template #content> Target image resolution </template>
                            </Tip>
                            <div class="flex align-baseline items-center gap-4">
                                <FormInput field-name="width" class="w-full text-center" placeholder="Width" />
                                <span>X</span>
                                <FormInput field-name="height" class="w-full text-center" placeholder="Height" />
                            </div>
                        </div>
                        <div class="flex flex-col gap-1">
                            <Tip>
                                <template #trigger> <Label>Iterations</Label> </template>
                                <template #content>
                                    How many steps generator takes to draw the fractal. More steps make the image
                                    smoother and less grainy.
                                </template>
                            </Tip>

                            <FormInput
                                field-name="iteration_count"
                                placeholder="Iterations"
                                class="w-full text-center"
                            />
                        </div>
                        <div class="flex flex-col gap-1">
                            <Tip>
                                <template #trigger> <Label>Symmetry</Label></template>
                                <template #content>
                                    Adds mirror or spinning symmetry (like snowflakes or kaleidoscopes).
                                </template>
                            </Tip>

                            <FormInput
                                field-name="symmetry_level"
                                placeholder="Symmetry level"
                                class="w-full text-center"
                            />
                        </div>
                        <div class="flex flex-col gap-1">
                            <Tip>
                                <template #trigger> <Label>Gamma</Label> </template>
                                <template #content>
                                    A brightness adjustment. Higher values reveal faint, hidden details but can wash out
                                    colors. Start with 2.2 for most screens.
                                </template>
                            </Tip>

                            <FormInput field-name="gamma" placeholder="Gamma" class="w-full text-center" />
                        </div>
                    </CardContent>
                </Card>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
