<script setup lang="ts">
import { computed } from "vue";
import { Button } from "./ui/button";
interface Props {
    progress: number;
}
const { progress } = defineProps<Props>();
const defaultText = "GENERATE";
const text = computed(() => {
    if (progress != 0 && progress < 1) {
        return progress.toFixed(2);
    }
    return defaultText;
});

const style = computed(() => {
    if(progress == 0 || progress >= 1) return {}
    return {
        backgroundImage: `linear-gradient(90deg,var(--accent) ${100 * progress}%,var(--background) ${0}%)`
    };
});
</script>
<template>
    <Button variant="outline" class="w-50 align-middle transition-all" :style="style"> {{ text }} </Button>
</template>
z
