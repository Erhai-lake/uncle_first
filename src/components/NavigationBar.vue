<script setup>
import { ref } from "vue"
import Theme from "@/services/Theme.js"

const THEME = ref(Theme.getTheme() || "light")

const switchTheme = () => {
	THEME.value = THEME.value === "light" ? "dark" : "light"
	Theme.applyTheme(THEME.value)
}
</script>

<template>
	<div class="navigation-bar">
		<div class="logo-container">
			<div class="logo"></div>
			<h1>标题</h1>
			<div></div>
			<div class="controls">
				<div><i class="fa-solid fa-globe"></i> 简体中文</div>
				<div @click="switchTheme">
					<i :class="{'fas fa-sun light-icon': THEME === 'light', 'fas fa-moon dark-icon': THEME === 'dark'}"></i>
					<span>{{ THEME === 'light' ? '白昼模式' : '黑夜模式' }}</span>
				</div>
			</div>
		</div>
	</div>
</template>

<style lang="less" scoped>
.navigation-bar {
	padding: 10px 20px;
	box-sizing: border-box;
	height: 68px;
	-webkit-user-select: none;
	-moz-user-select: none;
	-ms-user-select: none;
	user-select: none;
}

.logo-container {
	display: grid;
	grid-template-columns: 48px auto 1fr auto;
	align-items: center;
	gap: 10px;

	.logo {
		width: 48px;
		height: 48px;
		border-radius: 5px;
		background-image: url("https://placehold.co/48");
		background-size: cover;
		background-repeat: no-repeat;
		background-position: center;
	}

	h1 {
		font-size: 24px;
	}

	.controls {
		div {
			cursor: pointer;

			span{
				margin-left: 5px;
			}
		}
	}
}
</style>