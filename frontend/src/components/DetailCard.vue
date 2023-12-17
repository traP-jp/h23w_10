<template>
  <v-card class="full-screen-card">
    <MdPreview :editorId="props.editorId" :modelValue="props.content" />
    <v-card-actions class="d-flex justify-space-between">
      <div class="d-flex">
        <v-btn :to="`$/users/${props.user}`">
          <v-avatar size="small" class="avatar">
            <div class="mr--1"></div>
            <v-img :src="user?.icon_url ?? ''" aspect-ratio="1" contain class="avatar-image"></v-img>
          </v-avatar>
          <div style="margin-left: -8px">
            <v-chip variant="text" color="grey">{{ user?.name }}</v-chip>
          </div>
        </v-btn>
        <div style="margin-top: 3px">
          <v-chip variant="text" color="grey">
            {{ diffHuman(props.createdAt ?? new Date()).diff }}</v-chip>
        </div>
      </div>

      <div v-if="loginUser?.id === props.user.id">
        <v-btn @click="showModal(props.content, isQuestion)">編集</v-btn>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { MdPreview } from 'md-editor-v3'
import { diffHuman } from '@/lib/format'
import { type User } from '@/lib/api/users'
import { inject, type Ref } from 'vue'

const loginUser = inject<Ref<User | null>>('loginUser')
export interface Props {
  editorId: string
  content: string
  user: User
  createdAt: Date | undefined
  showModal: Function
  isQuestion: boolean
}

const props = defineProps<Props>()
</script>
