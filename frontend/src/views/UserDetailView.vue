<template>
  <div style="width: 80%; margin: auto">
    <h1>{{ userInfo?.name }} さんのページ</h1>
    <img :src="`${BASE}/imggen?user_id=${userInfo?.id}`" alt="User Image" />
  </div>
</template>

<script setup lang="ts">
import { BASE } from '@/lib/api/index'
import { ref, onMounted } from 'vue'
import { getUser, type User, type GetUserRequest } from '@/lib/api/users'
import { useRoute } from 'vue-router'

const route = useRoute()
const userInfo = ref<User>()
const userName = ref('')
const request: GetUserRequest = {
  id: route.params.id as string
}

onMounted(async () => {
  try {
    const res = await getUser(request)
    userInfo.value = res
    userName.value = res.name
  } catch (err) {
    console.error(err)
  }
})
</script>
