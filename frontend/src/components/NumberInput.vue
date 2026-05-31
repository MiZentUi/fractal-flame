<script setup lang="ts">

enum NumberType{
    REAL =  "real",
    INT = "integer",
    NAT = "natural",
    POS = "positive"
}

interface Props { numType?: "real" | "integer" | "natural" | "positive"}

import { ref, watch } from "vue";
import Input from "./ui/input/Input.vue";

const model = defineModel<number>();
const inputValue = ref<string>(model.value?.toString() || "");
const { numType = "real" } = defineProps<Props>();

const regExps = {
    [NumberType.REAL] : /^-?\d+\.?\d*$/,
    [NumberType.INT]: /^-?\d+$/,
    [NumberType.NAT]: /^\d+$/,
    [NumberType.POS]: /^\d+\.?\d*$/,
}

watch(inputValue, (n) => {
    model.value = Number(n);
});

watch(model, (n) => {
    inputValue.value = n ? n.toString() : "";
});

let lastInput = inputValue.value;
const update = (e: InputEvent) => {
    const emptyLineRegEx = /\s+|(^$)/;


    const target = e.target as HTMLInputElement;

    const substrs = target.value.match(regExps[numType as NumberType]);
    const newValue = emptyLineRegEx.test(target.value) ? "" : substrs ? substrs[0] : lastInput;
    target.value = inputValue.value = newValue;
    lastInput = inputValue.value;
};
</script>

<template>
    <Input :model-value="inputValue" :value="inputValue" @input="update" />
</template>
