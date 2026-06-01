<script setup lang="ts">
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import CardHeader from "@/components/ui/card/CardHeader.vue";
import FunctionSelector from "./FunctionSelector.vue";
import { Button } from "@/components/ui/button";
import { useField, useFieldArray } from "vee-validate";
import type { Function } from "@/api/generated/api.ts";

interface Props {
    functionNames: string[];
    fieldName: string;
}

const { fieldName, functionNames } = defineProps<Props>();
const field = useField(fieldName);
const { push, fields, remove } = useFieldArray<Function>(fieldName);

const pushNew = () => {
    push({ name: functionNames[Math.round((functionNames.length - 1) * Math.random())], weight: 1 });
    field.validate();
};
const removeV = (i: number) => {
    remove(i);
    field.validate();
};
</script>

<template>
    <Card :class="field.errors.value.length > 0 ? 'border-accent' : ''">
        <CardHeader> Functions </CardHeader>
        <CardContent class="flex flex-col gap-4 overflow-y-auto max-h-38 w-90">
            <FunctionSelector
                :function-names="functionNames"
                v-for="(field, i) in fields"
                :key="i"
                v-model="field.value"
                @remove="removeV(i)"
                :field-name="`${fieldName}[${i}]`"
            />
            <Button class="w-full" variant="outline" @click="pushNew">+</Button>
        </CardContent>
    </Card>
</template>
