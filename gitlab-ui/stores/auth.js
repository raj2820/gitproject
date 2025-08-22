import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: null,
    isAuthenticated: false
  }),

  getters: {
    userInitials: (state) => {
      if (!state.user?.username) return '?'
      return state.user.username.substring(0, 2).toUpperCase()
    }
  },

  actions: {
    async login(credentials) {
      try {
        const config = useRuntimeConfig()
        const response = await $fetch(`${config.public.apiBase}/auth/login`, {
          method: 'POST',
          body: credentials
        })

        this.user = response.user
        this.token = response.token
        this.isAuthenticated = true

        // Store in localStorage
        localStorage.setItem('auth_token', response.token)
        localStorage.setItem('auth_user', JSON.stringify(response.user))

        return { success: true, data: response }
      } catch (error) {
        console.error('Login error:', error)
        throw error
      }
    },

    async signup(userData) {
      try {
        const config = useRuntimeConfig()
        const response = await $fetch(`${config.public.apiBase}/auth/signup`, {
          method: 'POST',
          body: userData
        })

        this.user = response.user
        this.token = response.token
        this.isAuthenticated = true

        // Store in localStorage
        localStorage.setItem('auth_token', response.token)
        localStorage.setItem('auth_user', JSON.stringify(response.user))

        return { success: true, data: response }
      } catch (error) {
        console.error('Signup error:', error)
        throw error
      }
    },

    logout() {
      this.user = null
      this.token = null
      this.isAuthenticated = false

      // Clear localStorage
      localStorage.removeItem('auth_token')
      localStorage.removeItem('auth_user')
    },

    async checkAuth() {
      const token = localStorage.getItem('auth_token')
      const user = localStorage.getItem('auth_user')

      if (token && user) {
        try {
          this.token = token
          this.user = JSON.parse(user)
          this.isAuthenticated = true

          // Verify token is still valid
          const config = useRuntimeConfig()
          await $fetch(`${config.public.apiBase}/api/repos`, {
            headers: {
              'Authorization': `Bearer ${token}`
            }
          })

          return true
        } catch (error) {
          // Token is invalid, clear auth state
          this.logout()
          return false
        }
      }
      return false
    },

    getAuthHeaders() {
      if (!this.token) return {}
      return {
        'Authorization': `Bearer ${this.token}`
      }
    }
  }
}) 