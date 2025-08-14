<template>
  <div>
    <h1 class="text-xl font-bold mb-4">Users</h1>
    <ul>
      <li v-for="user in users" :key="user.id">
        {{ user.username }} ({{ user.email }})
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import api from '@/services/api'
import type { User } from '@/types/models'

const users = ref<User[]>([])

onMounted(async () => {
  const res = await api.get('/users')
  users.value = res.data
})
</script>
