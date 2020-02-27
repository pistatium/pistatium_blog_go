import Vue from 'vue'
import App from './App.vue'
import Entry from "./pages/Entry"
import Entries from "./pages/Entries"
import Admin from "./pages/Admin"
import vuetify from './plugins/vuetify'
import Portfolio from "./pages/Portfolio"
import VueRouter from 'vue-router'

import Adsense from 'vue-google-adsense/dist/Adsense.min.js'
import InArticleAdsense from 'vue-google-adsense/dist/InArticleAdsense.min.js'


Vue.use(require('vue-script2'))

Vue.use(Adsense)
Vue.use(InArticleAdsense)

Vue.config.productionTip = false

Vue.use(VueRouter)

const routes = [
    {path: '/admin/', component: Admin},
    {path: '/show/:id', component: Entry},
    {path: '/:page', component: Entries},
    {path: '/portfolio', component: Portfolio},
    {path: '/', component: Entries},
]

const router = new VueRouter({
    mode: 'history',
    scrollBehavior(to, from, savedPosition) {
        return {x: 0, y: 0}
    },
    routes
})

router.afterEach((to, from) => {
    gtag('config', 'UA-3817632-9', {'page_path': to.path});
})

new Vue({
    router,
    vuetify,
    render: h => h(App),
    data: {
        loading: false,
        entryHash: {},
        isLogin: false,
    },
}).$mount('#app')
