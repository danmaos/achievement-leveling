<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { RouterView, RouterLink, useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

onMounted(async () => {
  if (auth.isAuthenticated) {
    await auth.fetchUser()
  }
})

function handleLogout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-900 text-white">
    <nav v-if="auth.isAuthenticated" class="bg-gray-800 border-b border-gray-700">
      <div class="max-w-6xl mx-auto px-4 py-3 flex items-center justify-between">
        <div class="flex items-center gap-6">
          <RouterLink to="/" class="text-xl font-bold text-indigo-400">
            Achievement Leveling
          </RouterLink>
          <div class="flex gap-4">
            <RouterLink
              to="/"
              class="text-gray-300 hover:text-white transition"
            >
              Dashboard
            </RouterLink>
            <RouterLink
              to="/achievements"
              class="text-gray-300 hover:text-white transition"
            >
              Achievements
            </RouterLink>
            <RouterLink
              to="/profile"
              class="text-gray-300 hover:text-white transition"
            >
              Profile
            </RouterLink>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <img
            v-if="auth.user?.picture"
            :src="auth.user.picture"
            class="w-8 h-8 rounded-full"
            referrerpolicy="no-referrer"
          />
          <span class="text-sm text-gray-300">{{ auth.user?.name }}</span>
          <button
            @click="handleLogout"
            class="text-sm text-gray-400 hover:text-white ml-2"
          >
            Logout
          </button>
        </div>
      </div>
    </nav>
    <main>
      <RouterView />
    </main>
  </div>
</template>
