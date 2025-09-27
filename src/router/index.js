import {createRouter, createWebHashHistory} from "vue-router"

const routes = [
    {
        path: "/",
        name: "home",
        component: () => import("@/views/home.vue")
    }
]

// 路由实例
const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
