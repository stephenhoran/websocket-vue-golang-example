import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        connected: false,
        name: undefined,
        messages: []
    },
    mutations: {
        NAME(state, payload) {
            state.name = payload
        },
        SOCKET_ONOPEN(state) {
            state.connected = true
            window.console.log("Connected")
        },
        SOCKET_ONMESSAGE(state, event) {
            state.messages.push(event)
        }
    },
    actions: {
        setName({commit}, payload) {
            commit('NAME', payload)
        },
    },
    getters: {
        nickName: state => {
            return state.name
        },
        chatroomMessages: state => {
            let stringer = ""

            state.messages.forEach(event => {
                stringer += event.data + '\n'
            })

            return stringer
        }
    },
    modules: {}
})
