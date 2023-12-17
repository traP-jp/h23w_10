<template>
  <v-chip
    :to="`/questions?page=1&tag=${tag.id}`"
  >
    {{ props.tag.name }}
  </v-chip>
</template>

<script setup lang="ts">
import type { Tag } from '@/lib/api/tags'

export interface Props {
  tag: Tag
}
const props = defineProps<Props>()
</script>
