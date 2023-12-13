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

const showDrawer = ref(false)
const toggleDrawer = () => (showDrawer.value = !showDrawer.value)
</script>

<style scoped></style>
