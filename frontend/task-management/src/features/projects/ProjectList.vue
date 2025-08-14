<template>
  <div class="bg-white rounded-lg shadow p-6">
    <h2 class="text-2xl font-semibold mb-4">Project List</h2>
    <div v-if="projects.length" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div
        v-for="project in projects"
        :key="project.id"
        class="bg-indigo-50 border border-indigo-200 rounded-lg p-4 shadow-sm hover:shadow-md transition"
      >
        <h3 class="text-lg font-bold">{{ project.name }}</h3>
        <p class="text-sm text-gray-600">
          Status: <span class="capitalize">{{ project.status }}</span>
        </p>
        <p class="text-sm text-gray-600">
          Owner: {{ project.owner.firstName }} {{ project.owner.lastName }}
        </p>
        <p class="text-sm text-gray-600 mt-1">Members: {{ project.members?.length }}</p>
      </div>
    </div>
    <p v-else class="text-gray-500">No projects found.</p>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import api from '../../service/api'
import type { Project } from '../../types/models'

const projects = ref<Project[]>([])

onMounted(async () => {
  const res = await api.get('/projects')
  projects.value = res.data.data || []
})
</script>
