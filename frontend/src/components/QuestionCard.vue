<template>
  <v-card :class="$style.container" :href="`/questions/${props.question.id}`" tag="a">
    <v-card-title>
      <div>
        {{ props.question.title }}
      </div>
    </v-card-title>
    <v-card-text>
      <div>
        {{ props.question.content }}
      </div>
    </v-card-text>
    <v-card-actions>
      <div :class="$style.infoContainer">
        <div :class="$style.tagsContainer">
          <question-tag v-for="tag in props.question.tags" :key="tag.id" :tag="tag" />
        </div>
        <div :class="$style.statusContainer">
          <div :class="$style.status">
            <question-status :status="props.question.status" />
          </div>
          <div :class="`${$style.status} text-caption`">
            {{ props.question.user.name }}
            <v-tooltip :text="diffHuman(parseDate(props.question.created_at)).localeString" location="top">
              <template v-slot:activator="{ props: tooltipProps }">
                <span v-bind="tooltipProps">
                  {{ diffHuman(parseDate(props.question.created_at)).diff }}
                </span>
              </template>
            </v-tooltip>
          </div>
        </div>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { diffHuman } from '@/lib/format'
import QuestionTag from './QuestionTag.vue'
import QuestionStatus from './QuestionStatus.vue'
import type { Question } from '@/lib/api/questions'
import { parseDate } from '@/lib/parseDate'

export interface Props {
  question: Question
}
const props = defineProps<Props>()
</script>

<style module>
.container {
  width: 100%;
}
.infoContainer {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.tagsContainer {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.statusContainer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.status {
  display: flex;
  align-items: center;
  gap: 4px;
  flex-shrink: 0;
}
</style>
