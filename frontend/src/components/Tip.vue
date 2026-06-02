<script setup lang="ts">
import Card from "@/components/ui/card/Card.vue";
import CardContent from "@/components/ui/card/CardContent.vue";
import HoverCard from "@/components/ui/hover-card/HoverCard.vue";
import { HoverCardContent, HoverCardTrigger } from "reka-ui";
</script>
<template>
    <HoverCard :open-delay="100" :close-delay="100">
        <HoverCardTrigger as-child>
            <span class="hover:underline select-none">
                <slot name="trigger"></slot>
            </span>
        </HoverCardTrigger>
        <HoverCardContent class="z-50" align="start">
            <Card>
                <CardContent class="p-2 max-w-80 text-sm">
                    <slot name="content"></slot>
                </CardContent>
            </Card>
        </HoverCardContent>
    </HoverCard>
</template>
