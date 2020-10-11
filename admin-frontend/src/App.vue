<template>
<div id="app">
    <nav id="nav" v-show="state.showNav" class="uk-navbar-container" uk-navbar>
        <div class="uk-navbar-left">
            <ul class="uk-navbar-nav">
                <li><router-link to="/">Home</router-link></li>
                <li><router-link to="/about">About</router-link></li>
                <li><router-link v-if="state.showCustomersRoute" to="/customers">Customers</router-link></li>
                <li><router-link v-if="state.showUsersRoute" to="/users">Users</router-link></li>
            </ul>
        </div>
    </nav>
    <div class="uk-container">
        <router-view/>
    </div>
</div>
</template>

<script>
import { reactive, computed } from 'vue'
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';

export default {
    setup() {
        const router = useRouter();
        const store = useStore();

        const state = reactive({
            perms: computed(() => store.getters['users/getPermissions']),
            showNav: computed(() => router.currentRoute.value.name !== 'Login'),
            // Routes
            showCustomersRoute: computed(() => Boolean(state.perms.find(el => el.name === "customers.read"))),
            showUsersRoute: computed(() => Boolean(state.perms.find(el => el.name === "users.read")))
        })

        return {
            state
        }
    }
}
</script>

<style>
    .icon-style:hover {
        cursor: pointer;
    }
</style>
