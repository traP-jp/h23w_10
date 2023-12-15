<template>
  <v-card class="full-screen-card">
    <MdPreview :editorId="props.editorId" :modelValue="props.content" />
    <v-card-actions class="d-flex justify-space-between">
      <div class="d-flex justify-space-between text-end">
        <p>{{ props.userId }}|</p>
        <p>{{ props.createdAt ? props.createdAt.toLocaleDateString() : '' }}</p>
      </div>
      <div>
        <v-btn @click="showModal(props.content, isQuestion)">編集</v-btn>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { MdPreview } from 'md-editor-v3'

export interface Props {
  editorId: string
  content: string
  userId: string
  createdAt: Date | undefined
  showModal: Function
  isQuestion: boolean
}

const props = defineProps<Props>()
</script>
