<script setup>
import {computed, ref} from "vue"
import {useI18n} from "vue-i18n"
import Theme from "@/services/Theme.js"
import Database from "@/services/Database.js"
import config from "@/../public/config.json"

/**
 * 主题
 */
const THEME = ref(Theme.getTheme() || "light")

/**
 * 语言
 */
const {locale, t} = useI18n()

/**
 * Logo 主题
 */
const logoTheme = computed(() => {
	document.querySelector("link[rel='icon']").href = config.logo[THEME.value || "light"]
	return config.logo[THEME.value || "light"]
})

/**
 * 切换语言
 */
const switchLanguage = () => {
	const CURRENT_LOCALE = JSON.parse(JSON.stringify(locale.value))
	const NEW_LOCALE = CURRENT_LOCALE === "zh-CN" ? "en-US" : "zh-CN"
	locale.value = JSON.parse(JSON.stringify(NEW_LOCALE))
	Database.add("language", NEW_LOCALE)
	document.title = t("title")
}

/**
 * 切换主题
 */
const switchTheme = () => {
	THEME.value = THEME.value === "light" ? "dark" : "light"
	Theme.applyTheme(THEME.value)
}

/**
 * 外部链接判断
 */
const IS_EXTERNAL = (path) => {
		return /^https?:\/\//.test(path)
}
</script>

<template>
	<div class="navigation-bar">
		<div class="logo-container">
			<div class="logo" :style="{ backgroundImage: `url(${logoTheme})` }"></div>
			<h1>{{ t("title") }}</h1>
			<div></div>
			<div class="controls">
				<div @click="switchLanguage">
					<i class="fa-solid fa-globe"></i>
					<span>{{ locale === "zh-CN" ? "简体中文" : "English" }}</span>
				</div>
				<div @click="switchTheme">
					<i :class="{'fas fa-sun light-icon': THEME === 'light', 'fas fa-moon dark-icon': THEME === 'dark'}"></i>
					<span>{{ THEME === "light" ? t("theme.light") : t("theme.dark") }}</span>
				</div>
			</div>
		</div>
		<div class="hr"></div>
		<ul class="menu">
			<li v-for="(item, i) in config.menus" :key="i">
				<component
					:is="IS_EXTERNAL(item.path) ? 'a' : 'router-link'"
					:href="IS_EXTERNAL(item.path) ? item.path : null"
					:to="!IS_EXTERNAL(item.path) ? item.path : null"
					v-bind="IS_EXTERNAL(item.path) ? { target: '_' + item.target || 'self', rel: 'noopener noreferrer' } : {}">
					{{ t(item.name) }}
					<i v-if="IS_EXTERNAL(item.path)" class="fas fa-up-right-from-square external-icon"></i>
				</component>
				<ul v-if="item.children">
					<li v-for="(child, j) in item.children" :key="j">
						<component
							:is="IS_EXTERNAL(child.path) ? 'a' : 'router-link'"
							:href="IS_EXTERNAL(child.path) ? child.path : null"
							:to="!IS_EXTERNAL(child.path) ? child.path : null"
							v-bind="IS_EXTERNAL(child.path) ? { target: child.target || '_self', rel: 'noopener noreferrer' } : {}">
							{{ t(child.name) }}
							<i v-if="IS_EXTERNAL(child.path)" class="fas fa-up-right-from-square external-icon"></i>
						</component>
					</li>
				</ul>
			</li>
		</ul>
	</div>
</template>

<style lang="less" scoped>
.navigation-bar {
	padding: 10px 50px;
	box-sizing: border-box;
	-webkit-user-select: none;
	-moz-user-select: none;
	-ms-user-select: none;
	user-select: none;
}

.logo-container {
	margin-bottom: 10px;
	height: 68px;
	display: grid;
	grid-template-columns: 48px auto 1fr 100px;
	align-items: center;
	gap: 10px;

	.logo {
		width: 48px;
		height: 48px;
		border-radius: 5px;
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
			display: flex;
			align-items: center;

			i {
				width: 16px;
			}

			span {
				margin-left: 5px;
			}
		}
	}
}

.hr {
	border-bottom: 1px solid var(--border-color);
}

.menu {
	padding: 0;
	margin: 0;
	display: flex;
	list-style: none;

	> li {
		position: relative;
		margin: 0 20px;

		&:has(> ul)::after {
			content: "▼";
			font-size: 10px;
			position: absolute;
			right: -10px;
			top: 50%;
			transform: translateY(-50%);
			transition: transform 0.3s ease;
		}

		> a{
			padding: 12px 16px;
			display: block;
			font-weight: bold;
			text-decoration: none;
			transition: color 0.3s ease;

			&:hover{
				color: var(--primary-color);
			}
		}

		&:hover {
			> ul {
				opacity: 1;
				transform: translateY(0);
				pointer-events: auto;
			}

			&::after {
				transform: translateY(-50%) rotate(180deg);
			}
		}
	}

	ul {
		position: absolute;
		top: 100%;
		left: 0;
		padding: 8px 0;
		margin: 0;
		min-width: 140px;
		background-color: var(--background-color);
		list-style: none;
		opacity: 0;
		transform: translateY(10px);
		pointer-events: none;
		transition: all 0.3s ease;
		box-shadow: 0 4px 12px var(--box-shadow-color);
		border-radius: 6px;

		li {
			margin: 0;

			> a{
				padding: 10px 16px;
				display: block;
				text-decoration: none;
				transition: background 0.3s ease;

				&:hover{
					background-color: var(--button-hover-background-color);
					color: var(--primary-color);
				}
			}
		}
	}

	.external-icon {
		margin-left: 4px;
		font-size: 0.7em;
	}
}
</style>