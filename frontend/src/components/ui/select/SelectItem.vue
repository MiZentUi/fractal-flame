<script setup lang="ts">
import type { SelectItemProps } from "reka-ui";
import type { HTMLAttributes } from "vue";
import { reactiveOmit } from "@vueuse/core";
import { Check } from "lucide-vue-next";
import { SelectItem, SelectItemIndicator, SelectItemText, useForwardProps } from "reka-ui";
import { cn } from "@/lib/utils";

const props = defineProps<SelectItemProps & { class?: HTMLAttributes["class"] }>();

const delegatedProps = reactiveOmit(props, "class");

const forwardedProps = useForwardProps(delegatedProps);
</script>

<template>
    <SelectItem
        v-bind="forwardedProps"
        :class="
            cn(
                'relative flex w-full cursor-default aria-selected:text-accent select-none items-center rounded-sm py-1.5 pl-2 pr-8 text-sm outline-none focus:text-accent  data-[disabled]:pointer-events-none data-[disabled]:opacity-50',
                props.class,
            )
        "
    >
        <span class="absolute right-2 flex h-3.5 w-3.5 items-center justify-center">
            <SelectItemIndicator> • </SelectItemIndicator>
        </span>

        <SelectItemText>
            <slot />
        </SelectItemText>
    </SelectItem>
</template>
