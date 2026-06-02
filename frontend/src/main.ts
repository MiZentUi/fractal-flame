import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import 'viewerjs/dist/viewer.css'
import VueViewer from 'v-viewer'

import { createFractalsStore, storeSymbol } from "./store/useFractalStore.ts";
const app = createApp(App);
app.use(VueViewer);
app.use(router);
app.provide(storeSymbol, createFractalsStore());

app.mount("#app");
