<template>
  <v-app-bar app class="bg-white shadow-md p-4">
    <v-toolbar-title class="text-2xl font-bold text-indigo-600">Task Management</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn class="hover:underline" text to="/">Home</v-btn>
    <v-btn class="hover:underline" text to="/projects">Projects</v-btn>
    <v-btn class="hover:underline" text to="/tasks">Tasks</v-btn>
    <v-btn class="hover:underline" text to="/users">Users</v-btn>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { defineComponent } from 'vue'

defineComponent({
  name: 'Navbar',
})
</script>
