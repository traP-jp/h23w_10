<template>
  <div class="container">
    <div class="ma-2 ml-0"><h1>新しい質問をする</h1></div>
    <v-card>
      <v-form>
        <v-text-field
          label="タイトル"
          placeholder="質問のタイトルを入力"
          v-model="form.title"
          required
          class="ma-3"
        ></v-text-field>
        <div class="mx-3"><MdEditor v-model="form.content" :language="language" /></div>
        <v-autocomplete
          label="質問のタグを選択"
          :items="tags"
          item-title="name"
          item-value="id"
          class="mx-3 mt-6"
          v-model="form.tags"
          clearable
          multiple
          chips
          no-data-text="タグが見つかりません"
          @create-new-item="createNewTag"
        >
          <template v-slot:append-item>
            <v-list-item @click="createNewTag">
              <v-list-item-title> 新しいタグを作成 </v-list-item-title>
            </v-list-item>
          </template>
        </v-autocomplete>
        <div class="text-end mb-3 mr-3">
          <v-tooltip location="bottom" :disabled="canPostQuestion">
            <template v-slot:activator="{ props }">
              <span v-bind="props">
                <v-btn
                  rounded="xl"
                  color="green"
                  append-icon="mdi-send"
                  v-bind="props"
                  :disabled="!canPostQuestion"
                  @click="postNewQuestion"
                  >質問を送信</v-btn
                ></span
              > </template
            ><span>回答が入力されていません</span>
          </v-tooltip>
        </div></v-form
      ></v-card
    >
    <div class="createTagModal">
      <v-dialog v-model="isVisible"
        ><v-card>
          <v-form>
            <v-text-field
              label="タグ"
              placeholder="新しいタグの名前を入力"
              v-model="tagName"
              required
              class="ma-3"
            ></v-text-field>
            <div class="mb-3 mr-3" style="display: flex; justify-content: flex-end">
              <div class="mr-1">
                <v-btn
                  color="black"
                  rounded="xl"
                  append-icon="mdi-close"
                  @click="hideModal"
                >
                  閉じる
                </v-btn>
              </div>

              <v-tooltip location="bottom" :disabled="canPostTag">
                <template v-slot:activator="{ props }">
                  <span v-bind="props">
                    <v-btn
                      rounded="xl"
                      color="green"
                      append-icon="mdi-send"
                      v-bind="props"
                      :disabled="!canPostTag"
                      @click="postNewTag"
                      >送信</v-btn
                    ></span
                  > </template
                ><span>タグが入力されていません</span>
              </v-tooltip>
            </div></v-form
          ></v-card
        ></v-dialog
      >
    </div>
  </div>
</template>
<script setup lang="ts">
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { ref, computed, reactive } from 'vue'
import { postQuestion } from '@/lib/api/questions'
import { getTags, postTag, type Tag } from '@/lib/api/tags'
import { useRouter } from 'vue-router'

const tagName = ref('')

const language = 'en-US'
const form = reactive<{ title: string; content: string; tags: Tag[] }>({
  title: '',
  content: '# 質問内容を入力',
  tags: []
})
const tags = await getTags()
const router = useRouter()

const canPostQuestion = computed(() => {
  return form.title.length > 0 && form.content.length > 0
})

const canPostTag = computed(() => {
  return tagName.value.length > 0
})

const isVisible = ref(false)
const createNewTag = () => {
  isVisible.value = true
}
const hideModal = () => {
  tagName.value = ''
  isVisible.value = false
}

const postNewQuestion = async () => {
  const selectedTagIds: Omit<Tag, 'name'>[] = form.tags.map((tag) => ({ id: tag.id }))
  try {
    const res = await postQuestion({
      userId: 'masky', //TODO： ログインユーザーのIDを取得する
      title: form.title,
      content: form.content,
      tags: selectedTagIds,
      status: 'open'
    })
    router.push(`/questions/${res.id}`)
  } catch (err) {
    console.log(err)
  }
}

const postNewTag = async () => {
  try {
    const res = await postTag({
      name: tagName.value
    })
    tags.push(res)
    form.tags.push(res)
    tagName.value = ''
    isVisible.value = false
  } catch (err) {
    console.log(err)
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
</style>
