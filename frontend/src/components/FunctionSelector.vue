<script setup lang="ts">
interface Props {
    functionNames: string[];
}
interface Emits {
    remove: []
}

import type { Function } from "@/api/generated";
import { Button } from "@/components/ui/button";
import Input from "@/components/ui/input/Input.vue";
import Select from "@/components/Select.vue";
import { computed } from "vue";

const model = defineModel<Function>({ required: true });
const { functionNames } = defineProps<Props>();
const items = computed(() =>
    functionNames.map((el) => {
        return { name: el, value: el };
    }),
);
const emit = defineEmits<Emits>()

</script>

<template>
    <div class="flex gap-2 justify-between">
        <Select :items="items" class="w-40" v-model="model.name" placeholder="Function" />
        <Input class="w-18 placeholder:text-center" placeholder="Weight" v-model="model.weight" />
        <Button variant="ultrakill" class="p-0 h-fit text-xl self-center align-middle leading-0" @click="emit('remove')"> x </Button>
    </div>
</template>
