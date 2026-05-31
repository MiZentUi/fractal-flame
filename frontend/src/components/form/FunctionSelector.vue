<script setup lang="ts">
interface Props {
    functionNames: string[];
    fieldName: string
}
interface Emits {
    remove: [];
}

import { Button } from "@/components/ui/button";
import Select from "@/components/Select.vue";
import { computed } from "vue";
import { Label } from "../ui/label/index.ts";
import FormInput from "./FormInput.vue";
import { useField } from "vee-validate";

const { functionNames, fieldName } = defineProps<Props>();
const emit = defineEmits<Emits>();
const items = computed(() =>
    functionNames.map((el) => {
        return { name: el, value: el };
    }),
);

const {value: name} = useField(`${fieldName}.name`)
</script>

<template>
    <div class="flex gap-2 justify-between">
        <span class="flex flex-col gap-1 w-3/2">
            <Label>Function type</Label>
            <Select :items="items" class="w-full" v-model="name" placeholder="Function" />
        </span>

        <span class="flex flex-col gap-1 w-1/2">
            <Label>Weight</Label>
            <div class="flex gap-2 justify-between">
                <FormInput
                    :field-name="`${fieldName}.weight`"
                    class="w-18 placeholder:text-center"
                    placeholder="Weight"
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
