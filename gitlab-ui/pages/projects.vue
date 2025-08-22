<template>
  <div class="space-y-6">
    <!-- Page Header -->
    <div class="text-center">
      <h1 class="text-3xl font-bold text-gray-900">Project Management</h1>
      <p class="mt-2 text-gray-600">Organize and manage your development workflow</p>
    </div>

    <!-- Coming Soon Section -->
    <div class="card text-center py-12">
      <div class="w-16 h-16 bg-orange-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <svg class="w-8 h-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      </div>
      <h2 class="text-2xl font-semibold text-gray-900 mb-2">Coming Soon!</h2>
      <p class="text-gray-600 mb-6">
        We're working hard to bring you advanced project management features.
      </p>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto">
        <div class="text-center">
          <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-3">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
          <h3 class="font-semibold text-gray-900 mb-2">Issue Tracking</h3>
          <p class="text-sm text-gray-600">Track bugs, feature requests, and tasks</p>
        </div>
        
        <div class="text-center">
          <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mx-auto mb-3">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3a2 2 0 012-2h4a2 2 0 012 2v4m-6 0h6m-6 0l-1 8m1-8l1 8m-1 8v4a2 2 0 01-2 2H8a2 2 0 01-2-2v-4m6 0H8m6 0l1-8m-1 8l-1-8" />
            </svg>
          </div>
          <h3 class="font-semibold text-gray-900 mb-2">Merge Requests</h3>
          <p class="text-sm text-gray-600">Code review and collaboration workflow</p>
        </div>
        
        <div class="text-center">
          <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mx-auto mb-3">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <h3 class="font-semibold text-gray-900 mb-2">CI/CD Pipeline</h3>
          <p class="text-sm text-gray-600">Automated testing and deployment</p>
        </div>
      </div>
    </div>

    <!-- Roadmap -->
    <div class="card">
      <h2 class="text-xl font-semibold text-gray-900 mb-6">Development Roadmap</h2>
      
      <div class="space-y-6">
        <!-- Phase 1 -->
        <div class="border-l-4 border-green-500 pl-4">
          <div class="flex items-center mb-2">
            <div class="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center mr-3">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Phase 1 - Foundations ✅</h3>
          </div>
          <p class="text-gray-600 mb-2">User authentication, repository management, and Git over HTTP</p>
          <div class="flex flex-wrap gap-2">
            <span class="badge badge-success">Complete</span>
            <span class="badge badge-success">User Auth</span>
            <span class="badge badge-success">Repo Management</span>
            <span class="badge badge-success">Git HTTP</span>
          </div>
        </div>

        <!-- Phase 2 -->
        <div class="border-l-4 border-orange-500 pl-4">
          <div class="flex items-center mb-2">
            <div class="w-6 h-6 bg-orange-500 rounded-full flex items-center justify-center mr-3">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Phase 2 - Collaboration 🚧</h3>
          </div>
          <p class="text-gray-600 mb-2">Branches, commits, and basic collaboration features</p>
          <div class="flex flex-wrap gap-2">
            <span class="badge badge-warning">In Progress</span>
            <span class="badge badge-warning">Branches</span>
            <span class="badge badge-warning">Commits</span>
            <span class="badge badge-warning">Basic UI</span>
          </div>
        </div>

        <!-- Phase 3 -->
        <div class="border-l-4 border-gray-300 pl-4">
          <div class="flex items-center mb-2">
            <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center mr-3">
              <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Phase 3 - Advanced Features 📋</h3>
          </div>
          <p class="text-gray-600 mb-2">Merge requests, issue tracking, and CI/CD</p>
          <div class="flex flex-wrap gap-2">
            <span class="badge badge-private">Planned</span>
            <span class="badge badge-private">Merge Requests</span>
            <span class="badge badge-private">Issues</span>
            <span class="badge badge-private">CI/CD</span>
          </div>
        </div>

        <!-- Phase 4 -->
        <div class="border-l-4 border-gray-300 pl-4">
          <div class="flex items-center mb-2">
            <div class="w-6 h-6 bg-gray-300 rounded-full flex items-center justify-center mr-3">
              <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900">Phase 4 - Enterprise Features 🚀</h3>
          </div>
          <p class="text-gray-600 mb-2">Scaling, permissions, and enterprise features</p>
          <div class="flex flex-wrap gap-2">
            <span class="badge badge-private">Future</span>
            <span class="badge badge-private">Permissions</span>
            <span class="badge badge-private">Webhooks</span>
            <span class="badge badge-private">Scaling</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Current Status -->
    <div class="card">
      <h2 class="text-xl font-semibold text-gray-900 mb-4">Current Status</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div class="bg-green-50 border border-green-200 rounded-lg p-4">
          <div class="flex items-center">
            <svg class="w-5 h-5 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            <span class="font-medium text-green-800">Backend API</span>
          </div>
          <p class="text-sm text-green-700 mt-1">Go backend with PostgreSQL and Git integration</p>
        </div>
        
        <div class="bg-green-50 border border-green-200 rounded-lg p-4">
          <div class="flex items-center">
            <svg class="w-5 h-5 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            <span class="font-medium text-green-800">Frontend UI</span>
          </div>
          <p class="text-sm text-green-700 mt-1">Nuxt 3 with Tailwind CSS and modern design</p>
        </div>
        
        <div class="bg-orange-50 border border-orange-200 rounded-lg p-4">
          <div class="flex items-center">
            <svg class="w-5 h-5 text-orange-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="font-medium text-orange-800">Git Operations</span>
          </div>
          <p class="text-sm text-orange-700 mt-1">Basic Git over HTTP with Smart HTTP protocol</p>
        </div>
        
        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <div class="flex items-center">
            <svg class="w-5 h-5 text-blue-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="font-medium text-blue-800">Next Steps</span>
          </div>
          <p class="text-sm text-blue-700 mt-1">Implementing branches, commits, and collaboration features</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useAuthStore } from '~/stores/auth'

// Auth store
const authStore = useAuthStore()

// Check authentication on mount
onMounted(async () => {
  await authStore.checkAuth()
})
</script> 