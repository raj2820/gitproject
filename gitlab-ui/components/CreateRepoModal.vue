<template>
  <div class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex min-h-screen items-center justify-center p-4">
      <!-- Backdrop -->
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="$emit('close')"></div>

      <!-- Modal -->
      <div class="relative transform overflow-hidden rounded-lg bg-white shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-2xl">
        <!-- Header -->
        <div class="bg-gradient-to-r from-orange-500 to-red-500 px-6 py-4">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-white">Create New Repository</h3>
            <button @click="$emit('close')" class="text-white hover:text-gray-200">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Body -->
        <div class="px-6 py-6">
          <form @submit.prevent="handleSubmit" class="space-y-6">
            <!-- Repository Name -->
            <div>
              <label for="repo-name" class="block text-sm font-medium text-gray-700 mb-1">
                Repository name <span class="text-red-500">*</span>
              </label>
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-500">
                  {{ currentUser?.username }}/
                </span>
                <input
                  id="repo-name"
                  v-model="form.name"
                  type="text"
                  required
                  class="input-field pl-20"
                  placeholder="repository-name"
                  pattern="[a-zA-Z0-9_-]+"
                  title="Repository name can only contain letters, numbers, hyphens, and underscores"
                />
              </div>
              <p class="mt-1 text-sm text-gray-500">
                Great repository names are short and memorable. Need inspiration? How about <span class="text-orange-600 cursor-pointer hover:underline" @click="suggestName">suggesting a name</span>?
              </p>
            </div>

            <!-- Description -->
            <div>
              <label for="repo-description" class="block text-sm font-medium text-gray-700 mb-1">
                Description
              </label>
              <textarea
                id="repo-description"
                v-model="form.description"
                rows="3"
                class="input-field"
                placeholder="Add a description to help people understand your project"
              ></textarea>
            </div>

            <!-- Visibility -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-3">Visibility</label>
              <div class="space-y-3">
                <label class="flex items-center cursor-pointer">
                  <input
                    v-model="form.visibility"
                    type="radio"
                    value="public"
                    class="w-4 h-4 text-orange-600 border-gray-300 focus:ring-orange-500"
                  />
                  <div class="ml-3">
                    <div class="flex items-center">
                      <svg class="w-5 h-5 text-gray-400 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                      </svg>
                      <span class="font-medium text-gray-900">Public</span>
                    </div>
                    <p class="text-sm text-gray-500">Anyone on the internet can see this repository. You choose who can commit.</p>
                  </div>
                </label>

                <label class="flex items-center cursor-pointer">
                  <input
                    v-model="form.visibility"
                    type="radio"
                    value="private"
                    class="w-4 h-4 text-orange-600 border-gray-300 focus:ring-orange-500"
                  />
                  <div class="ml-3">
                    <div class="flex items-center">
                      <svg class="w-5 h-5 text-gray-400 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21" />
                      </svg>
                      <span class="font-medium text-gray-900">Private</span>
                    </div>
                    <p class="text-sm text-gray-500">You choose who can see and commit to this repository.</p>
                  </div>
                </label>
              </div>
            </div>

            <!-- README Options -->
            <div class="border-t border-gray-200 pt-6">
              <div class="flex items-center mb-4">
                <input
                  id="add-readme"
                  v-model="form.addReadme"
                  type="checkbox"
                  class="w-4 h-4 text-orange-600 border-gray-300 rounded focus:ring-orange-500"
                />
                <label for="add-readme" class="ml-2 text-sm font-medium text-gray-900">
                  Add a README file
                </label>
              </div>

              <div v-if="form.addReadme" class="ml-6 space-y-4">
                <div>
                  <label for="readme-type" class="block text-sm font-medium text-gray-700 mb-1">
                    README Type
                  </label>
                  <select
                    id="readme-type"
                    v-model="form.readmeType"
                    class="input-field"
                  >
                    <option value="markdown">Markdown (.md)</option>
                    <option value="text">Plain Text (.txt)</option>
                  </select>
                </div>

                <div>
                  <label for="readme-title" class="block text-sm font-medium text-gray-700 mb-1">
                    README Title
                  </label>
                  <input
                    id="readme-title"
                    v-model="form.readmeTitle"
                    type="text"
                    class="input-field"
                    placeholder="Repository name will be used if empty"
                  />
                </div>
              </div>
            </div>

            <!-- Default Branch Options -->
            <div class="border-t border-gray-200 pt-6">
              <div class="flex items-center mb-4">
                <input
                  id="custom-branch"
                  v-model="form.customBranch"
                  type="checkbox"
                  class="w-4 h-4 text-orange-600 border-gray-300 rounded focus:ring-orange-500"
                />
                <label for="custom-branch" class="ml-2 text-sm font-medium text-gray-900">
                  Customize default branch
                </label>
              </div>

              <div v-if="form.customBranch" class="ml-6 space-y-4">
                <div>
                  <label for="default-branch" class="block text-sm font-medium text-gray-700 mb-1">
                    Default Branch Name
                  </label>
                  <input
                    id="default-branch"
                    v-model="form.defaultBranch"
                    type="text"
                    class="input-field"
                    placeholder="main"
                    pattern="[a-zA-Z0-9_-]+"
                    title="Branch name can only contain letters, numbers, hyphens, and underscores"
                  />
                  <p class="mt-1 text-sm text-gray-500">
                    This will be the default branch for your repository. Common names are 'main', 'master', or 'develop'.
                  </p>
                </div>
              </div>
            </div>

            <!-- Error Message -->
            <div v-if="error" class="bg-red-50 border border-red-200 rounded-md p-3">
              <p class="text-sm text-red-600">{{ error }}</p>
            </div>

            <!-- Action Buttons -->
            <div class="flex justify-end space-x-3 pt-4">
              <button
                type="button"
                @click="$emit('close')"
                class="btn-secondary"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="loading || !form.name"
                class="btn-primary disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span v-if="loading" class="flex items-center">
                  <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  Creating...
                </span>
                <span v-else>Create repository</span>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useAuthStore } from '~/stores/auth'

// Props
const props = defineProps({})

// Emits
const emit = defineEmits(['close', 'created'])

// Auth store
const authStore = useAuthStore()

// Reactive state
const loading = ref(false)
const error = ref('')

// Form data
const form = reactive({
  name: '',
  description: '',
  visibility: 'public',
  addReadme: true,
  readmeType: 'markdown',
  readmeTitle: '',
  customBranch: false,
  defaultBranch: 'main'
})

// Computed properties
const currentUser = computed(() => authStore.user)

// Methods
const handleSubmit = async () => {
  if (!form.name.trim()) return

  loading.value = true
  error.value = ''

  try {
    const config = useRuntimeConfig()
    const response = await $fetch(`${config.public.apiBase}/api/repos`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...authStore.getAuthHeaders()
      },
      body: {
        name: form.name.trim(),
        description: form.description.trim(),
        visibility: form.visibility,
        add_readme: form.addReadme,
        readme_type: form.readmeType,
        readme_title: form.readmeTitle.trim() || form.name.trim(),
        custom_branch: form.customBranch,
        default_branch: form.customBranch ? form.defaultBranch : 'main'
      }
    })

    emit('created', response)
  } catch (err) {
    error.value = err.data?.error || err.message || 'Failed to create repository. Please try again.'
  } finally {
    loading.value = false
  }
}

const suggestName = () => {
  const suggestions = [
    'awesome-project',
    'my-first-repo',
    'learning-project',
    'demo-app',
    'portfolio-site'
  ]
  form.name = suggestions[Math.floor(Math.random() * suggestions.length)]
}
</script> 