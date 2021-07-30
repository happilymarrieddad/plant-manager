<template>
  <div class="customer-create">
    <ul class="uk-breadcrumb">
        <li><router-link to="/customers"><strong>Back to Customers</strong></router-link></li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.name" class="uk-input" type="text" placeholder="Customer Name">
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
    name: 'CustomerCreate',
    setup() {
        const store = useStore();
        const router = useRouter();

        const state = reactive({
            name: ''
        })

        const create = async () => {
            let err = await store.dispatch('customers/createOne', {
                name: state.name
            })
            if (err) {
                alert(err)
                return
            }

            router.push({ path: '/customers' })
        }

        return {
            create,
            state
        }
    }
}
</script>
