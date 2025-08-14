<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title>Task Management</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text to="/">Home</v-btn>
      <v-btn text to="/projects">Projects</v-btn>
      <v-btn text to="/tasks">Tasks</v-btn>
      <v-btn text to="/users">Users</v-btn>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>
<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',

  data() {
    return {
      //
    }
  },
})
</script>
