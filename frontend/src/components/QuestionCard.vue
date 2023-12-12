<template>
  <v-card>
    <v-card-title>
      <div>
        {{ props.question.title }}
      </div>
    </v-card-title>
    <v-card-text>
      <div>
        {{ props.question.content }}
      </div>
    </v-card-text>
    <v-card-actions>
      <div class="info-container">
        <div class="tags-container">
          <question-tag v-for="tag in props.question.tags" :key="tag.id" :tag="tag" />
        </div>
        <div class="status-container">
          <div class="status">
            <question-status :status="props.question.status" />
            {{ props.question.answers.length }}件の回答
          </div>
          <div class="status">
            {{ props.question.userId }}
            {{ diffHuman(props.question.createdAt).diff }}
          </div>
        </div>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { type Question } from '@/lib/api'
import { diffHuman } from '@/lib/format'
import QuestionTag from './QuestionTag.vue'
import QuestionStatus from './QuestionStatus.vue'

export interface Props {
  question: Question
}
const props = defineProps<Props>()
</script>

<style scoped>
.info-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.tags-container {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.status-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.status {
  display: flex;
  align-items: center;
  gap: 4px;
}
</style>
