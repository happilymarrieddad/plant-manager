import customerClient from '../../../../api/pb/js/customers_grpc_web_pb';
import customerObjs from '../../../../api/pb/js/customers_pb';

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
    async getAll({ state, rootState, commit }, { limit, offset } = {}) {
        state.loading = true;
        const client = new customerClient.V1CustomersClient(rootState.apiUrl)
        const request = new customerObjs.FindCustomersRequest();

        request.setLimit(limit || 25)
        request.setOffset(offset || 0)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve) => {
            client.findCustomers(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return resolve([null, err.message]);
                }
                const list = resp.toObject().customersList
                commit('setList', list)
                return resolve([list]);
            })
        })
    },
    async createOne({ state, rootState }, { name }) {
        state.loading = true;
        const client = new customerClient.V1CustomersClient(rootState.apiUrl)
        const request = new customerObjs.CreateCustomerRequest();

        request.setName(name)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve) => {
            client.createCustomer(request, { 'custom-header-1': 'value1' }, (err) => {
                state.loading = false;
                if (err) {
                    return resolve(err.message);
                }

                return resolve(null);
            })
        })
    },
    async destroyOne({ state, rootState }, id = 0) {
        state.loading = true;
        const client = new customerClient.V1CustomersClient(rootState.apiUrl)
        const request = new customerObjs.DestroyCustomerRequest()

        request.setId(id)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve) => {
            client.destroyCustomer(request, { 'custom-header-1': 'value1' }, (err) => {
                state.loading = false;
                if (err) {
                    return resolve(err.message);
                }

                return resolve();
            })
        })
    }
}

export default {
    namespaced: true,
    name: 'customers',
    state,
    mutations,
    actions,
    getters
}

