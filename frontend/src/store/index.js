import { createStore } from 'vuex'
import uikit from 'uikit';

import places from './modules/places'
import users from './modules/users'

export default createStore({
  state: {
    apiUrl: 'http://localhost:8080',
    errors: []
  },
  mutations: {
  },
  actions: {
    error({ state }, msg) {
      state.errors.push(msg);
      uikit.notification({
          message: msg,
          status: 'danger',
          pos: 'top-center',
          timeout: 5000
      });
    }
  },
  modules: {
      places,
      users
  }
})
