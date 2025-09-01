<template>
    <DateTimeSelect></DateTimeSelect>
</template>

<script setup>
    import { onMounted } from 'vue';
    import axios from 'axios';
    import DateTimeSelect from './DateTimeSelect.vue';

    onMounted(() => {
        const endpoint = import.meta.env.VITE_API_ENDPOINT;

        window.console.log(endpoint);

        // Get list of IPs
        axios.get(endpoint + '/getips')
            .then((response) => {
                window.console.log(response)
            })
            .catch(error => {
                window.console.error(error);
            });

    })
</script>