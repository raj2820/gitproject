<template>
  <div class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex min-h-screen items-center justify-center p-4">
      <!-- Backdrop -->
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="$emit('close')"></div>

      <!-- Modal -->
      <div class="relative transform overflow-hidden rounded-lg bg-white shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
        <!-- Header -->
        <div class="bg-gradient-to-r from-orange-500 to-red-500 px-6 py-4">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-white">
              {{ type === 'login' ? 'Sign In' : 'Create Account' }}
            </h3>
            <button @click="$emit('close')" class="text-white hover:text-gray-200">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Body -->
        <div class="px-6 py-6">
          <form @submit.prevent="handleSubmit" class="space-y-4">
            <!-- Username -->
            <div>
              <label for="username" class="block text-sm font-medium text-gray-700 mb-1">
                Username
              </label>
              <input
                id="username"
                v-model="form.username"
                type="text"
                required
                class="input-field"
                placeholder="Enter your username"
              />
            </div>

            <!-- Email (only for signup) -->
            <div v-if="type === 'signup'">
              <label for="email" class="block text-sm font-medium text-gray-700 mb-1">
                Email
              </label>
              <input
                id="email"
                v-model="form.email"
                type="email"
                required
                class="input-field"
                placeholder="Enter your email"
              />
            </div>

            <!-- Password -->
            <div>
              <label for="password" class="block text-sm font-medium text-gray-700 mb-1">
                Password
              </label>
              <input
                id="password"
                v-model="form.password"
                type="password"
                required
                class="input-field"
                placeholder="Enter your password"
              />
            </div>

            <!-- Error Message -->
            <div v-if="error" class="bg-red-50 border border-red-200 rounded-md p-3">
              <p class="text-sm text-red-600">{{ error }}</p>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="loading"
              class="w-full btn-primary disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <span v-if="loading" class="flex items-center justify-center">
                <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                {{ loading ? 'Processing...' : (type === 'login' ? 'Sign In' : 'Create Account') }}
              </span>
              <span v-else>
                {{ type === 'login' ? 'Sign In' : 'Create Account' }}
              </span>
            </button>
          </form>

          <!-- Switch Mode -->
          <div class="mt-6 text-center">
            <p class="text-sm text-gray-600">
              {{ type === 'login' ? "Don't have an account?" : 'Already have an account?' }}
              <button
                @click="switchMode"
                class="font-medium text-orange-600 hover:text-orange-500"
              >
                {{ type === 'login' ? 'Sign up' : 'Sign in' }}
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useAuthStore } from '~/stores/auth'

// Props
const props = defineProps({
  type: {
    type: String,
    required: true,
    validator: (value) => ['login', 'signup'].includes(value)
  }
})

// Emits
const emit = defineEmits(['close', 'success'])

// Auth store
const authStore = useAuthStore()

// Reactive state
const loading = ref(false)
const error = ref('')

// Form data
const form = reactive({
  username: '',
  email: '',
  password: ''
})

// Methods
const handleSubmit = async () => {
  loading.value = true
  error.value = ''

  try {
    if (props.type === 'login') {
      await authStore.login({
        username: form.username,
        password: form.password
      })
    } else {
      await authStore.signup({
        username: form.username,
        email: form.email,
        password: form.password
      })
    }

    emit('success', authStore.user)
  } catch (err) {
    error.value = err.data?.error || err.message || 'An error occurred. Please try again.'
  } finally {
    loading.value = false
  }
}

const switchMode = () => {
  // Reset form
  form.username = ''
  form.email = ''
  form.password = ''
  error.value = ''
  
  // Emit close to parent, parent will switch modal type
  emit('close')
}
</script> 