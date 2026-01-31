<script setup>
import {computed} from "vue"
import Theme from "@/services/Theme.js"
import {useI18n} from "vue-i18n"
import config from "../../public/config.json"

/**
 * 主题
 */
const THEME = Theme.getTheme()

/**
 * 语言
 */
const {t} = useI18n()

/**
 * Logo 主题
 */
const logoTheme = computed(() => {
	return config?.logo?.[THEME.value] || ""
})
</script>

<template>
	<footer class="footer-page">
		<div class="footer-main">
			<div class="footer-left">
				<div class="logo" :style="{ backgroundImage: `url(${logoTheme})` }"/>
				<h1 class="title">{{ t(config?.title || "title") }}</h1>
				<p class="description" v-if="config?.description">{{ t(config?.description) }}</p>
			</div>
			<div class="footer-right" v-if="config?.info.length > 0">
				<div class="contact-item" v-for="(item, i) in config?.info || []" :key="i">
					<span class="label">{{ t(item.label) }}：</span>
					<span>{{ t(item.value) }}</span>
				</div>
			</div>
		</div>
		<div class="footer-copyright" v-if="config?.copyright">
			<span>©</span>
			<span v-if="config?.copyright?.year">{{ config?.copyright?.year }}&nbsp;</span>
			<span v-if="config?.title">{{ t(config.title) }}&nbsp;</span>
			<span v-if="config?.copyright?.rights">{{ t(config?.copyright?.rights) }}&nbsp;</span>
			<span v-if="config?.copyright?.filing">{{ config?.copyright?.filing }}</span>
		</div>
	</footer>
</template>

<style scoped lang="less">
.footer-page {
	background-color: var(--footer-text-emphasis);
	font-size: 1.4rem;
}

.footer-main {
	padding: 4rem 2rem;
	margin: 0 auto;
	max-width: 120rem;
	display: flex;
	justify-content: space-between;
	gap: 6rem;
}

.footer-left {
	max-width: 48rem;
}

.logo {
	margin-bottom: 1.6rem;
	width: 4.8rem;
	height: 4.8rem;
	border-radius: 0.5rem;
	background-size: cover;
	background-repeat: no-repeat;
	background-position: center;
}

.title {
	margin-bottom: 1rem;
	font-size: 2.4rem;
}

.description {
	line-height: 1.6;
	color: var(--text-color-secondary);
}

.footer-right {
	display: flex;
	flex-direction: column;
	gap: 1.2rem;
}

.contact-item {
	line-height: 1.6;
}

.label {
	color: var(--text-color-secondary);
}

.footer-copyright {
	padding: 1.6rem 2rem;
	text-align: center;
	font-size: 1.2rem;
	color: var(--text-color-secondary);
	border-top: 0.1rem solid var(--border-color);
}
</style>