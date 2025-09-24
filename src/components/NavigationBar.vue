<script setup>
import {ref} from "vue"
import I18n from "@/services/I18n.js"
import Theme from "@/services/Theme.js"

const I18N = ref(I18n.getLanguage() || "zh-CN")
const THEME = ref(Theme.getTheme() || "light")

const switchLanguage = () => {
	I18N.value = I18N.value === "zh-CN" ? "en-US" : "zh-CN"
	I18n.applyLanguage(I18N.value)
}

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
				<div @click="switchLanguage">
					<i class="fa-solid fa-globe"></i>
					<span>{{ I18N === "zh-CN" ? "简体中文" : "English" }}</span>
				</div>
				<div @click="switchTheme">
					<i :class="{'fas fa-sun light-icon': THEME === 'light', 'fas fa-moon dark-icon': THEME === 'dark'}"></i>
					<span>{{ THEME === "light" ? $t("theme.light") : $t("theme.dark") }}</span>
				</div>
			</div>
		</div>
		<div class="hr"></div>
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
	margin-bottom: 10px;
	display: grid;
	grid-template-columns: 48px auto 1fr 100px;
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

			span {
				margin-left: 5px;
			}
		}
	}
}

.hr{
	border-bottom: 1px solid var(--border-color);
}
</style>