<template>
  <div class="dashboard-container">
    <!-- Header -->
    <div class="dashboard-header">
      <div class="header-content">
        <div class="breadcrumb">
          <svg class="breadcrumb-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 19l-7-7 7-7"
            />
          </svg>
          <span class="breadcrumb-text">Dashboard</span>
        </div>
        <div class="header-actions">
          <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
            />
          </svg>
          <svg class="action-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 17h5l-5 5-5-5h5V3h5v14z"
            />
          </svg>
          <div class="profile-avatar">U</div>
        </div>
      </div>
    </div>

    <!-- Filter Section -->
    <div class="filter-section">
      <div class="filter-dropdown-wrapper">
        <div class="dropdown-container" :class="{ 'dropdown-open': isDropdownOpen }">
          <button class="project-filter" @click="toggleDropdown" type="button">
            <span v-if="selectedProjects.length === 0" class="filter-placeholder">
              โปรเจคทั้งหมด
            </span>
            <span v-else-if="selectedProjects.length === 1" class="filter-selected">
              {{ getProjectName(selectedProjects[0]) }}
            </span>
            <span v-else class="filter-selected">
              {{ selectedProjects.length }} โปรเจคที่เลือก
            </span>
            <svg class="dropdown-arrow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M19 9l-7 7-7-7"
              />
            </svg>
          </button>

          <div v-if="isDropdownOpen" class="dropdown-menu" @click.stop>
            <div class="dropdown-header">
              <input
                v-model="searchQuery"
                type="text"
                class="search-input"
                placeholder="ค้นหาโปรเจค..."
                @input="onSearch"
              />
            </div>

            <div class="dropdown-actions">
              <button @click="selectAllProjects" class="action-btn">เลือกทั้งหมด</button>
              <button @click="clearAllProjects" class="action-btn">ยกเลิกทั้งหมด</button>
            </div>

            <div class="dropdown-options">
              <label v-for="project in filteredProjects" :key="project.id" class="option-item">
                <input
                  type="checkbox"
                  :value="project.id"
                  :checked="selectedProjects.includes(project.id)"
                  @change="toggleProject(project.id)"
                  class="option-checkbox"
                />
                <div class="option-content">
                  <span class="option-name">{{ project.name }}</span>
                  <span class="option-status" :class="getStatusClass(project.status)">
                    {{ project.status }}
                  </span>
                </div>
              </label>
            </div>

            <div v-if="filteredProjects.length === 0" class="no-results">ไม่พบโปรเจคที่ค้นหา</div>
          </div>
        </div>
      </div>
      <div class="last-updated">วันที่อัปเดตล่าสุด: {{ formatDate(new Date().toISOString()) }}</div>
    </div>

    <!-- Main Content -->
    <div class="main-content">
      <!-- Chart Section -->
      <div class="chart-section">
        <div class="section-header">
          <h2 class="section-title">Project Timeline</h2>
          <button class="view-all-btn">ดูเพิ่มเติม</button>
        </div>

        <div class="chart-wrapper">
          <div class="chart-content">
            <div class="chart-scroll-container">
              <div class="chart-bars" :style="{ height: barsHeight + 'px' }">
                <div
                  v-for="(project, index) in projects"
                  :key="index"
                  class="bar-row"
                  :style="{ height: barHeight + 'px' }"
                >
                  <div class="project-info">
                    <div class="project-name">{{ project.name || 'Unnamed Project' }}</div>
                  </div>
                  <div class="bar-container">
                    <div
                      class="bar"
                      :style="{
                        width: getBarWidth(project) + '%',
                      }"
                      :title="`${project.name}\nStart: ${formatDate(project.started_at)}\nEnd: ${formatDate(project.finished_at)}\nDuration: ${calculateDuration(project)} days`"
                    >
                      <div class="bar-label">{{ calculateDuration(project) }} วัน</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="chart-x-axis">
            <div class="x-axis-labels">
              <span>0</span>
              <span>{{ Math.round(maxDuration / 4) }}</span>
              <span>{{ Math.round(maxDuration / 2) }}</span>
              <span>{{ Math.round((maxDuration * 3) / 4) }}</span>
              <span>{{ Math.round(maxDuration) }}</span>
            </div>
            <div class="x-axis-title">Duration (days)</div>
          </div>
        </div>
      </div>

      <!-- Projects List Section -->
      <div class="projects-section">
        <div class="section-header">
          <h2 class="section-title">ความคืบหน้าโปรเจกต์</h2>
          <button class="view-all-btn">ดูเพิ่มเติม</button>
        </div>

        <div class="projects-list">
          <div v-for="(project, index) in displayedProjects" :key="project.id" class="project-card">
            <div class="project-header">
              <h3 class="project-title">{{ project.name }}</h3>
              <div class="project-stats">
                <div class="member-count">
                  <svg class="stat-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                    />
                  </svg>
                  {{ (project.members?.length || 0) + 1 }}
                </div>
                <span class="member-label">คน</span>
              </div>
            </div>

            <div class="progress-section">
              <div class="progress-header">
                <span class="progress-icon">📊</span>
                <span class="progress-text">Progress</span>
                <span class="progress-percentage">{{ getRandomProgress() }}%</span>
              </div>
              <div class="progress-bar">
                <div class="progress-fill" :style="{ width: getRandomProgress() + '%' }"></div>
              </div>
            </div>

            <div class="project-dates">
              <div class="date-info">
                <span class="date-label">Started:</span>
                <span class="date-value">{{ formatDate(project.started_at) }}</span>
              </div>
              <div class="date-info">
                <span class="date-label">Deadline:</span>
                <span class="date-value">{{ formatDate(project.finished_at) }}</span>
              </div>
            </div>

            <div class="team-avatars">
              <div class="avatar" v-for="(member, idx) in getTeamMembers(project)" :key="idx">
                {{ member }}
              </div>
              <div class="task-summary">
                <svg class="task-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                  />
                </svg>
                <span>{{ getRandomTasks() }} Task, {{ getRandomSubtasks() }} Subtask</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, computed } from 'vue'
