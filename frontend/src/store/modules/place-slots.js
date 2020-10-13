import placeSlotClient from '../../../../api/pb/js/place_slots_grpc_web_pb';
import placeSlotObjs from '../../../../api/pb/js/place_slots_pb';

const state = {
    loading: false
};

const mutations = {
    
};
const getters = {

};

const actions = {
    getOne({ state, rootState, commit }, id = 0) {
        return new Promise((resolve, reject) => {
            state.loading = true;
            const client = new placeSlotClient.V1PlaceSlotsClient(rootState.apiUrl)
            const request = new placeSlotObjs.GetPlaceSlotRequest()

            request.setId(id)
            request.setJwt(window.localStorage.getItem('jwt'))

            client.getPlaceSlot(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    commit('error', err.message, { root: true })
                    return reject(err.message);
                }
                return resolve(resp.toObject().placeslot);
            })
        })
    },
    updateOne({ state, rootState, commit }, { id, name }) {
        return new Promise((resolve, reject) => {
            state.loading = true;
            const client = new placeSlotClient.V1PlaceSlotsClient(rootState.apiUrl)
            const request = new placeSlotObjs.UpdatePlaceSlotRequest()

            request.setId(id)
            request.setName(name)
            request.setJwt(window.localStorage.getItem('jwt'))

            client.updatePlaceSlot(request, { 'custom-header-1': 'value1' }, (err) => {
                state.loading = false;
                if (err) {
                    commit('error', err.message, { root: true })
                    return reject(err.message);
                }

                return resolve(null);
            })
        })
    }
}

export default {
    namespaced: true,
    name: 'place_slots',
    state,
    mutations,
    actions,
    getters
}
