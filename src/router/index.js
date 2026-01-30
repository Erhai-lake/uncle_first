import {createRouter, createWebHashHistory} from "vue-router"

const routes = [
	{
		path: "/",
		name: "home",
		component: () => import("@/views/Home.vue"),
		meta: {
			requiresAuth: false
		}
	}
]

// 路由实例
const router = createRouter({
	history: createWebHashHistory(),
	routes
})

export default router
