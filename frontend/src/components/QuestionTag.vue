<script setup lang="ts">
import { type Tag } from '@/lib/api'

export interface Props {
  tag: Tag
}
const props = defineProps<Props>()
</script>

<template>
  <v-chip>
    {{ props.tag.name }}
  </v-chip>
</template>
