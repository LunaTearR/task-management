// features/task/TaskDetail.vue
<template>
  <div v-if="task">
    <h2 class="text-2xl font-bold mb-2">{{ task.title }}</h2>
    <p class="mb-4">{{ task.description }}</p>
    <p><strong>Status:</strong> {{ task.status }}</p>
    <p><strong>Priority:</strong> {{ task.priority }}</p>
  </div>
  <div v-else>
    <p>Loading...</p>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { getTaskById } from '../services/api'
import type { Task } from '../types/models'

const route = useRoute()
const task = ref<Task | null>(null)

onMounted(async () => {
  const id = Number(route.params.id)
  if (!isNaN(id)) {
    task.value = await getTaskById(id)
  }
})
</script>
