import {ref} from "vue"
import database from "@/services/Database.js"

const theme = ref(database.get("theme", "light"))

export default {
	/**
	 * 应用主题
	 * @param newTheme  主题名称
	 */
	applyTheme(newTheme = "light") {
		theme.value = newTheme
		document.documentElement.setAttribute("data-theme", newTheme)
		database.add("theme", newTheme)
	},
	/**
	 * 获取当前主题
	 * @returns {string} 主题名称
	 */
	getTheme() {
		return theme
	}
}
