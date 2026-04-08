import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '@/api/client'
import type { Achievement, AchievementCategory, UserAchievement, CompleteResult } from '@/types'

export const useAchievementsStore = defineStore('achievements', () => {
  const achievements = ref<Achievement[]>([])
  const categories = ref<AchievementCategory[]>([])
  const userProgress = ref<UserAchievement[]>([])
  const loading = ref(false)

  async function fetchAchievements(categoryId?: string) {
    loading.value = true
    try {
      const params = categoryId ? { category_id: categoryId } : {}
      const { data } = await api.get<Achievement[]>('/api/achievements', { params })
      achievements.value = data
    } finally {
      loading.value = false
    }
  }

  async function fetchCategories() {
    const { data } = await api.get<AchievementCategory[]>('/api/categories')
    categories.value = data
  }

  async function fetchProgress() {
    const { data } = await api.get<UserAchievement[]>('/api/progress')
    userProgress.value = data
  }

  async function completeAchievement(achievementId: string): Promise<CompleteResult> {
    const { data } = await api.post<CompleteResult>('/api/progress/complete', {
      achievement_id: achievementId,
    })
    await fetchProgress()
    return data
  }

  function isCompleted(achievementId: string): boolean {
    return userProgress.value.some((p) => p.achievement_id === achievementId)
  }

  function getCompletionCount(achievementId: string): number {
    const progress = userProgress.value.find((p) => p.achievement_id === achievementId)
    return progress?.times_completed ?? 0
  }

  return {
    achievements,
    categories,
    userProgress,
    loading,
    fetchAchievements,
    fetchCategories,
    fetchProgress,
    completeAchievement,
    isCompleted,
    getCompletionCount,
  }
})
