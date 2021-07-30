<template>
  <div class="place-create">
    <ul class="uk-breadcrumb">
        <li><router-link to="/places"><strong>Back to Places</strong></router-link></li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.name" class="uk-input" type="text" placeholder="Name">
            <input v-model="state.rows" class="uk-input" type="text" placeholder="Rows in Place (usually a place you put a container)">
            <input v-model="state.columns" class="uk-input" type="text" placeholder="Columns in Place (usually a place you put a container)">
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
    name: 'PlaceCreate',
    setup() {
        const store = useStore();
        const router = useRouter();

        const state = reactive({
            name: '',
            rows: null,
            columns: null
        })

        const create = async () => {
            let err = await store.dispatch('places/createOne', {
                name: state.name,
                rows: state.rows,
                columns: state.columns
            })
            if (err) {
                console.log(err)
                alert(err)
                return
            }

            router.push({ path: '/places' })
        }

        return {
            create,
            state
        }
    }
}
</script>
