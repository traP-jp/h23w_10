<template>
  <div class="container">
    <div class="heading">
      <h1>{{ question.title }}</h1>
    </div>
    <div class="post-metadata">
      <div class="left-align">
        <v-chip color="green">回答受付中</v-chip>
        <v-chip variant="text" color="grey">{{ answers.length }}件の回答</v-chip>
      </div>
      <div class="right-align">
        <v-chip variant="text" color="grey">{{ user.name }}</v-chip
        ><v-chip variant="text" color="grey"
          >投稿日:{{ question.createdAt.toLocaleDateString() }}</v-chip
        >
      </div>
    </div>
    <div class="post-metadata">
      <div class="left-align">
        <div class="tag-container">
          <v-for v-for="tag in tags" :key="tag.id">
            <div class="tag">
              <v-chip>{{ tag.name }}</v-chip>
            </div>
          </v-for>
        </div>
      </div>
    </div>
    <v-divider :thickness="1"></v-divider>
    <MdPreview :editorId="editorId" :modelValue="question.content" />
    <v-divider :thickness="2"></v-divider>
    <div class="heading">
      <h2>{{ answers.length }}件の回答</h2>
    </div>
    <div class="answers">
      <v-row no-gutters v-for="answer in answers" :key="answer.id"
        ><MdPreview :editorId="editorId" :modelValue="answer.content" />
        <v-divider :thickness="1"></v-divider
      ></v-row>
    </div>
    <v-divider :thickness="2"></v-divider>
    <div class="heading">
      <h2>回答を投稿する</h2>
    </div>
    <div class="sendNewAnswer">
      <MdEditor v-model="newAnswerContent" :language="language" />
      <div class="sendBtn">
        <v-tooltip location="bottom" :disabled="isFilled">
          <template v-slot:activator="{ props }">
            <span v-bind="props">
              <v-btn
                rounded="xl"
                color="green"
                append-icon="mdi-send"
                v-bind="props"
                @click="submitNewAnswer"
                :disabled="!newAnswerContent"
                >回答を送信</v-btn
              ></span
            > </template
          ><span>回答を入力してください</span>
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
          createdAt: new Date()
        },
        {
          id: '2',
          userID: '1',
          questionID: '1',
          content: '## 回答内容2',
          createdAt: new Date()
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
        status: 'open'
      }
    }
  },
  methods: {
    submitNewAnswer() {
      alert('回答を送信しました')
      this.newAnswerContent = ''
    }
  },
  computed: {
    isFilled() {
      return this.newAnswerContent.length > 0
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
.heading {
  margin: 10px;
  margin-left: 0px;
}
.post-metadata {
  display: flex;
  justify-content: space-between;
  margin: 10px;
  margin-left: 0px;
}
.left-align {
  text-align: left;
}
.right-align {
  text-align: right;
}
.tag-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
.sendNewAnswer {
  margin: 10px 20px;
}
.sendBtn {
  text-align: right;
  margin-top: 10px;
}
</style>
