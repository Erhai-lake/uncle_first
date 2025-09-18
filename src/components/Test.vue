<template>
	<div class="Test">
		<h2> Name: {{ NAME }} </h2>
		<h2> Age: {{ AGE }}</h2>
		<h2> Message: {{ Utils.show.showMessage(TEL) }}</h2>
		<button @click='Utils.show.showTel("你看你妈呢")'>Let me Look Look😋</button>
		<button @click='Utils.change.name("黑鬼")'>修改名称</button>
		<button @click=Utils.change.age(123456)>修改年龄</button>
	</div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const NAME = ref("Nigger")
const AGE = ref(114514)
const TEL = ref(1919810)

let Utils = {
	show: {
		showTel: function (message: string | number) {
			return alert(message)
		},
		showMessage: function (message: string | number) {
			return message
		}
	},
	change: {
		name: function (name: string) {
			NAME.value = name
		},
		age: function (age: number) {
			AGE.value = age
		}
	}
}
</script>

<style>
.Test {
	background-color: skyblue;
	box-shadow: 0 0 10px;
	border-radius: 10px;
	padding: 20px;
}

button {
	margin: 0 5px;
}
</style>