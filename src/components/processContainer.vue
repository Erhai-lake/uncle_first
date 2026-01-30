<script setup>
import {useI18n} from "vue-i18n"
import config from "@/../public/config.json"

/**
 * 语言
 */
const {t} = useI18n()
</script>

<template>
	<div class="process-container">
		<h2>{{t(config?.process?.title || "title")}}</h2>
		<p>{{t(config?.process?.description || "description")}}</p>
		<div class="card-container">
			<div class="card" v-for="(item, index) in config?.process?.cards || []" :key="index">
				<div :style="{backgroundImage: `url(${item.image || ''})`}"></div>
				<h3>{{t(item.title || 'title')}}</h3>
				<p>{{t(item.description || 'description')}}</p>
			</div>
		</div>
	</div>
</template>

<style scoped lang="less">
.process-container {
	padding: 8rem;
	text-align: center;
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

.card-container{
	padding: 2rem 0;
	display: flex;
	justify-content: center;
	align-items: center;
	flex-wrap: wrap;
	gap: 5rem;

	.card{
		padding: 3.2rem;
		width: 25rem;
		box-shadow: 0 3px 8px var(--box-shadow-color);
		border-radius: 16px;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 2rem;

		div{
			width: 7rem;
			height: 7rem;
			border-radius: 50%;
			background-color: var(--primary-color);
			background-size: cover;
			background-position: center;
			background-repeat: no-repeat;
		}

		h3{
			font-size: 2rem;
			font-weight: 600;
		}

		p{
			font-size: 2rem;
			color: var(--text-color-secondary);
		}
	}
}
</style>