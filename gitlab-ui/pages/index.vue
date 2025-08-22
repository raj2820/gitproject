<template>
  <div class="space-y-6">
    <!-- Welcome Header -->
    <div class="bg-gradient-to-r from-orange-500 to-red-500 rounded-lg p-8 text-white">
      <div class="max-w-3xl">
        <h1 class="text-3xl font-bold mb-2">
          Welcome back, {{ currentUser?.username || 'Developer' }}! 👋
        </h1>
        <p class="text-orange-100 text-lg">
          Ready to build something amazing? Create a new repository or continue working on your projects.
        </p>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="card hover:shadow-md transition-shadow cursor-pointer" @click="showCreateRepoModal = true">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
          </div>
          <div>
            <h3 class="font-semibold text-gray-900">New Repository</h3>
            <p class="text-sm text-gray-600">Create a new project</p>
          </div>
        </div>
      </div>

      <div class="card hover:shadow-md transition-shadow cursor-pointer" @click="navigateTo('/repos')">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
          </div>
          <div>
            <h3 class="font-semibold text-gray-900">Browse Repositories</h3>
            <p class="text-sm text-gray-600">View all projects</p>
          </div>
        </div>
      </div>

      <div class="card hover:shadow-md transition-shadow cursor-pointer" @click="navigateTo('/projects')">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
          <div>
            <h3 class="font-semibold text-gray-900">Project Management</h3>
            <p class="text-sm text-gray-600">Organize your work</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="card">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 00-2-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Total Repositories</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalRepos }}</p>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Public Repos</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.publicRepos }}</p>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-gray-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Private Repos</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.privateRepos }}</p>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-orange-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Member Since</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.memberSince }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Repositories -->
    <div class="card">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-xl font-semibold text-gray-900">Recent Repositories</h2>
        <button @click="navigateTo('/repos')" class="text-orange-600 hover:text-orange-500 font-medium">
          View all →
        </button>
      </div>

      <div v-if="loading" class="flex justify-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-orange-500"></div>
      </div>

      <div v-else-if="recentRepos.length === 0" class="text-center py-8">
        <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
          <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 00-2-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">No repositories yet</h3>
        <p class="text-gray-500 mb-4">Get started by creating your first repository</p>
        <button @click="showCreateRepoModal = true" class="btn-primary">
          Create Repository
        </button>
      </div>

      <div v-else class="space-y-4">
        <div v-for="repo in recentRepos" :key="repo.id" class="repo-item">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <div class="w-10 h-10 bg-orange-100 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 00-2-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                </svg>
              </div>
              <div>
                <h3 class="font-medium text-gray-900">{{ repo.name }}</h3>
                <p class="text-sm text-gray-500">{{ repo.description || 'No description' }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-3">
              <span :class="repo.visibility === 'public' ? 'badge badge-public' : 'badge badge-private'">
                {{ repo.visibility }}
              </span>
              <button @click="copyCloneUrl(repo)" class="text-gray-400 hover:text-gray-600">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create Repository Modal -->
    <CreateRepoModal v-if="showCreateRepoModal" @close="showCreateRepoModal = false" @created="handleRepoCreated" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '~/stores/auth'

// Auth store
const authStore = useAuthStore()

// Reactive state
const loading = ref(false)
const recentRepos = ref([])
const showCreateRepoModal = ref(false)

// Computed properties
const currentUser = computed(() => authStore.user)
const stats = computed(() => ({
  totalRepos: recentRepos.value.length,
  publicRepos: recentRepos.value.filter(repo => repo.visibility === 'public').length,
  privateRepos: recentRepos.value.filter(repo => repo.visibility === 'private').length,
  memberSince: currentUser.value?.created_at ? new Date(currentUser.value.created_at).getFullYear() : 'N/A'
}))

// Methods
const loadRecentRepos = async () => {
  if (!authStore.isAuthenticated) return

  loading.value = true
  try {
    const config = useRuntimeConfig()
    const response = await $fetch(`${config.public.apiBase}/api/repos`, {
      headers: authStore.getAuthHeaders()
    })
    recentRepos.value = response.slice(0, 5) // Show only 5 most recent
  } catch (error) {
    console.error('Failed to load repositories:', error)
  } finally {
    loading.value = false
  }
}

const copyCloneUrl = (repo) => {
  const url = `http://localhost:8080/git/${currentUser.value.username}/${repo.name}.git`
  navigator.clipboard.writeText(url).then(() => {
    // Show success message
    alert('Clone URL copied to clipboard!')
  })
}

const handleRepoCreated = () => {
  showCreateRepoModal.value = false
  loadRecentRepos()
}

// Lifecycle
onMounted(async () => {
  await authStore.checkAuth()
  if (authStore.isAuthenticated) {
    loadRecentRepos()
  }
})
</script> 