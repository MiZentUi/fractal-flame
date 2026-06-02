<script setup lang="ts">
import { ref } from "vue";

import { Sidebar, SidebarContent } from "@/components/ui/sidebar";
import { fractalsApi } from "@/api";
import MultipleFunctionSelector from "../form/MultipleFunctionSelector.vue";
import MultipleAffineSelector from "../form/MultipleAffineSelector.vue";
import { toast } from "vue-sonner";
import type { AxiosError } from "axios";

const functions = ref<string[]>([]);

fractalsApi.getFunctions().then((res) => (functions.value = res.data)).catch((err : AxiosError) => {
    toast.error(err.message)
});

</script>

<template>
    <Sidebar :side="'left'" class="mt-20 border-none">
        <SidebarContent>
            <div class="flex flex-col gap-6 px-4">
                <div class="h-60">
                    <MultipleFunctionSelector :function-names="functions" :field-name="'functions'"/>
                </div>
                <div class="h-62">
                    <MultipleAffineSelector field-name="affine_params"/>
                </div>
            </div>
        </SidebarContent>
    </Sidebar>
</template>
