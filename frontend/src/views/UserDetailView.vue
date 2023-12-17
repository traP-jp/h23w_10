<template>
  <div style="width: 80%; margin: auto">
    <v-row>
      <v-col cols="4">
        <h1>User Name: {{ userInfo?.name }}</h1>
        <img :src="userInfo?.icon_url" alt="User Icon" />
      </v-col>
      <v-col cols="8">
        <h1>円形に表示するやつ</h1>
        <img :src="userImage ?? ''" alt="show Circle Image" />
      </v-col>
    </v-row>
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
const userImage = ref<string | null>(null)
const request: GetUserRequest = {
  id: route.params.id as string
}

onMounted(async () => {
  try {
    const res = await getUser(request)
    userInfo.value = res
    userName.value = res.name
    const res2 = await fetch(`${BASE}/imggen`, {
      method: 'POST',
      body: JSON.stringify({
        // user_id: '1eea935c-0d3c-411b-a565-1b09565237f4'
        user_id: userInfo.value.id
      })
    })
    console.log(res2)
  } catch (err) {
    console.error(err)
  }
})
</script>
