import userObjs from '../../../../api/pb/js/users_pb';
import userClient from '../../../../api/pb/js/users_grpc_web_pb';

const state = () => ({
    loading: false,
    list: [],
    permissions: []
});

const mutations = {
    setPermissions(state, permissions = []) {
        state.permissions = permissions
    },
    setList(state, list = []) {
        state.list = list;
    }
};
const getters = {
    getPermissions(state) {
        return state.permissions;
    }
};

const actions = {
    async login({ state, rootState, commit }, { email, password }) {
        state.loading = true;
        const client = new userClient.V1UsersClient(rootState.apiUrl)
        const request = new userObjs.LoginRequest();

        request.setEmail(email)
        request.setPassword(password)

        return new Promise((resolve) => {
            client.login(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return resolve([err.message]);
                }
                const obj = resp.toObject()
                window.localStorage.setItem('jwt', obj.token);
                commit('setPermissions', obj.permissionsList);
                return resolve([null]);
            })
        })
    },
    async getAll({ state, rootState, commit }, { limit, offset } = {}) {
        state.loading = true;
        const client = new userClient.V1UsersClient(rootState.apiUrl)
        const request = new userObjs.FindUsersRequest();

        request.setLimit(limit || 25)
        request.setOffset(offset || 0)
        request.setJwt(window.localStorage.getItem('jwt'))

        return new Promise((resolve) => {
            client.findUsers(request, { 'custom-header-1': 'value1' }, (err, resp) => {
                state.loading = false;
                if (err) {
                    return resolve([null, err.message]);
                }
                const list = resp.toObject().usersList
                commit('setList', list)
                return resolve([list]);
            })
        })
    }
}

export default {
    namespaced: true,
    name: 'users',
    state, 
    mutations,
    actions,
    getters
}

