<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useAchievementsStore } from '@/stores/achievements'
import LevelBadge from '@/components/LevelBadge.vue'
import ProgressBar from '@/components/ProgressBar.vue'

const auth = useAuthStore()
const achievements = useAchievementsStore()

const levelTitles: Record<number, string> = {
  1: 'Novice', 2: 'Beginner', 3: 'Apprentice', 4: 'Journeyman', 5: 'Adept',
  6: 'Expert', 7: 'Veteran', 8: 'Master', 9: 'Grandmaster', 10: 'Legend',
}

const xpThresholds: Record<number, number> = {
  1: 0, 2: 100, 3: 300, 4: 600, 5: 1000,
  6: 1500, 7: 2100, 8: 2800, 9: 3600, 10: 4500,
}

const currentXP = computed(() => auth.user?.xp ?? 0)
const currentLevel = computed(() => auth.user?.level ?? 1)
const levelTitle = computed(() => levelTitles[currentLevel.value] ?? 'Unknown')
const nextLevelXP = computed(() => xpThresholds[currentLevel.value + 1] ?? xpThresholds[10])
const currentLevelXP = computed(() => xpThresholds[currentLevel.value] ?? 0)
const progressPercent = computed(() => {
  const range = nextLevelXP.value - currentLevelXP.value
  if (range <= 0) return 100
  return Math.min(100, ((currentXP.value - currentLevelXP.value) / range) * 100)
})

const completedCount = computed(() => achievements.userProgress.length)

const recentProgress = computed(() =>
  [...achievements.userProgress]
    .sort((a, b) => new Date(b.completed_at).getTime() - new Date(a.completed_at).getTime())
    .slice(0, 5)
)

onMounted(async () => {
  await Promise.all([
    achievements.fetchAchievements(),
    achievements.fetchProgress(),
  ])
})
</script>

<template>
  <div class="max-w-4xl mx-auto px-3 sm:px-4 py-4 sm:py-8">
    <div class="bg-gray-800 rounded-xl p-4 sm:p-6 mb-4 sm:mb-6">
      <div class="flex items-center gap-3 sm:gap-6">
        <LevelBadge :level="currentLevel" />
        <div class="flex-1 min-w-0">
          <h1 class="text-lg sm:text-2xl font-bold truncate">{{ auth.user?.name }}</h1>
          <p class="text-indigo-400 font-medium">{{ levelTitle }}</p>
          <div class="mt-3">
            <div class="flex justify-between text-sm text-gray-400 mb-1">
              <span>{{ currentXP }} XP</span>
              <span>{{ nextLevelXP }} XP</span>
            </div>
            <ProgressBar :percent="progressPercent" />
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-2 sm:gap-4 mb-4 sm:mb-6">
      <div class="bg-gray-800 rounded-xl p-3 sm:p-4 text-center">
        <p class="text-2xl sm:text-3xl font-bold text-indigo-400">{{ currentLevel }}</p>
        <p class="text-gray-400 text-xs sm:text-sm">Level</p>
      </div>
      <div class="bg-gray-800 rounded-xl p-3 sm:p-4 text-center">
        <p class="text-2xl sm:text-3xl font-bold text-green-400">{{ currentXP }}</p>
        <p class="text-gray-400 text-xs sm:text-sm">Total XP</p>
      </div>
      <div class="bg-gray-800 rounded-xl p-3 sm:p-4 text-center">
        <p class="text-2xl sm:text-3xl font-bold text-yellow-400">{{ completedCount }}</p>
        <p class="text-gray-400 text-xs sm:text-sm">Completed</p>
      </div>
    </div>

    <div class="bg-gray-800 rounded-xl p-4 sm:p-6">
      <h2 class="text-base sm:text-lg font-semibold mb-3 sm:mb-4">Recent Activity</h2>
      <div v-if="recentProgress.length === 0" class="text-gray-500 text-center py-4">
        No achievements completed yet. Start your journey!
      </div>
      <div v-else class="space-y-3">
        <div
          v-for="progress in recentProgress"
          :key="progress.id"
          class="flex items-center justify-between bg-gray-700/50 rounded-lg p-3"
        >
          <div>
            <p class="font-medium">{{ progress.achievement.title }}</p>
            <p class="text-sm text-gray-400">
              +{{ progress.achievement.xp_reward }} XP
              <span v-if="progress.times_completed > 1" class="ml-2">
                (x{{ progress.times_completed }})
              </span>
            </p>
          </div>
          <span class="text-xs text-gray-500">
            {{ new Date(progress.completed_at).toLocaleDateString() }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
