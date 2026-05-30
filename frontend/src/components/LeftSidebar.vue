<script setup lang="ts">
import { reactive, watch } from "vue";
import type { AffineParams, Function } from "@/api/generated";
import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import FunctionSelector from "@/components/FunctionSelector.vue";
import { AffineSelector } from "@/components/AffineSelector";
import { Button } from "@/components/ui/button";
import { randColor } from "@/utils/color";
import { useStore } from "@/store/useGenearatorStore";

const requesState = useStore()
console.log(requesState)

watch(requesState.functions, console.log);
watch(requesState.affine_params, console.log);

const addFunc = () => {
    requesState.functions?.push({
        name: "",
        weight: 1,
    });
};

const randParam = () => Number((Math.random() * 2 - 1).toFixed(2));

const addTransform = () => {
    requesState.affine_params?.push({
        color: randColor(),
        a: randParam(),
        b: randParam(),
        c: randParam(),
        d: randParam(),
        e: randParam(),
        f: randParam(),
    });
};
const removeTransform = (i: number) => {
    requesState.affine_params?.splice(i, 1);
};

const removeFunc = (i: number) => {
    requesState.functions?.splice(i, 1);
};
</script>

<template>
    <Sidebar :side="'left'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-6 px-4">
                <Card>
                    <CardHeader> Functions </CardHeader>
                    <CardContent class="flex flex-col gap-4 overflow-x-scroll h-38">
                        <FunctionSelector
                            @remove="removeFunc(i)"
                            :function-names="['test']"
                            v-model="requesState.functions[i]"
                            v-for="(_, i) in requesState.functions"
                            :key="i"
                        />
                        <Button class="w-full" variant="outline" @click="addFunc">+</Button>
                    </CardContent>
                </Card>

                <Card class="w-86">
                    <CardHeader> Affine transforms </CardHeader>
                    <CardContent class="flex flex-col gap-6 overflow-x-scroll h-62">
                        <AffineSelector
                            @remove="removeTransform(i)"
                            v-model="requesState.affine_params[i]"
                            v-for="(_, i) in requesState.affine_params"
                        ></AffineSelector>
                        <Button class="w-full" variant="outline" @click="addTransform">+</Button>
                    </CardContent>
                </Card>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
