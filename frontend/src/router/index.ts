import { createRouter, createWebHistory } from "vue-router";
import HomePage from "../views/HomePage.vue";
import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";
import AccountView from "@/views/AccountView.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            component: HomePage,
            name: "home",
            meta: {
                showFeed: true
            }
        },
        {
            path: "/login",
            component: LoginView,
            name: "login"
        },
        {
            path: "/register",
            component: RegisterView,
            name: "register"
        },
        {
            path: "/account",
            component: AccountView,
            name: "account"
        }
    ],
    scrollBehavior() {
        return { top: 0, behavior: "smooth" };
    },
});

export default router;
