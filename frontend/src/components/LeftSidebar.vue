<script setup lang="ts">
import { ref } from "vue";

import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import FunctionSelector from "@/components/FunctionSelector.vue";
import { AffineSelector } from "@/components/AffineSelector";
import { Button } from "@/components/ui/button";
import { randColor } from "@/utils/color";
import { type AffineParamsInput, type FunctionInput } from "@/store/useGenearatorStore";
import { useFieldArray } from "vee-validate";
import { fractalsApi } from "@/api";


const randParam = () => (Math.random() * 2 - 1).toFixed(2);

const randomAffine = () => {
    return {
        color: randColor(),
        a: randParam(),
        b: randParam(),
        c: randParam(),
        d: randParam(),
        e: randParam(),
        f: randParam(),
    };
};

const functions = ref<string[]>([]);

fractalsApi.getFunctions().then((res) => (functions.value = res.data));

const { push: pushFunc, fields: funcFields, remove: removeFunc } = useFieldArray<FunctionInput>("functions");
const {
    push: pushAffine,
    fields: affineFields,
    remove: removeAffine,
} = useFieldArray<AffineParamsInput>("affine_params");
</script>

<template>
    <Sidebar :side="'left'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-6 px-4">
                <div class="h-60">
                    <Card>
                        <CardHeader> Functions </CardHeader>
                        <CardContent class="flex flex-col gap-4 overflow-x-scroll max-h-38 w-90">
                            <FunctionSelector
                                :function-names="functions"
                                v-for="(field, i) in funcFields"
                                :key="i"
                                v-model="field.value"
                                @remove="removeFunc(i)"
                                :field-name="`functions[${i}]`"
                            />
                            <Button class="w-full" variant="outline" @click="pushFunc({ name: '', weight: 1 })"
                                >+</Button
                            >
                        </CardContent>
                    </Card>
                </div>

                <div class="h-62">
                    <Card class="w-86">
                        <CardHeader> Affine transforms </CardHeader>
                        <CardContent class="flex flex-col gap-6 overflow-x-scroll max-h-62">
                            <AffineSelector
                                v-model="field.value"
                                v-for="(field, i) in affineFields"
                                @remove="removeAffine(i)"
                                :key="i"
                                :field-name="`affine_params[${i}]`"
                            ></AffineSelector>
                            <Button class="w-full" variant="outline" @click="pushAffine(randomAffine())">+</Button>
                        </CardContent>
                    </Card>
                </div>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
