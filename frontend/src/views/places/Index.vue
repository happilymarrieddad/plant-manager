<template>
  <div class="places">
    <ul class="uk-breadcrumb">
        <li><span>Places</span></li>
        <li><router-link to="/places/create"><span uk-icon="plus-circle"></span></router-link></li>
    </ul>

    <table class="uk-table">
        <caption>Examples: ("Greenhouse", "Warehouse 47", "Machine 1")</caption>
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Rows/Columns</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="place in state.places" v-bind:key="place.id">
                <td v-html="place.id"></td>
                <td v-html="place.name"></td>
                <td v-html="`${place.rows}/${place.columns}`"></td>
                <td>
                    <router-link :to="'/places/edit/'+place.id"><span class="icon-style" uk-icon="settings"></span></router-link>
                    <span class="icon-style" @click="destroy(place.id)" uk-icon="trash"></span>
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
    name: 'Places',
    setup() {
        const store = useStore();

        const state = reactive({
            places: computed(() => store.state.places.list)
        })

        const fetchData = () => {
            store.dispatch('places/getAll')
        }

        const destroy = (id = 0) => {
            uikit.modal.confirm('Are you sure?').then(async () => {
                const err = await store.dispatch('places/destroyOne', id)
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
