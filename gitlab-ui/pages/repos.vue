<template>
  <div class="space-y-6">
    <!-- Page Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">Repositories</h1>
        <p class="mt-2 text-gray-600">Manage and organize your projects</p>
      </div>
      <button @click="showCreateRepoModal = true" class="btn-primary">
        <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        New Repository
      </button>
    </div>

    <!-- Search and Filters -->
    <div class="card">
      <div class="flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <div class="relative">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Search repositories..."
              class="input-field pl-10"
            />
          </div>
        </div>
        
        <div class="flex gap-2">
          <select v-model="visibilityFilter" class="input-field">
            <option value="">All</option>
            <option value="public">Public</option>
            <option value="private">Private</option>
          </select>
          
          <select v-model="sortBy" class="input-field">
            <option value="name">Name</option>
            <option value="created_at">Created</option>
            <option value="updated_at">Updated</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Repositories List -->
    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500"></div>
    </div>

    <div v-else-if="filteredRepos.length === 0" class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 00-2-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">
        {{ searchQuery || visibilityFilter ? 'No repositories found' : 'No repositories yet' }}
      </h3>
      <p class="text-gray-500 mb-4">
        {{ searchQuery || visibilityFilter ? 'Try adjusting your search or filters' : 'Get started by creating your first repository' }}
      </p>
      <button @click="showCreateRepoModal = true" class="btn-primary">
        Create Repository
      </button>
    </div>

    <div v-else class="grid gap-4">
      <div v-for="repo in filteredRepos" :key="repo.id" class="repo-item">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 00-2-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
              </svg>
            </div>
            
            <div>
              <h3 class="text-lg font-semibold text-gray-900">{{ repo.name }}</h3>
              <p class="text-sm text-gray-600">{{ repo.description || 'No description' }}</p>
              <div class="flex items-center space-x-4 mt-2 text-sm text-gray-500">
                <span :class="repo.visibility === 'public' ? 'badge badge-public' : 'badge badge-private'">
                  {{ repo.visibility }}
                </span>
                <span>Created {{ formatDate(repo.created_at) }}</span>
                <span v-if="repo.updated_at !== repo.created_at">
                  Updated {{ formatDate(repo.updated_at) }}
                </span>
              </div>
            </div>
          </div>
          
          <div class="flex items-center space-x-2">
            <button @click="copyCloneUrl(repo)" class="btn-secondary text-sm">
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
              Clone
            </button>
            
            <button @click="viewRepository(repo)" class="btn-primary text-sm">
              View
            </button>
            
            <button @click="deleteRepository(repo)" class="btn-danger text-sm">
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
              Delete
            </button>
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
const repositories = ref([])
const showCreateRepoModal = ref(false)
const searchQuery = ref('')
const visibilityFilter = ref('')
const sortBy = ref('name')

// Computed properties
const filteredRepos = computed(() => {
  let filtered = repositories.value

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(repo => 
      repo.name.toLowerCase().includes(query) ||
      (repo.description && repo.description.toLowerCase().includes(query))
    )
  }

  // Visibility filter
  if (visibilityFilter.value) {
    filtered = filtered.filter(repo => repo.visibility === visibilityFilter.value)
  }

  // Sort
  filtered.sort((a, b) => {
    switch (sortBy.value) {
      case 'name':
        return a.name.localeCompare(b.name)
      case 'created_at':
        return new Date(b.created_at) - new Date(a.created_at)
      case 'updated_at':
        return new Date(b.updated_at) - new Date(a.updated_at)
      default:
        return 0
    }
  })

  return filtered
})

// Methods
const loadRepositories = async () => {
  if (!authStore.isAuthenticated) return

  loading.value = true
  try {
    const config = useRuntimeConfig()
    const response = await $fetch(`${config.public.apiBase}/api/repos`, {
      headers: authStore.getAuthHeaders()
    })
    repositories.value = response
  } catch (error) {
    console.error('Failed to load repositories:', error)
    if (window.$notify) {
      window.$notify.error('Error', 'Failed to load repositories')
    }
  } finally {
    loading.value = false
  }
}

const copyCloneUrl = (repo) => {
  const url = `http://localhost:8080/git/${authStore.user.username}/${repo.name}.git`
  navigator.clipboard.writeText(url).then(() => {
    if (window.$notify) {
      window.$notify.success('Success', 'Clone URL copied to clipboard!')
    }
  })
}

const viewRepository = (repo) => {
  // Navigate to repository detail page (to be implemented)
  console.log('View repository:', repo.name)
}

const deleteRepository = async (repo) => {
  if (!confirm(`Are you sure you want to delete "${repo.name}"? This action cannot be undone.`)) {
    return
  }

  try {
    const config = useRuntimeConfig()
    await $fetch(`${config.public.apiBase}/api/repos/${repo.id}`, {
      method: 'DELETE',
      headers: authStore.getAuthHeaders()
    })
    
    repositories.value = repositories.value.filter(r => r.id !== repo.id)
    
    if (window.$notify) {
      window.$notify.success('Success', `Repository "${repo.name}" deleted successfully`)
    }
  } catch (error) {
    console.error('Failed to delete repository:', error)
    if (window.$notify) {
      window.$notify.error('Error', 'Failed to delete repository')
    }
  }
}

const handleRepoCreated = (repo) => {
  showCreateRepoModal.value = false
  repositories.value.unshift(repo)
  
  if (window.$notify) {
    window.$notify.success('Success', `Repository "${repo.name}" created successfully`)
  }
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  const now = new Date()
  const diffTime = Math.abs(now - date)
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays === 1) return 'today'
  if (diffDays === 2) return 'yesterday'
  if (diffDays <= 7) return `${diffDays - 1} days ago`
  if (diffDays <= 30) return `${Math.ceil(diffDays / 7)} weeks ago`
  if (diffDays <= 365) return `${Math.ceil(diffDays / 30)} months ago`
  return `${Math.ceil(diffDays / 365)} years ago`
}

// Lifecycle
onMounted(async () => {
  await authStore.checkAuth()
  if (authStore.isAuthenticated) {
    loadRepositories()
  }
})
</script> 