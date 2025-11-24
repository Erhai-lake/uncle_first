<script setup>
import {onMounted, onUnmounted, ref} from "vue"
import {useI18n} from "vue-i18n"
import CryptoJS from "crypto-js"
import ElInputText from "@/components/input/ElInputText.vue"
import ElButton from "@/components/input/ElButton.vue"
import EventBus from "@/services/EventBus.js"
import router from "@/router/index.js"

const {t} = useI18n()

// 登录状态
const isLogin = ref(false)
// 用户名
const username = ref("")
// 密码
const password = ref("")

// 初始化
const init = () => {
	username.value = ""
	password.value = ""
}

// 登录
const login = () => {
	const ENCRYPTED_PASSWORD = CryptoJS.SHA256(password.value).toString(CryptoJS.enc.Hex)
	console.log("尝试登录:", username.value, password.value, ENCRYPTED_PASSWORD)
	EventBus.emit("isLogin", false)
	router.push({name: "adminHome"})
	init()
}

// 取消
const cancel = () => {
	EventBus.emit("isLogin", false)
	init()
}

// 登录状态事件处理函数
const handleLoginEvent = (value) => {
	isLogin.value = value
}

onMounted(() => {
	init()
	EventBus.on("isLogin", handleLoginEvent)
})

onUnmounted(() => {
	init()
	EventBus.off("isLogin", handleLoginEvent)
})
</script>

<template>
	<transition name="login-fade">
		<div class="login" v-if="isLogin" @click="isLogin = false">
			<div class="login-container" @click.stop>
				<div class="form-item">
					<label>{{ t("login.username") }}</label>
					<el-input-text v-model="username" :placeholder="t('login.username-placeholder')"/>
				</div>
				<div class="form-item">
					<label>{{ t("login.password") }}</label>
					<el-input-text v-model="password" :placeholder="t('login.password-placeholder')" :password="true"/>
				</div>
				<el-button @click="login" class="button-login">{{ t("login.login") }}</el-button>
				<el-button @click="cancel">{{ t("login.cancel") }}</el-button>
			</div>
		</div>
	</transition>
</template>

<style lang="less" scoped>
// 添加过渡效果
.login-fade-enter-active,
.login-fade-leave-active {
	transition: opacity 0.3s ease;
}

.login-fade-enter-from,
.login-fade-leave-to {
	opacity: 0;
}

.login-fade-enter-to,
.login-fade-leave-from {
	opacity: 1;
}

.login {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background-color: var(--box-shadow-color);
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
