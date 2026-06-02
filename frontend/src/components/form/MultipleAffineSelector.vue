<script setup lang="ts">
interface Props {
    fieldName: string;
}

import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";

import AffineSelector from "@/components/form/AffineSelector.vue";
import { Button } from "@/components/ui/button";
import { randColor } from "@/utils/color";
import { useField, useFieldArray } from "vee-validate";
import type { AffineParams } from "@/api/generated";
import Tip from "../Tip.vue";

const randParam = () => Number((Math.random() * 2 - 1).toFixed(2));

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

const { fieldName } = defineProps<Props>();
const field = useField(fieldName);
const { push, fields, remove } = useFieldArray<AffineParams>(fieldName);

const pushNew = () => {
    push(randomAffine());
    field.validate();
};
const removeV = (i: number) => {
    remove(i);
    field.validate();
};
</script>

<template>
    <Card class="w-90" :class="field.errors.value.length > 0 ? 'border-accent' : ''">
        <CardHeader>
            <Tip>
                <template #trigger> Affine transforms </template>
                <template #content>
                    Basic moves: stretch, rotate, slide, or skew parts of the image. Each transform also has a weight and a color.
                </template>
            </Tip>
        </CardHeader>
        <CardContent class="flex flex-col gap-6 overflow-y-auto max-h-62 w-88 pr-4">
            <AffineSelector
                v-model="field.value"
                v-for="(field, i) in fields"
                @remove="removeV(i)"
                :key="i"
                :field-name="`affine_params[${i}]`"
            ></AffineSelector>
            <Button class="w-full" variant="outline" @click="pushNew">+</Button>
        </CardContent>
    </Card>
</template>
