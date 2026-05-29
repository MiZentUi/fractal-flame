<script setup lang="ts">
import { Button } from "@/components/ui/button";
import Input from "@/components/ui/input/Input.vue";
import Select from "@/components/Select.vue";


</script>

<template>
    <div class="flex gap-4 justify-between">
        <Select :items="[{name: 'test', value: 'ee'}]" class="w-35" placeholder="Function" />
        <Input class="w-20 placeholder:text-center" placeholder="Weight" />
        <Button variant="ultrakill" class="p-0 h-fit text-xl self-center align-middle leading-0"> x </Button>
    </div>
</template>
