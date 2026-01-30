<script setup>
import {ref} from "vue"
import config from "@/../public/config.json"
import {useI18n} from "vue-i18n"
import {Swiper, SwiperSlide} from "swiper/vue"
import {Navigation, Pagination, Autoplay} from "swiper/modules"
import "swiper/css"
import "swiper/css/navigation"
import "swiper/css/pagination"

/**
 * 语言
 */
const {t} = useI18n()

/**
 * 显示箭头
 */
const showArrows = ref(false)

/**
 * Swiper 实例
 */
const swiperInstance = ref(null)

/**
 * Swiper 事件处理
 */
const onSwiper = (swiper) => {
	swiperInstance.value = swiper
}

/**
 * 幻灯片序号
 */
const slideIndex = ref(0)

/**
 * 幻灯片切换事件处理
 */
const onSlideChange = (swiper) => {
	slideIndex.value = swiper.activeIndex
}
</script>

<template>
	<div class="swiper-container" @mouseenter="showArrows = true" @mouseleave="showArrows = false">
		<div class="left-content">
			<div class="content-wrapper">
				<h2>{{ t(config.carouselChart.images[slideIndex]?.title || config.carouselChart.default.title) }}</h2>
				<p
					v-for="(description, index) in config.carouselChart.images[slideIndex]?.description || config.carouselChart.default.description"
					:key="index">
					{{ t(description) }}
				</p>
			</div>
		</div>
		<div class="swiper-stage">
			<swiper
				:modules="[Navigation, Pagination, Autoplay]"
				:slides-per-view="1"
				:space-between="0"
				:loop="true"
				:pagination="{clickable: true, dynamicBullets: true, dynamicMainBullets: 3}"
				:navigation="{nextEl: '.custom-next', prevEl: '.custom-prev'}"
				@swiper="onSwiper"
				@slideChange="onSlideChange">
				<SwiperSlide v-for="(slide, index) in config.carouselChart.images" :key="index">
					<div
						class="slide-bg"
						:style="{backgroundImage: `url(${slide.image || config.carouselChart.default.image})`}"
						:aria-label="t(slide.title || config.carouselChart.default.title)"/>
				</SwiperSlide>
			</swiper>
			<div class="custom-navigation">
				<button class="custom-prev" :class="{ 'visible': showArrows }">
					<svg class="arrow-icon" viewBox="0 0 24 24">
						<path d="M15.41 7.41L14 6l-6 6 6 6 1.41-1.41L10.83 12z"/>
					</svg>
				</button>
				<button class="custom-next" :class="{ 'visible': showArrows }">
					<svg class="arrow-icon" viewBox="0 0 24 24">
						<path d="M10 6L8.59 7.41 13.17 12l-4.58 4.59L10 18l6-6z"/>
					</svg>
				</button>
			</div>
		</div>
	</div>
</template>

<style scoped lang="less">
.swiper-container {
	position: relative;
	height: 60rem;
	background-color: var(--box-shadow-color);
	box-shadow: 0 4px 20px var(--box-shadow-color);
	display: flex;
	overflow: hidden;

	&:hover {
		.custom-navigation button {
			opacity: 1;
		}
	}
}

.left-content {
	padding: 4rem 3rem;
	box-sizing: border-box;
	width: 30rem;
	color: var(--text-color);
	background-image: var(--swiper-container-text-background-color);
	display: flex;
	justify-content: center;
	flex-shrink: 0;
	overflow: auto;

	.content-wrapper {
		max-width: 100%;
		word-break: break-word;

		h2 {
			margin: 0 0 2rem 0;
			font-size: 2.8rem;
			font-weight: 600;
			line-height: 1.3;
		}

		p {
			margin: 0 0 1.5rem 0;
			font-size: 1.6rem;
			line-height: 1.6;
			opacity: 0.9;

			&:last-child {
				margin-bottom: 0;
			}
		}
	}
}

.swiper-stage {
	position: relative;
	flex: 1 1 auto;
	min-width: 0;

	:deep(.swiper) {
		height: 100%;
	}

	:deep(.swiper-wrapper) {
		height: 100%;
	}

	:deep(.swiper-slide) {
		height: 100%;
	}

	.slide-bg {
		width: 100%;
		height: 100%;
		background-size: cover;
		background-position: center;
		background-repeat: no-repeat;
	}

	.custom-navigation {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		pointer-events: none;
		z-index: 1;

		button {
			position: absolute;
			top: 50%;
			transform: translateY(-50%);
			width: 4.8rem;
			height: 4.8rem;
			background-color: var(--background-color-reverse);
			border: none;
			border-radius: 50%;
			cursor: pointer;
			display: flex;
			align-items: center;
			justify-content: center;
			transition: all 0.3s ease;
			opacity: 0;
			pointer-events: auto;
			box-shadow: 0 2px 10px var(--box-shadow-color);

			&:hover {
				background-color: var(--primary-color);
				transform: translateY(-50%) scale(1.1);
				box-shadow: 0 4px 15px var(--box-shadow-color);

				.arrow-icon {
					fill: var(--text-color);
				}
			}

			&:active {
				transform: translateY(-50%) scale(0.95);
			}

			&.custom-prev {
				left: 2rem;
			}

			&.custom-next {
				right: 2rem;
			}

			&.visible {
				opacity: 1;
			}

			.arrow-icon {
				width: 2.4rem;
				height: 2.4rem;
				fill: var(--text-color-reverse);
			}
		}
	}

	:deep(.swiper-pagination) {
		bottom: 2rem !important;

		.swiper-pagination-bullet {
			margin: 0 0.4rem;
			width: 0.8rem;
			height: 0.8rem;
			background: var(--background-color-reverse);
			opacity: 1;
			transition: all 0.3s ease;

			&.swiper-pagination-bullet-active {
				width: 2.4rem;
				background: var(--primary-color);
				border-radius: 0.4rem;
			}
		}
	}
}

@media (max-width: 76.8rem) {
	.swiper-container {
		height: auto;
		flex-direction: column;
	}

	.left-content {
		padding: 3rem 2rem;
		width: 100%;

		.content-wrapper {
			h2 {
				font-size: 2.4rem;
			}

			p {
				font-size: 1.6rem;
			}
		}
	}

	.swiper-stage {
		height: 30rem;

		.custom-navigation button {
			width: 4.8rem;
			height: 4.8rem;

			&.custom-prev {
				left: 2rem;
			}

			&.custom-next {
				right: 2rem;
			}

			.arrow-icon {
				width: 2.4rem;
				height: 2.4rem;
			}
		}
	}
}
</style>