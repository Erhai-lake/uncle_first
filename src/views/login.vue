<script setup>
import ElInputText from "@/components/input/ElInputText.vue"
import ElButton from "@/components/input/ElButton.vue"
import {ref} from "vue"

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
	<div class="login">
		<div class="login-container">
			<div class="form-item">
				<label>{{ $t("login.username") }}</label>
				<el-input-text v-model="username" :placeholder="$t('login.username-placeholder')"/>
			</div>
			<div class="form-item">
				<label>{{ $t("login.password") }}</label>
				<el-input-text v-model="password" :placeholder="$t('login.password-placeholder')"/>
			</div>
			<el-button @click="login">{{ $t("login.login") }}</el-button>
		</div>
	</div>
</template>

<style lang="less" scoped>
.login {
	width: 100%;
	height: 100%;
	display: flex;
	justify-content: center;
	align-items: center;
}

.login-container {
	padding: 20px;
	width: 300px;
	border: 1px solid var(--border-color);
	border-radius: 8px;
	box-shadow: 0 2px 8px var(--box-shadow-color);
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
		width: 100%;
	}
}
</style>
