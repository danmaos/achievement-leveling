<script setup lang="ts">
defineProps<{
  percent: number
}>()
</script>

<template>
  <div class="w-full bg-gray-700 rounded-full h-3 overflow-hidden">
    <div
      class="bg-gradient-to-r from-indigo-500 to-purple-500 h-full rounded-full transition-all duration-500"
      :style="{ width: `${percent}%` }"
    />
  </div>
</template>
