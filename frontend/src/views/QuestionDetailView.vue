<template>
  <div class="container">
    <div class="ma-2 ml-0">
      <h1>{{ question.title }}</h1>
    </div>
    <div class="post-metadata">
      <div class="text-start">
        <v-chip v-if="isOpen" color="green">回答受付中</v-chip>
        <v-chip v-else color="red">回答締め切り</v-chip>
        <v-chip variant="text" color="grey">{{ answers.length }}件の回答</v-chip>
      </div>
      <div class="text-end">
        <v-chip variant="text" color="grey">{{ user.name }}</v-chip
        ><v-chip variant="text" color="grey"
          >投稿日:{{ question.createdAt.toLocaleDateString() }}</v-chip
        >
      </div>
    </div>
    <div class="post-metadata">
      <div class="text-start">
        <div class="tag-container">
          <div v-for="tag in tags" :key="tag.id">
            <div class="tag">
              <v-chip>{{ tag.name }}</v-chip>
            </div>
          </div>
        </div>
      </div>
    </div>
    <v-divider :thickness="1"></v-divider>
    <v-card class="full-screen-card">
      <MdPreview :editorId="editorId" :modelValue="question.content" />
      <v-card-actions class="d-flex justify-space-between">
        <div class="d-flex justify-space-between text-end">
          <p>{{ question.userID }}|</p>
          <p>{{ question.createdAt.toLocaleDateString() }}</p>
        </div>
        <div>
          <v-btn @click="showModal">編集</v-btn>
          <v-btn
            density="compact"
            icon="mdi-thumb-up"
            color="green"
            @click="incrementScore(question)"
          ></v-btn>
          <v-chip class="mx-4" color="blue-grey lighten-2" text-color="white">{{
            question.score
          }}</v-chip>
          <v-btn
            density="compact"
            icon="mdi-thumb-down"
            color="red"
            @click="decrementScore(question)"
          ></v-btn>
        </div>
      </v-card-actions>
    </v-card>
    <v-dialog v-model="isVisible">
          <v-card>
            <v-card-title>
              <span class="headline">回答を編集する</span>
            </v-card-title>
            <v-card-text>
              <MdEditor v-model="question.content" :language="language" />
            </v-card-text>
            <v-card-actions>
              <v-btn color="black" append-icon="mdi-close" @click="hideModal"> 閉じる </v-btn>
              <v-btn
                color="green"
                append-icon="mdi-send"
                :disabled="!canSubmitNewContent(question.content)"
                @click="submitEditedData"
              >
                編集内容を保存する
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
    <v-divider :thickness="2"></v-divider>
    <div class="ma-2 ml-0">
      <h2>{{ answers.length }}件の回答</h2>
    </div>
    <div class="answers">
      <v-row no-gutters v-for="answer in answers" :key="answer.id">
        <v-card class="full-screen-card my-4">
          <MdPreview :editorId="editorId" :modelValue="answer.content" />
          <v-card-actions class="d-flex justify-space-between">
            <div class="d-flex justify-space-between text-end">
              <p>{{ answer.userID }}|</p>
              <p>{{ answer.createdAt.toLocaleDateString() }}</p>
            </div>
            <div>
              <v-btn @click="showModal">編集</v-btn>
              <v-btn
                density="compact"
                icon="mdi-thumb-up"
                color="green"
                @click="incrementScore(answer)"
              ></v-btn>
              <v-chip class="mx-4" color="blue-grey lighten-2" text-color="white">{{
                answer.score
              }}</v-chip>
              <v-btn
                density="compact"
                icon="mdi-thumb-down"
                color="red"
                @click="decrementScore(answer)"
              ></v-btn>
            </div>
          </v-card-actions>
        </v-card>
        <v-dialog v-model="isVisible">
          <v-card>
            <v-card-title>
              <span class="headline">回答を編集する</span>
            </v-card-title>
            <v-card-text>
              <MdEditor v-model="answer.content" :language="language" />
            </v-card-text>
            <v-card-actions>
              <v-btn color="black" append-icon="mdi-close" @click="hideModal"> 閉じる </v-btn>
              <v-btn
                color="green"
                append-icon="mdi-send"
                :disabled="!canSubmitNewContent(answer.content)"
                @click="submitEditedData"
              >
                編集内容を保存する
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-divider :thickness="1"></v-divider>
      </v-row>
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
              <v-btn
                rounded="xl"
                color="green"
                append-icon="mdi-send"
                v-bind="props"
                @click="submitNewAnswer"
                :disabled="!canSubmitNewContent(newAnswerContent)"
                >回答を送信</v-btn
              ></span
            > </template
          ><span>回答が入力されていないません</span>
        </v-tooltip>
      </div>
    </div>
  </div>
</template>

<script>
import { MdEditor, MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
export default {
  components: {
    MdEditor,
    MdPreview
  },
  data() {
    return {
      editorId: 'preview-only',
      newAnswerContent: '',
      language: 'en-US',
      isVisible: false,
      user: {
        id: '1',
        name: 'masky',
        iconURL: 'https://q.trap.jp/api/v3/public/icon/username',
        userType: 'trap'
      },
      tags: [
        { id: '1', name: 'Tag1' },
        { id: '2', name: 'Tag2' },
        { id: '3', name: 'Tag3' }
      ],
      answers: [
        {
          id: '1',
          userID: '1',
          questionID: '1',
          content: '## 回答内容1',
          createdAt: new Date(),
          score: 0
        },
        {
          id: '2',
          userID: '1',
          questionID: '1',
          content: '## 回答内容2',
          createdAt: new Date(),
          score: 2
        }
      ],
      question: {
        id: '1',
        userID: '1',
        title: '質問タイトル',
        content: '## 質問内容がここに入力される',
        createdAt: new Date(),
        tags: [],
        answers: [],
        status: 'open',
        score: 0
      }
    }
  },
  methods: {
    submitNewAnswer() {
      alert('回答を送信しました')
      this.newAnswerContent = ''
    },
    incrementScore(item) {
      item.score++
    },
    decrementScore(item) {
      item.score--
    },
    showModal() {
      this.isVisible = true
    },
    hideModal() {
      this.isVisible = false
    },
    submitEditedData() {
      alert('編集内容を保存しました')
      this.hideModal()
    },
    canSubmitNewContent(content) {
      return content.length > 0 // 本当は読み込んだ時点のデータとの差があったらtrueにしたい。一旦これで，別ブランチで実装する
    }
  },
  computed: {
    isOpen() {
      return this.question.status === 'open'
    }
  }
}
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
</style>
