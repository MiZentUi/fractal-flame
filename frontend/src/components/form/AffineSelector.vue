<script setup lang="ts">
interface Props {
    fieldName: string;
}

import { Button } from "@/components/ui/button";
import { useField } from "vee-validate";
import FormInput from "./FormInput.vue";

const inputStyles = "w-12 h-fit p-0.5 pb-0 pt-0 text-center placeholder:text-center";
const { fieldName } = defineProps<Props>();
const colorField = useField(`${fieldName}.color`)
const emit = defineEmits<{ remove: [] }>();
</script>

<template>
    <div class="grid grid-cols-[1fr] gap-2 border border-primary p-2 text-md">
        <div class="flex gap-2">
            <span class="w-full flex relative">
                <span
                    class="absolute self-center align-middle h-fit leading-5 w-full text-center pointer-events-none text-primary-foreground"
                >
                    {{ colorField.value }}
                </span>
                <FormInput
                    :field-name="`${fieldName}.color`"
                    type="color"
                    class="p-0 h-5 w-full p-1/2 py-0 self-center cursor-pointer"
                />
            </span>
            <Button
                variant="ultrakill"
                @click="emit('remove')"
                class="p-0 h-fit text-xl self-center align-middle leading-4"
            >
                x
            </Button>
        </div>
        <div class="font-stretch-50%">
            <span> x = </span>
            <span><FormInput :field-name="`${fieldName}.a`" :class="inputStyles" placeholder="A" /></span><span>*x + </span>
            <span><FormInput :field-name="`${fieldName}.b`" :class="inputStyles" placeholder="B" /></span>
            <span>*y + </span>
            <span><FormInput :field-name="`${fieldName}.c`" :class="inputStyles" placeholder="C" /></span>
        </div>

        <div class="font-stretch-50%">
            <span> y = </span>
            <span><FormInput :field-name="`${fieldName}.d`" :class="inputStyles" placeholder="D" /></span><span>*x + </span>
            <span><FormInput :field-name="`${fieldName}.e`" :class="inputStyles" placeholder="E" /></span>
            <span>*y + </span>
            <span><FormInput :field-name="`${fieldName}.f`" :class="inputStyles" placeholder="F" /></span>
        </div>
    </div>
</template>
