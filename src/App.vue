<script setup>
import {onMounted} from "vue"
import {useI18n} from "vue-i18n"
import Theme from "@/services/Theme.js"
import NavigationBar from "@/components/NavigationBar.vue"
import Login from "@/components/login.vue"
import database from "@/services/Database.js"

onMounted(() => {
	// 语言
	const {locale} = useI18n()
	locale.value = database.get("language", "zh-CN")
	// 主题
	Theme.applyTheme(Theme.getTheme())
})
</script>

<template>
	<div class="app">
		<navigation-bar/>
		<div class="router">
			<router-view/>
		</div>
	</div>
	<login/>
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
	padding: 0 50px;
	box-sizing: border-box;
	width: 100%;
	height: 100%;
	overflow: auto;
}
</style>