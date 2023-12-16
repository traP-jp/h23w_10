<template>
  <v-app>
    <v-navigation-drawer v-model="showDrawer" app>
      <v-list-item href="/" title="Top" />
      <v-list-item href="/questions" title="Questions" />
      <v-list-item href="/questions/new" title="Post Question" />
    </v-navigation-drawer>
    <v-app-bar>
      <v-app-bar-nav-icon @click="toggleDrawer" />
      <v-app-bar-title>staQoverflow</v-app-bar-title>
      <template v-if="!isLoggedIn">
        <v-icon>mdi-account</v-icon>
        <p>yu-za-na-me</p>
        <v-btn color="primary" @click="showUserInfo">ユーザー情報</v-btn>
      </template>
      <template v-if="!isLoggedIn">
        <v-btn color="primary" @click="handleLogin">ログイン</v-btn>
      </template>
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
import { ref, computed } from 'vue'
import { BASE } from '@/lib/api/index'
import { useRouter } from 'vue-router'

const showDrawer = ref(false)
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
const showUserInfo = async () => {
  try {
    const res = await fetch(`${BASE}/users/me`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('access_token')}`
      }
    })
    if (res.status === 401) {
      window.alert('ログイン画面に遷移します')
      handleLogin()
      return
    }
    const data = await res.json()
    router.push(`/users/${data.id}`)
  } catch (error) {
    console.error(error)
  }
}
const isLoggedIn = computed(() => localStorage.getItem('access_token'))
</script>

<style scoped></style>
