<template>
<div class="place-edit">
    <ul class="uk-breadcrumb">
        <li>
            <router-link to="/places"><strong>Places</strong></router-link>
        </li>
        <li v-html="state.name"></li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.name" class="uk-input" type="text" placeholder="Name">
            <input v-model="state.rows" disabled class="uk-input" type="text" placeholder="Rows in Place (usually a place you put a container)">
            <input v-model="state.columns" disabled class="uk-input" type="text" placeholder="Columns in Place (usually a place you put a container)">
        </div>

        <div class="uk-margin">
            <table class="uk-table">
                <tbody>
                    <tr v-for="row in state.slots" v-bind:key="row">
                        <td v-for="item in row" v-bind:key="item" />
                    </tr>
                </tbody>
                <tbody>
                    <tr v-for="row in state.slots" v-bind:key="row">
                        <th v-for="item in row" v-bind:key="item">
                            <span>{{item.rowid+1}}/{{item.columnid+1}}</span>
                            <span v-if="item.has_plant" uk-icon="image"></span>
                            <router-link :to="'/plant-slots/edit/'+item.id"><span class="icon-style" uk-icon="cog"></span></router-link>
                            <input v-model="item.name" class="uk-input" type="text" :placeholder="item.placeholder">
                        </th>
                    </tr>
                </tbody>
            </table>
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
    name: 'PlaceEdit',
    setup() {
        const store = useStore();
        const router = useRouter();

        const rawData = {
            slots: []
        }

        const state = reactive({
            loading: true,
            id: router.currentRoute.value.params.id,
            name: '',
            rows: null,
            columns: null,
            slots: [
                []
            ]
        })

        const fetchData = async () => {
            let res = await store.dispatch('places/getOne', state.id)

            state.name = res.name;
            state.rows = res.rows;
            state.columns = res.columns;

            let slots = [];
            res.slotsList.forEach((el, index) => {
                if (!slots[el.rowid]) {
                    slots[el.rowid] = [];
                }
                el.placeholder = `Place Slot Name (Ground Plot ${index+1}/Greenhouse ${index+1})`;
                // TODO: Find an easy way to set this
                el.has_plant = false;
                slots[el.rowid][el.columnid] = el;
            })
            state.slots = slots;
            rawData.slots = res.slotsList;
        }

        onMounted(() => {
            fetchData();
        });

        const save = async () => {
            await store.dispatch('places/updateOne', {
                id: state.id,
                name: state.name,
                slots: rawData.slots
            })

            router.push({
                path: '/places'
            })
        }

        return {
            save,
            state
        }
    }
}
</script>
