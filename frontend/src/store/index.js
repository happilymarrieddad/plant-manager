import { createStore } from 'vuex'

import users from './modules/users'

export default createStore({
  state: {
    apiUrl: 'http://localhost:8080'
  },
  mutations: {
  },
  actions: {
  },
  modules: {
      users
  }
})
