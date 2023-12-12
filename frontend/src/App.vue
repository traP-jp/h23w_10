<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink> | <RouterLink to="/about">About</RouterLink> |
        <RouterLink :to="'/questions/' + id">Question</RouterLink>
      </nav>
    </div>
    <RouterView />
  </header>
</template>
<script lang="ts">
import { RouterView } from 'vue-router'
export default {
  name: 'App',
  components: {
    RouterView
  },
  data() {
    return {
      id: 1
    }
  }
}
</script>

<style scoped></style>
