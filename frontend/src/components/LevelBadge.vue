<script setup lang="ts">
defineProps<{
  level: number
  size?: 'sm' | 'md' | 'lg'
}>()
</script>

<template>
  <div
    :class="[
      'rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center font-bold text-white shadow-lg',
      size === 'lg' ? 'w-20 h-20 text-3xl' : size === 'sm' ? 'w-10 h-10 text-sm' : 'w-14 h-14 text-xl'
    ]"
  >
    {{ level }}
  </div>
</template>
