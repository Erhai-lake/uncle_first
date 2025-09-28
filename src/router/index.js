import {createRouter, createWebHashHistory} from "vue-router"

const routes = [
	{
		path: "/",
		name: "home",
		component: () => import("@/views/home.vue"),
		meta: {
			requiresAuth: false
		}
	},
	// {
	// 	path: "/login",
	// 	name: "login",
	// 	component: () => import("@/views/login.vue"),
	// 	meta: {
	// 		requiresAuth: false
	// 	}
	// },
	// {
	// 	path: "/admin",
	// 	name: "admin",
	// 	component: () => import("@/views/admin.vue"),
	// 	meta: {
	// 		requiresAuth: true
	// 	}
	// }
]

// 路由实例
const router = createRouter({
	history: createWebHashHistory(),
	routes
})

router.beforeEach((to, from, next) => {
	const IS_LOGGED_IN = !!localStorage.getItem("token")
	if (to.meta.requiresAuth && !IS_LOGGED_IN) {
		// 没有登录 → 跳去登录页
		next({name: "login"})
	} else if (to.name === "login" && IS_LOGGED_IN) {
		// 已登录还想去登录页 → 重定向到后台
		next({name: "admin"})
	} else {
		// 其他情况放行
		next()
	}
})

export default router
