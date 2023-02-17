import Vue from 'vue'
import Router from 'vue-router'
import LoginPage from '@/pages/LoginPage'
import MainPage from '@/pages/MainPage'
import AddUrlPage from '@/pages/AddUrlPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'MainPage',
      component: MainPage
    },
    {
      path: '/login',
      name: 'LoginPage',
      component: LoginPage
    },
    {
      path: '/add_url',
      name: 'AddUrlPage',
      component: AddUrlPage
    }
  ]
})
