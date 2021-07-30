<template>
<div class="plant_type-edit">
    <ul class="uk-breadcrumb">
        <li>
            <router-link to="/plant_types"><strong>Plant Types</strong></router-link>
        </li>
        <li v-html="state.name"></li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.name" class="uk-input" type="text" placeholder="Name">
        </div>

        <div class="uk-margin">
            <table class="uk-table">
                <tbody>
                    <tr>
                        <td>ID</td>
                        <td>Name</td>
                        <td></td>
                    </tr>
                </tbody>
                <tbody>
                    <tr v-for="variety in state.varieties" v-bind:key="variety.id">
                        <th v-html="variety.id"></th>
                        <th v-html="variety.name"></th>
                        <th><span class="icon-style" uk-icon="trash" @click="removeVariety(variety.id)"></span></th>
                    </tr>
                </tbody>
                <tfoot>
                    <tr>
                        <td></td>
                        <td><input v-model="state.addVarietyName" class="uk-input" type="text" placeholder="New Variety Name"></td>
                        <td><span class="icon-style" uk-icon="plus" @click="addVariety"></span></td>
                    </tr>
                </tfoot>
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
import uikit from 'uikit';

export default {
    name: 'PlaceEdit',
    setup() {
        const store = useStore();
        const router = useRouter();

        const state = reactive({
            loading: true,
            id: router.currentRoute.value.params.id,
            name: '',
            varieties: [],
            addVarietyName: ''
        })

        const fetchData = async () => {
            let res = await store.dispatch('plant_types/getOne', state.id)

            state.name = res.name;
            console.log(res)
            state.varieties = res.varietiesList;
        }

        onMounted(() => {
            fetchData();
        });

        const save = async () => {
            await store.dispatch('plant_types/updateOne', {
                id: state.id,
                name: state.name
            })

            router.push({
                path: '/plant_types'
            })
        }

        const addVariety = async () => {
            await store.dispatch('varieties/createOne', {
                name: state.addVarietyName,
                plant_type_id: state.id
            })

            state.addVarietyName = ''
            fetchData()
        }

        const removeVariety = async (id) => {
            uikit.modal.confirm('Are you sure?').then(async () => {
                await store.dispatch('varieties/destroyOne', id)

                fetchData()
            });
        }

        return {
            save,
            state,
            addVariety,
            removeVariety
        }
    }
}
</script>
