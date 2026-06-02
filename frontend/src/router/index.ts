import { createRouter, createWebHistory } from "vue-router";
import { useLogin } from "@/utils/useLogin.ts";
const {isLogedin} = useLogin()

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            component: () => import("@/views/HomePage.vue"),
            name: "home",
            meta: {
                showFeed: true
            }
        },
        {
            path: "/login",
            component: () => import("@/views/LoginView.vue"),
            name: "login",
            meta: {
                noAuth: true
            }
        },
        {
            path: "/register",
            component: () => import("@/views/RegisterView.vue"),
            name: "register",
            meta: {
                noAuth: true
            }
        },
        {
            path: "/account",
            component: () => import("@/views/AccountView.vue"),
            name: "account",
            meta: {
                requiresAuth: true
            }
        }
    ],
    scrollBehavior() {
        return { top: 0, behavior: "smooth" };
    },
});

router.beforeEach((to) => {
  if ((to.meta.requiresAuth && !isLogedin.value) ||  (to.meta.noAuth && isLogedin.value) ) {
    return { name: 'home' };
  }
});

export default router;
