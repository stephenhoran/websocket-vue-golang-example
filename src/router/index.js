import Vue from 'vue'
import Router from 'vue-router'

import login from '@/components/login'
import chatroom from "@/components/chatroom"

Vue.use(Router)

const routes = [
    {path: "/", name: "root", component: login},
    {path: "/chatroom", name: "chatroom", component: chatroom}
]

export default new Router({routes})