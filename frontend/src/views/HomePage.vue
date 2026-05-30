<script setup lang="ts">
import FractalView from "@/components/FractalView.vue";
import LeftSidebar from "@/components/LeftSidebar.vue";
import RightSidebar from "@/components/RightSidebar.vue";
import SidebarTrigger from "@/components/SidebarTrigger.vue";
import { Button } from "@/components/ui/button";
import ButtonGroup from "@/components/ui/button-group/ButtonGroup.vue";

import { useStore } from "@/store/useGenearatorStore";

const requesState = useStore()
</script>
<template>
    <div class="bg-transparent sticky min-h-screen w-full">
        <ButtonGroup class="fixed z-50 bottom-16 right-4">
            <Button variant="outline" class="w-50 align-middle"> GENERATE </Button>
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
