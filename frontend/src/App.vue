<template>
  <v-app>
    <v-navigation-drawer v-model="showDrawer" app>
      <v-list-item href="/" title="Top" />
      <v-list-item href="/questions" title="Questions" />
      <v-list-item href="/questions/new" title="Post Question" />
    </v-navigation-drawer>
    <v-app-bar>
      <v-app-bar-nav-icon @click="toggleDrawer" />
      <router-link to="/">
        <v-btn variant="text" color="black">
          <v-app-bar-title>staQoverflow</v-app-bar-title>
        </v-btn>
      </router-link>
      <v-spacer></v-spacer>
      <div class="text-end mr-2">
        <template v-if="loginUser !== null">
          <v-btn color="primary" @click="showUserInfo">ユーザー情報</v-btn>
        </template>
        <template v-else>
          <v-btn color="primary" @click="handleLogin">ログイン</v-btn>
        </template>
      </div>
    </v-app-bar>
    <v-main>
      <Suspense>
        <RouterView />
      </Suspense>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import { ref, onMounted, provide } from 'vue'
import { getMe, type User } from '@/lib/api/users'
import { BASE } from '@/lib/api/index'
import { useRouter } from 'vue-router'

const showDrawer = ref(false)
const loginUser = ref<User | null>(null)
const toggleDrawer = () => (showDrawer.value = !showDrawer.value)

const handleLogin = async () => {
  try {
    const res = await fetch(`${BASE}/oauth2/params`, {})
    const data = await res.json()
    window.location.href = data.url
  } catch (error) {
    console.error(error)
  }
}

const router = useRouter()
const showUserInfo = () => {
  router.push(`/users/${loginUser.value?.id}`)
}

onMounted(async () => {
  try {
    const res = await getMe()
    loginUser.value = res
  } catch (error) {
    loginUser.value = null
    console.error(error)
  }
})

provide('loginUser', loginUser)
</script>

<style scoped></style>
