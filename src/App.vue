<script setup>
import I18n from "@/services/I18n.js"
import Theme from "@/services/Theme.js"
import NavigationBar from "@/components/NavigationBar.vue"

I18n.applyLanguage(I18n.getLanguage())
Theme.applyTheme(Theme.getTheme())
</script>

<template>
	<navigation-bar/>
</template>

<style scoped>
</style>