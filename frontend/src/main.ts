import './assets/main.css'
import "v3-infinite-loading/lib/style.css";
import InfiniteLoading from "v3-infinite-loading";
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)
app.component("infinite-loading", InfiniteLoading);
app.mount('#app')
