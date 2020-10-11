<template>
  <div class="login">
    <input v-model="state.email" placeholder="Email" />
    <input v-model="state.password" placeholder="Password" />
    <button type="button" @click="login">Login</button>
  </div>
</template>

<script>
import { reactive } from 'vue'
import { useStore } from "vuex";
import { useRouter } from 'vue-router';

export default {
    setup() {
        const store = useStore();
        const router = useRouter()

        const state = reactive({
            email: '',
            password: ''
        })

        const login = async () => {
            await store.dispatch('users/login', {
                email: state.email,
                password: state.password
            })
            router.push('/');
        }

        return { 
            login,
            state
        }
    }
}
</script>