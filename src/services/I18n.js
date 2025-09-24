import {createI18n} from "vue-i18n"

import zhCN from "@/services/locales/zh-CN.json"

export const I18n = createI18n({
    globalInjection: true,
    locale: "zh-CN",
    fallbackLocale: "zh-CN",
    messages: {
        "zh-CN": zhCN
    }
})
