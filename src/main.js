import {createApp} from "vue"
import {I18n} from "@/services/I18n.js"
import App from "@/App.vue"

// 全局样式
import "@/assets/styles/theme.less"

const APP = createApp(App)

APP.use(I18n)
APP.mount("#app")