import api from '../service/api'

const projects = ref<any[]>([])
const selectedProjects = ref<number[]>([])
const isDropdownOpen = ref(false)
const searchQuery = ref('')
const barHeight = 50
const barSpacing = 12

const filteredProjects = computed(() => {
  if (!searchQuery.value) {
    return projects.value
  }
  return projects.value.filter((project) =>
    project.name?.toLowerCase().includes(searchQuery.value.toLowerCase()),
  )
})

const displayedProjects = computed(() => {
  if (selectedProjects.value.length === 0) {
    return projects.value
  }
  return projects.value.filter((project) => selectedProjects.value.includes(project.id))
})

const maxDuration = computed(() => {
  if (displayedProjects.value.length === 0) return 100
  return Math.max(...displayedProjects.value.map((p) => calculateDuration(p)))
})

const barsHeight = computed(() => {
  return displayedProjects.value.length * (barHeight + barSpacing)
})

const calculateDuration = (project: any) => {
  if (!project.started_at || !project.finished_at) return 0
  const start = new Date(project.started_at).getTime()
  const end = new Date(project.finished_at).getTime()
  return Math.round((end - start) / (1000 * 60 * 60 * 24))
}

const getBarWidth = (project: any) => {
  const duration = calculateDuration(project)
  return Math.max((duration / maxDuration.value) * 100, 2)
}

const formatDate = (dateString: string) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString('th-TH', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
  })
}

const getRandomProgress = () => {
  return Math.floor(Math.random() * 40) + 60 // 60-100%
}

const getRandomTasks = () => {
  return Math.floor(Math.random() * 50) + 30 // 30-80 tasks
}

const getRandomSubtasks = () => {
  return Math.floor(Math.random() * 30) + 10 // 10-40 subtasks
}

const getTeamMembers = (project: any) => {
  const members = []
  const ownerInitial = project.owner?.username?.charAt(0)?.toUpperCase() || 'U'
  members.push(ownerInitial)

  if (project.members) {
    const memberInitials = project.members
      .slice(0, 2)
      .map((m: any) => m.username?.charAt(0)?.toUpperCase() || 'M')
    members.push(...memberInitials)
  }

  return members.slice(0, 3)
}

// Dropdown functions
const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

const toggleProject = (projectId: number) => {
  const index = selectedProjects.value.indexOf(projectId)
  if (index > -1) {
    selectedProjects.value.splice(index, 1)
  } else {
    selectedProjects.value.push(projectId)
  }
}

const selectAllProjects = () => {
  selectedProjects.value = filteredProjects.value.map((p) => p.id)
}

const clearAllProjects = () => {
  selectedProjects.value = []
}

