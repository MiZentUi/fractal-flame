<script setup lang="ts">
import {
    Pagination,
    PaginationContent,
    PaginationItem,
    PaginationNext,
    PaginationPrevious,
} from "@/components/ui/pagination";
import PaginationLast from "./ui/pagination/PaginationLast.vue";
import PaginationFirst from "./ui/pagination/PaginationFirst.vue";

const model = defineModel<number>();
const { pageCount, itemsPerPage } = defineProps<{ pageCount: number; itemsPerPage: number }>();
</script>

<template>
    <div class="flex flex-col m-4">
        <Pagination
            @update:page="(val) => (model = val)"
            v
            v-slot="{ page }"
            :items-per-page="itemsPerPage"
            :total="pageCount * itemsPerPage"
            :default-page="1"
        >
            <PaginationContent v-slot="{ items }">
                <PaginationFirst />
                <PaginationPrevious />

                <template v-for="(item, index) in items" :key="index">
                    <PaginationItem v-if="item.type === 'page'" :value="item.value" :is-active="item.value === page">
                        {{ item.value }}
                    </PaginationItem>
                </template>

                <PaginationNext />
                <PaginationLast />
            </PaginationContent>
        </Pagination>
    </div>
</template>
