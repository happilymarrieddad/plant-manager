import plantTypeClient from '../../../../api/pb/js/plant_types_grpc_web_pb';
import plantTypeObjs from '../../../../api/pb/js/plant_types_pb';

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
        const client = new plantTypeClient.V1PlantTypesClient(rootState.apiUrl)
        const request = new plantTypeObjs.FindPlantTypesRequest();

        request.setLimit(limit || 25)
        request.setOffset(offset || 0)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.findPlantTypes(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }
                const list = resp.toObject().planttypesList
                commit('setList', list)
                return resolve(list);
            })
        })
    },
    createOne({ state, rootState }, { name }) {
        state.loading = true;
        const client = new plantTypeClient.V1PlantTypesClient(rootState.apiUrl)
        const request = new plantTypeObjs.CreatePlantTypeRequest();

        request.setName(name)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.createPlantType(request, { 'custom-header-1': 'value1' }, (err) => {
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
        const client = new plantTypeClient.V1PlantTypesClient(rootState.apiUrl)
        const request = new plantTypeObjs.DestroyPlantTypeRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.destroyPlantType(request, { 'custom-header-1': 'value1' }, (err) => {
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
        const client = new plantTypeClient.V1PlantTypesClient(rootState.apiUrl)
        const request = new plantTypeObjs.GetPlantTypeRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.getPlantType(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }

                return resolve(resp.toObject().planttype);
            })
        })
    },
    updateOne({ state, rootState, commit }, { id, name }) {
        state.loading = true;
        const client = new plantTypeClient.V1PlantTypesClient(rootState.apiUrl)
        const request = new plantTypeObjs.UpdatePlantTypeRequest()

        request.setId(id)
        request.setName(name)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.updatePlantType(request, { 'custom-header-1': 'value1' }, (err) => {
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
    name: 'plant_types',
    state,
    mutations,
    actions,
    getters
}
