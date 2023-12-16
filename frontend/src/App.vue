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
      <template v-if="loginUser !== null">
        <v-btn color="primary" @click="showUserInfo">ユーザー情報</v-btn>
      </template>
      <template v-else>
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
import { ref, onMounted } from 'vue'
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
    console.log('onMounted')
    const res = await getMe()
    loginUser.value = res
  } catch (error) {
    loginUser.value = null
    console.error(error)
  }
})
</script>

<style scoped></style>
