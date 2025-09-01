<template>
    <v-sheet
          border="dashed md"
          color="surface-light"
          height="200"
          rounded="lg"
          width="100%"
        >
          <DashboardLoads></DashboardLoads>
        </v-sheet>
    <h1>Dashboard</h1>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

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

    // Get latest loads
    axios.get(endpoint + '/getlatestloads')
    .then((response) => {
        window.console.log(response)
    })
    .catch(error => {
        window.console.error(error);
    });
});
</script>