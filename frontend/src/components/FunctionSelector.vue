<script setup lang="ts">
interface Props {
    functionNames: string[];
}
interface Emits {
    remove: [];
}

import type { Function } from "@/api/generated";
import { Button } from "@/components/ui/button";
import Input from "@/components/ui/input/Input.vue";
import Select from "@/components/Select.vue";
import { computed } from "vue";
import { Label } from "./ui/label";
import NumberInput from "./NumberInput.vue";
import { fractalsApi } from "@/api/index.ts";
import { ref } from "vue";

const model = defineModel<Function>({ required: true });
const { functionNames } = defineProps<Props>();
const items = computed(() =>
    functions.value.map((el) => {
        return { name: el, value: el };
    }),
);
// const items = computed(() =>
//     functionNames.map((el) => {
//         return { name: el, value: el };
//     }),
// );
const emit = defineEmits<Emits>();
const functions = ref<string[]>([])


fractalsApi.getFunctions().then((res) => functions.value = res.data)

</script>

<template>
    <div class="flex gap-2 justify-between">
        <span class="flex flex-col gap-1 w-3/2">
            <Label>Function type</Label>
            <Select :items="items" class="w-full" v-model="model.name" placeholder="Function" />
        </span>

        <span class="flex flex-col gap-1 w-1/2">
            <Label>Weight</Label>
            <div class="flex gap-2 justify-between">
                <NumberInput
                    num-type="positive"
                    class="w-18 placeholder:text-center"
                    placeholder="Weight"
                    v-model="model.weight"
                />
                <Button
                    variant="ultrakill"
                    class="p-0 h-fit text-xl self-center align-middle leading-0"
                    @click="emit('remove')"
                >
                    x
                </Button>
            </div>
        </span>
    </div>
</template>
