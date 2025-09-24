<script setup>
import Test from "@/components/Test.vue"
</script>

<template>
	<div>
		<p>{{ $t("hello") }}</p>
		<p>{{ $t("welcome", {name: "BYD"}) }}</p>
		<Test/>
	</div>
</template>

<style scoped>
</style>