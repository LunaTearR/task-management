<template>
  <div>
    <div v-for="user in users" :key="user.username" class="user-item">
      <div class="item">
        <i class="icon" :class="iconClass(user.username)"></i>
        <div class="details">
          <h3>{{ user.username }}</h3>
          <p>{{ user.email }}</p>
        </div>
      </div>
    </div>
    <div v-if="users.length === 0">No users found...</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

interface User {
  username: string
  email: string
  first_name: string
  last_name: string
}

const users = ref<User[]>([])

const apiClient = axios.create({
  baseURL: 'http://localhost:8084/api/v1',
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
})

const fetchUsers = async () => {
  try {
    const response = await apiClient.get<{ data: User[] }>('/users')
    console.log('Response data:', response.data)
    users.value = response.data.data
  } catch (err) {
    console.error('Error fetching users:', err)
  }
}

onMounted(fetchUsers) 

const iconClass = (username: string) => `icon-${username}`
</script>

<style scoped>
.user-item {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 8px;
  margin-bottom: 10px;
}
.item {
  display: flex;
  align-items: center;
}
.icon {
  margin-right: 12px;
}
.details {
  padding-left: 12px;
}
</style>
