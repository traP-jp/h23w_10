<template>
  <div class="container">
    <div v-if="!loading" class="questions">
      <QuestionCard v-for="question in data.questions" :key="question.id" :question="question" />
    </div>
    <div v-else>Loading...</div>
    <v-pagination v-model="page" :length="Math.ceil(data.total / limit)" :total-visible="9" />
  </div>
</template>

<script setup lang="ts">
import QuestionCard from '@/components/QuestionCard.vue'
import { getQuestions, type GetQuestionsResponse } from '@/lib/api/questions'
import { watch } from 'vue'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const limit = 5
// get path param
const route = useRoute()

// use ref to make page reactive
const page = ref(typeof route.query.page === 'string' ? parseInt(route.query.page) : 1)
const data = ref<GetQuestionsResponse>({
  questions: [],
  total: 0
})
const loading = ref(false)

watch(
  page,
  async (page) => {
    loading.value = true
    try {
      const res = await getQuestions({ limit, offset: (page - 1) * limit })
      console.log(res)
      data.value = res
    } catch (err) {
      console.error(err)
    } finally {
      loading.value = false
    }
  },
  {
    immediate: true
  }
)
</script>

<style scoped>
.container {
  width: 100%;
  max-width: 1000px;
}
.questions {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  gap: 8px;
}
</style>
