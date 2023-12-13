<template>
  <div :class="$style.container">
    <div :class="$style.filterContainer">
      <v-select
        label="filter by status"
        clearable
        :items="[
          { title: '解決済み', value: 'closed' },
          { title: '回答受付中', value: 'open' }
        ]"
        v-model="req.status"
      />
      <v-autocomplete
        label="filter by tag"
        :items="tags"
        v-model="req.tag"
        item-title="name"
        item-value="id"
        clearable
      />
    </div>
    <div
      v-if="data.questions"
      :class="[
        $style.questions,
        {
          [$style.loading]: loading
        }
      ]"
    >
      <QuestionCard v-for="question in data.questions" :key="question.id" :question="question" />
      <div v-if="data.questions.length === 0">条件に一致する質問が見つかりませんでした</div>
    </div>
    <div v-else>Loading...</div>
    <v-pagination v-model="req.page" :length="Math.ceil(data.total / limit)" :total-visible="9" />
  </div>
</template>

<script setup lang="ts">
import QuestionCard from '@/components/QuestionCard.vue'
import { getQuestions, type GetQuestionsResponse, type QuestionStatus } from '@/lib/api/questions'
import { getTags, type Tag } from '@/lib/api/tags'
import { watch, reactive, ref } from 'vue'
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

const data = ref<{
  questions: GetQuestionsResponse['questions'] | undefined
  total: GetQuestionsResponse['total']
}>({
  questions: undefined,
  total: 0
})
const req = reactive<Req>({
  page: typeof route.query.page === 'string' ? parseInt(route.query.page) : 1,
  tag: typeof route.query.tag === 'string' ? route.query.tag : undefined,
  status:
    typeof route.query.status === 'string' ? (route.query.status as QuestionStatus) : undefined
})

// queryが変更されたらreqを更新する (ブラウザバック等の対応)
watch(
  () => route.query,
  (query) => {
    req.page = typeof query.page === 'string' ? parseInt(query.page) : 1
    req.tag = typeof query.tag === 'string' ? query.tag : undefined
    req.status = typeof query.status === 'string' ? (query.status as QuestionStatus) : undefined
  },
  {
    immediate: true
  }
)

// tagが変更されたらpageを1にする
watch(
  () => req.tag,
  () => {
    req.page = 1
  },
  {
    immediate: true
  }
)

// フィルター, ページが変更されたらAPIを叩く
watch(
  req,
  async (req) => {
    loading.value = true

    const newQuery: Partial<Record<keyof Req, string>> = {
      page: req.page.toString()
    }
    if (req.tag) {
      newQuery.tag = req.tag
    }
    if (req.status) {
      newQuery.status = req.status
    }
    history.pushState(null, '', `${route.path}?${new URLSearchParams(newQuery).toString()}`)

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

<style module>
.container {
  width: 80%;
  margin: auto;
  padding-top: 24px;
  padding-bottom: 50px;
}
.loading {
  opacity: 0.5;
}
.filterContainer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
}
@media screen and (max-width: 600px) {
  .filterContainer {
    flex-direction: column;
    align-items: stretch;
    gap: 0;
  }
}
.questions {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  gap: 8px;
}
</style>
