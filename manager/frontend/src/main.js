// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueCookies from 'vue-cookies'

import "modern-normalize/modern-normalize.css"
import {
  Row,
  Col,
  Container,
  Header,
  Main,
  Footer,
  Table,
  TableColumn,
} from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import axios from 'axios'
import "nes.css/css/nes.min.css"

Vue.config.productionTip = false
Vue.prototype.axios = axios

Vue.component(Row.name, Row)
Vue.component(Col.name, Col)
Vue.component(Container.name, Container)
Vue.component(Header.name, Header)
Vue.component(Main.name, Main)
Vue.component(Footer.name, Footer)
Vue.component(Table.name, Table)
Vue.component(TableColumn.name, TableColumn)

Vue.use(VueCookies)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
