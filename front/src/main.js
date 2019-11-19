import Vue from 'vue'
import App from './App.vue'
import Entry from "./pages/Entry";
import Entries from "./pages/Entries";
import vuetify from './plugins/vuetify'
import VueRouter from 'vue-router'

Vue.config.productionTip = false

Vue.use(VueRouter)

const routes = [
    {path: '/show/:id', component: Entry},
    {path: '/:page', component: Entries},
    {path: '/', component: Entries},
]

const router = new VueRouter({
    mode: 'history',
    scrollBehavior(to, from, savedPosition) {
        return {x: 0, y: 0}
    },
    routes
})

new Vue({
    router,
    vuetify,
    render: h => h(App),
    data: {
        loading: false,
        entryHash: {},
    },
}).$mount('#app')
