<template>
  <div class="container">
    <v-autocomplete
      label="filter by tag"
      :items="tags"
      v-model="req.tag"
      item-title="name"
      item-value="id"
      clearable
      @change="req.page = 1"
    />
    <div v-if="!loading" class="questions">
      <QuestionCard v-for="question in data.questions" :key="question.id" :question="question" />
      <!-- todo: 検索結果0件の時のfallback -->
    </div>
    <!-- todo: loading時のlayout shift対策 -->
    <div v-else>Loading...</div>
    <v-pagination v-model="req.page" :length="Math.ceil(data.total / limit)" :total-visible="9" />
  </div>
</template>

<script setup lang="ts">
import QuestionCard from '@/components/QuestionCard.vue'
import { getQuestions, type GetQuestionsResponse, type QuestionStatus } from '@/lib/api/questions'
import { getTags, type Tag } from '@/lib/api/tags'
import { watch } from 'vue'
import { reactive } from 'vue'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const limit = 10
const route = useRoute()
const tags = await getTags()

type Req = {
  page: number
  tag?: Tag['id']
  status?: QuestionStatus
}

const loading = ref(false)

const data = ref<GetQuestionsResponse>({
  questions: [],
  total: 0
})
const req = reactive<Req>({
  page: typeof route.query.page === 'string' ? parseInt(route.query.page) : 1,
  tag: typeof route.query.tag === 'string' ? route.query.tag : undefined,
  status:
    typeof route.query.status === 'string' ? (route.query.status as QuestionStatus) : undefined
})

watch(
  req,
  async (req) => {
    loading.value = true
    try {
      const res = await getQuestions({
        limit,
        offset: (req.page - 1) * limit,
        tag: req.tag,
        status: req.status
      })
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
