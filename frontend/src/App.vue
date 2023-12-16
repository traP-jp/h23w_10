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
      <!-- todo: add user icon / login button -->
      <v-btn color="primary" @click="handleLogin">ログイン</v-btn>
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
import { ref } from 'vue'
import { BASE } from '@/lib/api/index'
const showDrawer = ref(false)
const toggleDrawer = () => (showDrawer.value = !showDrawer.value)
const handleLogin = async () => {
  try {
    const res = await fetch(`${BASE}/oauth2/params`, {})
    const data = await res.json()
    console.log(data)
    window.location.href = data.url
  } catch (error) {
    console.error(error)
  }
  // ここにログイン処理を書く
}
</script>

<style scoped></style>
