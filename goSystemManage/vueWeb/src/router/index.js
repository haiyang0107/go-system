import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/home/home'
import Login from '@/views/login/login'

Vue.use(Router)
const router = new Router({
  // mode: 'history',
  routes: [
    {
      path: '/', // 默认进入路由
      redirect: '/home' // 重定向
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '**', // 错误路由
      redirect: '/home' // 重定向
    }
  ]
})
router.beforeEach((to, from, next) => {
  console.log('路由守卫')
  // to: Route: 即将要进入的目标 路由对象
  // from: Route: 当前导航正要离开的路由
  // next: Function: 一定要调用该方法来 resolve 这个钩子。执行效果依赖 next 方法的调用参数。
  const nextRoute = ['home']
  let isLogin = false // 是否登录
  let isLoginFlag = sessionStorage.getItem('isLoginFlag')
  if (isLoginFlag) {
    isLogin = true
  }
  // 未登录状态；当路由到nextRoute指定页时，跳转至login
  if (nextRoute.indexOf(to.name) >= 0) {
    if (!isLogin) {
      console.log('未登录，请先登录')
      router.push({ name: 'login' })
    }
  }
  // 已登录状态；当路由到login时，跳转至home
  if (to.name === 'login') {
    if (isLogin) {
      router.push({ name: 'home' })
    }
  }
  next()
})
export default router
