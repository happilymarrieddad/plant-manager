import { createStore } from 'vuex'

import customers from './modules/customers'
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
      customers,
      users
  }
})
