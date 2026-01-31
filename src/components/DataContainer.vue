<script setup>
import {useI18n} from "vue-i18n"
import config from "@/../public/config.json"

/**
 * 配置
 */
const data = config?.data || []

/**
 * 语言
 */
const {t} = useI18n()
</script>

<template>
	<div class="data-container" v-if="data?.length > 0">
		<div v-for="(item, index) in data || []" :key="index">
			<h2 v-if="item.number">{{item.number}}</h2>
			<p v-if="item.description">{{t(item.description)}}</p>
		</div>
	</div>
</template>

<style scoped lang="less">
.data-container {
	padding: 10rem 0.1rem;
	text-align: center;
	color: var(--text-color-reverse);
	background-image: linear-gradient(135deg, var(--primary-color), var(--data-container-background-color));
	display: flex;
	justify-content: space-evenly;
	align-items: center;
	flex-wrap: wrap;

	div{
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		gap: 2rem;

		h2 {
			font-size: 5.6rem;
			font-weight: bold;
		}

		p {
			font-size: 1.8rem;
			font-weight: 400;
		}
	}
}
</style>