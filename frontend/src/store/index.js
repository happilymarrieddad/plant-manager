import { createStore } from 'vuex'
import uikit from 'uikit';

import places from './modules/places'
import place_slots from './modules/place-slots'
import plant_types from './modules/plant-types'
import users from './modules/users'
import varieties from './modules/varieties'

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
      place_slots,
      plant_types,
      users,
      varieties
  }
})
