<template>
  <div class="customers">
    <ul class="uk-breadcrumb">
        <li><span>Customers</span></li>
        <li><router-link to="/customers/create"><span uk-icon="plus-circle"></span></router-link></li>
    </ul>

    <table class="uk-table">
        <caption></caption>
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="customer in state.customers" v-bind:key="customer.id">
                <td v-html="customer.id"></td>
                <td v-html="customer.name"></td>
                <td>
                    <span class="icon-style" @click="destroy(customer.id)" uk-icon="trash"></span>
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
    name: 'Customers',
    setup() {
        const store = useStore();

        const state = reactive({
            customers: computed(() => store.state.customers.list)
        })

        const fetchData = () => {
            store.dispatch('customers/getAll')
        }

        const destroy = (id = 0) => {
            uikit.modal.confirm('Are you sure?').then(async () => {
                const err = await store.dispatch('customers/destroyOne', id)
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
