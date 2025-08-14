import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8084/api/v1',
  timeout: 5000,
})

export default api
import type { Task, Project, User } from '../types/models'

const API_URL = '/tasks'

export async function getTasks(): Promise<Task[]> {
  const response = await api.get(API_URL)
  return response.data
}

export async function getTaskById(id: number): Promise<Task> {
  const response = await api.get(`${API_URL}/${id}`)
  return response.data
}

export async function getProjects(): Promise<Project[]> {
  const response = await api.get('/projects')
  return response.data
}

export async function getUsers(): Promise<User[]> {
  const response = await api.get('/users')
  return response.data
}

export async function getUserById(ids: [number]): Promise<User[]> {
  const response = await api.post(`/users/by-ids`, { ids })
  return response.data
}