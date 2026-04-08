import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/api/client'
import type { User } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('token'))

  const isAuthenticated = computed(() => !!token.value)

  function setToken(newToken: string) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
  }

  async function fetchUser() {
    if (!token.value) return
    try {
      const { data } = await api.get<User>('/auth/me')
      user.value = data
    } catch {
      logout()
    }
  }

  function loginWithGoogle() {
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
    window.location.href = `${apiUrl}/auth/google/login`
  }

  return { user, token, isAuthenticated, setToken, logout, fetchUser, loginWithGoogle }
})
