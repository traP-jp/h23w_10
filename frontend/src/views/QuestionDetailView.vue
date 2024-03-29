<template>
  <div class="container" v-if="question !== null">
    <div class="ma-2 ml-0">
      <h1>{{ question.title }}</h1>
    </div>
    <div class="post-metadata">
      <div class="text-start">
        <question-status :status="question.status" />
        <v-chip variant="text" color="grey">{{ answers.length }}件の回答</v-chip>
      </div>
      <div class="text-end">
        <v-chip variant="text" color="grey">{{ question.user.name }}</v-chip>
        <v-chip variant="text" color="grey">投稿日:{{
          question.created_at ? parseDate(question.created_at).toLocaleDateString() : ''
        }}</v-chip>
      </div>
    </div>
    <div class="post-metadata">
      <div class="text-start">
        <div class="tag-container">
          <question-tag v-for="tag in question.tags" :key="tag.id" :tag="tag" />
        </div>
      </div>
    </div>
    <v-divider :thickness="1"></v-divider>
    <DetailCard :editorId="editorId" :content="question.content" :user="question.user"
      :createdAt="parseDate(question.created_at)" :showModal="showModal" :isQuestion="true"
      :isQuestionResolved="isQuestionResolved" @update:isQuestionResolved="isQuestionResolved = $event" />
    <v-divider :thickness="2"></v-divider>
    <div class="ma-2 ml-0">
      <h2>{{ answers.length }}件の回答</h2>
    </div>
    <div class="answers">
      <DetailCard v-for="answer in answers" :key="answer.id" :editorId="editorId" :content="answer.content"
        :user="answer.user" :createdAt="parseDate(answer.created_at)" :showModal="showModal" :isQuestion="false"
        :isQuestionResolved="isQuestionResolved" @update:isQuestionResolved="isQuestionResolved = $event" />
    </div>
    <div class="d-flex justify-center my-4"
      v-if="question.status === 'open' && loginUser && question.user.id === loginUser.id">
      <v-btn color="green" rounded="xl" @click="changeQuestionStatus">この質問を解決済みにする</v-btn>
    </div>
    <v-divider :thickness="2"></v-divider>
    <div class="ma-2 ml-0">
      <h2>回答を投稿する</h2>
    </div>
    <div class="my-2 mx-5">
      <MdEditor v-model="newAnswerContent" :language="language" />
      <div class="text-end mt-5">
        <v-tooltip location="bottom" :disabled="canSubmitNewContent(newAnswerContent)">
          <template v-slot:activator="{ props }">
            <span v-bind="props">
              <v-btn rounded="xl" color="green" append-icon="mdi-send" v-bind="props" @click="submitNewAnswer"
                :disabled="!canSubmitNewContent(newAnswerContent)">回答を送信</v-btn></span>
          </template><span>回答が入力されていません</span>
        </v-tooltip>
      </div>
    </div>
    <v-dialog v-model="isVisible">
      <v-card>
        <v-card-title>
          <span class="headline">回答を編集する</span>
        </v-card-title>
        <v-card-text>
          <v-text-field
          label="タイトル"
          placeholder="質問のタイトルを入力"
          v-model="newTitle"
          required
          class="ma-3"
        ></v-text-field>
          <MdEditor v-model="modalContent" :language="language" />
        </v-card-text>
        <div class="my-3" style="display: flex; justify-content: flex-end">
          <v-btn color="black" rounded="xl" append-icon="mdi-close" class="mr-1" @click="hideModal">
            閉じる
          </v-btn>
          <v-btn color="green" rounded="xl" append-icon="mdi-send" :disabled="!canSubmitNewContent(modalContent)"
            @click="submitEditedData" class="ml-3 mr-6">
            編集内容を保存する
          </v-btn>
        </div>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { ref, onMounted, inject, type Ref } from 'vue'
import { type User } from '@/lib/api/users'
import { getQuestion, putQuestion, type Question } from '@/lib/api/questions'
import { postAnswer, type Answer } from '@/lib/api/answers'
import { type Tag } from '@/lib/api/tags'
import QuestionStatus from '@/components/QuestionStatus.vue'
import QuestionTag from '@/components/QuestionTag.vue'
import DetailCard from '@/components/DetailCard.vue'
import { useRoute } from 'vue-router'
import { parseDate } from '@/lib/parseDate'

const editorId = 'preview-only'
const modalContent = ref('')
const newTitle = ref('')
const newAnswerContent = ref('')
const language = 'en-US'
const isVisible = ref(false)
const isQuestion = ref(true)
const tags = ref<Tag[]>([])
const isQuestionResolved = ref(false)
const answers = ref<Answer[]>([])
const question = ref<Question | null>(null)
const loginUser = inject<Ref<User | null>>('loginUser')

const submitNewAnswer = async () => {
  if (!question.value) throw new Error('question is null')
  if (!loginUser?.value) throw new Error('loginUser is null')
  await postAnswer({
    content: newAnswerContent.value,
    question_id: question.value.id,
    user_id: loginUser.value.id
  })
  newAnswerContent.value = ''
}

const showModal = (content: string, isQuestionValue: boolean) => {
  modalContent.value = content
  isQuestion.value = isQuestionValue
  isVisible.value = true
}

const hideModal = () => {
  isVisible.value = false
}

const submitEditedData = async () => {
  if (isQuestion.value) {
  if (!question.value) throw new Error('question is null')
    await putQuestion({
      id: question.value.id,
      content: modalContent.value,
      title: newTitle.value,
      status: question.value.status,
      tags: question.value.tags?.map((tag) => ({ id: tag.id })) ?? []
    })
  } else {
    // Todo: APIに置き換える
    console.log('回答が編集されました')
  }
  hideModal()
}

const canSubmitNewContent = (content: string) => {
  return content.length > 0
}

const changeQuestionStatus = async () => {
  if (!question.value) throw new Error('question is null')
  await putQuestion({
    id: question.value.id,
    content: question.value.content,
    title: question.value.title,
    status: 'closed',
    tags: question.value.tags?.map((tag) => ({ id: tag.id })) ?? []
  })
  question.value.status = 'closed'
}

onMounted(() => {
  const route = useRoute()
  const id: string = route.params.id as string
  getQuestion({ id })
    .then((response) => {
      question.value = response
      answers.value = question.value.answers ?? []
      tags.value = question.value.tags ?? []
      newTitle.value = question.value.title
      if (question.value.status === 'closed') {
        isQuestionResolved.value = true
      }
    })
    .catch((error) => {
      console.error(error)
    })
})
</script>
<style scoped>
.container {
  width: 80%;
  margin: auto;
  margin-bottom: 50px;
  flex-direction: column;
  align-items: center;
}

.post-metadata {
  display: flex;
  justify-content: space-between;
  margin: 10px;
  margin-left: 0px;
}

.tag-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.full-screen-card {
  width: 100%;
  margin: 20px auto;
}

.flex-space-between {
  display: flex;
  justify-content: space-between;
}
</style>
