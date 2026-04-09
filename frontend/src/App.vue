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
      <div class="max-w-6xl mx-auto px-3 sm:px-4 py-2 sm:py-3">
        <div class="flex items-center justify-between">
          <RouterLink to="/" class="text-lg sm:text-xl font-bold text-indigo-400 shrink-0">
            Achievement Leveling
          </RouterLink>
          <div class="flex items-center gap-2 sm:gap-3">
            <img
              v-if="auth.user?.picture"
              :src="auth.user.picture"
              class="w-7 h-7 sm:w-8 sm:h-8 rounded-full"
              referrerpolicy="no-referrer"
            />
            <span class="text-sm text-gray-300 hidden sm:inline">{{ auth.user?.name }}</span>
            <button
              @click="handleLogout"
              class="text-xs sm:text-sm text-gray-400 hover:text-white"
            >
              Logout
            </button>
          </div>
        </div>
        <div class="flex gap-4 mt-2 text-sm overflow-x-auto">
          <RouterLink
            to="/"
            class="text-gray-300 hover:text-white transition whitespace-nowrap"
          >
            Dashboard
          </RouterLink>
          <RouterLink
            to="/achievements"
            class="text-gray-300 hover:text-white transition whitespace-nowrap"
          >
            Achievements
          </RouterLink>
          <RouterLink
            to="/profile"
            class="text-gray-300 hover:text-white transition whitespace-nowrap"
          >
            Profile
          </RouterLink>
        </div>
      </div>
    </nav>
    <main>
      <RouterView />
    </main>
  </div>
</template>
