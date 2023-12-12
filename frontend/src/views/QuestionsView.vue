<script setup lang="ts">
import QuestionCard from '@/components/QuestionCard.vue'
import { type Question } from '@/lib/api'

const data: Question[] = [
  {
    id: 'test',
    userId: 'test',
    title: 'test',
    content: 'test',
    createdAt: new Date(),
    tags: [
      {
        id: 'test',
        name: 'Python'
      },
      {
        id: 'test',
        name: 'Programming'
      },
      {
        id: 'test',
        name: 'Python'
      },
      {
        id: 'test',
        name: 'Python'
      }
    ],
    answers: [],
    status: 'open'
  },
  {
    id: 'test',
    userId: 'test',
    title: 'test',
    content: 'test',
    createdAt: new Date(),
    tags: [],
    answers: [],
    status: 'open'
  }
]
</script>
<template>
  <div>
    <QuestionCard v-for="question in data" :key="question.id" :question="question" />
  </div>
</template>
