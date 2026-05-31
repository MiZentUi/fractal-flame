<script setup lang="ts">
import type { AffineParams } from "@/api/generated";
import { Button } from "@/components/ui/button";
import Input from "@/components/ui/input/Input.vue";
import NumberInput from "../NumberInput.vue";

const inputStyles = "w-12 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center";

const model = defineModel<AffineParams>({ required: true });
const emit = defineEmits<{ remove: [] }>();
</script>

<template>
    <div class="grid grid-cols-[1fr] gap-2 border border-primary p-2 text-md">
        <div class="flex gap-2">
            <span class="w-full flex relative">
                <span
                    class="absolute self-center align-middle h-fit leading-5 w-full text-center pointer-events-none text-primary-foreground"
                >
                    {{ model.color }}
                </span>
                <Input
                    type="color"
                    class="p-0 h-5 w-full p-1/2 py-0 self-center cursor-pointer"
                    v-model:model-value="model.color"
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
            <span><NumberInput v-model="model.a" :class="inputStyles" placeholder="A" /></span><span>*x + </span>
            <span><NumberInput v-model="model.b" :class="inputStyles" placeholder="B" /></span>
            <span>*y + </span>
            <span><NumberInput v-model="model.c" :class="inputStyles" placeholder="C" /></span>
        </div>

        <div class="font-stretch-50%">
            <span> y = </span>
            <span><NumberInput v-model="model.d" :class="inputStyles" placeholder="D" /></span><span>*x + </span>
            <span><NumberInput v-model="model.e" :class="inputStyles" placeholder="E" /></span>
            <span>*y + </span>
            <span><NumberInput v-model="model.f" :class="inputStyles" placeholder="F" /></span>
        </div>
    </div>
</template>
