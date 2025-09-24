import {createI18n} from "vue-i18n"

import zhCN from "@/services/locales/zh-CN.json"
import enUS from "@/services/locales/en-US.json"

export const I18n = createI18n({
	globalInjection: true,
	locale: "zh-CN",
	fallbackLocale: "zh-CN",
	messages: {
		"zh-CN": zhCN,
		"en-US": enUS
	}
})

export default {
	/**
	 * 应用语言
	 * @param lang 语言名称
	 */
	applyLanguage(lang = "zh-CN") {
		I18n.global.locale = lang
		sessionStorage.setItem("language", lang)
	},
	/**
	 * 获取当前语言
	 * @returns {string} 语言名称
	 */
	getLanguage() {
		return sessionStorage.getItem("language") || "zh-CN"
	}
}
