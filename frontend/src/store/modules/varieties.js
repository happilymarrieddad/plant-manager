import varietyClient from '../../../../api/pb/js/varieties_grpc_web_pb';
import varietyObjs from '../../../../api/pb/js/varieties_pb';

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
    getAll({ state, rootState, commit }, { plant_type_id, limit, offset } = {}) {
        state.loading = true;
        const client = new varietyClient.V1VarietiesClient(rootState.apiUrl)
        const request = new varietyObjs.FindVarietiesRequest();

        request.setPlanttypeid(plant_type_id)
        request.setLimit(limit || 25)
        request.setOffset(offset || 0)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.findVarieties(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }
                const list = resp.toObject().varietiesList
                commit('setList', list)
                return resolve(list);
            })
        })
    },
    createOne({ state, rootState }, { name, plant_type_id }) {
        state.loading = true;
        const client = new varietyClient.V1VarietiesClient(rootState.apiUrl)
        const request = new varietyObjs.CreateVarietyRequest();

        request.setName(name)
        request.setPlanttypeid(plant_type_id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.createVariety(request, { 'custom-header-1': 'value1' }, (err) => {
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
        const client = new varietyClient.V1VarietiesClient(rootState.apiUrl)
        const request = new varietyObjs.DestroyVarietyRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.destroyVariety(request, { 'custom-header-1': 'value1' }, (err) => {
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
        const client = new varietyClient.V1VarietiesClient(rootState.apiUrl)
        const request = new varietyObjs.GetVarietyRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve, reject) => {
            client.getVariety(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return reject(err.message);
                }

                return resolve(resp.toObject().variety);
            })
        })
    },
    updateOne({ state, rootState, commit }, { id, name }) {
        state.loading = true;
        const client = new varietyClient.V1VarietiesClient(rootState.apiUrl)

        return new Promise((resolve, reject) => {
            const request = new varietyObjs.UpdateVarietyRequest()

            request.setId(id)
            request.setName(name)
            request.setJwt(window.localStorage.getItem('jwt'))

            client.updateVariety(request, { 'custom-header-1': 'value1' }, (err) => {
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
    name: 'varieties',
    state,
    mutations,
    actions,
    getters
}

