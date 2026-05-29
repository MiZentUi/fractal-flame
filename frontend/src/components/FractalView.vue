<script setup lang="ts">

const props = defineProps(["href"])

</script>
<template>
    <div class="flex justify-center items-center bg-primary-foreground">
        <img :src="props.href" alt="Generated fractal" class="max-h-screen select-none" draggable="false">
    </div>
</template>
