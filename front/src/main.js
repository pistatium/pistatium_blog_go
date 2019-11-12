import Vue from 'vue'
import App from './App.vue'
import Entry from "./pages/Entry";
import Entries from "./pages/Entries";
import vuetify from './plugins/vuetify'
import VueRouter from 'vue-router'

Vue.config.productionTip = false

Vue.use(VueRouter)

const routes = [
  { path: '/show', component: Entry },
  { path: '/', component: Entries },
]

const router = new VueRouter({
  mode: 'history',
  routes
})

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
