<script setup lang="ts">
import { fractalsApi } from "@/api";
import FractalView from "@/components/FractalView.vue";
import LeftSidebar from "@/components/LeftSidebar.vue";
import RightSidebar from "@/components/RightSidebar.vue";
import SidebarTrigger from "@/components/SidebarTrigger.vue";
import { Button } from "@/components/ui/button";
import ButtonGroup from "@/components/ui/button-group/ButtonGroup.vue";

import { useStore } from "@/store/useGenearatorStore";
import { watch } from "vue";
import * as yup from "yup";
import { useForm } from "vee-validate";
import { type FractalRequest } from "@/api/generated";
import { toast } from 'vue-sonner'

export type ConditionalSchema<T> = T extends string
    ? yup.StringSchema
    : T extends number
      ? yup.NumberSchema
      : T extends boolean
        ? yup.BooleanSchema
        : T extends Record<any, any>
          ? yup.AnyObjectSchema
          : T extends Array<any>
            ? yup.ArraySchema<any, any>
            : yup.AnySchema;

export type Shape<Fields> = {
    [Key in keyof Fields]: ConditionalSchema<Fields[Key]>;
};

const formSchema = yup.object<Shape<FractalRequest>>({
    affine_params: yup.array(
        yup.object({
            color: yup.string(),
            a: yup.number(),
            b: yup.number(),
            c: yup.number(),
            d: yup.number(),
            e: yup.number(),
            f: yup.number(),
        }),
    ),
    functions: yup.array(
        yup.object({
            name: yup.string(),
            weight: yup.number().positive(),
        }),
    ),
    gamma: yup.number().positive(),
    height: yup.number().integer().positive().max(2000),
    width: yup.number().integer().positive().max(2000),
    iteration_count: yup.number().integer().positive().max(100000000),
    symmetry_level: yup.number().integer().positive().max(20),
});

const { values, errors } = useForm({
    validationSchema: formSchema,
});

watch(values, console.log);
watch(errors, console.error);
</script>
<template>
    <div class="bg-transparent sticky min-h-screen w-full">
        <ButtonGroup class="fixed z-50 bottom-16 right-4">
            <Button variant="outline" class="w-50 align-middle" @click="toast.info('ddd')"> GENERATE </Button>
            <SidebarTrigger variant="outline" class="w-10"></SidebarTrigger>
        </ButtonGroup>

        <FractalView
            class="h-screen absolute top-0 left-0 z-[-9] w-screen"
            href="https://github.com/MiZentUi/fractal-flame/blob/coursework/images/result6.png?raw=true"
        />
        <LeftSidebar />
        <RightSidebar />
    </div>
</template>
z
