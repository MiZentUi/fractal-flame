<script setup lang="ts">
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";

interface Item {
    name: string;
    value: any;
}
interface Props {
    items: Item[];
    placeholder?: string;
    label?: string;

}

const props = defineProps<Props>();
const model = defineModel<any>();


</script>

<template>
    <Select v-model="model" :disabled="!props.items?.length">
        <SelectTrigger :class="$attrs.class">
            <SelectValue :placeholder="props.placeholder || ''"/>
        </SelectTrigger>
        <SelectContent class="w-(--reka-select-trigger-width)">
            <SelectGroup>
                <SelectLabel v-if="props.label">{{ props.label }}</SelectLabel>
                <SelectItem v-for="el in props.items" :value="el.value"> {{ el.name }} </SelectItem>
            </SelectGroup>
        </SelectContent>
    </Select>
</template>
