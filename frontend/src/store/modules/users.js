import userObjs from '../../../../api/pb/js/users_pb';
import userClient from '../../../../api/pb/js/users_grpc_web_pb';

const state = () => ({
    loading: false,
    permissions: []
});

const mutations = {
    setPermissions(state, permissions = []) {
        state.permissions = permissions
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
                    return resolve([null, err]);
                }
                window.localStorage.setItem('jwt', resp.getToken());
                commit('setPermissions', resp.toObject().permissionsList);
                return resolve([null]);
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

