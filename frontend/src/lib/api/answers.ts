import { BASE, useMock } from '.'
import type { Question } from './questions'
import type { User } from './users'

export type Answer = {
  id: string
  user: User
  question_id: Question['id']
  content: string
  created_at: Date
}

export type PostAnswerRequest = Omit<Answer, 'id' | 'created_at'>
export type PostAnswerResponse = Answer

export const postAnswer = async (req: PostAnswerRequest): Promise<PostAnswerResponse> => {
  if (useMock) {
    const { postAnswerMock } = await import('./mock')
    return postAnswerMock(req)
  }

  const res = await fetch(`${BASE}/questions/${req.question_id}/answers`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(req)
  })
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: PostAnswerResponse = await res.json()
  return json
}

export type PutAnswerRequest = Omit<Answer, 'user' | 'created_at'>
export type PutAnswerResponse = Answer

export const putAnswer = async (req: PutAnswerRequest): Promise<PutAnswerResponse> => {
  const res = await fetch(`${BASE}/questions/${req.question_id}/answers/${req.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(req)
  })
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: Answer = await res.json()
  return json
}