<template>
  <div style="width: 80%; margin: auto">
    <v-row>
      <v-col cols="4">
        <h1>User Name: {{ userInfo?.name }}</h1>
        <img :src="userInfo?.iconURL" alt="User Icon" />
      </v-col>
      <v-col cols="8">
        <h1>円形に表示するやつ</h1>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
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
    console.log(userInfo.value.name)
    console.log(userInfo.value.iconURL)
  } catch (err) {
    console.error(err)
  }
})
</script>
