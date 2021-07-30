<template>
<div class="plant_type-create">
    <ul class="uk-breadcrumb">
        <li>
            <router-link to="/plant_types"><strong>Plant Types</strong></router-link>
        </li>
    </ul>

    <form>
        <div class="uk-margin">
            <input v-model="state.name" class="uk-input" type="text" placeholder="Name">
        </div>

        <div uk-form-custom>
            <button type="button" @click="create">Create</button>
        </div>
    </form>
</div>
</template>

<script>
import {
    reactive
} from 'vue'
import {
    useStore
} from "vuex";
import {
    useRouter
} from 'vue-router';

export default {
    name: 'PlantTypeCreate',
    setup() {
        const store = useStore();
        const router = useRouter();

        const state = reactive({
            name: ''
        })

        const create = async () => {
            await store.dispatch('plant_types/createOne', {
                name: state.name
            })

            router.push({
                path: '/plant_types'
            })
        }

        return {
            create,
            state
        }
    }
}
</script>
