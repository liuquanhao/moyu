import Vue from 'vue'
import Router from 'vue-router'
import SystemPage from '@/pages/SystemPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'SystemPage',
      component: SystemPage
    }
  ]
})
