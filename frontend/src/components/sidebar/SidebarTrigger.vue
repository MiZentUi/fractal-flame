<script setup lang="ts">
import { useSidebar } from '@/components/ui/sidebar'
import {Button } from '@/components/ui/button'
import { ChevronsLeftRight, ChevronsRightLeft } from 'lucide-vue-next';

const { open, toggleSidebar } = useSidebar()



</script>

<template>
  <Button class="text-2xl p-0" @click="toggleSidebar">
        <ChevronsLeftRight v-if="open" :strokeWidth="3" :size="60"></ChevronsLeftRight>
        <ChevronsRightLeft v-else :strokeWidth="3" :size="60"></ChevronsRightLeft>
  </Button>
</template>
