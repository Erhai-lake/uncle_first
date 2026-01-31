<script setup>
import {onMounted} from "vue"
import {useI18n} from "vue-i18n"
import Theme from "@/services/Theme.js"
import Database from "@/services/Database.js"
import NavigationBar from "@/components/NavigationBar.vue"

/**
 * 主题
 */
const THEME = Theme.getTheme()

onMounted(() => {
	// 语言
	const {locale, t} = useI18n()
	locale.value = Database.get("language", "zh-CN")
	// 主题
	Theme.applyTheme(THEME.value)
	// 设置标题
	document.title = t("title")
})
</script>

<template>
	<div class="app">
		<navigation-bar/>
		<div class="router">
			<router-view/>
		</div>
	</div>
</template>

<style scoped lang="less">
.app {
	width: 100%;
	height: 100vh;
	display: grid;
	grid-template-rows: auto 1fr;
	overflow: auto;
}

.router {
	box-sizing: border-box;
	width: 100%;
	height: 100%;
	overflow: auto;
}
</style>