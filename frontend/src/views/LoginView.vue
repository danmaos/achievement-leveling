<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import GoogleLoginButton from '@/components/GoogleLoginButton.vue'

const auth = useAuthStore()
</script>

<template>
  <div class="min-h-screen flex items-center justify-center px-4">
    <div class="text-center w-full">
      <h1 class="text-3xl sm:text-5xl font-bold mb-2 text-indigo-400">Achievement Leveling</h1>
      <p class="text-gray-400 mb-6 sm:mb-8 text-base sm:text-lg">Track your achievements. Earn XP. Level up.</p>
      <div class="bg-gray-800 rounded-xl p-6 sm:p-8 shadow-2xl max-w-sm mx-auto">
        <h2 class="text-xl font-semibold mb-6">Get Started</h2>
        <GoogleLoginButton @click="auth.loginWithGoogle()" />
      </div>
    </div>
  </div>
</template>
