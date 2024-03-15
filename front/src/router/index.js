import Vue from 'vue'
import VueRouter from 'vue-router'
import userLogin from '@/pages/sign-in/login'
import indexPage from "@/pages/index"
import curUser from '@/utils/cur-user'
import superAdmin from './modules/superAdmin'

import questionManagerRouter from "./modules/questionManager";
Vue.use(VueRouter)

const routes = [
	{
		path: '/',
		name: 'indexPage',
		component: indexPage
	},
	{
		path: '/ldxf-login',
		name: 'login',
		component: userLogin
	},
	{
		path: '/ldxf-cms',
		name: 'cms',
		component: indexPage,
		children: [
			...questionManagerRouter,
			...superAdmin,
		]
	}
]


const router = new VueRouter({
	mode: 'history',
	// base: process.env.BASE_URL,
	routes
})

/**
 * 路由守卫
 * */
router.beforeEach((to, from, next) => {
	/** 判断token是否存在*/
	const token = curUser.getToken()
	/** 页面是否需要登陆才能操作？ */
	if (!token && to.path !== '/ldxf-login') {
		next('/ldxf-login')
	} else {
		next()
	}
})

export default router
