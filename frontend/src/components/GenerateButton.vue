<script setup lang="ts">
import { computed, ref } from "vue";
import { Button } from "./ui/button";
interface Props {
    progress: number;
    showProgress?: boolean;
}
const { progress, showProgress } = defineProps<Props>();
const isHovered = ref<boolean>(false);
const text = computed(() => {
    if (isHovered.value && showProgress) {
            return "CANCEL";
        }
        if (showProgress) {
            return (progress * 100 ).toFixed(0) + "%";
        }
        return "GENERATE";
});


const style = computed(() => {
    if (!showProgress) return {};
    return {
        backgroundImage: `linear-gradient(90deg,var(--accent) ${100 * progress}%,var(--background) ${0}%)`,
    };
});
</script>
<template>
    <Button
        ref="elem"
        variant="outline"
        class="w-50 align-middle transition-all"
        :style="style"
        @mouseenter="isHovered = true"
        @mouseleave="isHovered = false"
    >
        {{ text }}
    </Button>
</template>
