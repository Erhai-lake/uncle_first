import {createApp} from "vue"
import router from "@/router"
import {I18n} from "@/services/I18n.js"
import App from "@/App.vue"

// 全局样式
import "@/assets/styles/theme.less"

const APP = createApp(App)

APP.use(router)
APP.use(I18n)
APP.mount("#app")
