import {createI18n} from "vue-i18n"

import zhCN from "@/assets/lang/zh-CN.json"
import enUS from "@/assets/lang/en-US.json"

export const I18n = createI18n({
	legacy: false,
	globalInjection: true,
	locale: "zh-CN",
	fallbackLocale: "zh-CN",
	messages: {
		"zh-CN": zhCN,
		"en-US": enUS
	}
})
