import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import { createStore, stateSymbol } from "./utils/useStore.ts";

const app = createApp(App);
app.use(router);
app.provide(stateSymbol, createStore());
app.mount("#app");
