<script setup lang="ts">
import { useScroll } from "@/utils/useScroll";

const { scrollProgress } = useScroll();
</script>
<template>
    <footer
        :style="{ backgroundColor: 'rgba(0,0,0,' + scrollProgress + ')' }"
        class="h-(--footer-height) z-1 px-4 py-4 text-left text-sm text-accent bg-primary-foreground"
    >
        &copy; 2026 FRACTAL FLAME. All rights reserved.
    </footer>
</template>
