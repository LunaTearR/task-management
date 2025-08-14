export interface Task {
  id: number
  title: string
  description: string
  priority: string
  criticalLevel: string
  status: string
}

export interface User {
  id: number
  username: string
  email: string
  firstName: string
  lastName: string
}

export interface Project {
  id: number
  name: string
  description?: string
  status: string
  owner: User
  members?: User[]
  started_at: string
  finished_at: string
  deadline_at: string
}
