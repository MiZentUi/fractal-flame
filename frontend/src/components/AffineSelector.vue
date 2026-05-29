<script setup lang="ts">
import { Button } from "@/components/ui/button";
import Input from "@/components/ui/input/Input.vue";
import { ref, watch } from "vue";

const selectedColor = ref<string>("#000000")
watch(selectedColor, (n) => {
    console.log(n)
})

</script>

<template>
    <div class="grid grid-cols-[1fr] gap-2 border border-primary p-2">
        <div class="flex gap-2">
            <span class="w-full flex relative">
                <span class="absolute align-middle h-fit leading-5 w-full text-center pointer-events-none invert grayscale" :style="{color: selectedColor}">
                    {{selectedColor}}
                </span>
                <Input type="color" class="p-0 h-5 w-full p-1/2 py-0 self-center" v-model:model-value="selectedColor" />
            </span>
            <Button variant="ultrakill" class="p-0 h-fit text-xl self-center align-middle leading-4"> x </Button>
        </div>
        <div class="font-stretch-50%">
            <span> x = </span>
            <span><Input class="w-10 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center" placeholder="A" /></span
            ><span>*x + </span>
            <span><Input class="w-10 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center" placeholder="B" /></span>
            <span>*y + </span>
            <span><Input class="w-10 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center" placeholder="C" /></span>
        </div>

        <div class="font-stretch-50%">
            <span> y = </span>
            <span><Input class="w-10 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center" placeholder="D" /></span
            ><span>*x + </span>
            <span><Input class="w-10 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center" placeholder="E" /></span>
            <span>*y + </span>
            <span><Input class="w-10 h-fit p-0.5 pb-0 pt-0 text-end placeholder:text-center" placeholder="F" /></span>
        </div>
    </div>
</template>
