<template>
  <div class="place-edit">
    <ul class="uk-breadcrumb">
        <li><router-link to="/places"><strong>Back to Places</strong></router-link></li>
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
                        <th v-for="item in row" v-bind:key="item">
                            {{item.rowId+1}}/{{item.columnId+1}} <input
                                v-model="item.name"
                                class="uk-input"
                                type="text"
                                :placeholder="item.placeholder"
                            >
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
import { reactive, onMounted } from 'vue'
import { useStore } from "vuex";
import { useRouter } from 'vue-router';

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
            slots: [[]]
        })

        const fetchData = async () => {
            let [res, err] = await store.dispatch('places/getOne', state.id)
            if (err) {
                store.dispatch('error', err)
                return
            }

            state.name = res.name;
            state.rows = res.rows;
            state.columns = res.columns;
        
            let slots = [[]];
            res.slotsList.forEach((el,index) => {
                if (!slots[el.rowId]) {
                    slots[el.rowId] = [];
                }
                el.placeholder = `Place Slot Name (Ground Plot ${index+1}/Greenhouse ${index+1})`
                slots[el.rowId][el.columnId] = el;
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

            router.push({ path: '/places' })
        }

        return {
            save,
            state
        }
    }
}
</script>
