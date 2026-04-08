<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useAchievementsStore } from '@/stores/achievements'
import { useAuthStore } from '@/stores/auth'
import AchievementCard from '@/components/AchievementCard.vue'

const store = useAchievementsStore()
const auth = useAuthStore()
const selectedCategory = ref<string>('')
const completingId = ref<string>('')
const levelUpMessage = ref('')

onMounted(async () => {
  await Promise.all([
    store.fetchCategories(),
    store.fetchAchievements(),
    store.fetchProgress(),
  ])
})

async function filterByCategory(categoryId: string) {
  selectedCategory.value = categoryId
  await store.fetchAchievements(categoryId || undefined)
}

async function handleComplete(achievementId: string) {
  completingId.value = achievementId
  try {
    const result = await store.completeAchievement(achievementId)
    await auth.fetchUser()
    if (result.leveled_up && result.new_level) {
      levelUpMessage.value = `Level Up! You are now Level ${result.new_level.level}: ${result.new_level.title}!`
      setTimeout(() => { levelUpMessage.value = '' }, 4000)
    }
  } finally {
    completingId.value = ''
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <h1 class="text-2xl font-bold mb-6">Achievements</h1>

    <div
      v-if="levelUpMessage"
      class="bg-indigo-600 text-white p-4 rounded-xl mb-6 text-center font-bold text-lg animate-pulse"
    >
      {{ levelUpMessage }}
    </div>

    <div class="flex gap-2 mb-6 flex-wrap">
      <button
        @click="filterByCategory('')"
        :class="[
          'px-4 py-2 rounded-lg text-sm font-medium transition',
          selectedCategory === '' ? 'bg-indigo-600 text-white' : 'bg-gray-700 text-gray-300 hover:bg-gray-600'
        ]"
      >
        All
      </button>
      <button
        v-for="cat in store.categories"
        :key="cat.id"
        @click="filterByCategory(cat.id)"
        :class="[
          'px-4 py-2 rounded-lg text-sm font-medium transition',
          selectedCategory === cat.id ? 'text-white' : 'bg-gray-700 text-gray-300 hover:bg-gray-600'
        ]"
        :style="selectedCategory === cat.id ? { backgroundColor: cat.color } : {}"
      >
        {{ cat.name }}
      </button>
    </div>

    <div v-if="store.loading" class="text-center py-12">
      <div class="animate-spin w-8 h-8 border-4 border-indigo-400 border-t-transparent rounded-full mx-auto"></div>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <AchievementCard
        v-for="achievement in store.achievements"
        :key="achievement.id"
        :achievement="achievement"
        :completed="store.isCompleted(achievement.id)"
        :completion-count="store.getCompletionCount(achievement.id)"
        :completing="completingId === achievement.id"
        @complete="handleComplete(achievement.id)"
      />
    </div>
  </div>
</template>
