<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <div>
    <header>
      <div class="wrapper">
        <nav>
          <RouterLink to="/">Home</RouterLink> |
          <RouterLink to="/about">About</RouterLink>
        </nav>
      </div>
    </header>
    <Suspense>
      <RouterView />
    </Suspense>
  </div>
</template>

<style scoped></style>
