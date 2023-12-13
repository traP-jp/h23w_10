<template>
  <div :class="$style.container">
    <div :class="$style.header">
      <h1>staQ overflow</h1>
      <v-btn href="/questions/new" color="primary" size="x-large" append-icon="mdi-send"
        >質問する</v-btn
      >
    </div>
    <p>traP部員に質問しよう</p>
    <h2>Latest Questions</h2>
    <div v-if="!loading" :class="$style.questions">
      <QuestionCard v-for="question in data.questions" :key="question.id" :question="question" />
      <v-btn href="/questions" append-icon="mdi-arrow-right">全ての質問を見る</v-btn>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script setup lang="ts">
import QuestionCard from '@/components/QuestionCard.vue'
import { getQuestions, type GetQuestionsResponse } from '@/lib/api/questions'
import { ref } from 'vue'

const loading = ref(true)
const data = ref<GetQuestionsResponse>({
  questions: [],
  total: 0
})
getQuestions({
  limit: 4,
  offset: 0
})
  .then((res) => {
    data.value = res
  })
  .catch((err) => {
    console.error(err)
  })
  .finally(() => {
    loading.value = false
  })
</script>

<style module>
.container {
  width: 80%;
  margin: auto;
  padding-top: 24px;
  padding-bottom: 50px;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.questions {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  gap: 8px;
}
</style>
