import database from "@/services/Database.js"

export default {
	/**
	 * 应用主题
	 * @param theme 主题名称
	 */
	applyTheme(theme = "light") {
		document.documentElement.setAttribute("data-theme", theme)
		database.add("theme", theme)
		this.getTheme()
	},
	/**
	 * 获取当前主题
	 * @returns {string} 主题名称
	 */
	getTheme() {
		return database.get("theme", "light")
	}
}
