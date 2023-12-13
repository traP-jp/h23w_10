<template>
  <div>
    <header>
      <div class="wrapper">
        <nav>
          <RouterLink to="/">Home</RouterLink> |
          <RouterLink :to="'/questions/' + id">Question</RouterLink>
        </nav>
      </div>
    </header>
    <RouterView />
  </div>
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
