<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation Header -->
    <nav class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <!-- Logo and Brand -->
          <div class="flex items-center">
            <div class="flex-shrink-0 flex items-center">
              <div class="w-8 h-8 bg-gradient-to-r from-orange-500 to-red-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </div>
              <span class="ml-3 text-xl font-bold text-gradient">GitLab Tool</span>
            </div>
          </div>

          <!-- Navigation Links -->
          <div class="hidden md:flex items-center space-x-8">
            <NuxtLink to="/" class="nav-link" active-class="nav-link-active">
              Dashboard
            </NuxtLink>
            <NuxtLink to="/repos" class="nav-link" active-class="nav-link-active">
              Repositories
            </NuxtLink>
            <NuxtLink to="/projects" class="nav-link" active-class="nav-link-active">
              Projects
            </NuxtLink>
          </div>

          <!-- User Menu -->
          <div class="flex items-center space-x-4">
            <template v-if="isAuthenticated">
              <!-- Notifications -->
              <button class="p-2 text-gray-400 hover:text-gray-500 relative">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M9 11h.01M9 8h.01M9 5h.01" />
                </svg>
                <span class="absolute top-0 right-0 block h-2 w-2 rounded-full bg-red-400 ring-2 ring-white"></span>
              </button>

              <!-- User Avatar -->
              <div class="relative">
                <button @click="toggleUserMenu" class="flex items-center space-x-2 text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500">
                  <div class="w-8 h-8 bg-orange-500 rounded-full flex items-center justify-center">
                    <span class="text-white font-medium">{{ userInitials }}</span>
                  </div>
                  <span class="text-gray-700">{{ currentUser?.username }}</span>
                  <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>

                <!-- User Dropdown Menu -->
                <div v-if="showUserMenu" class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-50">
                  <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Your profile</a>
                  <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Settings</a>
                  <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Help</a>
                  <div class="border-t border-gray-100"></div>
                  <button @click="logout" class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Sign out
                  </button>
                </div>
              </div>
            </template>

            <!-- Auth Buttons -->
            <template v-else>
              <button @click="showLoginModal = true" class="btn-secondary">
                Sign in
              </button>
              <button @click="showSignupModal = true" class="btn-primary">
                Sign up
              </button>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <NuxtPage />
    </main>

    <!-- Auth Modals -->
    <AuthModal v-if="showLoginModal" type="login" @close="showLoginModal = false" @success="handleAuthSuccess" />
    <AuthModal v-if="showSignupModal" type="signup" @close="showSignupModal = false" @success="handleAuthSuccess" />

    <!-- Global Notifications -->
    <NotificationContainer />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '~/stores/auth'

// Auth store
const authStore = useAuthStore()

// Reactive state
const showUserMenu = ref(false)
const showLoginModal = ref(false)
const showSignupModal = ref(false)

// Computed properties
const isAuthenticated = computed(() => authStore.isAuthenticated)
const currentUser = computed(() => authStore.user)
const userInitials = computed(() => {
  if (!currentUser.value?.username) return '?'
  return currentUser.value.username.substring(0, 2).toUpperCase()
})

// Methods
const toggleUserMenu = () => {
  showUserMenu.value = !showUserMenu.value
}

const handleAuthSuccess = (userData) => {
  showLoginModal.value = false
  showSignupModal.value = false
  // The auth store will handle the rest
}

const logout = () => {
  authStore.logout()
  showUserMenu.value = false
}

// Close user menu when clicking outside
onMounted(() => {
  document.addEventListener('click', (e) => {
    if (!e.target.closest('.relative')) {
      showUserMenu.value = false
    }
  })
})
</script> 