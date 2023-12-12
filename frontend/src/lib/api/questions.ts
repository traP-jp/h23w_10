import { BASE } from '.'
import type { Answer } from './answers'
import type { Tag } from './tags'
import type { User } from './users'

export type QuestionStatus = 'open' | 'closed'

export type Question = {
  id: string
  userId: User['id']
  title: string
  content: string
  createdAt: Date
  tags: Tag[]
  answers: Answer[]
  status: QuestionStatus
}

export type GetQuestionsRequest = {
  limit: number
  offset: number
  tag?: Tag['id']
  status?: QuestionStatus
}
export type GetQuestionsResponse = {
  questions: Question[]
  total: number
}

export const getQuestions = async (req: GetQuestionsRequest): Promise<GetQuestionsResponse> => {
  if (import.meta.env.DEV) {
    const { getQuestionsMock } = await import('./mock')
    return getQuestionsMock(req)
  }

  const params = new URLSearchParams({
    limit: String(req.limit),
    offset: String(req.offset)
  })
  if (req.tag) {
    params.set('tag', req.tag)
  }
  if (req.status) {
    params.set('status', req.status)
  }
  const res = await fetch(`${BASE}/questions?${params}`)
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: GetQuestionsResponse = await res.json()
  return json
}

export type GetQuestionRequest = {
  id: Question['id']
}
export type GetQuestionResponse = Question

export const getQuestion = async (req: GetQuestionRequest): Promise<GetQuestionResponse> => {
  if (import.meta.env.DEV) {
    const { getQuestionMock } = await import('./mock')
    return getQuestionMock(req)
  }

  const res = await fetch(`${BASE}/questions/${req.id}`)
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: GetQuestionResponse = await res.json()
  return json
}

export type PostQuestionRequest = {
  userId: User['id']
  title: string
  content: string
  tags: Omit<Tag, 'name'>[]
  status?: QuestionStatus
}
export type PostQuestionResponse = Question

export const postQuestion = async (req: PostQuestionRequest): Promise<PostQuestionResponse> => {
  if (import.meta.env.DEV) {
    const { postQuestionMock } = await import('./mock')
    return postQuestionMock(req)
  }

  req.status = req.status ?? 'open'
  const res = await fetch(`${BASE}/questions`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(req)
  })
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: PostQuestionResponse = await res.json()
  return json
}

// TODO: PUT /questions/:id