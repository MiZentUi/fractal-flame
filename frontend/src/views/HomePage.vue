<script setup lang="ts">
import FractalView from "@/components/FractalView.vue";
import LeftSidebar from "@/components/sidebar/LeftSidebar.vue";
import RightSidebar from "@/components/sidebar/RightSidebar.vue";
import SidebarTrigger from "@/components/sidebar/SidebarTrigger.vue";
import ButtonGroup from "@/components/ui/button-group/ButtonGroup.vue";

import { ref, watch } from "vue";
import { useForm } from "vee-validate";
import { toast } from "vue-sonner";
import { useSidebar } from "@/components/ui/sidebar";
import { formSchema } from "@/shemas/generatorSchema";
import GenerateButton from "@/components/GenerateButton.vue";
import { createFractalEventsUrl, fractalsApi } from "@/api";
import type { FractalRequest, TaskState } from "@/api/generated";

const { values, errors, validate } = useForm<FractalRequest>({
    validationSchema: formSchema,
    initialValues: {
        affine_params: [],
        functions: [],
        gamma: 2.2,
        height: 1080,
        width: 1920,
        iteration_count: 1000000,
        symmetry_level: 1,
    },
    keepValuesOnUnmount: true
});

watch(values, console.log);
watch(errors, console.error);
const { open } = useSidebar();
const progress = ref<number>(0);
const base64Image = ref<string>("");
const altText = ref<string>("Your fractal will be here");

const subbmit = async () => {
    const { valid, errors } = await validate();
    if (!valid) {
        open.value = true;
        toast.info(String(errors));
        return;
    }

    fractalsApi.generation(values).then((res) => {
        const source = new EventSource(createFractalEventsUrl(res.data.fractal_id), { withCredentials: true });
        altText.value = "Pending...";
        base64Image.value = "";
        // Listen for generic messages
        source.onmessage = (event) => {
            console.log("New message:", event.data);
        };

        source.onopen = () => {
            console.log("Connected");
        };

        // Listen for specific event types (if named by the server)
        source.addEventListener("update", (event) => {
            console.log("Update received:", event.data);
            const data = JSON.parse(event.data) as TaskState;
            progress.value = data.progress;
            base64Image.value = data.preview;
        });
        source.addEventListener("complete", (event) => {
            console.log("completed:", event.data);
            source.close();
        });

        // Handle errors or connection loss
        source.onerror = (err) => {
            console.error("EventSource failed:", err);
        };
    });
};
</script>
<template>
    <div class="bg-transparent sticky min-h-screen w-full">
        <ButtonGroup class="fixed z-50 bottom-16 right-4">
            <GenerateButton :progress="progress" @click="subbmit" />
            <SidebarTrigger variant="outline" class="w-10"></SidebarTrigger>
        </ButtonGroup>

        <FractalView
            :alt="altText"
            class="h-screen absolute top-0 left-0 z-[-9] w-full"
            :href="`data:image/png;base64, ${base64Image}`"
        />
        <LeftSidebar />
        <RightSidebar />
    </div>
</template>
