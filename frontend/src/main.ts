import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
// import { createStore, stateSymbol } from "./store/useGenearatorStore.ts";
import { createFractalsStore, storeSymbol } from "./store/useFractalStore.ts";
const app = createApp(App);
app.use(router);
// app.provide(stateSymbol, createStore());
app.provide(storeSymbol, createFractalsStore());

app.mount("#app");
