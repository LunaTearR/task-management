<template>
  <div>
    <h1 class="text-xl font-bold mb-4">Tasks</h1>
    <ul>
      <li v-for="task in tasks" :key="task.id">
        {{ task.title }} ({{ task.priority }})
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import api from '../services/api'
import type { Task } from '../types/models'

const tasks = ref<Task[]>([])

onMounted(async () => {
  const res = await api.get('/tasks')
  tasks.value = res.data
})
</script>
