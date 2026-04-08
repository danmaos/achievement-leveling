<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()

onMounted(async () => {
  const token = route.query.token as string
  if (token) {
    auth.setToken(token)
    await auth.fetchUser()
    router.replace('/')
  } else {
    router.replace('/login')
  }
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center">
    <div class="text-center">
      <div class="animate-spin w-8 h-8 border-4 border-indigo-400 border-t-transparent rounded-full mx-auto mb-4"></div>
      <p class="text-gray-400">Signing you in...</p>
    </div>
  </div>
</template>
