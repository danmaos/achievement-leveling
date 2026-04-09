<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import LevelBadge from '@/components/LevelBadge.vue'

const auth = useAuthStore()

const levelTitles: Record<number, string> = {
  1: 'Novice', 2: 'Beginner', 3: 'Apprentice', 4: 'Journeyman', 5: 'Adept',
  6: 'Expert', 7: 'Veteran', 8: 'Master', 9: 'Grandmaster', 10: 'Legend',
}
</script>

<template>
  <div class="max-w-2xl mx-auto px-3 sm:px-4 py-4 sm:py-8">
    <div class="bg-gray-800 rounded-xl p-5 sm:p-8 text-center">
      <img
        v-if="auth.user?.picture"
        :src="auth.user.picture"
        class="w-24 h-24 rounded-full mx-auto mb-4 border-4 border-indigo-500"
        referrerpolicy="no-referrer"
      />
      <h1 class="text-2xl font-bold">{{ auth.user?.name }}</h1>
      <p class="text-gray-400 mb-6">{{ auth.user?.email }}</p>

      <div class="flex justify-center mb-6">
        <LevelBadge :level="auth.user?.level ?? 1" size="lg" />
      </div>

      <p class="text-indigo-400 text-lg font-semibold mb-2">
        {{ levelTitles[auth.user?.level ?? 1] }}
      </p>

      <div class="grid grid-cols-2 gap-4 mt-6">
        <div class="bg-gray-700/50 rounded-lg p-4">
          <p class="text-2xl font-bold text-green-400">{{ auth.user?.xp ?? 0 }}</p>
          <p class="text-sm text-gray-400">Total XP</p>
        </div>
        <div class="bg-gray-700/50 rounded-lg p-4">
          <p class="text-2xl font-bold text-indigo-400">{{ auth.user?.level ?? 1 }}</p>
          <p class="text-sm text-gray-400">Current Level</p>
        </div>
      </div>

      <p class="text-xs text-gray-500 mt-6">
        Member since {{ auth.user?.created_at ? new Date(auth.user.created_at).toLocaleDateString() : '...' }}
      </p>
    </div>
  </div>
</template>
