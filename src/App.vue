<script setup>
import I18n from "@/services/I18n.js"
import Theme from "@/services/Theme.js"
import NavigationBar from "@/components/NavigationBar.vue"

I18n.applyLanguage(I18n.getLanguage())
Theme.applyTheme(Theme.getTheme())
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
	padding: 0 50px;
	box-sizing: border-box;
	width: 100%;
	height: 100%;
	overflow: auto;
}
</style>