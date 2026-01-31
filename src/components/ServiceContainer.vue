<script setup>
import {useI18n} from "vue-i18n"
import config from "@/../public/config.json"

/**
 * 配置
 */
const service = config?.service || {}

/**
 * 语言
 */
const {t} = useI18n()
</script>

<template>
	<div
		class="service-container"
		v-if="service?.title || service?.description || service?.cards?.length > 0">
		<h2 v-if="service?.title">{{ t(service?.title) }}</h2>
		<p v-if="service?.description">{{ t(service?.description) }}</p>
		<div class="card-container" v-if="service?.cards?.length > 0">
			<div class="card" v-for="(item, index) in service?.cards || []" :key="index">
				<div v-if="item.image" :style="{backgroundImage: `url(${item.image})`}"></div>
				<h3 v-if="item.title">{{ t(item.title) }}</h3>
				<p v-if="item.description">{{ t(item.description) }}</p>
			</div>
		</div>
	</div>
</template>

<style scoped lang="less">
.service-container {
	padding: 8rem;
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 4rem;

	h2 {
		font-size: 2.8rem;
		font-weight: 600;
	}

	p {
		font-size: 2rem;
		font-weight: 400;
		color: var(--text-color-secondary);
	}
}

.card-container {
	padding: 2rem 0;
	width: 100%;
	display: flex;
	justify-content: space-evenly;
	align-items: center;
	flex-wrap: wrap;
	gap: 5rem;

	.card {
		padding: 3.2rem;
		width: 45rem;
		height: 45rem;
		box-shadow: 0 3px 8px var(--box-shadow-color);
		border-radius: 16px;
		display: flex;
		flex-direction: column;
		gap: 2rem;

		div {
			flex: 1;
			border-radius: 0 48px 0 48px;
			background-color: var(--primary-color);
			background-size: cover;
			background-position: center;
			background-repeat: no-repeat;
		}

		h3 {
			font-size: 2rem;
			font-weight: 600;
		}

		p {
			font-size: 2rem;
			color: var(--text-color-secondary);
		}
	}
}
</style>