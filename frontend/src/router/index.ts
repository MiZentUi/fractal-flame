import { createRouter, createWebHistory } from "vue-router";
import HomePage from "../views/HomePage.vue";
import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";
import AccountView from "@/views/AccountView.vue";
import { useLogin } from "@/utils/useLogin.ts";
const {isLogedin} = useLogin()

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
            name: "login",
            meta: {
                noAuth: true
            }
        },
        {
            path: "/register",
            component: RegisterView,
            name: "register",
            meta: {
                noAuth: true
            }
        },
        {
            path: "/account",
            component: AccountView,
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
