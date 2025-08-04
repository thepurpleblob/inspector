<template>
  <v-layout class="rounded rounded-md border">
    <v-app-bar color="university-blue">
      <v-img :src="UofG" max-height="50px"></v-img>
      <v-toolbar-title class="ml-2">
          Moodle inspector
      </v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer 
      @update:model-value="drawer_update"
      v-model="drawerstate"

      >
      <v-list-item 
        title="Moodle inspector"
        subtitle="Menu"
        ></v-list-item>
      <v-divider></v-divider>
      <NavList></NavList>
    </v-navigation-drawer>

    <v-main class="d-flex align-center justify-center" height="300">
      <v-container>
        <span class="mdi mdi-turbine"></span>
        <v-sheet
          border="dashed md"
          color="surface-light"
          height="200"
          rounded="lg"
          width="100%"
        ></v-sheet>
      </v-container>
    </v-main>
  </v-layout>
</template>

<script setup>
import { ref } from 'vue';
import UofG from './assets/uofg.png'
import NavList from '@/components/NavList.vue';

const drawerstate = ref(null)

function drawer_update() {
    window.console.log('drawer updated');
    window.console.log(drawerstate.value);
}

</script>

<style>
  body {
    font-family: "PT Sans";
  }
</style>
