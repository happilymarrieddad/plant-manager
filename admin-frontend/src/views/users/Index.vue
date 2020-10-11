<template>
  <div class="users">
    <ul class="uk-breadcrumb">
        <li><span>Users</span></li>
        <!--
            <li><router-link to="/users/create"><span uk-icon="plus-circle"></span></router-link></li>
        -->
    </ul>

    <table class="uk-table">
        <caption></caption>
        <thead>
            <tr>
                <th>ID</th>
                <th>Email</th>
                <th>First Name</th>
                <th>Last Name</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="user in state.users" v-bind:key="user.id">
                <td v-html="user.id"></td>
                <td v-html="user.email"></td>
                <td v-html="user.firstname"></td>
                <td v-html="user.lastname"></td>
                <td>
                    <!--
                        <span class="icon-style" @click="destroy(user.id)" uk-icon="trash"></span>
                    -->
                </td>
            </tr>
        </tbody>
    </table>
  </div>
</template>

<script>
import { reactive, computed, onMounted } from 'vue'
import { useStore } from "vuex";
import uikit from 'uikit';

export default {
    name: 'Users',
    setup() {
        const store = useStore();

        const state = reactive({
            users: computed(() => store.state.users.list)
        })

        const fetchData = () => {
            store.dispatch('users/getAll')
        }

        // TODO
        const destroy = (id = 0) => {
            uikit.modal.confirm('Are you sure?').then(async () => {
                const err = await store.dispatch('users/destroyOne', id)
                if (err) {
                    alert(err)
                    return
                }

                fetchData()
            });
        }

        onMounted(fetchData);

        return {
            destroy,
            state
        }
    }
}
</script>
