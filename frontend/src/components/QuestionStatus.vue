<template>
  <v-chip :color="props.status === 'open' ? 'green' : 'red'">
    {{ props.status === 'open' ? '回答受付中' : '解決済み' }}
  </v-chip>
</template>

<script setup lang="ts">
import type { QuestionStatus } from '@/lib/api/questions'

export interface Props {
  status: QuestionStatus
}
const props = defineProps<Props>()
</script>
