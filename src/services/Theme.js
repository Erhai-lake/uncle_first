export default {
	/**
	 * 应用主题
	 * @param theme 主题名称
	 */
	applyTheme(theme = "light") {
		document.documentElement.setAttribute("data-theme", theme)
		sessionStorage.setItem("theme", theme)
		this.getTheme()
	},
	/**
	 * 获取当前主题
	 * @returns {string} 主题名称
	 */
	getTheme() {
		console.log(sessionStorage.getItem("theme"))
		return sessionStorage.getItem("theme") || "light"
	}
}
