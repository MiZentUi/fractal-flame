<script setup lang="ts">

const props = defineProps(["href", "alt"])

</script>
<template>
    <div class="flex justify-center items-center bg-primary-foreground h-full">
        <img :src="props.href" :alt="props.alt || 'Fractal'" class="max-h-full select-none" draggable="false">
    </div>
</template>
