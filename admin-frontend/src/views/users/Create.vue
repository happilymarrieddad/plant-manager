<template>
  <div class="user-create">
    <ul class="uk-breadcrumb">
        <li><router-link to="/users"><strong>Back to Users</strong></router-link></li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.email" class="uk-input" type="text" placeholder="Email">
            <input v-model="state.firstName" class="uk-input" type="text" placeholder="First Name">
            <input v-model="state.lastName" class="uk-input" type="text" placeholder="Last Name">
        </div>

        <div uk-form-custom>
            <button type="button" @click="create">Create</button>
        </div>
    </form>
  </div>
</template>

<script>
import { reactive } from 'vue'
import { useStore } from "vuex";
import { useRouter } from 'vue-router';

export default {    
    name: 'UserCreate',
    setup() {
        const store = useStore();
        const router = useRouter();

        const state = reactive({
            email: '',
            firstName: '',
            lastName: ''
        })

        const create = async () => {
            let err = await store.dispatch('users/createOne', {
                email: state.email,
                firstName: state.firstName,
                lastName: state.lastName
            })
            if (err) {
                alert(err)
                return
            }

            router.push({ path: '/users' })
        }

        return {
            create,
            state
        }
    }
}
</script>
