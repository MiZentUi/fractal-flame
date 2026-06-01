<script setup lang="ts">
import Card from "./ui/card/Card.vue";
import CardContent from "./ui/card/CardContent.vue";
import CardHeader from "./ui/card/CardHeader.vue";
import Pagination from "./Pagination.vue";
import FractalView from "./FractalView.vue";
import { reactive, ref, watch } from "vue";
import { createFractalImageUrl, fractalsApi, usersApi } from "@/api";
import { useScroll } from "@/utils/useScroll.ts";
import type { UserResponse, FractalsResponse } from "@/api/generated/api.ts";
import MultiFractalView from "./MultiFractalView.vue";

const fractals = ref<FractalsResponse>({
    items: [],
    page: 0,
    page_count: 0,
});

const { scrollProgress } = useScroll();
const page = ref<number>(1);
const users = reactive<Map<number, UserResponse>>(new Map());
watch(fractals, () => {
    for (const i in fractals.value.items) {
        const frac = fractals.value.items[i];
        if (frac.user_id) usersApi.getUser(frac.user_id).then((user) => users.set(frac.id, user.data));
    }
});

const fetchPage = async (num: number) => {
    const fratcalsResp = await fractalsApi.getFractals(num, 12, "created", "desc");
    fractals.value = fratcalsResp.data;
};

watch(page, fetchPage);
fetchPage(1);
</script>
<template>
    <MultiFractalView
        header="Works of others"
        v-model="page"
        :users="users"
        :fractals="fractals"
        :style="{ backgroundColor: 'rgba(0,0,0,' + scrollProgress + ')' }"
        class="h-[calc(100vh-var(--footer-height)-var(--header-height))]"
    />

</template>
