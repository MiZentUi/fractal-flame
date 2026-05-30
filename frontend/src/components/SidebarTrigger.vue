<script setup lang="ts">
import { useSidebar } from '@/components/ui/sidebar'
import {Button } from '@/components/ui/button'
import { watch, ref } from 'vue';

const { open, toggleSidebar } = useSidebar()
const innerText = ref("⏴⏵")
watch(open, (val) => {
    innerText.value = val ? "⏵⏴" : "⏴⏵"
})


</script>

<template>
  <Button class="table-cell text-2xl p-0 pb-1" @click="toggleSidebar">
      {{ innerText }}
  </Button>
</template>