const getProjectName = (projectId: number) => {
  const project = projects.value.find((p) => p.id === projectId)
  return project?.name || 'Unknown Project'
}

const onSearch = () => {
  // Search is handled by computed property filteredProjects
}

const getStatusClass = (status: string) => {
  return {
    'status-active': status === 'Active',
    'status-pending': status === 'Pending',
    'status-delayed': status === 'Delayed',
  }
}

// Close dropdown when clicking outside
const handleClickOutside = (event: Event) => {
  const dropdown = document.querySelector('.dropdown-container')
  if (dropdown && !dropdown.contains(event.target as Node)) {
    isDropdownOpen.value = false
  }
}

onMounted(async () => {
  try {
    const res = await api.get('/projects')
    projects.value = res.data.data.filter(
      (p: { started_at: any; finished_at: any }) => p.started_at && p.finished_at,
    )

    // Add click outside listener
    document.addEventListener('click', handleClickOutside)
  } catch (error) {
    console.error('Error fetching projects:', error)
  }
})

// Cleanup event listener
const cleanup = () => {
  document.removeEventListener('click', handleClickOutside)
}
</script>

<style scoped>
.dashboard-container {
  background-color: #f8fafc;
  min-height: 100vh;
  padding: 0;
}

/* Header */
.dashboard-header {
  background: white;
  border-bottom: 1px solid #e2e8f0;
  padding: 1rem 2rem;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.breadcrumb {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #64748b;
}

.breadcrumb-icon {
  width: 16px;
  height: 16px;
}

.breadcrumb-text {
  font-size: 1.125rem;
  font-weight: 600;
  color: #334155;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.action-icon {
  width: 20px;
  height: 20px;
  color: #64748b;
  cursor: pointer;
}

.action-icon:hover {
  color: #334155;
}

.profile-avatar {
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.875rem;
}

/* Filter Section */
.filter-section {
  padding: 1.5rem 2rem 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Filter Section */
.filter-section {
  padding: 1.5rem 2rem 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.filter-dropdown-wrapper {
  position: relative;
}

.dropdown-container {
  position: relative;
}

.project-filter {
  background: white;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  padding: 0.75rem 2.5rem 0.75rem 1rem;
  font-size: 0.875rem;
  color: #374151;
  cursor: pointer;
  min-width: 250px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  transition: all 0.2s ease;
}

.project-filter:hover {
  border-color: #10b981;
}

.dropdown-container.dropdown-open .project-filter {
  border-color: #10b981;
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.1);
}

.filter-placeholder {
  color: #6b7280;
}

.filter-selected {
  color: #374151;
  font-weight: 500;
}

.dropdown-arrow {
  width: 16px;
  height: 16px;
  color: #6b7280;
  transition: transform 0.2s ease;
  flex-shrink: 0;
}

.dropdown-container.dropdown-open .dropdown-arrow {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: white;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
  z-index: 50;
  margin-top: 0.25rem;
  max-height: 320px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.dropdown-header {
  padding: 0.75rem;
  border-bottom: 1px solid #e5e7eb;
}

.search-input {
  width: 100%;
  padding: 0.5rem 0.75rem;
  border: 1px solid #d1d5db;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  outline: none;
  transition: border-color 0.2s ease;
}

.search-input:focus {
  border-color: #10b981;
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.1);
}

.dropdown-actions {
  padding: 0.5rem 0.75rem;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  gap: 0.5rem;
}

.action-btn {
  background: none;
  border: none;
  color: #10b981;
  font-size: 0.75rem;
  font-weight: 500;
  cursor: pointer;
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  transition: background-color 0.2s ease;
}

.action-btn:hover {
  background: #f0fdf4;
}

.dropdown-options {
  flex: 1;
  overflow-y: auto;
  max-height: 200px;
}

.option-item {
  display: flex;
  align-items: center;
  padding: 0.75rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
  border-bottom: 1px solid #f3f4f6;
}

.option-item:hover {
  background: #f9fafb;
}

.option-item:last-child {
  border-bottom: none;
}

.option-checkbox {
  margin-right: 0.75rem;
  width: 16px;
  height: 16px;
  accent-color: #10b981;
}

.option-content {
  flex: 1;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.option-name {
  font-size: 0.875rem;
  color: #374151;
  font-weight: 500;
}

.option-status {
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
  border-radius: 0.375rem;
  font-weight: 600;
}

.option-status.status-active {
  background: #dcfce7;
  color: #16a34a;
}

.option-status.status-pending {
  background: #fef3c7;
  color: #d97706;
}

.option-status.status-delayed {
  background: #fee2e2;
  color: #dc2626;
}

.no-results {
  padding: 1rem;
  text-align: center;
  color: #6b7280;
  font-size: 0.875rem;
}

.last-updated {
  font-size: 0.875rem;
  color: #64748b;
}

/* Main Content */
.main-content {
  padding: 0 2rem 2rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

/* Section Headers */
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.section-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #1f2937;
}

.view-all-btn {
  color: #06b6d4;
  font-size: 0.875rem;
  font-weight: 500;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
}

.view-all-btn:hover {
  color: #0891b2;
}

/* Chart Section */
.chart-section {
  background: white;
  border-radius: 0.75rem;
  padding: 1.5rem;
  border: 1px solid #e2e8f0;
}

.chart-wrapper {
  margin-top: 1rem;
}

.chart-content {
  height: 300px;
  display: flex;
  flex-direction: column;
}

.chart-scroll-container {
  flex: 1;
  overflow-y: auto;
  padding-right: 0.5rem;
}

.chart-bars {
  width: 100%;
}

.bar-row {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
  gap: 1rem;
}

.project-info {
  width: 200px;
  flex-shrink: 0;
}

.project-name {
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.bar-container {
  flex: 1;
  position: relative;
}

.bar {
  height: 32px;
  background: linear-gradient(90deg, #10b981, #059669);
  border-radius: 6px;
  position: relative;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding-right: 0.75rem;
  min-width: 60px;
}

.bar:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.bar-label {
  color: white;
  font-size: 0.75rem;
  font-weight: 600;
}

.chart-x-axis {
  margin-top: 1rem;
  padding-top: 0.75rem;
  border-top: 1px solid #e2e8f0;
  margin-left: 216px;
}

.x-axis-labels {
  display: flex;
  justify-content: space-between;
  font-size: 0.75rem;
  color: #6b7280;
  margin-bottom: 0.5rem;
}

.x-axis-title {
  text-align: center;
  font-size: 0.875rem;
  color: #64748b;
  font-weight: 500;
}

/* Projects Section */
.projects-section {
  background: white;
  border-radius: 0.75rem;
  padding: 1.5rem;
  border: 1px solid #e2e8f0;
}

.projects-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 1.5rem;
  max-height: 500px;
  overflow-y: auto;
  padding-right: 0.5rem;
}

.project-card {
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 0.75rem;
  padding: 1.25rem;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.project-card:hover {
  border-color: #10b981;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.project-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1rem;
}

.project-title {
  font-size: 1rem;
  font-weight: 600;
  color: #1f2937;
  line-height: 1.4;
}

.project-stats {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #10b981;
}

.member-count {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  font-size: 1.25rem;
  font-weight: 600;
}

.stat-icon {
  width: 16px;
  height: 16px;
}

.member-label {
  font-size: 0.875rem;
  color: #64748b;
}

.progress-section {
  margin-bottom: 1rem;
}

.progress-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.progress-icon {
  font-size: 1rem;
}

.progress-text {
  font-size: 0.875rem;
  color: #64748b;
  flex: 1;
}

.progress-percentage {
  font-size: 0.875rem;
  font-weight: 600;
  color: #10b981;
}

.progress-bar {
  height: 8px;
  background: #e2e8f0;
  border-radius: 4px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #10b981, #059669);
  border-radius: 4px;
  transition: width 0.3s ease;
}

.project-dates {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;
  padding: 0.75rem;
  background: #f8fafc;
  border-radius: 0.5rem;
}

.date-info {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.date-label {
  font-size: 0.75rem;
  color: #64748b;
  font-weight: 500;
}

.date-value {
  font-size: 0.875rem;
  color: #374151;
  font-weight: 600;
}

.team-avatars {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.avatar {
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, #10b981, #059669);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.75rem;
  font-weight: 600;
  border: 2px solid white;
  margin-right: -8px;
  position: relative;
}

.avatar:first-child {
  margin-right: -8px;
}

.task-summary {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-left: auto;
  color: #64748b;
  font-size: 0.875rem;
}

.task-icon {
  width: 16px;
  height: 16px;
}

/* Scrollbar */
.chart-scroll-container::-webkit-scrollbar {
  width: 6px;
}

.chart-scroll-container::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 3px;
}

.chart-scroll-container::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.chart-scroll-container::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>
