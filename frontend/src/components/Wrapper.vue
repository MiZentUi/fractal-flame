<script setup lang="ts">
import Footer from "./Footer.vue";
import Header from "./Header.vue";
import { SidebarProvider } from "./ui/sidebar/index.ts";
import Feed from "./Feed.vue";
import { Toaster } from "@/components/ui/sonner";
</script>

<template>
    <SidebarProvider>
        <div class="min-h-screen flex flex-col justify-between 6 w-full">
            <Header class="fixed z-10" />
            <section class="w-screen mx-auto min-h-[calc(100vh-18*var(--spacing))]">
                <slot></slot>
            </section>
            <Feed v-if="$route.meta?.showFeed"></Feed>
            <Footer />
        </div>
    </SidebarProvider>
    <Toaster :position="'top-center'" />
</template>
