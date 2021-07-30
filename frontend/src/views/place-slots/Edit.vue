<template>
<div class="place-edit">
    <ul class="uk-breadcrumb">
        <li>
            <router-link to="/places"><strong>Places</strong></router-link>
        </li>
        <li>
            <router-link :to="'/places/edit/'+state.place_id"><strong v-html="state.place_name"></strong></router-link>
        </li>
        <li v-html="state.name"></li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.name" class="uk-input" type="text" placeholder="Name">
        </div>

        <div class="uk-margin">

        </div>

        <div uk-form-custom>
            <button type="button" @click="save">Save</button>
        </div>
    </form>
</div>
</template>

<script>
import {
    reactive,
    onMounted
} from 'vue'
import {
    useStore
} from "vuex";
import {
    useRouter
} from 'vue-router';

export default {
    name: 'PlaceSlotsEdit',
    setup() {
        const store = useStore();
        const router = useRouter();

        const state = reactive({
            id: router.currentRoute.value.params.id,
            name: "",
            row_id: 0,
            column_id: 0,
            place_id: 0,
            customer_id: 0,
            place_name: ''
        })

        const fetchData = async () => {
            const placeSlot = await store.dispatch('place_slots/getOne', state.id)
            const place = await store.dispatch('places/getOne', placeSlot.placeid)

            state.name = placeSlot.name;
            state.row_id = placeSlot.rowid;
            state.column_id = placeSlot.columnid;
            state.place_id = placeSlot.placeid;
            state.customer_id = placeSlot.customerid;
            state.place_name = place.name
        }

        const save = async () => {

        }

        onMounted(() => {
            fetchData();
        });

        return {
            save,
            state
        }
    }
}
</script>
