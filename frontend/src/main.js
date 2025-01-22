import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import CanvasJSChart from "@canvasjs/vue-charts";
import { dataStore, getAllData } from "./script";

const datas = await getAllData();
dataStore.setDatas(datas);

const app = createApp(App);

app.use(CanvasJSChart);
app.mount("#app");
