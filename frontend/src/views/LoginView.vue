<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import GoogleLoginButton from '@/components/GoogleLoginButton.vue'

const auth = useAuthStore()
</script>

<template>
  <div class="min-h-screen flex items-center justify-center">
    <div class="text-center">
      <h1 class="text-5xl font-bold mb-2 text-indigo-400">Achievement Leveling</h1>
      <p class="text-gray-400 mb-8 text-lg">Track your achievements. Earn XP. Level up.</p>
      <div class="bg-gray-800 rounded-xl p-8 shadow-2xl max-w-sm mx-auto">
        <h2 class="text-xl font-semibold mb-6">Get Started</h2>
        <GoogleLoginButton @click="auth.loginWithGoogle()" />
      </div>
    </div>
  </div>
</template>
