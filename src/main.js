import {createApp} from "vue"
import {i18n} from "@/services/i18n"
import App from "@/App.vue"

const APP = createApp(App)

APP.use(i18n)
APP.mount("#app")
