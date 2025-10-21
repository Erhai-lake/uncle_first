<script setup>
import {onUnmounted, ref} from "vue"
import {useI18n} from "vue-i18n"
import Theme from "@/services/Theme.js"
import NavigationBar from "@/components/NavigationBar.vue"
import Login from "@/components/login.vue"
import EventBus from "@/services/EventBus.js"

// 语言
const {locale} = useI18n()
locale.value = JSON.parse(JSON.stringify(sessionStorage.getItem("language") || "zh-CN"))

// 主题
Theme.applyTheme(Theme.getTheme())

const isLogin = ref(false)
const handleLoginEvent = (value) => {
	isLogin.value = value
}
EventBus.on("isLogin", handleLoginEvent)
onUnmounted(() => {
	EventBus.off("isLogin", handleLoginEvent)
})
</script>

<template>
	<div class="app">
		<navigation-bar/>
		<div class="router">
			<router-view/>
		</div>
	</div>
	<div class="login-container" v-if="isLogin">
		<login/>
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

.login-container {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background-color: var(--box-shadow-color);
	display: flex;
	justify-content: center;
	align-items: center;
}
</style>