<script setup lang="ts">
import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import Label from "./ui/label/Label.vue";
import FormInput from "./FormInput.vue";
</script>

<template>
    <Sidebar :side="'right'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-4 px-4 w-75">
                <Card>
                    <CardHeader> Genetator settings </CardHeader>
                    <CardContent class="flex flex-col gap-4">
                        <div>
                            <Label>Resolution</Label>
                            <div class="flex align-baseline items-center gap-4">
                                <FormInput
                                    field-name="width"
                                    class="w-full text-center"
                                    placeholder="Width"
                                />
                                <span>X</span>
                                <FormInput
                                    field-name="height"
                                    class="w-full text-center"
                                    placeholder="Height"
                                />
                            </div>
                        </div>
                        <div class="flex flex-col gap-1">
                            <Label>Iterations</Label>
                            <FormInput
                                field-name="iteration_count"
                                placeholder="Iterations"
                                class="w-full text-center"
                            />
                        </div>
                        <div class="flex flex-col gap-1">
                            <Label>Symmetry</Label>
                            <FormInput
                                field-name="symmetry_level"
                                placeholder="Symmetry level"
                                class="w-full text-center"
                            />
                        </div>
                        <div class="flex flex-col gap-1">
                            <Label>Gamma</Label>
                            <FormInput
                                field-name="gamma"
                                placeholder="Gamma"
                                class="w-full text-center"
                            />
                        </div>
                    </CardContent>
                </Card>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
