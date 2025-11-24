import {createRouter, createWebHashHistory} from "vue-router"
import Database from "@/services/Database.js"
import EventBus from "@/services/EventBus.js"

const routes = [
	{
		path: "/",
		name: "home",
		component: () => import("@/views/Home.vue"),
		meta: {
			requiresAuth: false
		}
	},
	{
		path: "/admin",
		name: "adminHome",
		component: () => import("@/views/admin/Home.vue"),
		meta: {
			requiresAuth: true
		}
	}
]

// 路由实例
const router = createRouter({
	history: createWebHashHistory(),
	routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
	const IS_LOGGED_IN = !!Database.get("token")
	if (to.meta.requiresAuth && !IS_LOGGED_IN) {
		// 未登录
		EventBus.emit("isLogin", true)
	} else if (IS_LOGGED_IN) {
		// 已登录
		next({name: "adminHome"})
	} else {
		// 其他情况放行
		next()
	}
})

export default router
