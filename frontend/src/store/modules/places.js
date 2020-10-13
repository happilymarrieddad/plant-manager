import placeClient from '../../../../api/pb/js/places_grpc_web_pb';
import placeObjs from '../../../../api/pb/js/places_pb';

const state = {
    loading: false,
    list: []
};

const mutations = {
    setList(state, list = []) {
        state.list = list;
    }
};
const getters = {
    
};

const actions = {
    getAll({ state, rootState, commit }, { limit, offset } = {}) {
        state.loading = true;
        const client = new placeClient.V1PlacesClient(rootState.apiUrl)
        const request = new placeObjs.FindPlacesRequest();

        request.setLimit(limit || 25)
        request.setOffset(offset || 0)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.findPlaces(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }
                const list = resp.toObject().placesList
                commit('setList', list)
                return resolve(list);
            })
        })
    },
    createOne({ state, rootState }, { name, rows, columns }) {
        state.loading = true;
        const client = new placeClient.V1PlacesClient(rootState.apiUrl)
        const request = new placeObjs.CreatePlaceRequest();

        request.setName(name)
        request.setRows(rows)
        request.setColumns(columns)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.createPlace(request, { 'custom-header-1': 'value1' }, (err) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }

                return resolve(null);
            })
        })
    },
    destroyOne({ state, rootState }, id = 0) {
        state.loading = true;
        const client = new placeClient.V1PlacesClient(rootState.apiUrl)
        const request = new placeObjs.DestroyPlaceRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.destroyPlace(request, { 'custom-header-1': 'value1' }, (err) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }

                return resolve(null);
            })
        })
    },
    getOne({ state, rootState }, id = 0) {
        state.loading = true;
        const client = new placeClient.V1PlacesClient(rootState.apiUrl)
        const request = new placeObjs.GetPlaceRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.getPlace(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }

                return resolve(resp.toObject().place);
            })
        })
    },
    updateOne({ state, rootState, dispatch, commit }, { id, name, slots = [] }) {
        state.loading = true;
        const client = new placeClient.V1PlacesClient(rootState.apiUrl)

        const proms = slots.map(slot => {
            return (slot.id && slot.name) ? dispatch('place_slots/updateOne', slot, { root: true }) : Promise.resolve()
        })

        return Promise.all([
            new Promise((resolve, reject) => {
                const request = new placeObjs.UpdatePlaceRequest()
        
                request.setId(id)
                request.setName(name)
                request.setJwt(window.localStorage.getItem('jwt'))

                client.updatePlace(request, { 'custom-header-1': 'value1' }, (err) => {
                    state.loading = false;
                    if (err) {
                        commit('error', err.message, { root: true })
                        return reject(err.message);
                    }

                    return resolve(null);
                })
            }),
            ...proms
        ])
    }
}

export default {
    namespaced: true,
    name: 'places',
    state,
    mutations,
    actions,
    getters
}

