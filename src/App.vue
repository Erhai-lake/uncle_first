<script setup>
import Theme from "@/services/Theme.js"
import Test from "@/components/Test.vue"
import NavigationBar from "@/components/NavigationBar.vue"

Theme.applyTheme(Theme.getTheme())
</script>

<template>
	<navigation-bar/>
	<div style="display: none;">
		<p>{{ $t("hello") }}</p>
		<p>{{ $t("welcome", {name: "BYD"}) }}</p>
		<Test/>
	</div>
</template>

<style scoped>
</style>