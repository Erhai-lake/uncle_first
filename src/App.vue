<script setup>
import Test from "@/components/Test.vue"
import NavigationBar from "@/components/NavigationBar.vue"
</script>

<template>
	<navigation-bar/>
	<div>
		<p>{{ $t("hello") }}</p>
		<p>{{ $t("welcome", {name: "BYD"}) }}</p>
		<Test/>
	</div>
</template>

<style scoped>
</style>