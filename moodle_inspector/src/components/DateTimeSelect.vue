<template>
    <v-btn @click="dialog = true">
        Select range
    </v-btn>

    <v-dialog v-model="dialog" width="auto">
        <v-card
            max-width="400"
            prepend-icon="mdi-update"
            text="Select the date and time ranges."
            title="Select range"
            >
            <template v-slot:actions>
            <v-btn
                class="ms-auto"
                text="Ok"
                @click="dialog = false"
            ></v-btn>
            </template>
        </v-card>
    </v-dialog>
</template>

<script setup>
    import { ref } from 'vue';

    const dialog = ref(false);
</script>