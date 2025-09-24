import {createApp} from "vue"
import {I18n} from "@/services/I18n.js"
import App from "@/App.vue"

const APP = createApp(App)

APP.use(I18n)
APP.mount("#app")
