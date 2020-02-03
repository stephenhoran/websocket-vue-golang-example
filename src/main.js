import Vue from 'vue'
import App from './App.vue'
import vuetify from '@/plugins/vuetify'
import store from './store'
import router from '@/router'
import VueNativeSock from 'vue-native-websocket'

Vue.use(VueNativeSock, 'ws://localhost:8088', {store: store})

Vue.config.productionTip = false

new Vue({
    vuetify,
    store,
    router,
    render: h => h(App)
}).$mount('#app')
