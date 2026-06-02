<script setup lang="ts">
import FractalView from "@/components/FractalView.vue";
import LeftSidebar from "@/components/sidebar/LeftSidebar.vue";
import RightSidebar from "@/components/sidebar/RightSidebar.vue";
import SidebarTrigger from "@/components/sidebar/SidebarTrigger.vue";
import ButtonGroup from "@/components/ui/button-group/ButtonGroup.vue";

import { onUnmounted, ref } from "vue";
import { useForm } from "vee-validate";
import { toast } from "vue-sonner";
import { useSidebar } from "@/components/ui/sidebar";
import { formSchema } from "@/shemas/generatorSchema";
import GenerateButton from "@/components/GenerateButton.vue";
import { createFractalEventsUrl, fractalsApi } from "@/api";
import type { FractalRequest, TaskState } from "@/api/generated";

const { open } = useSidebar();
const progress = ref<number>(0);
const base64Image = ref<string>("");
const altText = ref<string>("Your fractal will be here");
const currentEventSouce = ref<EventSource>();
const streaming = ref<boolean>(false)


const { values, validate } = useForm<FractalRequest>({
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
    keepValuesOnUnmount: true,
});

const onSSEUpdate = (event: MessageEvent<any>) => {
    const data = JSON.parse(event.data) as TaskState;
    progress.value = data.progress;
    base64Image.value = data.preview;
};

const onSSEComplite = (_: MessageEvent<any>) => {
    //TODO: make it update feed
    closeImageEventStream();
};
const onSSEError = (_: Event) => {
    toast.error("Generation failed")
    closeImageEventStream();
};

const connectToImageEventStream = (url: string) => {
    closeImageEventStream();
    openImageEventStream(url)
    altText.value = "Pending...";
    base64Image.value = "";
    currentEventSouce.value?.addEventListener("update", onSSEUpdate);
    currentEventSouce.value?.addEventListener("complete", onSSEComplite);
    currentEventSouce.value?.addEventListener("error", onSSEError);
};

const closeImageEventStream = () => {
    currentEventSouce.value?.close();
    streaming.value = false
    currentEventSouce.value = undefined;
}

const openImageEventStream = (url: string) => {
    currentEventSouce.value = new EventSource(url);
    streaming.value = true
}

const subbmit = async () => {
    if (streaming.value) {
        closeImageEventStream();
        return;
    }
    const { valid, errors } = await validate();
    if (!valid) {
        open.value = true;
        for (const [_, value] of Object.entries(errors)) {
            toast.info(value as string);
        }
        return;
    }
    try {
        const resp = await fractalsApi.generation(values, {withCredentials: true});
    progress.value = 0;
    connectToImageEventStream(createFractalEventsUrl(resp.data.fractal_id));
    }
    catch (e) {
        toast.error(String(e))
    }
};

onUnmounted(() => {
    closeImageEventStream();
});
</script>
<template>
    <div class="bg-transparent sticky min-h-screen w-full">
        <ButtonGroup class="fixed z-50 bottom-16 right-4">
            <GenerateButton :progress="progress" :show-progress="streaming" @click="subbmit"/>
            <SidebarTrigger variant="outline" class="w-10"></SidebarTrigger>
        </ButtonGroup>

        <FractalView
            :alt="altText"
            class="h-screen fixed top-0 left-0 z-[-9] w-full"
            :href="base64Image ? `data:image/png;base64, ${base64Image}` : undefined"
            :noOpen="false"
        />
        <LeftSidebar />
        <RightSidebar />
    </div>
</template>
