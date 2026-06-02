<script setup lang="ts">
import { useSidebar } from '@/components/ui/sidebar'
import {Button } from '@/components/ui/button'
import { computed } from 'vue';

const { open, toggleSidebar } = useSidebar()
const innerText = computed(() => open.value ? "⏴⏵" : "⏵⏴")



</script>

<template>
  <Button class="table-cell text-2xl p-0 pb-1" @click="toggleSidebar">
      {{ innerText }}
  </Button>
</template>
