<script setup>
import {ref} from "vue"
import {useI18n} from "vue-i18n"
import ElInputText from "@/components/input/ElInputText.vue"
import ElButton from "@/components/input/ElButton.vue"
import EventBus from "@/services/EventBus.js"

const {t} = useI18n()

// 表单数据
const username = ref("")
const password = ref("")

/**
 * 登录
 */
const login = () => {
	console.log("尝试登录:", username.value, password.value)
}
</script>

<template>
	<div class="login" @click.stop>
		<div class="form-item">
			<label>{{ t("login.username") }}</label>
			<el-input-text v-model="username" :placeholder="t('login.username-placeholder')"/>
		</div>
		<div class="form-item">
			<label>{{ t("login.password") }}</label>
			<el-input-text v-model="password" :placeholder="t('login.password-placeholder')"/>
		</div>
		<el-button @click="login" class="button-login">{{ t("login.login") }}</el-button>
		<el-button @click="EventBus.emit('isLogin', false)">{{ t("login.cancel") }}</el-button>
	</div>
</template>

<style lang="less" scoped>
.login {
	padding: 20px;
	width: 300px;
	border: 1px solid var(--border-color);
	border-radius: 8px;
	box-shadow: 0 2px 8px var(--box-shadow-color);
	background-color: var(--background-color);
	text-align: center;

	.form-item {
		margin-bottom: 15px;
		text-align: left;

		label {
			margin-bottom: 5px;
			display: block;
			font-size: 14px;
		}

		input {
			width: 100%;
			box-sizing: border-box;
		}
	}

	button {
		margin-top: 10px;
		width: 100%;

		&.button-login {
			background-color: var(--primary-color);

			&:hover {
				opacity: 0.8;
			}

			&:active {
				opacity: 0.5;
			}
		}
	}
}
</style>
