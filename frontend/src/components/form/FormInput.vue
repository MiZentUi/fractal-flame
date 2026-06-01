<script setup lang="ts">
import { useField } from "vee-validate";
import Input from "../ui/input/Input.vue";

const { fieldName } = defineProps<{ fieldName: string }>();
const { value, errors } = useField<string>(fieldName);

</script>

<template>
    <Input :error="errors.length > 0" v-model="value"/>
</template>
