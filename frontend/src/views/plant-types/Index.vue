<template>
<div class="plant_types">
    <ul class="uk-breadcrumb">
        <li><span>Places</span></li>
        <li>
            <router-link to="/plant_types/create"><span uk-icon="plus-circle"></span></router-link>
        </li>
    </ul>

    <table class="uk-table">
        <caption>Examples: ("Blueberry", "Tomato")</caption>
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="plantType in state.plant_types" v-bind:key="plantType.id">
                <td v-html="plantType.id"></td>
                <td v-html="plantType.name"></td>
                <td>
                    <router-link :to="'/plant_types/edit/'+plantType.id"><span class="icon-style" uk-icon="settings"></span></router-link>
                    <span class="icon-style" @click="destroy(plantType.id)" uk-icon="trash"></span>
                </td>
            </tr>
        </tbody>
    </table>
</div>
</template>

<script>
import {
    reactive,
    computed,
    onMounted
} from 'vue'
import {
    useStore
} from "vuex";
import uikit from 'uikit';

export default {
    name: 'Places',
    setup() {
        const store = useStore();

        const state = reactive({
            plant_types: computed(() => store.state.plant_types.list)
        })

        const fetchData = () => {
            store.dispatch('plant_types/getAll')
        }

        const destroy = (id = 0) => {
            uikit.modal.confirm('Are you sure?').then(async () => {
                await store.dispatch('plant_types/destroyOne', id)

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
