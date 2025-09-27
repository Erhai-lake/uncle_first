<script setup>
import I18n from "@/services/I18n.js"
import Theme from "@/services/Theme.js"
import NavigationBar from "@/components/NavigationBar.vue"

I18n.applyLanguage(I18n.getLanguage())
Theme.applyTheme(Theme.getTheme())
</script>

<template>
    <navigation-bar/>
    <div class="router">
      <router-view/>
    </div>
</template>

<style scoped>
.router{
  padding: 0 20px;
}
</style>