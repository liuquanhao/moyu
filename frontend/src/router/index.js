import Vue from 'vue'
import Router from 'vue-router'
import HardwareInfo from '@/components/HardwareInfo'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HardwareInfo',
      component: HardwareInfo
    }
  ]
})
