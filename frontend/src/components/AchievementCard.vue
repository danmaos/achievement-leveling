<script setup lang="ts">
import type { Achievement } from '@/types'

defineProps<{
  achievement: Achievement
  completed: boolean
  completionCount: number
  completing: boolean
}>()

defineEmits<{
  complete: []
}>()

const difficultyColors: Record<string, string> = {
  easy: 'text-green-400 bg-green-400/10',
  medium: 'text-yellow-400 bg-yellow-400/10',
  hard: 'text-orange-400 bg-orange-400/10',
  epic: 'text-purple-400 bg-purple-400/10',
}
</script>

<template>
  <div
    :class="[
      'bg-gray-800 rounded-xl p-3 sm:p-5 border transition',
      completed ? 'border-green-500/30' : 'border-gray-700 hover:border-gray-600'
    ]"
  >
    <div class="flex items-start justify-between mb-3">
      <div>
        <h3 class="font-semibold text-base sm:text-lg">{{ achievement.title }}</h3>
        <span
          v-if="achievement.category"
          class="text-xs px-2 py-0.5 rounded-full"
          :style="{ backgroundColor: achievement.category.color + '20', color: achievement.category.color }"
        >
          {{ achievement.category.name }}
        </span>
      </div>
      <span
        :class="['text-xs font-medium px-2 py-1 rounded-full', difficultyColors[achievement.difficulty]]"
      >
        {{ achievement.difficulty }}
      </span>
    </div>

    <p class="text-gray-400 text-sm mb-4">{{ achievement.description }}</p>

    <div class="flex items-center justify-between gap-2">
      <div class="flex items-center gap-1 sm:gap-2 shrink-0">
        <span class="text-indigo-400 font-bold text-sm sm:text-base">+{{ achievement.xp_reward }} XP</span>
        <span v-if="achievement.repeatable" class="text-xs text-gray-500 hidden sm:inline">(repeatable)</span>
      </div>

      <div v-if="completed && !achievement.repeatable" class="flex items-center gap-1 text-green-400 text-sm">
        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
        </svg>
        Completed
      </div>

      <button
        v-else
        @click="$emit('complete')"
        :disabled="completing"
        :class="[
          'px-3 sm:px-4 py-1.5 rounded-lg text-xs sm:text-sm font-medium transition',
          completing
            ? 'bg-gray-600 text-gray-400 cursor-not-allowed'
            : 'bg-indigo-600 hover:bg-indigo-500 text-white'
        ]"
      >
        <span v-if="completing">...</span>
        <span v-else-if="completionCount > 0">Complete Again ({{ completionCount }})</span>
        <span v-else>Complete</span>
      </button>
    </div>
  </div>
</template>
