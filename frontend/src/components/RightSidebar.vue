<script setup lang="ts">
import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import Input from "./ui/input/Input.vue";
import { useStore } from "@/store/useGenearatorStore.ts";
import Label from "./ui/label/Label.vue";
import NumberInput from "./NumberInput.vue";

const requesState = useStore();
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
                                <NumberInput
                                    num-type="natural"
                                    v-model="requesState.width"
                                    class="w-full text-center"
                                    placeholder="Width"
                                />
                                <span>X</span>
                                <NumberInput
                                    num-type="natural"
                                    v-model="requesState.height"
                                    class="w-full text-center"
                                    placeholder="Height"
                                />
                            </div>
                        </div>
                        <div class="flex flex-col gap-1">
                            <Label>Iterations</Label>
                            <NumberInput
                                num-type="natural"
                                v-model="requesState.iteration_count"
                                placeholder="Iterations"
                                class="w-full text-center"
                            />
                        </div>
                        <div class="flex flex-col gap-1">
                            <Label>Symmetry</Label>
                            <NumberInput
                                num-type="natural"
                                v-model="requesState.symmetry_level"
                                placeholder="Symmetry level"
                                class="w-full text-center"
                            />
                        </div>
                        <div class="flex flex-col gap-1">
                            <Label>Gamma</Label>
                            <NumberInput
                                num-type="real"
                                v-model="requesState.gamma"
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
